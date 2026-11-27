// Package ox is a minimal dependency Go (and TinyGo compatible) command-line
// flag and argument parsing library.
package ox

//go:generate stringer -type OnErr

import (
	"bytes"
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"
	"time"
	"unicode"

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
	// DefaultFlagNameMapper is the default flag name mapper.
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
	// DefaultEmitExecExitError toggles emitting captured stderr output from
	// [exec.ExitError] to the [Context.Stderr].
	DefaultEmitExecExitError = false
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
				// error code
				code := 1
				if v, ok := err.(interface{ ErrorCode() int }); ok {
					code = v.ErrorCode()
				}
				// build error message
				var sb strings.Builder
				_, _ = sb.WriteString(strings.TrimRightFunc(err.Error(), unicode.IsSpace))
				if e, ok := asExitError(err); ok && ctx.EmitExecExitError && len(e.Stderr) != 0 {
					code = e.ExitCode()
					if b := bytes.TrimRightFunc(e.Stderr, unicode.IsSpace); len(b) != 0 {
						sb.WriteString(": ")
						sb.Write(b)
					}
				}
				// emit error message
				_, _ = fmt.Fprintf(w, text.ErrorMessage, sb.String())
				if v, ok := err.(interface{ ErrorDetails() string }); ok {
					_, _ = fmt.Fprint(w, v.ErrorDetails())
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
	// DefaultMaxDist is the default maximum Levenshtein distance for
	// suggestions and for the corrections of completion. 0 turns off the
	// distance match. [MaxDist] changes it for one command.
	DefaultMaxDist = 2
	// DefaultFoldCase when true makes suggestions, and the completion of
	// command and flag names, compare names without case. The parse of a
	// command name always compares with case. Set it to false to compare
	// with case.
	DefaultFoldCase = true
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
	// DefaultCompActiveHelpSuffix is the suffix of the environment variable
	// that turns off ActiveHelp for a program, such as APP_ACTIVE_HELP=0.
	DefaultCompActiveHelpSuffix = "_ACTIVE_HELP"
	// DefaultCompActiveHelpMarker starts a completion that is a line of
	// ActiveHelp, which the completion scripts show as help and not as a
	// completion. See [NewActiveHelp].
	DefaultCompActiveHelpMarker = "_activeHelp_ "
	// DefaultCompWrite is the default completion write func, that writes the
	// completion script from the passed template.
	DefaultCompWrite = func(ctx *Context, cmd *Command, noDescriptions bool, shell, templ string) error {
		rootName := cmd.RootName()
		varName := regexp.MustCompile(`[^A-Za-z0-9_]`).ReplaceAllString(rootName, "_")
		activeHelp := compActiveHelpVar(rootName)
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
		// the scripts discard standard error, and read the directive
		comps, dir, err := ctx.Comps()
		if err != nil {
			_, _ = fmt.Fprintf(ctx.Stderr, text.ErrorMessage, err)
		}
		if os.Getenv(compActiveHelpVar(ctx.Root.Name)) == "0" {
			comps = slices.DeleteFunc(comps, isActiveHelp)
		}
		for _, comp := range comps {
			if !noDescriptions && comp.Usage != "" {
				_, _ = fmt.Fprintln(ctx.Stdout, comp.Name+"\t"+comp.Usage)
			} else {
				_, _ = fmt.Fprintln(ctx.Stdout, comp.Name)
			}
		}
		_, _ = fmt.Fprintf(ctx.Stdout, ":%d\n", dir)
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
	// DefaultSizePrec is the default [Size] display precision.
	DefaultSizePrec = -2
	// DefaultRatePrec is the default [Rate] display precision.
	DefaultRatePrec = -2
	// DefaultRateUnit is the default [Rate] unit used to display the rate, if
	// none is specified or when precision is negative.
	DefaultRateUnit = time.Second
)

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

// RunError creates a [Context] from the options, as [RunContext] does, then
// parses, validates and runs the command. It returns the first error, and
// does not write it, exit or panic, whatever [Context.OnErr] is. So the
// program chooses the message and the exit status. When help, the version or
// a completion stops the run with [ErrExit], RunError returns nil.
func RunError(ctx context.Context, opts ...Option) error {
	c, err := NewContext(opts...)
	if err == nil {
		err = c.Parse()
	}
	if err == nil {
		err = c.Validate()
	}
	if err == nil {
		err = c.Run(ctx)
	}
	if errors.Is(err, ErrExit) {
		return nil
	}
	return err
}

// AppFrom runs the program that v describes, as [Run] does. v is a pointer
// to a struct, and [App] builds the command tree from it. AppFrom adds
// [Defaults], so the program gets help, a version and completion. The options
// in opts come after.
//
// Pass [App] to [RunError] to choose the [context.Context], or to get the
// error.
func AppFrom[T *E, E any](v T, opts ...Option) {
	RunContext(DefaultContext, prepend(opts, App(v), Defaults())...)
}

// Context is the execution (run) context used by [Run]/[RunContext].
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
	// EmitExecExitError toggles emitting captured stderr output from
	// [exec.ExitError] to [Context.Stderr].
	EmitExecExitError bool
	// OnErr is the on error handling. Used by the default Handle func.
	OnErr OnErr
	// Continue is the func used to decide if it should continue on an error.
	Continue func(*Command, error) bool
	// Handler is the func used to handle errors within [Run]/[RunContext].
	Handler func(error) bool
	// Override override expansion keys.
	Override func(ConfigType, string) (any, bool, error)
	// Root is the root command created within [Run]/[RunContext].
	Root *Command
	// Exec is the exec target, determined by the Root's definition and after
	// Args.
	Exec *Command
	// Vars are the variables parsed from the flag definitions of the Root
	// command and its sub-commands.
	Vars Vars
	// Interpolate interpolates variables based on the context.
	Interpolate func(*Context, any) (any, error)
	// Lookup retrieves config variables based on the context.
	Lookup func(*Context, ConfigType, string) (any, error)
	// Loader is the config loader func.
	Loader func(*Context, ConfigType, string) (any, bool, error)

	// files are the config files that were read.
	files []configFile
	// loaded is true when the config files were read.
	loaded bool
}

// NewContext creates a new run context.
func NewContext(opts ...Option) (*Context, error) {
	ctx := &Context{
		Exit: os.Exit,
		Panic: func(v any) {
			panic(v)
		},
		Comp:              DefaultComp,
		Stdin:             os.Stdin,
		Stdout:            os.Stdout,
		Stderr:            os.Stderr,
		Args:              DefaultStripTestFlags(os.Args[1:]),
		EmitExecExitError: DefaultEmitExecExitError,
		Interpolate:       DefaultInterpolate,
		Lookup:            DefaultLookup,
		Loader:            DefaultLoader,
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

// Parse parses the context args, then fills each flag that the args did not
// set, with [Context.Fill].
func (ctx *Context) Parse() error {
	if ctx.Root.Comp && len(ctx.Args) != 0 && strings.HasPrefix(ctx.Args[0], DefaultCompName) {
		return ctx.Comp(ctx)
	}
	var err error
	if ctx.Exec, ctx.Args, err = Parse(ctx, ctx.Root, ctx.Args); err != nil {
		return err
	}
	if err := ctx.Fill(ctx.Exec); err != nil {
		return newCommandError(ctx.Exec.Name, err)
	}
	return nil
}

// Comps returns the completions for the context. The first arg is the name
// of the completion command, and the last arg is the word to complete.
//
// Comps ignores each error in the args before the last arg, as cobra does.
// After `--`, it completes only positional arguments. It never offers a
// hidden or a deprecated flag, nor a hidden or a deprecated command.
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
		Continue: func(*Command, error) bool {
			return true
		},
		Args: ctx.Args[1:n],
		Vars: make(Vars),
	}
	var dashdash bool
	c.Exec, c.Args, dashdash, _ = parse(c, c.Root, c.Args)
	// a value from the environment or a config file is a help, and its
	// error is not
	_ = c.Fill(c.Exec)
	var comps []Completion
	var dir CompDirective
	var err error
	switch toComplete := ctx.Args[n]; {
	case dashdash:
		comps, dir, err = c.compArgs(toComplete, false)
	case strings.HasPrefix(toComplete, "-") && strings.Contains(toComplete, "="):
		// --flag=value and -f=value
		name, value, _ := strings.Cut(toComplete, "=")
		short := !strings.HasPrefix(name, "--")
		if short {
			r := []rune(name)
			name = string(r[len(r)-1:])
		}
		g := c.Exec.Flag(strings.TrimPrefix(name, "--"), true, short)
		if g == nil {
			return nil, CompDefault, nil
		}
		comps, dir, err = c.compFlagValue(g, value)
	case strings.HasPrefix(toComplete, "-"):
		comps, dir = c.Exec.CompFlags(strings.TrimLeft(toComplete, "-"), false, false, !strings.HasPrefix(toComplete, "--"))
	default:
		if g := compPrevFlag(c.Exec, prev(ctx.Args[1:], n-1)); g != nil {
			comps, dir, err = c.compFlagValue(g, toComplete)
		} else {
			comps, dir, err = c.compArgs(toComplete, true)
		}
	}
	if err != nil {
		return nil, CompError, err
	}
	return comps, dir &^ CompNoFilter, nil
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
			if value, err = asString[string](v); err != nil {
				return fmt.Errorf("populate %s: cannot expand %v (%T): %w", g.Name, g.Def, g.Def, err)
			}
		}
		if err := ctx.Vars.Set(ctx, g, value, false); err != nil {
			return fmt.Errorf("populate %s: %w", g.Name, err)
		}
	}
	return nil
}

