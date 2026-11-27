package ox

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/xo/ox/text"
)

func TestSpecString(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{"hook with spec", []Option{Option(HookT), NoArg(false, nil), Default(func(string) {}), Spec("NAME=VALUE")}, "set NAME=VALUE"},
		{"hook with spec and noarg takes no value", []Option{Option(HookT), NoArg(true, ""), Default(func() {}), Spec("NAME=VALUE")}, "set"},
		{"hook without spec", []Option{Option(HookT), NoArg(true, ""), Default(func() {})}, "set"},
		{"hook without spec or noarg", []Option{Option(HookT), NoArg(false, nil), Default(func(string) {})}, "set"},
		{"bool with spec", []Option{Option(BoolT), Spec("false")}, "set"},
		{"string", nil, "set string"},
		{"string with spec", []Option{Spec("NAME=VALUE")}, "set NAME=VALUE"},
		{"slice", []Option{Option(SliceT), Elem(IntT)}, "set int"},
		{"map", []Option{Option(MapT), MapKey(StringT), Elem(IntT)}, "set string=int"},
		{"optional value", []Option{OptionalValue("options")}, "set[=options]"},
		{"optional value on a hook", []Option{Option(HookT), Default(func(string) {}), OptionalValue("NAME")}, "set[=NAME]"},
		{"optional value with no spec", []Option{OptionalValue("")}, "set"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g, err := NewFlag("set", "usage", test.opts...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s := g.SpecString(); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestHelpLayout writes help with the default layout and with the layout of
// [HelpColumn]. The case "default layout" holds the output of the default
// layout from before HelpColumn, aliases in help and NoDefault existed.
func TestHelpLayout(t *testing.T) {
	flags := func() *FlagSet {
		return Flags().
			String("name", "the name to use\nwhen it runs", Short("n"), Aliases("nom", "N"), Default("bob"), Section(0)).
			Bool("quiet", "run quietly", Short("q"), Aliases("silent"), Section(0)).
			Hook("set", "set a variable", func(string) {}, Short("s"), Section(1)).
			String("text", "a long description that goes on and on until it is longer than the wrap width of ox, which is ninety five", Section(1)).
			Int("count", "the count", Default("3"), Section(1))
	}
	values := func() *FlagSet {
		return Flags().
			String("mode", "the mode", Short("m"), OptionalValue("MODE"), HelpValue("fast", "go fast"), HelpValue("very-slow", "go\nslow"), Section(0)).
			Bool("quiet", "run quietly", Section(0))
	}
	tests := []struct {
		name   string
		spacer string
		opts   []Option
		exp    string
	}{
		{
			"default layout", " ",
			[]Option{flags()},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command with two lines\n\nFlags:\n  -h, --help         show help, then exit\n\nFirst:\n  -n, --name string  the name to use when it runs (default: bob)\n  -q, --quiet        run quietly\n\nSecond:\n  -s, --set          set a variable\n      --text string  a long description that goes on and on until it is longer than the wrap\n                     width of ox, which is ninety five\n      --count int    the count (default: 3)\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"column", " ",
			[]Option{flags(), HelpColumn(24)},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help                  show help for any command\n  sub                   a sub command\n                        with two lines\n\nFlags:\n  -h, --help            show help, then exit\n\nFirst:\n  -n, -N, --name, --nom string\n                        the name to use\n                        when it runs (default: bob)\n  -q, --quiet, --silent\n                        run quietly\n\nSecond:\n  -s, --set             set a variable\n      --text string     a long description that goes on and on until it is longer than the wrap width of ox, which is ninety five\n      --count int       the count (default: 3)\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"column with equals", "=",
			[]Option{flags(), HelpColumn(24)},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help                  show help for any command\n  sub                   a sub command\n                        with two lines\n\nFlags:\n  -h, --help            show help, then exit\n\nFirst:\n  -n, -N, --name=, --nom=string\n                        the name to use\n                        when it runs (default: bob)\n  -q, --quiet, --silent\n                        run quietly\n\nSecond:\n  -s, --set             set a variable\n      --text=string     a long description that goes on and on until it is longer than the wrap width of ox, which is ninety five\n      --count=int       the count (default: 3)\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"column as help option", " ",
			[]Option{flags(), Help(Sections("First", "Second"), HelpColumn(22))},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help                show help for any command\n  sub                 a sub command\n                      with two lines\n\nFlags:\n  -h, --help          show help, then exit\n\nFirst:\n  -n, -N, --name, --nom string\n                      the name to use\n                      when it runs (default: bob)\n  -q, --quiet, --silent\n                      run quietly\n\nSecond:\n  -s, --set           set a variable\n      --text string   a long description that goes on and on until it is longer than the wrap width of ox, which is ninety five\n      --count int     the count (default: 3)\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"no default", " ",
			[]Option{
				Flags().
					String("name", "the name (default: the user)", Default("bob"), NoDefault()).
					Int("count", "the count", Default("3")),
			},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command with two lines\n\nFlags:\n      --name string  the name (default: the user)\n      --count int    the count (default: 3)\n  -h, --help         show help, then exit\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"values", " ",
			[]Option{values()},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command with two lines\n\nFlags:\n  -h, --help            show help, then exit\n\nFirst:\n  -m, --mode[=MODE]     the mode\n      --mode=fast       go fast\n      --mode=very-slow  go slow\n      --quiet           run quietly\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"values in a column", "=",
			[]Option{values(), HelpColumn(24)},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help                  show help for any command\n  sub                   a sub command\n                        with two lines\n\nFlags:\n  -h, --help            show help, then exit\n\nFirst:\n  -m, --mode[=MODE]     the mode\n      --mode=fast       go fast\n      --mode=very-slow  go\n                        slow\n      --quiet           run quietly\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"no default from a tag", " ",
			[]Option{
				From(&struct {
					Name  string `ox:"the name (default: the user),default:bob,nodefault"`
					Count int    `ox:"the count,default:3"`
				}{}),
			},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command with two lines\n\nFlags:\n      --name string  the name (default: the user)\n      --count int    the count (default: 3)\n  -h, --help         show help, then exit\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setSpacer(t, test.spacer)
			opts := []Option{
				Usage("app", "an app"),
				Help(Sections("First", "Second")),
				Sub(Usage("sub", "a sub command\nwith two lines")),
			}
			if s := runHelp(t, append(opts, test.opts...)...); s != test.exp {
				t.Errorf("expected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", test.exp, s, test.exp, s)
			}
		})
	}
}

// TestHelpInASection holds that a program can put --help and --version in a
// section of its own, with its own short names and usage, when it defines
// flags with Special("hook:help") and Special("hook:version").
func TestHelpInASection(t *testing.T) {
	s := runHelp(
		t,
		Usage("app", "an app"),
		Version(),
		Help(Sections("General options", "Other options")),
		Flags().
			String("host", "the host", Short("h"), Spec("HOSTNAME"), Section(1)).
			Bool("version", "output version information, then exit", Short("V"), Special("hook:version"), Section(0)).
			Bool("help", "show this help, then exit", Short("?"), Special("hook:help"), Section(0)),
	)
	exp := "app an app\n\nUsage:\n  app [flags] [args]\n\nGeneral options:\n  -V, --version        output version information, then exit\n  -?, --help           show this help, then exit\n\nOther options:\n  -h, --host HOSTNAME  the host\n"
	if s != exp {
		t.Errorf("expected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", exp, s, exp, s)
	}
	var buf bytes.Buffer
	RunContext(
		context.Background(),
		Usage("app", "an app"),
		VersionString("1.2.3"),
		Version(),
		Help(),
		Flags().
			Bool("version", "output version information, then exit", Short("V"), Special("hook:version")).
			Bool("help", "show this help, then exit", Short("?"), Special("hook:help")),
		Args("-V"),
		Pipe(nil, &buf, &buf),
	)
	if s, exp := buf.String(), "app 1.2.3\n"; s != exp {
		t.Errorf("expected %q, got: %q", exp, s)
	}
}

// TestHelpPsql builds the options of psql 18.7 and compares its help with the
// output of `psql --help` in testdata/psql-help.txt. One line of psql cannot
// match, so the test changes it before it compares:
//
//   - `-1 ("one"), --single-transaction`: the text `("one")` is an exception
//     that only psql has, so ox writes `-1, --single-transaction`.
func TestHelpPsql(t *testing.T) {
	buf, err := os.ReadFile("testdata/psql-help.txt")
	if err != nil {
		t.Fatal(err)
	}
	exp := string(buf)
	for _, r := range []struct{ old, new string }{
		{"  -1 (\"one\"), --single-transaction\n", "  -1, --single-transaction\n"},
	} {
		if !strings.Contains(exp, r.old) {
			t.Fatalf("testdata/psql-help.txt does not hold %q", r.old)
		}
		exp = strings.Replace(exp, r.old, r.new, 1)
	}
	setSpacer(t, "=")
	s := runHelp(
		t,
		Usage("psql", ""),
		Version(),
		Help(
			Banner("psql is the PostgreSQL interactive terminal."),
			Spec("[OPTION]... [DBNAME [USERNAME]]"),
			Sections(
				"General options",
				"Input and output options",
				"Output format options",
				"Connection options",
			),
			Footer(`For more information, type "\?" (for internal commands) or "\help" (for SQL
commands) from within psql, or consult the psql section in the PostgreSQL
documentation.

Report bugs to <pgsql-bugs@lists.postgresql.org>.
PostgreSQL home page: <https://www.postgresql.org/>`),
			HelpColumn(27),
		),
		Flags().
			String("command", "run only single command (SQL or internal) and exit", Short("c"), Spec("COMMAND"), Section(0)).
			String("dbname", "database name to connect to", Short("d"), Spec("DBNAME"), Section(0)).
			String("file", "execute commands from file, then exit", Short("f"), Spec("FILENAME"), Section(0)).
			Bool("list", "list available databases, then exit", Short("l"), Section(0)).
			Hook("set", "set psql variable NAME to VALUE\n(e.g., -v ON_ERROR_STOP=1)", func(string) {}, NoArg(false, nil), Short("v"), Aliases("variable"), Spec("NAME=VALUE"), Section(0)).
			Bool("version", "output version information, then exit", Short("V"), Special("hook:version"), Section(0)).
			Bool("no-psqlrc", "do not read startup file (~/.psqlrc)", Short("X"), Section(0)).
			Bool("single-transaction", "execute as a single transaction (if non-interactive)", Short("1"), Section(0)).
			Bool("help", "show this help, then exit", Short("?"), Special("hook:help"), Section(0),
				OptionalValue("options"),
				HelpValue("commands", "list backslash commands, then exit"),
				HelpValue("variables", "list special variables, then exit"),
			).
			Bool("echo-all", "echo all input from script", Short("a"), Section(1)).
			Bool("echo-errors", "echo failed commands", Short("b"), Section(1)).
			Bool("echo-queries", "echo commands sent to server", Short("e"), Section(1)).
			Bool("echo-hidden", "display queries that internal commands generate", Short("E"), Section(1)).
			String("log-file", "send session log to file", Short("L"), Spec("FILENAME"), Section(1)).
			Bool("no-readline", "disable enhanced command line editing (readline)", Short("n"), Section(1)).
			String("output", "send query results to file (or |pipe)", Short("o"), Spec("FILENAME"), Section(1)).
			Bool("quiet", "run quietly (no messages, only query output)", Short("q"), Section(1)).
			Bool("single-step", "single-step mode (confirm each query)", Short("s"), Section(1)).
			Bool("single-line", "single-line mode (end of line terminates SQL command)", Short("S"), Section(1)).
			Bool("no-align", "unaligned table output mode", Short("A"), Section(2)).
			Bool("csv", "CSV (Comma-Separated Values) table output mode", Section(2)).
			String("field-separator", `field separator for unaligned output (default: "|")`, Short("F"), Spec("STRING"), Default("|"), NoDefault(), Section(2)).
			Bool("html", "HTML table output mode", Short("H"), Section(2)).
			Map("pset", `set printing option VAR to ARG (see \pset command)`, Short("P"), Spec("VAR[=ARG]"), Section(2)).
			String("record-separator", "record separator for unaligned output (default: newline)", Short("R"), Spec("STRING"), Section(2)).
			Bool("tuples-only", "print rows only", Short("t"), Section(2)).
			String("table-attr", "set HTML table tag attributes (e.g., width, border)", Short("T"), Spec("TEXT"), Section(2)).
			Bool("expanded", "turn on expanded table output", Short("x"), Section(2)).
			Bool("field-separator-zero", "set field separator for unaligned output to zero byte", Short("z"), Section(2)).
			Bool("record-separator-zero", "set record separator for unaligned output to zero byte", Short("0"), Section(2)).
			String("host", "database server host or socket directory", Short("h"), Spec("HOSTNAME"), Section(3)).
			String("port", "database server port", Short("p"), Spec("PORT"), Section(3)).
			String("username", "database user name", Short("U"), Spec("USERNAME"), Section(3)).
			Bool("no-password", "never prompt for password", Short("w"), Section(3)).
			Bool("password", "force password prompt (should happen automatically)", Short("W"), Section(3)),
	)
	if s != exp {
		el, sl := strings.Split(exp, "\n"), strings.Split(s, "\n")
		for i := range max(len(el), len(sl)) {
			var a, b string
			if i < len(el) {
				a = el[i]
			}
			if i < len(sl) {
				b = sl[i]
			}
			if a != b {
				t.Errorf("line %d:\nexpected: %q\ngot:      %q", i+1, a, b)
			}
		}
	}
}

// runHelp runs a command with the options and --help, and returns the help.
func runHelp(t *testing.T, opts ...Option) string {
	t.Helper()
	var buf bytes.Buffer
	RunContext(
		context.Background(),
		append(opts, Args("--help"), Pipe(nil, &buf, &buf))...,
	)
	return buf.String()
}

// setSpacer sets [text.FlagSpecSpacer] for the test, and sets it back after
// the test.
func setSpacer(t *testing.T, spacer string) {
	t.Helper()
	prev := text.FlagSpecSpacer
	text.FlagSpecSpacer = spacer
	t.Cleanup(func() {
		text.FlagSpecSpacer = prev
	})
}

// TestOptionalValue holds that a flag with [OptionalValue] takes a value only
// after an =, and takes the empty string without one.
func TestOptionalValue(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"no value", []string{"--mode"}, "mode=\nargs=[]"},
		{"a value", []string{"--mode=fast"}, "mode=fast\nargs=[]"},
		{"a value after a space is an argument", []string{"--mode", "fast"}, "mode=\nargs=[fast]"},
		{"short", []string{"-m"}, "mode=\nargs=[]"},
		{"short with a value", []string{"-m=fast"}, "mode=fast\nargs=[]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			RunContext(
				context.Background(),
				Usage("app", "an app"),
				Flags().String("mode", "the mode", Short("m"), OptionalValue("MODE")),
				Exec(func(ctx context.Context, args []string) {
					c, _ := Ctx(ctx)
					fmt.Fprintf(&buf, "mode=%s\nargs=%v", c.Vars["mode"], args)
				}),
				Args(test.args...),
				Pipe(nil, &buf, &buf),
			)
			if s := buf.String(); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestUsageString holds that [Command.UsageString] returns the help that
// --help writes, and that [ArgUsage] writes the section of the positional
// arguments in each layout.
func TestUsageString(t *testing.T) {
	flags := func() *FlagSet {
		return Flags().
			String("name", "the name", Short("n"), Default("bob")).
			Bool("quiet", "run quietly", Short("q"))
	}
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{
			"no args",
			[]Option{flags()},
			"app an app\n\nUsage:\n  app [flags] [args]\n\nFlags:\n  -n, --name string  the name (default: bob)\n  -q, --quiet        run quietly\n  -h, --help         show help, then exit\n",
		},
		{
			"args",
			[]Option{flags(), Spec("[DSN] [FILE]"), ArgUsage("DSN", "a URL, or the name of a connection"), ArgUsage("FILE", "a file of commands")},
			"app an app\n\nUsage:\n  app [DSN] [FILE]\n\nArguments:\n  DSN     a URL, or the name of a connection\n  FILE    a file of commands\n\nFlags:\n  -n, --name string  the name (default: bob)\n  -q, --quiet        run quietly\n  -h, --help         show help, then exit\n",
		},
		{
			"args as help options",
			[]Option{flags(), Help(Spec("[DSN]"), ArgUsage("DSN", "a URL,\nor the name of a connection"))},
			"app an app\n\nUsage:\n  app [DSN]\n\nArguments:\n  DSN     a URL, or the name of a connection\n\nFlags:\n  -n, --name string  the name (default: bob)\n  -q, --quiet        run quietly\n  -h, --help         show help, then exit\n",
		},
		{
			"args in a column",
			[]Option{flags(), Help(Spec("[DSN]"), ArgUsage("DSN", "a URL,\nor the name of a connection"), HelpColumn(20))},
			"app an app\n\nUsage:\n  app [DSN]\n\nArguments:\n  DSN               a URL,\n                    or the name of a connection\n\nFlags:\n  -n, --name string\n                    the name (default: bob)\n  -q, --quiet       run quietly\n  -h, --help        show help, then exit\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := append([]Option{Usage("app", "an app"), Help()}, test.opts...)
			if s := runHelp(t, opts...); s != test.exp {
				t.Errorf("--help\nexpected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", test.exp, s, test.exp, s)
			}
			ctx, err := NewContext(opts...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s := ctx.Root.UsageString(); s != test.exp {
				t.Errorf("UsageString\nexpected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", test.exp, s, test.exp, s)
			}
		})
	}
}

// TestHelpArgsSpec holds that the usage line follows [ValidArgs] (B11).
func TestHelpArgsSpec(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{"no valid args", nil, "app [flags] [args]"},
		{"no args", []Option{ValidArgs(0, 0)}, "app [flags]"},
		{"one or two args", []Option{ValidArgs(1, 2)}, "app [flags] [args]"},
		{"values", []Option{ValidArgs(0, 2, "bash", "zsh")}, "app [flags] [bash|zsh]..."},
		{"one value", []Option{ValidArgs(1, 1, "a", "b")}, "app [flags] [a|b]"},
		{"no maximum", []Option{ValidArgs(1, 0, "a")}, "app [flags] [a]..."},
		{"a spec wins", []Option{ValidArgs(0, 0), Spec("[FILE]")}, "app [FILE]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := runHelp(t, append([]Option{Usage("app", "an app"), Help()}, test.opts...)...)
			_, after, _ := strings.Cut(s, "Usage:\n  ")
			if line, _, _ := strings.Cut(after, "\n"); line != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, line)
			}
		})
	}
}

