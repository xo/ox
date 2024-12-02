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

// Args is a [Run]/[RunContext]/[Context] option to set the command-line
// arguments.
func Args(args ...string) ContextOption {
	return option{
		name: "Args",
		ctx: func(ctx *Context) error {
			ctx.Args = args
			return nil
		},
	}
}

// Pipe is a [Run]/[RunContext]/[Context] option to set the standard in, out,
// and error.
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
func Exec[F ExecType](f F) CommandOption {
	return option{
		name: "Exec",
		cmd: func(cmd *Command) error {
			var err error
			cmd.Exec, err = NewExec(f)
			return err
		},
	}
}

// Defaults is a [Command] option to add [Help], [Comp], and [Version] to the
// command.
func Defaults() CommandOption {
	opts := []Option{Help(), Comp(), Version()}
	return option{
		name: "Defaults",
		cmd: func(cmd *Command) error {
			return applyOpts(cmd, opts...)
		},
		post: func(cmd *Command) error {
			return applyPostOpts(cmd, opts...)
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
				typ := reflect.TypeOf(val).String()
				if strings.HasPrefix(typ, "*struct {") {
					typ = "*struct"
				}
				return fmt.Errorf("cannot build flags for %s: %w", typ, err)
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
		cmd: func(cmd *Command) error {
			return cmd.Sub(opts...)
		},
	}
}

// Version is a [Command] option to hook --version with version output.
func Version(opts ...Option) CommandOption {
	return option{
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			return NewVersionFor(cmd, opts...)
		},
	}
}

// Help is a [Command] option to hook `--help` with help output, as well as to
// add a `help` sub command to the tree.
func Help(opts ...Option) CommandOption {
	return option{
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			return NewHelpFor(cmd, opts...)
		},
	}
}

