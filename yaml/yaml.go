// Package yaml provides a yaml config reader for ox.
package yaml

import (
	"sync"

	"github.com/goccy/go-yaml"
	"github.com/xo/ox"
)

func init() {
	ox.RegisterConfigLoader("yaml", func(opts ...any) (ox.ConfigLoader, error) {
		/*
			d := new(decoder)
			for _, opt := range opts {
				switch v := opt.(type) {
				case func(*decoder) error:
					if err := v(d); err != nil {
						return nil, err
					}
				case yaml.DecodeOption:
					d.opts = append(d.opts, v)
				}
			}
			return d, nil
		*/
		return nil, nil
	})
}

type decoder struct {
	opts []yaml.DecodeOption
	once sync.Once
}

// func From(r io.Reader)

func File(name string) func(*decoder) error {
	return func(d *decoder) error {
		return nil
	}
}

/*
// decoder is a yaml decoder.
type decoder struct{}

// Decode satisfies the [ox.ConfigLoader] interface.
func (d *decoder) Decode(_ context.Context, r io.Reader, v any) error {
	return yaml.NewDecoder(r, d.opts...).Decode(v)
}
*/
