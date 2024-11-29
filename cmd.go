package ox

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Command is a command.
type Command struct {
	// Parent is the parent command.
	Parent *Command
	// Exec is the func to be executed.
	Exec ExecFunc
	// Name is the command name.
	Name string
	// Usage is the command usage.
	Usage string
	// Aliases are the command aliases.
	Aliases []string
	// Flags are the command's flags.
	Flags *FlagSet
	// Commands are sub commands.
	Commands []*Command
	// Args are the command's argument validation func's.
	Args []func([]string) error
	// OnErr indicates whether to continue, panic, or to error when
	// flag/argument parsing fails.
	OnErr OnErr
	// Hidden indicates the command is hidden from help output.
	Hidden bool
	// Deprecated indicates the command is deprecated.
	Deprecated bool
	// Help is the help emitter.
	Help io.WriterTo
}

// NewCommand creates a new command.
func NewCommand(opts ...Option) (*Command, error) {
	cmd := new(Command)
	if err := applyOpts(cmd, opts...); err != nil {
		return nil, err
	}
	switch noName := cmd.Name == ""; {
	case noName && cmd.Parent == nil:
		cmd.Name = filepath.Base(os.Args[0])
	case noName:
		return nil, fmt.Errorf("command: %w", ErrUsageNotSet)
	}
	if err := applyPostOpts(cmd, opts...); err != nil {
		return nil, err
	}
	return cmd, nil
}

// Sub creates a sub command.
func (cmd *Command) Sub(opts ...Option) error {
	sub, err := NewCommand(prepend(opts, Parent(cmd))...)
	if err != nil {
		return err
	}
	cmd.Commands = append(cmd.Commands, sub)
	return nil
}

// Tree returns the parents for the command.
func (cmd *Command) Tree() []string {
	var v []string
	for c := cmd; c != nil; c = c.Parent {
		v = append(v, c.Name)
	}
	slices.Reverse(v)
	return v
}

// Path returns the execution path for the command.
func (cmd *Command) Path() []string {
	return cmd.Tree()[1:]
}

// Populate populates vars with all the command's flags values, overwriting any
// set variables if applicable. When all is true, all flag values will be
// populated, otherwise only flags with default values will be. When overwrite
// is true, existing vars will be set to either to flag's empty or default
// value.
func (cmd *Command) Populate(all, overwrite bool, vars Vars) error {
	if cmd.Flags == nil {
		return nil
	}
	for _, g := range cmd.Flags.Flags {
		if _, ok := vars[g.Name]; ok && overwrite {
			delete(vars, g.Name)
		}
		var value string
		switch {
		case g.Type == HookT, g.Def == nil && !all:
			continue
		case g.Def != nil:
			var err error
			if value, err = asString[string](g.Def); err != nil {
				return err
			}
		}
		if err := vars.Set(g, value, false); err != nil {
			return fmt.Errorf("cannot populate %s with %q: %w", g.Name, value, err)
		}
	}
	return nil
}

// Command returns the sub command with the name.
func (cmd *Command) Command(name string) *Command {
	for _, c := range cmd.Commands {
		if c.Name == name || slices.Contains(c.Aliases, name) {
			return c
		}
	}
	return nil
}

// Flag finds a flag from the command or the command's parents.
func (cmd *Command) Flag(name string, short bool) *Flag {
	if cmd.Flags != nil {
		if i := slices.IndexFunc(cmd.Flags.Flags, func(g *Flag) bool {
			return g.Name == name && !short ||
				g.Short == name && short ||
				slices.ContainsFunc(g.Aliases, func(s string) bool {
					return s == name && len(s) == 1 == short
				})
		}); i != -1 {
			return cmd.Flags.Flags[i]
		}
	}
	if cmd.Parent != nil {
		return cmd.Parent.Flag(name, short)
	}
	return nil
}

// Validate validates the passed args.
func (cmd *Command) Validate(args []string) error {
	// validate args
	for _, f := range cmd.Args {
		if err := f(args); err != nil {
			return newCommandError(cmd.Name, err)
		}
	}
	return nil
}

// FlagSet is a set of command-line flag definitions.
type FlagSet struct {
	Flags []*Flag
}

// Flags creates a new flag set from the options.
func Flags() *FlagSet {
	return new(FlagSet)
}

// Option returns a [CommandOption] for the flag set.
func (fs *FlagSet) Option() option {
	return option{
		name: "FlagSet",
		ctx: func(*Context) error {
			return nil
		},
		cmd: func(cmd *Command) error {
			cmd.Flags = fs
			return nil
		},
	}
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

// Path adds a path variable to the flag set.
func (fs *FlagSet) Path(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(PathT))...)
}

// Count adds a count variable to the flag set.
func (fs *FlagSet) Count(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(CountT))...)
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

// URL adds a [url.URL] variable to the flag set.
func (fs *FlagSet) URL(name, desc string, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(URLT))...)
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
	// Type is the flag [Type].
	Type Type
	// Key is the flag's key type when the flag is a [MapT].
	Key Type
	// Elem is the flag's element type when the flag is a [SliceT] or [MapT].
	Elem Type
	// Name is the flag's name.
	Name string
	// Usage is the flag's usage.
	Usage string
	// Short is the flag's short name.
	Short string
	// Aliases are the flag's aliases.
	Aliases []string
	// Def is the default value for the flag.
	Def any
	// NoArg indicates that flag does takes an optional argument.
	NoArg bool
	// NoArgDef is the default value for the flag when no argument was passed for the flag.
	NoArgDef any
	// Binds are variables that will be set.
	Binds []Binder
	// Hidden indicates the command is hidden from help output.
	Hidden bool
	// Deprecated indicates the command is deprecated.
	Deprecated bool
	// Keys are config look up keys.
	Keys map[string]string
}

