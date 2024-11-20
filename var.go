package kobra

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"maps"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// DefaultLayout is the default timestamp layout.
var DefaultLayout = time.RFC3339

// stringVal is a string value.
type stringVal struct {
	typ Type
	v   []byte
}

// NewString creates a new string value.
func NewString(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &stringVal{
			typ: StringT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *stringVal) Type() Type {
	return val.typ
}

func (val *stringVal) Val() any {
	return val.v
}

func (val *stringVal) Get() (string, error) {
	return string(val.v), nil
}

func (val *stringVal) Set(_ context.Context, s string) error {
	switch val.typ {
	case Base64T:
		var err error
		if val.v, err = base64.StdEncoding.DecodeString(s); err != nil {
			return errors.Join(ErrInvalidValue, err)
		}
	case HexT:
		var err error
		if val.v, err = hex.DecodeString(s); err != nil {
			return errors.Join(ErrInvalidValue, err)
		}
	default:
		val.v = []byte(s)
	}
	return nil
}

func (val *stringVal) MarshalText() ([]byte, error) {
	return slices.Clone(val.v), nil
}

func (val *stringVal) String() string {
	return string(val.v)
}

func (val *stringVal) Base64() string {
	return base64.StdEncoding.EncodeToString(val.v)
}

func (val *stringVal) Hex() string {
	return hex.EncodeToString(val.v)
}

func (val *stringVal) Bytes() []byte {
	return val.v
}

func (val *stringVal) Runes() []rune {
	return []rune(string(val.v))
}

// runeVal wraps a rune value.
type runeVal struct {
	typ Type
	v   rune
}

// NewRune creates a rune value.
func NewRune(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &runeVal{
			typ: RuneT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *runeVal) Type() Type {
	return val.typ
}

func (val *runeVal) Val() any {
	return val.v
}

func (val *runeVal) Get() (string, error) {
	return string(val.v), nil
}

func (val *runeVal) Rune() rune {
	return val.v
}

func (val *runeVal) Set(_ context.Context, s string) error {
	n, b := 0, []byte(s)
	if val.v, n = utf8.DecodeRune(b); val.v == utf8.RuneError || n != len(b) {
		return errors.Join(ErrInvalidValue, ErrInvalidRune)
	}
	return nil
}

// boolVal wraps a bool value.
type boolVal struct {
	typ Type
	v   bool
}

// NewBool creates a bool value.
func NewBool(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &boolVal{
			typ: BoolT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *boolVal) Type() Type {
	return val.typ
}

func (val *boolVal) Val() any {
	return val.v
}

func (val *boolVal) Get() (string, error) {
	return strconv.FormatBool(val.v), nil
}

func (val *boolVal) Bool() bool {
	return val.v
}

func (val *boolVal) Set(_ context.Context, s string) error {
	if s == "" {
		val.v = false
		return nil
	}
	b, ok := asBool(s)
	if !ok {
		return ErrInvalidValue
	}
	val.v = b
	return nil
}

// intVal wraps a int value.
type intVal[T inti] struct {
	typ     Type
	bitSize int
	v       T
}

// NewInt creates a int value.
func NewInt[T inti](opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &intVal[T]{
			typ: IntT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *intVal[T]) Type() Type {
	return val.typ
}

func (val *intVal[T]) Val() any {
	return val.v
}

func (val *intVal[T]) Get() (string, error) {
	return strconv.FormatInt(int64(val.v), 10), nil
}

func (val *intVal[T]) Int64() int64 {
	return int64(val.v)
}

func (val *intVal[T]) Set(_ context.Context, s string) error {
	if s == "" {
		val.v = 0
		return nil
	}
	i, err := strconv.ParseInt(s, 10, val.bitSize)
	if err != nil {
		return ErrInvalidValue
	}
	val.v = T(i)
	return nil
}

// uintVal wraps a uint value.
type uintVal[T uinti] struct {
	typ     Type
	bitSize int
	v       T
}

// NewUint creates a uint value.
func NewUint[T uinti](opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &uintVal[T]{
			typ: UintT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *uintVal[T]) Type() Type {
	return val.typ
}

func (val *uintVal[T]) Val() any {
	return val.v
}

func (val *uintVal[T]) Get() (string, error) {
	return strconv.FormatUint(uint64(val.v), 10), nil
}

func (val *uintVal[T]) Uint64() uint64 {
	return uint64(val.v)
}

func (val *uintVal[T]) Set(_ context.Context, s string) error {
	if s == "" {
		val.v = 0
		return nil
	}
	u, err := strconv.ParseUint(s, 10, val.bitSize)
	if err != nil {
		return err
	}
	val.v = T(u)
	return nil
}

// floatVal wraps a float value.
type floatVal[T floati] struct {
	typ     Type
	bitSize int
	v       T
}

// NewFloat creates a new float value.
func NewFloat[T floati](opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &floatVal[T]{
			typ:     Float64T,
			bitSize: 64,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *floatVal[T]) Type() Type {
	return val.typ
}

func (val *floatVal[T]) Val() any {
	return val.v
}

func (val *floatVal[T]) Get() (string, error) {
	return strconv.FormatFloat(float64(val.v), 'f', -1, val.bitSize), nil
}

func (val *floatVal[T]) Float64() float64 {
	return float64(val.v)
}

func (val *floatVal[T]) Set(_ context.Context, s string) error {
	if s == "" {
		val.v = 0.0
		return nil
	}
	f, err := strconv.ParseFloat(s, val.bitSize)
	if err != nil {
		return ErrInvalidValue
	}
	val.v = T(f)
	return nil
}

// complexVal wraps a complex value.
type complexVal[T complexi] struct {
	typ     Type
	bitSize int
	v       T
}

// NewComplex creates a complex value.
func NewComplex[T complexi](opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &complexVal[T]{
			typ:     Complex128T,
			bitSize: 128,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *complexVal[T]) Type() Type {
	return val.typ
}

func (val *complexVal[T]) Val() any {
	return val.v
}

func (val *complexVal[T]) Get() (string, error) {
	return strconv.FormatComplex(complex128(val.v), 'f', -1, val.bitSize), nil
}

func (val *complexVal[T]) Complex128() complex128 {
	return complex128(val.v)
}

func (val *complexVal[T]) Set(_ context.Context, s string) error {
	if s == "" {
		val.v = 0i
		return nil
	}
	c, err := strconv.ParseComplex(s, val.bitSize)
	if err != nil {
		return err
	}
	val.v = T(c)
	return nil
}

// timeVal wraps a time value.
type timeVal struct {
	typ    Type
	layout string
	v      time.Time
}

// NewTime creates a time value.
func NewTime(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &timeVal{
			typ:    TimeT,
			layout: DefaultLayout,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *timeVal) Type() Type {
	return val.typ
}

func (val *timeVal) Val() any {
	return val.v
}

func (val *timeVal) Time() time.Time {
	return val.v
}

func (val *timeVal) Get() (string, error) {
	return val.v.Format(val.layout), nil
}

func (val *timeVal) Set(_ context.Context, s string) error {
	if s == "" {
		val.v = time.Time{}
		return nil
	}
	var err error
	val.v, err = time.Parse(val.layout, s)
	return err
}

// durationVal wraps a [time.Duration] value.
type durationVal struct {
	typ Type
	v   time.Duration
}

// NewDuration creates a duration value.
func NewDuration(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &durationVal{
			typ: DurationT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *durationVal) Type() Type {
	return val.typ
}

func (val *durationVal) Val() any {
	return val.v
}

func (val *durationVal) Get() (string, error) {
	return val.v.String(), nil
}

func (val *durationVal) Duration() time.Duration {
	return val.v
}

func (val *durationVal) Set(_ context.Context, s string) error {
	if s == "" {
		val.v = 0
		return nil
	}
	var err error
	val.v, err = time.ParseDuration(s)
	return err
}

// marshalVal wraps a marshal value.
type marshalVal struct {
	typ       Type
	v         any
	marshal   func() ([]byte, error)
	unmarshal func([]byte) error
	text      bool
}

func newMarshalValue(v any, marshal func() ([]byte, error), unmarshal func([]byte) error, text bool, opts ...Option) (Value, error) {
	val := &marshalVal{
		v:         v,
		marshal:   marshal,
		unmarshal: unmarshal,
		text:      text,
	}
	for _, o := range opts {
		if err := o.apply(val); err != nil {
			return nil, err
		}
	}
	return val, nil
}

// NewText creates a value for a type that supports [encoding.TextMarshaler]
// and [enoding.TextUnmarshaler].
func NewText(f func() (any, error), opts ...Option) func() (Value, error) {
	registerMarshaler(f, textNew)
	return func() (Value, error) {
		v, err := f()
		if err != nil {
			return nil, errors.Join(ErrInvalidValue, err)
		}
		i, ok := v.(texti)
		if !ok {
			return nil, ErrInvalidTextType
		}
		return newMarshalValue(v, i.MarshalText, i.UnmarshalText, true, prepend(opts, Option(StringT))...)
	}
}

// NewBinary creates a value for a type that supports [encoding.BinaryMarshaler]
// and [enoding.BinaryUnmarshaler].
func NewBinary(f func() (any, error), opts ...Option) func() (Value, error) {
	registerMarshaler(f, binaryNew)
	return func() (Value, error) {
		v, err := f()
		if err != nil {
			return nil, errors.Join(ErrInvalidValue, err)
		}
		i, ok := v.(binaryi)
		if !ok {
			return nil, ErrInvalidBinaryType
		}
		return newMarshalValue(v, i.MarshalBinary, i.UnmarshalBinary, false, prepend(opts, Option(StringT))...)
	}
}

func (val *marshalVal) Type() Type {
	return val.typ
}

func (val *marshalVal) Val() any {
	return val.v
}

func (val *marshalVal) Get() (string, error) {
	b, err := val.marshal()
	if err != nil {
		return "", errors.Join(ErrInvalidValue, err)
	}
	return string(b), nil
}

func (val *marshalVal) MarshalText() ([]byte, error) {
	return val.marshal()
}

func (val *marshalVal) UnmarshalText(b []byte) error {
	return val.unmarshal(b)
}

func (val *marshalVal) MarshalBinary() ([]byte, error) {
	return val.marshal()
}

func (val *marshalVal) UnmarshalBinary(b []byte) error {
	return val.unmarshal(b)
}

func (val *marshalVal) Set(_ context.Context, s string) error {
	if err := val.unmarshal([]byte(s)); err != nil {
		return errors.Join(ErrInvalidValue, err)
	}
	return nil
}

// countVal is a count value.
type countVal struct {
	typ Type
	v   int64
}

// NewCount creates a count value.
func NewCount(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &countVal{
			typ: CountT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *countVal) Type() Type {
	return val.typ
}

func (val *countVal) Val() any {
	return val.v
}

func (val *countVal) Get() (string, error) {
	return strconv.FormatInt(val.v, 10), nil
}

func (val *countVal) Int64() int64 {
	return val.v
}

func (val *countVal) Count() int {
	return int(val.v)
}

func (val *countVal) Set(context.Context, string) error {
	val.v++
	return nil
}

// pathVal is a path value.
type pathVal struct {
	typ Type
	dir string
	v   string
}

// NewPath creates a path value.
func NewPath(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &pathVal{
			typ: PathT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *pathVal) Type() Type {
	return val.typ
}

func (val *pathVal) Val() any {
	return val.v
}

func (val *pathVal) Get() (string, error) {
	return val.v, nil
}

func (val *pathVal) Path() string {
	return val.v
}

func (val *pathVal) Set(_ context.Context, s string) error {
	val.v = s
	return nil
}

type hookVal struct {
	typ Type
	v   func(context.Context) error
}

// NewHook returns a hook.
func NewHook(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &hookVal{
			typ: HookT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *hookVal) Type() Type {
	return val.typ
}

func (val *hookVal) Val() any {
	return val.v
}

func (val *hookVal) Get() (string, error) {
	return "hook", nil
}

func (val *hookVal) Set(ctx context.Context, _ string) error {
	return val.v(ctx)
}

// sliceVal is a slice value.
type sliceVal struct {
	typ Type
	v   []*VarSet
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
		return val, nil
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

func (val *sliceVal) Set(ctx context.Context, s string) error {
	var vs *VarSet
	vs, err := vs.Set(ctx, val.typ, "", s, true)
	if err != nil {
		return err
	}
	val.v = append(val.v, vs)
	return nil
}

func (val *sliceVal) String() string {
	s := make([]string, len(val.v))
	for i, v := range val.v {
		s[i] = toString(v.Var.Val())
	}
	return "[" + strings.Join(s, " ") + "]"
}

// Index returns the i'th variable from the slice.
func (val *sliceVal) Index(i int) *VarSet {
	return val.v[i]
}

func (val *sliceVal) Len() int {
	return len(val.v)
}

// mapVal is a map value.
type mapVal struct {
	typ Type
	v   map[string]*VarSet
}

// NewMap creates a map value of type.
func NewMap(opts ...Option) func() (Value, error) {
	return func() (Value, error) {
		val := &mapVal{
			typ: StringT,
		}
		for _, o := range opts {
			if err := o.apply(val); err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}

func (val *mapVal) Type() Type {
	return "map[string]" + val.typ
}

func (val *mapVal) Val() any {
	return val.v
}

func (val *mapVal) Get() (string, error) {
	return string(val.Type()), nil
}

func (val *mapVal) Set(ctx context.Context, s string) error {
	if val.v == nil {
		val.v = make(map[string]*VarSet)
	}
	k, value, ok := strings.Cut(s, "=")
	if !ok {
		return ErrInvalidMapValue
	}
	var vs *VarSet
	vs, err := vs.Set(ctx, val.typ, "", value, true)
	if err != nil {
		return err
	}
	val.v[k] = vs
	return nil
}

func (val *mapVal) String() string {
	i, s := 0, make([]string, len(val.v))
	for _, k := range slices.Sorted(maps.Keys(val.v)) {
		s[i] = k + ":" + toString(val.v[k].Var)
		i++
	}
	return "[" + strings.Join(s, " ") + "]"
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
func (vars Vars) Set(ctx context.Context, g *Flag, value string, wasSet bool) error {
	// fmt.Fprintf(os.Stdout, "SETTING: %q (%s/%s): %q\n", g.Descs[0].Name, g.Type, g.Sub, value)
	vs, err := vars[g.Descs[0].Name].Set(ctx, g.Type, g.Sub, value, wasSet)
	if err != nil {
		return err
	}
	vars[g.Descs[0].Name] = vs
	return nil
}

// VarSet wraps a variable and its set state.
type VarSet struct {
	Type   Type
	Var    Value
	WasSet bool
}

// Set sets a var.
func (vs *VarSet) Set(ctx context.Context, typ, sub Type, value string, wasSet bool) (*VarSet, error) {
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
			return nil, errors.Join(ErrCouldNotCreateValue, err)
		}
		if sub != "" {
			if err := sub.apply(vs.Var); err != nil {
				return nil, err
			}
		}
	}
	if err := vs.Var.Set(ctx, value); err != nil {
		return nil, err
	}
	vs.WasSet = wasSet
	return vs, nil
}

// Value is the value interface.
type Value interface {
	Type() Type
	Set(context.Context, string) error
	Get() (string, error)
	Val() any
}
