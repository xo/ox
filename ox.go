// Package ox is a context-based, minimal dependency Go (and TinyGo compatible)
// command-line flag and argument parsing library.
package ox

//go:generate stringer -type OnErr

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// Run creates a [Command] for f using [os.Args] by default, unless arguments
// were specified using a [Option] such as [RunArgs].
func Run[T ExecFunc](f T, opts ...Option) {
	ctx := &RunContext{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Args:   os.Args[1:],
	}
	for _, o := range opts {
		if err := o.apply(ctx); err != nil {
			ctx.Handle(err)
		}
	}
	if ctx.Vars == nil {
		ctx.Vars = make(Vars)
	}
	root, err := NewCommand(f, opts...)
	if err != nil {
		ctx.Handle(err)
	}
	cmd, args, err := Parse(root, ctx.Args, ctx.Vars)
	if err != nil {
		ctx.Handle(err)
	}
	if err := cmd.Validate(args); err != nil {
		ctx.Handle(err)
	}
	if err := cmd.Exec(context.Background(), args); err != nil {
		ctx.Handle(err)
	}
}

// RunContext are run context options.
type RunContext struct {
	Args   []string
	Stdin  io.Writer
	Stdout io.Writer
	Stderr io.Writer
	OnErr  OnErr
	Vars   Vars
}

// Handle handles an error.
func (ctx *RunContext) Handle(err error) {
	var w io.Writer = os.Stderr
	if ctx.Stderr != nil {
		w = ctx.Stderr
	}
	switch {
	case errors.Is(err, ErrExit), ctx.OnErr == OnErrContinue:
	case ctx.OnErr == OnErrExit:
		fmt.Fprintln(w, "error:", err)
		os.Exit(1)
	case ctx.OnErr == OnErrPanic:
		panic(err)
	}
}

// OnErr is the on error handling option type.
type OnErr uint8

// On error handling options.
const (
	OnErrExit OnErr = iota
	OnErrContinue
	OnErrPanic
)

// apply satisfies the [Option] interface.
func (e OnErr) apply(val any) error {
	switch v := val.(type) {
	case *Command:
		v.OnErr = e
	default:
		return fmt.Errorf("%s option: %w", e, ErrAppliedToInvalidType)
	}
	return nil
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
	// ErrIsNil is the is nil error.
	// ErrIsNil Error = "is nil"
	// ErrTypeMismatch is the type mismatch error.
	ErrTypeMismatch Error = "type mismatch"
	// ErrExit is the exit error.
	ErrExit Error = "exit"
)

var (
	// DefaultLayout is the default timestamp layout.
	DefaultLayout = time.RFC3339
	// DefaultTagName is the default struct tag name used in [FromFlags] and
	// related func's for building flag options.
	DefaultTagName = "ox"
)

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
