package kobra

import (
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
	BigIntT     Type = "big int"
	BigFloatT   Type = "big float"
	BigRatT     Type = "big rat"
	TimestampT  Type = "timestamp"
	DateTimeT   Type = "datetime"
	DateT       Type = "date"
	TimeT       Type = "time"
	DurationT   Type = "duration"
	URLT        Type = "url"
	AddrT       Type = "addr"
	AddrPortT   Type = "addrport"
	CIDRT       Type = "cidr"
	CountT      Type = "count"
	PathT       Type = "path"
	HookT       Type = "hook"
	SliceT      Type = "slice"
	MapT        Type = "map"
	UUIDT       Type = "uuid"
	ColorT      Type = "color"
)

// apply satisfies the [Option] interface.
func (typ Type) apply(val any) error {
	switch v := val.(type) {
	case *Flag:
		if v.Type != SliceT && v.Type != MapT && v.Type != HookT {
			v.Type = typ
		} else {
			v.Sub = typ
		}
	case *marshalDesc:
		v.Type = typ
	case *stringVal:
		v.typ = typ
	case *runeVal:
		v.typ = typ
	case *boolVal:
		v.typ = typ
	case *intVal[int64]:
		v.typ, v.bitSize = typ, 64
	case *intVal[int32]:
		v.typ, v.bitSize = typ, 32
	case *intVal[int16]:
		v.typ, v.bitSize = typ, 16
	case *intVal[int8]:
		v.typ, v.bitSize = typ, 8
	case *intVal[int]:
		v.typ, v.bitSize = typ, 0
	case *uintVal[uint64]:
		v.typ, v.bitSize = typ, 64
	case *uintVal[uint32]:
		v.typ, v.bitSize = typ, 32
	case *uintVal[uint16]:
		v.typ, v.bitSize = typ, 16
	case *uintVal[uint8]:
		v.typ, v.bitSize = typ, 8
	case *uintVal[uint]:
		v.typ, v.bitSize = typ, 0
	case *floatVal[float64]:
		v.typ, v.bitSize = typ, 64
	case *floatVal[float32]:
		v.typ, v.bitSize = typ, 32
	case *complexVal[complex128]:
		v.typ, v.bitSize = typ, 128
	case *complexVal[complex64]:
		v.typ, v.bitSize = typ, 64
	case *timeVal:
		v.typ = typ
	case *durationVal:
		v.typ = typ
	case *marshalVal:
		v.typ = typ
	case *countVal:
		v.typ = typ
	case *pathVal:
		v.typ = typ
	case *mapVal:
		v.typ = typ
	case *sliceVal:
		v.typ = typ
	default:
		return fmt.Errorf("Type(%s) as option: %w", typ, ErrAppliedToInvalidType)
	}
	return nil
}

// New creates a new value for the registered type.
func (typ Type) New() (Value, error) {
	if typ == HookT {
		return nil, nil
	}
	h, ok := types[typ]
	if !ok {
		return nil, ErrTypeNotDefined
	}
	return h.New()
}

// String satisfies the [fmt.Stringer] interface.
func (typ Type) String() string {
	return string(typ)
}

// types are registered type descriptions.
var types map[Type]TypeDesc

// textNew are the new text funcs.
var textNew map[reflect.Type]marshalDesc

// binaryNew are the new binary funcs.
var binaryNew map[reflect.Type]marshalDesc

