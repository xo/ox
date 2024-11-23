package ox

import (
	"cmp"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// As converts a value.
func As[T any](val any) (T, error) {
	var layout string
	if v, ok := val.(interface{ Layout() string }); ok {
		layout = v.Layout()
	}
	if v, ok := val.(Value); ok {
		val = v.Val()
	}
	if v, err := conv[T](val, layout); err == nil {
		if value, ok := v.(T); ok {
			return value, nil
		}
	}
	var res T
	return res, fmt.Errorf("%w: %T->%T", ErrInvalidConversion, val, res)
}

// SliceAs converts a value to a slice.
func SliceAs[T any](val Value) ([]T, error) {
	/*
		if v, ok := val.(*sliceVal); ok {
			s := make([]T, len(v.v))
			for i := range v.v {
				if v, err := As[T](v.v[i]); err == nil {
					s[i] = v
				}
			}
		}
	*/
	return nil, ErrInvalidConversion
}

// MapAs converts a value to a map.
func MapAs[K cmp.Ordered, T any](val Value) (map[K]T, error) {
	/*
		if v, ok := val.(*mapVal[K]); ok {
			m := make(map[K]T)
			for key := range v.v {
				if k, err := As[K](key); err == nil {
					if v, err := As[T](v.v[key]); err == nil {
						m[k] = v
					}
				}
			}
			return m, nil
		}
	*/
	return nil, ErrInvalidConversion
}

// conv converts a value.
func conv[T any](val any, layout string) (any, error) {
	var res T
	var v any = res
	switch v.(type) {
	case []byte:
		return toBytes(val, layout), nil
	case string:
		return toString(val, layout), nil
	case []rune:
		return []rune(toString(val, layout)), nil
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
	if z, err := convUnmarshal(reflect.TypeOf(val), val, layout); err == nil {
		return z, nil
	}
	if convValue(reflect.ValueOf(&res), val, layout) {
		return res, nil
	}
	return nil, fmt.Errorf("%w: %T->%T", ErrInvalidConversion, val, res)
}

// convValue attempts to convert using reflect -- expects:
// reflect.Value(&<target>).
func convValue(v reflect.Value, val any, layout string) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			ok = false
		}
	}()
	switch v = v.Elem(); v.Kind() {
	case reflect.String:
		v.SetString(toString(val, layout))
		return true
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		if i, ok := asInt64(val); ok && !v.OverflowInt(i) {
			v.SetInt(i)
			return true
		}
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		if u, ok := asUint64(val); ok && !v.OverflowUint(u) {
			v.SetUint(u)
			return true
		}
	case reflect.Float64, reflect.Float32:
		if f, ok := asFloat64(val); ok && !v.OverflowFloat(f) {
			v.SetFloat(f)
			return true
		}
	case reflect.Complex128, reflect.Complex64:
		if c, ok := asComplex128(val); ok && !overflowComplex(v, c) {
			v.SetComplex(c)
			return true
		}
	}
	return false
}

// convUnmarshal creates a new value and unmarshals the value to it.
func convUnmarshal(typ reflect.Type, val any, layout string) (any, error) {
	/*
		d, ok, asText := newDesc{}, false, false
		if d, ok = text[typ]; ok {
			asText = true
		} else if d, ok = binary[typ]; ok {
			asText = false
		}
		if !ok {
			return nil, fmt.Errorf("%w: no text or binary type registered", ErrInvalidConversion)
		}
		v, err := d.New()
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrCouldNotCreateValue, err)
		}
		var unmarshal func([]byte) error
		var b []byte
		if asText {
			z := v.(interface{ UnmarshalText([]byte) error })
			unmarshal, b = z.UnmarshalText, []byte(toString(val, layout))
		} else {
			z := v.(interface{ UnmarshalBinary([]byte) error })
			unmarshal, b = z.UnmarshalBinary, toBytes(val, layout)
		}
		if err := unmarshal(b); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrInvalidConversion, err)

		return v, nil
	*/
	return nil, nil
}

// asBytes converts the value to a []byte.
func asBytes(val any, layout string) ([]byte, bool) {
	switch v := val.(type) {
	case []byte:
		return v, true
	case string:
		return []byte(v), true
	case []rune:
		return []byte(string(v)), true
	case time.Time:
		return []byte(v.Format(layout)), true
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
func asString(val any, layout string) (string, bool) {
	switch v := val.(type) {
	case string:
		return v, true
	case []byte:
		return string(v), true
	case []rune:
		return string(v), true
	case bool:
		return strconv.FormatBool(v), true
	case time.Time:
		return v.Format(layout), true
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
	if s, ok := asString(val, ""); ok {
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			return i, true
		}
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
	if s, ok := asString(val, ""); ok {
		if u, err := strconv.ParseUint(s, 10, 64); err == nil {
			return u, true
		}
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
	if s, ok := asString(val, ""); ok {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f, true
		}
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
	if s, ok := asString(val, ""); ok {
		if c, err := strconv.ParseComplex(s, 128); err == nil {
			return c, true
		}
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
	if v, ok := asString(val, ""); ok {
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
func toBytes(val any, layout string) []byte {
	if v, ok := asBytes(val, layout); ok {
		return v
	}
	if v, ok := asString(val, layout); ok {
		return []byte(v)
	}
	return nil
}

// toString converts a value to string.
func toString(val any, layout string) string {
	if v, ok := asString(val, layout); ok {
		return v
	}
	if v, ok := asBytes(val, layout); ok {
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
	var v T
	return v
}

// toFloat converts the value to a float.
func toFloat[T floati](val any) T {
	if v, ok := asFloat64(val); ok {
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
