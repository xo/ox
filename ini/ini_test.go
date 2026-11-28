package ini

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/xo/ox"
)

// TestDecode decodes INI and looks up dotted keys.
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
			"top level keys",
			"port = 8080\nname: db\n  on =  true  \nempty =\n",
			map[string]any{"port": "8080", "name": "db", "on": "true", "empty": "", "other": nil},
			false,
		},
		{
			"comments and blank lines",
			"; a comment\n# a comment\n\n  ; an indented comment\na = b\n",
			map[string]any{"a": "b"},
			false,
		},
		{
			"sections",
			"[server]\nhost = db\n\n[migrate]\nsteps = 3\n\n[migrate.source]\ndir = ./up\n",
			map[string]any{"server.host": "db", "migrate.steps": "3", "migrate.source.dir": "./up", "host": nil, "migrate.source.host": nil},
			false,
		},
		{
			"a section twice",
			"[a]\nb = 1\n[c]\nd = 2\n[a]\ne = 3\n",
			map[string]any{"a.b": "1", "a.e": "3", "c.d": "2"},
			false,
		},
		{
			"spaces in a section",
			"[ a . b ]\nc = d\n",
			map[string]any{"a.b.c": "d"},
			false,
		},
		{
			"a dotted key",
			"[migrate]\nsource.dir = ./up\n",
			map[string]any{"migrate.source.dir": "./up"},
			false,
		},
		{
			"a key twice is a list",
			"tag = a\ntag = b\ntag = c\n",
			map[string]any{"tag": []any{"a", "b", "c"}, "tag.1": "b"},
			false,
		},
		{
			"quotes",
			"a = \" spaces \"\nb = ' single '\nc = \"tab\\there\"\nd = \"\nf = it's\n",
			map[string]any{"a": " spaces ", "b": " single ", "c": "tab\there", "d": "\"", "f": "it's"},
			false,
		},
		{
			"a value with = and :",
			"url = http://host:80/?a=b\ntime: 10:30\n",
			map[string]any{"url": "http://host:80/?a=b", "time": "10:30"},
			false,
		},
		{"empty", "", map[string]any{"a": nil}, false},
		{"no =", "a b\n", nil, true},
		{"no key", "= b\n", nil, true},
		{"an open section", "[a\n", nil, true},
		{"an empty section", "[]\n", nil, true},
		{"an empty part of a section", "[a..b]\n", nil, true},
		{"a bad quote", "a = \"\\q\"\n", nil, true},
		{"a value and a section", "a = 1\n[a]\nb = 2\n", nil, true},
		{"a section and a value", "[a]\nb = 1\n[]\n", nil, true},
		{"a section then a value", "[a.b]\nc = 1\n[a]\nb = 2\n", nil, true},
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

// TestDecodeError tests that an error names its line and wraps
// ox.ErrInvalidConfig.
func TestDecodeError(t *testing.T) {
	_, err := Decode([]byte("a = 1\n\nb\n"))
	if !errors.Is(err, ox.ErrInvalidConfig) {
		t.Fatalf("expected %v, got: %v", ox.ErrInvalidConfig, err)
	}
	if exp := `line 3: invalid config: no = in "b"`; err.Error() != exp {
		t.Errorf("expected %q, got: %q", exp, err.Error())
	}
}