// NewFlag creates a new command-line flag.
func NewFlag(name, usage string, opts ...Option) (*Flag, error) {
	g := &Flag{
		Type:  StringT,
		Key:   StringT,
		Elem:  StringT,
		Name:  name,
		Usage: usage,
	}
	if err := applyOpts(g, opts...); err != nil {
		return nil, err
	}
	if extra, ok := typeFlagOpts[g.Type]; ok {
		if err := applyOpts(g, extra...); err != nil {
			return nil, err
		}
	}
	if g.Name == "" {
		return nil, ErrInvalidFlagName
	}
	if g.NoArg && g.NoArgDef == nil {
		return nil, fmt.Errorf("flag %s: %w: NoArgDef cannot be nil", g.Name, ErrInvalidValue)
	}
	return g, nil
}

// FlagFrom creates flags for the value.
func FlagsFrom[T *E, E any](val T) ([]*Flag, error) {
	v := reflect.ValueOf(val).Elem()
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%w: %T is not a *struct", ErrInvalidType, val)
	}
	typ := v.Type()
	var flags []*Flag
	for i := range typ.NumField() {
		// make sure the field is exported
		f := typ.Field(i)
		if r := []rune(f.Name); !unicode.IsUpper(r[0]) {
			continue
		}
		s, ok := f.Tag.Lookup(DefaultTagName)
		if !ok {
			continue
		}
		tag := strings.Split(s, ",")
		if len(tag) == 0 || tag[0] == "-" {
			continue
		}
		// build flag opts
		opts, err := buildFlagOpts(typ, v.Field(i).Addr(), f.Name, tag[1:])
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", f.Name, err)
		}
		// TODO: add name mapping here
		g, err := NewFlag(DefaultFlagNameMapper(f.Name), tag[0], opts...)
		if err != nil {
			return nil, fmt.Errorf("field %s: cannot create flag: %w", f.Name, err)
		}
		flags = append(flags, g)
	}
	return flags, nil
}

// New creates a new value for the flag's type.
func (g *Flag) New() (Value, error) {
	switch g.Type {
	case SliceT:
		return NewSlice(g.Elem), nil
	case MapT:
		return NewMap(g.Key, g.Elem)
	}
	return g.Type.New()
}

// ExecType is the interface for func's that can be used with [Run],
// [NewCommand], and [Exec].
type ExecType interface {
	func(context.Context, []string) error |
		func(context.Context, []string) |
		func(context.Context) error |
		func(context.Context) |
		func([]string) error |
		func([]string) |
		func() error |
		func()
}

// ExecFunc wraps a exec func.
type ExecFunc func(context.Context, []string) error

// NewExec creates a [ExecFunc] func.
func NewExec[T ExecType](f T) (ExecFunc, error) {
	var v any = f
	switch f := v.(type) {
	case func(context.Context, []string) error:
		return f, nil
	case func(context.Context, []string):
		return func(ctx context.Context, args []string) error {
			f(ctx, args)
			return nil
		}, nil
	case func(context.Context) error:
		return func(ctx context.Context, _ []string) error {
			return f(ctx)
		}, nil
	case func(context.Context):
		return func(ctx context.Context, _ []string) error {
			f(ctx)
			return nil
		}, nil
	case func([]string) error:
		return func(_ context.Context, args []string) error {
			return f(args)
		}, nil
	case func([]string):
		return func(_ context.Context, args []string) error {
			f(args)
			return nil
		}, nil
	case func() error:
		return func(context.Context, []string) error {
			return f()
		}, nil
	case func():
		return func(context.Context, []string) error {
			f()
			return nil
		}, nil
	}
	return nil, fmt.Errorf("%w: invalid exec func %T", ErrInvalidType, f)
}

// buildFlagOpts builds flag options for v from the passed struct tag values.
func buildFlagOpts(typ reflect.Type, value reflect.Value, name string, s []string) ([]Option, error) {
	var b *bool
	var opts []Option
	for _, opt := range s {
		key, val, _ := strings.Cut(opt, ":")
		switch key {
		case "name":
			opts = append(opts, Name(val))
		case "type":
			opts = append(opts, Type(val))
		case "usage":
			name, usage, _ := strings.Cut(val, "|")
			opts = append(opts, Usage(name, usage))
		case "short":
			opts = append(opts, Short(val))
		case "alias":
			name, usage, _ := strings.Cut(val, "|")
			opts = append(opts, Alias(name, usage))
		case "default":
			opts = append(opts, Default(val))
		case "key":
			typ, key, _ := strings.Cut(val, "|")
			opts = append(opts, Key(typ, key))
		case "set":
			for i := range typ.NumField() {
				field := typ.Field(i)
				if field.Name == name || field.Name != val {
					continue
				}
				break
			}
		default:
			return nil, fmt.Errorf("%w: %q", ErrUnknownTagOption, key)
		}
	}
	// prepend bind to opt
	return prepend(opts, BindRef(value, b)), nil
}

// newCommandError creates a command error.
func newCommandError(name string, err error) error {
	return fmt.Errorf("command %s: %w", name, err)
}

// newFlagError creates a flag error.
func newFlagError(name string, err error) error {
	if utf8.RuneCountInString(name) == 1 {
		return fmt.Errorf("-%s: %w", name, err)
	}
	return fmt.Errorf("--%s: %w", name, err)
}
