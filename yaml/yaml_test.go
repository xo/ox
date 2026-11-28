package yaml

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/xo/ox"
)

// TestDecode decodes YAML and looks up dotted keys and YAML paths.
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
			"port: 8080\nname: db\non: true\nrate: 1.5\ntags: [a, b]\n",
			map[string]any{"port": 8080, "name": "db", "on": true, "rate": 1.5, "tags": []any{"a", "b"}, "tags.1": "b", "tags.2": nil},
			false,
		},
		{
			"nested maps",
			"migrate:\n  steps: 3\n  source:\n    dir: ./up\n",
			map[string]any{"migrate.steps": 3, "migrate.source.dir": "./up", "migrate.source": map[string]any{"dir": "./up"}, "migrate.other": nil, "steps": nil},
			false,
		},
		{
			"a literal key with a dot",
			"a.b: literal\na:\n  b: nested\n",
			map[string]any{"a.b": "literal"},
			false,
		},
		{
			"a YAML path",
			"store:\n  book:\n    - author: jon\n    - author: ken\n",
			map[string]any{"$.store.book[0].author": "jon", "$.store.book[*].author": []any{"jon", "ken"}, "$.store.pen": nil, "store.book.1.author": "ken"},
			false,
		},
		{"empty", "", map[string]any{"a": nil}, false},
		{"a syntax error", "a: [b\n", nil, true},
		{"not a map", "- a\n- b\n", nil, true},
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

// TestLookupInvalidPath tests that a YAML path that does not parse is an
// invalid key.
func TestLookupInvalidPath(t *testing.T) {
	cfg, err := Decode([]byte("a: b\n"))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if _, _, err := cfg.Lookup("$.["); !errors.Is(err, ox.ErrInvalidKey) {
		t.Errorf("expected %v, got: %v", ox.ErrInvalidKey, err)
	}
}
