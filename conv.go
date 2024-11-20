package kobra

import (
	"encoding"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// As converts a value.
func As[T any](val Value) (T, error) {
	v, err := conv[T](val.Val())
	if err != nil {
		var res T
		return res, err
	}
	z, ok := v.(T)
	if !ok {
		return z, ErrInvalidConversion
	}
	return z, nil
}

// toType returns the corresponding type for a value.
func toType(val any) Type {
	switch v := val.(type) {
	case interface{ Type() Type }:
		return v.Type()
	case []byte:
		return BytesT
	case string:
		return StringT
	case []rune:
		return RunesT
	case time.Time:
		return TimeT
	case time.Duration:
		return DurationT
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
	case interface{ Val() any }:
		return toType(v.Val())
	}
	typ := reflect.TypeOf(val)
	if v, ok := textNew[typ]; ok {
		return v.Type
	}
	if v, ok := binaryNew[typ]; ok {
		return v.Type
	}
	return Type(typ.Kind().String())
}

// conv converts a value.
func conv[T any](val any) (any, error) {
	var res T
	var v any = res
	switch v.(type) {
	case []byte:
		return toBytes(val), nil
	case string:
		return toString(val), nil
	case []rune:
		return []rune(toString(val)), nil
	case time.Time:
		return toTime(val), nil
	case time.Duration:
		return toDuration(val), nil
	case int64:
		return toInt[int64](val), nil
	case int32:
		return toInt[int32](val), nil
	case int16:
		return toInt[int16](val), nil
	case int8:
		return toInt[int8](val), nil
	case int:
		return toInt[int](val), nil
	case uint64:
		return toUint[uint64](val), nil
	case uint32:
		return toUint[uint32](val), nil
	case uint16:
		return toUint[uint16](val), nil
	case uint8:
		return toUint[uint8](val), nil
	case uint:
		return toUint[uint](val), nil
	case float64:
		return toFloat[float64](val), nil
	case float32:
		return toFloat[float32](val), nil
	case complex128:
		return toComplex[complex128](val), nil
	case complex64:
		return toComplex[complex64](val), nil
	}
	// type could be both text and binary unmarshaler, but the new func
	// should only be registered for either text or binary
	if _, ok := v.(interface{ UnmarshalText([]byte) error }); ok {
		if z, err := unmarshalTextValue[T](val); err == nil {
			return z, nil
		}
	}
	if _, ok := v.(interface{ UnmarshalBinary([]byte) error }); ok {
		if z, err := unmarshalBinaryValue[T](val); err == nil {
			return z, nil
		}
	}
	if convValue(reflect.ValueOf(&res), val) {
		return res, nil
	}
	return nil, ErrInvalidConversion
}

// convValue attempts to convert using reflect -- expects:
// reflect.Value(&<target>).
func convValue(v reflect.Value, val any) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			ok = false
		}
	}()
	switch v = v.Elem(); v.Kind() {
	case reflect.String:
		v.SetString(toString(val))
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		if i := toInt[int64](val); !v.OverflowInt(i) {
			v.SetInt(i)
		}
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		if u := toUint[uint64](val); !v.OverflowUint(u) {
			v.SetUint(u)
		}
	case reflect.Float64, reflect.Float32:
		if f := toFloat[float64](val); !v.OverflowFloat(f) {
			v.SetFloat(f)
		}
	case reflect.Complex128, reflect.Complex64:
		if c := toComplex[complex128](val); !overflowComplex(v, c) {
			v.SetComplex(c)
		}
	}
	return true
}

// unmarshalTextValue creates a new value and unmarshals the value to it.
func unmarshalTextValue[T any](val any) (T, error) {
	var res T
	d, ok := textNew[reflect.TypeOf(res)]
	if !ok {
		return res, ErrInvalidTextType
	}
	v, err := d.New()
	if err != nil {
		return res, err
	}
	z, ok := v.(interface{ UnmarshalText([]byte) error })
	if !ok {
		return res, ErrInvalidTextType
	}
	if err := z.UnmarshalText([]byte(toString(val))); err != nil {
		return res, err
	}
	if res, ok = v.(T); !ok {
		return res, ErrInvalidTextType
	}
	return res, nil
}

