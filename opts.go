package ox

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"slices"
	"strings"
	"unicode/utf8"
)

// Args is a [Context] option to set the command-line arguments to use.
func Args(args []string) ContextOption {
	return option{
		name: "Args",
		ctx: func(ctx *Context) error {
			ctx.Args = args
			return nil
		},
	}
}

// Pipe is a [Context] option to set the standard in, out, and error to use.
func Pipe(stdin io.Reader, stdout, stderr io.Writer) ContextOption {
	return option{
		name: "Pipe",
		ctx: func(ctx *Context) error {
			ctx.Stdin, ctx.Stdout, ctx.Stderr = stdin, stdout, stderr
			return nil
		},
	}
}

// Exec is a [Command] option to set the exec func.
func Exec[F ExecType](f F) Option {
	return option{
		name: "Exec",
		cmd: func(cmd *Command) error {
			cmd.Exec = NewExec(f)
			return nil
		},
	}
}

// From is a [Command] option to build the command's flags from val.
func From[T *E, E any](val T) CommandOption {
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

// Sub is a [Command] option to create a sub command.
func Sub(opts ...Option) CommandOption {
	return option{
		name: "Sub",
		cmd: func(c *Command) error {
			return c.Sub(opts...)
		},
	}
}

// Version is a [Command] option to hook --version with version output.
func Version(version string, opts ...Option) CommandOption {
	return option{
		cmd: func(*Command) error {
			return nil
		},
	}
}

// Help is a [Command] option to hook --help with help output.
func Help(opts ...Option) CommandOption {
	return option{
		cmd: func(*Command) error {
			return nil
		},
	}
}

// Comp is a [Command] option to enable command completion.
func Comp() CommandOption {
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

// UserConfigFile is a [Command] option to load a config file from the user's
// config directory.
func UserConfigFile() CommandOption {
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

// ArgsFunc is a [Command] option to set the command's argument validation funcs.
func ArgsFunc(funcs ...func([]string) error) CommandOption {
	return option{
		name: "ArgsFunc",
		cmd: func(c *Command) error {
			c.Args = append(c.Args, funcs...)
			return nil
		},
	}
}

// ValidArgs is a [Command] option to the set the command's range of allowed
// minimum/maximum argruments and allowed argument values. A minimum/maximum <
// 0 means no minimum/maximum.
func ValidArgs(minimum, maximum int, values ...string) CommandOption {
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

// Parent is a [Command] option to set the command's parent.
func Parent(parent *Command) CommandOption {
	return option{
		name: "Parent",
		cmd: func(c *Command) error {
			c.Parent, c.OnErr = parent, parent.OnErr
			return nil
		},
	}
}

// Name is a [Command]/[Flag] option to set the command/flag's name.
func Name(name string) CommandFlagOption {
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
func Usage(name, usage string) CommandFlagOption {
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

// Alias is a [Command]/[Flag] option to add a alias for the command/flag.
func Alias(name, usage string) CommandFlagOption {
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

// Short is a [Flag] option to add a flag's short (single character) alias.
func Short(name string) FlagOption {
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

// MapKey is a [Flag] option to set the flag's map key type.
func MapKey(typ Type) FlagOption {
	return option{
		name: "MapKey",
		flag: func(g *Flag) error {
			g.Key = typ
			return nil
		},
	}
}

// BindValue is a [Flag] option to add a bound value to a flag.
func BindValue(value reflect.Value, b *bool) FlagOption {
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

// BindSet is a [Flag] option to add a bound variable to a flag.
func BindSet[T *E, E any](v T, b *bool) FlagOption {
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

// Bind is a [Flag] option to add a bound variable to a flag.
func Bind[T *E, E any](v T) FlagOption {
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
func Default(def any) FlagOption {
	return option{
		name: "Default",
		flag: func(g *Flag) error {
			g.Def = def
			return nil
		},
	}
}

// NoArg is a [Flag] option to set that the flag expects no argument.
func NoArg(noArg bool) FlagOption {
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
func Key(typ, key string) FlagOption {
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

// Option is the interface for options that can be passed when creating a
// [Context], [Command], or [Flag].
//
// The [Option] type is aliased as [ContextOption], [CommandOption],
// [CommandFlagOption], and [FlagOption] and provided for ease-of-use,
// readibility, and categorization within documentation.
//
// A [ContextOption] can be applied to a [Context] and passed to [Run].
//
// A [CommandOption] can be applied to a [Command] and passed to [NewCommand].
//
// A [CommandFlagOption] can be applied to either a [Command] or [Flag], and
// can be passed to [NewCommand] and [NewFlag].
//
// A [FlagOption] can be applied to a [Flag] and passed to [NewFlag].
//
// Additionally, a [OnErr] can be used as a [CommandOption] or in calls to
// [NewCommand], and a [Type] can be used as a [FlagOption] in calls to
// [NewFlag].
type Option interface {
	apply(any) error
}

// ContextOption are [Option]'s that apply to a [Context].
type ContextOption = Option

// CommandOption are [Option]'s that apply to a [Command].
type CommandOption = Option

// CommandFlagOption are [Option]'s that apply to either a [Command] or [Flag].
type CommandFlagOption = Option

// FlagOption are [Option]'s that apply to a [Flag].
type FlagOption = Option

// option wraps an option.
type option struct {
	name string
	ctx  func(*Context) error
	cmd  func(*Command) error
	set  func(*FlagSet) error
	flag func(*Flag) error
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
	case *Context:
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
