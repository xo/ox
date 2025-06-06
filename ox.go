// Package ox is a minimal dependency Go (and TinyGo compatible) command-line
// flag and argument parsing library.
package ox

//go:generate stringer -type OnErr

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/xo/ox/strcase"
	"github.com/xo/ox/text"
)

var (
	// DefaultContext is the default [context.Context].
	DefaultContext = context.Background()
	// DefaultStructTagName is the default struct tag name used in
	// [From]/[FlagsFrom] and other funcs.
	DefaultStructTagName = "ox"
	// DefaultLayout is the default timestamp layout used for formatting and
	// parsing [Time] values.
	DefaultLayout = time.RFC3339
	// DefaultStripGoTestFlags when enabled strips Go's `-test.` prefix'd
	// flags.
	DefaultStripGoTestFlags = true
	// DefaultWrapWidth is the default wrap width.
	DefaultWrapWidth = 95
	// DefaultWrapTo wraps a line of text, split into sections by "\n\n", using
	// [Wrap], emitting the output to a writer.
	DefaultWrapTo = func(w io.StringWriter, s string, wrapWidth, prefixWidth int) {
		for i, line := range strings.Split(s, "\n\n") {
			if i != 0 {
				_, _ = w.WriteString("\n\n")
				_, _ = w.WriteString(strings.Repeat(" ", prefixWidth))
			}
			_, _ = w.WriteString(Wrap(line, wrapWidth, prefixWidth))
		}
	}
	// DefaultWidth returns the width of a string used by [Wrap].
	DefaultWidth = func(s string) int {
		return len(s)
	}
	// DefaultFlagNameMaper is the default flag name mapper.
	DefaultFlagNameMapper = func(s string) string {
		return strings.ReplaceAll(strcase.CamelToSnake(s), "_", "-")
	}
	// DefaultVersionTrimPrefix is used by [DefaultVersionMapper] to trim the
	// `v` prefix on build version.
	DefaultVersionTrimPrefix = true
	// DefaultVersion is the default version func.
	DefaultVersion = func(ctx *Context) error {
		var name string
		if ctx != nil && ctx.Root != nil {
			name = ctx.Root.Name
		}
		var ver string
		switch {
		case ctx != nil && ctx.Exec != nil && ctx.Exec.Version != "":
			ver = ctx.Exec.Version
		case ctx != nil && ctx.Root != nil && ctx.Root.Version != "":
			ver = ctx.Root.Version
		default:
			ver = BuildVersion()
		}
		name, ver = DefaultVersionMapper(name, ver)
		var w io.Writer
		if ctx != nil {
			w = ctx.Stdout
		}
		if w == nil {
			w = os.Stdout
		}
		_, _ = fmt.Fprintln(w, name, ver)
		return ErrExit
	}
	// DefaultVersionMapper maps the passed name, version.
	DefaultVersionMapper = func(name, ver string) (string, string) {
		if name == "" {
			name = filepath.Base(os.Args[0])
			if runtime.GOOS == "windows" && strings.HasSuffix(strings.ToLower(name), ".exe") {
				name = name[:len(name)-4]
			}
		}
		if DefaultVersionTrimPrefix && regexp.MustCompile(`^v[0-9]+\.`).MatchString(ver) {
			ver = strings.TrimPrefix(ver, "v")
		}
		return name, ver
	}
	// DefaultContinueHandler is the default continue handler used by
	// [Run]/[RunContext].
	DefaultContinueHandler = func(ctx *Context) func(*Command, error) bool {
		var onErr OnErr
		if ctx != nil {
			onErr = ctx.OnErr
		}
		return func(cmd *Command, err error) bool {
			if cmd != nil {
				onErr = cmd.OnErr
			}
			return err == nil || onErr == OnErrContinue
		}
	}
	// DefaultErrorHandler is the default error handler used by
	// [Run]/[RunContext].
	DefaultErrorHandler = func(ctx *Context) func(error) bool {
		return func(err error) bool {
			switch {
			case errors.Is(err, ErrExit):
			case ctx.OnErr == OnErrContinue:
				return false
			case ctx.OnErr == OnErrExit:
				var w io.Writer = os.Stderr
				if ctx.Stderr != nil {
					w = ctx.Stderr
				}
				fmt.Fprintf(w, text.ErrorMessage, err)
				if v, ok := err.(interface{ ErrorDetails() string }); ok {
					fmt.Fprint(w, v.ErrorDetails())
				}
				code := 1
				if v, ok := err.(interface{ ErrorCode() int }); ok {
					code = v.ErrorCode()
				}
				ctx.Exit(code)
			case ctx.OnErr == OnErrPanic:
				ctx.Panic(err)
			}
			return true
		}
	}
	// DefaultSuggestionsEnabled sets whether command name suggestions are
	// given.
	DefaultSuggestionsEnabled = true
	// DefaultMaxDist is the default maximum distance for suggestions.
	DefaultMaxDist = 2
	// DefaultFlagWidth is the default minimum flag width. Does not include the
	// padding.
	DefaultFlagWidth = 8
	// DefaultCommandWidth is the default minimum command width.
	DefaultCommandWidth = 6
	// DefaultCompTemplates are the default completion templates.
	DefaultCompTemplates = templates
	// DefaultCompName is the name for the completion script hook.
	DefaultCompName = "__complete"
	// DefaultCompNameNoDesc is the name for the completion script hook without
	// descriptions.
	DefaultCompNameNoDesc = "__completeNoDesc"
	// DefaultCompEnvNoDescName is the name for the environment variable for
	// disabling completion descriptions.
	DefaultCompEnvNoDescName = "COMPLETION_DESCRIPTIONS"
	// DefaultCompActiveHelpSuffix is the active help suffix.
	DefaultCompActiveHelpSuffix = "_ACTIVE_HELP"
	// DefaultCompWrite is the default completion write func, that writes the
	// completion script from the passed template.
	DefaultCompWrite = func(ctx *Context, cmd *Command, noDescriptions bool, shell, templ string) error {
		rootName := cmd.RootName()
		varName := regexp.MustCompile(`[^A-Za-z0-9_]`).ReplaceAllString(rootName, "_")
		activeHelp := strings.ToUpper(varName) + DefaultCompActiveHelpSuffix
		compName := DefaultCompName
		if noDescriptions {
			compName = DefaultCompNameNoDesc
		}
		_, _ = fmt.Fprintf(
			ctx.Stdout,
			templ,
			rootName,
			varName,
			compName,
			activeHelp,
		)
		return ErrExit
	}
	// DefaultComp is the default completion func.
	DefaultComp = func(ctx *Context) error {
		noDescriptions := ctx.Args[0] == DefaultCompNameNoDesc
		for _, name := range []string{ctx.Root.Name + "_" + DefaultCompEnvNoDescName, DefaultCompEnvNoDescName} {
			if s := os.Getenv(name); s != "" {
				if b, err := asBool(s); err == nil {
					noDescriptions = !b
				}
			}
		}
		comps, dir, err := ctx.Comps()
		if err != nil {
			ctx.Handler(err)
			ctx.Exit(1)
		}
		for _, comp := range comps {
			if !noDescriptions {
				_, _ = fmt.Fprintln(ctx.Stdout, comp.Name+"\t"+comp.Usage)
			} else {
				_, _ = fmt.Fprintln(ctx.Stdout, comp.Name)
			}
		}
		_, _ = fmt.Fprintf(ctx.Stdout, ":%d\n", dir)
		_, _ = fmt.Fprintf(ctx.Stderr, "COMP ENDED: %s\n", dir)
		return ErrExit
	}
	// DefaultLoader is the default config loader.
	DefaultLoader = func(ctx *Context, typ, key string) (any, bool, error) {
		if loader, ok := loaders[typ]; ok {
			v, err := loader(ctx, key)
			return v, true, err
		}
		return nil, false, nil
	}
	// DefaultStripTestFlags strips flags starting with `-test.` from args.
	DefaultStripTestFlags = func(args []string) []string {
		if !DefaultStripGoTestFlags {
			return args
		}
		var v []string
		for i := 0; i < len(args); i++ {
			switch hasPrefix := strings.HasPrefix(args[i], "-test."); {
			case hasPrefix && !strings.Contains(args[i], "="):
				i++
				continue
			case hasPrefix:
				continue
			}
			v = append(v, args[i])
		}
		return v
	}
	// DefaultSizePrec is the default [Size] display precision.
	DefaultSizePrec = -2
	// DefaultRatePrec is the default [Rate] display precision.
	DefaultRatePrec = -2
	// DefaultRateUnit is the default [Rate] unit used to display the rate, if
	// none is specified or when precision is negative.
	DefaultRateUnit = time.Second
	// DefaultMaxPathLinks are the max symlinks to follow in [Path].
	DefaultMaxPathLinks = 16
	// DefaultGetwd is the func used to read the working directory by [Path].
	// Normally [os.Getwd].
	DefaultGetwd = os.Getwd
	// DefaultLstat is the func used to lstat a path by [Path]. Normally
	// [os.Lstat].
	DefaultLstat = os.Lstat
	// DefaultReadlink is the func used to read a link by [Path]. Normally
	// [os.Readlink].
	DefaultReadlink = os.Readlink
)

