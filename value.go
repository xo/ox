package ox

import (
	"cmp"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"
	"sync/atomic"
	"time"
)

// Value is the value interface.
type Value interface {
	Type() Type
	Val() any
	SetSet(bool)
	WasSet() bool
	Set(string) error
	Get() (string, error)
	String() string
}

// anyVal wraps a value.
type anyVal[T any] struct {
	typ Type
	set bool
	v   T
}

// NewVal returns a [Value] func storing the value as type T.
func NewVal[T any](opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &anyVal[T]{
			typ: typeType[T](),
		}
		if err := Apply(val, opts...); err != nil {
			return nil, err
		}
		return val, nil
	}
}

func (val *anyVal[T]) Type() Type {
	return val.typ
}

func (val *anyVal[T]) Val() any {
	return val.v
}

func (val *anyVal[T]) SetType(typ Type) {
	val.typ = typ
}

func (val *anyVal[T]) SetSet(set bool) {
	val.set = set
}

func (val *anyVal[T]) WasSet() bool {
	return val.set
}

func (val *anyVal[T]) Set(s string) error {
	var value any = s
	switch val.typ {
	case Base64T:
		b, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidValue, err)
		}
		value = b
	case HexT:
		b, err := hex.DecodeString(s)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidValue, err)
		}
		value = b
	case ByteT:
		s, err := asByte(s)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidValue, err)
		}
		value = s
	case RuneT:
		s, err := asRune(s)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidValue, err)
		}
		value = s
	case CountT:
		inc(&val.v, 1)
		return nil
	}
	v, err := as[T](value, layout(val.v))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidValue, err)
	}
	var ok bool
	if val.v, ok = v.(T); !ok {
		return fmt.Errorf("%w: %T->%T", ErrInvalidConversion, value, val.v)
	}
	return nil
}

func (val *anyVal[T]) Get() (string, error) {
	if invalid(val.v) {
		return "", nil
	}
	v, err := as[string](val.v, layout(val.v))
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInvalidValue, err)
	}
	if s, ok := v.(string); ok {
		return s, nil
	}
	return "", fmt.Errorf("%w: %T->string", ErrInvalidConversion, v)
}

func (val *anyVal[T]) String() string {
	s, _ := val.Get()
	return s
}

// NewTime returns a Value func for a [FormattedTime] value.
func NewTime(typ Type, layout string) func() (Value, error) {
	return func() (Value, error) {
		val := &anyVal[FormattedTime]{
			typ: typ,
			v:   FormattedTime{layout: layout},
		}
		if val.typ == "" {
			val.typ = TimestampT
		}
		if val.v.layout == "" {
			val.v.layout = DefaultLayout
		}
		return val, nil
	}
}

// sliceVal is a slice value.
type sliceVal struct {
	typ   Type
	v     []Value
	set   bool
	split func(string) []string
}

// NewSlice creates a slice value of type.
func NewSlice(typ Type) Value {
	return &sliceVal{
		typ: typ,
		split: func(s string) []string {
			return SplitBy(s, ',')
		},
	}
}

// NewArray creates a slice value of type.
func NewArray(typ Type) Value {
	return &sliceVal{
		typ: typ,
	}
}

func (val *sliceVal) Type() Type {
	return "[]" + val.typ
}

func (val *sliceVal) Val() any {
	return val.v
}

func (val *sliceVal) Get() (string, error) {
	return val.String(), nil
}

func (val *sliceVal) SetSet(set bool) {
	val.set = set
}

func (val *sliceVal) WasSet() bool {
	return val.set
}

func (val *sliceVal) Set(s string) error {
	var strs []string
	if val.split != nil {
		strs = val.split(s)
	} else {
		strs = []string{s}
	}
	for i, str := range strs {
		v, err := val.typ.New()
		if err != nil {
			return err
		}
		switch err := v.Set(str); {
		case err != nil && 1 < len(strs):
			return fmt.Errorf("set %d: %w", i, err)
		case err != nil:
			return err
		}
		val.v = append(val.v, v)
	}
	return nil
}

