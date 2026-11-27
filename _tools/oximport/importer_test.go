package main

import (
	"context"
	"testing"
)

// pluginRunner writes the same help for each command, as a helm plugin does,
// because helm passes the arguments of a plugin to the plugin.
type pluginRunner struct {
	// runs is the number of runs.
	runs int
}

// Run satisfies the [Runner] interface.
func (r *pluginRunner) Run(_ context.Context, args ...string) ([]byte, error) {
	r.runs++
	switch {
	case args[0] == "__complete":
		return []byte("diff\tshow a diff\n:4\n"), nil
	case len(args) == 1:
		return []byte("The package manager.\n\nUsage:\n  helm [command]\n\nAvailable Commands:\n  diff        show a diff\n\nFlags:\n  -h, --help   help for helm\n"), nil
	}
	return []byte("Show a diff.\n\nUsage:\n  diff [flags]\n\nAvailable Commands:\n  diff        show a diff\n\nFlags:\n  -h, --help   help for diff\n"), nil
}

// Home satisfies the [Runner] interface.
func (r *pluginRunner) Home() string {
	return archiveHome
}

func TestImportStopsWhenHelpRepeats(t *testing.T) {
	r := new(pluginRunner)
	imp := &Importer{Runner: r, Depth: -1}
	root, err := imp.Import(context.Background(), "helm")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	c := root.Command("diff")
	if c == nil || c.Stub {
		t.Fatalf("expected helm diff to be read")
	}
	if d := c.Command("diff"); d == nil || !d.Stub {
		t.Errorf("expected the diff command under helm diff to be a stub")
	}
	if r.runs > 20 {
		t.Errorf("expected few runs, got: %d", r.runs)
	}
}

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		help string
		exp  Format
	}{
		{"Usage: du [OPTION]... [FILE]...\n  -a, --all   all\n", FormatGNU},
		{"usage: gofmt [flags] [path ...]\n  -cpuprofile string\n    \twrite cpu profile to this file\n  -d\tdisplay diffs\n", FormatGoFlag},
		{"psql is a terminal.\n\nUsage:\n  psql [OPTION]...\n\nGeneral options:\n  -c, --command=COMMAND    run a command\n", FormatPsql},
		{"List releases.\n\nUsage:\n  helm list [flags]\n\nFlags:\n  -a, --all   all\n", ""},
	}
	for _, test := range tests {
		t.Run(string(test.exp), func(t *testing.T) {
			if got := detectFormat(test.help); got != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, got)
			}
		})
	}
}
