package toml

import (
	"encoding/json"
	"testing"
)

// TestDecode decodes TOML and looks up dotted keys.
func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		data string
		// keys holds the value of each key, and nil for a key that the file
		// does not have
		keys map[string]any
		err  bool
	}{
		{
			"values",
			"port = 8080\nname = \"db\"\non = true\nrate = 1.5\ntags = [\"a\", \"b\"]\n",
			map[string]any{"port": 8080, "name": "db", "on": true, "rate": 1.5, "tags": []any{"a", "b"}, "tags.0": "a", "tags.2": nil},
			false,
		},
		{
			"tables",
			"[migrate]\nsteps = 3\n\n[migrate.source]\ndir = \"./up\"\n",
			map[string]any{"migrate.steps": 3, "migrate.source.dir": "./up", "migrate.other": nil, "steps": nil},
			false,
		},
		{
			"dotted keys",
			"migrate.source.dir = \"./up\"\n",
			map[string]any{"migrate.source.dir": "./up"},
			false,
		},
		{
			"an array of tables",
			"[[server]]\nhost = \"a\"\n[[server]]\nhost = \"b\"\n",
			map[string]any{"server.0.host": "a", "server.1.host": "b", "server.2.host": nil},
			false,
		},
		{"empty", "", map[string]any{"a": nil}, false},
		{"a syntax error", "port = \n", nil, true},
		{"a key twice", "a = 1\na = 2\n", nil, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg, err := Decode([]byte(test.data))
			switch {
			case test.err && err == nil:
				t.Fatalf("expected an error")
			case !test.err && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case test.err:
				return
			}
			for key, exp := range test.keys {
				v, ok, err := cfg.Lookup(key)
				switch {
				case err != nil:
					t.Fatalf("%s: expected no error, got: %v", key, err)
				case exp == nil && ok:
					t.Errorf("%s: expected no value, got: %v", key, v)
				case exp == nil:
				case !ok:
					t.Errorf("%s: expected a value", key)
				default:
					if a, b := mustJSON(t, v), mustJSON(t, exp); a != b {
						t.Errorf("%s: expected %s, got: %s", key, b, a)
					}
				}
			}
		})
	}
}

// mustJSON returns v as JSON, so that the test compares a slice or a map.
func mustJSON(t *testing.T, v any) string {
	t.Helper()
	buf, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(buf)
}