// Run creates and builds the execution [Context] based on the passed
// [Option]s. Accepts any [ContextOption], [CommandOption] or
// [CommandFlagOption]. After building a execution [Context] and root
// [Command], either [os.Args] or the provided [Args] will be [Parse]'d and
// validated.
//
// After args have been parsed and validated, [Context] will have its
// [Context.Exec] set to either the root command, or to the user's requested
// command. The determined [Command.Exec] will be executed, or if it was not
// set, then the [Command.Help]'s will be written to the [Context.Stdout].
//
// See [RunContext] to pass a [context.Context] to the executed command, or set
// the [DefaultContext].
func Run(opts ...Option) {
	RunContext(DefaultContext, opts...)
}

// RunContext creates a [Context] and builds a execution [Context] and root
// [Command] based on the passed options. Accepts any [ContextOption],
// [CommandOption] or [CommandFlagOption]. After building a execution [Context]
// and root [Command], either [os.Args] or the provided [Args] will be
// [Parse]'d and validated.
//
// After args have been parsed and validated, [Context] will have its
// [Context.Exec] set to either the root command, or to the user's requested
// command. The determined [Command.Exec] will be executed, or if it was not
// set, then the [Command.Help]'s will be written to the [Context.Stdout].
func RunContext(ctx context.Context, opts ...Option) {
	// build context
	c, err := NewContext(opts...)
	if err != nil && c.Handler(err) {
		return
	}
	// parse
	if err = c.Parse(); err != nil && c.Handler(err) {
		return
	}
	// validate
	if err = c.Validate(); err != nil && c.Handler(err) {
		return
	}
	// execute
	if err = c.Run(ctx); err != nil && c.Handler(err) {
		return
	}
}