// unmarshalBinaryValue creates a new value and unmarshals the value to it.
func unmarshalBinaryValue[T any](val any) (T, error) {
	var res T
	d, ok := binaryNew[reflect.TypeOf(res)]
	if !ok {
		return res, ErrInvalidBinaryType
	}
	v, err := d.New()
	if err != nil {
		return res, err
	}
	z, ok := v.(interface{ UnmarshalBinary([]byte) error })
	if !ok {
		return res, ErrInvalidBinaryType
	}
	if err := z.UnmarshalBinary([]byte(toString(val))); err != nil {
		return res, err
	}
	if res, ok = v.(T); !ok {
		return res, ErrInvalidBinaryType
	}
	return res, nil
}

// asBytes converts the value to a []byte.
func asBytes(val any) ([]byte, bool) {
	switch v := val.(type) {
	case []byte:
		return v, true
	case string:
		return []byte(v), true
	case []rune:
		return []byte(string(v)), true
	case interface{ MarshalBinary() ([]byte, error) }:
		if b, err := v.MarshalBinary(); err == nil {
			return b, true
		}
	case interface{ MarshalText() ([]byte, error) }:
		if b, err := v.MarshalText(); err == nil {
			return b, true
		}
	case interface{ Bytes() []byte }:
		return v.Bytes(), true
	case interface{ String() string }:
		return []byte(v.String()), true
	case interface{ Runes() []rune }:
		return []byte(string(v.Runes())), true
	case interface{ Rune() rune }:
		return []byte(string(v.Rune())), true
	case interface{ Byte() byte }:
		return []byte{v.Byte()}, true
	}
	return nil, false
}

// asString converts the value to a string.
func asString(val any) (string, bool) {
	switch v := val.(type) {
	case string:
		return v, true
	case []byte:
		return string(v), true
	case []rune:
		return string(v), true
	case bool:
		return strconv.FormatBool(v), true
	case interface{ MarshalText() ([]byte, error) }:
		if b, err := v.MarshalText(); err == nil {
			return string(b), true
		}
	case interface{ MarshalBinary() ([]byte, error) }:
		if b, err := v.MarshalBinary(); err == nil {
			return string(b), true
		}
	case interface{ String() string }:
		return v.String(), true
	case interface{ Bytes() []byte }:
		return string(v.Bytes()), true
	case interface{ Runes() []rune }:
		return string(v.Runes()), true
	case interface{ Rune() rune }:
		return string(v.Rune()), true
	case interface{ Byte() byte }:
		return string(v.Byte()), true
	}
	return "", false
}

// asInt64 converts the value to a int64.
func asInt64(val any) (int64, bool) {
	switch v := val.(type) {
	case int64:
		return v, true
	case int32:
		return int64(v), true
	case int16:
		return int64(v), true
	case int8:
		return int64(v), true
	case int:
		return int64(v), true
	case interface{ Int64() int64 }:
		return v.Int64(), true
	case interface{ Int32() int32 }:
		return int64(v.Int32()), true
	case interface{ Int16() int16 }:
		return int64(v.Int16()), true
	case interface{ Int8() int8 }:
		return int64(v.Int8()), true
	case interface{ Int() int }:
		return int64(v.Int()), true
	}
	return 0, false
}

// asUint64 converts the value to a uint64.
func asUint64(val any) (uint64, bool) {
	switch v := val.(type) {
	case uint64:
		return v, true
	case uint32:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case uint:
		return uint64(v), true
	case interface{ Uint64() uint64 }:
		return v.Uint64(), true
	case interface{ Uint32() uint32 }:
		return uint64(v.Uint32()), true
	case interface{ Uint16() uint16 }:
		return uint64(v.Uint16()), true
	case interface{ Uint8() uint8 }:
		return uint64(v.Uint8()), true
	case interface{ Uint() uint }:
		return uint64(v.Uint()), true
	}
	return 0, false
}

// asFloat64 converts the value to a float64.
func asFloat64(val any) (float64, bool) {
	switch v := val.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case interface{ Float64() float64 }:
		return v.Float64(), true
	case interface{ Float32() float32 }:
		return float64(v.Float32()), true
	}
	return 0, false
}

// asComplex128 converts the value to a complex128.
func asComplex128(val any) (complex128, bool) {
	switch v := val.(type) {
	case complex128:
		return v, true
	case complex64:
		return complex128(v), true
	case interface{ Complex128() complex128 }:
		return v.Complex128(), true
	case interface{ Complex64() complex64 }:
		return complex128(v.Complex64()), true
	}
	return 0, false
}