// Expand expands a variable using [Context.Interpolate]. A nil context
// returns the variable as it is.
func (ctx *Context) Expand(v any) (any, error) {
	if ctx != nil && ctx.Interpolate != nil {
		return ctx.Interpolate(ctx, v)
	}
	return v, nil
}

// compArgs returns the completions of a positional argument. When commands is
// true, it adds the sub commands of the command.
func (ctx *Context) compArgs(toComplete string, commands bool) ([]Completion, CompDirective, error) {
	cmd := ctx.Exec
	var comps []Completion
	if commands && len(ctx.Args) == 0 {
		comps, _ = cmd.CompCommands(toComplete, false, false)
	}
	switch {
	case cmd.CompFunc != nil:
		v, dir, err := compFunc(ctx, cmd.CompFunc, toComplete)
		return compNearest(append(comps, v...)), dir, err
	case len(cmd.ValidArgs) != 0:
		if cmd.maxArgs < 0 || len(ctx.Args) < cmd.maxArgs {
			comps = compNearest(append(comps, compValues(cmd, toComplete, cmd.ValidArgs)...))
		}
		if len(comps) != 0 {
			return comps, CompNoFileComp | CompKeepOrder, nil
		}
		return nil, CompNoFileComp, nil
	case len(comps) != 0:
		return comps, CompNoFileComp | CompKeepOrder, nil
	}
	return nil, CompDefault, nil
}

