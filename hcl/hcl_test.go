package hcl

import (
	"encoding/json"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		data string
		keys map[string]any
		err  bool
	}{
		{
			"attributes",
			"port = 8080\nname = \"db\"\non = true\ntags = [\"a\", \"b\"]\nrate = 1.5\n",
			map[string]any{"port": json.Number("8080"), "name": "db", "on": true, "tags": []any{"a", "b"}, "tags.1": "b", "rate": json.Number("1.5")},
			false,
		},
		{
			"a block",
			"server {\n  host = \"db\"\n}\n",
			map[string]any{"server.host": "db"},
			false,
		},
		{
			"a block with labels",
			"server \"db\" \"main\" {\n  port = 5432\n}\n",
			map[string]any{"server.db.main.port": json.Number("5432")},
			false,
		},
		{
			"two blocks of a type",
			"server \"a\" {\n  port = 1\n}\nserver \"b\" {\n  port = 2\n}\n",
			map[string]any{"server.a.port": json.Number("1"), "server.b.port": json.Number("2")},
			false,
		},
		{
			"two blocks with the same labels",
			"server {\n  port = 1\n}\nserver {\n  host = \"db\"\n}\n",
			map[string]any{"server.port": json.Number("1"), "server.host": "db"},
			false,
		},
		{
			"an object",
			"server = {\n  port = 1\n}\n",
			map[string]any{"server.port": json.Number("1")},
			false,
		},
		{"an expression", "port = 8000 + 80\n", map[string]any{"port": json.Number("8080")}, false},
		{"a syntax error", "port = \n", nil, true},
		{"a variable", "port = var.port\n", nil, true},
		{"a function", "port = max(1, 2)\n", nil, true},
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
				case !ok:
					t.Fatalf("%s: expected a value", key)
				}
				if a, b := mustJSON(t, v), mustJSON(t, exp); a != b {
					t.Errorf("%s: expected %s, got: %s", key, b, a)
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
