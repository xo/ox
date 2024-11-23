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
)

// Value is the value interface.
type Value interface {
	Type() Type
	Val() any
	SetSet(bool)
	WasSet() bool
	Set(string) error
	Get() (string, error)
}

// anyVal wraps a value.
type anyVal[T any] struct {
	typ   Type
	v     T
	noArg bool
	set   bool
}

// NewVal returns a [Value] func.
func NewVal[T any](opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		typ := typeType[T]()
		val := &anyVal[T]{
			typ:   typ,
			noArg: typ == BoolT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		if typ != val.typ && val.typ == CountT {
			val.noArg = true
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
	}
	v, err := as[T](value, val.typ.Layout())
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
	return As[string](val.v)
}

// sliceVal is a slice value.
type sliceVal struct {
	typ Type
	v   []Value
}

// NewSlice creates a slice value of type.
func NewSlice(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &sliceVal{
			typ: StringT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return nil, nil
		// return val, nil
	}
}

func (val *sliceVal) Type() Type {
	return "[]" + val.typ
}

func (val *sliceVal) Val() any {
	return val.v
}

func (val *sliceVal) Get() (string, error) {
	return string(val.Type()), nil
}

func (val *sliceVal) Set(s string) error {
	return nil
}

func (val *sliceVal) String() string {
	s := make([]string, len(val.v))
	for i, v := range val.v {
		s[i] = toString[string](v, val.typ.Layout())
	}
	return "[" + strings.Join(s, " ") + "]"
}

// Index returns the i'th variable from the slice.
func (val *sliceVal) Index(i int) Value {
	return val.v[i]
}

func (val *sliceVal) Len() int {
	return len(val.v)
}

// mapVal is a map value.
type mapVal[K cmp.Ordered] struct {
	v map[K]Value
}

// NewMap creates a map value of type.
func NewMap(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		/*
			val := &mapVal[string]{}
			for _, o := range opts {
				if err := o.apply(val); err != nil {
					return nil, err
				}
			}
			return val, nil
		*/
		return nil, nil
	}
}

func (val *mapVal[K]) Type() Type {
	// return "map[" + val.key + "]" + val.typ
	return ""
}

func (val *mapVal[K]) Val() any {
	return val.v
}

func (val *mapVal[K]) Get() (string, error) {
	return "", nil
}

func (val *mapVal[K]) Set(s string) error {
	if val.v == nil {
		val.v = make(map[K]Value)
	}
	key, value, ok := strings.Cut(s, "=")
	key, value = key, value
	if !ok || key == "" {
		return fmt.Errorf("%w: %s", ErrInvalidValue, "bad map key")
	}
	/*
		k, err := NewValue(val.key, "", "")
		if err != nil {
			return err
		}
		if err := k.Set(ctx, key); err != nil {
			return err
		}
		v, err := NewValue(val.typ, "", "")
		if err != nil {
			return err
		}
		if err := v.Set(ctx, value); err != nil {
			return err
		}
		val.v[k] = v
	*/
	return nil
}

func (val *mapVal[K]) String() string {
	/*
		s := make([]string, len(val.v))
		for i, k := range slices.Sorted(maps.Keys(val.v)) {
			value, _ := val.v[k].Get()
			s[i] = toString(k) + ":" + value
				s[i] = k + ":" + toString(val.v[k].Var)
				i++
		}
		return "[" + strings.Join(s, " ") + "]"
	*/
	return ""
}

// Vars is the type for storing variables in the context.
type Vars map[string]*VarSet

// String satisfies the [fmt.Stringer] interfaec.
func (vars Vars) String() string {
	var v []string
	for _, k := range slices.Sorted(maps.Keys(vars)) {
		if val, ok := vars[k].Var.(interface{ String() string }); ok {
			v = append(v, k+":"+val.String())
		} else {
			if s, err := vars[k].Var.Get(); err == nil {
				v = append(v, k+":"+s)
			}
		}
	}
	return "[" + strings.Join(v, " ") + "]"
}

// Set sets a variable in the vars.
func (vars Vars) Set(g *Flag, value any, wasSet bool) error {
	/*
		// fmt.Fprintf(os.Stdout, "SETTING: %q (%s/%s): %q\n", g.Name(), g.Type, g.Sub, value)
		name := g.Name()
		if name == "" {
			return ErrInvalidFlagName
		}
		vs, err := vars[name].Set(g.Type, g.Sub, value, wasSet)
		if err != nil {
			return err
		}
		for i, val := range g.Binds {
			if err := val.Set(value); err != nil {
				return fmt.Errorf("flag %s: bind %d (%T): cannot set %q: %w", g.Name(), i, val.Get(), value, err)
			}
		}
		vars[name] = vs
	*/
	return nil
}

// VarSet wraps a variable and its set state.
type VarSet struct {
	Type   Type
	Var    Value
	WasSet bool
}

