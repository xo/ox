package metadoc

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/xo/ox"
)

// update writes the golden files in testdata from the output of the tests.
var update = flag.Bool("update", false, "write the golden files")

// date is the date of the pages of the tests.
var date = time.Date(2026, 11, 28, 0, 0, 0, 0, time.UTC)

// treeOpts returns the options of the command tree of the tests: a root with
// flags, an inherited flag, a local flag, a hidden flag, a deprecated flag,
// allowed values, a sub command with arguments, an example and a sub command
// of its own, and a hidden and a deprecated command.
func treeOpts(extra ...ox.Option) []ox.Option {
	return append([]ox.Option{
		ox.Usage("app", "copy files between hosts"),
		ox.VersionString("v1.2.3"),
		ox.Version(),
		ox.Help(),
		ox.Banner("App copies files.\n\nIt reads its config from FILE."),
		ox.Example("  app --config x.yaml get a.txt\n  app get -r 5 a.txt /tmp"),
		ox.Footer("See https://example.com/app for more."),
		ox.Flags().
			String("config", "the config file", ox.Short("c"), ox.Spec("FILE"), ox.Default("$HOME/.apprc")).
			Bool("debug", "write debug output", ox.Short("d"), ox.Aliases("trace")).
			String("format", "the output format", ox.Valid("json", "yaml"), ox.Default("json")).
			String("token", "the token", ox.Local()).
			String("secret", "a hidden flag", ox.Hidden(true)).
			String("old-flag", "a deprecated flag", ox.Deprecated(true)),
		ox.Sub(
			ox.Usage("get", "get a file"),
			ox.Arg("path", "the path of the file"),
			ox.Arg("dest", "where to write the file", ox.Default(".")),
			ox.Flags().
				Int("retries", "the number of retries", ox.Short("r"), ox.Default(3)).
				Bool("force", "overwrite a file", ox.Required()),
			ox.Example("  app get a.txt"),
			ox.Sub(
				ox.Usage("all", "get each file"),
				ox.Flags().
					Bool("dry-run", "write what it gets, and get nothing"),
				ox.Exec(func() {}),
			),
			ox.Exec(func() {}),
		),
		ox.Sub(
			ox.Usage("put", "put a file"),
			ox.Exec(func() {}),
		),
		ox.Sub(
			ox.Usage("hidden", "a hidden command"),
			ox.Hidden(true),
			ox.Exec(func() {}),
		),
		ox.Sub(
			ox.Usage("old", "a deprecated command"),
			ox.Deprecated(true),
			ox.Exec(func() {}),
		),
	}, extra...)
}

// newTree builds the command tree of the tests.
func newTree(t *testing.T, extra ...ox.Option) *ox.Command {
	t.Helper()
	root, err := ox.NewCommand(treeOpts(extra...)...)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return root
}

// golden compares buf with the file in testdata, or writes the file when
// -update is set.
func golden(t *testing.T, name string, buf []byte) {
	t.Helper()
	path := filepath.Join("testdata", name)
	if *update {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if err := os.WriteFile(path, buf, 0o644); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		return
	}
	exp, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if s := string(buf); s != string(exp) {
		t.Errorf("%s: expected:\n%s\ngot:\n%s", name, exp, s)
	}
}

func TestWrite(t *testing.T) {
	tests := []struct {
		name   string
		format Format
		path   []string
		tree   bool
		file   string
	}{
		{"man tree", ManPage(Date(date)), nil, true, "tree.1"},
		{"markdown tree", Markdown(), nil, true, "tree.md"},
		{"man root", ManPage(Date(date)), nil, false, "app.1"},
		{"man sub", ManPage(Date(date)), []string{"get"}, false, "app-get.1"},
		{"man deep", ManPage(Date(date), Section(8), Source("App 2"), Manual("App Manual")), []string{"get", "all"}, false, "app-get-all.8"},
		{"markdown root", Markdown(), nil, false, "app.md"},
		{"markdown sub", Markdown(), []string{"get"}, false, "app-get.md"},
		{"markdown deep", Markdown(), []string{"get", "all"}, false, "app-get-all.md"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := newTree(t).Lookup(test.path...)
			var buf bytes.Buffer
			var err error
			if test.tree {
				err = Write(&buf, cmd, test.format)
			} else {
				err = WriteCommand(&buf, cmd, test.format)
			}
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			golden(t, test.file, buf.Bytes())
		})
	}
}

