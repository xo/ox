package ox

import (
	"cmp"
	"encoding"
	"errors"
	"fmt"
	"maps"
	"math"
	"reflect"
	"slices"
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
	switch any(res).(type) {
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
		v, err := asTime(val, layout)
		if err != nil {
			return nil, err
		}
		return v.Time(), nil
	case time.Duration:
		return asDuration(val)
	case Size:
		return asSize(val)
	case Rate:
		return asRate(val)
	}
	// a value of type T, or a pointer to one
	if v, ok := val.(T); ok {
		return v, nil
	}
	refType := reflect.TypeFor[T]()
	if rv := reflect.ValueOf(val); rv.IsValid() && rv.Kind() == reflect.Pointer && !rv.IsNil() && rv.Elem().Type() == refType {
		return rv.Elem().Interface(), nil
	}
	// a registered type
	if typ := typeRef(res); isUnmarshal(typ) {
		v, err := asUnmarshal(typ, val)
		if err != nil {
			return nil, err
		}
		return derefTo(refType, v), nil
	}
	// a type that reads its own text. This registers nothing, so that two
	// goroutines can convert at the same time.
	if kind := unmarshalOf(refType); kind != unmarshalNone {
		v, err := unmarshalTo(refType, kind, val)
		if err != nil {
			return nil, err
		}
		return v.Interface(), nil
	}
	return nil, fmt.Errorf("%w: %T->%s", ErrInvalidConversion, val, refType)
}

// isUnmarshal returns true when typ is a registered type with a text, a
// binary or a Set method.
func isUnmarshal(typ Type) bool {
	_, text := typeLookup(typeTextNews, typ)
	_, binary := typeLookup(typeBinaryNews, typ)
	_, set := typeLookup(typeSetNews, typ)
	return text || binary || set
}

// derefTo returns the value that v points to when refType is not a pointer
// and v is a pointer to refType. Else it returns v.
func derefTo(refType reflect.Type, v any) any {
	if rv := reflect.ValueOf(v); refType.Kind() != reflect.Pointer && rv.Kind() == reflect.Pointer && !rv.IsNil() && rv.Elem().Type() == refType {
		return rv.Elem().Interface()
	}
	return v
}

// unmarshalTo reads val into a new value of refType, with the method of kind.
// When refType is not a pointer, the value is not a pointer.
func unmarshalTo(refType reflect.Type, kind unmarshalKind, val any) (reflect.Value, error) {
	buf, err := asString[[]byte](val)
	if err != nil {
		return reflect.Value{}, err
	}
	base := refType
	if base.Kind() == reflect.Pointer {
		base = base.Elem()
	}
	v := reflect.New(base)
	if len(buf) != 0 {
		if err := unmarshalWith(v.Interface(), kind, buf); err != nil {
			return reflect.Value{}, fmt.Errorf("%w: %w", ErrInvalidConversion, err)
		}
	}
	if refType.Kind() != reflect.Pointer {
		return v.Elem(), nil
	}
	return v, nil
}

// unmarshalWith reads buf into v, with the method of kind.
func unmarshalWith(v any, kind unmarshalKind, buf []byte) error {
	switch kind {
	case unmarshalText:
		return v.(encoding.TextUnmarshaler).UnmarshalText(buf)
	case unmarshalBinary:
		return v.(encoding.BinaryUnmarshaler).UnmarshalBinary(buf)
	case unmarshalSet:
		return v.(setter).Set(string(buf))
	}
	return ErrInvalidConversion
}

