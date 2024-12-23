package ox

import (
	"context"
	"io"
)

// ConfigDecoder is the interface for configuration decoders.
type ConfigDecoder interface {
	Decode(context.Context, io.Reader, any) error
}

// configTypes
var configTypes map[string]func(...any) (ConfigDecoder, error)

func init() {
	configTypes = make(map[string]func(...any) (ConfigDecoder, error))
	configs = make(map[string][]ConfigValueDecoder)
}

// RegisterConfigFileType registers a config file type.
func RegisterConfigFileType(typ string, handler func(...any) (ConfigDecoder, error)) {
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
