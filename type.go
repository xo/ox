package ox

import (
	"encoding"
	"fmt"
	"math/big"
	"net/netip"
	"net/url"
	"reflect"
	"regexp"
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

	BigIntT   Type = "bigint"
	BigFloatT Type = "bigfloat"
	BigRatT   Type = "bigrat"
	AddrT     Type = "addr"
	AddrPortT Type = "addrport"
	CIDRT     Type = "cidr"
	RegexpT   Type = "regexp"
	URLT      Type = "url"

	UUIDT  Type = "uuid"
	ColorT Type = "color"
	GlobT  Type = "glob"

	SliceT Type = "slice"
	ArrayT Type = "array"
	MapT   Type = "map"

	HookT Type = "hook"
)

// Option returns an [option] for the type.
func (typ Type) Option() option {
	return option{
		name: "Type(" + typ.String() + ")",
		flag: func(g *Flag) error {
			if g.Type != SliceT && g.Type != ArrayT && g.Type != MapT && g.Type != HookT {
				g.Type = typ
			} else if g.Type != HookT {
				g.Elem = typ
			}
			return nil
		},
		typ: func(t interface{ SetType(Type) }) error {
			t.SetType(typ)
			return nil
		},
	}
}

// String satisfies the [fmt.Stringer] interface.
func (typ Type) String() string {
	return string(typ)
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
		f, ok := typeNews[typ]
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

// typeNews are registered type creation funcs.
var typeNews map[Type]func() (Value, error)

// typeFlagOpts are registered type flag options.
var typeFlagOpts map[Type][]Option

// reflectTypes are reflect type name lookups for registered types.
var reflectTypes map[string]Type

// typeTextNews are registered type creation funcs for text marshalable types.
var typeTextNews map[Type]func() (any, error)

// typeBinaryNews are registered type creation funcs for binary marshalable
// types.
var typeBinaryNews map[Type]func() (any, error)

func init() {
	typeNews = make(map[Type]func() (Value, error))
	typeFlagOpts = make(map[Type][]Option)
	reflectTypes = make(map[string]Type)
	typeTextNews = make(map[Type]func() (any, error))
	typeBinaryNews = make(map[Type]func() (any, error))
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
	typeNews[typ], typeFlagOpts[typ] = f, opts
}

// RegisterTypeName registers a type name.
func RegisterTypeName(typ Type, names ...string) {
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
		if typ, ok := reflectTypes[s]; ok {
			return typ
		}
		return Type(s)
	}
	return ""
}

// defaultType returns the type, map key type, and element type of v.
func defaultType(refType reflect.Type) (Type, Type, Type, error) {
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
		return Complex128T, StringT, StringT, nil
	case reflect.Pointer:
		if typ := reflectType(refType); typ != "" {
			return typ, StringT, StringT, nil
		}
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
	if refType == timeType {
		return TimestampT, StringT, StringT, nil
	}
	return "", "", "", ErrInvalidType
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
	if typ, ok := reflectTypes[s]; ok {
		return typ
	}
	return ""
}

var timeType = reflect.TypeOf(time.Time{})
