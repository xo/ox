package ox

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"unicode/utf8"
)

// Parent is a [Command] option to set the Parent for the command.
func Parent(parent *Command) Option {
	return option{
		name: "parent",
		cmd: func(c *Command) error {
			c.Parent, c.OnErr = parent, parent.OnErr
			return nil
		},
	}
}

// From is a [Command] option to build the command's flags from val.
func From[T *E, E any](val T) Option {
	return option{
		name: "From",
		cmd: func(cmd *Command) error {
			if val == nil {
				return fmt.Errorf("%T: %w: cannot be nil", val, ErrInvalidValue)
			}
			flags, err := FlagsFrom(val)
			if err != nil {
				return fmt.Errorf("cannot build flags for %T: %w", val, err)
			}
			if cmd.Flags == nil {
				cmd.Flags = Flags()
			}
			cmd.Flags.Flags = append(cmd.Flags.Flags, flags...)
			return nil
		},
	}
}

// RunArgs is a [Run] option to set the command-line arguments to use.
func RunArgs(args []string) Option {
	return option{
		name: "RunArgs",
		ctx: func(opts *RunContext) error {
			opts.Args = args
			return nil
		},
	}
}

// Version is a [Command] option to hook --version with version output.
func Version(version string, opts ...Option) Option {
	return Hook(func(ctx context.Context) error {
		_, _ = fmt.Fprintf(Stdout(ctx), "%s %s\n", RootName(ctx), version)
		return ErrExit
	},
		prepend(opts, Usage("version", "display version and exit"))...,
	)
}

// Help is a [Command] option to hook --help with help output.
func Help(opts ...Option) Option {
	return Hook(func(ctx context.Context) error {
		_, _ = fmt.Fprintf(Stdout(ctx), "%s help!\n", RootName(ctx))
		return ErrExit
	},
		prepend(opts, Usage("help", "display help and exit"))...,
	)
}

// Comp is a [Command] option to enable command completion.
func Comp() Option {
	return option{
		name: "Comp",
		cmd: func(c *Command) error {
			if c.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			return nil
		},
	}
}

// Name is a [Command]/[Flag] option to set the command/flag's name.
func Name(name string) Option {
	return option{
		name: "Name",
		cmd: func(c *Command) error {
			c.Descs[0].Name = name
			return nil
		},
		flag: func(g *Flag) error {
			g.Descs[0].Name = name
			return nil
		},
	}
}

// Usage is a [Command]/[Flag] option to set the command/flag's name and usage.
func Usage(name, usage string) Option {
	return option{
		name: "Usage",
		cmd: func(c *Command) error {
			c.Descs[0].Name, c.Descs[0].Usage = name, usage
			return nil
		},
		flag: func(g *Flag) error {
			g.Descs[0].Name, g.Descs[0].Usage = name, usage
			return nil
		},
	}
}

// Short is a [Flag] option to set the short name for a flag.
func Short(name string) Option {
	return option{
		name: "Short",
		flag: func(g *Flag) error {
			if utf8.RuneCountInString(name) != 1 {
				return fmt.Errorf("%w: %q", ErrInvalidShortName, name)
			}
			g.Descs = append(g.Descs, Desc{Name: name})
			return nil
		},
	}
}

// Alias is a [Command]/[Flag] option to set the command/flag's alias.
func Alias(name, usage string) Option {
	return option{
		name: "Alias",
		cmd: func(c *Command) error {
			c.Descs = append(c.Descs, Desc{
				Name:  name,
				Usage: usage,
			})
			return nil
		},
		flag: func(g *Flag) error {
			g.Descs = append(g.Descs, Desc{
				Name:  name,
				Usage: usage,
			})
			return nil
		},
	}
}

// ArgsFunc is a [Command] option to set the command's argument validation funcs.
func ArgsFunc(funcs ...func([]string) error) Option {
	return option{
		name: "ArgsFunc",
		cmd: func(c *Command) error {
			c.Args = append(c.Args, funcs...)
			return nil
		},
	}
}

// Args is a [Command] option to the set the range of a command's minimum/maximum
// arg count and allowed arg values. A minimum/maximum < 0 means no
// minimum/maximum.
func Args(minimum, maximum int, values ...string) Option {
	return option{
		name: "Args",
		cmd: func(c *Command) error {
			for i, s := range values {
				if s == "" {
					return fmt.Errorf("arg %q (%d) is %w", s, i, ErrInvalidValue)
				}
			}
			c.Args = append(c.Args, func(args []string) error {
				switch n := len(args); {
				case minimum < 0 && maximum < 0:
				case minimum == 0 && maximum == 0 && n != 0:
					return fmt.Errorf("%w: takes no args", ErrInvalidArgCount)
				case minimum <= 0 && maximum < n:
					return fmt.Errorf("%w: takes max %d arg(s)", ErrInvalidArgCount, maximum)
				case maximum <= 0 && n < minimum:
					return fmt.Errorf("%w: takes min %d arg(s)", ErrInvalidArgCount, minimum)
				case 0 <= minimum && 1 <= maximum && (n < minimum || maximum < n):
					return fmt.Errorf("%w: takes %d-%d args", ErrInvalidArgCount, minimum, maximum)
				}
				if len(values) != 0 {
					for i, arg := range args {
						if !slices.Contains(values, arg) {
							return fmt.Errorf("%w: %q (%d) not any of %s", ErrInvalidArg, arg, i, strings.Join(values, ", "))
						}
					}
				}
				return nil
			})
			return nil
		},
	}
}

