// Package toml registers the TOML config file format for ox. Import it for
// its side effect:
//
//	import _ "github.com/xo/ox/toml"
//
// Then a program reads a TOML config file with [ox.ConfigFiles], and a flag
// default reads a key with ${toml::key}.
package toml

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/xo/ox"
)

func init() {
	ox.RegisterConfigFormat(ox.TOMLT, Decode, "toml")
}

// Decode decodes TOML data as an [ox.Config]. A key is a dotted path, such
// as "server.port".
func Decode(data []byte) (ox.Config, error) {
	var m map[string]any
	if err := toml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return ox.MapConfig(m), nil
}
