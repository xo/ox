package ox

import (
	"context"
	"fmt"
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
	// Exec is the func to be executed.
	Exec Exec
	// Descs is the command's name/usage descriptions.
	Descs []Desc
	// Parent is the parent command.
	Parent *Command
	// Commands are sub commands.
	Commands []*Command
	// Flags are the command's flags.
	Flags *FlagSet
	// Args are the command's argument validation func's.
	Args []func([]string) error
	// OnErr indicates whether to continue, panic, or to error when
	// flag/argument parsing fails.
	OnErr OnErr
}

// NewCommand creates a new command.
func NewCommand[T ExecFunc](f T, opts ...Option) (*Command, error) {
	cmd := &Command{
		Exec:  NewExec(f),
		Descs: make([]Desc, 1),
	}
	for _, o := range opts {
		if err := o.apply(cmd); err != nil {
			return nil, err
		}
	}
	switch noName := cmd.Name() == ""; {
	case noName && cmd.Parent == nil:
		cmd.Descs[0].Name = filepath.Base(os.Args[0])
	case noName:
		return nil, fmt.Errorf("command: %w", ErrUsageNotSet)
	}
	return cmd, nil
}

// Sub creates a sub command.
func (cmd *Command) Sub(f func(context.Context, []string) error, opts ...Option) error {
	sub, err := NewCommand(f, prepend(opts, Parent(cmd))...)
	if err != nil {
		return err
	}
	cmd.Commands = append(cmd.Commands, sub)
	return nil
}

// Command returns the sub command with the name.
func (cmd *Command) Command(name string) *Command {
	for _, sub := range cmd.Commands {
		for _, d := range sub.Descs {
			if d.Name == name {
				return sub
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

// Validate validates the passed args.
func (cmd *Command) Validate(args []string) error {
	// validate args
	for _, f := range cmd.Args {
		if err := f(args); err != nil {
			return newCommandError(cmd.Name(), err)
		}
	}
	return nil
}

// Tree returns the parents of the command.
func (cmd *Command) Tree() []string {
	var v []string
	for c := cmd; c != nil; {
		v, c = append(v, c.Name()), c.Parent
	}
	slices.Reverse(v)
	return v
}

// Name returns the command's name.
func (cmd *Command) Name() string {
	if len(cmd.Descs) != 0 {
		return cmd.Descs[0].Name
	}
	return ""
}

// FlagSet is a set of command-line flag definitions.
type FlagSet struct {
	Flags []*Flag
}

// Flags creates a new flag set from the options.
func Flags(opts ...Option) *FlagSet {
	return new(FlagSet)
}

// apply satisfies the [Option] interface.
func (fs *FlagSet) apply(val any) error {
	switch v := val.(type) {
	case *Command:
		v.Flags = fs
	case *Context:
	default:
		return fmt.Errorf("FlagSet: %w", ErrAppliedToInvalidType)
	}
	return nil
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
	Key   Type
	Sub   Type
	Descs []Desc
	Def   any
	NoArg bool
	Binds []BoundValue
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
	if opts, ok := typeFlagOpts[g.Type]; ok {
		for _, o := range opts {
			if err := o.apply(g); err != nil {
				return nil, err
			}
		}
	}
	if g.Name() == "" {
		return nil, ErrInvalidFlagName
	}
	return g, nil
}

// FlagFrom creates flags for the value.
func FlagsFrom[T *E, E any](val T) ([]*Flag, error) {
	v := reflect.ValueOf(val).Elem()
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%w: not a *struct", ErrInvalidType)
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
		g, err := NewFlag(strings.ToLower(f.Name), tag[0], opts...)
		if err != nil {
			return nil, fmt.Errorf("field %s: cannot create flag: %w", f.Name, err)
		}
		flags = append(flags, g)
	}
	return flags, nil
}

// Name returns the flag's primary name.
func (g *Flag) Name() string {
	if len(g.Descs) != 0 {
		return g.Descs[0].Name
	}
	return ""
}

// Short returns the flag's first short alias, and whether or not a short alias
// was found.
func (g *Flag) Short() (string, bool) {
	if 1 < len(g.Descs) {
		for _, d := range g.Descs[1:] {
			if utf8.RuneCountInString(d.Name) == 1 {
				return d.Name, true
			}
		}
	}
	return "", false
}

// Parse parses the command-line arguments into vars.
func Parse(root *Command, args []string, vars Vars) (*Command, []string, error) {
	if root.Parent != nil {
		return nil, nil, fmt.Errorf("Parse: %w", ErrCanOnlyBeUsedWithRootCommand)
	}
	if err := Populate(root, false, false, vars); err != nil {
		return nil, nil, newCommandError(root.Name(), err)
	}
	if len(args) == 0 {
		return root, nil, nil
	}
	return parse(root, args, vars)
}

// Populate populates vars with all the command's flags values, and overwriting
// if applicable. When all is true, all flag values will be populated,
// otherwise only flags with default values will be. When overwrite is true,
// existing vars will be set to either to flag's empty or default value.
func Populate(cmd *Command, all, overwrite bool, vars Vars) error {
	if cmd.Flags == nil {
		return nil
	}
	for _, g := range cmd.Flags.Flags {
		name := g.Name()
		if _, ok := vars[name]; ok && overwrite {
			delete(vars, name)
		}
		var value string
		switch {
		case g.Type == HookT, g.Def == nil && !all:
			continue
		case g.Def != nil:
			s, err := asString[string](g.Def, g.Type.Layout())
			if err != nil {
				return err
			}
			value = s
		}
		if err := vars.Set(g, value, false); err != nil {
			return fmt.Errorf("cannot populate %s with %q: %w", name, value, err)
		}
	}
	return nil
}

// parse parses the args into m based on the flags on the command.
func parse(cmd *Command, args []string, vars Vars) (*Command, []string, error) {
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
				if err := Populate(c, false, false, vars); err != nil {
					return nil, nil, newCommandError(c.Name(), err)
				}
				cmd = c
			} else {
				v = append(v, s)
			}
		case s == "--":
			return cmd, append(v, args...), nil
		case s[1] == '-':
			if args, err = parseLong(cmd, s, args, vars); err != nil {
				return nil, nil, err
			}
		default:
			if args, err = parseShort(cmd, s, args, vars); err != nil {
				return nil, nil, err
			}
		}
	}
	return cmd, v, nil
}

