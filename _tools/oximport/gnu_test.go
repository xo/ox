package main

import (
	"slices"
	"strings"
	"testing"
)

func TestGNU(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		spec     string
		sections int
		flags    []testFlag
	}{
		{"du", 28, "[OPTION]... [FILE]...\n  or:  du [OPTION]... --files0-from=F", 0, []testFlag{
			{"null", "0", "bool", "", "", -1, nil},
			{"block-size", "B", "string", "SIZE", "", -1, nil},
			{"max-depth", "d", "int", "N", "", -1, nil},
			{"files0-from", "", "string", "F", "", -1, nil},
			{"H", "H", "bool", "", "", -1, nil},
			{"k", "k", "bool", "", "", -1, nil},
			{"time", "", "string", "WORD", "", -1, nil},
			{"help", "", "bool", "", "", -1, nil},
		}},
		{"cp", 36, "[OPTION]... [-T] SOURCE DEST\n  or:  cp [OPTION]... SOURCE... DIRECTORY\n  or:  cp [OPTION]... -t DIRECTORY SOURCE...", 0, []testFlag{
			{"archive", "a", "bool", "", "", -1, nil},
			{"backup", "", "string", "CONTROL", "", -1, nil},
			{"b", "b", "bool", "", "", -1, nil},
			{"recursive", "R", "bool", "", "", -1, []string{"r"}},
			{"target-directory", "t", "string", "DIRECTORY", "", -1, nil},
		}},
		{"tar", 156, "[OPTION...] [FILE]...", 18, []testFlag{
			{"catenate", "A", "bool", "", "", 0, []string{"concatenate"}},
			{"diff", "d", "bool", "", "", 0, []string{"compare"}},
			{"listed-incremental", "g", "string", "FILE", "", 1, nil},
			{"occurrence", "", "string", "NUMBER", "", 1, nil},
			{"exclude-caches-under", "", "bool", "", "", 2, nil},
			{"help", "?", "bool", "", "", 17, nil},
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := importTestdata(t, test.name)
			if root.Format != FormatGNU {
				t.Fatalf("expected format %q, got: %q", FormatGNU, root.Format)
			}
			if n := len(root.Flags); n != test.n {
				t.Errorf("expected %d flags, got: %d", test.n, n)
			}
			if root.Spec != test.spec {
				t.Errorf("expected spec %q, got: %q", test.spec, root.Spec)
			}
			if n := len(root.Sections); n != test.sections {
				t.Errorf("expected %d sections, got: %d %q", test.sections, n, root.Sections)
			}
			if !strings.Contains(root.Footer, "Report bugs to") && test.name != "tar" {
				t.Errorf("expected the footer to have the bug address, got: %q", root.Footer)
			}
			checkFlags(t, root, test.flags)
		})
	}
}

func TestGNUDescriptions(t *testing.T) {
	root := importTestdata(t, "du")
	g := root.Flag("time")
	if g == nil {
		t.Fatal("expected --time")
	}
	exp := "show time of the last modification of any file in the directory,\n" +
		"or any of its subdirectories\n" +
		"show time as WORD instead of modification time:\n" +
		"atime, access, use, ctime or status"
	if g.Usage != exp {
		t.Errorf("expected usage %q, got: %q", exp, g.Usage)
	}
	if !g.Optional {
		t.Errorf("expected --time to take an optional value")
	}
	if !strings.HasPrefix(root.Banner, "Summarize device usage") {
		t.Errorf("expected banner, got: %q", root.Banner)
	}
	tar := importTestdata(t, "tar")
	if exp := []string{"Main operation mode", "Operation modifiers"}; !slices.Equal(tar.Sections[:2], exp) {
		t.Errorf("expected sections %q, got: %q", exp, tar.Sections[:2])
	}
	if g := tar.Flag("exclude-caches-under"); g != nil {
		if exp := "exclude everything under directories containing\nCACHEDIR.TAG"; g.Usage != exp {
			t.Errorf("expected usage %q, got: %q", exp, g.Usage)
		}
	}
	if !strings.HasPrefix(tar.Example, "  tar -cf archive.tar foo bar") {
		t.Errorf("expected example, got: %q", tar.Example)
	}
}
