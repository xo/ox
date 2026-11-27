package main

import (
	"slices"
	"strings"
	"testing"
)

// TestPsql reads the help of psql 18.7, and checks each fault of the old
// generator that B40 lists.
func TestPsql(t *testing.T) {
	root := importTestdata(t, "psql")
	if root.Format != FormatPsql {
		t.Fatalf("expected format %q, got: %q", FormatPsql, root.Format)
	}
	if exp := "psql is the PostgreSQL interactive terminal."; root.Banner != exp {
		t.Errorf("expected banner %q, got: %q", exp, root.Banner)
	}
	if exp := "[OPTION]... [DBNAME [USERNAME]]"; root.Spec != exp {
		t.Errorf("expected spec %q, got: %q", exp, root.Spec)
	}
	exp := []string{"General options", "Input and output options", "Output format options", "Connection options"}
	if !slices.Equal(root.Sections, exp) {
		t.Errorf("expected sections %q, got: %q", exp, root.Sections)
	}
	if exp := "PostgreSQL home page: <https://www.postgresql.org/>"; !strings.HasSuffix(root.Footer, exp) {
		t.Errorf("expected footer to end with %q, got: %q", exp, root.Footer)
	}
	if n := len(root.Flags); n != 35 {
		t.Errorf("expected 35 flags, got: %d", n)
	}
	checkFlags(t, root, []testFlag{
		{"command", "c", "string", "COMMAND", "", 0, nil},
		{"set", "v", "map", "NAME=VALUE", "", 0, []string{"variable"}},
		{"version", "V", "bool", "", "", 0, nil},
		{"single-transaction", "1", "bool", "", "", 0, nil},
		{"help", "?", "bool", "options", "", 0, nil},
		{"field-separator", "F", "string", "STRING", "", 2, nil},
		{"pset", "P", "map", "VAR[=ARG]", "", 2, nil},
		{"field-separator-zero", "z", "bool", "", "", 2, nil},
		{"record-separator-zero", "0", "bool", "", "", 2, nil},
		{"csv", "", "bool", "", "", 2, nil},
		{"password", "W", "bool", "", "", 3, nil},
	})
	if g := root.Flag("set"); g != nil {
		if exp := "set psql variable NAME to VALUE\n(e.g., -v ON_ERROR_STOP=1)"; g.Usage != exp {
			t.Errorf("expected usage %q, got: %q", exp, g.Usage)
		}
	}
	if g := root.Flag("single-transaction"); g != nil {
		if exp := `("one")`; g.Note != exp {
			t.Errorf("expected note %q, got: %q", exp, g.Note)
		}
		if exp := "execute as a single transaction (if non-interactive)"; g.Usage != exp {
			t.Errorf("expected usage %q, got: %q", exp, g.Usage)
		}
	}
	if g := root.Flag("help"); g != nil {
		if exp := []string{"options", "commands", "variables"}; !slices.Equal(g.Values, exp) {
			t.Errorf("expected values %q, got: %q", exp, g.Values)
		}
		if g.Hook != "help" || !g.Optional {
			t.Errorf("expected an optional help hook, got: hook %q optional %t", g.Hook, g.Optional)
		}
	}
	if g := root.Flag("field-separator"); g != nil {
		if exp := `field separator for unaligned output (default: "|")`; g.Usage != exp {
			t.Errorf("expected usage %q, got: %q", exp, g.Usage)
		}
	}
}
