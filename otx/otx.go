// Package otx provides [context.Context] access methods for xo/ox.
package otx

import (
	"cmp"
	"context"
	"math/big"
	"net/netip"
	"net/url"
	"time"

	"github.com/xo/ox"
)

// Vars returns all variables from the context.
func Vars(ctx context.Context) (ox.Vars, bool) {
	if c, ok := ox.Ctx(ctx); ok && c != nil {
		return c.Vars, true
	}
	return nil, false
}

// Any returns a variable, its set status, and if it was defined from the
// context.
func Any(ctx context.Context, name string) (ox.Value, bool) {
	if vars, ok := Vars(ctx); ok {
		if val, ok := vars[name]; ok {
			return val, true
		}
	}
	return nil, false
}

// Get returns a variable.
func Get[T any](ctx context.Context, name string) T {
	if val, ok := Any(ctx, name); ok {
		if v, err := ox.As[T](val); err == nil {
			return v
		}
	}
	var v T
	return v
}

// Slice returns the slice variable from the context as a slice of type E.
func Slice[E any](ctx context.Context, name string) []E {
	if val, ok := Any(ctx, name); ok {
		if v, err := ox.AsSlice[E](val); err == nil {
			return v
		}
	}
	return nil
}

// Map returns the map variable from the context.
func Map[K cmp.Ordered, T any](ctx context.Context, name string) map[K]T {
	if val, ok := Any(ctx, name); ok {
		if m, err := ox.AsMap[K, T](val); err == nil {
			return m
		}
	}
	return make(map[K]T)
}

// Bytes returns a variable as []byte from the context.
func Bytes(ctx context.Context, name string) []byte {
	return Get[[]byte](ctx, name)
}

// String returns the string variable from the context.
func String(ctx context.Context, name string) string {
	return Get[string](ctx, name)
}

// Runes returns a variable as []rune from the context.
func Runes(ctx context.Context, name string) []rune {
	return Get[[]rune](ctx, name)
}

// Bool returns the bool variable from the context.
func Bool(ctx context.Context, name string) bool {
	return Get[bool](ctx, name)
}

// Byte returns the byte variable from the context.
func Byte(ctx context.Context, name string) byte {
	if b := Get[[]byte](ctx, name); len(b) != 0 {
		return b[0]
	}
	return 0
}

// Rune returns the rune variable from the context.
func Rune(ctx context.Context, name string) rune {
	if r := Get[[]rune](ctx, name); len(r) != 0 {
		return r[0]
	}
	return 0
}

// Int64 returns the int64 variable from the context.
func Int64(ctx context.Context, name string) int64 {
	return Get[int64](ctx, name)
}

// Int32 returns the int32 variable from the context.
func Int32(ctx context.Context, name string) int32 {
	return Get[int32](ctx, name)
}

// Int16 returns the int16 variable from the context.
func Int16(ctx context.Context, name string) int16 {
	return Get[int16](ctx, name)
}

// Int returns the int variable from the context.
func Int(ctx context.Context, name string) int {
	return Get[int](ctx, name)
}

// Uint64 returns the uint64 variable from the context.
func Uint64(ctx context.Context, name string) uint64 {
	return Get[uint64](ctx, name)
}

// Uint32 returns the uint32 variable from the context.
func Uint32(ctx context.Context, name string) uint32 {
	return Get[uint32](ctx, name)
}

// Uint16 returns the uint16 variable from the context.
func Uint16(ctx context.Context, name string) uint16 {
	return Get[uint16](ctx, name)
}

// Uint8 returns the uint8 variable from the context.
func Uint8(ctx context.Context, name string) uint8 {
	return Get[uint8](ctx, name)
}

// Uint returns the uint variable from the context.
func Uint(ctx context.Context, name string) uint {
	return Get[uint](ctx, name)
}

// Float64 returns the float64 variable from the context.
func Float64(ctx context.Context, name string) float64 {
	return Get[float64](ctx, name)
}

// Float32 returns the float32 variable from the context.
func Float32(ctx context.Context, name string) float32 {
	return Get[float32](ctx, name)
}

// Complex128 returns the complex128 variable from the context.
func Complex128(ctx context.Context, name string) complex128 {
	return Get[complex128](ctx, name)
}

// Complex64 returns the complex64 variable from the context.
func Complex64(ctx context.Context, name string) complex64 {
	return Get[complex64](ctx, name)
}

// BigInt returns the [big.Int] variable from the context.
func BigInt(ctx context.Context, name string) *big.Int {
	return Get[*big.Int](ctx, name)
}

// BigFloat returns the [big.Float] variable from the context.
func BigFloat(ctx context.Context, name string) *big.Float {
	return Get[*big.Float](ctx, name)
}

// BigRat returns the [big.Rat] variable from the context.
func BigRat(ctx context.Context, name string) *big.Rat {
	return Get[*big.Rat](ctx, name)
}

// Time returns the [ox.Time] variable from the context.
func Time(ctx context.Context, name string) ox.FormattedTime {
	return Get[ox.FormattedTime](ctx, name)
}

// TimeTime returns the [time.Time] variable from the context.
func TimeTime(ctx context.Context, name string) time.Time {
	return Get[ox.FormattedTime](ctx, name).Time()
}

// Duration returns the [time.Duration] variable from the context.
func Duration(ctx context.Context, name string) time.Duration {
	return Get[time.Duration](ctx, name)
}

// URL returns the [url.URL] variable from the context.
func URL(ctx context.Context, name string) *url.URL {
	return Get[*url.URL](ctx, name)
}

// Addr returns the [netip.Addr] variable from the context.
func Addr(ctx context.Context, name string) *netip.Addr {
	return Get[*netip.Addr](ctx, name)
}

// AddrPort returns the [netip.AddrPort] variable from the context.
func AddrPort(ctx context.Context, name string) *netip.AddrPort {
	return Get[*netip.AddrPort](ctx, name)
}

// CIDR returns the [netip.Prefix] variable from the context.
func CIDR(ctx context.Context, name string) *netip.Prefix {
	return Get[*netip.Prefix](ctx, name)
}

// Path returns the path variable from the context.
func Path(ctx context.Context, name string) string {
	return Get[string](ctx, name)
}

// Count returns the count variable from the context.
func Count(ctx context.Context, name string) int {
	return Get[int](ctx, name)
}
