// Package glob provides a ox type for glob processing.
package glob

import (
	"context"
	"slices"

	"github.com/gobwas/glob"
	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

func init() {
	ox.RegisterTypeName(ox.GlobT, "*glob.GlobValue")
	ox.RegisterTextType(New)
}

// Glob returns the glob var from the context.
func Glob(ctx context.Context, name string) *GlobValue {
	return otx.Get[*GlobValue](ctx, name)
}

// GlobValue wraps a [glob.Glob] value.
type GlobValue struct {
	glob.Glob
	text []byte
}

// New creates a new glob value.
func New() (*GlobValue, error) {
	return new(GlobValue), nil
}

// String satisfies the [fmt.Stringer] interface.
func (val *GlobValue) String() string {
	return string(val.text)
}

// UnmarshalText satisfies the [ox.BinaryMarshalUnmarshaler] interface.
func (val *GlobValue) UnmarshalText(buf []byte) error {
	var err error
	if val.Glob, err = glob.Compile(string(buf)); err != nil {
		return err
	}
	val.text = slices.Clone(buf)
	return nil
}

// MarshalText satisfies the [ox.BinaryMarshalUnmarshaler] interface.
func (val *GlobValue) MarshalText() ([]byte, error) {
	return slices.Clone(val.text), nil
}
