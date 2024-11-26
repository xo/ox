package ox

import (
	"cmp"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// As converts a [Value] to type T.
func As[T any](value Value) (T, error) {
	val := value.Val()
	var res T
	v, err := as[T](val, layout(val))
	if err != nil {
		return res, err
	}
	var ok bool
	if res, ok = v.(T); !ok {
		return res, fmt.Errorf("%w: %T->%T", ErrInvalidConversion, v, res)
	}
	return res, nil
}

// AsSlice converts a [Value] to a slice of type E.
func AsSlice[E any](value Value) ([]E, error) {
	val, ok := value.(*sliceVal)
	if !ok {
		return nil, ErrInvalidConversion
	}
	s := make([]E, len(val.v))
	for i := range val.v {
		v, err := As[E](val.v[i])
		if err != nil {
			var res E
			return nil, fmt.Errorf("cannot convert slice value %d: %w: %T->%T", i, err, v, res)
		}
		s[i] = v
	}
	return s, nil
}

// AsMap converts a value to a map with a key of type K, and elements of type
// E.
func AsMap[K cmp.Ordered, E any](value Value) (map[K]E, error) {
	val, ok := value.(*mapVal)
	if !ok {
		return nil, ErrInvalidConversion
	}
	m := make(map[K]E)
	for _, keyval := range val.v.Keys() {
		key, err := as[K](keyval, "")
		if err != nil {
			return nil, err
		}
		k, ok := key.(K)
		if !ok {
			return nil, fmt.Errorf("%w: %T->%T", ErrInvalidConversion, key, k)
		}
		val := val.v.Get(keyval)
		v, err := As[E](val)
		if err != nil {
			return nil, err
		}
		m[k] = v
	}
	return m, nil
}

// To converts a [Value] to type T.
func To[T any](value Value) T {
	v, _ := As[T](value)
	return v
}

// ToSlice converts a [Value] to a slice of type E.
func ToSlice[E any](value Value) []E {
	v, _ := AsSlice[E](value)
	return v
}

// ToMap converts a value to a map with a key of type K, and elements of type
// E.
func ToMap[K cmp.Ordered, E any](value Value) map[K]E {
	v, _ := AsMap[K, E](value)
	return v
}

// as converts a value.
func as[T any](val any, layout string) (any, error) {
	var res T
	var v any = res
	switch v.(type) {
	case []byte:
		return asString[[]byte](val)
	case string:
		return asString[string](val)
	case []rune:
		return asString[[]rune](val)
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
	case FormattedTime:
		return asTime(val, layout)
	case time.Time:
		return asTime(val, layout)
	case time.Duration:
		return asDuration(val)
	}
	// unmarshal
	if v, err := asUnmarshal[T](val); err == nil {
		return v, nil
	}
	// reflect
	if asValue(reflect.ValueOf(&res), val) {
		return res, nil
	}
	return nil, fmt.Errorf("%w: %T->%T", ErrInvalidConversion, val, res)
}

// asString converts the value to a string.
func asString[T stringi](val any) (T, error) {
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
	case FormattedTime:
		return T(v.String()), nil
	case time.Duration:
		if v == 0 {
			return T(""), nil
		}
		return T(v.String()), nil
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
	case interface{ MarshalText() ([]byte, error) }:
		if b, err := v.MarshalText(); err == nil {
			return T(string(b)), nil
		}
	case interface{ MarshalBinary() ([]byte, error) }:
		if b, err := v.MarshalBinary(); err == nil {
			return T(string(b)), nil
		}
	}
	var res T
	return res, ErrInvalidValue
}