// Context is the execution (run) context used by [Run]/[RunContext].
type Context struct {
	// Exit is the exit func.
	Exit func(int)
	// Panic is the panic func.
	Panic func(any)
	// Comp is the completion func.
	Comp func(*Context) error
	// Loader is the config loader func.
	Loader func(*Context, string, string) (any, bool, error)
	// Stdin is the standard in to use, normally [os.Stdin].
	Stdin io.Reader
	// Stdout is the standard out to use, normally [os.Stdout].
	Stdout io.Writer
	// Stderr is the standard error to use, normally [os.Stderr].
	Stderr io.Writer
	// Args are the arguments to parse, normally [os.Args][1:].
	Args []string
	// OnErr is the on error handling. Used by the default Handle func.
	OnErr OnErr
	// Continue is the func used to decide if it should continue on an error.
	Continue func(*Command, error) bool
	// Handler is the func used to handle errors within [Run]/[RunContext].
	Handler func(error) bool
	// Override override expansion keys.
	Override func(string, string) (string, bool)
	// Root is the root command created within [Run]/[RunContext].
	Root *Command
	// Exec is the exec target, determined by the Root's definition and after
	// Args.
	Exec *Command
	// Vars are the variables parsed from the flag definitions of the Root
	// command and its sub-commands.
	Vars Vars
}

