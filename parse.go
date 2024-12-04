package ox

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

// Parse parses the command-line arguments into vars.
func Parse(ctx *Context, root *Command, args []string, vars Vars) (*Command, []string, error) {
	if root.Parent != nil {
		return nil, nil, fmt.Errorf("Parse: %w", ErrCanOnlyBeUsedWithRootCommand)
	}
	if err := root.Populate(ctx, false, false, vars); err != nil {
		return nil, nil, newCommandError(root.Name, err)
	}
	if len(args) == 0 {
		return root, nil, nil
	}
	return parse(ctx, root, args, vars)
}

// parse parses the args into m based on the flags on the command.
func parse(ctx *Context, cmd *Command, args []string, vars Vars) (*Command, []string, error) {
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
				if err := c.Populate(ctx, false, false, vars); err != nil {
					return nil, nil, newCommandError(c.Name, err)
				}
				cmd = c
			} else {
				v = append(v, s)
			}
		case s == "--":
			return cmd, append(v, args...), nil
		case s[1] == '-':
			if args, err = parseLong(ctx, cmd, s, args, vars); err != nil {
				return nil, nil, err
			}
		default:
			if args, err = parseShort(ctx, cmd, s, args, vars); err != nil {
				return nil, nil, err
			}
		}
	}
	return cmd, v, nil
}

// parseLong parses a long flag ('--arg' '--arg v' '--arg k=v' '--arg=' '--arg=v').
func parseLong(ctx *Context, cmd *Command, s string, args []string, vars Vars) ([]string, error) {
	arg, value, ok := strings.Cut(strings.TrimPrefix(s, "--"), "=")
	g := cmd.Flag(arg, false)
	switch {
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
	if err := vars.Set(ctx, g, value, true); err != nil {
		return nil, newFlagError(arg, err)
	}
	return args, nil
}

// parseShort parses short flags ('-a' '-aaa' '-av' '-a v' '-a=' '-a=v').
func parseShort(ctx *Context, cmd *Command, s string, args []string, vars Vars) ([]string, error) {
	for v := []rune(s[1:]); len(v) != 0; v = v[1:] {
		arg := string(v[0])
		switch g, n := cmd.Flag(arg, true), len(v[1:]); {
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
			if err := vars.Set(ctx, g, value, true); err != nil {
				return nil, newFlagError(arg, err)
			}
		case n == 0 && len(args) == 0: // missing argument to -a
			return nil, newFlagError(arg, ErrMissingArgument)
		case n != 0: // -a=, -a=v, -av
			if slices.Index(v, '=') == 1 {
				v = v[1:]
			}
			if err := vars.Set(ctx, g, string(v[1:]), true); err != nil {
				return nil, newFlagError(arg, err)
			}
			return args, nil
		default: // -a v
			if err := vars.Set(ctx, g, args[0], true); err != nil {
				return nil, newFlagError(arg, err)
			}
			return args[1:], nil
		}
	}
	return args, nil
}

// Vars is a map of argument variables.
type Vars map[string]Value

// String satisfies the [fmt.Stringer] interfaec.
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
	}
	if err := v.Set(s); err != nil {
		return err
	}
	v.SetSet(set)
	for i, val := range g.Binds {
		if err := val.Bind(s); err != nil {
			return fmt.Errorf("bind %d (%s): cannot set %q: %w", i, val, s, err)
		}
	}
	vars[name] = v
	return nil
}