// Set sets a var.
func (vs *VarSet) Set(typ, sub Type, value string, wasSet bool) (*VarSet, error) {
	if vs == nil {
		vs = &VarSet{
			Type: typ,
		}
	}
	if vs.Type != typ {
		return nil, ErrTypeMismatch
	}
	if vs.Var == nil {
		var err error
		if vs.Var, err = typ.New(); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrCouldNotCreateValue, err)
		}
		if sub != "" {
			if err := sub.apply(vs.Var); err != nil {
				return nil, err
			}
		}
	}
	if err := vs.Var.Set(value); err != nil {
		return nil, err
	}
	vs.WasSet = wasSet
	return vs, nil
}

// BoundValue is the bound value interface.
type BoundValue interface {
	Set(any) error
	Get() any
}

// boundVal is a bound value.
type boundVal[T *E, E any] struct {
	v T
	b *bool
}

// NewBind binds a value.
func NewBind[T *E, E any](v T, b *bool) (BoundValue, error) {
	return &boundVal[T, E]{
		v: v,
		b: b,
	}, nil
}

func (val *boundVal[T, E]) Set(v any) error {
	typ := reflect.TypeOf(*val.v)
	switch typ.Kind() {
	case reflect.Slice:
		/*
			if sliceSet(reflect.ValueOf(&val.v), s) {
				return nil
			}
		*/
	case reflect.Map:
		/*
			if mapSet(reflect.ValueOf(&val.v), s) {
				return nil
			}
		*/
	default:
		/*
			if v, err := conv[E](s); err == nil {
				var ok bool
				if *val.v, ok = v.(E); ok {
					if val.b != nil {
						*val.b = true
					}
					return nil
				}
			}
		*/
	}
	return fmt.Errorf("%w: %s->%T", ErrInvalidConversion, typ, *val.v)
}

func (val *boundVal[T, E]) Get() any {
	return *val.v
}

// boundRefVal is a value bound with reflection.
type boundRefVal struct {
	v reflect.Value
	b *bool
}

// NewBindValue binds [reflect.Value] value and its set flag.
func NewBindValue(value reflect.Value, b *bool) (BoundValue, error) {
	switch {
	case value.Kind() != reflect.Pointer, value.IsNil():
		return nil, fmt.Errorf("%w: not a pointer or is nil", ErrInvalidValue)
	}
	return &boundRefVal{
		v: value,
		b: b,
	}, nil
}

func (val *boundRefVal) Set(v any) error {
	typ := val.v.Elem().Type()
	switch typ.Kind() {
	case reflect.Slice:
		/*
			if val.sliceSet(s) {
				return nil
			}
		*/
	case reflect.Map:
		/*
			if val.mapSet(s) {
				return nil
			}
		*/
	default:
		/*
			if convValue(val.v, s) {
				return nil
			}
			if v, err := convUnmarshal(typ, s); err == nil {
				reflect.Indirect(val.v).Set(reflect.ValueOf(v))
				return nil
			}
		*/
	}
	return fmt.Errorf("%w: cannot convert %T->%s", ErrInvalidConversion, v, typ)
}

func (val *boundRefVal) sliceSet(s string) bool {
	/*
		z := reflect.New(val.v.Elem().Type().Elem())
		fmt.Fprintf(os.Stdout, "type: %s\n", z.Type())
		if !convValue(z, s) {
			return false
		}
		val.v.Elem().Set(reflect.Append(val.v.Elem(), reflect.Indirect(z)))
		return true
	*/
	return false
}

func (val *boundRefVal) mapSet(s string) bool {
	return false
}

func (val *boundRefVal) Get() any {
	return val.v.Elem().Interface()
}

// sliceSet sets a a value on a slice.
func sliceSet(val reflect.Value, s string) bool {
	/*
		z := reflect.New(v.Type().Elem())
		if convValue(z, s) {
			v = reflect.Append(v, reflect.Indirect(z))
		}
	*/
	return false
}

// mapSet sets a value on a map.
func mapSet(val reflect.Value, s string) bool {
	return false
	/*
		key, value, ok := strings.Cut(s, "=")
		if !ok || key == "" {
			return fmt.Errorf("%w: %s", ErrInvalidValue, "bad map key")
		}
		m := reflect.ValueOf(val.v).Elem()
		// create map if nil
		if m.IsNil() {
			m.Set(reflect.MakeMap(m.Type()))
			var ok bool
			if *val.v, ok = m.Interface().(E); !ok {
				return fmt.Errorf("%w: %T->%T", ErrInvalidConversion, m.Interface(), *val.v)
			}
		}
		// convert key
		k := reflect.New(m.Type().Key())
		if !convValue(k, key) {
			return fmt.Errorf("%w (key): %T->%T", ErrInvalidConversion, key, k.Interface())
		}
		// convert value
		v := reflect.New(m.Type().Elem())
		if !convValue(v, value) {
			return fmt.Errorf("%w (value): %T->%T", ErrInvalidConversion, value, v.Interface())
		}
		m.SetMapIndex(reflect.Indirect(k), reflect.Indirect(v))
		return nil
	*/
}
