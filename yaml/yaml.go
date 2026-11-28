// Package yaml registers the YAML config file format for ox. Import it for
// its side effect:
//
//	import _ "github.com/xo/ox/yaml"
//
// Then a program reads a YAML config file with [ox.ConfigFiles], and a flag
// default reads a key with ${yaml::key}.
package yaml

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"

	"github.com/xo/ox"
)

func init() {
	ox.RegisterConfigFormat(ox.YAMLT, Decode, "yaml", "yml")
}

// Decode decodes YAML data as an [ox.Config].
func Decode(data []byte) (ox.Config, error) {
	var m map[string]any
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return &Config{
		data: data,
		m:    ox.MapConfig(m),
	}, nil
}

// Config is a decoded YAML config file.
type Config struct {
	data []byte
	m    ox.MapConfig
}

// Lookup satisfies the [ox.Config] interface. A key that starts with "$" is a
// YAML path, such as "$.store.book[*].author". Any other key is a dotted
// path, such as "server.port".
func (cfg *Config) Lookup(key string) (any, bool, error) {
	if !strings.HasPrefix(key, "$") {
		return cfg.m.Lookup(key)
	}
	path, err := yaml.PathString(key)
	if err != nil {
		return nil, false, fmt.Errorf("%w: %w", ox.ErrInvalidKey, err)
	}
	var v any
	switch err := path.Read(bytes.NewReader(cfg.data), &v); {
	case yaml.IsNotFoundNodeError(err):
		return nil, false, nil
	case err != nil:
		return nil, false, err
	}
	return v, true, nil
}
