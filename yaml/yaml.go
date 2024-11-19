// Package yaml provides a yaml config reader for kobra.
package yaml

import (
	"context"
	"io"

	"github.com/goccy/go-yaml"
	"github.com/xo/kobra"
)

func init() {
	kobra.RegisterConfigFileType("yaml", func(opts ...any) (kobra.ConfigDecoder, error) {
		d := new(decoder)
		for _, opt := range opts {
			if o, ok := opt.(yaml.DecodeOption); ok {
				d.opts = append(d.opts, o)
			}
		}
		return d, nil
	})
}

// decoder is a yaml decoder.
type decoder struct {
	opts []yaml.DecodeOption
}

// Decode satisfies the [kobra.ConfigDecoder] interface.
func (d *decoder) Decode(_ context.Context, r io.Reader, v any) error {
	return yaml.NewDecoder(r, d.opts...).Decode(v)
}
