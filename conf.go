package ox

import (
	"context"
	"io"
)

// ConfigLoader is the interface for configuration decoders.
type ConfigLoader interface {
	Load(context.Context, io.Reader) (ConfigValueDecoder, error)
}

// configTypes
var configTypes map[string]func(...any) (ConfigLoader, error)

func init() {
	configTypes = make(map[string]func(...any) (ConfigLoader, error))
	configs = make(map[string][]ConfigValueDecoder)
}

// RegisterConfigLoader registers a config file type.
func RegisterConfigLoader(typ string, handler func(...any) (ConfigLoader, error)) {
	configTypes[typ] = handler
}

// ConfigGetter
type ConfigValueDecoder interface {
	Decode(string, any) error
}

var configs map[string][]ConfigValueDecoder

// RegisterConfig
func RegisterConfig(typ string, h ConfigValueDecoder) {
	configs[typ] = append(configs[typ], h)
}