// TestHelpDefault holds that help writes a default that is not a string
// (B46).
func TestHelpDefault(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{"string", []Option{Default("x")}, "(default: x)"},
		{"int", []Option{Option(IntT), Default(3)}, "(default: 3)"},
		{"uint", []Option{Option(UintT), Default(uint(5432))}, "(default: 5432)"},
		{"float", []Option{Option(Float64T), Default(1.5)}, "(default: 1.5)"},
		{"duration", []Option{Option(DurationT), Default(5 * time.Second)}, "(default: 5s)"},
		{"size", []Option{Option(SizeT), Default(Size(2 * MiB))}, "(default: 2 MiB)"},
		{"slice", []Option{Option(SliceT), Default([]string{"a", "b"})}, "(default: [a b])"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := runHelp(t, Usage("app", "an app"), Help(), Flags().Var("v", "a value", test.opts...))
			if !strings.Contains(s, "a value "+test.exp+"\n") {
				t.Errorf("expected %q in:\n%s", test.exp, s)
			}
		})
	}
}

// TestFooterAndVersion holds two faults that B50 fixed: an empty footer in
// Help hides the footer and keeps the spec, and Version hooks a hook:version
// flag on a root that has sub commands.
func TestFooterAndVersion(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		args []string
		exp  string
	}{
		{
			"empty footer", []Option{Help(Footer("")), Sub(Usage("sub", "a sub"))}, []string{"--help"},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub\n\nFlags:\n  -h, --help      show help, then exit\n",
		},
		{
			"version flag beside the version command", []Option{
				Help(), Version(), VersionString("1.2.3"), Sub(Usage("sub", "a sub")),
				Flags().Bool("version", "show the version", Short("V"), Special("hook:version")),
			}, []string{"-V"},
			"app 1.2.3\n",
		},
		{
			"version command", []Option{
				Help(), Version(), VersionString("1.2.3"), Sub(Usage("sub", "a sub")),
				Flags().Bool("version", "show the version", Short("V"), Special("hook:version")),
			}, []string{"version"},
			"app 1.2.3\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := RunError(
				context.Background(),
				append([]Option{Usage("app", "an app"), Args(test.args...), Pipe(nil, &buf, &buf)}, test.opts...)...,
			)
			switch {
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case buf.String() != test.exp:
				t.Errorf("expected %q, got: %q", test.exp, buf.String())
			}
		})
	}
}
