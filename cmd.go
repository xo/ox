package kobra

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// FlagSet is a set of command-line flag definitions.
type FlagSet struct {
	Flags []*Flag
}

// Flags creates a new flag set from the options.
func Flags() *FlagSet {
	return new(FlagSet)
}

// apply satisfies the [Option] interface.
func (fs *FlagSet) apply(v any) error {
	if cmd, ok := v.(*Command); ok {
		cmd.Flags = fs
		return nil
	}
	return ErrOptionAppliedToInvalidType
}

// Var adds a variable to the flag set.
func (fs *FlagSet) Var(name, desc string, opts ...Option) *FlagSet {
	if fs == nil {
		fs = Flags()
	}
	g, err := NewFlag(name, desc, opts...)
	if err != nil {
		panic(err)
	}
	fs.Flags = append(fs.Flags, g)
	return fs
}

// String adds a string variable to the flag set.
func (fs *FlagSet) String(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(StringT))...)
}

// Bytes adds a []byte variable to the flag set.
func (fs *FlagSet) Bytes(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(BytesT))...)
}

// Base64 adds a base64 encoded []byte variable to the flag set.
func (fs *FlagSet) Base64(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Base64T))...)
}

// Hex adds a hex encoded []byte variable to the flag set.
func (fs *FlagSet) Hex(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(HexT))...)
}

// Bool adds a bool variable to the flag set.
func (fs *FlagSet) Bool(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(BoolT))...)
}

// Byte adds a byte variable to the flag set.
func (fs *FlagSet) Byte(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(ByteT))...)
}

// Rune adds a rune variable to the flag set.
func (fs *FlagSet) Rune(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(RuneT))...)
}

// Int64 adds a int64 variable to the flag set.
func (fs *FlagSet) Int64(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Int64T))...)
}

// Int32 adds a int32 variable to the flag set.
func (fs *FlagSet) Int32(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Int32T))...)
}

// Int16 adds a int16 variable to the flag set.
func (fs *FlagSet) Int16(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Int16T))...)
}

// Int8 adds a int8 variable to the flag set.
func (fs *FlagSet) Int8(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Int8T))...)
}

// Int adds a int variable to the flag set.
func (fs *FlagSet) Int(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(IntT))...)
}

// Uint64 adds a uint64 variable to the flag set.
func (fs *FlagSet) Uint64(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Uint64T))...)
}

// Uint32 adds a uint32 variable to the flag set.
func (fs *FlagSet) Uint32(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Uint32T))...)
}

// Uint16 adds a uint16 variable to the flag set.
func (fs *FlagSet) Uint16(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Uint16T))...)
}

// Uint8 adds a uint8 variable to the flag set.
func (fs *FlagSet) Uint8(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Uint8T))...)
}

// Uint adds a uint variable to the flag set.
func (fs *FlagSet) Uint(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(UintT))...)
}

// Float64 adds a float64 variable to the flag set.
func (fs *FlagSet) Float64(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Float64T))...)
}

// Float32 adds a float32 variable to the flag set.
func (fs *FlagSet) Float32(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Float32T))...)
}

// Complex128 adds a complex128 variable to the flag set.
func (fs *FlagSet) Complex128(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Complex128T))...)
}

// Complex64 adds a complex64 variable to the flag set.
func (fs *FlagSet) Complex64(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(Complex64T))...)
}

// BigInt adds a [math/big.Int] variable to the flag set.
func (fs *FlagSet) BigInt(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(BigIntT))...)
}

// BigFloat adds a [math/big.Float] variable to the flag set.
func (fs *FlagSet) BigFloat(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(BigFloatT))...)
}

// BigRat adds a [math/big.Rat] variable to the flag set.
func (fs *FlagSet) BigRat(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(BigRatT))...)
}

// Timestamp adds a [time.Time] variable to the flag set in the expected format of
// [time.RFC3339].
func (fs *FlagSet) Timestamp(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(TimestampT))...)
}

// DateTime adds a [time.Time] variable to the flag set in the expected format of
// [time.DateTime].
func (fs *FlagSet) DateTime(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(DateTimeT))...)
}

// Date adds a [time.Time] variable to the flag set in the expected format of
// [time.DateOnly].
func (fs *FlagSet) Date(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(DateT))...)
}

// Duration adds a [time.Duration] variable to the flag set.
func (fs *FlagSet) Duration(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(DurationT))...)
}

// URL adds a [url.URL] variable to the flag set.
func (fs *FlagSet) URL(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(URLT))...)
}

// Addr adds a [netip.Addr] variable to the flag set.
func (fs *FlagSet) Addr(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(AddrT))...)
}

// AddrPort adds a [netip.AddrPort] variable to the flag set.
func (fs *FlagSet) AddrPort(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(AddrPortT))...)
}

