package main

import (
	"context"
	"path/filepath"
	"slices"
	"testing"
)

func TestParseOptionLine(t *testing.T) {
	tests := []struct {
		line     string
		spaceArg bool
		longs    []string
		shorts   []string
		spec     string
		optional bool
		note     string
		desc     string
	}{
		{"  -c, --command=COMMAND    run only single command", false, []string{"command"}, []string{"c"}, "COMMAND", false, "", "run only single command"},
		{"  -v, --set=, --variable=NAME=VALUE", false, []string{"set", "variable"}, []string{"v"}, "NAME=VALUE", false, "", ""},
		{`  -1 ("one"), --single-transaction`, false, []string{"single-transaction"}, []string{"1"}, "", false, `("one")`, ""},
		{"  -?, --help[=options]     show this help, then exit", false, []string{"help"}, []string{"?"}, "options", true, "", "show this help, then exit"},
		{"      --help=commands      list backslash commands, then exit", false, []string{"help"}, nil, "commands", false, "", "list backslash commands, then exit"},
		{"  -z, --field-separator-zero", false, []string{"field-separator-zero"}, []string{"z"}, "", false, "", ""},
		{"  -0, --record-separator-zero", false, []string{"record-separator-zero"}, []string{"0"}, "", false, "", ""},
		{"  -R, -r, --recursive", false, []string{"recursive"}, []string{"R", "r"}, "", false, "", ""},
		{"  -A, --catenate, --concatenate   append tar files to an archive", false, []string{"catenate", "concatenate"}, []string{"A"}, "", false, "", "append tar files to an archive"},
		{"      --exclude-caches-under exclude everything under directories containing", false, []string{"exclude-caches-under"}, nil, "", false, "", "exclude everything under directories containing"},
		{"  -H", false, nil, []string{"H"}, "", false, "", ""},
		{"      --backup[=CONTROL]", false, []string{"backup"}, nil, "CONTROL", true, "", ""},
		{"  -R, --repo [HOST/]OWNER/REPO   Select another repository", true, []string{"repo"}, []string{"R"}, "[HOST/]OWNER/REPO", false, "", "Select another repository"},
		{"  -d, --draft             Filter by draft state", true, []string{"draft"}, []string{"d"}, "", false, "", "Filter by draft state"},
		{`      --color string[="auto"]   use color`, true, []string{"color"}, nil, `string[="auto"]`, false, "", "use color"},
	}
	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			opt, ok := parseOptionLine(test.line, test.spaceArg)
			if !ok {
				t.Fatalf("expected an option line")
			}
			switch {
			case !slices.Equal(opt.longs, test.longs):
				t.Errorf("expected longs %q, got: %q", test.longs, opt.longs)
			case !slices.Equal(opt.shorts, test.shorts):
				t.Errorf("expected shorts %q, got: %q", test.shorts, opt.shorts)
			case opt.spec != test.spec:
				t.Errorf("expected spec %q, got: %q", test.spec, opt.spec)
			case opt.optional != test.optional:
				t.Errorf("expected optional %t, got: %t", test.optional, opt.optional)
			case opt.note != test.note:
				t.Errorf("expected note %q, got: %q", test.note, opt.note)
			case opt.desc != test.desc:
				t.Errorf("expected desc %q, got: %q", test.desc, opt.desc)
			}
		})
	}
}

func TestParseOptionLineNot(t *testing.T) {
	for _, line := range []string{
		"",
		"Usage: du [OPTION]... [FILE]...",
		"                           (e.g., -v ON_ERROR_STOP=1)",
		"         -T option; NUMBER defaults to 1",
		"  --",
		"  - a list item",
	} {
		t.Run(line, func(t *testing.T) {
			if _, ok := parseOptionLine(line, false); ok {
				t.Errorf("expected no option line")
			}
		})
	}
}

// importTestdata imports a program from its archive in testdata.
func importTestdata(t *testing.T, name string, path ...string) *Command {
	t.Helper()
	a, err := readArchive(filepath.Join("testdata", name+".txtar"))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	imp := &Importer{
		Runner: &replayRunner{archive: a},
		Values: true,
		Path:   path,
		Depth:  0,
		Logf:   t.Logf,
	}
	root, err := imp.Import(context.Background(), name)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return root
}

// testFlag is what a test expects of a flag.
type testFlag struct {
	name    string
	short   string
	typ     string
	spec    string
	def     string
	section int
	aliases []string
}

// checkFlags checks the flags of cmd.
func checkFlags(t *testing.T, cmd *Command, flags []testFlag) {
	t.Helper()
	for _, test := range flags {
		g := cmd.Flag(test.name)
		switch {
		case g == nil:
			t.Errorf("%s: expected flag --%s", cmd.Name, test.name)
		case g.Name != test.name:
			t.Errorf("%s: expected name %q, got: %q", cmd.Name, test.name, g.Name)
		case g.Short != test.short:
			t.Errorf("--%s: expected short %q, got: %q", test.name, test.short, g.Short)
		case string(g.Type) != test.typ:
			t.Errorf("--%s: expected type %q, got: %q", test.name, test.typ, g.Type)
		case g.Spec != test.spec:
			t.Errorf("--%s: expected spec %q, got: %q", test.name, test.spec, g.Spec)
		case g.Default != test.def:
			t.Errorf("--%s: expected default %q, got: %q", test.name, test.def, g.Default)
		case g.Section != test.section:
			t.Errorf("--%s: expected section %d, got: %d", test.name, test.section, g.Section)
		case !slices.Equal(g.Aliases, test.aliases):
			t.Errorf("--%s: expected aliases %q, got: %q", test.name, test.aliases, g.Aliases)
		}
	}
}
