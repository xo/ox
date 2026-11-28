package ox

import (
	"encoding"
	"fmt"
	"math/big"
	"net/netip"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Type is a variable type.
type Type string

// Types.
const (
	BytesT      Type = "bytes"
	StringT     Type = "string"
	RunesT      Type = "runes"
	Base64T     Type = "base64"
	HexT        Type = "hex"
	BoolT       Type = "bool"
	ByteT       Type = "byte"
	RuneT       Type = "rune"
	Int64T      Type = "int64"
	Int32T      Type = "int32"
	Int16T      Type = "int16"
	Int8T       Type = "int8"
	IntT        Type = "int"
	Uint64T     Type = "uint64"
	Uint32T     Type = "uint32"
	Uint16T     Type = "uint16"
	Uint8T      Type = "uint8"
	UintT       Type = "uint"
	Float64T    Type = "float64"
	Float32T    Type = "float32"
	Complex128T Type = "complex128"
	Complex64T  Type = "complex64"
	SizeT       Type = "size"
	RateT       Type = "rate"

	TimestampT Type = "timestamp"
	DateTimeT  Type = "datetime"
	DateT      Type = "date"
	TimeT      Type = "time"

	DurationT Type = "duration"

	CountT Type = "count"
	PathT  Type = "path"
	FileT  Type = "file"
	DirT   Type = "dir"

	BigIntT   Type = "bigint"
	BigFloatT Type = "bigfloat"
	BigRatT   Type = "bigrat"
	AddrT     Type = "addr"
	AddrPortT Type = "addrport"
	CIDRT     Type = "cidr"
	RegexpT   Type = "regexp"
	URLT      Type = "url"
	UUIDT     Type = "uuid"

	SliceT Type = "slice"
	ArrayT Type = "array"
	MapT   Type = "map"

	HookT Type = "hook"
)

// Option returns an [option] for the type.
// The option carries the type, and not a func, so that a flag of each type
// does not allocate a func (B55).
func (typ Type) Option() option {
	return option{
		setType: typ,
	}
}

// String satisfies the [fmt.Stringer] interface.
func (typ Type) String() string {
	return string(typ)
}

// Formatter returns the value format string for the type.
func (typ Type) Formatter() string {
	switch typ {
	case BytesT, StringT, RunesT, ByteT, RuneT:
		return `%q`
	}
	return `%v`
}

// New creates a new [Value] for the registered type.
func (typ Type) New() (Value, error) {
	var v Value
	var err error
	switch typ {
	case HookT:
		return nil, fmt.Errorf("%w (hook)", ErrCouldNotCreateValue)
	case SliceT:
		v = NewSlice(StringT)
	case ArrayT:
		v = NewArray(StringT)
	case MapT:
		v, err = NewMap(StringT, StringT)
	default:
		f, ok := typeLookup(typeNews, typ)
		if !ok {
			return nil, fmt.Errorf("%w: type not registered (%q)", ErrCouldNotCreateValue, string(typ))
		}
		v, err = f()
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCouldNotCreateValue, err)
	}
	return v, nil
}

// applyType applies the type of a [Type] option to val: a [Flag] gets the
// type, or the element type of a slice, an array or a map, and a value with
// a SetType method gets the type.
func (typ Type) applyType(val any) error {
	switch v := val.(type) {
	case *Flag:
		if v.Type != SliceT && v.Type != ArrayT && v.Type != MapT && v.Type != HookT {
			v.Type = typ
		} else if v.Type != HookT {
			v.Elem = typ
		}
		return nil
	case *Command, *Context, *CommandHelp, *Argument:
	case interface{ SetType(Type) }:
		v.SetType(typ)
		return nil
	}
	return fmt.Errorf("Type(%s) (%T): %w", typ, val, ErrAppliedToInvalidType)
}

// typeMu guards the maps of registered types below. The first use of a type
// that reads its own text registers it (D41), so that two goroutines that
// build commands at the same time write to the maps. Each read takes the
// read lock with [typeLookup].
var typeMu sync.RWMutex

// typeLookup returns the value of k in m, one of the maps that typeMu guards.
func typeLookup[K comparable, V any](m map[K]V, k K) (V, bool) {
	typeMu.RLock()
	defer typeMu.RUnlock()
	v, ok := m[k]
	return v, ok
}

// typeNews are registered type creation funcs.
//
// The maps are made here, and not in init, so that the init of any file can
// register a type, such as the init in type_go127.go.
var typeNews = make(map[Type]func() (Value, error))

// typeFlagOpts are registered type flag options.
var typeFlagOpts = make(map[Type][]Option)