// CIDR adds a [netip.Prefix] variable to the flag set.
func (fs *FlagSet) CIDR(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(CIDRT))...)
}

// Path adds a path variable to the flag set.
func (fs *FlagSet) Path(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(PathT))...)
}

// Count adds a count variable to the flag set.
func (fs *FlagSet) Count(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(CountT))...)
}

// UUID adds a uuid variable to the flag set.
func (fs *FlagSet) UUID(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(UUIDT))...)
}

// Color adds a color variable to the flag set.
func (fs *FlagSet) Color(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(ColorT))...)
}

// Slice adds a slice variable to the flag set.
func (fs *FlagSet) Slice(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(SliceT))...)
}

// Map adds a map variable to the flag set.
func (fs *FlagSet) Map(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(MapT))...)
}

// Hook adds a hook to the flag set.
func (fs *FlagSet) Hook(name, desc string, f func(context.Context) error, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(HookT), Default(f))...)
}

// Flag is a command-line flag variable definition.
type Flag struct {
	Type  Type
	Sub   Type
	Descs []Desc
	Def   any
	NoArg bool
	Keys  map[string]string
}

// NewFlag creates a new command-line flag.
func NewFlag(name, usage string, opts ...Option) (*Flag, error) {
	g := &Flag{
		Type: StringT,
		Descs: []Desc{{
			Name:  name,
			Usage: usage,
		}},
	}
	for _, o := range opts {
		if err := o.apply(g); err != nil {
			return nil, err
		}
	}
	if len(g.Descs) == 0 || len(g.Descs[0].Name) == 0 {
		return nil, ErrInvalidFlagName
	}
	// NOTE: forgive me linus, for i have hacked...
	for _, o := range types[g.Type].Opts {
		if err := o.apply(g); err != nil {
			return nil, err
		}
	}
	return g, nil
}

// Command is a command.
type Command struct {
	// Exec is the func to be executed.
	Exec func(context.Context, []string) error
	// Descs is the command's name and usage descriptions.
	Descs []Desc
	// Parent is the parent command.
	Parent *Command
	// Commands are sub commands.
	Commands []*Command
	// Flags are the command's flags.
	Flags *FlagSet
	// Args are the command's argument validation func's.
	Args []func([]string) error
	// OnErr indicates whether to continue, panic, or to error when flag
	// parsing fails.
	OnErr OnErr
}

// NewCommand creates a new command.
func NewCommand(f func(context.Context, []string) error, opts ...Option) (*Command, error) {
	c := &Command{
		Exec:  f,
		Descs: []Desc{{}},
	}
	for _, o := range opts {
		if err := o.apply(c); err != nil {
			return nil, err
		}
	}
	switch n := len(c.Descs[0].Name); {
	case n == 0 && c.Parent == nil:
		c.Descs[0].Name = filepath.Base(os.Args[0])
	case n == 0:
		return nil, ErrCommandUsageNotSet
	}
	return c, nil
}

// Run executes the command with the context, validating passed arguments.
func (cmd *Command) Run(ctx context.Context, args []string) error {
	// validate args
	for _, f := range cmd.Args {
		if err := f(args); err != nil {
			return err
		}
	}
	return cmd.Exec(ctx, args)
}

// Sub creates a sub command.
func (cmd *Command) Sub(f func(context.Context, []string) error, opts ...Option) error {
	c, err := NewCommand(f, prepend(opts, parent(cmd))...)
	if err != nil {
		return err
	}
	cmd.Commands = append(cmd.Commands, c)
	return nil
}

// Command returns the sub command with the name.
func (cmd *Command) Command(name string) *Command {
	for _, c := range cmd.Commands {
		for _, d := range c.Descs {
			if d.Name == name {
				return c
			}
		}
	}
	return nil
}

// Flag finds a flag from the command or the command's parents.
func (cmd *Command) Flag(name string) *Flag {
	if cmd.Flags != nil {
		for _, g := range cmd.Flags.Flags {
			for _, d := range g.Descs {
				if d.Name == name {
					return g
				}
			}
		}
	}
	if cmd.Parent != nil {
		return cmd.Parent.Flag(name)
	}
	return nil
}

// Get returns a sub command.
func (cmd *Command) Get(name string) *Command {
	for _, c := range cmd.Commands {
		for _, d := range c.Descs {
			if d.Name == name {
				return c
			}
		}
	}
	return nil
}

// Parse parses the command-line arguments, adding to the context.
func Parse(ctx context.Context, root *Command, args []string) (*Command, []string, Vars, error) {
	switch {
	case root.Parent != nil:
		return nil, nil, nil, ErrParseCanOnlyBeUsedWithRootCommand
	case len(args) == 0:
		return root, nil, Vars{}, nil
	}
	vars := make(Vars)
	cmd, args, err := parse(ctx, root, args, vars)
	return cmd, args, vars, err
}

