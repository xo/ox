// Package ox is a minimal dependency Go (and TinyGo compatible) command-line
// flag and argument parsing library.
package ox

//go:generate stringer -type OnErr

import (
	"cmp"
	"context"
	"time"
)

var (
	// DefaultTagName is the default struct tag name used in [FromFlags] and
	// related func's.
	DefaultTagName = "ox"
	// DefaultContext is the default [context.Context].
	DefaultContext = context.Background()
	// DefaultLayout is the default timestamp layout used for formatting and
	// parsing [TimestampT] values.
	DefaultLayout = time.RFC3339
)

// Run creates a [Context] and a [Command] based on the options and executes
// the command.
func Run(opts ...Option) {
	RunContext(DefaultContext, opts...)
}

// RunContext creates a [Command] for f using [os.Args] by default, unless
// arguments were specified using a [ContextOption].
func RunContext(ctx context.Context, opts ...Option) {
	c, err := NewContext(opts...)
	if err != nil {
		c.Handle(err)
	}
	root, err := NewCommand(opts...)
	if err != nil {
		c.Handle(err)
	}
	cmd, args, err := Parse(root, c.Args, c.Vars)
	if err != nil {
		c.Handle(err)
	}
	if err := cmd.Validate(args); err != nil {
		c.Handle(err)
	}
	if err := cmd.Exec(WithContext(ctx, c), args); err != nil {
		c.Handle(err)
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