// parseLong parses a long flag ('--arg' '--arg v' '--arg k=v' '--arg=' '--arg=v').
func parseLong(cmd *Command, s string, args []string, vars Vars) ([]string, error) {
	arg, value, ok := strings.Cut(strings.TrimPrefix(s, "--"), "=")
	g := cmd.Flag(arg)
	switch {
	case g == nil:
		return nil, newFlagError(arg, ErrUnknownFlag)
	case ok: // --arg=v
	case g.NoArg: // --arg
		value = toBoolString(g.Def)
	case len(args) != 0: // --arg v
		value, args = args[0], args[1:]
	default: // missing argument to --arg
		return nil, newFlagError(arg, ErrMissingArgument)
	}
	if err := vars.Set(g, value, true); err != nil {
		return nil, newFlagError(arg, err)
	}
	return args, nil
}

// parseShort parses short flags ('-a' '-aaa' '-av' '-a v' '-a=' '-a=v').
func parseShort(cmd *Command, s string, args []string, vars Vars) ([]string, error) {
	for v := []rune(s[1:]); len(v) != 0; v = v[1:] {
		arg := string(v[0])
		switch g, n := cmd.Flag(arg), len(v[1:]); {
		case g == nil:
			return nil, newFlagError(arg, ErrUnknownFlag)
		case g.NoArg: // -a
			var value string
			if slices.Index(v, '=') == 1 {
				value, v = string(v[2:]), v[len(v)-1:]
			} else {
				value = toBoolString(g.Def)
			}
			if err := vars.Set(g, value, true); err != nil {
				return nil, newFlagError(arg, err)
			}
		case n == 0 && len(args) == 0: // missing argument to -a
			return nil, newFlagError(arg, ErrMissingArgument)
		case n != 0: // -a=, -a=v, -av
			if slices.Index(v, '=') == 1 {
				v = v[1:]
			}
			if err := vars.Set(g, string(v[1:]), true); err != nil {
				return nil, newFlagError(arg, err)
			}
			return args, nil
		default: // -a v
			if err := vars.Set(g, args[0], true); err != nil {
				return nil, newFlagError(arg, err)
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
		return fmt.Errorf("Desc: %w", ErrAppliedToInvalidType)
	}
	return nil
}

// buildFlagOpts builds flag options for v from the passed struct tag values.
func buildFlagOpts(typ reflect.Type, val reflect.Value, name string, s []string) ([]Option, error) {
	var b *bool
	var opts []Option
	for _, opt := range s {
		key, value, _ := strings.Cut(opt, ":")
		switch key {
		case "name":
			opts = append(opts, Name(value))
		case "type":
			opts = append(opts, Type(value))
		case "usage":
			name, usage, _ := strings.Cut(value, "|")
			opts = append(opts, Usage(name, usage))
		case "short":
			opts = append(opts, Short(value))
		case "alias":
			name, usage, _ := strings.Cut(value, "|")
			opts = append(opts, Alias(name, usage))
		case "default":
			opts = append(opts, Default(value))
		case "key":
			typ, key, _ := strings.Cut(value, "|")
			opts = append(opts, Key(typ, key))
		case "set":
			for i := range typ.NumField() {
				field := typ.Field(i)
				if field.Name == name || field.Name != value {
					continue
				}
				break
			}
		default:
			return nil, fmt.Errorf("%w: %q", ErrUnknownTagOption, key)
		}
	}
	// prepend bind to opt
	return prepend(opts, BindValue(val, b)), nil
}

// ExecFunc is the interface for func's that can be used with [Run] and
// [NewCommand].
type ExecFunc interface {
	func(context.Context, []string) error |
		func([]string) error |
		func(context.Context) error |
		func() error |
		func()
}

// Exec wraps a exec func.
type Exec func(context.Context, []string) error

// NewExec creates a [Exec] func.
func NewExec[T ExecFunc](f T) Exec {
	var v any = f
	switch f := v.(type) {
	case func(context.Context, []string) error:
		return f
	case func([]string) error:
		return func(_ context.Context, args []string) error {
			return f(args)
		}
	case func(context.Context) error:
		return func(ctx context.Context, _ []string) error {
			return f(ctx)
		}
	case func() error:
		return func(context.Context, []string) error {
			return f()
		}
	case func():
		return func(context.Context, []string) error {
			f()
			return nil
		}
	}
	return func(context.Context, []string) error {
		return fmt.Errorf("%w: no exec func", ErrInvalidType)
	}
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