// MapKey is a [Flag] option to set the map key type.
func MapKey(opts ...Option) Option {
	return option{
		name: "MapKey",
		flag: func(g *Flag) error {
			return nil
		},
	}
}

// Sub is a [Command] option to create a sub command.
func Sub(f func(context.Context, []string) error, opts ...Option) Option {
	return option{
		name: "Sub",
		cmd: func(c *Command) error {
			return c.Sub(f, opts...)
		},
	}
}

// UserConfigFile is a [Command] option to load a config file from the user's
// config directory.
func UserConfigFile() Option {
	return option{
		name: "UserConfigFile",
		cmd: func(c *Command) error {
			if c.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			dir, err := userConfigDir()
			if err != nil {
				return err
			}
			dir = dir
			return nil
		},
	}
}

// BindValue is a [Flag] option to set a bound value and a set flag.
func BindValue(value reflect.Value, b *bool) Option {
	return option{
		name: "BindValue",
		flag: func(g *Flag) error {
			val, err := NewBindValue(value, b)
			if err != nil {
				return err
			}
			g.Binds = append(g.Binds, val)
			return nil
		},
	}
}

// BindSet is a [Flag] option to bind a variable and its set flag.
func BindSet[T *E, E any](v T, b *bool) Option {
	return option{
		name: "BindSet",
		flag: func(g *Flag) error {
			val, err := NewBind(v, b)
			if err != nil {
				return err
			}
			g.Binds = append(g.Binds, val)
			return nil
		},
	}
}

// Bind is a [Flag] option to bind a variable.
func Bind[T *E, E any](v T) Option {
	return option{
		name: "Bind",
		flag: func(g *Flag) error {
			val, err := NewBind(v, nil)
			if err != nil {
				return err
			}
			g.Binds = append(g.Binds, val)
			return nil
		},
	}
}

// Default is a [Flag] option to set the flag's default value.
func Default(def any) Option {
	return option{
		name: "Default",
		flag: func(g *Flag) error {
			g.Def = def
			return nil
		},
	}
}

// NoArg is a [Flag] option to set that the flag expects no argument.
func NoArg(noArg bool) Option {
	return option{
		name: "NoArg",
		flag: func(g *Flag) error {
			g.NoArg = noArg
			return nil
		},
	}
}

// Key is a [Flag] option to set the flag's config lookup key for a registered
// config file type.
func Key(typ, key string) Option {
	return option{
		name: "Key",
		flag: func(g *Flag) error {
			if g.Keys == nil {
				g.Keys = make(map[string]string)
			}
			g.Keys[typ] = key
			return nil
		},
	}
}

// Hook is a [Flag] option to set a hook for a flag, that exits normally.
func Hook(f func(context.Context) error, opts ...Option) Option {
	return option{
		name: "Hook",
		flag: func(g *Flag) error {
			g.Type, g.Def = HookT, f
			return nil
		},
	}
}

// HookDump is an option to set a hook for a flag that Fprint's s and v to the
// set standard out and then exits normally.
func HookDump(s string, v ...any) Option {
	return Hook(func(ctx context.Context) error {
		_, _ = fmt.Fprintf(Stdout(ctx), s, v...)
		return ErrExit
	})
}

/*
// MustExist is a option to indicate that a path value must exist on disk.
func MustExist(mustExist bool) Option {
	return option{
		name: "MustExist",
		flag: func(g *Flag) error {
			return nil
		},
	}
}

// Relative is a option to indacet that a path value is relative to the base
// path.
func Relative(dir string) Option {
	return option{
		name: "Relative",
		flag: func(g *Flag) error {
			return nil
		},
	}
}
*/

// Option is a option.
type Option interface {
	apply(any) error
}

// option wraps an option.
type option struct {
	name string
	cmd  func(*Command) error
	set  func(*FlagSet) error
	flag func(*Flag) error
	ctx  func(*RunContext) error
}

// apply satisfies the [Option] interface.
func (opt option) apply(val any) error {
	var err error
	switch v := val.(type) {
	case *Command:
		if opt.cmd != nil {
			err = opt.cmd(v)
		}
	case *Flag:
		if opt.flag != nil {
			err = opt.flag(v)
		}
	case *RunContext:
		if opt.ctx != nil {
			err = opt.ctx(v)
		}
	default:
		err = ErrAppliedToInvalidType
	}
	if err != nil {
		return fmt.Errorf("%s: %w", opt.name, err)
	}
	return nil
}