// Comp is a [Command] option to enable command completion.
func Comp() CommandOption {
	return option{
		name: "Comp",
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
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
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
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

// ArgsFunc is a [Command] option to add funcs to command's argument validation
// funcs.
func ArgsFunc(funcs ...func([]string) error) CommandOption {
	return option{
		name: "ArgsFunc",
		cmd: func(cmd *Command) error {
			cmd.Args = append(cmd.Args, funcs...)
			return nil
		},
	}
}

// ValidArgs is a [Command] option to add a argument validation func to the
// command that validates the range of allowed minimum/maximum argruments and
// allowed argument values. A minimum/maximum < 0 means no minimum/maximum.
func ValidArgs(minimum, maximum int, values ...string) CommandOption {
	return option{
		name: "Args",
		cmd: func(cmd *Command) error {
			for i, s := range values {
				if s == "" {
					return fmt.Errorf("%w: argument %d cannot be empty string", ErrInvalidArg, i)
				}
			}
			cmd.Args = append(cmd.Args, func(args []string) error {
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
		cmd: func(cmd *Command) error {
			cmd.Parent, cmd.OnErr = parent, parent.OnErr
			return nil
		},
	}
}

// Name is a [Command]/[Flag] option to set the command/flag's name.
func Name(name string) CommandFlagOption {
	return option{
		name: "Name",
		cmd: func(cmd *Command) error {
			cmd.Name = name
			return nil
		},
		flag: func(g *Flag) error {
			g.Name = name
			return nil
		},
	}
}

// Usage is a [Command]/[Flag] option to set the command/flag's name and usage.
func Usage(name, usage string) CommandFlagOption {
	return option{
		name: "Usage",
		cmd: func(cmd *Command) error {
			cmd.Name, cmd.Usage = name, usage
			return nil
		},
		flag: func(g *Flag) error {
			g.Name, g.Usage = name, usage
			return nil
		},
	}
}

// Alias is a [Command]/[Flag] option to add a alias for the command/flag.
func Alias(name, usage string) CommandFlagOption {
	return option{
		name: "Alias",
		cmd: func(cmd *Command) error {
			return nil
		},
		flag: func(g *Flag) error {
			g.Name, g.Usage = name, usage
			return nil
		},
	}
}

// Short is a [Flag] option to add a flag's short (single character) alias.
func Short(short string) FlagOption {
	return option{
		name: "Short",
		flag: func(g *Flag) error {
			if utf8.RuneCountInString(short) != 1 {
				return fmt.Errorf("%w: %q", ErrInvalidShortName, short)
			}
			g.Short = short
			return nil
		},
	}
}

// Banner is a [Command]/[Flag] option to set the command/flag's banner.
func Banner(banner string) CommandFlagOption {
	return option{
		name: "Banner",
		flag: func(g *Flag) error {
			return nil
		},
	}
}

// Spec is a [Command]/[Flag] option to set the command/flag's spec.
func Spec(spec string) CommandFlagOption {
	return option{
		name: "Spec",
		flag: func(g *Flag) error {
			return nil
		},
	}
}

// Footer is a [Command]/[Flag] option to set the command/flag's footer.
func Footer(footer string) CommandFlagOption {
	return option{
		name: "Footer",
		flag: func(g *Flag) error {
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

// BindRef is a [Flag] option to add a reflected value to a flag.
func BindRef(value reflect.Value, b *bool) FlagOption {
	return option{
		name: "RefValue",
		flag: func(g *Flag) error {
			val, err := NewRef(value, b)
			if err != nil {
				return err
			}
			g.Binds = append(g.Binds, val)
			return nil
		},
	}
}

// Default is a [Flag] option to set the flag's default value.
//
// Special strings can be used are expanded when the flag is created:
//
//	$HOME - the current user's home directory
//	$USER - the current user's user name
//	$CACHE - the current user's cache directory
//	$APPCACHE - the current user's cache directory, with the root command's name added as a subdir
//	$ENV{KEY} - the environment value for $KEY
//	$CFG{[TYPE::]KEY} - the registered config file loader type and key value
//
// For example:
//
//	ox.Default("$USER") - expands to "ken" if the user running the application is "ken"
//	ox.Default("$HOME") - expands to /home/$USER on most Unix systems
//	ox.Default("$CACHE") - expands to /home/$USER/.cache on most Unix systems
//	ox.Default("$APPCACHE") - expands to /home/$USER/.cache/myApp if the root command's name is "myApp" on most Unix systems
//	ox.Default("$ENV{MY_VAR}") - expands to value of the environment var $MY_VAR
//	ox.Default("$CFG{yaml::a.b.c}") - expands to the registered YAML config file's key of a.b.c
//	ox.Default("$CFG{a.b.c}") - expands to the first registered config file that returns a non-nil value for key a.b.c
func Default(def any) FlagOption {
	return option{
		name: "Default",
		flag: func(g *Flag) error {
			g.Def = def
			return nil
		},
	}
}

// NoArg is a [Flag] option to set that the flag expects no argument, and the
// default value to set.
func NoArg(noArg bool, def any) FlagOption {
	return option{
		name: "NoArg",
		flag: func(g *Flag) error {
			g.NoArg, g.NoArgDef = noArg, def
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

// FlagOption is a [Flag] option to set the flag's help section.
func Section(section int) FlagOption {
	return option{
		name: "Section",
		flag: func(g *Flag) error {
			return nil
		},
	}
}

// Hook is a [Flag] option to set a hook for a flag, that exits normally.
func Hook(f func(context.Context) error) CommandFlagOption {
	return option{
		name: "Hook",
		flag: func(g *Flag) error {
			g.Type, g.Def = HookT, f
			return nil
		},
	}
}

// Sections is a [Help] option to add section names for help.
func Sections(sections ...string) HelpOption {
	return option{
		/*
			help: func(*HelpDesc) error {
				return nil
			},
		*/
	}
}

// opt is the interface for options that can be passed when creating a
// [Context], [Command], or [Flag].
//
// The [opt] type is aliased as [ContextOption], [CommandOption],
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
	Option() option
}

// ContextOption are [opt]'s that apply to a [Context].
type ContextOption = Option

// CommandOption are [opt]'s that apply to a [Command].
type CommandOption = Option

// CommandFlagOption are [opt]'s that apply to either a [Command] or [Flag].
type CommandFlagOption = Option

// FlagOption are [opt]'s that apply to a [Flag].
type FlagOption = Option

// HelpOption are [opt]'s that apply to [Help].
type HelpOption = Option

// option wraps an option.
type option struct {
	name string
	ctx  func(*Context) error
	cmd  func(*Command) error
	post func(*Command) error
	flag func(*Flag) error
	typ  func(v interface{ SetType(Type) }) error
}

func (opt option) Option() option {
	return opt
}

// apply satisfies the [opt] interface.
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
	case interface{ SetType(Type) }:
		if opt.typ != nil {
			err = opt.typ(v)
		}
	default:
		err = ErrAppliedToInvalidType
	}
	if err != nil {
		return fmt.Errorf("%s: %w", opt.name, err)
	}
	return nil
}

// applyOpts applies the options to v.
func applyOpts(v any, opts ...Option) error {
	for _, o := range opts {
		if err := o.Option().apply(v); err != nil {
			return err
		}
	}
	return nil
}

// postOpts returns the post options.
func applyPostOpts(cmd *Command, opts ...Option) error {
	for _, o := range opts {
		if opt := o.Option(); opt.post != nil {
			if err := opt.post(cmd); err != nil {
				return err
			}
		}
	}
	return nil
}
