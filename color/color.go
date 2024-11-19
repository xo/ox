// Package color provides a color type for kobra.
package color

import (
	"context"
	"image/color"

	"github.com/kenshaw/colors"
	"github.com/xo/kobra"
)

// Default is the default color.
var Default color.Color = colors.Transparent

func init() {
	kobra.RegisterType(kobra.ColorT, kobra.NewTypeDesc(kobra.NewText(func() (any, error) {
		c := colors.FromColor(Default)
		return &c, nil
	}, kobra.ColorT)))
}

// Color retrieves a color from the context.
func Color(ctx context.Context, name string) *colors.Color {
	return kobra.Get[*colors.Color](ctx, name)
}