// parse parses the args into m based on the flags on the command.
func parse(ctx context.Context, cmd *Command, args []string, vars Vars) (*Command, []string, error) {
	// fmt.Fprintf(os.Stdout, "args: %q\n", args)
	var v []string
	var s string
	var n int
	var err error
	for len(args) != 0 {
		switch s, n, args = args[0], len(args[0]), args[1:]; {
		case n == 0, n == 1, s[0] != '-':
			// lookup sub command
			var c *Command
			if len(v) == 0 {
				c = cmd.Command(s)
			}
			if c != nil {
				cmd = c
			} else {
				v = append(v, s)
			}
		case s == "--":
			return cmd, append(v, args...), nil
		case s[1] == '-':
			if args, err = parseLong(ctx, cmd, s, args, vars); err != nil {
				return nil, nil, err
			}
		default:
			if args, err = parseShort(ctx, cmd, s, args, vars); err != nil {
				return nil, nil, err
			}
		}
	}
	return cmd, v, nil
}

// parseLong parses a long flag ('--arg' '--arg v' '--arg k=v' '--arg=' '--arg=v').
func parseLong(ctx context.Context, cmd *Command, s string, args []string, vars Vars) ([]string, error) {
	arg, value, ok := strings.Cut(strings.TrimPrefix(s, "--"), "=")
	g := cmd.Flag(arg)
	switch {
	case g == nil:
		return nil, newFlagError(arg, false, ErrUnknownFlag)
	case ok: // --arg=v
	case g.NoArg: // --arg
		value = toBoolString(g.Def)
	case len(args) != 0: // --arg v
		value, args = args[0], args[1:]
	default: // missing argument to --arg
		return nil, newFlagError(arg, false, ErrMissingArgument)
	}
	if err := vars.Set(ctx, g, value); err != nil {
		return nil, newFlagError(arg, false, err)
	}
	return args, nil
}

// parseShort parses short flags ('-a' '-aaa' '-av' '-a v' '-a=' '-a=v').
func parseShort(ctx context.Context, cmd *Command, s string, args []string, vars Vars) ([]string, error) {
	for v := []rune(s[1:]); len(v) != 0; v = v[1:] {
		arg := string(v[0])
		switch g, n := cmd.Flag(arg), len(v[1:]); {
		case g == nil:
			return nil, newFlagError(arg, true, ErrUnknownFlag)
		case g.NoArg: // -a
			if err := vars.Set(ctx, g, toBoolString(g.Def)); err != nil {
				return nil, newFlagError(arg, true, err)
			}
		case n == 0 && len(args) == 0: // missing argument to -a
			return nil, newFlagError(arg, true, ErrMissingArgument)
		case n != 0: // -a=, -a=v, -av
			if slices.Index(v, '=') == 1 {
				v = v[1:]
			}
			if err := vars.Set(ctx, g, string(v[1:])); err != nil {
				return nil, newFlagError(arg, true, err)
			}
			return args, nil
		default: // -a v
			if err := vars.Set(ctx, g, args[0]); err != nil {
				return nil, newFlagError(arg, true, err)
			}
			return args[1:], nil
		}
	}
	return args, nil
}

// Desc contains a command/flag description.
type Desc struct {
	Name       string
	Usage      string
	Short      string
	Key        map[string]string
	Hidden     bool
	Deprecated bool
}

// apply satisfies the [Option] interface.
func (d Desc) apply(val any) error {
	switch v := val.(type) {
	case *Command:
		v.Descs = append(v.Descs, d)
	case *Flag:
		v.Descs = append(v.Descs, d)
	default:
		return ErrOptionAppliedToInvalidType
	}
	return nil
}

// FlagError is a flag parse error.
type FlagError struct {
	Arg   string
	Short bool
	Err   error
}

// newFlagError creates a flag error.
func newFlagError(arg string, short bool, err error) error {
	return &FlagError{
		Arg:   arg,
		Short: short,
		Err:   err,
	}
}

// Error satisfies the [error] interface.
func (err *FlagError) Error() string {
	if err.Short {
		return "-" + err.Arg + ": " + err.Err.Error()
	}
	return "--" + err.Arg + ": " + err.Err.Error()
}

// Unwrap satisfies the [errors.Unwrap] interface.
func (err *FlagError) Unwrap() error {
	return err.Err
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
	if v, ok := val.(*Command); ok {
		v.OnErr = e
		return nil
	}
	return ErrOptionAppliedToInvalidType
}

// Handle handles an error.
func (e OnErr) Handle(ctx context.Context, err error) {
	switch {
	case errors.Is(err, ErrExit), e == OnErrContinue:
	case e == OnErrExit:
		fmt.Fprintln(Stderr(ctx), "error:", err)
		os.Exit(1)
	case e == OnErrPanic:
		panic(err)
	}
}
