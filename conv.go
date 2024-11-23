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
	var res T
	v, err := as[T](val, layout)
	if err != nil {
		return res, err
	}
	var ok bool
	if res, ok = v.(T); !ok {
		return res, fmt.Errorf("%w: %T->%T", ErrInvalidConversion, v, res)
	}
	return res, nil
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

// as converts a value.
func as[T any](val any, layout string) (any, error) {
	var res T
	var v any = res
	switch v.(type) {
	case []byte:
		return asString[[]byte](val, layout)
	case string:
		return asString[string](val, layout)
	case []rune:
		return asString[[]rune](val, layout)
	case bool:
		return asBool(val)
	case int64:
		return asInt[int64](val)
	case int32:
		return asInt[int32](val)
	case int16:
		return asInt[int16](val)
	case int8:
		return asInt[int8](val)
	case int:
		return asInt[int](val)
	case uint64:
		return asUint[uint64](val)
	case uint32:
		return asUint[uint32](val)
	case uint16:
		return asUint[uint16](val)
	case uint8:
		return asUint[uint8](val)
	case uint:
		return asUint[uint](val)
	case float64:
		return asFloat[float64](val)
	case float32:
		return asFloat[float32](val)
	case complex128:
		return asComplex[complex128](val)
	case complex64:
		return asComplex[complex64](val)
	case time.Time:
		return asTime(val, layout)
	case time.Duration:
		return asDuration(val)
	}
	/*
		// type could be both text and binary unmarshaler, but the new func
		// should only be registered for either text or binary
		if z, err := convUnmarshal(reflect.TypeOf(val), val, layout); err == nil {
			return z, nil
		}
		if convValue(reflect.ValueOf(&res), val, layout) {
			return res, nil
		}
	*/
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
	case reflect.Slice:
		// []byte/[]rune
	case reflect.String:
		if s, err := asString[string](val, layout); err == nil {
			v.SetString(s)
			return true
		}
	case reflect.Bool:
		if b, err := asBool(val); err == nil {
			v.SetBool(b)
			return true
		}
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		if i, err := asInt[int64](val); err == nil && !v.OverflowInt(i) {
			v.SetInt(i)
			return true
		}
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		if u, err := asUint[uint64](val); err == nil && !v.OverflowUint(u) {
			v.SetUint(u)
			return true
		}
	case reflect.Float64, reflect.Float32:
		if f, err := asFloat[float64](val); err == nil && !v.OverflowFloat(f) {
			v.SetFloat(f)
			return true
		}
	case reflect.Complex128, reflect.Complex64:
		if c, err := asComplex[complex128](val); err == nil && !overflowComplex(v, c) {
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

// asString converts the value to a string.
func asString[T stringi](val any, layout string) (T, error) {
	switch v := val.(type) {
	case string:
		return T(v), nil
	case []byte:
		return T(string(v)), nil
	case []rune:
		return T(string(v)), nil
	case bool:
		return T(strconv.FormatBool(v)), nil
	case int64:
		return T(strconv.FormatInt(v, 10)), nil
	case int32:
		return T(strconv.FormatInt(int64(v), 10)), nil
	case int16:
		return T(strconv.FormatInt(int64(v), 10)), nil
	case int8:
		return T(strconv.FormatInt(int64(v), 10)), nil
	case int:
		return T(strconv.FormatInt(int64(v), 10)), nil
	case uint64:
		return T(strconv.FormatUint(v, 10)), nil
	case uint32:
		return T(strconv.FormatUint(uint64(v), 10)), nil
	case uint16:
		return T(strconv.FormatUint(uint64(v), 10)), nil
	case uint8:
		return T(strconv.FormatUint(uint64(v), 10)), nil
	case uint:
		return T(strconv.FormatUint(uint64(v), 10)), nil
	case float64:
		return T(strconv.FormatFloat(v, 'f', -1, bitSize[float64]())), nil
	case float32:
		return T(strconv.FormatFloat(float64(v), 'f', -1, bitSize[float32]())), nil
	case complex128:
		return T(strconv.FormatComplex(v, 'f', -1, bitSize[complex128]())), nil
	case complex64:
		return T(strconv.FormatComplex(complex128(v), 'f', -1, bitSize[complex64]())), nil
	case time.Time:
		return T(v.Format(layout)), nil
	case time.Duration:
		return T(v.String()), nil
	case interface{ MarshalText() ([]byte, error) }:
		if b, err := v.MarshalText(); err == nil {
			return T(string(b)), nil
		}
	case interface{ MarshalBinary() ([]byte, error) }:
		if b, err := v.MarshalBinary(); err == nil {
			return T(string(b)), nil
		}
	case interface{ String() string }:
		return T(v.String()), nil
	case interface{ Bytes() []byte }:
		return T(string(v.Bytes())), nil
	case interface{ Runes() []rune }:
		return T(string(v.Runes())), nil
	case interface{ Rune() rune }:
		return T(string(v.Rune())), nil
	case interface{ Byte() byte }:
		return T(string(v.Byte())), nil
	}
	var res T
	return res, ErrInvalidConversion
}

// asBool converts the value to a bool.
func asBool(val any) (bool, error) {
	if v, ok := val.(bool); ok {
		return v, nil
	}
	if v, ok := val.(interface{ Bool() bool }); ok {
		return v.Bool(), nil
	}
	if v, err := asInt[int64](val); err == nil {
		return 0 < v, nil
	}
	if v, err := asUint[uint64](val); err == nil {
		return 0 < v, nil
	}
	if v, err := asFloat[float64](val); err == nil {
		return 0 < v, nil
	}
	// NOTE: not doing asComplex here
	s, err := asString[string](val, "")
	switch {
	case err != nil:
		return false, err
	case s == "":
		return false, nil
	}
	return strconv.ParseBool(strings.ToLower(s))
}

// asInt converts the value to a int64.
func asInt[T inti](val any) (T, error) {
	switch v := val.(type) {
	case int64:
		return T(v), nil
	case int32:
		return T(v), nil
	case int16:
		return T(v), nil
	case int8:
		return T(v), nil
	case int:
		return T(v), nil
	case interface{ Int64() int64 }:
		return T(v.Int64()), nil
	case interface{ Int32() int32 }:
		return T(v.Int32()), nil
	case interface{ Int16() int16 }:
		return T(v.Int16()), nil
	case interface{ Int8() int8 }:
		return T(v.Int8()), nil
	case interface{ Int() int }:
		return T(v.Int()), nil
	}
	s, err := asString[string](val, "")
	switch {
	case err != nil:
		return 0, err
	case s == "":
		return 0, nil
	}
	v, err := strconv.ParseInt(s, 10, bitSize[T]())
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

// asUint converts the value to a uint64.
func asUint[T uinti](val any) (T, error) {
	switch v := val.(type) {
	case uint64:
		return T(v), nil
	case uint32:
		return T(v), nil
	case uint16:
		return T(v), nil
	case uint8:
		return T(v), nil
	case uint:
		return T(v), nil
	case interface{ Uint64() uint64 }:
		return T(v.Uint64()), nil
	case interface{ Uint32() uint32 }:
		return T(v.Uint32()), nil
	case interface{ Uint16() uint16 }:
		return T(v.Uint16()), nil
	case interface{ Uint8() uint8 }:
		return T(v.Uint8()), nil
	case interface{ Uint() uint }:
		return T(v.Uint()), nil
	}
	s, err := asString[string](val, "")
	switch {
	case err != nil:
		return 0, err
	case s == "":
		return 0, nil
	}
	v, err := strconv.ParseUint(s, 10, bitSize[T]())
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

// asFloat converts the value to a float64.
func asFloat[T floati](val any) (T, error) {
	switch v := val.(type) {
	case float64:
		return T(v), nil
	case float32:
		return T(v), nil
	case interface{ Float64() float64 }:
		return T(v.Float64()), nil
	case interface{ Float32() float32 }:
		return T(v.Float32()), nil
	}
	s, err := asString[string](val, "")
	switch {
	case err != nil:
		return 0, err
	case s == "":
		return 0, nil
	}
	v, err := strconv.ParseFloat(s, bitSize[T]())
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

// asComplex converts the value to a complex128.
func asComplex[T complexi](val any) (T, error) {
	switch v := val.(type) {
	case complex128:
		return T(v), nil
	case complex64:
		return T(v), nil
	case interface{ Complex128() complex128 }:
		return T(v.Complex128()), nil
	case interface{ Complex64() complex64 }:
		return T(v.Complex64()), nil
	}
	s, err := asString[string](val, "")
	switch {
	case err != nil:
		return 0, err
	case s == "":
		return 0, nil
	}
	v, err := strconv.ParseComplex(s, bitSize[T]())
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

// asTime converts the value to a [time.Time].
func asTime(val any, layout string) (time.Time, error) {
	switch v := val.(type) {
	case time.Time:
		return v, nil
	case interface{ Time() time.Time }:
		return v.Time(), nil
	}
	s, err := asString[string](val, layout)
	switch {
	case err != nil:
		return time.Time{}, err
	case s == "":
		return time.Time{}, nil
	}
	return time.Parse(layout, s)
}

// asDuration converts the value to a [time.Duration].
func asDuration(val any) (time.Duration, error) {
	switch v := val.(type) {
	case time.Duration:
		return v, nil
	case interface{ Duration() time.Duration }:
		return v.Duration(), nil
	}
	s, err := asString[string](val, "")
	switch {
	case err != nil:
		return 0, err
	case s == "":
		return 0, nil
	}
	return time.ParseDuration(s)
}

// toBytes converts a value to []byte.
func toBytes(val any, layout string) []byte {
	v, _ := asString[[]byte](val, layout)
	return v
}

// toString converts a value to string.
func toString(val any, layout string) string {
	v, _ := asString[string](val, layout)
	return v
}

// toRunes converts a value to a []rune.
func toRunes(val any, layout string) []rune {
	v, _ := asString[[]rune](val, layout)
	return v
}

// toBool converts the value to a bool.
func toBool(val any) bool {
	v, _ := asBool(val)
	return v
}

// toBoolString converts the value to a string, returning "true" when the value
// is nil.
func toBoolString(val any) string {
	if val != nil {
		return strconv.FormatBool(toBool(val))
	}
	return "true"
}

// toInt converts the value to a int.
func toInt[T inti](val any) T {
	v, _ := asInt[T](val)
	return v
}

// toUint converts the value to a uint.
func toUint[T uinti](val any) T {
	v, _ := asUint[T](val)
	return v
}

// toFloat converts the value to a float.
func toFloat[T floati](val any) T {
	v, _ := asFloat[T](val)
	return v
}

// toComplex converts the value to a complex.
func toComplex[T complexi](val any) T {
	v, _ := asComplex[T](val)
	return v
}

// toTime converts the value to a [time.Time].
func toTime(val any, layout string) time.Time {
	v, _ := asTime(val, layout)
	return v
}

// toDuration converts the value to a [time.Duration].
func toDuration(val any) time.Duration {
	v, _ := asDuration(val)
	return v
}

// stringi is the string interface.
type stringi interface {
	~[]byte | ~string | ~[]rune
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
