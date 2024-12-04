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
	// DefaultWrap wraps a line of text with [Wrap] using [DefaultWrapWidth]
	// for the width.
	DefaultWrap = func(s string, prefixWidth int) string {
		return Wrap(s, DefaultWrapWidth, prefixWidth)
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
	// DefaultVersionMapper is the default version mapper.
	DefaultVersionMapper = func(name, ver string) (string, string) {
		if name == "" {
			name = filepath.Base(os.Args[0])
		}
		if regexp.MustCompile(`^v[0-9]+\.`).MatchString(ver) {
			ver = strings.TrimPrefix(ver, "v")
		}
		return name, ver
	}
	// DefaultVersion is the default version func.
	DefaultVersion = func(ctx *Context) error {
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
		var name string
		if ctx.Root != nil {
			name = ctx.Root.Name
		}
		name, ver = DefaultVersionMapper(name, ver)
		fmt.Fprintln(ctx.Stdout, name, ver)
		return ErrExit
	}
)

// Run creates a [Context] and builds a [Command] and its [FlagSet] based on
// the options and then executing the command. See [RunContext] to pass a
// [context.Context] to the executed command, or set the [DefaultContext].
func Run(opts ...Option) {
	RunContext(DefaultContext, opts...)
}

// RunContext creates a [Command] for f using [os.Args] by default, unless
// arguments were specified using a [ContextOption].
func RunContext(ctx context.Context, opts ...Option) {
	c, err := NewContext(opts...)
	if err != nil && c.Handle(err) {
		return
	}
	// parse flags into context
	if err = c.Parse(); err != nil && c.Handle(err) {
		return
	}
	// validate args
	if err = c.Validate(); err != nil && c.Handle(err) {
		return
	}
	// execute
	if err = c.Run(ctx); err != nil && c.Handle(err) {
		return
	}
}

// Context is a [Run]/[RunContext] context.
type Context struct {
	// Args are the arguments to parse, normally [os.Args][1:].
	Args []string
	// Stdin is the standard in to use, normally [os.Stdin].
	Stdin io.Reader
	// Stdout is the standard out to use, normally [os.Stdout].
	Stdout io.Writer
	// Stderr is the standard error to use, normally [os.Stderr].
	Stderr io.Writer
	// OnErr is the on error handling. Used by the default Handle func.
	OnErr OnErr
	// Handle is the func used to handle errors within [Run]/[RunContext].
	Handle func(error) bool
	// Root is the root command created within [Run]/[RunContext].
	Root *Command
	// Exec is the exec target, determined by the Root's definition and after Args.
	Exec *Command
	// Vars are the variables parsed from the flag definitions of the Root
	// command and its sub-commands.
	Vars Vars
	// Override are overriding expansions.
	Override map[string]string
}

// NewContext creates a new run context.
func NewContext(opts ...Option) (*Context, error) {
	ctx := &Context{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Args:   stripTestFlags(os.Args[1:]),
	}
	ctx.Handle = func(err error) bool {
		var w io.Writer = os.Stderr
		if ctx.Stderr != nil {
			w = ctx.Stderr
		}
		switch {
		case errors.Is(err, ErrExit):
		case ctx.OnErr == OnErrContinue:
			return false
		case ctx.OnErr == OnErrExit:
			fmt.Fprintf(w, "%s%v\n", text.ErrorMessagePrefix, err)
			os.Exit(1)
		case ctx.OnErr == OnErrPanic:
			panic(err)
		}
		return true
	}
	if err := applyOpts(ctx, opts...); err != nil {
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
	var err error
	ctx.Exec, ctx.Args, err = Parse(ctx, ctx.Root, ctx.Args, ctx.Vars)
	return err
}

// Validate validates the args.
func (ctx *Context) Validate() error {
	return ctx.Exec.Validate(ctx.Args)
}

// Run runs the command.
func (ctx *Context) Run(parent context.Context) error {
	if ctx.Exec.Exec != nil {
		return ctx.Exec.Exec(WithContext(parent, ctx), ctx.Args)
	}
	_, _ = ctx.Exec.HelpContext(ctx).WriteTo(ctx.Stdout)
	return nil
}

// Expand expands s.
//
//	$HOME - the current user's home directory
//	$USER - the current user's user name
//	$CACHE - the current user's cache directory
//	$APPCACHE - the current user's cache directory, with the root command's name added as a subdir
//	$NUMCPU - the value of [runtime.NumCPU]
//	$ARCH - the value of [runtime.GOARCH]
//	$OS - $ the value of [runtime.GOOS]
//	$ENV{KEY} - the environment value for $KEY
//	$CFG{[TYPE::]KEY} - the registered config file loader type and key value
//
// TODO: finish implementation for $ENV/$CFG
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
	// ErrUnknownFlag is the unknown flag error.
	ErrUnknownFlag Error = "unknown flag"
	// ErrMissingArgument is the missing argument error.
	ErrMissingArgument Error = "missing argument"
	// ErrInvalidType is the invalid type error.
	ErrInvalidType Error = "invalid type"
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

// ldist is a levenshtein distance implementation.
func ldist[T []E, E cmp.Ordered](a, b T) int {
	m, n := len(a), len(b)
	v := make([]int, m+1)
	for i := range m + 1 {
		v[i] = i
	}
	var i, d, l int
	for j := range n {
		v[0] = j + 1
		for i, d = 0, j; i < m; i, d = i+1, l {
			if l = v[i+1]; a[i] != b[j] {
				d++
			}
			v[i+1] = min(v[i+1]+1, v[i]+1, d)
		}
	}
	return v[m]
}

// stripTestFlags strips flags starting with `-test.` from args.
func stripTestFlags(args []string) []string {
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

// prepend is a generic prepend.
func prepend[S ~[]E, E any](v S, s ...E) S {
	return append(s, v...)
}
