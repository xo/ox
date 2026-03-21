// Package toml provides a toml config reader for ox.
package toml

import (
	"context"
	"io"

	"github.com/pelletier/go-toml/v2"
)

func init() {
	/*
		ox.RegisterConfigLoader("toml", func(opts ...any) (ox.ConfigLoader, error) {
				d := new(decoder)
				for _, opt := range opts {
					if o, ok := opt.(func(*toml.Decoder)); ok {
						d.opts = append(d.opts, o)
					}
				}
				return d, nil
			return nil, nil
		})
	*/
}

type decoder struct {
	opts []func(*toml.Decoder)
}

// Decode satisfies the [ox.ConfigLoader] interface.
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