// asString converts the value to a string.
func asString[T stringi](val any) (T, error) {
	switch v := val.(type) {
	case nil:
		return T(""), nil
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
	case Size:
		return T(FormatSize(int64(v), 'z', -2, true)), nil
	case Rate:
		return T(FormatRate(v, 'z', -2, true)), nil
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

// marshal marshals val to a string representation.
//
// Does not do recursive marshaling.
func marshal[T stringi](val any) (T, error) {
	switch v, err := asString[T](val); {
	case err == nil:
		return v, nil
	case !errors.Is(err, ErrInvalidValue):
		return v, err
	}
	v := reflect.ValueOf(val)
	// dereference
	switch v.Kind() {
	case reflect.Interface, reflect.Pointer:
		v = v.Elem()
	}
	// marshal
	switch v.Kind() {
	case reflect.Slice:
		return marshalSlice[T](v)
	case reflect.Map:
		return marshalMap[T](v)
	case reflect.Struct:
		return marshalStruct[T](v)
	}
	var z T
	return z, fmt.Errorf("%w: %T", ErrInvalidConversion, z)
}

// marshalSlice marshals a slice into a simple text form.
func marshalSlice[T stringi](v reflect.Value) (T, error) {
	var sb strings.Builder
	for i := range v.Len() {
		if i != 0 {
			_ = sb.WriteByte(',')
		}
		s, err := asString[string](v.Index(i).Interface())
		if err != nil {
			var z T
			return z, fmt.Errorf("%w: index %d: %w of type %T", ErrInvalidConversion, i, err, v.Index(i).Interface())
		}
		_, _ = sb.WriteString(s)
	}
	return asString[T](sb.String())
}

// marshalMap marshals a map into a simple text form.
func marshalMap[T stringi](v reflect.Value) (T, error) {
	m := make(map[string]string)
	for it := v.MapRange(); it.Next(); {
		mk, err := asString[string](it.Key().Interface())
		if err != nil {
			var z T
			return z, fmt.Errorf("%w: map key %q: %w", ErrInvalidConversion, it.Key().Interface(), err)
		}
		mv, err := asString[string](it.Value().Interface())
		if err != nil {
			var z T
			return z, fmt.Errorf("%w: map key %q: %w: %w", ErrInvalidConversion, mk, ErrInvalidValue, err)
		}
		m[mk] = mv
	}
	var sb strings.Builder
	for i, k := range slices.Sorted(maps.Keys(m)) {
		if i != 0 {
			sb.WriteByte(',')
		}
		_, _ = fmt.Fprintf(&sb, "%s=%s", k, m[k])
	}
	return asString[T](sb.String())
}

// marshalStruct marshals a struct into a simple text form.
func marshalStruct[T stringi](v reflect.Value) (T, error) {
	var sb strings.Builder
	for i := range v.NumField() {
		if i != 0 {
			sb.WriteByte(',')
		}
		f := v.Field(i)
		s, err := asString[string](f.Interface())
		if err != nil {
			var z T
			return z, fmt.Errorf("%w struct field %s: %w", ErrInvalidConversion, f.Type().Name(), err)
		}
		_, _ = fmt.Fprintf(&sb, "%s=%s", f.Type().Name(), s)
	}
	return asString[T](sb.String())
}

// asBool converts the value to a bool.
func asBool(val any) (bool, error) {
	switch v := val.(type) {
	case bool:
		return v, nil
	case interface{ Bool() bool }:
		return v.Bool(), nil
	case string:
		// the words of strconv.ParseBool first, so that a word does not try
		// each number parse, which allocates an error. A number such as 2
		// falls through to the number parses
		if v == "" {
			return false, nil
		}
		if b, err := strconv.ParseBool(strings.ToLower(v)); err == nil {
			return b, nil
		}
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

// asTimeAny converts the value to a [time.Time]. A string is parsed with the
// layouts of [TimestampT], [DateTimeT], [DateT] and [TimeT], in turn.
func asTimeAny(val any) (time.Time, error) {
	switch v := val.(type) {
	case time.Time:
		return v, nil
	case FormattedTime:
		return v.Time(), nil
	case interface{ Time() time.Time }:
		return v.Time(), nil
	}
	s, err := asString[string](val)
	switch {
	case err != nil:
		return time.Time{}, err
	case s == "":
		return time.Time{}, nil
	}
	for _, layout := range []string{time.RFC3339Nano, time.DateTime, time.DateOnly, time.TimeOnly} {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("%w: %q is not a time", ErrInvalidConversion, s)
}

// asDuration converts the value to a [time.Duration].
func asDuration(val any) (time.Duration, error) {
	switch v := val.(type) {
	case time.Duration:
		return v, nil
	case interface{ Duration() time.Duration }:
		return v.Duration(), nil
	case string:
		// a duration with a unit first, so that it does not try each number
		// parse, which allocates an error. A number without a unit, such as
		// 30, which time.ParseDuration refuses, falls through to the number
		// parses
		if v == "" {
			return 0, nil
		}
		if d, err := time.ParseDuration(v); err == nil {
			return d, nil
		}
	}
	// a number is seconds, and seconds out of the range of a duration are
	// an error
	if v, err := asInt[int64](val); err == nil {
		if v < math.MinInt64/int64(time.Second) || math.MaxInt64/int64(time.Second) < v {
			return 0, fmt.Errorf("%w: %d seconds overflows a duration", ErrInvalidConversion, v)
		}
		return time.Duration(v) * time.Second, nil
	}
	if v, err := asUint[uint64](val); err == nil {
		if math.MaxInt64/uint64(time.Second) < v {
			return 0, fmt.Errorf("%w: %d seconds overflows a duration", ErrInvalidConversion, v)
		}
		return time.Duration(v) * time.Second, nil
	}
	if v, err := asFloat[float64](val); err == nil {
		d := v * float64(time.Second)
		if math.IsNaN(d) || d < math.MinInt64 || math.MaxInt64 <= d {
			return 0, fmt.Errorf("%w: %v seconds overflows a duration", ErrInvalidConversion, v)
		}
		return time.Duration(d), nil
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

// asSize converts the value to a [Size].
func asSize(val any) (Size, error) {
	switch v := val.(type) {
	case Size:
		return v, nil
	case interface{ Size() Size }:
		return v.Size(), nil
	}
	if v, err := asInt[int64](val); err == nil {
		return Size(v), nil
	}
	if v, err := asUint[uint64](val); err == nil {
		if v > math.MaxInt64 {
			return 0, fmt.Errorf("%w: %d overflows a size", ErrInvalidConversion, v)
		}
		return Size(v), nil
	}
	if v, err := asFloat[float64](val); err == nil {
		// a float out of the range of an int64 has no Size
		if math.IsNaN(v) || v < math.MinInt64 || math.MaxInt64 <= v {
			return 0, fmt.Errorf("%w: %v overflows a size", ErrInvalidConversion, v)
		}
		return Size(v), nil
	}
	s, err := asString[string](val)
	switch {
	case err != nil:
		return 0, err
	case s == "":
		return 0, nil
	}
	return ParseSize(s)
}

// asRate converts the value to a [Rate].
func asRate(val any) (Rate, error) {
	switch v := val.(type) {
	case Rate:
		return v, nil
	case interface{ Rate() Rate }:
		return v.Rate(), nil
	}
	if v, err := asInt[int64](val); err == nil {
		return Rate{Size(v), DefaultRateUnit}, nil
	}
	if v, err := asUint[uint64](val); err == nil {
		if v > math.MaxInt64 {
			return Rate{}, fmt.Errorf("%w: %d overflows a rate", ErrInvalidConversion, v)
		}
		return Rate{Size(v), DefaultRateUnit}, nil
	}
	if v, err := asFloat[float64](val); err == nil {
		return Rate{Size(v), DefaultRateUnit}, nil
	}
	s, err := asString[string](val)
	switch {
	case err != nil:
		return Rate{}, err
	case s == "":
		return Rate{}, nil
	}
	return ParseRate(s)
}

// asUnmarshal creates a new value as T, and unmarshals the value to it.
func asUnmarshal(typ Type, val any) (any, error) {
	buf, err := asString[[]byte](val)
	if err != nil {
		return nil, err
	}
	v, f, err := asNew(typ)
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
	var f func() (any, error)
	var kind unmarshalKind
	for _, k := range []struct {
		news map[Type]func() (any, error)
		kind unmarshalKind
	}{
		{typeTextNews, unmarshalText},
		{typeBinaryNews, unmarshalBinary},
		{typeSetNews, unmarshalSet},
	} {
		if g, ok := typeLookup(k.news, typ); ok {
			f, kind = g, k.kind
			break
		}
	}
	if f == nil {
		return nil, nil, fmt.Errorf("%w (%s): no text/binary marshaler", ErrInvalidConversion, typ)
	}
	v, err := f()
	if err != nil {
		return nil, nil, fmt.Errorf("%w (%s): %w", ErrCouldNotCreateValue, typ, err)
	}
	return v, func(buf []byte) error {
		return unmarshalWith(v, kind, buf)
	}, nil
}

// asValue converts a value -- expects reflect.Value(&<target>).
func asValue(value reflect.Value, val any) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	// the types whose kind does not give the conversion
	switch el := value.Elem(); el.Type() {
	case reflect.TypeFor[time.Duration]():
		d, err := asDuration(val)
		if err != nil {
			return err
		}
		el.SetInt(int64(d))
		return nil
	case reflect.TypeFor[Size]():
		size, err := asSize(val)
		if err != nil {
			return err
		}
		el.SetInt(int64(size))
		return nil
	case reflect.TypeFor[Rate]():
		rate, err := asRate(val)
		if err != nil {
			return err
		}
		el.Set(reflect.ValueOf(rate))
		return nil
	case reflect.TypeFor[time.Time]():
		t, err := asTimeAny(val)
		if err != nil {
			return err
		}
		el.Set(reflect.ValueOf(t))
		return nil
	}
	if ok, err := unmarshalInto(value, val); ok {
		return err
	}
	switch el := value.Elem(); el.Kind() {
	case reflect.Slice:
		// reflect cannot tell []rune from []int32, so a []int32 takes the
		// runes of the value
		var v any
		switch el.Type().Elem().Kind() {
		case reflect.Uint8:
			v, err = asString[[]byte](val)
		case reflect.Int32:
			v, err = asString[[]rune](val)
		default:
			return ErrInvalidConversion
		}
		if err != nil {
			return err
		}
		el.Set(reflect.ValueOf(v).Convert(el.Type()))
		return nil
	case reflect.String:
		s, err := asString[string](val)
		if err != nil {
			return err
		}
		el.SetString(s)
		return nil
	case reflect.Bool:
		b, err := asBool(val)
		if err != nil {
			return err
		}
		el.SetBool(b)
		return nil
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		switch i, err := asInt[int64](val); {
		case err != nil:
			return err
		case el.OverflowInt(i):
		default:
			el.SetInt(i)
			return nil
		}
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		switch u, err := asUint[uint64](val); {
		case err != nil:
			return err
		case el.OverflowUint(u):
		default:
			el.SetUint(u)
			return nil
		}
	case reflect.Float64, reflect.Float32:
		switch f, err := asFloat[float64](val); {
		case err != nil:
			return err
		case el.OverflowFloat(f):
		default:
			el.SetFloat(f)
			return nil
		}
	case reflect.Complex128, reflect.Complex64:
		switch c, err := asComplex[complex128](val); {
		case err != nil:
			return err
		case overflowComplex(el, c):
		default:
			el.SetComplex(c)
			return nil
		}
	case reflect.Pointer:
		v, err := asUnmarshal(reflectType(el.Type()), val)
		if err != nil {
			return err
		}
		reflect.Indirect(value).Set(reflect.ValueOf(v))
		return nil
	}
	return ErrInvalidConversion
}

// unmarshalInto sets the value that value points to from val, when the type
// of the value, or a pointer to it, reads its own text. It returns false for
// any other type. When val is already of the type, or a pointer to it, it
// sets val, or a copy of what val points to.
func unmarshalInto(value reflect.Value, val any) (bool, error) {
	el := value.Elem()
	kind := unmarshalOf(el.Type())
	if kind == unmarshalNone {
		return false, nil
	}
	if rv := reflect.ValueOf(val); rv.IsValid() {
		switch {
		case rv.Type().AssignableTo(el.Type()):
			el.Set(rv)
			return true, nil
		case rv.Kind() == reflect.Pointer && !rv.IsNil() && rv.Elem().Type().AssignableTo(el.Type()):
			el.Set(rv.Elem())
			return true, nil
		}
	}
	v, err := unmarshalTo(el.Type(), kind, val)
	if err != nil {
		return true, err
	}
	el.Set(v)
	return true, nil
}

// toString converts a value to string.
func toString[T stringi](val any) T {
	v, _ := asString[T](val)
	return v
}

// bitSize returns the bitsize for T.
func bitSize[T inti | uinti | floati | complexi]() int {
	switch any(T(0)).(type) {
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