// asBool converts the value to a bool.
func asBool(val any) (bool, bool) {
	if v, ok := val.(bool); ok {
		return v, true
	}
	if v, ok := val.(interface{ Bool() bool }); ok {
		return v.Bool(), true
	}
	if v, ok := asInt64(val); ok {
		return 0 < v, true
	}
	if v, ok := asUint64(val); ok {
		return 0 < v, true
	}
	if v, ok := asFloat64(val); ok {
		return 0 < v, true
	}
	if v, ok := asString(val); ok {
		switch strings.TrimSpace(strings.ToLower(v)) {
		case "true", "t", "1":
			return true, true
		case "false", "f", "0", "":
			return false, true
		}
	}
	// NOTE: probably should do more conversions here
	// NOTE: AsComplex128, AsBigInt, AsBigFloat, ...
	return false, false
}

// toBytes converts a value to []byte.
func toBytes(val any) []byte {
	if v, ok := asBytes(val); ok {
		return v
	}
	if v, ok := asString(val); ok {
		return []byte(v)
	}
	return nil
}

// toString converts a value to string.
func toString(val any) string {
	if v, ok := asString(val); ok {
		return v
	}
	if v, ok := asBytes(val); ok {
		return string(v)
	}
	if v, ok := asInt64(val); ok {
		return strconv.FormatInt(v, 10)
	}
	if v, ok := asUint64(val); ok {
		return strconv.FormatUint(v, 10)
	}
	if v, ok := asFloat64(val); ok {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}
	if v, ok := asComplex128(val); ok {
		return strconv.FormatComplex(v, 'f', -1, 128)
	}
	return ""
}

// toBool converts the value to a bool.
func toBool(val any) bool {
	b, _ := asBool(val)
	return b
}

// toBoolString converts the value to either a "true" or "false" string.
func toBoolString(val any) string {
	if val == nil {
		return "true"
	}
	return strconv.FormatBool(toBool(val))
}

// toInt converts the value to a int.
func toInt[T inti](val any) T {
	if v, ok := asInt64(val); ok {
		return T(v)
	}
	if v, ok := asUint64(val); ok {
		return T(v)
	}
	if v, err := strconv.ParseInt(toString(val), 10, 64); err == nil {
		return T(v)
	}
	var v T
	return v
}

// toUint converts the value to a uint.
func toUint[T uinti](val any) T {
	if v, ok := asUint64(val); ok {
		return T(v)
	}
	if v, ok := asInt64(val); ok {
		return T(v)
	}
	if v, err := strconv.ParseUint(toString(val), 10, 64); err == nil {
		return T(v)
	}
	var v T
	return v
}

// toFloat converts the value to a float.
func toFloat[T floati](val any) T {
	if v, ok := asFloat64(val); ok {
		return T(v)
	}
	if v, err := strconv.ParseFloat(toString(val), 64); err == nil {
		return T(v)
	}
	var v T
	return v
}

// toComplex converts the value to a complex.
func toComplex[T complexi](val any) T {
	if v, ok := asComplex128(val); ok {
		return T(v)
	}
	if v, err := strconv.ParseComplex(toString(val), 128); err == nil {
		return T(v)
	}
	var v T
	return v
}

// toTime converts the value to a [time.Time].
func toTime(val any) time.Time {
	switch v := val.(type) {
	case time.Time:
		return v
	case interface{ Time() time.Time }:
		return v.Time()
	}
	return time.Time{}
}

// toDuration converts the value to a [time.Duration].
func toDuration(val any) time.Duration {
	switch v := val.(type) {
	case time.Duration:
		return v
	case interface{ Duration() time.Duration }:
		return v.Duration()
	}
	return time.Duration(0)
}

// inti is the int inteface.
type inti interface {
	~int64 | ~int32 | ~int16 | ~int8 | ~int
}

// uinti is the uint interface.
type uinti interface {
	~uint64 | ~uint32 | ~uint16 | ~uint8 | ~uint
}

// floati is the float interface.
type floati interface {
	~float64 | ~float32
}

// complexi is the complex interface.
type complexi interface {
	~complex128 | ~complex64
}

// texti is the text marshal interface.
type texti interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// binaryi is the binary marshal interface.
type binaryi interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