// NewContext creates a new run context.
func NewContext(opts ...Option) (*Context, error) {
	ctx := &Context{
		Exit: os.Exit,
		Panic: func(v any) {
			panic(v)
		},
		Comp:   DefaultComp,
		Loader: DefaultLoader,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Args:   DefaultStripTestFlags(os.Args[1:]),
	}
	ctx.Continue = DefaultContinueHandler(ctx)
	ctx.Handler = DefaultErrorHandler(ctx)
	if err := Apply(ctx, opts...); err != nil {
		return ctx, err
	}
	if ctx.Vars == nil {
		ctx.Vars = make(Vars)
	}
	if ctx.Root == nil {
		var err error
		if ctx.Root, err = NewCommand(opts...); err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}

// Parse parses the context args.
func (ctx *Context) Parse() error {
	if ctx.Root.Comp && len(ctx.Args) != 0 && strings.HasPrefix(ctx.Args[0], DefaultCompName) {
		return ctx.Comp(ctx)
	}
	var err error
	if err = ctx.Populate(ctx.Root, false, false); err != nil {
		return newCommandError(ctx.Root.Name, err)
	}
	ctx.Exec, ctx.Args, err = Parse(ctx, ctx.Root, ctx.Args)
	return err
}

// Comps returns the completions for the context.
func (ctx *Context) Comps() ([]Completion, CompDirective, error) {
	n := len(ctx.Args)
	if n < 2 {
		return nil, CompError, ErrInvalidArgCount
	}
	n--
	c := &Context{
		Root:   ctx.Root,
		Stdout: io.Discard,
		Stderr: io.Discard,
		Panic:  func(any) {},
		Exit:   func(int) {},
		Continue: func(cmd *Command, err error) bool {
			fmt.Fprintf(ctx.Stderr, "COMP CONTINUE ERR: %v\n", err)
			switch {
			case errors.Is(err, ErrUnknownFlag),
				errors.Is(err, ErrMissingArgument),
				errors.Is(err, ErrExit):
				return true
			}
			return false
		},
		Args: ctx.Args[1:n],
		Vars: make(Vars),
	}
	fmt.Fprintf(ctx.Stderr, "COMP ARGS: %v -- %s\n", ctx.Args[1:n], ctx.Args[n])
	if err := c.Parse(); err != nil {
		fmt.Fprintf(ctx.Stderr, "COMP PARSE ERR: %v\n", err)
		return nil, CompError, err
	}
	fmt.Fprintf(ctx.Stderr, "COMP COMMAND: %s\n", c.Exec.Name)
	var comps []Completion
	var dir CompDirective
	// TODO: expose flags to allow hidden/deprecated
	hidden, deprecated := false, true
	// build completions
	switch prev, arg := prev(ctx.Args, n), ctx.Args[n]; {
	case strings.HasPrefix(arg, "-"):
		comps, dir = c.Exec.CompFlags(strings.TrimLeft(arg, "-"), hidden, deprecated, !strings.HasPrefix(arg, "--"))
	case strings.HasPrefix(prev, "-"):
		// TODO: logic incorrect; need to strip the `-`/`--` and `=` from the flag
		switch g := c.Exec.Flag(prev, true, len([]rune(prev)) == 2); {
		case g == nil, g.NoArg, strings.Contains(prev, "="):
			comps, dir = c.Exec.CompCommands(arg, hidden, deprecated)
		default:
			// TODO: handle completion for flags with completion definitions
		}
	default:
		comps, dir = c.Exec.CompCommands(arg, hidden, deprecated)
	}
	return comps, dir, nil
}

// Validate validates the args.
func (ctx *Context) Validate() error {
	return ctx.Exec.Validate(ctx.Args)
}

// Run runs the command.
func (ctx *Context) Run(parent context.Context) error {
	switch {
	case ctx.Exec.Exec != nil:
		return ctx.Exec.Exec(WithContext(parent, ctx), ctx.Args)
	case DefaultSuggestionsEnabled && len(ctx.Exec.Commands) != 0 && len(ctx.Args) != 0:
		if err := ctx.Exec.Suggest(ctx.Args...); err != nil {
			return err
		}
	}
	_, _ = ctx.Exec.HelpContext(ctx).WriteTo(ctx.Stdout)
	return nil
}

// Populate populates the context's vars with the command's flag's default
// values when not already set. When all is true, all flag values will be
// populated with the default or zero value. When overwrite is true, any
// existing flag vars will be overwritten.
func (ctx *Context) Populate(cmd *Command, all, overwrite bool) error {
	if cmd.Flags == nil {
		return nil
	}
	for _, g := range cmd.Flags.Flags {
		if _, ok := ctx.Vars[g.Name]; ok && overwrite {
			delete(ctx.Vars, g.Name)
		}
		var value string
		switch {
		case g.Type == HookT, g.Def == nil && !all:
			continue
		case g.Def != nil:
			v, err := ctx.Expand(g.Def)
			if err != nil {
				return fmt.Errorf("populate %s: cannot expand %v (%T): %w", g.Name, g.Def, g.Def, err)
			}
			value, _ = asString[string](v)
		}
		if err := ctx.Vars.Set(ctx, g, value, false); err != nil {
			return fmt.Errorf("populate %s: %w", g.Name, err)
		}
	}
	return nil
}

// Expand expands variables in v, where any variable can be the following:
//
//	$HOME - the current user's home directory (ex: ~/)
//	$USER - the current user's user name (ex: user)
//	$APPNAME - the root command's name (ex: appName)
//	$CONFIG - the current user's config directory (ex: ~/.config)
//	$APPCONFIG - the current user's config directory, with the root command's name added as a subdir (ex: ~/.config/appName)
//	$CACHE - the current user's cache directory (ex: ~/.cache)
//	$APPCACHE - the current user's cache directory, with the root command's name added as a subdir (ex: ~/.cache/appName)
//	$NUMCPU - the value of [runtime.NumCPU] (ex: 4)
//	$NUMCPU2 - the value of [runtime.NumCPU], plus 2 (ex: 6)
//	$NUMCPU2X - the value of [runtime.NumCPU], times 2 (ex: 8)
//	$ARCH - the value of [runtime.GOARCH] (ex: amd64)
//	$OS - the value of [runtime.GOOS] (ex: windows)
//	$ENV{KEY} - the environment value for $KEY
//	$CONFIG_TYPE{KEY} - the registered config file loader type and key value, for example: `$YAML{my_key}`, `$TOML{my_key}`
func (ctx *Context) Expand(v any) (any, error) {
	s, ok := v.(string)
	if !ok {
		return v, nil
	}
	b := []byte(s)
	matches := keyRE.FindAllSubmatchIndex(b, -1)
	slices.Reverse(matches)
	for _, m := range matches {
		typ, key := string(b[m[2]:m[3]]), ""
		if m[4] != -1 {
			key = string(b[m[4]:m[5]])
		}
		switch z, ok, err := ctx.ExpandKey(typ, key); {
		case err != nil:
			return nil, err
		case ok:
			str, err := asString[string](z)
			if err != nil {
				return nil, err
			}
			b = slices.Replace(b, m[0], m[1], []byte(str)...)
		}
	}
	return string(b), nil
}

// ExpandKey expands a config key, using the context Loader.
//
// See [Expand] for more information.
func (ctx *Context) ExpandKey(typ, key string) (any, bool, error) {
	if ctx.Override != nil {
		if s, ok := ctx.Override(typ, key); ok {
			return s, true, nil
		}
	}
	var f func() (string, error)
	switch typ {
	case "HOME":
		f = userHomeDir
	case "USER":
		f = func() (string, error) {
			u, err := user.Current()
			if err != nil {
				return "", err
			}
			return u.Username, nil
		}
	case "APPNAME":
		return ctx.Root.Name, true, nil
	case "CONFIG":
		f = userConfigDir
	case "APPCONFIG":
		f = func() (string, error) {
			dir, err := userConfigDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ctx.Root.Name), nil
		}
	case "CACHE":
		f = userCacheDir
	case "APPCACHE":
		f = func() (string, error) {
			dir, err := userCacheDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ctx.Root.Name), nil
		}
	case "NUMCPU":
		return strconv.Itoa(runtime.NumCPU()), true, nil
	case "NUMCPU2":
		return strconv.Itoa(runtime.NumCPU() + 2), true, nil
	case "NUMCPU2X":
		return strconv.Itoa(runtime.NumCPU() * 2), true, nil
	case "ARCH":
		return runtime.GOARCH, true, nil
	case "OS":
		return runtime.GOOS, true, nil
	default:
		if ctx.Loader == nil {
			return nil, false, nil
		}
		v, ok, err := ctx.Loader(ctx, typ, key)
		if err != nil {
			return "", false, fmt.Errorf("expand $%s: %w", keyname(typ, key), err)
		}
		return v, ok, nil
	}
	s, err := f()
	if err != nil {
		return "", false, fmt.Errorf("expand $%s: %w", keyname(typ, key), err)
	}
	return s, true, nil
}

