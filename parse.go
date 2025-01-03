package ox

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

// Parse parses the args into context.
func Parse(ctx *Context, cmd *Command, args []string) (*Command, []string, error) {
	var v []string
	var s string
	var n int
	var err error
	for len(args) != 0 {
		switch s, n, args = args[0], len(args[0]), args[1:]; {
		case n == 0, n == 1, s[0] != '-':
			// lookup sub command
			var c *Command
			if len(v) == 0 {
				c = cmd.Command(s)
			}
			if c != nil {
				if err := ctx.Populate(c, false, false); err != nil {
					return nil, nil, newCommandError(c.Name, err)
				}
				cmd = c
			} else {
				v = append(v, s)
			}
		case s == "--":
			return cmd, append(v, args...), nil
		case s[1] == '-':
			if args, err = ParseFlagLong(ctx, cmd, s, args); err != nil {
				if !ctx.Continue(cmd, err) {
					return nil, nil, err
				}
			}
		default:
			if args, err = ParseFlagShort(ctx, cmd, s, args); err != nil {
				if !ctx.Continue(cmd, err) {
					return nil, nil, err
				}
			}
		}
	}
	return cmd, v, nil
}

// ParseFlagLong parses a long flag ('--arg' '--arg v' '--arg k=v' '--arg=' '--arg=v').
func ParseFlagLong(ctx *Context, cmd *Command, s string, args []string) ([]string, error) {
	arg, value, ok := strings.Cut(strings.TrimPrefix(s, "--"), "=")
	g := cmd.Flag(arg, true, false)
	switch {
	case g == nil && ctx.Continue(cmd, newFlagError(arg, ErrUnknownFlag)):
		return args, nil
	case g == nil:
		return nil, newFlagError(arg, ErrUnknownFlag)
	case ok: // --arg=v
	case g.NoArg: // --arg
		var err error
		if value, err = asString[string](g.NoArgDef); err != nil {
			return nil, newFlagError(arg, err)
		}
	case len(args) != 0: // --arg v
		value, args = args[0], args[1:]
	default: // missing argument to --arg
		return nil, newFlagError(arg, ErrMissingArgument)
	}
	if err := ctx.Vars.Set(ctx, g, value, true); err != nil {
		return nil, newFlagError(arg, err)
	}
	return args, nil
}

// ParseFlagShort parses short flags ('-a' '-aaa' '-av' '-a v' '-a=' '-a=v').
func ParseFlagShort(ctx *Context, cmd *Command, s string, args []string) ([]string, error) {
	for v := []rune(s[1:]); len(v) != 0; v = v[1:] {
		arg := string(v[0])
		switch g, n := cmd.Flag(arg, true, true), len(v[1:]); {
		case g == nil && ctx.Continue(cmd, newFlagError(arg, ErrUnknownFlag)):
			return args, nil
		case g == nil:
			return nil, newFlagError(arg, ErrUnknownFlag)
		case g.NoArg: // -a
			var value string
			var err error
			if slices.Index(v, '=') == 1 {
				value, v = string(v[2:]), v[len(v)-1:]
			} else if value, err = asString[string](g.NoArgDef); err != nil {
				return nil, newFlagError(arg, err)
			}
			if err := ctx.Vars.Set(ctx, g, value, true); err != nil {
				return nil, newFlagError(arg, err)
			}
		case n == 0 && len(args) == 0: // missing argument to -a
			return nil, newFlagError(arg, ErrMissingArgument)
		case n != 0: // -a=, -a=v, -av
			if slices.Index(v, '=') == 1 {
				v = v[1:]
			}
			if err := ctx.Vars.Set(ctx, g, string(v[1:]), true); err != nil {
				return nil, newFlagError(arg, err)
			}
			return args, nil
		default: // -a v
			if err := ctx.Vars.Set(ctx, g, args[0], true); err != nil {
				return nil, newFlagError(arg, err)
			}
			return args[1:], nil
		}
	}
	return args, nil
}

// Vars is a map of argument variables.
type Vars map[string]Value

// String satisfies the [fmt.Stringer] interface.
func (vars Vars) String() string {
	var v []string
	for _, k := range slices.Sorted(maps.Keys(vars)) {
		if s, err := vars[k].Get(); err == nil {
			v = append(v, k+":"+s)
		}
	}
	return "[" + strings.Join(v, " ") + "]"
}

// Set sets a variable in the vars.
func (vars Vars) Set(ctx *Context, g *Flag, s string, set bool) error {
	name := g.Name
	if name == "" {
		return ErrInvalidFlagName
	}
	var err error
	v, ok := vars[name]
	if !ok {
		if v, err = g.New(ctx); err != nil {
			return err
		}
		if s, ok := v.(interface{ SetSplit(func(string) []string) }); ok && g.Split != "" {
			var f func(string) []string
			switch r := []rune(g.Split); len(r) {
			case 1:
				f = func(s string) []string {
					return SplitBy(s, r[0])
				}
			default:
				f = func(s string) []string {
					return strings.Split(s, g.Split)
				}
			}
			s.SetSplit(f)
		}
	}
	if err := v.Set(s); err != nil {
		return err
	}
	v.SetSet(set)
	for i, bind := range g.Binds {
		if err := bind.Bind(s); err != nil {
			return fmt.Errorf("bind %d (%s): cannot set %q: %w", i, bind, s, err)
		}
	}
	vars[name] = v
	return nil
}
