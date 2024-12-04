package ox

import (
	"context"
	"fmt"
	"io"
	"maps"
	"reflect"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/xo/ox/text"
)

// Option is the interface for options that can be passed when creating a
// [Context], [Command], [Flag], or [Help].
//
// The [Option] type is aliased as [ContextOption], [CommandOption],
// [CommandFlagOption], [FlagOption], [HelpOption], [FlagHelpOption] and
// provided for ease-of-use, readibility, and categorization within
// documentation.
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
// A [HelpOption] can be passed to [Help] or to [NewCommandHelp].
//
// A [FlagHelpOption] can be applied to a [Flag], passed to [NewFlag], or
// passed to [Help].
//
// Additionally, a [OnErr] and [SectionMap] can be used as a [CommandOption] or
// in calls to [NewCommand], and a [Type] can be used as a [FlagOption] in
// calls to [NewFlag].
type Option interface {
	Option() option
}

// ContextOption are [Option]'s that apply to a [Context].
type ContextOption = Option

// CommandOption are [Option]'s that apply to a [Command].
type CommandOption = Option

// CommandFlagOption are [Option]'s that apply to either a [Command] or [Flag].
type CommandFlagOption = Option

// FlagOption are [Option]'s that apply to a [Flag].
type FlagOption = Option

// HelpOption are [Option]'s that apply to [Help].
type HelpOption = Option

// FlagHelpOption are [Option]'s that apply to [Flag] and [Help].
type FlagHelpOption = Option

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

