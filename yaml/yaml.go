// Package yaml provides a yaml config reader for ox.
package yaml

import (
	"fmt"
	"sync"

	"github.com/goccy/go-yaml"
	"github.com/xo/ox"
)

func init() {
	ox.RegisterConfigLoader("yaml", LoadKey)
}

// DefaultDecoder is the default yaml decoder.
var DefaultDecoder = new(decoder)

// LoadKey loads a key from a yaml file.
func LoadKey(ctx *ox.Context, key string) (any, bool, error) {
	/*
		v, err := ctx.Load(
			DefaultDecoder,
			"yaml",
			ox.Extension("yaml", "yml"),
		)
		if err != nil {
			return nil, false, err
		}
	*/
	path, err := yaml.PathString(key)
	if err != nil {
		return nil, false, fmt.Errorf("%w: %w", ox.ErrInvalidKey, err)
	}
	path = path
	return "", false, nil
}

type decoder struct {
	opts []yaml.DecodeOption
	once sync.Once
}
