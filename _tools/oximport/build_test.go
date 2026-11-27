package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// buildTests are the programs in testdata, with the path of commands that the
// archive holds.
var buildTests = []struct {
	name string
	path []string
}{
	{"psql", nil},
	{"du", nil},
	{"cp", nil},
	{"tar", nil},
	{"goimports", nil},
	{"gh", []string{"pr", "list"}},
	{"helm", []string{"list"}},
	{"kubectl", []string{"get"}},
}

// TestBuild writes each style of each program in testdata, builds the hand
// and reflection styles, and runs each with --help. It compares that help with
// the help of the program, and logs the number of lines that differ. For
// psql, it logs each line that differs. The importer does not write every
// option of the help of psql, so a difference does not fail the test.
func TestBuild(t *testing.T) {
	if testing.Short() {
		t.Skip("builds programs")
	}
	goBin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("go is not installed")
	}
	dir := newModule(t)
	type program struct {
		root *Command
		bins map[Style]string
	}
	programs := make(map[string]*program)
	for _, test := range buildTests {
		root := importTestdata(t, test.name, test.path...)
		programs[test.name] = &program{root: root, bins: make(map[Style]string)}
		for _, style := range Styles {
			src, err := Write(root, style)
			if err != nil {
				t.Fatalf("%s %s: expected no error, got: %v", test.name, style, err)
			}
			name := filepath.Join(dir, test.name, string(style), "main.go")
			if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(name, src, 0o644); err != nil {
				t.Fatal(err)
			}
		}
	}
	run := func(args ...string) {
		t.Helper()
		cmd := exec.Command(goBin, args...)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(), "GOWORK=off", "GOFLAGS=-mod=mod", "GOPROXY=off")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("go %s: %v\n%s", strings.Join(args, " "), err, out)
		}
	}
	run("generate", "./...")
	run("vet", "./...")
	for _, test := range buildTests {
		for _, style := range Styles {
			bin := filepath.Join(dir, "bin", test.name+"-"+string(style))
			run("build", "-o", bin, "./"+test.name+"/"+string(style))
			programs[test.name].bins[style] = bin
		}
	}
	for _, test := range buildTests {
		t.Run(test.name, func(t *testing.T) {
			p := programs[test.name]
			a, err := readArchive(filepath.Join("testdata", test.name+".txtar"))
			if err != nil {
				t.Fatal(err)
			}
			paths := [][]string{nil}
			for i := range test.path {
				paths = append(paths, test.path[:i+1])
			}
			for _, style := range Styles {
				for _, path := range paths {
					args := append(append([]string{}, path...), "--help")
					want, err := (&replayRunner{archive: a}).Run(context.Background(), args...)
					if err != nil {
						t.Fatal(err)
					}
					cmd := exec.Command(p.bins[style], args...)
					got, err := cmd.Output()
					if err != nil {
						t.Fatalf("%s %s: expected no error, got: %v", style, argsKey(args), err)
					}
					diff := lineDiff(splitLines(string(want)), splitLines(string(got)))
					lost, extra := 0, 0
					for _, line := range diff {
						if strings.HasPrefix(line, "-") {
							lost++
						} else {
							extra++
						}
					}
					t.Logf("%s %s %s: ox writes %d of %d lines the same, and %d other lines", style, test.name, argsKey(args), len(splitLines(string(want)))-lost, len(splitLines(string(want))), extra)
					if test.name == "psql" && style == StyleHand {
						t.Logf("diff of psql --help (-) and the ox program (+):\n%s", strings.Join(diff, "\n"))
					}
					checkHelp(t, findCommand(p.root, path), string(got))
				}
			}
		})
	}
}

// checkHelp checks that the help names each flag of the command that the
// program defines.
func checkHelp(t *testing.T, cmd *Command, help string) {
	t.Helper()
	for _, g := range written(cmd.root(), nil).find(cmd.Path()).Flags {
		if !strings.Contains(help, "--"+g.Name) {
			t.Errorf("%s: expected --%s in the help", strings.Join(cmd.Tree(), " "), g.Name)
		}
		if g.Short != "" && !strings.Contains(help, "-"+g.Short+", --"+g.Name) {
			t.Errorf("%s: expected -%s, --%s in the help", strings.Join(cmd.Tree(), " "), g.Short, g.Name)
		}
	}
}

// find returns the command at path.
func (cmd *Command) find(path []string) *Command {
	c := cmd
	for _, name := range path {
		c = c.Command(name)
	}
	return c
}

// findCommand returns the command at path.
func findCommand(root *Command, path []string) *Command {
	return root.find(path)
}

// newModule creates a module that requires the ox module of this repository.
func newModule(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	repo, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	mod := fmt.Sprintf("module oximporttest\n\ngo 1.27\n\nrequire github.com/xo/ox v0.0.0\n\nreplace github.com/xo/ox => %s\n", repo)
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(mod), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

// lineDiff returns the lines that differ between a and b, with "- " before
// each line of a and "+ " before each line of b.
func lineDiff(a, b []string) []string {
	// the length of the longest common subsequence of a[i:] and b[j:]
	n, m := len(a), len(b)
	lcs := make([][]int, n+1)
	for i := range lcs {
		lcs[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if a[i] == b[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			} else {
				lcs[i][j] = max(lcs[i+1][j], lcs[i][j+1])
			}
		}
	}
	var diff []string
	i, j := 0, 0
	for i < n || j < m {
		switch {
		case i < n && j < m && a[i] == b[j]:
			i, j = i+1, j+1
		case j < m && (i == n || lcs[i][j+1] >= lcs[i+1][j]):
			diff, j = append(diff, "+ "+b[j]), j+1
		default:
			diff, i = append(diff, "- "+a[i]), i+1
		}
	}
	return diff
}

func TestLineDiff(t *testing.T) {
	got := lineDiff([]string{"a", "b", "c"}, []string{"a", "x", "c", "d"})
	exp := []string{"+ x", "- b", "+ d"}
	if !bytes.Equal([]byte(strings.Join(got, "|")), []byte(strings.Join(exp, "|"))) {
		t.Errorf("expected %q, got: %q", exp, got)
	}
}
