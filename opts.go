package ox

import (
	"cmp"
	"context"
	"fmt"
	"io"
	"maps"
	"reflect"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Option is the interface for options that can be passed when creating a
// [Context], [Command], [Flag], or [Help].
//
// The [Option] type is aliased as [ContextOption], [CommandOption],
// [CommandFlagOption], [FlagOption], [HelpOption], [FlagHelpOption] and
// provided for ease-of-use, readibility, and categorization within
// documentation.
//
// A [ContextOption] can be applied to a [Context] and passed to [Run]/[RunContext].
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

// ArgOption are [Option]'s that apply to an [Argument].
type ArgOption = Option

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

// Interpolate is a [Run]/[RunContext]/[Context] option to set a interpolation
// func.
func Interpolate(interpolate func(*Context, any) (any, error)) ContextOption {
	return option{
		name: "Interpolate",
		ctx: func(ctx *Context) error {
			ctx.Interpolate = interpolate
			return nil
		},
	}
}

// Lookup is a [Run]/[RunContext]/[Context] option to set the func that looks
// up a config variable.
func Lookup(lookup func(*Context, ConfigType, string) (any, error)) ContextOption {
	return option{
		name: "Lookup",
		ctx: func(ctx *Context) error {
			ctx.Lookup = lookup
			return nil
		},
	}
}

// Override is a [Run]/[RunContext]/[Context] option to set a override
// expansion func.
func Override(override func(ConfigType, string) (any, bool, error)) ContextOption {
	return option{
		name: "Override",
		ctx: func(ctx *Context) error {
			ctx.Override = override
			return nil
		},
	}
}

// OverrideMap is a [Run]/[RunContext]/[Context] option to override expansion
// variables using a map.
func OverrideMap(override map[string]string) ContextOption {
	return option{
		name: "OverrideMap",
		ctx: func(ctx *Context) error {
			ctx.Override = func(typ ConfigType, key string) (any, bool, error) {
				if typ != "" {
					key = strings.ToUpper(string(typ)) + "::" + key
				}
				s, ok := override[key]
				return s, ok, nil
			}
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

// Before is a [Command] option to set a func that runs before the exec func
// of the command, and of each sub command below it. f takes any signature
// that [Exec] takes, and gets the same [context.Context] and args as the exec
// func. The before funcs run from the root down to the command that runs. When
// one returns an error, the run stops: no later before func, no exec func and
// no after func runs. Help, the version and completion do not run it.
func Before[F ExecType](f F) CommandOption {
	return option{
		name: "Before",
		cmd: func(cmd *Command) error {
			var err error
			cmd.Before, err = NewExec(f)
			return err
		},
	}
}

// After is a [Command] option to set a func that runs after the exec func of
// the command, and of each sub command below it. f takes any signature that
// [Exec] takes. The after funcs run from the command that runs up to the root,
// also when the exec func returns an error. Each one runs, whatever an earlier
// one returns. The run returns the errors of the exec func and of the after
// funcs, joined with [errors.Join]. Help, the version and completion do not
// run it.
func After[F ExecType](f F) CommandOption {
	return option{
		name: "After",
		cmd: func(cmd *Command) error {
			var err error
			cmd.After, err = NewExec(f)
			return err
		},
	}
}

// Defaults is a [Command] option to add [Version], [Comp], and [Help] options
// to the command. Note that the order is important.
func Defaults(opts ...Option) CommandOption {
	options := []Option{Version(), Comp(), Help(opts...)}
	return option{
		name: "Defaults",
		cmd: func(cmd *Command) error {
			return Apply(cmd, options...)
		},
		post: func(cmd *Command) error {
			return ApplyPost(cmd, options...)
		},
	}
}

// From is a [Command] option to build the command's flags from a value of type
// *struct using reflection. Adds flags for all exported fields of the passed
// value having a `ox` tag and a non-empty description.
//
// Example:
//
//	args := struct{
//		MyFlag      string   `ox:"my flag,short:f,default:$USER"`
//		MyVerbosity int      `ox:"my verbosity,type:count,short:v"`
//		MyURL       *url.URL `ox:"my url,set:MyURLSet"`
//		MyURLSet    bool
//		MyOtherFlag string   `ox:"a long\\, long description,short:F"`
//		MyFloat     float64  `ox:"my float,hidden,name:MYF"`
//	}{}
//
//	ox.Run(
//		ox.From(&args),
//	)
//
// See [FlagsFrom] for more information for available tag options.
func From[T *E, E any](val T) CommandOption {
	return option{
		name: "From",
		cmd: func(cmd *Command) error {
			if val == nil {
				return fmt.Errorf("%T: %w: cannot be nil", val, ErrInvalidValue)
			}
			flags, args, err := appendFlags(nil, nil, reflect.ValueOf(val), nil)
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
			addTagGroups(cmd, flags)
			return cmd.addArgs(args...)
		},
	}
}

// App is a [Command] option to build the command, its flags and its sub
// commands from v, a pointer to a struct:
//
//   - The doc comment of the struct type gives the usage and the help of the
//     command. A section starts with a line that holds only its name and a
//     colon. Each line after it, with one tab removed, is the text of the
//     section, up to the next section. The sections are `Description:`, the
//     usage of the command, `Usage:`, whose first line is the name of the
//     command after the names of its parents, then its spec, `Aliases:`,
//     `Section:`, `Banner:`, `Examples:`, `Sections:`, the names of the flag
//     sections, `Command sections:`, the names of the sub command sections,
//     and `Footer:`. A sub command without a usage line takes its name from
//     its type, without the suffix `Command` and without the name of the type
//     of its parent, such as `list` for `AppListCommand`.
//   - Each field with an `ox` tag is a flag, as [From] makes it. When the tag
//     has no description, the doc comment of the field gives the usage.
//   - When the struct has a method `Commands() []any`, each value that it
//     returns is a sub command, built the same way.
//   - When the struct has a method `Exec`, with a signature that [Exec]
//     takes, such as `Exec(context.Context, []string) error`, it is the exec
//     func of the command.
//   - When the struct has a method `Before` or `After`, with a signature that
//     [Exec] takes, it is the func of [Before] or [After] of the command.
//
// Go keeps no doc comment in a program, so the program registers them with
// [RegisterDoc]. The tool `github.com/xo/ox/cmd/oxdoc` writes those calls
// from `go generate`.
func App[T *E, E any](v T) CommandOption {
	return appOption(v, "")
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

// Extend is a root [Command] option that runs f on the root command after
// each other option, so that f finds every flag and sub command, and the
// help, version and completion of [Defaults]. A package outside ox uses it to
// add to the command tree, as [github.com/xo/ox/metadoc.Command] does. The
// name names the option in an error.
func Extend(name string, f func(*Command) error) CommandOption {
	return option{
		name: name,
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			return nil
		},
		last: f,
	}
}

// Version is a [Command] option to hook --version with version output derived
// from Go's build info.
func Version() CommandOption {
	return option{
		name: "Version",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			switch {
			case cmd.Parent != nil:
				return ErrCanOnlyBeUsedWithRootCommand
			case len(cmd.Commands) != 0:
				if err := NewVersion(cmd); err != nil {
					return err
				}
				// a flag that the program defines with hook:version works
				// beside the version command
				if cmd.FlagSpecial("hook:version") == nil {
					return nil
				}
			}
			return NewVersionFlag(cmd)
		},
	}
}

// VersionString is a [Command] option to set the version displayed for the
// command.
func VersionString(version string) CommandOption {
	return option{
		name: "VersionString",
		cmd: func(cmd *Command) error {
			cmd.Version = version
			return nil
		},
	}
}

// AppNameVersion is a [Command] option to set the root app name and version.
// Mostly used with [AppFrom].
func AppNameVersion(name, version string) CommandOption {
	return option{
		name: "AppNameVersion",
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			cmd.Name, cmd.Version = name, version
			return nil
		},
	}
}

// Help is a [Command]/[Flag] option to add help output to a root command. By
// default, adds a `--help` flag to all commands in the command tree including
// the root command.
//
// Additionalyl, a `help` subcommand will be added to the root command if the
// root command has at least one other subcommand not named `help`.
func Help(opts ...Option) CommandFlagOption {
	return option{
		name: "Help",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if err := NewHelpFlag(cmd, opts...); err != nil {
				return err
			}
			// add help sub command
			if len(cmd.Commands) != 0 && cmd.Parent == nil {
				if err := NewHelp(cmd); err != nil {
					return err
				}
			}
			// add to all sub commands
			if err := AddHelpFlag(cmd); err != nil {
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
		post: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			f := NewCompFlags
			if len(cmd.Commands) != 0 {
				f = NewComp
			}
			if err := f(cmd); err != nil {
				return err
			}
			cmd.Comp = true
			return nil
		},
	}
}

// ConfigFiles is a root [Command] option to add config files to read, in
// order. Each path is interpolated, such as "$APPCONFIG/config.yaml". A later
// file wins over an earlier one. A file that does not exist is skipped. The
// format comes from the extension of the file. See [RegisterConfigFormat].
func ConfigFiles(paths ...string) CommandOption {
	return option{
		name: "ConfigFiles",
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			cmd.ConfigFiles = append(cmd.ConfigFiles, paths...)
			return nil
		},
	}
}

// ConfigRequired is a root [Command] option that makes it an error when no
// config file exists.
func ConfigRequired() CommandOption {
	return option{
		name: "ConfigRequired",
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			cmd.ConfigRequired = true
			return nil
		},
	}
}

