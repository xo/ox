// Package color provides a color type for ox.
package color

import (
	"context"
	"image/color"

	"github.com/kenshaw/colors"
	"github.com/xo/ox"
)

// Default is the default color.
var Default color.Color = colors.Transparent

func init() {
	ox.RegisterTextType[*colors.Color](New)
}

func New() (*colors.Color, error) {
	c := colors.FromColor(Default)
	return &c, nil
}

// Color retrieves a color from the context.
func Color(ctx context.Context, name string) *colors.Color {
	return ox.Get[*colors.Color](ctx, name)
}