// reflectTypes are reflect type name lookups for registered types.
var reflectTypes = make(map[string]Type)

// valueTypes are the types for reflect types that are not pointers and that
// the kind of the reflect type does not give, such as [time.Time].
var valueTypes = map[reflect.Type]Type{
	reflect.TypeFor[time.Time]():     TimestampT,
	reflect.TypeFor[time.Duration](): DurationT,
	reflect.TypeFor[Size]():          SizeT,
	reflect.TypeFor[Rate]():          RateT,
}

// typeTextNews are registered type creation funcs for text marshalable types.
var typeTextNews = make(map[Type]func() (any, error))

// typeBinaryNews are registered type creation funcs for binary marshalable
// types.
var typeBinaryNews = make(map[Type]func() (any, error))

// typeSetNews are the creation funcs of the types that [registerUnmarshal]
// found with a Set method, as a [flag.Value] has.
var typeSetNews = make(map[Type]func() (any, error))

// unmarshalTypes are the types that [registerUnmarshal] registered, by the
// type that is not a pointer.
var unmarshalTypes = make(map[reflect.Type]Type)

func init() {
	// register basic types
	RegisterType(BytesT, NewVal[[]byte]())
	RegisterType(StringT, NewVal[string]())
	RegisterType(RunesT, NewVal[[]rune]())
	RegisterType(Base64T, NewVal[[]byte](Base64T))
	RegisterType(HexT, NewVal[[]byte](HexT))
	RegisterType(BoolT, NewVal[bool](), NoArg(true, true))
	RegisterType(ByteT, NewVal[string](ByteT))
	RegisterType(RuneT, NewVal[string](RuneT))
	RegisterType(Int64T, NewVal[int64]())
	RegisterType(Int32T, NewVal[int32]())
	RegisterType(Int16T, NewVal[int16]())
	RegisterType(Int8T, NewVal[int8]())
	RegisterType(IntT, NewVal[int]())
	RegisterType(Uint64T, NewVal[uint64]())
	RegisterType(Uint32T, NewVal[uint32]())
	RegisterType(Uint16T, NewVal[uint16]())
	RegisterType(Uint8T, NewVal[uint8]())
	RegisterType(UintT, NewVal[uint]())
	RegisterType(Float64T, NewVal[float64]())
	RegisterType(Float32T, NewVal[float32]())
	RegisterType(Complex128T, NewVal[complex128]())
	RegisterType(Complex64T, NewVal[complex64]())
	RegisterType(RateT, NewVal[Rate]())
	RegisterType(SizeT, NewVal[Size]())
	RegisterType(TimestampT, NewTime(TimestampT, ""))
	RegisterType(DateTimeT, NewTime(DateTimeT, time.DateTime))
	RegisterType(DateT, NewTime(DateT, time.DateOnly))
	RegisterType(TimeT, NewTime(TimeT, time.TimeOnly))
	RegisterType(DurationT, NewVal[time.Duration]())
	RegisterType(CountT, NewVal[uint64](CountT), NoArg(true, ""))
	RegisterType(PathT, NewVal[string](PathT))
	RegisterType(FileT, NewVal[string](FileT))
	RegisterType(DirT, NewVal[string](DirT))
	// register text marshal types
	RegisterTextType(func() (*big.Int, error) {
		return big.NewInt(0), nil
	})
	RegisterTextType(func() (*big.Float, error) {
		return big.NewFloat(0), nil
	})
	RegisterTextType(func() (*big.Rat, error) {
		return big.NewRat(0, 1), nil
	})
	RegisterTextType(func() (*netip.Addr, error) {
		return new(netip.Addr), nil
	})
	RegisterTextType(func() (*netip.AddrPort, error) {
		return new(netip.AddrPort), nil
	})
	RegisterTextType(func() (*netip.Prefix, error) {
		return new(netip.Prefix), nil
	})
	RegisterTextType(func() (*regexp.Regexp, error) {
		return new(regexp.Regexp), nil
	})
	// register binary marshal types
	RegisterBinaryType(func() (*url.URL, error) {
		return new(url.URL), nil
	})
}

// RegisterType registers a type.
func RegisterType(typ Type, f func() (Value, error), opts ...Option) {
	typeMu.Lock()
	defer typeMu.Unlock()
	typeNews[typ], typeFlagOpts[typ] = f, opts
}

// RegisterTypeName registers a type name.
func RegisterTypeName(typ Type, names ...string) {
	typeMu.Lock()
	defer typeMu.Unlock()
	for _, name := range names {
		reflectTypes[name] = typ
	}
}