func (val *sliceVal) String() string {
	s := make([]string, len(val.v))
	for i, v := range val.v {
		s[i] = toString[string](v)
	}
	return "[" + strings.Join(s, " ") + "]"
}

// SetSplit sets the split func.
func (val *sliceVal) SetSplit(split func(string) []string) {
	val.split = split
}

// Index returns the i'th variable from the slice.
func (val *sliceVal) Index(i int) Value {
	return val.v[i]
}

// Len returns the slice length.
func (val *sliceVal) Len() int {
	return len(val.v)
}

// mapVal is a map value.
type mapVal struct {
	key   Type
	typ   Type
	set   bool
	v     valueMap
	split func(string) []string
}

// NewMap creates a map value of type.
func NewMap(key, typ Type) (Value, error) {
	val := &mapVal{
		key: key,
		typ: typ,
		split: func(s string) []string {
			return SplitBy(s, ',')
		},
	}
	var err error
	if val.v, err = makeMap(val.key, val.typ); err != nil {
		return nil, err
	}
	return val, nil
}

func (val *mapVal) Type() Type {
	return "map[" + val.key + "]" + val.typ
}

func (val *mapVal) Val() any {
	return val.v
}

func (val *mapVal) SetSet(set bool) {
	val.set = set
}

func (val *mapVal) WasSet() bool {
	return val.set
}

func (val *mapVal) Set(s string) error {
	var strs []string
	if val.split != nil {
		strs = val.split(s)
	} else {
		strs = []string{s}
	}
	for i, str := range strs {
		switch err := val.v.Set(str); {
		case err != nil && 1 < len(strs):
			return fmt.Errorf("set %d: %w", i, err)
		case err != nil:
			return err
		}
	}
	return nil
}

func (val *mapVal) String() string {
	return val.v.String()
}

func (val *mapVal) Get() (string, error) {
	return val.v.String(), nil
}

// SetSplit sets the split func.
func (val *mapVal) SetSplit(split func(string) []string) {
	val.split = split
}

type valueMap interface {
	Set(s string) error
	String() string
	Keys() []any
	Get(any) Value
}

// makeMap makes a map for the key and type.
func makeMap(key, typ Type) (valueMap, error) {
	switch key {
	case StringT:
		return newValueMap[string](typ), nil
	case Int64T:
		return newValueMap[int64](typ), nil
	case Int32T:
		return newValueMap[int32](typ), nil
	case Int16T:
		return newValueMap[int16](typ), nil
	case Int8T:
		return newValueMap[int8](typ), nil
	case IntT:
		return newValueMap[int](typ), nil
	case Uint64T:
		return newValueMap[uint64](typ), nil
	case Uint32T:
		return newValueMap[uint32](typ), nil
	case Uint16T:
		return newValueMap[uint16](typ), nil
	case Uint8T:
		return newValueMap[uint8](typ), nil
	case UintT:
		return newValueMap[uint](typ), nil
	case Float64T:
		return newValueMap[float64](typ), nil
	case Float32T:
		return newValueMap[float32](typ), nil
	}
	return nil, fmt.Errorf("%w: bad map key type %s", ErrInvalidType, key)
}

// valMap wraps a map.
type valMap[K cmp.Ordered] struct {
	typ Type
	v   map[K]Value
}

// newValueMap creates a new value map.
func newValueMap[K cmp.Ordered](typ Type) valueMap {
	return &valMap[K]{
		typ: typ,
	}
}

func (val *valMap[K]) Set(s string) error {
	if val.v == nil {
		val.v = make(map[K]Value)
	}
	keystr, value, ok := strings.Cut(s, "=")
	if !ok || keystr == "" {
		return fmt.Errorf("%w %q: %s", ErrInvalidValue, s, "missing map key")
	}
	key, err := as[K](keystr, "")
	if err != nil {
		return fmt.Errorf("bad map key: %q: %w", keystr, err)
	}
	k, ok := key.(K)
	if !ok {
		return fmt.Errorf("%w: string->%T", ErrInvalidConversion, k)
	}
	v, ok := val.v[k]
	if !ok {
		if v, err = val.typ.New(); err != nil {
			return err
		}
	}
	if err := v.Set(value); err != nil {
		return err
	}
	val.v[k] = v
	return nil
}