// asBool converts the value to a bool.
func asBool(val any) (bool, error) {
	switch v := val.(type) {
	case bool:
		return v, nil
	case interface{ Bool() bool }:
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
	s, err := asString[string](val)
	switch {
	case err != nil:
		return false, err
	case s == "":
		return false, nil
	}
	return strconv.ParseBool(strings.ToLower(s))
}

// asBytes converts to []byte. It is not like asString, which should be used in
// general for forcing conversions.
func asBytes(val any) ([]byte, error) {
	switch v := val.(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	case []rune:
		return []byte(string(v)), nil
	case byte:
		return []byte{v}, nil
	case rune:
		return []byte(string(v)), nil
	}
	return asString[[]byte](val)
}

// asByte converts the value to a single character string.
func asByte(val any) (string, error) {
	b, err := asBytes(val)
	switch n := len(b); {
	case err != nil:
		return "", err
	case n == 0:
		return "", nil
	case n != 1:
		return "", ErrInvalidValue
	}
	return string(b), nil
}

// asRune converts the value to a single rune string.
func asRune(val any) (string, error) {
	b, err := asBytes(val)
	n := len(b)
	switch {
	case err != nil:
		return "", err
	case n == 0:
		return "", nil
	}
	r, size := utf8.DecodeRune(b)
	switch {
	case r == utf8.RuneError, size != n:
		return "", ErrInvalidValue
	}
	return string(r), nil
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
	s, err := asString[string](val)
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
	s, err := asString[string](val)
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
	s, err := asString[string](val)
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
	s, err := asString[string](val)
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
func asTime(val any, layout string) (FormattedTime, error) {
	switch v := val.(type) {
	case FormattedTime:
		return v, nil
	case time.Time:
		return FormattedTime{layout: layout, v: v}, nil
	case interface{ Time() time.Time }:
		return FormattedTime{layout: layout, v: v.Time()}, nil
	}
	s, err := asString[string](val)
	switch {
	case err != nil:
		return FormattedTime{layout: layout}, err
	case s == "":
		return FormattedTime{layout: layout}, nil
	}
	v, err := time.Parse(layout, s)
	if err != nil {
		return FormattedTime{layout: layout}, err
	}
	return FormattedTime{layout: layout, v: v}, nil
}

// asDuration converts the value to a [time.Duration].
func asDuration(val any) (time.Duration, error) {
	switch v := val.(type) {
	case time.Duration:
		return v, nil
	case interface{ Duration() time.Duration }:
		return v.Duration(), nil
	}
	if v, err := asInt[int64](val); err == nil {
		return time.Duration(v) * time.Second, nil
	}
	if v, err := asUint[uint64](val); err == nil {
		return time.Duration(v) * time.Second, nil
	}
	if v, err := asFloat[float64](val); err == nil {
		return time.Duration(v * float64(time.Second)), nil
	}
	s, err := asString[string](val)
	switch {
	case err != nil:
		return 0, err
	case s == "":
		return 0, nil
	}
	return time.ParseDuration(s)
}

// asUnmarshal creates a new value as T, and unmarshals the value to it.
func asUnmarshal[T any](val any) (any, error) {
	buf, err := asString[[]byte](val)
	if err != nil {
		return nil, err
	}
	v, f, err := asNew(typeType[T]())
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		// as with other types, return the "zero value" when val is ""
		return v, nil
	}
	if err := f(buf); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidConversion, err)
	}
	return v, nil
}

// asNew creates a new value for the type, returning the created value and the
// unmarshal func for the value.
func asNew(typ Type) (any, func([]byte) error, error) {
	var (
		f      func() (any, error)
		ok     bool
		asText bool
	)
	if f, ok = typeTextNews[typ]; ok {
		asText = true
	} else if f, ok = typeBinaryNews[typ]; ok {
		asText = false
	}
	if !ok {
		return nil, nil, fmt.Errorf("%w (%s): no text/binary marshaler", ErrInvalidConversion, typ)
	}
	v, err := f()
	if err != nil {
		return nil, nil, fmt.Errorf("%w (%s): %w", ErrCouldNotCreateValue, typ, err)
	}
	var u func([]byte) error
	if asText {
		u = v.(interface{ UnmarshalText([]byte) error }).UnmarshalText
	} else {
		u = v.(interface{ UnmarshalBinary([]byte) error }).UnmarshalBinary
	}
	return v, u, nil
}

// asValue attempts to convert using reflect -- expects:
// reflect.Value(&<target>).
func asValue(v reflect.Value, val any) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			ok = false
		}
	}()
	// fmt.Fprintf(os.Stdout, "asValue: %s\n", v.Type())
	switch v = v.Elem(); v.Kind() {
	case reflect.Slice:
		// TODO: implement []byte/[]rune
	case reflect.String:
		if s, err := asString[string](val); err == nil {
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

// toString converts a value to string.
func toString[T stringi](val any) T {
	v, _ := asString[T](val)
	return v
}

/*
// toBool converts the value to a bool.
func toBool(val any) bool {
	v, _ := asBool(val)
	return v
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
*/

// bitSize returns the bitsize for T.
func bitSize[T inti | uinti | floati | complexi]() int {
	var res T
	var v any = res
	switch v.(type) {
	case complex128:
		return 128
	case int64, uint64, float64, complex64:
		return 64
	case int32, uint32, float32:
		return 32
	case int16, uint16:
		return 16
	case int8, uint8:
		return 8
	}
	return 0
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