// Override is a [Run]/[RunContext]/[Context] option to override expansion
// variables.
func Override(override map[string]string) ContextOption {
	return option{
		name: "Override",
		ctx: func(ctx *Context) error {
			ctx.Override = override
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
func Defaults(opts ...Option) CommandOption {
	options := []Option{Help(opts...), Comp(), Version()}
	return option{
		name: "Defaults",
		cmd: func(cmd *Command) error {
			return applyOpts(cmd, options...)
		},
		post: func(cmd *Command) error {
			return applyPostOpts(cmd, options...)
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
		name: "Version",
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

// Help is a [Command] option to add help output to a root command. By default,
// it adds a `--help` flag to all commands in the command tree. The root
// command will have a `help` sub command added if there are any other defined
// sub commands.
func Help(opts ...Option) CommandOption {
	return option{
		name: "Help",
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
			if err := NewHelpFor(cmd, opts...); err != nil {
				return err
			}
			// add help sub command
			if len(cmd.Commands) != 0 {
				if err := cmd.Sub(
					Exec(func(ctx context.Context) error {
						c, _ := Ctx(ctx)
						_, _ = cmd.HelpContext(c).WriteTo(c.Stdout)
						return ErrExit
					}),
					Usage(text.HelpCommandName, text.HelpCommandDesc),
				); err != nil {
					return err
				}
			}
			if err := addHelp(cmd); err != nil {
				return err
			}
			return nil
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
func Usage(name, usage string, aliases ...string) CommandFlagOption {
	return option{
		name: "Usage",
		cmd: func(cmd *Command) error {
			if cmd.Parent == nil && len(aliases) != 0 {
				return fmt.Errorf("%w: cannot apply to root command", ErrInvalidType)
			}
			cmd.Name, cmd.Usage, cmd.Aliases = name, usage, aliases
			return nil
		},
		flag: func(g *Flag) error {
			g.Name, g.Usage, g.Aliases = name, usage, aliases
			return nil
		},
	}
}

// Aliases is a [Command]/[Flag] option to add a alias for the command/flag.
func Aliases(aliases ...string) CommandFlagOption {
	return option{
		name: "Aliases",
		cmd: func(cmd *Command) error {
			if cmd.Parent == nil {
				return fmt.Errorf("%w: cannot apply to root command", ErrInvalidType)
			}
			cmd.Aliases = append(cmd.Aliases, aliases...)
			return nil
		},
		flag: func(g *Flag) error {
			g.Aliases = append(g.Aliases, aliases...)
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

// MapKey is a [Flag] option to set the flag's map key type.
func MapKey(mapKey Type) FlagOption {
	return option{
		name: "MapKey",
		flag: func(g *Flag) error {
			g.MapKey = mapKey
			return nil
		},
	}
}

// Elem is a [Flag] option to set the flag's element type.
func Elem(elem Type) FlagOption {
	return option{
		name: "Elem",
		flag: func(g *Flag) error {
			g.Elem = elem
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
//
// See [Context.Expand] for more expansion details.
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

// Special is a [Flag] option to set a flag's special value.
func Special(special string) FlagOption {
	return option{
		name: "Special",
		cmd: func(cmd *Command) error {
			cmd.Special = special
			return nil
		},
		flag: func(g *Flag) error {
			g.Special = special
			return nil
		},
	}
}

// Section is a [Command]/[Flag] option to set the flag's help section.
func Section(section int) CommandFlagOption {
	return option{
		name: "Section",
		cmd: func(cmd *Command) error {
			cmd.Section = section
			return nil
		},
		flag: func(g *Flag) error {
			g.Section = section
			return nil
		},
	}
}

// Hook is a [Flag] option to set a hook for a flag.
func Hook(f func(context.Context) error) FlagOption {
	return option{
		name: "Hook",
		flag: func(g *Flag) error {
			g.Type, g.Def = HookT, f
			return nil
		},
	}
}

// Banner is a [Help] option to set the banner.
func Banner(banner string) HelpOption {
	return option{
		name: "Banner",
		help: func(help *CommandHelp) error {
			help.Banner, help.NoBanner = banner, banner == ""
			return nil
		},
	}
}

// Spec is a [Flag]/[Help] option to set the use spec.
func Spec(spec string) FlagHelpOption {
	return option{
		name: "Spec",
		flag: func(g *Flag) error {
			g.Spec = spec
			return nil
		},
		help: func(help *CommandHelp) error {
			help.Spec, help.NoSpec = spec, spec == ""
			return nil
		},
	}
}

// Footer is a [Help] option to set the command/flag's footer.
func Footer(footer string) HelpOption {
	return option{
		name: "Footer",
		help: func(help *CommandHelp) error {
			help.Footer, help.NoSpec = footer, footer == ""
			return nil
		},
	}
}

// Sections is a [Help] option to set section names for commands when used with
// a [Command] or flag sections when used with [Help].
func Sections(sections ...string) HelpOption {
	return option{
		name: "Sections",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				return fmt.Errorf("%w: cannot be used with nil CommandHelp", ErrInvalidType)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.CommandSections = sections
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.Sections = sections
			return nil
		},
	}
}

// SectionMap is a [Command] option to apply refined sections to commands and
// flags.
type SectionMap map[string]int

// Option satisfies the [Option] interface.
func (m SectionMap) Option() option {
	return option{
		name: "SectionMap",
		cmd: func(*Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			for _, k := range slices.Sorted(maps.Keys(m)) {
				if strings.HasPrefix(k, "flag:") {
					if g := cmd.Flag(strings.TrimPrefix(k, "flag:"), false); g != nil {
						g.Section = m[k]
					}
				} else if c := cmd.Command(k); c != nil {
					c.Section = m[k]
				}
			}
			return nil
		},
	}
}

// option wraps an option.
type option struct {
	name string
	ctx  func(*Context) error
	cmd  func(*Command) error
	post func(*Command) error
	flag func(*Flag) error
	help func(*CommandHelp) error
	typ  func(v interface{ SetType(Type) }) error
}

// Option satisfies the [Option] interface.
func (opt option) Option() option {
	return opt
}

// apply applies the option to val.
func (opt option) apply(val any) error {
	var err error = ErrAppliedToInvalidType
	switch v := val.(type) {
	case *Command:
		if opt.ctx != nil {
			err = nil
		}
		if opt.cmd != nil {
			err = opt.cmd(v)
		}
	case *Flag:
		if opt.flag != nil {
			err = opt.flag(v)
		}
	case *Context:
		if opt.cmd != nil || opt.post != nil {
			err = nil
		}
		if opt.ctx != nil {
			err = opt.ctx(v)
		}
	case *CommandHelp:
		if opt.help != nil {
			err = opt.help(v)
		}
	case interface{ SetType(Type) }:
		if opt.typ != nil {
			err = opt.typ(v)
		}
	}
	if err != nil {
		return fmt.Errorf("%s (%T): %w", opt.name, val, err)
	}
	return nil
}

// addHelp adds help for all sub commands on cmd.
func addHelp(cmd *Command) error {
	if len(cmd.Commands) == 0 {
		return nil
	}
	for _, c := range cmd.Commands {
		if err := NewHelpFor(c); err != nil {
			return err
		}
		if err := addHelp(c); err != nil {
			return err
		}
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

// postOpts applies post options to the command.
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