func TestWriteDir(t *testing.T) {
	tests := []struct {
		name   string
		format Format
		files  []string
	}{
		{"man", ManPage(Date(date)), []string{"app-get-all.1", "app-get.1", "app-help.1", "app-put.1", "app-version.1", "app.1"}},
		{"markdown", Markdown(), []string{"app-get-all.md", "app-get.md", "app-help.md", "app-put.md", "app-version.md", "app.md"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dir := filepath.Join(t.TempDir(), "a", "b")
			if err := WriteDir(dir, newTree(t), test.format); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			entries, err := os.ReadDir(dir)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			files := make([]string, 0, len(entries))
			for _, e := range entries {
				files = append(files, e.Name())
				info, err := e.Info()
				if err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
				// the umask can take bits away, and never adds one
				if mode := info.Mode().Perm(); mode&^0o644 != 0 {
					t.Errorf("expected mode 0644 or less, got: %o", mode)
				}
				buf, err := os.ReadFile(filepath.Join(dir, e.Name()))
				if err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
				// each file is the same as the page of its command
				golden(t, e.Name(), buf)
			}
			if !slices.Equal(files, test.files) {
				t.Errorf("expected %v, got: %v", test.files, files)
			}
		})
	}
}

func TestWriteDirName(t *testing.T) {
	root, err := ox.NewCommand(
		ox.Usage("app", "an app"),
		ox.Sub(ox.Usage("../x", "a bad name")),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	dir := t.TempDir()
	err = WriteDir(dir, root, Markdown())
	if !errors.Is(err, ox.ErrInvalidValue) {
		t.Fatalf("expected ErrInvalidValue, got: %v", err)
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 0 {
		t.Errorf("expected no file, got: %d", len(entries))
	}
}

func TestDate(t *testing.T) {
	today := time.Now().UTC().Format(time.DateOnly)
	tests := []struct {
		name  string
		opts  []ManOption
		epoch string
		exp   string
		err   bool
	}{
		{"date wins", []ManOption{Date(date)}, "0", "2026-11-28", false},
		{"source date epoch", nil, "1700000000", "2023-11-14", false},
		{"source date epoch is utc", nil, "1700006399", "2023-11-14", false},
		{"empty source date epoch", nil, "", today, false},
		{"bad source date epoch", nil, "yesterday", "", true},
		{"bad source date epoch with date", []ManOption{Date(date)}, "yesterday", "2026-11-28", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("SOURCE_DATE_EPOCH", test.epoch)
			root, err := ox.NewCommand(ox.Usage("app", "an app"))
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			var buf bytes.Buffer
			err = Write(&buf, root, ManPage(test.opts...))
			switch {
			case test.err && !errors.Is(err, ox.ErrInvalidValue):
				t.Fatalf("expected ErrInvalidValue, got: %v", err)
			case test.err:
				return
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			}
			exp := `.TH "APP" "1" "` + test.exp + `" "app"` + "\n"
			if s := buf.String(); !strings.HasPrefix(s, exp) {
				t.Errorf("expected prefix %q, got: %q", exp, s)
			}
		})
	}
}

func TestRoff(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{`a\b`, `a\eb`},
		{`--dir`, `\-\-dir`},
		{".TH x", `\&.TH x`},
		{"'x", `\&'x`},
		{"a\n.b\n'c\n d", "a\n\\&.b\n\\&'c\n d"},
		{"a.b 'c'", "a.b 'c'"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if s := roff(test.s); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

func TestRoffArg(t *testing.T) {
	if s, exp := roffArg("a \"b\"\nc"), `"a \(dqb\(dq c"`; s != exp {
		t.Errorf("expected %q, got: %q", exp, s)
	}
}

func TestMarkdownEscape(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"a <b> & c", `a \<b\> \& c`},
		{"*a* _b_ `c`", "\\*a\\* \\_b\\_ \\`c\\`"},
		{"a|b [c]", `a\|b \[c\]`},
		{`a\b`, `a\\b`},
		{"# a", `\# a`},
		{"- a\n+ b\n= c", "\\- a\n\\+ b\n\\= c"},
		{"1. a\n  2) b", "1\\. a\n  2\\) b"},
		{"a-b 1.2", "a-b 1.2"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if s := mdEscape(test.s); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

func TestCodeSpan(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"--dir DIR", "`--dir DIR`"},
		{"a`b", "``a`b``"},
		{"`a``", "``` `a`` ```"},
		{"a\nb", "`a b`"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if s := codeSpan(test.s); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

func TestEscapeInPages(t *testing.T) {
	root, err := ox.NewCommand(
		ox.Usage("app", "an app"),
		ox.Banner(".TH not a request\n'not a request\na\\b <b> *c* _d_ `e` |f|"),
		ox.Flags().
			String("my-flag", "a flag with a-hyphen", ox.Spec("A|B")),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		format Format
		exp    []string
	}{
		{ManPage(Date(date)), []string{
			"\\&.TH not a request\n\\&'not a request\na\\eb <b> *c* _d_ `e` |f|\n",
			`\fB\-\-my\-flag\fR A|B` + "\na flag with a\\-hyphen\n",
		}},
		{Markdown(), []string{
			".TH not a request\n'not a request\na\\\\b \\<b\\> \\*c\\* \\_d\\_ \\`e\\` \\|f\\|\n",
			"- `--my-flag A|B`: a flag with a-hyphen\n",
		}},
	}
	for _, test := range tests {
		t.Run(test.format.name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := Write(&buf, root, test.format); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			for _, exp := range test.exp {
				if s := buf.String(); !strings.Contains(s, exp) {
					t.Errorf("expected %q in:\n%s", exp, s)
				}
			}
		})
	}
}

func TestFrontMatterAndLink(t *testing.T) {
	format := Markdown(
		FrontMatter(func(cmd *ox.Command) string {
			return "---\ntitle: " + strings.Join(cmd.Tree(), " ") + "\n---"
		}),
		Link(func(cmd *ox.Command) string {
			return "/docs/" + strings.Join(cmd.Tree(), "_") + "/"
		}),
	)
	var buf bytes.Buffer
	if err := WriteCommand(&buf, newTree(t).Lookup("get"), format); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	s := buf.String()
	for _, exp := range []string{
		"---\ntitle: app get\n---\n\n# app get\n",
		"- [app get all](/docs/app_get_all/): get each file\n",
		"## See Also\n\n- [app](/docs/app/): copy files between hosts\n",
	} {
		if !strings.Contains(s, exp) {
			t.Errorf("expected %q in:\n%s", exp, s)
		}
	}
}

func TestCommand(t *testing.T) {
	tests := []struct {
		name string
		opts []ox.Option
		err  error
	}{
		{"both formats", []ox.Option{Command()}, nil},
		{"one format", []ox.Option{Command(ManPage())}, nil},
		{"before the sub commands", []ox.Option{Command(), ox.Sub(ox.Usage("last", "the last command"))}, nil},
		{"name collision", []ox.Option{ox.Sub(ox.Usage("__docs", "a command")), Command()}, ox.ErrExist},
		{"name collision after", []ox.Option{Command(), ox.Sub(ox.Usage("__docs", "a command"))}, ox.ErrExist},
		{"format given twice", []ox.Option{Command(ManPage(), ManPage())}, ox.ErrInvalidValue},
		{"empty format", []ox.Option{Command(Format{})}, ox.ErrInvalidValue},
		{"not the root", []ox.Option{ox.Sub(ox.Usage("sub", "a sub command"), Command())}, ox.ErrCanOnlyBeUsedWithRootCommand},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ox.NewCommand(append([]ox.Option{ox.Usage("app", "an app")}, test.opts...)...)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.err):
				t.Fatalf("expected %v, got: %v", test.err, err)
			}
		})
	}
}