// keyname returns the key name.
func keyname(typ, key string) string {
	if key != "" {
		return typ + "{" + key + "}"
	}
	return typ
}

// contextKey is the context key type.
type contextKey int

// context keys.
const (
	ctxKey contextKey = iota
)

// WithContext adds a [Context] to the parent [context.Context].
func WithContext(parent context.Context, ctx *Context) context.Context {
	return context.WithValue(parent, ctxKey, ctx)
}

// Ctx returns the context from the [context.Context].
func Ctx(ctx context.Context) (*Context, bool) {
	c, ok := ctx.Value(ctxKey).(*Context)
	return c, ok
}

// OnErr is the on error handling type.
type OnErr uint8

// On error handling types.
const (
	OnErrExit OnErr = iota
	OnErrContinue
	OnErrPanic
)

// Option returns an [opt] for the error handling type.
func (e OnErr) Option() option {
	return option{
		name: "OnErr(" + e.String() + ")",
		cmd: func(cmd *Command) error {
			cmd.OnErr = e
			return nil
		},
	}
}

// Error is a package error.
type Error string

// Error satisfies the [error] interface.
func (err Error) Error() string {
	return string(err)
}

// Errors.
const (
	// ErrUsageNotSet is the usage not set error.
	ErrUsageNotSet Error = "usage not set"
	// ErrCanOnlyBeUsedWithRootCommand is the can only be used with root command error.
	ErrCanOnlyBeUsedWithRootCommand Error = "can only be used with root command"
	// ErrAppliedToInvalidType is the applied to invalid type error.
	ErrAppliedToInvalidType Error = "applied to invalid type"
	// ErrInvalidArgCount is the invalid arg count error.
	ErrInvalidArgCount Error = "invalid arg count"
	// ErrInvalidFlagName is the invalid flag name error.
	ErrInvalidFlagName Error = "invalid flag name"
	// ErrInvalidShortName is the invalid short name error.
	ErrInvalidShortName Error = "invalid short name"
	// ErrCouldNotCreateValue is the could not create value error.
	ErrCouldNotCreateValue Error = "could not create value"
	// ErrUnknownCommand is the unknown command error.
	ErrUnknownCommand Error = "unknown command"
	// ErrUnknownFlag is the unknown flag error.
	ErrUnknownFlag Error = "unknown flag"
	// ErrMissingArgument is the missing argument error.
	ErrMissingArgument Error = "missing argument"
	// ErrInvalidType is the invalid type error.
	ErrInvalidType Error = "invalid type"
	// ErrInvalidSize is the invalid size error.
	ErrInvalidSize Error = "invalid size"
	// ErrInvalidRate is the invalid rate error.
	ErrInvalidRate Error = "invalid rate"
	// ErrUnknownSize is the unknown size error.
	ErrUnknownSize Error = "unknown size"
	// ErrUnknownTagOption is the unknown tag option error.
	ErrUnknownTagOption Error = "unknown tag option"
	// ErrInvalidValue is the invalid value error.
	ErrInvalidValue Error = "invalid value"
	// ErrInvalidArg is the invalid arg error.
	ErrInvalidArg Error = "invalid arg"
	// ErrInvalidConversion is the invalid conversion error.
	ErrInvalidConversion Error = "invalid conversion"
	// ErrTypeMismatch is the type mismatch error.
	ErrTypeMismatch Error = "type mismatch"
	// ErrExit is the exit error.
	ErrExit Error = "exit"
)