func (val valMap[K]) String() string {
	s := make([]string, len(val.v))
	for i, k := range slices.Sorted(maps.Keys(val.v)) {
		value, _ := val.v[k].Get()
		s[i] = toString[string](k) + ":" + value
	}
	return "[" + strings.Join(s, " ") + "]"
}

func (val valMap[K]) Keys() []any {
	keys := make([]any, len(val.v))
	for i, k := range slices.Sorted(maps.Keys(val.v)) {
		keys[i] = k
	}
	return keys
}

func (val valMap[K]) Get(key any) Value {
	if k, ok := key.(K); ok {
		return val.v[k]
	}
	return nil
}

// Binder is the interface for binding values.
type Binder interface {
	Bind(string) error
}

// bindVal is a bound value.
type bindVal[T *E, E any] struct {
	v   T
	set *bool
}

// NewBind binds a value and its set flag.
func NewBind[T *E, E any](v T, set *bool) (Binder, error) {
	return &bindVal[T, E]{
		v:   v,
		set: set,
	}, nil
}

func (val *bindVal[T, E]) Bind(s string) error {
	typ := reflect.TypeOf(*val.v)
	switch typ.Kind() {
	case reflect.Slice:
		switch err := sliceSet(reflect.ValueOf(val.v), s); {
		case err != nil:
			return err
		case val.set != nil:
			*val.set = true
		}
		return nil
	case reflect.Map:
		switch err := mapSet(reflect.ValueOf(val.v), s); {
		case err != nil:
			return err
		case val.set != nil:
			*val.set = true
		}
		return nil
	default:
		if v, err := as[E](s, layout(val.v)); err == nil {
			var ok bool
			if *val.v, ok = v.(E); ok {
				if val.set != nil {
					*val.set = true
				}
				return nil
			}
		}
	}
	return fmt.Errorf("%w: %s->%T", ErrInvalidConversion, typ, *val.v)
}

func (val *bindVal[T, E]) String() string {
	return reflect.TypeOf(val.v).String()
}

func (val *bindVal[T, E]) Get() any {
	return *val.v
}

// bindRefVal is a reflection bound value.
type bindRefVal struct {
	v   reflect.Value
	set *bool
}

// NewBindRef binds [reflect.Value] value and its set flag.
func NewBindRef(value reflect.Value, set *bool) (Binder, error) {
	switch {
	case value.Kind() != reflect.Pointer, value.IsNil():
		return nil, fmt.Errorf("%w: not a pointer or is nil", ErrInvalidValue)
	}
	return &bindRefVal{
		v:   value,
		set: set,
	}, nil
}

func (val *bindRefVal) String() string {
	return val.v.Type().String()
}

func (val *bindRefVal) Bind(s string) error {
	typ := val.v.Elem().Type()
	switch typ.Kind() {
	case reflect.Slice:
		switch err := sliceSet(val.v, s); {
		case err != nil:
			return err
		case val.set != nil:
			*val.set = true
		}
		return nil
	case reflect.Map:
		switch err := mapSet(val.v, s); {
		case err != nil:
			return err
		case val.set != nil:
			*val.set = true
		}
		return nil
	case reflect.Pointer:
		v, err := asUnmarshal(reflectType(typ), s)
		if err != nil {
			return err
		}
		reflect.Indirect(val.v).Set(reflect.ValueOf(v))
		if val.set != nil {
			*val.set = true
		}
		return nil
	}
	switch err := asValue(val.v, s); {
	case err != nil:
		return err
	case val.set != nil:
		*val.set = true
	}
	return nil
}

func (val *bindRefVal) Get() any {
	return val.v.Elem().Interface()
}

// hookVal is a hook func.
type hookVal struct {
	typ Type
	ctx *Context
	set bool
	v   func(*Context, string) error
}

