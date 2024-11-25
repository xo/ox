package ox

import (
	"encoding"
	"fmt"
	"math/big"
	"net/netip"
	"net/url"
	"reflect"
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
	TimestampT  Type = "timestamp"
	DateTimeT   Type = "datetime"
	DateT       Type = "date"
	TimeT       Type = "time"
	DurationT   Type = "duration"
	CountT      Type = "count"
	PathT       Type = "path"

	BigIntT   Type = "big int"
	BigFloatT Type = "big float"
	BigRatT   Type = "big rat"
	URLT      Type = "url"
	AddrT     Type = "addr"
	AddrPortT Type = "addrport"
	CIDRT     Type = "cidr"

	UUIDT  Type = "uuid"
	ColorT Type = "color"

	SliceT Type = "slice"
	MapT   Type = "map"

	HookT Type = "hook"
)

// apply satisfies the [Option] interface.
func (typ Type) apply(val any) error {
	switch v := val.(type) {
	case *Flag:
		if v.Type != SliceT && v.Type != MapT && v.Type != HookT {
			v.Type = typ
		} else if v.Type != HookT {
			v.Sub = typ
		}
	case interface{ SetType(Type) }:
		v.SetType(typ)
	default:
		return fmt.Errorf("Type(%s) used as option: %w", typ, ErrAppliedToInvalidType)
	}
	return nil
}

// String satisfies the [fmt.Stringer] interface.
func (typ Type) String() string {
	return string(typ)
}

// New creates a new [Value] for the registered type.
func (typ Type) New() (Value, error) {
	if typ == HookT {
		return nil, nil
	}
	f, ok := typeNews[typ]
	if !ok {
		return nil, fmt.Errorf("%w: type not registered", ErrCouldNotCreateValue)
	}
	v, err := f()
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
	typeNews = map[Type]func() (Value, error){
		BytesT:      NewVal[[]byte](),
		StringT:     NewVal[string](),
		RunesT:      NewVal[[]rune](),
		Base64T:     NewVal[[]byte](Base64T),
		HexT:        NewVal[[]byte](HexT),
		BoolT:       NewVal[bool](),
		ByteT:       NewVal[string](ByteT),
		RuneT:       NewVal[string](RuneT),
		Int64T:      NewVal[int64](),
		Int32T:      NewVal[int32](),
		Int16T:      NewVal[int16](),
		Int8T:       NewVal[int8](),
		IntT:        NewVal[int](),
		Uint64T:     NewVal[uint64](),
		Uint32T:     NewVal[uint32](),
		Uint16T:     NewVal[uint16](),
		Uint8T:      NewVal[uint8](),
		UintT:       NewVal[uint](),
		Float64T:    NewVal[float64](),
		Float32T:    NewVal[float32](),
		Complex128T: NewVal[complex128](),
		Complex64T:  NewVal[complex64](),
		TimestampT:  NewTime(TimestampT, ""),
		DateTimeT:   NewTime(DateTimeT, time.DateTime),
		DateT:       NewTime(DateT, time.DateOnly),
		TimeT:       NewTime(TimeT, time.TimeOnly),
		DurationT:   NewVal[time.Duration](),
		CountT:      NewVal[uint64](CountT),
		PathT:       NewVal[string](PathT),
	}
	typeFlagOpts = map[Type][]Option{
		BoolT:  {NoArg(true, true)},
		CountT: {NoArg(true, "")},
	}
	reflectTypes = make(map[string]Type)
	// text marshal types
	typeTextNews = make(map[Type]func() (any, error))
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
	// binary marshal types
	typeBinaryNews = make(map[Type]func() (any, error))
	RegisterBinaryType(func() (*url.URL, error) {
		return new(url.URL), nil
	})
}

// RegisterTypeOptions registers [Flag] options used within [NewFlag], used for
// setting default flag values.
func RegisterTypeOptions(typ Type, opts ...Option) {
	typeFlagOpts[typ] = opts
}

// RegisterTypeName registers a type name.
func RegisterTypeName(typ Type, names ...string) {
	for _, name := range names {
		reflectTypes[name] = typ
	}
}

// RegisterTextType registers a new text type.
func RegisterTextType[T TextMarshaler](f func() (T, error)) {
	registerMarshaler[T](func() (any, error) { return f() }, typeTextNews)
}

// RegisterBinaryType registers a new binary type.
func RegisterBinaryType[T BinaryMarshaler](f func() (T, error)) {
	registerMarshaler[T](func() (any, error) { return f() }, typeBinaryNews)
}

// registerMarshaler registers a type marshaler.
func registerMarshaler[T any](f func() (any, error), descs map[Type]func() (any, error)) {
	typ := typeType[T]()
	if _, ok := typeNews[typ]; ok {
		return
	}
	if _, ok := descs[typ]; ok {
		return
	}
	typeNews[typ], descs[typ] = NewVal[T](typ), f
}

// TextMarshaler is the text marshal interface.
type TextMarshaler interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// BinaryMarshaler is the binary marshal interface.
type BinaryMarshaler interface {
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
	case *url.URL:
		return URLT
	}
	s := reflect.TypeOf(val).String()
	if typ, ok := reflectTypes[s]; ok {
		return typ
	}
	return Type(s)
}
