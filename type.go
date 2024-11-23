package ox

import (
	"encoding"
	"fmt"
	"math/big"
	"net/netip"
	"net/url"
	"os"
	"reflect"
	"strings"
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

	SliceT Type = "slice"
	MapT   Type = "map"

	BigIntT   Type = "big int"
	BigFloatT Type = "big float"
	BigRatT   Type = "big rat"
	URLT      Type = "url"
	AddrT     Type = "addr"
	AddrPortT Type = "addrport"
	CIDRT     Type = "cidr"
	UUIDT     Type = "uuid"
	ColorT    Type = "color"

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

// New creates a new value for the registered type.
func (typ Type) New() (Value, error) {
	if typ == HookT {
		return nil, nil
	}
	f, ok := types[typ]
	if !ok {
		return nil, fmt.Errorf("%w: type not registered", ErrCouldNotCreateValue)
	}
	v, err := f()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCouldNotCreateValue, err)
	}
	return v, nil
}

// Layout returns the time layout for the type.
func (typ Type) Layout() string {
	return layouts[typ]
}

// String satisfies the [fmt.Stringer] interface.
func (typ Type) String() string {
	if i := strings.LastIndex(string(typ), "."); i != -1 {
		return strings.ToLower(string(typ[i+1:]))
	}
	return string(typ)
}

// types are registered type descriptions.
var types map[Type]func() (Value, error)

// layouts are time time layouts.
var layouts map[Type]string

// text holds new text types.
var text map[Type]newDesc

// binary holds new binary types.
var binary map[Type]newDesc

func init() {
	types = map[Type]func() (Value, error){
		BytesT:      NewVal[[]byte](),
		StringT:     NewVal[string](),
		RunesT:      NewVal[[]rune](),
		Base64T:     NewVal[[]byte](Base64T),
		HexT:        NewVal[[]byte](HexT),
		BoolT:       NewVal[bool](),
		ByteT:       NewVal[byte](ByteT),
		RuneT:       NewVal[rune](RuneT),
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
		TimestampT:  NewVal[time.Time](),
		DateTimeT:   NewVal[time.Time](DateTimeT),
		DateT:       NewVal[time.Time](DateT),
		TimeT:       NewVal[time.Time](TimeT),
		DurationT:   NewVal[time.Duration](),
		CountT:      NewVal[uint64](CountT),
		PathT:       NewVal[string](PathT),
		SliceT:      NewSlice(),
		MapT:        NewMap(),
		// HookT:       NewTypeDesc(NewHook(), NoArg(true)),
	}
	layouts = map[Type]string{
		TimestampT: time.RFC3339Nano,
		DateTimeT:  time.DateTime,
		DateT:      time.DateOnly,
		TimeT:      time.TimeOnly,
	}
	text = make(map[Type]newDesc)
	// text marshal types
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
	binary = make(map[Type]newDesc)
	// binary marshal types
	RegisterBinaryType(func() (*url.URL, error) {
		return new(url.URL), nil
	})
}

// newDesc is a new description.
type newDesc struct {
	Type Type
	New  func() (any, error)
}

// RegisterLayout registers a time layout for the type.
func RegisterLayout(typ Type, layout string) {
	layouts[typ] = layout
}

// RegisterTextType registers a new text type.
func RegisterTextType[T TextMarshaler](f func() (T, error), opts ...Option) {
	registerMarshaler[T](func() (any, error) { return f() }, text, opts...)
}

// RegisterBinaryType registers a new binary type.
func RegisterBinaryType[T BinaryMarshaler](f func() (T, error), opts ...Option) {
	registerMarshaler[T](func() (any, error) { return f() }, binary, opts...)
}

// registerMarshaler registers a type marshaler.
func registerMarshaler[T any](f func() (any, error), m map[Type]newDesc, opts ...Option) {
	typ := typeType[T]()
	fmt.Fprintf(os.Stdout, "type: %s\n", typ)
	/*
		if _, ok := m[typ]; ok {
			return
		}
		d := newDesc{
			New: f,
		}
		for _, o := range opts {
			_ = o.apply(&d)
		}
		// set type
		if d.Type == "" {
			s := typ.String()
			if i := strings.LastIndex(s, "."); i != -1 {
				s = s[i+1:]
			}
			d.Type = Type(strings.ToLower(s))
		}
		fmt.Fprintf(os.Stdout, "type: %s -- %s\n", typ, d.Type)
		if _, ok := types[d.Type]; d.Type != "" && !ok {
			types[d.Type] = NewVal[T]()
		}
		m[typ] = d
	*/
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

// varType returns the type for v.
func varType(val any) (Type, bool) {
	switch val.(type) {
	case []byte:
		return BytesT, true
	case string:
		return StringT, true
	case []rune:
		return RunesT, true
	case bool:
		return BoolT, true
	case int64:
		return Int64T, true
	case int32:
		return Int32T, true
	case int16:
		return Int16T, true
	case int8:
		return Int8T, true
	case int:
		return IntT, true
	case uint64:
		return Uint64T, true
	case uint32:
		return Uint32T, true
	case uint16:
		return Uint16T, true
	case uint8:
		return Uint8T, true
	case uint:
		return UintT, true
	case float64:
		return Float64T, true
	case float32:
		return Float32T, true
	case complex128:
		return Complex128T, true
	case complex64:
		return Complex64T, true
	case time.Time:
		return TimestampT, true
	case time.Duration:
		return DurationT, true
	}
	return typeRef(val), false
}

// typeType returns the type for T.
func typeType[T any]() Type {
	var v T
	return typeRef(v)
}

// typeRef returns the type for val.
func typeRef(val any) Type {
	switch val.(type) {
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
	return Type(reflect.TypeOf(val).String())
}