// SuggestionError is a suggestion error.
type SuggestionError struct {
	Parent  *Command
	Arg     string
	Command *Command
}

// NewSuggestionError creates a suggestion error.
func NewSuggestionError(parent *Command, arg string, command *Command) error {
	return &SuggestionError{
		Parent:  parent,
		Arg:     arg,
		Command: command,
	}
}

// Error satisfies the [error] interface.
func (err *SuggestionError) Error() string {
	return fmt.Sprintf(text.SuggestionErrorMessage, ErrUnknownCommand, err.Arg, err.Parent.Name)
}

// Unwrap satisfies the [errors.Unwrap] interface.
func (err *SuggestionError) Unwrap() error {
	return ErrUnknownCommand
}

// ErrorDetails returns error details. Used by [DefaultErrorHandler].
func (err *SuggestionError) ErrorDetails() string {
	return fmt.Sprintf(text.SuggestionErrorDetails, err.Command.Name)
}

// Wrap wraps a line of text to the specified width, adding a prefix of empty
// prefixWidth to each wrapped line.
func Wrap(s string, width, prefixWidth int) string {
	words := strings.Fields(strings.TrimSpace(s))
	if len(words) == 0 {
		return ""
	}
	prefix, wrapped := strings.Repeat(" ", prefixWidth), words[0]
	left := width - prefixWidth - len(wrapped)
	for _, word := range words[1:] {
		if left < DefaultWidth(word)+1 {
			wrapped += "\n" + prefix + word
			left = width - DefaultWidth(word)
		} else {
			wrapped += " " + word
			left -= 1 + DefaultWidth(word)
		}
	}
	return wrapped
}

