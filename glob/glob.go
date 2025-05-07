// Package glob provides a ox type for glob processing.
package glob

import (
	"context"

	"github.com/kenshaw/glob"
	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

func init() {
	ox.RegisterTypeName(ox.GlobT, "*glob.Glob")
	ox.RegisterTextType(New)
}

// New creates a new glob.
func New() (*glob.Glob, error) {
	return glob.New(), nil
}

// Glob returns the glob var from the context.
func Glob(ctx context.Context, name string) *glob.Glob {
	return otx.Get[*glob.Glob](ctx, name)
}