// compFlagValue returns the completions of the value of the flag.
func (ctx *Context) compFlagValue(g *Flag, toComplete string) ([]Completion, CompDirective, error) {
	switch {
	case g.CompFunc != nil:
		return compFunc(ctx, g.CompFunc, toComplete)
	case len(g.Allowed) == 0 && isPath(g.Type):
		comps, dir := g.compPath()
		return comps, dir, nil
	case len(g.Allowed) == 0:
		return nil, CompDefault, nil
	}
	// a slice completes the value after the last separator
	var head string
	if (g.Type == SliceT || g.Type == ArrayT) && g.Split != "" {
		if i := strings.LastIndex(toComplete, g.Split); i != -1 {
			head, toComplete = toComplete[:i+len(g.Split)], toComplete[i+len(g.Split):]
		}
	}
	values := make([]string, len(g.Allowed))
	for i, v := range g.Allowed {
		values[i] = fmt.Sprint(v)
	}
	comps := compValues(ctx.Exec, toComplete, values)
	if len(comps) == 0 {
		return nil, CompNoFileComp, nil
	}
	for i := range comps {
		comps[i].Name = head + comps[i].Name
	}
	return comps, CompNoFileComp | CompKeepOrder, nil
}

// compFunc calls f, and keeps only the completions that start with
// toComplete, unless f returns [CompNoFilter].
func compFunc(ctx *Context, f func(*Context, string) ([]Completion, CompDirective, error), toComplete string) ([]Completion, CompDirective, error) {
	comps, dir, err := f(ctx, toComplete)
	if err != nil {
		return nil, CompError, err
	}
	if dir&CompNoFilter == 0 {
		comps = slices.DeleteFunc(comps, func(c Completion) bool {
			return !strings.HasPrefix(c.Name, toComplete) && !isActiveHelp(c)
		})
	}
	for i := range comps {
		comps[i].Dist = 0
	}
	return comps, dir, nil
}

// compValues returns each value that starts with toComplete. When none
// does, it returns each value within the maximum distance of the command, as
// a correction.
func compValues(cmd *Command, toComplete string, values []string) []Completion {
	var comps []Completion
	for _, s := range values {
		if strings.HasPrefix(s, toComplete) {
			comps = append(comps, NewCompletion(s, "", 0))
		}
	}
	if maxDist := maxDist(cmd); len(comps) == 0 && 0 < maxDist && toComplete != "" {
		r := []rune(foldCase(toComplete))
		for _, s := range values {
			if dist := Ldist([]rune(foldCase(s)), r); dist <= maxDist {
				comps = append(comps, NewCompletion(s, "", dist))
			}
		}
	}
	return compNearest(comps)
}