// EnvPrefix is a root [Command] option to fill each flag from the
// environment variable PREFIX_NAME, with the name of the flag in upper case
// and each "-" as "_", such as APP_LOG_LEVEL for --log-level. The command
// line wins over the environment, and the environment wins over the config
// files and the default. See [NoEnv] and [EnvVar].
func EnvPrefix(prefix string) CommandOption {
	return option{
		name: "EnvPrefix",
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			cmd.EnvPrefix = prefix
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
// Completion offers the values at each position up to the maximum. See
// [Command.ValidArgs].
func ValidArgs(minimum, maximum int, values ...string) CommandOption {
	return option{
		name: "ValidArgs",
		cmd: func(cmd *Command) error {
			for i, s := range values {
				if s == "" {
					return fmt.Errorf("%w: argument %d cannot be empty string", ErrInvalidArg, i)
				}
			}
			cmd.ValidArgs = append(cmd.ValidArgs, values...)
			cmd.maxArgs, cmd.noArgs = maximum, minimum == 0 && maximum == 0
			if maximum == 0 && 0 < minimum {
				cmd.maxArgs = -1
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

// OnlyOneOf is a [Command] option to allow one of the named flags at most.
// When two or more of them get a value from the command line, the
// environment or a config file, [Context.Validate] returns an error that
// wraps [ErrOnlyOneOf] and names them. A default does not count. The struct
// tag option is `onlyoneof:GROUP`, which adds the flag to the group GROUP.
func OnlyOneOf(names ...string) CommandOption {
	return groupCommandOption("OnlyOneOf", RuleOnlyOneOf, names)
}

// AllOrNoneOf is a [Command] option to need each of the named flags, or none
// of them. When some of them get a value from the command line, the
// environment or a config file, and others do not, [Context.Validate]
// returns an error that wraps [ErrAllOrNoneOf] and names them. A default does
// not count. The struct tag option is `allornoneof:GROUP`.
func AllOrNoneOf(names ...string) CommandOption {
	return groupCommandOption("AllOrNoneOf", RuleAllOrNoneOf, names)
}

// AtLeastOneOf is a [Command] option to need one of the named flags at least.
// When none of them gets a value from the command line, the environment or a
// config file, [Context.Validate] returns an error that wraps
// [ErrAtLeastOneOf] and names them. A default does not count. The struct tag
// option is `atleastoneof:GROUP`.
func AtLeastOneOf(names ...string) CommandOption {
	return groupCommandOption("AtLeastOneOf", RuleAtLeastOneOf, names)
}

// groupCommandOption returns the option that adds a [FlagGroup] to a
// command. After the command has its flags, each name must be the long name
// or a long alias of a flag of the command or of a parent.
func groupCommandOption(name string, rule Rule, names []string) CommandOption {
	return option{
		name: name,
		cmd: func(*Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			grp := FlagGroup{
				Rule:  rule,
				Names: make([]string, len(names)),
			}
			for i, s := range names {
				g := cmd.Flag(s, true, false)
				if g == nil {
					return fmt.Errorf("%s: %w", name, newFlagError("--"+s, ErrUnknownFlag))
				}
				grp.Names[i] = g.Name
			}
			cmd.FlagGroups = append(cmd.FlagGroups, grp)
			return nil
		},
	}
}

// Arg is a [Command] option to add a positional argument to the command, after
// the arguments that it has. opts are each [ArgOption] and [FlagOption] of
// [NewArgument], such as a [Type], [Elem], [Default], [Valid], [Bind],
// [CompFunc], [Spec] or [Optional].
//
// The parse converts each argument to its type, and checks it, as for a flag.
// So a [FileT] must be a file that exists. [Context.Vars] holds the value by
// the name of the argument, so that the exec func reads it with `otx.Get`.
// An argument is required, unless it has a default or [Optional]. A
// required argument cannot follow an optional one. The last argument can be
// variadic, with the type [SliceT] and an [Elem], and it then takes each
// argument that is left. Help writes each argument in the usage line and in
// the `Arguments:` section, and completion completes it by its type.
func Arg(name, usage string, opts ...Option) CommandOption {
	return option{
		name: "Arg",
		cmd: func(cmd *Command) error {
			a, err := NewArgument(name, usage, opts...)
			if err != nil {
				return err
			}
			return cmd.addArgs(a)
		},
	}
}

// Optional is an [ArgOption] that makes a positional argument optional, so
// that the command runs without it.
func Optional() ArgOption {
	return option{
		name: "Optional",
		arg: func(a *Argument) error {
			a.Optional = true
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

// Usage is a [Command]/[Flag] option to set the command/flag's name, usage,
// and aliases.
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

// Aliases is a [Command]/[Flag] option to add aliases for the command/flag.
func Aliases(aliases ...string) CommandFlagOption {
	return option{
		name: "Aliases",
		cmd: func(cmd *Command) error {
			cmd.Aliases = append(cmd.Aliases, aliases...)
			return nil
		},
		flag: func(g *Flag) error {
			g.Aliases = append(g.Aliases, aliases...)
			return nil
		},
	}
}

// HiddenAliases is a [Flag] option to add aliases that help does not write.
// A hidden alias parses like any alias, and completion offers it.
func HiddenAliases(aliases ...string) FlagOption {
	return option{
		name: "HiddenAliases",
		flag: func(g *Flag) error {
			g.HiddenAliases = append(g.HiddenAliases, aliases...)
			return nil
		},
	}
}

// DeprecatedAliases is a [Flag] option to add aliases that neither help nor
// completion shows. A deprecated alias parses like any alias, so an old name
// keeps working.
func DeprecatedAliases(aliases ...string) FlagOption {
	return option{
		name: "DeprecatedAliases",
		flag: func(g *Flag) error {
			g.DeprecatedAliases = append(g.DeprecatedAliases, aliases...)
			return nil
		},
	}
}

// CompFunc is a [Command]/[Flag] option to set the func that completes the
// positional arguments of the command, or the value of the flag. f returns
// the completions for toComplete, the word under the cursor. ctx.Exec is the
// command, ctx.Vars holds the flags parsed so far, and ctx.Args holds the
// positional arguments before toComplete.
//
// Completion keeps only the results that start with toComplete, unless f
// returns [CompNoFilter]. Completion corrects no result. If f returns no
// completions and does not return [CompNoFileComp], the shell completes file
// names.
func CompFunc(f func(ctx *Context, toComplete string) ([]Completion, CompDirective, error)) CommandFlagOption {
	return option{
		name: "CompFunc",
		cmd: func(cmd *Command) error {
			cmd.CompFunc = f
			return nil
		},
		flag: func(g *Flag) error {
			g.CompFunc = f
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

// Split is a [Flag] option to add a flag's split separator.
func Split(split string) FlagOption {
	return option{
		name: "Split",
		flag: func(g *Flag) error {
			g.Split = split
			return nil
		},
	}
}

// Strftime is a [Flag] option to read and write the time of a flag in a
// strftime layout, such as `%d/%m/%Y %H:%M`. A flag of type [TimestampT],
// [DateTimeT], [DateT] or [TimeT] keeps its type, and a flag of another type
// gets the type [TimestampT]. A slice, an array, a map or a hook flag refuses
// it.
//
// Strftime converts the layout to a Go layout, and sets [Flag.Layout]. It
// takes only the directives that Go writes and reads in the same way:
//
//	%Y  the year, such as 2006
//	%y  the year in 2 digits, such as 06
//	%m  the month in 2 digits, such as 01
//	%b  the short name of the month, such as Jan
//	%B  the name of the month, such as January
//	%d  the day of the month in 2 digits, such as 02
//	%e  the day of the month, padded with a space, such as " 2"
//	%j  the day of the year in 3 digits, such as 002
//	%a  the short name of the day of the week, such as Mon
//	%A  the name of the day of the week, such as Monday
//	%H  the hour of 24 hours in 2 digits, such as 15
//	%I  the hour of 12 hours in 2 digits, such as 03
//	%p  AM or PM
//	%M  the minute in 2 digits, such as 04
//	%S  the second in 2 digits, such as 05
//	%f  the microseconds in 6 digits, after a . or a ,
//	%L  the milliseconds in 3 digits, after a . or a ,
//	%z  the offset of the zone, such as -0700
//	%Z  the name of the zone, such as MST
//	%T  the same as %H:%M:%S
//	%R  the same as %H:%M
//	%D  the same as %m/%d/%y
//	%F  the same as %Y-%m-%d
//	%%  a %
//
// Strftime refuses any other directive, and text between the directives that
// Go reads as a part of a time, such as 1 or Jan. Its error wraps
// [ErrInvalidLayout].
func Strftime(layout string) FlagOption {
	return option{
		name: "Strftime",
		flag: func(g *Flag) error {
			goLayout, err := strftimeLayout(layout)
			if err != nil {
				return err
			}
			switch g.Type {
			case TimestampT, DateTimeT, DateT, TimeT:
			case SliceT, ArrayT, MapT, HookT:
				return fmt.Errorf("Strftime(%q) (%s): %w", layout, g.Type, ErrAppliedToInvalidType)
			default:
				g.Type = TimestampT
			}
			g.Layout = goLayout
			return nil
		},
	}
}

// Suggested is a [Command] option to add suggested names for the command.
func Suggested(suggested ...string) CommandOption {
	return option{
		name: "Suggested",
		cmd: func(cmd *Command) error {
			cmd.Suggested = append(cmd.Suggested, suggested...)
			return nil
		},
	}
}

// StopAtArg is a [Command] option to end the flags of the command at its
// first positional argument, as pflag's SetInterspersed(false) does. Each
// later argument is a positional argument, even one that starts with `-`, and
// a later `--` too. A wrapper command, such as `app exec cmd --its-own-flag`,
// uses it.
//
// A sub command name comes before the first positional argument, so
// StopAtArg on a command with sub commands still finds them. A sub command
// does not inherit StopAtArg. A `--` before the first positional argument
// ends the flags, as it does without StopAtArg.
func StopAtArg() CommandOption {
	return option{
		name: "StopAtArg",
		cmd: func(cmd *Command) error {
			cmd.StopAtArg = true
			return nil
		},
	}
}

// Abbrev is a [Command] option to let the command line give a long flag by a
// unique prefix of its name, as GNU getopt_long does: `--verb` means
// `--verbose` when no other long flag of the command starts with `verb`. A
// sub command inherits it.
//
// An exact name always wins over a prefix. A prefix that matches more than
// one flag fails with [ErrAmbiguousFlag], and the error names the flags
// that it matches. A prefix matches the long names that completion offers:
// the name, the aliases and the hidden aliases of each flag that the command
// accepts, including the flags that it inherits. A prefix does not match a
// hidden flag, a deprecated flag or a deprecated alias, which need their
// exact name. A short flag needs its exact name too.
func Abbrev() CommandOption {
	return option{
		name: "Abbrev",
		cmd: func(cmd *Command) error {
			cmd.Abbrev = true
			return nil
		},
	}
}

// PlusWords is a [Command] option that reads each word that starts with `+`
// as the plus forms of short flags, as bash does, and not only a word whose
// first letter names a flag with [AllowPlus] (D52). So a letter without a
// plus form, such as the `z` of `+z`, is an error that wraps
// [ErrUnknownFlag], and a lone `+` is skipped. A command takes each word so
// when it or a parent has PlusWords, so a sub command inherits it. A program
// that takes an argument such as the `+%Y` of `date` does not use it.
func PlusWords() CommandOption {
	return option{
		name: "PlusWords",
		cmd: func(cmd *Command) error {
			cmd.PlusWords = true
			return nil
		},
	}
}

// LetterClusters is a [Command] option that reads short letters as bash
// does. A letter that takes a value takes it only from the next word, so
// `-oe errexit` sets e, and o takes errexit, where `getopt` gives o the value
// e. Each letter that takes a value takes the next word in turn. An `=` is a
// letter, so `-e=` is an error that wraps [ErrUnknownFlag], as bash refuses
// it. It applies to the `+` forms too. A command reads its letters so when it
// or a parent has LetterClusters, so a sub command inherits it. A long flag
// does not change.
func LetterClusters() CommandOption {
	return option{
		name: "LetterClusters",
		cmd: func(cmd *Command) error {
			cmd.LetterClusters = true
			return nil
		},
	}
}

// BashLongFlags is a [Command] option that reads long flags as bash does. A
// long flag after the first short flag, or after the first plus form, is an
// error that wraps [ErrUnknownFlag], as `bash -e --noediting` fails. A long
// flag with a value after `=`, such as `--norc=1` or `--rcfile=x`, is an
// error that wraps [ErrInvalidValue], so a long flag takes its value only
// from the next word. `--` still ends the flags. A command reads its long
// flags so when it or a parent has BashLongFlags, so a sub command inherits
// it (D54).
func BashLongFlags() CommandOption {
	return option{
		name: "BashLongFlags",
		cmd: func(cmd *Command) error {
			cmd.BashLongFlags = true
			return nil
		},
	}
}

// NegateBools is a root [Command] option to give each bool flag of the
// command tree a negated form, such as `--no-debug` for `--debug`. The
// negated form sets the flag to false, and takes no value. It counts as a
// value from the command line, as the flag does. Without a prefix, the
// prefix is [DefaultNegatePrefix]. Each prefix adds a form, so
// `NegateBools("no", "without")` gives `--no-debug` and `--without-debug`.
// A prefix is a word of letters, digits and hyphens, with no hyphen at
// either end.
//
// Only a bool flag that the command line sets to true without a value gets
// a negated form. A count flag and a hook flag, such as `--help`, do not.
// The negated form comes from the long name of the flag. An alias and a
// short name get none. A flag whose name starts with a prefix and a hyphen,
// such as `--no-psqlrc`, gets none, so that no flag reads `--no-no-psqlrc`.
// Help writes the flag as `--[no-]debug`, with the first prefix, and
// completion offers each form.
//
// The negated form of a hidden flag is hidden, the negated form of a local
// flag is local (see [Local]), and a sub command takes the negated form of
// each flag that it inherits. A negated form with the name of another flag
// of a command, or of its long alias, is an error when the command is
// built. Give one of the flags [Negatable] false to keep it out. The
// environment and the config files set the flag itself, and are not
// changed.
//
// NegateBools is not one of the [Defaults], because it changes how the
// command line parses.
func NegateBools(prefixes ...string) CommandOption {
	return option{
		name: "NegateBools",
		cmd: func(cmd *Command) error {
			if cmd.Parent != nil {
				return ErrCanOnlyBeUsedWithRootCommand
			}
			v := prefixes
			if len(v) == 0 {
				v = []string{DefaultNegatePrefix}
			}
			for i, p := range v {
				switch {
				case !validPrefix(p):
					return fmt.Errorf("%w: prefix %q", ErrInvalidValue, p)
				case slices.Contains(v[:i], p):
					return fmt.Errorf("%w: prefix %q is given twice", ErrInvalidValue, p)
				}
			}
			cmd.NegateBools = v
			return nil
		},
	}
}

// validPrefix returns true when p is a prefix that [NegateBools] takes: a
// word of letters, digits and hyphens, with no hyphen at either end.
func validPrefix(p string) bool {
	if p == "" || p[0] == '-' || p[len(p)-1] == '-' {
		return false
	}
	for _, r := range p {
		if r != '-' && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Negatable is a [Flag] option to choose whether a bool flag has a negated
// form, such as `--no-debug`. Negatable(false) keeps the flag out of
// [NegateBools]. Negatable(true) gives the flag a negated form without
// [NegateBools], with the prefix of [DefaultNegatePrefix]. With
// [NegateBools], the flag takes the prefixes of the root command. A flag
// that is not a bool flag refuses Negatable(true) with
// [ErrAppliedToInvalidType].
func Negatable(negatable bool) FlagOption {
	return option{
		name: "Negatable",
		flag: func(g *Flag) error {
			g.Negatable, g.NoNegate = negatable, !negatable
			return nil
		},
	}
}

// AllowPlus is a [Flag] option that lets the short name of the flag take a
// `+` as well as a `-`, as bash takes `+o name` to turn an option off. The
// flag needs a short name, and a count refuses it. The plus form of a plain
// bool flag, such as `+e`, sets it to false, and takes no value. The plus form
// of any other flag takes its value as the `-` form does, such as `+o name`,
// and a hook of the flag reads the sign with [Context.IsPlus]. Hooks run in
// the order of the command line, so a program that needs the order of each
// `-` and `+`, as a shell does, gives each option a hook.
//
// A word such as `+ex` is read as flags only when its first letter names a
// flag with AllowPlus. Each letter must then name such a flag. Any other word
// that starts with `+`, such as the `+%Y` of `date`, stays an argument. The
// struct tag option is `plus`.
func AllowPlus() FlagOption {
	return option{
		name: "AllowPlus",
		flag: func(g *Flag) error {
			g.AllowPlus = true
			return nil
		},
	}
}

// OptionalAtEnd is a [Flag] option that lets a flag that takes a value be
// the last word of the command line without one, as bash takes a lone `-o`
// to list its options. The flag is set from an empty value, and its hook
// reads [Context.NoValue] as true. Anywhere else, the flag takes its value as
// before. A flag that takes no value refuses it with
// [ErrAppliedToInvalidType].
func OptionalAtEnd() FlagOption {
	return option{
		name: "OptionalAtEnd",
		flag: func(g *Flag) error {
			g.OptionalAtEnd = true
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

// TypeFor is a [Flag] option that sets the type of the flag from the Go type
// T, as [From] does for the type of a field. T can be a type that reads its
// own text, with UnmarshalText, UnmarshalBinary or Set, or a pointer to one,
// and it needs no registration. Help writes the name of such a type in lower
// case, so a flag of type colors.Color shows `--fg color`. T can also be a
// type that [From] takes, such as int, []string or map[string]int.
//
//	ox.Flags().
//		Var("fg", "the foreground color", ox.TypeFor[colors.Color]())
func TypeFor[T any]() FlagOption {
	return option{
		name: "TypeFor",
		flag: func(g *Flag) error {
			refType := reflect.TypeFor[T]()
			typ, mapKey, elem, err := defaultType(refType)
			if err != nil {
				return fmt.Errorf("%s: %w", refType, err)
			}
			g.Type, g.MapKey, g.Elem = typ, mapKey, elem
			return nil
		},
	}
}

// BindRef is a [Flag] option to bind a reflected value to a flag.
func BindRef(value reflect.Value, set *bool) FlagOption {
	return option{
		name: "BindRef",
		flag: func(g *Flag) error {
			val, err := NewBindRef(value, set)
			if err != nil {
				return err
			}
			g.Binds = append(g.Binds, val)
			return nil
		},
	}
}

// Bind is a [Flag] option to bind a variable to a flag.
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

// BindSet is a [Flag] option to bind a variable and a set variable to a flag.
func BindSet[T *E, E any](v T, set *bool) FlagOption {
	return option{
		name: "BindSet",
		flag: func(g *Flag) error {
			val, err := NewBind(v, set)
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
// Special strings can be used that will be expanded when the flag is created:
//
//	$HOME - the current user's home directory
//	$USER - the current user's user name
//	$CACHE - the current user's cache directory
//	$STATE - the current user's state directory
//	$APPCACHE - the current user's cache directory, with the root command's name added as a subdir
//	$APPSTATE - the current user's state directory, with the root command's name added as a subdir
//	$ENV{KEY} - the environment value for $KEY
//	$CFG{[TYPE::]KEY} - the registered config file loader type and key value
//
// For example:
//
//	ox.Default("$USER") - expands to "ken" if the user running the application is "ken"
//	ox.Default("$HOME") - expands to /home/$USER
//	ox.Default("$CACHE") - expands to /home/$USER/.cache
//	ox.Default("$STATE") - expands to /home/$USER/.local/state
//	ox.Default("$APPCACHE") - expands to /home/$USER/.cache/myApp if the root command's name is "myApp"
//	ox.Default("$APPSTATE") - expands to /home/$USER/.local/state/myApp if the root command's name is "myApp"
//	ox.Default("~/.${APPNAME|lower}rc") - expands to ~/.myapprc if the root command's name is "myApp"
//	ox.Default("$ENV{MY_VAR}") - expands to value of the environment var $MY_VAR
//	ox.Default("$CFG{yaml::a.b.c}") - expands to the registered YAML config file's key of a.b.c
//	ox.Default("$CFG{a.b.c}") - expands to the first registered config file that returns a non-nil value and nil error for key a.b.c
//
// See [Context.Interpolate] for further details.
func Default(def any) FlagOption {
	return option{
		name: "Default",
		flag: func(g *Flag) error {
			g.Def = def
			return nil
		},
	}
}

// NoDefault is a [Flag] option to hide the "(default: ...)" text of the flag
// in help. Use it when the usage of the flag already names its default.
func NoDefault() FlagOption {
	return option{
		name: "NoDefault",
		flag: func(g *Flag) error {
			g.NoDefault = true
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

// OptionalValue is a [Flag] option to set that the flag takes an optional
// value, and the spec of that value. Help writes the spec in brackets, such
// as `--help[=options]`. The flag takes no argument, as with [NoArg], so a
// value must follow an `=`, such as `--help=commands`. It keeps the default
// that [NoArg] sets, and sets the empty string when there is none.
func OptionalValue(spec string) FlagOption {
	return option{
		name: "OptionalValue",
		flag: func(g *Flag) error {
			g.Spec, g.Optional, g.NoArg = spec, true, true
			if g.NoArgDef == nil {
				g.NoArgDef = ""
			}
			return nil
		},
	}
}

// HelpValue is a [Flag] option to add a value of the flag that help writes on
// a line of its own, with its usage, such as psql's `--help=commands`. Help
// writes the values after the flag, in the order of the options.
//
// When the command line gives the flag this value, each func in funcs runs in
// order, in place of the flag, and the run stops as help does. The first
// error of a func stops the run with that error. So `--help=commands` can
// write a topic of its own to [Context.Stdout], and `--help` still writes the
// help of the command. Without funcs, the flag takes the value as it takes
// any other value.
func HelpValue(value, usage string, funcs ...func(*Context) error) FlagOption {
	return option{
		name: "HelpValue",
		flag: func(g *Flag) error {
			g.HelpValues = append(g.HelpValues, HelpLine{Name: value, Usage: usage, Funcs: funcs})
			return nil
		},
	}
}

// Abs is a [Flag] option to make the value of a path flag absolute. A
// relative path is joined to the working directory.
func Abs() FlagOption {
	return option{
		name: "Abs",
		flag: func(g *Flag) error {
			g.Abs = true
			return nil
		},
	}
}

// Real is a [Flag] option to resolve each symbolic link in the value of a
// path flag, with [Realpath]. A path that does not exist is made absolute, as
// with [Abs].
func Real() FlagOption {
	return option{
		name: "Real",
		flag: func(g *Flag) error {
			g.Real = true
			return nil
		},
	}
}

// Home is a [Flag] option to show the default of a path flag under the home
// directory as `~` in help, such as `~/.config/app`. It does not change the
// value.
func Home() FlagOption {
	return option{
		name: "Home",
		flag: func(g *Flag) error {
			g.Home = true
			return nil
		},
	}
}

// Ext is a [Flag] option to set the file extensions that completion offers
// for a path flag, such as `Ext(".yaml", ".yml")`. bash and zsh complete only
// the files with these extensions, and fish and PowerShell complete every
// file. It does not refuse a value with another extension.
func Ext(exts ...string) FlagOption {
	return option{
		name: "Ext",
		flag: func(g *Flag) error {
			g.Exts = append(g.Exts, exts...)
			return nil
		},
	}
}

// NotExist is a [Flag] option to make the value of a path flag a path that
// must not exist yet, such as an output file.
func NotExist() FlagOption {
	return option{
		name: "NotExist",
		flag: func(g *Flag) error {
			g.NotExist = true
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

// ConfigFlag is a [Flag] option to mark a flag whose value names the config
// file. When the flag has a value, its file replaces the files of
// [ConfigFiles], and the file must exist. Its default can read the
// environment, such as "${env::APP_CONFIG}", and not a config file.
func ConfigFlag() FlagOption {
	return option{
		name: "ConfigFlag",
		flag: func(g *Flag) error {
			g.ConfigFile = true
			return nil
		},
	}
}

// EnvVar is a [Flag] option to fill the flag from the environment variable
// name, whatever the prefix of [EnvPrefix]. A hook flag is filled from the
// environment only with EnvVar.
func EnvVar(name string) FlagOption {
	return option{
		name: "EnvVar",
		flag: func(g *Flag) error {
			g.EnvVar = name
			return nil
		},
	}
}

// NoEnv is a [Flag] option to keep the flag out of the environment.
func NoEnv() FlagOption {
	return option{
		name: "NoEnv",
		flag: func(g *Flag) error {
			g.NoEnv = true
			return nil
		},
	}
}

// NoConfig is a [Flag] option to keep the flag out of the config files.
func NoConfig() FlagOption {
	return option{
		name: "NoConfig",
		flag: func(g *Flag) error {
			g.NoConfig = true
			return nil
		},
	}
}

// Required is a [Flag] option to make the parse fail when the flag gets no
// value from the command line, the environment or a config file. A default
// does not count. [Context.Validate] returns an error that wraps
// [ErrRequiredFlag] and names each flag that has no value. Help marks the flag
// with the text of `text.FlagRequired`. The struct tag option is `required`.
func Required() FlagOption {
	return option{
		name: "Required",
		flag: func(g *Flag) error {
			g.Required = true
			return nil
		},
	}
}

// groupOption is the [Flag] option of a struct tag that adds the flag to the
// group name, such as `onlyoneof:format`.
func groupOption(rule Rule, name string) FlagOption {
	return option{
		name: "group",
		flag: func(g *Flag) error {
			if name == "" {
				return fmt.Errorf("%w: empty group name", ErrInvalidValue)
			}
			g.groups = append(g.groups, groupTag{
				rule: rule,
				name: name,
			})
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

// Valid is a [Flag] option to add a allowed value validator to a flag.
func Valid[T cmp.Ordered](values ...T) FlagOption {
	return option{
		name: "Valid",
		flag: func(g *Flag) error {
			g.Valid = append(g.Valid, valid(values...))
			for _, val := range values {
				g.Allowed = append(g.Allowed, val)
			}
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

// Hidden is a [Command]/[Flag] option to set the command/flag's hidden.
func Hidden(hidden bool) CommandFlagOption {
	return option{
		name: "Hidden",
		cmd: func(cmd *Command) error {
			cmd.Hidden = hidden
			return nil
		},
		flag: func(g *Flag) error {
			g.Hidden = hidden
			return nil
		},
	}
}

// Deprecated is a [Command]/[Flag] option to set the command/flag's deprecated.
func Deprecated(deprecated bool) CommandFlagOption {
	return option{
		name: "Deprecated",
		cmd: func(cmd *Command) error {
			cmd.Deprecated = deprecated
			return nil
		},
		flag: func(g *Flag) error {
			g.Deprecated = deprecated
			return nil
		},
	}
}

// Local is a [Flag] option to keep the flag to its command. A sub command
// does not inherit a local flag, so the command line of a sub command refuses
// it with [ErrUnknownFlag], and completion does not offer it there. A flag
// that is not local works after each sub command, as before.
//
// Local changes only the command line. The environment, a config file or the
// default still give the flag its value, as [Context.Fill] does for each flag
// of the command and its parents.
func Local() FlagOption {
	return option{
		name: "Local",
		flag: func(g *Flag) error {
			g.Local = true
			return nil
		},
	}
}

// Hook is a [Flag] option to set a hook for a flag. The flag takes no
// argument. When the flag is on the command line, the parse calls f with a
// [context.Context] that holds the [Context]. Read it with [Ctx]. The parent
// of that [context.Context] is [DefaultContext], because the parse runs
// before [Context.Run].
func Hook(f func(context.Context) error) FlagOption {
	return option{
		name: "Hook",
		flag: func(g *Flag) error {
			g.Type, g.Def, g.NoArg, g.NoArgDef = HookT, f, true, ""
			return nil
		},
	}
}

// Banner is a [Help] option to set the banner.
func Banner(banner string) HelpOption {
	return option{
		name: "Banner",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.Banner, help.NoBanner = banner, banner == ""
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
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
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.Spec, help.NoSpec = spec, spec == ""
			}
			return nil
		},
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
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.Footer, help.NoFooter = footer, footer == ""
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.Footer, help.NoFooter = footer, footer == ""
			return nil
		},
	}
}

// Example is a [Help] option to set the example.
func Example(example string) HelpOption {
	return option{
		name: "Example",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.Example, help.NoExample = example, example == ""
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.Example, help.NoExample = example, example == ""
			return nil
		},
	}
}

// Sort is a [Command] option to set the sort for flags when used with a
// [Command].
func Sort(sort bool) CommandOption {
	return option{
		name: "Sort",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.Sort = sort
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.Sort = sort
			return nil
		},
	}
}

// CommandSort is a [Command] option to set the sort for sub commands when used
// with a [Command].
func CommandSort(commandSort bool) CommandOption {
	return option{
		name: "CommandSort",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.CommandSort = commandSort
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.CommandSort = commandSort
			return nil
		},
	}
}

// MaxDist is a [Command] option to set the maximum Levenshtein distance of the
// suggestions and the corrections of the command. 0 keeps [DefaultMaxDist], and
// a negative distance turns off the distance match for the command.
func MaxDist(maxDist int) CommandOption {
	return option{
		name: "MaxDist",
		cmd: func(cmd *Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.MaxDist = maxDist
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.MaxDist = maxDist
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
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.CommandSections = sections
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.Sections = sections
			return nil
		},
	}
}

// ArgUsage is a [Help] option to add the name and the usage of a positional
// argument, such as `DSN`. Help writes each one in an `Arguments:` section
// after the usage, in the order of the options.
func ArgUsage(name, usage string) HelpOption {
	f := func(help *CommandHelp) {
		help.Args = append(help.Args, HelpLine{Name: name, Usage: usage})
	}
	return option{
		name: "ArgUsage",
		cmd: func(*Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				f(help)
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			f(help)
			return nil
		},
	}
}

// HelpColumn is a [Command]/[Help] option to start each description in help
// after column characters, as `psql --help` does. An option or a command that
// does not fit before the column goes on a line of its own, and its
// description starts on the next line. A description is not wrapped, and each
// line break in it starts a new line at the column. Help also writes the
// aliases of each flag. Without HelpColumn, or with a column of zero or less,
// help uses the default layout.
func HelpColumn(column int) CommandOption {
	return option{
		name: "HelpColumn",
		cmd: func(*Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.Column = column
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.Column = column
			return nil
		},
	}
}

// InheritedFlags is a [Command]/[Help] option to show the flags that a sub
// command inherits from its parents in its help, after its own flags, in a
// section with the name of [text.GlobalFlags], as cobra does. A local flag
// of a parent is not inherited (see [Local]), and a flag of a parent with the
// name of a flag of the sub command is not shown. When the root command has
// it, [Help] gives it to each sub command.
func InheritedFlags(show bool) CommandOption {
	return option{
		name: "InheritedFlags",
		cmd: func(*Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.InheritedFlags = show
			}
			return nil
		},
		flag: func(g *Flag) error {
			if strings.HasPrefix(g.Special, `hook:`) {
				return ErrAppliedToInvalidType
			}
			return nil
		},
		help: func(help *CommandHelp) error {
			help.InheritedFlags = show
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
				if after, ok := strings.CutPrefix(k, "flag:"); ok {
					if g := cmd.Flag(after, false, false); g != nil {
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
	last func(*Command) error
	flag func(*Flag) error
	help func(*CommandHelp) error
	typ  func(v interface{ SetType(Type) }) error
	arg  func(*Argument) error
	// setType is the type of a [Type] option, which has no funcs.
	setType Type
}

// Option satisfies the [Option] interface.
func (opt option) Option() option {
	return opt
}

// apply applies the option to val.
func (opt option) apply(val any) error {
	if opt.setType != "" {
		return opt.setType.applyType(val)
	}
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
		if opt.cmd != nil || opt.post != nil || opt.last != nil {
			err = nil
		}
		if opt.ctx != nil {
			err = opt.ctx(v)
		}
	case *CommandHelp:
		if opt.help != nil {
			err = opt.help(v)
		}
	case *Argument:
		if opt.arg != nil {
			err = opt.arg(v)
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

// Apply applies the options to v.
func Apply(v any, opts ...Option) error {
	for _, o := range opts {
		if err := o.Option().apply(v); err != nil {
			return err
		}
	}
	return nil
}

// ApplyPost applies post options to the command, then the funcs of each
// [Extend] option.
func ApplyPost(cmd *Command, opts ...Option) error {
	// each Option call can build its funcs again, so call it once
	var last []option
	for _, o := range opts {
		opt := o.Option()
		if opt.post != nil {
			if err := opt.post(cmd); err != nil {
				return err
			}
		}
		if opt.last != nil {
			last = append(last, opt)
		}
	}
	for _, opt := range last {
		if err := opt.last(cmd); err != nil {
			return fmt.Errorf("%s: %w", opt.name, err)
		}
	}
	return nil
}
