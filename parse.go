package ox

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

// Parse parses flags and args into context.
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
	return vars.set(ctx, g, s, set, true)
}

// fill sets a flag from a value that is not from the command line, such as
// a value from the environment, a config file or a default. It assigns the
// value: a slice gets only the elements of val, and a count gets the number
// in val, where the command line adds one for each use.
func (vars Vars) fill(ctx *Context, g *Flag, val any) error {
	strs, split, err := fillValues(g, val)
	if err != nil {
		return err
	}
	delete(vars, g.Name)
	if g.Type == CountT {
		n, err := strconv.ParseUint(strs[0], 10, 16)
		if err != nil {
			return fmt.Errorf("%w: count %q: %w", ErrInvalidValue, strs[0], err)
		}
		for range n {
			if err := vars.set(ctx, g, "", false, true); err != nil {
				return err
			}
		}
		return nil
	}
	for _, s := range strs {
		if err := vars.set(ctx, g, s, false, split); err != nil {
			return err
		}
	}
	return nil
}

// set sets a flag. When split is false, a slice or a map takes s whole.
func (vars Vars) set(ctx *Context, g *Flag, s string, set, split bool) error {
	name := g.Name
	if name == "" {
		return ErrInvalidFlagName
	}
	// instantiate
	v, ok := vars[name]
	if !ok {
		var err error
		if v, err = g.New(ctx); err != nil {
			return err
		}
		setValid(v, g.Validate)
		setSplit(v, g.Split)
	}
	if !split {
		noSplit(v)
	}
	// set
	if err := v.Set(s); err != nil {
		return err
	}
	v.SetSet(set)
	// bind
	for i, bind := range g.Binds {
		setSplit(bind, g.Split)
		if !split {
			noSplit(bind)
		}
		if err := bind.Bind(s); err != nil {
			return fmt.Errorf("bind %d (%s): cannot set %q: %w", i, bind, s, err)
		}
		bind.SetSet(set)
	}
	vars[name] = v
	return nil
}

// setValid sets the valid func on v.
func setValid(v any, valid func(any) error) {
	if valid == nil {
		return
	}
	if z, ok := v.(interface{ SetValid(func(any) error) }); ok {
		z.SetValid(valid)
	}
}

// noSplit turns off the split of a slice or a map value.
func noSplit(v any) {
	if z, ok := v.(interface{ SetSplit(func(string) []string) }); ok {
		z.SetSplit(nil)
	}
}

// setSplit sets the split func of a slice or a map value.
func setSplit(v any, split string) {
	if split == "" {
		return
	}
	z, ok := v.(interface{ SetSplit(func(string) []string) })
	if !ok {
		return
	}
	var f func(string) []string
	switch r := []rune(split); len(r) {
	case 1:
		f = func(s string) []string {
			return SplitBy(s, r[0])
		}
	default:
		f = func(s string) []string {
			return strings.Split(s, split)
		}
	}
	z.SetSplit(f)
}