// RegisterTextType registers a new text type.
func RegisterTextType[T TextMarshalUnmarshaler](f func() (T, error)) {
	registerMarshaler[T](func() (any, error) { return f() }, typeTextNews)
}

// RegisterBinaryType registers a new binary type.
func RegisterBinaryType[T BinaryMarshalUnmarshaler](f func() (T, error)) {
	registerMarshaler[T](func() (any, error) { return f() }, typeBinaryNews)
}

// registerMarshaler registers a type marshaler.
func registerMarshaler[T any](f func() (any, error), m map[Type]func() (any, error)) {
	typ := typeType[T]()
	typeMu.Lock()
	defer typeMu.Unlock()
	if _, ok := typeNews[typ]; ok {
		return
	}
	if _, ok := m[typ]; ok {
		return
	}
	typeNews[typ], m[typ] = NewVal[T](typ), f
}

// TextMarshalUnmarshaler is the text marshal interface.
type TextMarshalUnmarshaler interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// BinaryMarshalUnmarshaler is the binary marshal interface.
type BinaryMarshalUnmarshaler interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// typeType returns the type for T.
func typeType[T any]() Type {
	var v T
	return typeRef(v)
}

// typeRef returns the type for val.
func typeRef(val any) Type {
	switch val.(type) {
	case []byte:
		return BytesT
	case string:
		return StringT
	case []rune:
		return RunesT
	case bool:
		return BoolT
	case int64:
		return Int64T
	case int32:
		return Int32T
	case int16:
		return Int16T
	case int8:
		return Int8T
	case int:
		return IntT
	case uint64:
		return Uint64T
	case uint32:
		return Uint32T
	case uint16:
		return Uint16T
	case uint8:
		return Uint8T
	case uint:
		return UintT
	case float64:
		return Float64T
	case float32:
		return Float32T
	case complex128:
		return Complex128T
	case complex64:
		return Complex64T
	case time.Time:
		return TimestampT
	case time.Duration:
		return DurationT
	case Size:
		return SizeT
	case Rate:
		return RateT
	case *big.Int:
		return BigIntT
	case *big.Float:
		return BigFloatT
	case *big.Rat:
		return BigRatT
	case *netip.Addr:
		return AddrT
	case *netip.AddrPort:
		return AddrPortT
	case *netip.Prefix:
		return CIDRT
	case *regexp.Regexp:
		return RegexpT
	case *url.URL:
		return URLT
	}
	typ := reflect.TypeOf(val)
	if typ != nil {
		s := typ.String()
		if typ, ok := typeLookup(reflectTypes, s); ok {
			return typ
		}
		return Type(s)
	}
	return ""
}

// defaultType returns the type, map key type, and element type of v.
func defaultType(refType reflect.Type) (Type, Type, Type, error) {
	if typ, ok := valueTypes[refType]; ok {
		return typ, StringT, StringT, nil
	}
	// a type that reads its own text comes before the kind of the type, so
	// that a named int such as slog.Level reads "debug"
	if typ := registerUnmarshal(refType); typ != "" {
		return typ, StringT, StringT, nil
	}
	switch refType.Kind() {
	case reflect.String:
		return StringT, StringT, StringT, nil
	case reflect.Bool:
		return BoolT, StringT, StringT, nil
	case reflect.Int64:
		return Int64T, StringT, StringT, nil
	case reflect.Int32:
		return Int32T, StringT, StringT, nil
	case reflect.Int16:
		return Int16T, StringT, StringT, nil
	case reflect.Int8:
		return Int8T, StringT, StringT, nil
	case reflect.Int:
		return IntT, StringT, StringT, nil
	case reflect.Uint64:
		return Uint64T, StringT, StringT, nil
	case reflect.Uint32:
		return Uint32T, StringT, StringT, nil
	case reflect.Uint16:
		return Uint16T, StringT, StringT, nil
	case reflect.Uint8:
		return Uint8T, StringT, StringT, nil
	case reflect.Uint:
		return UintT, StringT, StringT, nil
	case reflect.Float64:
		return Float64T, StringT, StringT, nil
	case reflect.Float32:
		return Float32T, StringT, StringT, nil
	case reflect.Complex128:
		return Complex128T, StringT, StringT, nil
	case reflect.Complex64:
		return Complex64T, StringT, StringT, nil
	case reflect.Slice:
		elem := refType.Elem()
		if elem.Kind() == reflect.Slice && elem.Elem().Kind() == reflect.Uint8 {
			return SliceT, StringT, BytesT, nil
		}
		if typ, _, _, err := defaultType(elem); err == nil {
			return SliceT, StringT, typ, nil
		}
	case reflect.Map:
		if mapKey, _, _, err := defaultType(refType.Key()); err == nil {
			if elem, _, _, err := defaultType(refType.Elem()); err == nil {
				return MapT, mapKey, elem, nil
			}
		}
	}
	return "", "", "", ErrInvalidType
}

