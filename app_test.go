package ox

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// testApp is the root command of the tests of App.
type testApp struct {
	Verbose bool   `ox:",short:v,section:0"`
	Name    string `ox:"the name from the tag,short:n"`
	Quiet   bool   `ox:""`
	Skip    bool   `ox:"-"`
	out     *bytes.Buffer
}

func (app *testApp) Commands() []any {
	return []any{&testAppRunCommand{app: app}, new(testAppListCommand)}
}

func (app *testApp) Exec(_ context.Context, args []string) error {
	fmt.Fprintf(app.out, "root: verbose=%t name=%s args=%v\n", app.Verbose, app.Name, args)
	return nil
}

// testAppRunCommand is a sub command with an Exec func.
type testAppRunCommand struct {
	Count int `ox:",short:c"`
	app   *testApp
}

func (cmd *testAppRunCommand) Exec(args []string) {
	fmt.Fprintf(cmd.app.out, "run: count=%d verbose=%t args=%v\n", cmd.Count, cmd.app.Verbose, args)
}

// testAppListCommand is a sub command with no doc and no Exec func.
type testAppListCommand struct{}

// testAppBad has a section that is not a number.
type testAppBad struct{}

func init() {
	RegisterDoc[testApp](`testApp is the app.

Description:

	the test app

Usage:

	app [flags] [command] [NAME]

Banner:

	The test app.

Sections:

	General
	Other

Command sections:

	Commands

Examples:

	  app run -c 2

Footer:

	See the docs.
`, map[string]string{
		"Verbose": "be verbose,\nwith a comma\n\nThe second paragraph is not usage.\n",
		"Name":    "the doc that the tag replaces\n",
	})
	RegisterDoc[testAppRunCommand](`testAppRunCommand is app run.

Description:

	run the app

Usage:

	app run [N]

Section:

	0

Aliases:

	r
	go
`, map[string]string{
		"Count": "the count\n",
	})
	RegisterDoc[testAppBad]("Section:\n\n\tone\n", nil)
}

func TestApp(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"root", []string{"-v", "-n", "bob", "x"}, "root: verbose=true name=bob args=[x]\n"},
		{"sub", []string{"-v", "run", "-c", "3", "y"}, "run: count=3 verbose=true args=[y]\n"},
		{"sub alias", []string{"r", "--count", "4"}, "run: count=4 verbose=false args=[]\n"},
		{
			"root help", []string{"--help"},
			"The test app.\n\nUsage:\n  app [flags] [command] [NAME]\n\nExamples:\n  app run -c 2\n\nAvailable Commands:\n  help    show help for any command\n  list    \n\nCommands:\n  run     run the app\n\nFlags:\n  -n, --name string  the name from the tag\n  -h, --help         show help, then exit\n\nGeneral:\n  -v, --verbose      be verbose, with a comma\n\nSee the docs.\n",
		},
		{
			"sub help", []string{"run", "--help"},
			"run run the app\n\nUsage:\n  app run [N]\n\nAliases:\n  run, r, go\n\nFlags:\n  -c, --count int  the count\n  -h, --help       show help, then exit\n",
		},
		{"sub without a doc", []string{"list", "--help"}, "list\n\nUsage:\n  app list [flags] [args]\n\nFlags:\n  -h, --help      show help, then exit\n"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			app := &testApp{out: &buf}
			if err := RunError(
				context.Background(),
				App(app),
				Help(),
				Args(test.args...),
				Pipe(nil, &buf, &buf),
			); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s := buf.String(); s != test.exp {
				t.Errorf("expected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", test.exp, s, test.exp, s)
			}
		})
	}
}

func TestAppErrors(t *testing.T) {
	tests := []struct {
		name string
		opt  Option
		exp  error
	}{
		{"not a struct", App(new(int)), ErrInvalidType},
		{"bad section", App(new(testAppBad)), ErrInvalidValue},
		{"bad flag", App(new(struct {
			bad string `ox:"x"`
		})), ErrInvalidType},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := RunError(context.Background(), test.opt, Args())
			if !errors.Is(err, test.exp) {
				t.Errorf("expected error %v, got: %v", test.exp, err)
			}
		})
	}
}

func TestParseDoc(t *testing.T) {
	tests := []struct {
		name string
		doc  string
		exp  map[string]string
	}{
		{"empty", "", map[string]string{}},
		{"no section", "App is an app.\n", map[string]string{}},
		{"sections", "App is an app.\n\nUsage:\n\n\tapp [flags]\n\nFooter:\n\n\tline one\n\n\tline three\n\nUnknown:\n", map[string]string{"Usage": "app [flags]", "Footer": "line one\n\nline three\n\nUnknown:"}},
		{"text that is not indented", "Banner:\nthe banner\n", map[string]string{"Banner": "the banner"}},
		{"an empty section", "Banner:\n\nFooter:\n\tx", map[string]string{"Banner": "", "Footer": "x"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := parseDoc(test.doc)
			if len(m) != len(test.exp) {
				t.Fatalf("expected %q, got: %q", test.exp, m)
			}
			for k, v := range test.exp {
				if m[k] != v {
					t.Errorf("%s: expected %q, got: %q", k, v, m[k])
				}
			}
		})
	}
}

// TestAppFrom holds that AppFrom adds the defaults, such as the version
// command. A run that exits stops the test.
func TestAppFrom(t *testing.T) {
	var buf bytes.Buffer
	AppFrom(&testApp{out: &buf}, Args("version"), VersionString("1.2.3"), Pipe(nil, &buf, &buf))
	if s, exp := buf.String(), "app 1.2.3\n"; s != exp {
		t.Errorf("expected %q, got: %q", exp, s)
	}
	if strings.Contains(buf.String(), "root:") {
		t.Errorf("expected no run")
	}
}

func TestIndent(t *testing.T) {
	tests := []struct {
		name string
		s    string
		exp  string
	}{
		{"none", "a\nb", "  a\n  b"},
		{"common", "    a\n      b\n\n    c", "  a\n    b\n\n  c"},
		{"blank lines", "\ta\n \n\tb", "  a\n\n  b"},
		{"tabs and spaces", "\t a\n\tb", "   a\n  b"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if s := indent(test.s, "  "); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}