// errUnknown is the start of the error of an unknown command, which does not
// wrap ox.ErrUnknownCommand.
var errUnknown = errors.New(`unknown command "html"`)

func TestRun(t *testing.T) {
	t.Setenv("SOURCE_DATE_EPOCH", "1795824000")
	dir := t.TempDir()
	tests := []struct {
		name string
		args []string
		exp  string
		err  error
	}{
		{"man", []string{"__docs", "man"}, "tree.1", nil},
		{"markdown", []string{"__docs", "markdown"}, "tree.md", nil},
		{"man dir", []string{"__docs", "man", "--dir", filepath.Join(dir, "man")}, "", nil},
		{"markdown dir", []string{"__docs", "markdown", "--dir=" + filepath.Join(dir, "md")}, "", nil},
		{"unknown format", []string{"__docs", "html"}, "", errUnknown},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			err := ox.RunError(context.Background(), treeOpts(
				Command(),
				ox.Args(test.args...),
				ox.Pipe(nil, &stdout, &stderr),
			)...)
			switch {
			case test.err != nil && (err == nil || !strings.HasPrefix(err.Error(), test.err.Error())):
				t.Fatalf("expected %v, got: %v", test.err, err)
			case test.err != nil:
				return
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case test.exp != "":
				golden(t, test.exp, stdout.Bytes())
			case stdout.Len() != 0:
				t.Errorf("expected no output, got: %q", stdout.String())
			}
		})
	}
	for _, name := range []string{"man/app-get-all.1", "md/app-get.md"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	}
}

