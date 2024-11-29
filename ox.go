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
	"strings"
	"time"

	"github.com/xo/ox/strcase"
)

var (
	// DefaultTagName is the default struct tag name used in [FromFlags] and
	// related func's.
	DefaultTagName = "ox"
	// DefaultContext is the default [context.Context].
	DefaultContext = context.Background()
	// DefaultLayout is the default timestamp layout used for formatting and
	// parsing [Time] values.
	DefaultLayout = time.RFC3339
	// DefaultFlagNameMaper is the default flag name mapper.
	DefaultFlagNameMapper = func(s string) string {
		return strings.ReplaceAll(strcase.CamelToSnake(s), "_", "-")
	}
	// DefaultVersionName is the default version command/flag name.
	DefaultVersionName = "version"
	// DefaultVersionShort is the default version short flag name.
	DefaultVersionShort = "v"
	// DefaultHelpName is the default help command/flag name.
	DefaultHelpName = "help"
	// DefaultHelpShort is the default help short flag name.
	DefaultHelpShort = "?"
	// DefaultWrapWidth is the default wrap width.
	DefaultWrapWidth = 95
	// DefaultWrap wraps a line of text with [Wrap] using [DefaultWrapWidth]
	// for the width.
	DefaultWrap = func(s string, prefixWidth int) string {
		return Wrap(s, DefaultWrapWidth, prefixWidth)
	}
	// DefaultErrorMessagePrefix is the default 'error: ' message prefix.
	DefaultErrorMessagePrefix = "error: "
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
	if err != nil {
		if c.Handle(err) {
			return
		}
	}
	if c.Root, err = NewCommand(opts...); err != nil {
		if c.Handle(err) {
			return
		}
	}
	if c.Exec, c.Args, err = Parse(c.Root, c.Args, c.Vars); err != nil {
		if c.Handle(err) {
			return
		}
	}
	if err = c.Exec.Validate(c.Args); err != nil {
		if c.Handle(err) {
			return
		}
	}
	if err = c.Exec.Exec(WithContext(ctx, c), c.Args); err != nil {
		if c.Handle(err) {
			return
		}
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
}

// NewContext creates a new run context.
func NewContext(opts ...Option) (*Context, error) {
	ctx := &Context{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Args:   os.Args[1:],
	}
	if err := applyOpts(ctx, opts...); err != nil {
		return ctx, err
	}
	if ctx.Vars == nil {
		ctx.Vars = make(Vars)
	}
	if ctx.Handle == nil {
		ctx.Handle = func(err error) bool {
			var w io.Writer = os.Stderr
			if ctx.Stderr != nil {
				w = ctx.Stderr
			}
			switch {
			case errors.Is(err, ErrExit):
			case errors.Is(err, ErrHelp):
			case ctx.OnErr == OnErrContinue:
				return false
			case ctx.OnErr == OnErrExit:
				fmt.Fprintf(w, "%s%v\n", DefaultErrorMessagePrefix, err)
				os.Exit(1)
			case ctx.OnErr == OnErrPanic:
				panic(err)
			}
			return true
		}
	}
	return ctx, nil
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
	// ErrHelp is the help error.
	ErrHelp Error = "help"
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
		if left < len(word)+1 {
			wrapped += "\n" + prefix + word
			left = width - len(word)
		} else {
			wrapped += " " + word
			left -= 1 + len(word)
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

// prepend is a generic prepend.
func prepend[S ~[]E, E any](v S, s ...E) S {
	return append(s, v...)
}
