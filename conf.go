package ox

import (
	"os"
)

// ConfigLoader is the interface for configuration decoders.
type ConfigLoader interface{}

// loaders are config loaders.
var loaders map[string]func(*Context, string) (string, error)

func init() {
	loaders = make(map[string]func(*Context, string) (string, error))
	RegisterConfigLoader("ENV", func(_ *Context, key string) (string, error) {
		return os.Getenv(key), nil
	})
}

// RegisterConfigLoader registers a config file type.
func RegisterConfigLoader(typ string, f func(*Context, string) (string, error)) {
	loaders[typ] = f
}