// newHook creates a new hook for the func f in v.
func newHook(ctx *Context, v any) (Value, error) {
	val := &hookVal{
		typ: HookT,
		ctx: ctx,
	}
	switch f := v.(type) {
	case func(*Context, string) error:
		val.v = f
	case func(*Context, string):
		val.v = func(ctx *Context, s string) error {
			f(ctx, s)
			return nil
		}
	case func(*Context) error:
		val.v = func(ctx *Context, _ string) error {
			return f(ctx)
		}
	case func(*Context):
		val.v = func(ctx *Context, _ string) error {
			f(ctx)
			return nil
		}
	case func(string) error:
		val.v = func(_ *Context, s string) error {
			return f(s)
		}
	case func(string):
		val.v = func(_ *Context, s string) error {
			f(s)
			return nil
		}
	case func() error:
		val.v = func(*Context, string) error {
			return f()
		}
	case func():
		val.v = func(*Context, string) error {
			f()
			return nil
		}
	default:
		return nil, fmt.Errorf("%w: invalid hook func %T", ErrInvalidValue, v)
	}
	return val, nil
}

func (val *hookVal) Type() Type {
	return val.typ
}

func (val *hookVal) Val() any {
	return val.v
}

func (val *hookVal) SetSet(set bool) {
	val.set = set
}

func (val *hookVal) WasSet() bool {
	return val.set
}

func (val *hookVal) Set(s string) error {
	return val.v(val.ctx, s)
}

func (val *hookVal) Get() (string, error) {
	return "(hook)", nil
}

func (val *hookVal) String() string {
	return "(hook)"
}

// FormattedTime wraps a time value with a specific layout.
type FormattedTime struct {
	layout string
	v      time.Time
}

// Layout returns the layout.
func (val FormattedTime) Layout() string {
	return val.layout
}

// Time returns the time value.
func (val FormattedTime) Time() time.Time {
	return val.v
}

// Set sets parses the time value from the string.
func (val *FormattedTime) Set(s string) error {
	var err error
	val.v, err = time.Parse(val.layout, s)
	return err
}

// Format formats the time value as the provided layout.
func (val FormattedTime) Format(layout string) string {
	return val.v.Format(layout)
}

// String satisfies the [fmt.Stringer] interface.
func (val FormattedTime) String() string {
	return val.v.Format(val.layout)
}

// IsValid returns true when the time is not zero.
func (val FormattedTime) IsValid() bool {
	return !val.v.IsZero()
}

// sliceSet sets a value on a slice -- expects reflect.ValueOf(&<target>).
func sliceSet(value reflect.Value, s string) error {
	v := reflect.New(value.Elem().Type().Elem())
	err := asValue(v, s)
	if err != nil {
		return err
	}
	reflect.Indirect(value).Set(reflect.Append(value.Elem(), reflect.Indirect(v)))
	return nil
}

// mapSet sets a value on a map -- expects reflect.ValueOf(&<target>).
func mapSet(val reflect.Value, s string) error {
	key, value, ok := strings.Cut(s, "=")
	if !ok || key == "" {
		return ErrInvalidValue
	}
	typ := val.Elem().Type()
	// create map if nil
	if val.Elem().IsNil() {
		reflect.Indirect(val).Set(reflect.MakeMap(typ))
	}
	// convert key
	k := reflect.New(typ.Key())
	if err := asValue(k, key); err != nil {
		return err
	}
	// convert value
	v := reflect.New(typ.Elem())
	if err := asValue(v, value); err != nil {
		return err
	}
	reflect.Indirect(val).SetMapIndex(reflect.Indirect(k), reflect.Indirect(v))
	return nil
}

// inc increments the value.
func inc(val any, delta uint64) {
	if v, ok := val.(*uint64); ok {
		atomic.AddUint64(v, delta)
	}
}

// invalid returns true if the value is invalid.
func invalid(val any) bool {
	switch v := val.(type) {
	case interface{ IsValid() bool }:
		// netip.{Addr,AddrPort,Prefix} and FormattedTime
		return !v.IsValid()
	case interface{ IsZero() bool }:
		return v.IsZero()
	}
	return false
}

// layout returns the layout for the value.
func layout(val any) string {
	if v, ok := val.(interface{ Layout() string }); ok {
		return v.Layout()
	}
	return ""
}