// compNearest keeps only the completions with the smallest distance, so
// that a correction shows only when nothing matches. It keeps the order.
func compNearest(comps []Completion) []Completion {
	if len(comps) == 0 {
		return nil
	}
	d := slices.MinFunc(comps, func(a, b Completion) int {
		return a.Dist - b.Dist
	}).Dist
	return slices.DeleteFunc(comps, func(c Completion) bool {
		return d < c.Dist
	})
}

// compPrevFlag returns the flag whose value follows s, such as `--config` or
// `-vc`. It returns nil when s is not a flag that takes a value, or when s
// holds the value, as in `--config=file` or `-cfile`.
func compPrevFlag(cmd *Command, s string) *Flag {
	switch {
	case len(s) < 2, s[0] != '-', s == "--", strings.Contains(s, "="):
		return nil
	case strings.HasPrefix(s, "--"):
		if g := cmd.Flag(s[2:], true, false); g != nil && !g.NoArg {
			return g
		}
		return nil
	}
	r := []rune(s[1:])
	for i, c := range r {
		switch g := cmd.Flag(string(c), true, true); {
		case g == nil:
			return nil
		case g.NoArg:
		case i == len(r)-1:
			return g
		default:
			return nil
		}
	}
	return nil
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

// Option returns the option that sets the error handling of the [Context]
// and of the [Command], such as `ox.Run(ox.OnErrPanic, ...)`.
func (e OnErr) Option() option {
	return option{
		name: "OnErr(" + e.String() + ")",
		ctx: func(ctx *Context) error {
			ctx.OnErr = e
			return nil
		},
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
	// ErrUnknownLoader is the unknown loader error.
	ErrUnknownLoader Error = "unknown loader"
	// ErrUnknownKey is the unknown key error.
	ErrUnknownKey Error = "unknown key"
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
	// ErrInvalidOp is the invalid op error.
	ErrInvalidOp Error = "invalid op"
	// ErrInvalidKey is the invalid key error.
	ErrInvalidKey Error = "invalid key"
	// ErrInvalidFilter is the invalid filter error.
	ErrInvalidFilter Error = "invalid filter"
	// ErrTypeMismatch is the type mismatch error.
	ErrTypeMismatch Error = "type mismatch"
	// ErrNotImplemented is the not implemented error.
	ErrNotImplemented Error = "not implemented"
	// ErrInvalidConfig is the invalid config error, for a config file that
	// its format cannot decode.
	ErrInvalidConfig Error = "invalid config"
	// ErrConfigNotFound is the config not found error.
	ErrConfigNotFound Error = "config not found"
	// ErrExit is the exit error.
	ErrExit Error = "exit"
	// ErrDuplicateFlag is the duplicate flag error, for two flags of a
	// command with the same name.
	ErrDuplicateFlag Error = "duplicate flag"
	// ErrNotExist is the does not exist error of a [FileT] or a [DirT].
	ErrNotExist Error = "does not exist"
	// ErrExist is the exists error of a path with [NotExist].
	ErrExist Error = "exists"
	// ErrIsDir is the is a directory error of a [FileT].
	ErrIsDir Error = "is a directory"
	// ErrNotDir is the not a directory error of a [DirT].
	ErrNotDir Error = "not a directory"
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

// BuildVersion returns the version of the main module from the build info,
// or "(devel)".
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

// prepend is a generic prepend.
func prepend[S ~[]E, E any](v S, s ...E) S {
	return append(s, v...)
}

// compActiveHelpVar returns the name of the environment variable that turns
// off ActiveHelp for the program.
func compActiveHelpVar(name string) string {
	return strings.ToUpper(regexp.MustCompile(`[^A-Za-z0-9_]`).ReplaceAllString(name, "_")) + DefaultCompActiveHelpSuffix
}

// isActiveHelp returns true when c is a line of ActiveHelp.
func isActiveHelp(c Completion) bool {
	return strings.HasPrefix(c.Name, DefaultCompActiveHelpMarker)
}

// asExitError is a standin for Go 1.26's errors.AsType.
//
// TODO: remove when tinygo supports go1.26.
func asExitError(err error) (*exec.ExitError, bool) {
	if e := (*exec.ExitError)(nil); errors.As(err, &e) {
		return e, true
	}
	return nil, false
}