// TestHidden makes sure that the hidden command changes neither help nor
// completion, and that the docs leave it out.
func TestHidden(t *testing.T) {
	// run returns the output and the error of a run
	run := func(opts ...ox.Option) string {
		t.Helper()
		var stdout, stderr bytes.Buffer
		if err := ox.RunError(context.Background(), append(opts, ox.Pipe(nil, &stdout, &stderr))...); err != nil {
			return stdout.String() + "error: " + err.Error()
		}
		return stdout.String()
	}
	trees := []struct {
		name string
		opts func() []ox.Option
	}{
		{"with sub commands", func() []ox.Option {
			return []ox.Option{
				ox.Usage("app", "an app"),
				ox.Defaults(),
				ox.Sub(ox.Usage("sub", "a sub command"), ox.Exec(func() {})),
			}
		}},
		{"without sub commands", func() []ox.Option {
			return []ox.Option{
				ox.Usage("app", "an app"),
				ox.Defaults(),
				ox.Flags().String("name", "a name"),
				ox.Exec(func() {}),
			}
		}},
		{"without exec", func() []ox.Option {
			return []ox.Option{
				ox.Usage("app", "an app"),
				ox.Defaults(),
			}
		}},
	}
	for _, tree := range trees {
		for _, args := range [][]string{
			{"--help"},
			{"__complete", ""},
			{"__complete", "__"},
			{"__complete", "--"},
			{"help"},
			{"__doc"},
		} {
			t.Run(tree.name+" "+strings.Join(args, " "), func(t *testing.T) {
				exp := run(append(tree.opts(), ox.Args(args...))...)
				s := run(append(tree.opts(), Command(), ox.Args(args...))...)
				if s != exp {
					t.Errorf("expected:\n%s\ngot:\n%s", exp, s)
				}
				if strings.Contains(s, DefaultName) {
					t.Errorf("expected no %s in:\n%s", DefaultName, s)
				}
			})
		}
	}
	// the docs keep help, version and completion, and leave out __docs
	for _, name := range []string{"man", "markdown"} {
		s := run(append(trees[0].opts(), Command(), ox.Args(DefaultName, name))...)
		for _, exp := range []string{"app help", "app version", "app completion", "app completion bash"} {
			if !strings.Contains(s, exp) {
				t.Errorf("%s: expected %q in:\n%s", name, exp, s)
			}
		}
		if strings.Contains(s, "docs") || strings.Contains(s, "__complete") {
			t.Errorf("%s: expected no docs and no __complete in:\n%s", name, s)
		}
	}
}

// TestBeforeDoesNotRun makes sure that the funcs of ox.Before and ox.After
// do not run for the hidden command, as for help.
func TestBeforeDoesNotRun(t *testing.T) {
	var ran bool
	var stdout bytes.Buffer
	err := ox.RunError(context.Background(),
		ox.Usage("app", "an app"),
		ox.Before(func() error {
			ran = true
			return errors.New("before")
		}),
		ox.Flags().String("token", "a token", ox.Required()),
		ox.Exec(func() {}),
		Command(Markdown()),
		ox.Args(DefaultName, "markdown"),
		ox.Pipe(nil, &stdout, &stdout),
	)
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case ran:
		t.Errorf("expected the before func not to run")
	case !strings.HasPrefix(stdout.String(), "# app\n"):
		t.Errorf("expected the docs, got: %q", stdout.String())
	}
}

// TestNegatedFlags tests that a bool flag with a negated form shows it as
// help does, such as --[no-]color, in a sub command too (D47).
func TestNegatedFlags(t *testing.T) {
	root, err := ox.NewCommand(
		ox.Usage("app", "an app"),
		ox.NegateBools(),
		ox.Flags().
			Bool("color", "use color").
			Bool("no-rc", "do not read the startup file").
			Bool("force", "overwrite each file", ox.Negatable(false)),
		ox.Sub(
			ox.Usage("sub", "a sub command"),
			ox.Exec(func() {}),
		),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		format Format
		exp    []string
		not    []string
	}{
		{ManPage(Date(date)), []string{`\fB\-\-[no\-]color\fR`, `\fB\-\-no\-rc\fR`, `\fB\-\-force\fR`}, []string{`[no\-]no`, `[no\-]force`}},
		{Markdown(), []string{"`--[no-]color`", "`--no-rc`", "`--force`"}, []string{"[no-]no", "[no-]force"}},
	}
	for _, test := range tests {
		t.Run(test.format.name, func(t *testing.T) {
			for _, cmd := range []*ox.Command{root, root.Command("sub")} {
				var buf bytes.Buffer
				if err := WriteCommand(&buf, cmd, test.format); err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
				s := buf.String()
				for _, exp := range test.exp {
					if !strings.Contains(s, exp) {
						t.Errorf("%s: expected %q in:\n%s", cmd.Name, exp, s)
					}
				}
				for _, not := range test.not {
					if strings.Contains(s, not) {
						t.Errorf("%s: expected no %q in:\n%s", cmd.Name, not, s)
					}
				}
			}
		})
	}
}
