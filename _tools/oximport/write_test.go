package main

import (
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/xo/ox"
)

func TestParseComplete(t *testing.T) {
	tests := []struct {
		out       string
		ok        bool
		names     []string
		directive int
	}{
		{"open\nclosed\n:4\n", true, []string{"open", "closed"}, 4},
		{"--help\tShow help\n-h\tShow help\n:4\nCompletion ended with directive: ShellCompDirectiveNoFileComp\n", false, nil, 0},
		{"--help\tShow help\n-h\tShow help\n:4\n", true, []string{"--help", "-h"}, 4},
		{":0\n", true, nil, 0},
		{"_activeHelp_ a line\nyaml\n:8\n", true, []string{"yaml"}, 8},
		{"", false, nil, 0},
		{"psql: error: connection to server failed\n", false, nil, 0},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			comp, ok := parseComplete([]byte(test.out))
			switch {
			case ok != test.ok:
				t.Fatalf("expected ok %t, got: %t", test.ok, ok)
			case !ok:
			case !slices.Equal(comp.names, test.names):
				t.Errorf("expected names %q, got: %q", test.names, comp.names)
			case comp.directive != test.directive:
				t.Errorf("expected directive %d, got: %d", test.directive, comp.directive)
			}
		})
	}
}

func TestJoinFlagsRequired(t *testing.T) {
	cmd := &Command{Name: "create", Flags: []*Flag{
		{Name: "help", Short: "h", Hook: "help", Type: ox.BoolT},
		{Name: "name", Type: ox.StringT},
		{Name: "size", Type: ox.IntT},
	}}
	comp, _ := parseComplete([]byte("--name\tthe name\n-n\tthe name\n:4\n"))
	joinFlags(cmd, comp)
	if g := cmd.Flag("name"); !g.Required || g.Short != "n" {
		t.Errorf("expected --name to be required with -n, got: %t %q", g.Required, g.Short)
	}
	if g := cmd.Flag("size"); g.Required {
		t.Errorf("expected --size not to be required")
	}
}

func TestArchive(t *testing.T) {
	a := newArchive("a comment\nVersion: 1.0")
	a.set("--help", []byte("Usage: prog\n\n  -a  all\n\n"))
	a.set(argsKey([]string{"__complete", "pr", ""}), []byte("x\ty\n:4\n"))
	a.set("--version", nil)
	name := filepath.Join(t.TempDir(), "prog.txtar")
	if err := a.write(name); err != nil {
		t.Fatal(err)
	}
	b, err := readArchive(name)
	if err != nil {
		t.Fatal(err)
	}
	if b.comment != a.comment {
		t.Errorf("expected comment %q, got: %q", a.comment, b.comment)
	}
	if !slices.Equal(a.keys, b.keys) {
		t.Errorf("expected keys %q, got: %q", a.keys, b.keys)
	}
	for _, key := range a.keys {
		if string(a.outputs[key]) != string(b.outputs[key]) {
			t.Errorf("%s: expected %q, got: %q", key, a.outputs[key], b.outputs[key])
		}
	}
	if exp := `__complete pr ""`; a.keys[1] != exp {
		t.Errorf("expected key %q, got: %q", exp, a.keys[1])
	}
}

// TestTestdata checks that no archive holds the name of the user who
// recorded it.
func TestTestdata(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil || home == "/" {
		t.Skip("no home directory")
	}
	files, err := filepath.Glob(filepath.Join("testdata", "*.txtar"))
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range files {
		buf, err := os.ReadFile(name)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(buf), home) {
			t.Errorf("%s: expected no %s", name, home)
		}
	}
}

// TestStructTag checks that the description and the options of a tag reach
// ox unchanged. See "Escaping in a struct tag" in docs/DESIGN.md.
func TestStructTag(t *testing.T) {
	tests := []string{
		"read a file, then exit",
		"host:port to listen on",
		`say "hi", loudly`,
		`a back\slash`,
		"two\nlines",
		"a `quoted` word, here",
	}
	for _, s := range tests {
		t.Run(s, func(t *testing.T) {
			tag := structTag([]string{s, "default:http://x:80/a,b"})
			lit, err := strconv.Unquote(tag)
			if err != nil {
				t.Fatalf("expected a Go string, got: %s", tag)
			}
			v, ok := reflect.StructTag(lit).Lookup("ox")
			if !ok {
				t.Fatalf("expected an ox tag in %s", lit)
			}
			parts := ox.SplitBy(v, ',')
			if exp := []string{s, "default:http://x:80/a,b"}; !slices.Equal(parts, exp) {
				t.Errorf("expected %q, got: %q", exp, parts)
			}
		})
	}
}

// TestFlagsFrom checks that ox builds each flag of a reflection style from
// its struct, with the same name.
func TestFlagsFrom(t *testing.T) {
	v := struct {
		SingleTransaction bool              `ox:"execute as a single transaction,short:1"`
		Set               map[string]string `ox:"set psql variable NAME to VALUE\n(e.g.\\, -v ON_ERROR_STOP=1),spec:NAME=VALUE,short:v,alias:variable"`
		H                 bool              `ox:"equivalent to --dereference-args (-D),name:H,short:H"`
		Wait              string            `ox:"time to wait,type:duration,default:5s"`
	}{}
	flags, err := ox.FlagsFrom(&v)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	var names []string
	for _, g := range flags {
		names = append(names, g.Name)
	}
	if exp := []string{"single-transaction", "set", "H", "wait"}; !slices.Equal(names, exp) {
		t.Errorf("expected %q, got: %q", exp, names)
	}
	if exp := "set psql variable NAME to VALUE\n(e.g., -v ON_ERROR_STOP=1)"; flags[1].Usage != exp {
		t.Errorf("expected usage %q, got: %q", exp, flags[1].Usage)
	}
	for _, g := range []*Flag{{Name: "single-transaction"}, {Name: "H"}, {Name: "help-all"}, {Name: "kube-as-group"}} {
		f := fields([]*Flag{g})[0]
		if name := ox.DefaultFlagNameMapper(f.name); name != g.Name && !slices.Contains(f.opts, "name:"+g.Name) {
			t.Errorf("--%s: field %s maps to %s and has no name option", g.Name, f.name, name)
		}
	}
}

func TestGoName(t *testing.T) {
	tests := []struct{ name, exp string }{
		{"single-transaction", "SingleTransaction"},
		{"files0-from", "Files0From"},
		{"0", "Flag0"},
		{"?", "Flag"},
		{"kube_as_group", "KubeAsGroup"},
	}
	for _, test := range tests {
		if got := goName(test.name); got != test.exp {
			t.Errorf("%s: expected %q, got: %q", test.name, test.exp, got)
		}
	}
}
