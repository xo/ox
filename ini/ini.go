// Package ini registers the INI config file format for ox. Import it for its
// side effect:
//
//	import _ "github.com/xo/ox/ini"
//
// Then a program reads an INI config file with [ox.ConfigFiles], and a flag
// default reads a key with ${ini::key}. The package uses only the standard
// library, so it is in the module of ox, and a program that does not import
// it does not link it.
//
// The format:
//
//	; a comment, and # starts a comment too
//	verbose = true
//
//	[server]
//	host = db.example.com
//	port: 5432
//
//	[migrate.source]
//	dir = "./migrations"
//	tag = a
//	tag = b
//
// A key before the first section is at the top. A section nests its keys
// under its name, and a dot in the name of a section or of a key nests it
// again. So the key of dir above is "migrate.source.dir". A key that is
// written more than once is a list, so tag above is ["a", "b"]. A value in
// double quotes is a Go string, and a value in single quotes is the text
// between them. Any other value is the text after the = or the :, without
// the spaces around it.
package ini

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/xo/ox"
)

func init() {
	ox.RegisterConfigFormat(ox.INIT, Decode, "ini")
}

// Decode decodes INI data as an [ox.Config]. A key is a dotted path, such as
// "server.port".
func Decode(data []byte) (ox.Config, error) {
	m := make(map[string]any)
	section := m
	s := bufio.NewScanner(bytes.NewReader(data))
	for n := 1; s.Scan(); n++ {
		line := strings.TrimSpace(s.Text())
		switch {
		case line == "", line[0] == ';', line[0] == '#':
			continue
		case line[0] == '[':
			name, ok := strings.CutSuffix(line, "]")
			if name = strings.TrimSpace(name[1:]); !ok || name == "" {
				return nil, fmt.Errorf("line %d: %w: invalid section %q", n, ox.ErrInvalidConfig, line)
			}
			var err error
			if section, err = nest(m, strings.Split(name, ".")); err != nil {
				return nil, fmt.Errorf("line %d: %w", n, err)
			}
			continue
		}
		i := strings.IndexAny(line, "=:")
		if i <= 0 {
			return nil, fmt.Errorf("line %d: %w: no = in %q", n, ox.ErrInvalidConfig, line)
		}
		key, val := strings.TrimSpace(line[:i]), strings.TrimSpace(line[i+1:])
		v, err := unquote(val)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w: %s: %w", n, ox.ErrInvalidConfig, key, err)
		}
		parts := strings.Split(key, ".")
		parent, err := nest(section, parts[:len(parts)-1])
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", n, err)
		}
		if err := set(parent, parts[len(parts)-1], v); err != nil {
			return nil, fmt.Errorf("line %d: %w", n, err)
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return ox.MapConfig(m), nil
}

// nest returns the map at the path of names under m, and makes each map that
// is missing.
func nest(m map[string]any, names []string) (map[string]any, error) {
	for _, name := range names {
		if name = strings.TrimSpace(name); name == "" {
			return nil, fmt.Errorf("%w: an empty name", ox.ErrInvalidConfig)
		}
		switch v := m[name].(type) {
		case nil:
			z := make(map[string]any)
			m[name], m = z, z
		case map[string]any:
			m = v
		default:
			return nil, fmt.Errorf("%w: %s is a value and a section", ox.ErrInvalidConfig, name)
		}
	}
	return m, nil
}

// set sets the key of m to v. A key that m already has becomes a list.
func set(m map[string]any, key string, v string) error {
	switch z := m[key].(type) {
	case nil:
		m[key] = v
	case string:
		m[key] = []any{z, v}
	case []any:
		m[key] = append(z, v)
	default:
		return fmt.Errorf("%w: %s is a section and a value", ox.ErrInvalidConfig, key)
	}
	return nil
}

// unquote returns the value of s. A value in double quotes is a Go string,
// and a value in single quotes is the text between them.
func unquote(s string) (string, error) {
	switch {
	case len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"':
		return strconv.Unquote(s)
	case len(s) >= 2 && s[0] == '\'' && s[len(s)-1] == '\'':
		return s[1 : len(s)-1], nil
	}
	return s, nil
}
