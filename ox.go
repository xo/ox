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
	"strconv"
	"strings"
	"time"

	"github.com/xo/ox/strcase"
	"github.com/xo/ox/text"
)

var (
	// DefaultContext is the default [context.Context].
	DefaultContext = context.Background()
	// DefaultTagName is the default struct tag name used in [FromFlags] and
	// related func's.
	DefaultTagName = "ox"
	// DefaultLayout is the default timestamp layout used for formatting and
	// parsing [Time] values.
	DefaultLayout = time.RFC3339
	// DefaultStripGoTestFlags when enabled strips Go's `-test.` prefix'd
	// flags.
	DefaultStripGoTestFlags = true
	// DefaultWrapWidth is the default wrap width.
	DefaultWrapWidth = 95
	// DefaultWrap wraps sections of text (separated by "\n\n") using [Wrap].
	DefaultWrap = func(s string, wrapWidth, prefixWidth int) string {
		sb := new(strings.Builder)
		for i, line := range strings.Split(s, "\n\n") {
			if i != 0 {
				_, _ = sb.WriteString("\n\n")
				_, _ = sb.WriteString(strings.Repeat(" ", prefixWidth))
			}
			_, _ = sb.WriteString(Wrap(line, wrapWidth, prefixWidth))
		}
		return sb.String()
	}
	// DefaultWidth returns the width of a string used by [Wrap].
	DefaultWidth = func(s string) int {
		return len(s)
	}
	// DefaultFlagNameMaper is the default flag name mapper.
	DefaultFlagNameMapper = func(s string) string {
		return strings.ReplaceAll(strcase.CamelToSnake(s), "_", "-")
	}
	// DefaultVersionString is the default version string.
	DefaultVersionString = "0.0.0-dev"
	// DefaultVersionTrimPrefix is used by [DefaultVersionMapper] to trim the
	// `v` prefix on build version.
	DefaultVersionTrimPrefix = true
	// DefaultVersion is the default version func.
	DefaultVersion = func(ctx *Context) error {
		ver := BuildVersion()
		var name string
		if ctx != nil && ctx.Root != nil {
			name = ctx.Root.Name
		}
		name, ver = DefaultVersionMapper(name, ver)
		var w io.Writer
		if ctx != nil {
			w = ctx.Stdout
		}
		if w == nil {
			w = os.Stdout
		}
		fmt.Fprintln(w, name, ver)
		return ErrExit
	}
	// DefaultVersionMapper maps the passed name, version.
	DefaultVersionMapper = func(name, ver string) (string, string) {
		if name == "" {
			name = filepath.Base(os.Args[0])
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
	// DefaultPrec is the default [Size]/[Rate] display precision.
	DefaultPrec = -2
	// DefaultIEC is the default [Size]/[Rate] IEC display setting.
	DefaultIEC = true
	// DefaultUnit is the default [Rate] unit.
	DefaultUnit = time.Second
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

// Context is a [Run]/[RunContext] execution context.
type Context struct {
	// Exit is the exit func.
	Exit func(int)
	// Panic is the panic func.
	Panic func(any)
	// Comp is the completion func.
	Comp func(*Context) error
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
	// Override are overriding expansions.
	Override map[string]string
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
			fmt.Fprintf(os.Stderr, "COMP CONTINUE ERR: %v\n", err)
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
	fmt.Fprintf(os.Stderr, "COMP COMMAND: %s\n", c.Exec.Name)
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

// Populate populates the context's vars with all the command's flags values,
// overwriting any set variables if applicable. When all is true, all flag
// values will be populated, otherwise only flags with default values will be.
//
// When overwrite is true, existing vars will be set either to flag's empty or
// default value.
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
			var err error
			if value, err = ctx.Expand(g.Def); err != nil {
				return err
			}
		}
		if err := ctx.Vars.Set(ctx, g, value, false); err != nil {
			return fmt.Errorf("cannot populate %s with %q: %w", g.Name, value, err)
		}
	}
	return nil
}

// Expand expands variables in v, where any variable can be the following:
//
//	$APPNAME - the root command's name
//	$HOME - the current user's home directory
//	$USER - the current user's user name
//	$CONFIG - the current user's config directory
//	$APPCONFIG - the current user's config directory, with the root command's name added as a subdir
//	$CACHE - the current user's cache directory
//	$APPCACHE - the current user's cache directory, with the root command's name added as a subdir
//	$NUMCPU - the value of [runtime.NumCPU]
//	$ARCH - the value of [runtime.GOARCH]
//	$OS - the value of [runtime.GOOS]
//	$ENV{KEY} - the environment value for $KEY
//	$CONFIG_TYPE{KEY} - the registered config file loader type and key value, for example: `$YAML{my_key}`, `$TOML{my_key}`
//
// TODO: finish implementation for $CONFIG_TYPE, expand anywhere in string
func (ctx *Context) Expand(v any) (string, error) {
	s, ok := v.(string)
	if !ok {
		return asString[string](v)
	}
	if ctx.Override != nil {
		if s, ok := ctx.Override[s]; ok {
			return s, nil
		}
	}
	var f func() (string, error)
	switch s {
	case "$HOME":
		f = userHomeDir
	case "$USER":
		f = func() (string, error) {
			u, err := user.Current()
			if err != nil {
				return "", err
			}
			return u.Username, nil
		}
	case "$CONFIG":
		f = userConfigDir
	case "$APPCONFIG":
		f = func() (string, error) {
			dir, err := userConfigDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ctx.Root.Name), nil
		}
	case "$CACHE":
		f = userCacheDir
	case "$APPCACHE":
		f = func() (string, error) {
			dir, err := userCacheDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ctx.Root.Name), nil
		}
	case "$APPNAME":
		return ctx.Root.Name, nil
	case "$NUMCPU":
		return strconv.Itoa(runtime.NumCPU()), nil
	case "$ARCH":
		return runtime.GOARCH, nil
	case "$OS":
		return runtime.GOOS, nil
	default:
		return s, nil
	}
	s, err := f()
	if err != nil {
		return "", fmt.Errorf("unable to expand %q: %w", s, err)
	}
	return s, nil
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