// unmarshalKind is the method with which a type reads its text.
type unmarshalKind int

// Unmarshal kinds.
const (
	unmarshalNone unmarshalKind = iota
	unmarshalText
	unmarshalBinary
	unmarshalSet
)

// setter is a type that sets itself from a string, as a [flag.Value] does.
type setter interface {
	Set(string) error
}

// unmarshalOf returns the method with which refType, or a pointer to it,
// reads its text. UnmarshalText comes first, then UnmarshalBinary, then Set.
// A type of [valueTypes], such as [time.Time], has none.
// A type assertion, and not [reflect.Type.Implements], finds the method, so
// that it works with TinyGo.
func unmarshalOf(refType reflect.Type) unmarshalKind {
	// time.Time, Size and Rate have their own parse, which reads more forms
	// than their UnmarshalText
	if _, ok := valueTypes[refType]; ok {
		return unmarshalNone
	}
	if refType.Kind() == reflect.Pointer {
		refType = refType.Elem()
	}
	if refType.Kind() == reflect.Pointer || refType.Kind() == reflect.Interface {
		return unmarshalNone
	}
	switch reflect.New(refType).Interface().(type) {
	case encoding.TextUnmarshaler:
		return unmarshalText
	case encoding.BinaryUnmarshaler:
		return unmarshalBinary
	case setter:
		return unmarshalSet
	}
	return unmarshalNone
}

// registerUnmarshal returns the [Type] of refType when refType, or a pointer
// to it, reads its own text, and "" for any other type. A type that is
// registered, such as *url.URL or *colors.Color with the color package,
// keeps its [Type]. A type that is not registered is registered here, as
// [RegisterTextType] or [RegisterBinaryType] would register it. Its name is
// the name of the type in lower case, such as "color" for colors.Color and
// for *colors.Color, unless another type has that name.
//
// A value of the type is always a pointer, and a bind or a conversion to the
// type that is not a pointer copies the value that it points to.
func registerUnmarshal(refType reflect.Type) Type {
	base := refType
	if base.Kind() == reflect.Pointer {
		base = base.Elem()
	}
	ptr := reflect.PointerTo(base)
	if typ := reflectType(ptr); typ != "" {
		return typ
	}
	kind := unmarshalOf(base)
	if kind == unmarshalNone {
		return ""
	}
	typeMu.Lock()
	defer typeMu.Unlock()
	// another goroutine can have registered it
	if typ, ok := unmarshalTypes[base]; ok {
		return typ
	}
	typ := unmarshalName(base)
	f := func() (any, error) {
		return reflect.New(base).Interface(), nil
	}
	switch kind {
	case unmarshalText:
		typeTextNews[typ] = f
	case unmarshalBinary:
		typeBinaryNews[typ] = f
	case unmarshalSet:
		typeSetNews[typ] = f
	}
	typeNews[typ] = func() (Value, error) {
		return &marshalVal{typ: typ}, nil
	}
	unmarshalTypes[base] = typ
	reflectTypes[ptr.String()] = typ
	return typ
}

// unmarshalName returns the name of the [Type] of base: the name of base in
// lower case, such as "color", then its package and name, such as
// "colors.color", then its package path and name. The caller holds the lock
// of typeMu.
func unmarshalName(base reflect.Type) Type {
	for _, name := range []string{
		base.Name(),
		base.String(),
		base.PkgPath() + "." + base.Name(),
	} {
		typ := Type(strings.ToLower(name))
		if _, ok := typeNews[typ]; !ok && name != "" && name != "." {
			return typ
		}
	}
	return Type(base.String())
}

// reflectType returns the [Type] for the reflect type.
func reflectType(refType reflect.Type) Type {
	s := refType.String()
	switch s {
	case "*big.Int":
		return BigIntT
	case "*big.Float":
		return BigFloatT
	case "*big.Rat":
		return BigRatT
	case "*netip.Addr":
		return AddrT
	case "*netip.AddrPort":
		return AddrPortT
	case "*netip.Prefix":
		return CIDRT
	case "*regexp.Regexp":
		return RegexpT
	case "*url.URL":
		return URLT
	}
	if typ, ok := typeLookup(reflectTypes, s); ok {
		return typ
	}
	return ""
}
