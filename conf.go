package ox

import (
	"context"
	"io"
)

// ConfigDecoder
type ConfigDecoder interface {
	Decode(context.Context, io.Reader, any) error
}

// configTypes
var configTypes map[string]func(...any) (ConfigDecoder, error)

func init() {
	configTypes = make(map[string]func(...any) (ConfigDecoder, error))
}

// RegisterConfigFileType registers a config file type.
func RegisterConfigFileType(typ string, handler func(...any) (ConfigDecoder, error)) {
	configTypes[typ] = handler
}