// Ldist is a Levenshtein distance implementation.
func Ldist[T []E, E cmp.Ordered](a, b T) int {
	m, n := len(a), len(b)
	v := make([]int, m+1)
	var i, j, d, l int
	for i = range m + 1 {
		v[i] = i
	}
	for i = range n {
		for v[0], j, d = i+1, 1, i; j <= m; j, d = j+1, l {
			if l = v[j]; a[j-1] != b[i] {
				d++
			}
			v[j] = min(v[j-1]+1, v[j]+1, d)
		}
	}
	return v[m]
}

// BuildVersion returns the Go build version, or [DefaultVersionString].
func BuildVersion() string {
	ver := "(devel)"
	if info, ok := debug.ReadBuildInfo(); ok {
		mod := &info.Main
		if mod.Replace != nil {
			mod = mod.Replace
		}
		if mod.Version != "" && mod.Version != "(devel)" {
			ver = mod.Version
		}
	}
	return ver
}

// keyRE matches keys used by [Context.Expand].
var keyRE = regexp.MustCompile(`\$([A-Z][A-Z0-9_]{1,32})(?:\{([^}]{1,128})\})?`)

// prepend is a generic prepend.
func prepend[S ~[]E, E any](v S, s ...E) S {
	return append(s, v...)
}
