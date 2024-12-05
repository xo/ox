package ox

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
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
	// Section is the help section.
	Section int
	// Help is the help emitter.
	Help io.WriterTo
	// Special is the special value.
	Special string
	// Templates are the completion template files.
	Templates fs.FS
}

// NewCommand creates a new command.
func NewCommand(opts ...Option) (*Command, error) {
	cmd := &Command{
		Section: -1,
	}
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

// RootName returns the root name for the command.
func (cmd *Command) RootName() string {
	return cmd.Tree()[0]
}

// Populate populates vars with all the command's flags values, overwriting any
// set variables if applicable. When all is true, all flag values will be
// populated, otherwise only flags with default values will be. When overwrite
// is true, existing vars will be set to either to flag's empty or default
// value.
func (cmd *Command) Populate(ctx *Context, all, overwrite bool, vars Vars) error {
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
			if value, err = ctx.Expand(g.Def); err != nil {
				return err
			}
		}
		if err := vars.Set(ctx, g, value, false); err != nil {
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

// CommandSpecial returns the sub command with the special value.
func (cmd *Command) CommandSpecial(special string) *Command {
	for _, c := range cmd.Commands {
		if c.Special == special {
			return c
		}
	}
	return nil
}

// Flag finds a flag from the command or the command's parents.
func (cmd *Command) Flag(name string, parents, short bool) *Flag {
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
	if parents && cmd.Parent != nil {
		return cmd.Parent.Flag(name, parents, short)
	}
	return nil
}

// FlagSpecial returns the flag with a special value.
func (cmd *Command) FlagSpecial(special string) *Flag {
	if cmd.Flags != nil && len(cmd.Flags.Flags) != 0 {
		for _, g := range cmd.Flags.Flags {
			if g.Special == special {
				return g
			}
		}
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

// HelpContext adds the context to the command's help.
func (cmd *Command) HelpContext(ctx *Context) io.WriterTo {
	help := cmd.Help
	if help == nil {
		help, _ = NewCommandHelp(cmd)
	}
	if v, ok := help.(interface{ SetContext(*Context) }); ok {
		v.SetContext(ctx)
	}
	return help
}

// WriteTo satisfies the [io.WriterTo] interface.
func (cmd *Command) WriteTo(w io.Writer) (int64, error) {
	help := cmd.Help
	if help == nil {
		var err error
		if help, err = NewCommandHelp(cmd); err != nil {
			return 0, err
		}
	}
	return help.WriteTo(w)
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
func (fs *FlagSet) Hook(name, desc string, f any, opts ...Option) *FlagSet {
	return fs.Var(name, desc, prepend(opts, Option(HookT), NoArg(true, ""), Default(f))...)
}

// Flag is a command-line flag variable definition.
type Flag struct {
	// Type is the flag [Type].
	Type Type
	// MapKey is the flag's map key type when the flag is a [MapT].
	MapKey Type
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
	// NoArgDef is the default value for the flag when no argument was passed
	// for the flag.
	NoArgDef any
	// Binds are variables that will be set.
	Binds []Binder
	// Hidden indicates the command is hidden from help output.
	Hidden bool
	// Deprecated indicates the command is deprecated.
	Deprecated bool
	// Keys are config look up keys.
	Keys map[string]string
	// Spec is the help spec.
	Spec string
	// Section is the help section.
	Section int
	// Special is the special value.
	Special string
}

// NewFlag creates a new command-line flag.
func NewFlag(name, usage string, opts ...Option) (*Flag, error) {
	g := &Flag{
		Type:    StringT,
		MapKey:  StringT,
		Elem:    StringT,
		Name:    name,
		Usage:   usage,
		Section: -1,
	}
	if err := applyOpts(g, opts...); err != nil {
		return nil, err
	}
	if extra, ok := typeFlagOpts[g.Type]; ok {
		if err := applyOpts(g, extra...); err != nil {
			return nil, err
		}
	}
	switch {
	case g.Name == "":
		return nil, ErrInvalidFlagName
	case g.NoArg && g.NoArgDef == nil:
		return nil, fmt.Errorf("flag %s: %w: NoArgDef cannot be nil", g.Name, ErrInvalidValue)
	case g.Type == HookT && g.Def == nil:
		return nil, fmt.Errorf("flag %s: %w: Def cannot be nil", g.Name, ErrInvalidValue)
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
		// check field is exported
		f := typ.Field(i)
		s, ok := f.Tag.Lookup(DefaultTagName)
		if !ok {
			continue
		}
		if r := []rune(f.Name); !unicode.IsUpper(r[0]) {
			return nil, fmt.Errorf("%w: field %q is not exported but has tag `%s`", ErrInvalidType, f.Name, DefaultTagName)
		}
		tag := split(s, ',')
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
func (g *Flag) New(ctx *Context) (Value, error) {
	switch g.Type {
	case SliceT:
		return NewSlice(g.Elem), nil
	case MapT:
		return NewMap(g.MapKey, g.Elem)
	case HookT:
		return newHook(ctx, g.Def)
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
func buildFlagOpts(refType reflect.Type, value reflect.Value, name string, s []string) ([]Option, error) {
	var b *bool
	typ, mapKey, elem, err := defaultType(value.Elem().Type())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", value.Type().String(), err)
	}
	opts := []Option{typ, MapKey(mapKey), Elem(elem)}
	for _, opt := range s {
		key, val, _ := strings.Cut(opt, ":")
		switch key {
		case "name":
			opts = append(opts, Name(val))
		case "type":
			opts = append(opts, Type(val))
		case "elem":
			opts = append(opts, Elem(Type(val)))
		case "mapkey":
			opts = append(opts, MapKey(Type(val)))
		case "usage":
			name, usage, _ := strings.Cut(val, "|")
			opts = append(opts, Usage(name, usage))
		case "short":
			opts = append(opts, Short(val))
		case "alias":
			opts = append(opts, Aliases(val))
		case "default":
			opts = append(opts, Default(val))
		case "key":
			typ, key, _ := strings.Cut(val, "|")
			opts = append(opts, Key(typ, key))
		case "section":
			section, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("%w: section %s: %w", ErrInvalidConversion, val, err)
			}
			opts = append(opts, Section(section))
		case "spec":
			opts = append(opts, Spec(val))
		case "hook":
			opts = append(opts, Special("hook:"+val))
		case "set":
			for i := range refType.NumField() {
				field := refType.Field(i)
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

// split splits a string by the cut rune, skipping escaped runes.
func split(str string, cut rune) []string {
	r := []rune(str)
	v := make([][]rune, 1)
	var c, next rune
	for i, j, n := 0, 0, len(r); i < n; i++ {
		c = r[i]
		switch next = peek(r, i+1, n); {
		case r[i] == '\\' && next != 0:
			i, c = i+1, next
		case r[i] == cut:
			j, v = j+1, append(v, nil)
			continue
		}
		v[j] = append(v[j], c)
	}
	s := make([]string, len(v))
	for i, r := range v {
		s[i] = string(r)
	}
	return s
}

// peek peeks at the next rune in r.
func peek(r []rune, i, n int) rune {
	if i < n {
		return r[i]
	}
	return 0
}