// Wrap wraps a line of text to the specified width, and adding a prefix of
// empty prefixWidth to each wrapped line.
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
	ver := DefaultVersionString
	if info, ok := debug.ReadBuildInfo(); ok && ver == "0.0.0-dev" {
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

// ParseRate parses a byte rate string.
func ParseRate(s string) (int64, int, bool, time.Duration, error) {
	unit := time.Second
	if i := strings.LastIndexByte(s, '/'); i != -1 {
		switch s[i+1:] {
		// unitMap is the duration unit map.
		case "ns":
			unit = time.Nanosecond
		case "us", "µs", "μs": // U+00B5 = micro symbol, U+03BC = Greek letter mu
			unit = time.Microsecond
		case "ms":
			unit = time.Millisecond
		case "s":
			unit = time.Second
		case "m":
			unit = time.Minute
		case "h":
			unit = time.Hour
		default:
			return 0, 0, false, 0, fmt.Errorf("%w %q", ErrInvalidRate, s)
		}
		s = s[:i]
	}
	sz, prec, iec, err := ParseSize(s)
	if err != nil {
		return 0, 0, false, 0, err
	}
	return sz, prec, iec, unit, nil
}

// FormatSize formats a byte rate.
func FormatRate(size int64, prec int, iec bool, unit time.Duration) string {
	return FormatSize(size, prec, iec) + "/" + unitString(unit)
}

// ParseSize parses a byte size string.
func ParseSize(s string) (int64, int, bool, error) {
	m := sizeRE.FindStringSubmatch(s)
	if m == nil {
		return 0, 0, false, fmt.Errorf("%w %q", ErrInvalidSize, s)
	}
	f, err := strconv.ParseFloat(m[2], 64)
	switch {
	case err != nil:
		return 0, 0, false, fmt.Errorf("%w: %w", ErrInvalidSize, err)
	case m[1] == "-":
		f = -f
	}
	sz, iec, err := sizeType(m[3])
	if err != nil {
		return 0, 0, false, fmt.Errorf("%w: %w", ErrInvalidSize, err)
	}
	prec := DefaultPrec
	if i := strings.LastIndexByte(m[2], '.'); i != -1 {
		prec = len(m[2]) - i - 1
	}
	return int64(f * float64(sz)), prec, iec, nil
}

// FormatSize formats a byte size.
//
// Formatting rules follow [strconv.FormatFloat] rules, with additional option
// for a precision of -2, which will remove trailing zeroes and the decimal if
// the decimal places are all zeroes.
func FormatSize(size int64, prec int, iec bool) string {
	n, t, suffix := KB, "kMGTPE", "B"
	if iec {
		n, t, suffix = KiB, "KMGTPE", "iB"
	}
	var neg string
	if size < 0 {
		neg, size = "-", -size
	}
	if size < n {
		return neg + strconv.FormatInt(size, 10) + " B"
	}
	e, d := 0, n
	for i := size / n; n <= i; i /= n {
		d *= n
		e++
	}
	precabs := prec
	if precabs < -1 {
		precabs = -precabs
	}
	s := strconv.FormatFloat(float64(size)/float64(d), 'f', precabs, 64)
	if prec < -1 {
		s = strings.TrimRight(s, ".0")
	}
	return neg + s + " " + string(t[e]) + suffix
}

// sizeType returns the size of s.
func sizeType(s string) (int64, bool, error) {
	switch strings.ToLower(s) {
	case "", "b":
		return B, false, nil
	case "kb":
		return KB, false, nil
	case "mb":
		return MB, false, nil
	case "gb":
		return GB, false, nil
	case "tb":
		return TB, false, nil
	case "pb":
		return PB, false, nil
	case "eb":
		return EB, false, nil
	case "kib":
		return KiB, true, nil
	case "mib":
		return MiB, true, nil
	case "gib":
		return GiB, true, nil
	case "tib":
		return TiB, true, nil
	case "pib":
		return PiB, true, nil
	case "eib":
		return EiB, true, nil
	}
	return 0, false, fmt.Errorf("%w %q", ErrUnknownSize, s)
}

// Byte sizes.
const (
	B  int64 = 1
	KB int64 = 1_000
	MB int64 = 1_000_000
	GB int64 = 1_000_000_000
	TB int64 = 1_000_000_000_000
	PB int64 = 1_000_000_000_000_000
	EB int64 = 1_000_000_000_000_000_000
)

// IEC byte sizes.
const (
	KiB int64 = 1_024
	MiB int64 = 1_048_576
	GiB int64 = 1_073_741_824
	TiB int64 = 1_099_511_627_776
	PiB int64 = 1_125_899_906_842_624
	EiB int64 = 1_152_921_504_606_846_976
)

// unitString returns the string for a time unit (duration).
func unitString(unit time.Duration) string {
	switch {
	case unit == 0, unit > time.Hour:
		return unitString(DefaultUnit)
	case unit > time.Minute:
		return "h"
	case unit > time.Second:
		return "m"
	case unit > time.Millisecond:
		return "s"
	case unit > time.Microsecond:
		return "ms"
	case unit > time.Nanosecond:
		return "µs"
	}
	return "ns"
}

// sizeRE matches sizes.
var sizeRE = regexp.MustCompile(`(?i)^([-+])?([0-9]+(?:\.[0-9]*)?)(?: ?([a-z]+))?$`)

// prepend is a generic prepend.
func prepend[S ~[]E, E any](v S, s ...E) S {
	return append(s, v...)
}