func init() {
	textNew = make(map[reflect.Type]marshalDesc)
	binaryNew = make(map[reflect.Type]marshalDesc)
	types = map[Type]TypeDesc{
		BytesT:      NewTypeDesc(NewString(BytesT)),
		StringT:     NewTypeDesc(NewString()),
		RunesT:      NewTypeDesc(NewString(RunesT)),
		Base64T:     NewTypeDesc(NewString(Base64T)),
		HexT:        NewTypeDesc(NewString(HexT)),
		BoolT:       NewTypeDesc(NewBool(), NoArg(true)),
		ByteT:       NewTypeDesc(NewRune(ByteT)),
		RuneT:       NewTypeDesc(NewRune()),
		Int64T:      NewTypeDesc(NewInt[int64](Int64T)),
		Int32T:      NewTypeDesc(NewInt[int32](Int32T)),
		Int16T:      NewTypeDesc(NewInt[int16](Int16T)),
		Int8T:       NewTypeDesc(NewInt[int8](Int8T)),
		IntT:        NewTypeDesc(NewInt[int]()),
		Uint64T:     NewTypeDesc(NewUint[uint64](Uint64T)),
		Uint32T:     NewTypeDesc(NewUint[uint32](Uint32T)),
		Uint16T:     NewTypeDesc(NewUint[uint16](Uint16T)),
		Uint8T:      NewTypeDesc(NewUint[uint8](Uint8T)),
		UintT:       NewTypeDesc(NewUint[uint]()),
		Float64T:    NewTypeDesc(NewFloat[float64]()),
		Float32T:    NewTypeDesc(NewFloat[float32](Float32T)),
		Complex128T: NewTypeDesc(NewComplex[complex128]()),
		Complex64T:  NewTypeDesc(NewComplex[complex64](Complex64T)),
		BigIntT:     NewTypeDesc(NewText(func() (any, error) { return big.NewInt(0), nil }, BigIntT)),
		BigFloatT:   NewTypeDesc(NewText(func() (any, error) { return big.NewFloat(0), nil }, BigFloatT)),
		BigRatT:     NewTypeDesc(NewText(func() (any, error) { return big.NewRat(0, 1), nil }, BigRatT)),
		TimestampT:  NewTypeDesc(NewTime()),
		DateTimeT:   NewTypeDesc(NewTime(DateTimeT, Layout(time.DateTime))),
		DateT:       NewTypeDesc(NewTime(DateT, Layout(time.DateOnly))),
		TimeT:       NewTypeDesc(NewTime(TimeT, Layout(time.TimeOnly))),
		DurationT:   NewTypeDesc(NewDuration()),
		URLT:        NewTypeDesc(NewBinary(func() (any, error) { return new(url.URL), nil }, URLT)),
		AddrT:       NewTypeDesc(NewText(func() (any, error) { return new(netip.Addr), nil }, AddrT)),
		AddrPortT:   NewTypeDesc(NewText(func() (any, error) { return new(netip.AddrPort), nil }, AddrPortT)),
		CIDRT:       NewTypeDesc(NewText(func() (any, error) { return new(netip.Prefix), nil }, CIDRT)),
		CountT:      NewTypeDesc(NewCount(), NoArg(true)),
		PathT:       NewTypeDesc(NewPath()),
		HookT:       NewTypeDesc(NewHook(), NoArg(true)),
		SliceT:      NewTypeDesc(NewSlice()),
		MapT:        NewTypeDesc(NewMap()),
	}
}

// RegisterType registers a type description.
func RegisterType(typ Type, h TypeDesc) {
	types[typ] = h
}

// TypeDesc is a type description.
type TypeDesc struct {
	New  func() (Value, error)
	Opts []Option
}

// NewTypeDesc creates a type description.
func NewTypeDesc(f func() (Value, error), opts ...Option) TypeDesc {
	return TypeDesc{
		New:  f,
		Opts: opts,
	}
}

// marshalDesc is a marshaler description.
type marshalDesc struct {
	Type Type
	New  func() (any, error)
}

// registerMarshaler registers a type marshaler.
func registerMarshaler(f func() (any, error), m map[reflect.Type]marshalDesc, opts ...Option) {
	v, err := f()
	if err != nil {
		return
	}
	typ := reflect.TypeOf(v)
	if _, ok := m[typ]; ok {
		return
	}
	d := marshalDesc{
		New: f,
	}
	for _, o := range opts {
		_ = o.apply(&d)
	}
	m[typ] = d
}
