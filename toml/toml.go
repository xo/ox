// Package toml provides a toml config reader for kobra.
package toml

import (
	"context"
	"io"

	"github.com/pelletier/go-toml/v2"
	"github.com/xo/kobra"
)

func init() {
	kobra.RegisterConfigFileType("toml", func(opts ...any) (kobra.ConfigDecoder, error) {
		d := new(decoder)
		for _, opt := range opts {
			if o, ok := opt.(func(*toml.Decoder)); ok {
				d.opts = append(d.opts, o)
			}
		}
		return d, nil
	})
}

type decoder struct {
	opts []func(*toml.Decoder)
}

// Decode satisfies the [kobra.ConfigDecoder] interface.
func (d *decoder) Decode(_ context.Context, r io.Reader, v any) error {
	dec := toml.NewDecoder(r)
	for _, o := range d.opts {
		o(dec)
	}
	return dec.Decode(v)
}

func DisallowUnknownFields(dec *toml.Decoder) {
	dec.DisallowUnknownFields()
}

func EnableUnmarshalerInterface(dec *toml.Decoder) {
	dec.EnableUnmarshalerInterface()
}
