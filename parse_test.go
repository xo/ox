package ox

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	for i, test := range parseTests() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Logf("invoked: %v", test.v)
			c := &Context{
				Root:   testCommand(t),
				Stdout: new(bytes.Buffer),
				Continue: func(*Command, error) bool {
					return false
				},
				Args: test.v[1:],
				Vars: make(Vars),
			}
			t.Logf("cmd: %s", c.Root.Name)
			t.Logf("args: %q", c.Args)
			if err := c.Parse(); err != nil {
				fmt.Fprintln(c.Stdout, "error:", err)
			}
			if c.Exec != nil {
				t.Logf("exec: %s", c.Exec.Name)
			}
			t.Logf("vars: %s", c.Vars)
			if c.Exec != nil {
				ctx := WithContext(context.Background(), c)
				if err := c.Exec.Exec(ctx, c.Args); err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
			}
			s := strings.TrimSpace(c.Stdout.(*bytes.Buffer).String())
			if exp := strings.Join(test.exp, "\n"); s != exp {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", exp, s)
			}
		})
	}
}

// TestParseErrors tests that each parse error wraps the error that a program
// finds with errors.Is. usql depends on these errors. See B34.
func TestParseErrors(t *testing.T) {
	tests := []struct {
		name string
		args []string
		env  string
		exp  error
	}{
		{"unknown long flag", []string{"--nope"}, "", ErrUnknownFlag},
		{"unknown long flag with value", []string{"--nope=foo"}, "", ErrUnknownFlag},
		{"unknown short flag", []string{"-z"}, "", ErrUnknownFlag},
		{"unknown short flag after a known one", []string{"-bz"}, "", ErrUnknownFlag},
		{"unknown flag of a sub command", []string{"one", "--nope", "a"}, "", ErrUnknownFlag},
		{"missing argument to a long flag", []string{"--str"}, "", ErrMissingArgument},
		{"missing argument to a short flag", []string{"-s"}, "", ErrMissingArgument},
		{"missing argument after a known short flag", []string{"-bs"}, "", ErrMissingArgument},
		{"invalid int with equals", []string{"--int=foo"}, "", ErrInvalidValue},
		{"invalid int as next arg", []string{"--int", "foo"}, "", ErrInvalidValue},
		{"invalid int in a short flag", []string{"-nfoo"}, "", ErrInvalidValue},
		{"invalid int as next arg of a short flag", []string{"-n", "foo"}, "", ErrInvalidValue},
		{"invalid bool", []string{"--bool=maybe"}, "", ErrInvalidValue},
		{"invalid bool in a short flag", []string{"-b=maybe"}, "", ErrInvalidValue},
		{"value that is not valid", []string{"--level=9"}, "", ErrInvalidValue},
		{"invalid date", []string{"--date=foo"}, "", ErrInvalidValue},
		{"map value without a key", []string{"--map=foo"}, "", ErrInvalidValue},
		{"invalid slice element", []string{"--ints=1,x"}, "", ErrInvalidValue},
		{"invalid value for a bound variable", []string{"--bound=foo"}, "", ErrInvalidValue},
		{"invalid uint", []string{"--uint=foo"}, "", ErrInvalidValue},
		{"invalid value from the environment", nil, "foo", ErrInvalidValue},
		{"arg to a command that takes none", []string{"none", "a"}, "", ErrInvalidArgCount},
		{"no arg to a command that takes one", []string{"one"}, "", ErrInvalidArgCount},
		{"two args to a command that takes one", []string{"one", "a", "b"}, "", ErrInvalidArgCount},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.env != "" {
				t.Setenv("OX_PARSE_ERRORS_ENV", test.env)
			}
			var bound int
			ctx, err := NewContext(
				Args(test.args...),
				Usage("cmd", ""),
				Flags().
					String("str", "", Short("s")).
					String("bound", "", Bind(&bound)).
					Bool("bool", "", Short("b")).
					Int("int", "", Short("n")).
					Int("level", "", Valid(1, 2, 3)).
					Uint("uint", "").
					Int("env", "", EnvVar("OX_PARSE_ERRORS_ENV")).
					Date("date", "").
					Map("map", "").
					Slice("ints", "", IntT),
				Sub(Usage("none", ""), ValidArgs(0, 0)),
				Sub(Usage("one", ""), ValidArgs(1, 1)),
			)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			// a program validates the args after the parse, as RunContext does
			if err = ctx.Parse(); err == nil {
				err = ctx.Validate()
			}
			switch {
			case err == nil:
				t.Fatalf("expected error %v, got: nil", test.exp)
			case !errors.Is(err, test.exp):
				t.Errorf("expected error %v, got: %v", test.exp, err)
			}
			t.Logf("err: %v", err)
		})
	}
}

// TestHook tests that a flag built with the Hook option takes no argument,
// and that its func gets a context.Context that holds the Context. See B26.
func TestHook(t *testing.T) {
	errHook := errors.New("hook error")
	tests := []struct {
		name  string
		args  []string
		calls int
		rest  []string
		exp   error
	}{
		{"long", []string{"--hook"}, 1, nil, nil},
		{"short", []string{"-k"}, 1, nil, nil},
		{"twice", []string{"-k", "--hook"}, 2, nil, nil},
		{"before an arg", []string{"--hook", "a"}, 1, []string{"a"}, nil},
		{"with a short bool flag", []string{"-kb"}, 1, nil, nil},
		{"not given", []string{"a"}, 0, []string{"a"}, nil},
		{"error", []string{"--fail"}, 0, nil, errHook},
		{"func without an error", []string{"--plain"}, 1, nil, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var c *Context
			calls := 0
			f := func(ctx context.Context) {
				var ok bool
				if c, ok = Ctx(ctx); !ok || c == nil {
					t.Errorf("expected a context in the hook")
				}
				calls++
			}
			ctx, err := NewContext(
				Args(test.args...),
				Usage("cmd", ""),
				Flags().
					Var("hook", "a hook", Short("k"), Hook(func(ctx context.Context) error {
						f(ctx)
						return nil
					})).
					Var("fail", "a hook that fails", Hook(func(context.Context) error {
						return errHook
					})).
					Hook("plain", "a hook without an error", f).
					Bool("bool", "", Short("b")),
			)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			switch err := ctx.Parse(); {
			case test.exp == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case test.exp != nil && !errors.Is(err, test.exp):
				t.Fatalf("expected error %v, got: %v", test.exp, err)
			}
			if calls != test.calls {
				t.Errorf("expected %d calls, got: %d", test.calls, calls)
			}
			if calls != 0 && c != ctx {
				t.Errorf("expected the hook to get the parse context")
			}
			if test.exp == nil && !slices.Equal(ctx.Args, test.rest) {
				t.Errorf("expected args %q, got: %q", test.rest, ctx.Args)
			}
		})
	}
}

type parseTest struct {
	v   []string
	exp []string
}

func parseTests() []parseTest {
	return []parseTest{
		{
			ss(``),
			[]string{
				`exec: cmd`,
				`name: cmd`,
				`path: []`,
				`args: []`,
				`vars: [int:15]`,
			},
		},
		{
			ss(`a//`),
			[]string{
				`exec: cmd`,
				`name: cmd`,
				`path: []`,
				`args: []`,
				`vars: [int:15]`,
			},
		},
		{
			ss(`a//--foo=a//one////--foo//b//two//--foo//c//three//four//blah//yay`),
			[]string{
				`exec: one`,
				`name: one`,
				`path: [one]`,
				`args: [ two three four blah yay]`,
				`vars: [foo:[a b c] int:15]`,
			},
		},
		{
			ss(`a//--foo=a,b//--foo//c,d//-f=e,f`),
			[]string{
				`exec: cmd`,
				`name: cmd`,
				`path: []`,
				`args: []`,
				`vars: [foo:[a b c d e f] int:15]`,
			},
		},
		{
			ss(`a//-m=A=100//-m//FOO=200//one//two`),
			[]string{
				`exec: two`,
				`name: two`,
				`path: [one two]`,
				`args: []`,
				`vars: [int:15 map:[A:100 FOO:200]]`,
			},
		},
		{
			ss(`a//-b//one//-b=false//two//-b=t//three//-b=1`),
			[]string{
				`exec: three`,
				`name: three`,
				`path: [one two three]`,
				`args: []`,
				`vars: [bvar:true int:15]`,
			},
		},
		{
			ss(`a//--inc//one//--inc//--inc//two//--inc//three`),
			[]string{
				`exec: three`,
				`name: three`,
				`path: [one two three]`,
				`args: []`,
				`vars: [inc:4 int:15]`,
			},
		},
		{
			ss(`a//-fa=b//-iiib//one//two//-bbb//three`),
			[]string{
				`exec: three`,
				`name: three`,
				`path: [one two three]`,
				`args: []`,
				`vars: [bvar:true foo:[a=b] inc:3 int:15]`,
			},
		},
		{
			ss(`a//one//two//four//-b`),
			[]string{
				`exec: two`,
				`name: two`,
				`path: [one two]`,
				`args: [four]`,
				`vars: [bvar:true int:15]`,
			},
		},
		{
			ss(`a//one//four//-iii//--inc//fun`),
			[]string{
				`exec: four`,
				`name: four`,
				`path: [one four]`,
				`args: [fun]`,
				`vars: [inc:4 int:15]`,
			},
		},
		{
			ss(`a//five//-u//file:a//-ufile:b//-c=1.2.3.4/24//--cidr//2.4.6.8/0//foo//bar`),
			[]string{
				`exec: five`,
				`name: five`,
				`path: [five]`,
				`args: [foo bar]`,
				`vars: [cidr:[1.2.3.4/24 2.4.6.8/0] int:15 url:[file:a file:b] val:125]`,
			},
		},
		{
			ss(`a//five//--time//A=07:15:13//a//-d2001-12-25//-d=2002-01-15//--time=B=12:15:32//b`),
			[]string{
				`exec: five`,
				`name: five`,
				`path: [five]`,
				`args: [a b]`,
				`vars: [date:[2001-12-25 2002-01-15] int:15 time:[A:07:15:13 B:12:15:32] val:125]`,
			},
		},
		{
			ss(`a//--//five//--a=b`),
			[]string{
				`exec: cmd`,
				`name: cmd`,
				`path: []`,
				`args: [five --a=b]`,
				`vars: [int:15]`,
			},
		},
		{
			ss(`a//five//-T//255=07:15:32//-T//128=12:15:20//-C=16.1=A//-iiiiC25=J//-C//16.100=C//-C//17.0=1//-C17.0000=//-C//17=//--//--//-a//-b=c`),
			[]string{
				`exec: five`,
				`name: five`,
				`path: [five]`,
				`args: [-- -a -b=c]`,
				`vars: [countmap:[16.1:2 17:3 25:1] inc:4 int:15 timemap:[128:12:15:20 255:07:15:32] val:125]`,
			},
		},
		{
			ss(`a//six//--map//a=b,c=d//--slice//e,f`),
			[]string{
				`exec: six`,
				`name: six`,
				`path: [six]`,
				`args: []`,
				`vars: [int:15 map:[a:b c:d] slice:[e f]]`,
			},
		},
		{
			ss(`a//six//-s//15mib//-r//186mb/s//-d//15us`),
			[]string{
				`exec: six`,
				`name: six`,
				`path: [six]`,
				`args: []`,
				"vars: [duration:15µs int:15 rate:177.38 MiB/s size:15 MiB]",
			},
		},
		// parse errors
		{
			ss(`a//-n//16`),
			[]string{
				`error: -n: invalid value: 16: allowed values: 1, 5, 10, 15`,
			},
		},
		{
			ss(`a//--int=16`),
			[]string{
				`error: --int: invalid value: 16: allowed values: 1, 5, 10, 15`,
			},
		},
		{
			ss(`a//-b//one//-b=fal//two//-b=t//three//-b=1`),
			[]string{
				`error: -b: invalid value: strconv.ParseBool: parsing "fal": invalid syntax`,
			},
		},
		{
			ss(`a//six//--map//a=b,c=d//--slice//e,g`),
			[]string{
				`error: --slice: set 2: invalid value: "g": allowed values: "foo", "bar", "e", "f"`,
			},
		},
	}
}

func ss(s string) []string {
	return strings.Split(s, "//")
}

func testCommand(t *testing.T) *Command {
	t.Helper()
	cmd, err := NewCommand(
		Exec(testDump(t, "cmd")),
		Usage("cmd", ""),
		Flags().
			String("avar", "", Short("a")).
			Bool("bvar", "", Short("b")).
			Count("inc", "", Short("i")).
			Map("map", "", RunesT, Short("m")).
			Slice("foo", "", Short("f")).
			Int("int", "", Short("n"), Default(float64(15.0)), Valid(1, 5, 10, 15)),
		Sub(
			Exec(testDump(t, "one")),
			Usage("one", ""),
			Sub(
				Exec(testDump(t, "two")),
				Usage("two", ""),
				Sub(
					Exec(testDump(t, "three")),
					Usage("three", ""),
				),
			),
			Sub(
				Exec(testDump(t, "four")),
				Usage("four", ""),
			),
			Suggested("remove"),
		),
		Sub(
			Exec(testDump(t, "five")),
			Usage("five", ""),
			Flags().
				String("val", "", Short("l"), Default(125)).
				Slice("cidr", "", Short("c"), CIDRT).
				Slice("url", "", Short("u"), URLT).
				Slice("date", "", Short("d"), DateT).
				Map("time", "", Short("t"), TimeT).
				Map("timemap", "", Short("T"), MapKey(Uint64T), TimeT).
				Map("countmap", "", Short("C"), MapKey(Float64T), CountT),
		),
		Sub(
			Exec(testDump(t, "six")),
			Usage("six", ""),
			Flags().
				Duration("duration", "", Short("d")).
				Size("size", "", Short("s")).
				Rate("rate", "", Short("r")).
				Slice("slice", "", Short("S"), Valid("foo", "bar", "e", "f")).
				Array("array", "", Short("a")).
				Map("map", "", Short("m")),
		),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return cmd
}

func testDump(t *testing.T, name string) func(context.Context, []string) {
	t.Helper()
	return func(ctx context.Context, args []string) {
		c, ok := Ctx(ctx)
		if !ok {
			t.Fatalf("expected non-nil context")
		}
		_, _ = fmt.Fprintln(c.Stdout, "exec:", name)
		_, _ = fmt.Fprintln(c.Stdout, "name:", c.Exec.Name)
		_, _ = fmt.Fprintln(c.Stdout, "path:", c.Exec.Path())
		_, _ = fmt.Fprintln(c.Stdout, "args:", args)
		_, _ = fmt.Fprintln(c.Stdout, "vars:", c.Vars)
	}
}

// TestNoArgExpand holds that the value of a flag without a value expands its
// variables (B19).
func TestNoArgExpand(t *testing.T) {
	t.Setenv("OX_TEST_DIR", "/tmp/x")
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"long", []string{"--rc"}, "/tmp/x/.apprc"},
		{"short", []string{"-r"}, "/tmp/x/.apprc"},
		{"a value", []string{"--rc=other"}, "other"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var s string
			err := RunError(
				context.Background(),
				Usage("app", "an app"),
				Flags().String("rc", "the rc file", Short("r"), NoArg(true, "$ENV{OX_TEST_DIR}/.${APPNAME|lower}rc"), Bind(&s)),
				Exec(func() {}),
				Args(test.args...),
			)
			switch {
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case s != test.exp:
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestDuplicateFlag holds that a command refuses two flags with the same
// name (B19).
func TestDuplicateFlag(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  error
	}{
		{"names", []Option{Flags().String("a", "").Int("a", "")}, ErrDuplicateFlag},
		{"name and alias", []Option{Flags().String("abc", "").Int("b", "", Aliases("abc"))}, ErrDuplicateFlag},
		{"name and hidden alias", []Option{Flags().String("abc", "").Int("b", "", HiddenAliases("abc"))}, ErrDuplicateFlag},
		{"shorts", []Option{Flags().String("a", "", Short("x")).Int("b", "", Short("x"))}, ErrDuplicateFlag},
		{"short and alias", []Option{Flags().String("a", "", Short("x")).Int("b", "", Aliases("x"))}, ErrDuplicateFlag},
		{"different", []Option{Flags().String("a", "", Short("x")).Int("b", "", Short("y"))}, nil},
		{"a one letter name and its short name", []Option{Flags().Bool("H", "", Short("H"))}, nil},
		{"a one letter name and a short name", []Option{Flags().Bool("H", "").Bool("x", "", Short("H"))}, nil},
		{"an alias that repeats the name", []Option{Flags().Bool("abc", "", Aliases("abc"))}, nil},
		{"a parent and a sub command", []Option{Flags().String("a", ""), Sub(Usage("sub", ""), Flags().String("a", ""))}, nil},
		{"two sub commands", []Option{Sub(Usage("one", ""), Flags().String("a", "")), Sub(Usage("two", ""), Flags().String("a", ""))}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewCommand(append([]Option{Usage("app", "")}, test.opts...)...)
			switch {
			case test.exp == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.exp):
				t.Errorf("expected error %v, got: %v", test.exp, err)
			}
		})
	}
}

// TestFlagErrorNames holds that a flag error names the flag as the command
// line gave it, so that a one-letter long flag is --d and not -d (B50).
func TestFlagErrorNames(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"one-letter long", []string{"--d"}, "--d: missing argument"},
		{"short", []string{"-d"}, "-d: missing argument"},
		{"long", []string{"--dir"}, "--dir: missing argument"},
		{"unknown one-letter long", []string{"--x"}, "--x: unknown flag"},
		{"unknown short", []string{"-x"}, "-x: unknown flag"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := RunError(
				context.Background(),
				Usage("app", "an app"),
				Flags().String("d", "", Short("d")).String("dir", ""),
				Exec(func() {}),
				Args(test.args...),
			)
			if err == nil || err.Error() != test.exp {
				t.Errorf("expected %q, got: %v", test.exp, err)
			}
		})
	}
}

// TestDeprecatedWarning holds that a deprecated flag writes a warning, and
// that a deprecated alias does not (B50).
func TestDeprecatedWarning(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"deprecated", []string{"--old", "x"}, "warning: --old is deprecated\n"},
		{"deprecated short", []string{"-o", "x"}, "warning: -o is deprecated\n"},
		{"deprecated alias", []string{"--older", "x"}, ""},
		{"not deprecated", []string{"--new", "x"}, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			err := RunError(
				context.Background(),
				Usage("app", "an app"),
				Flags().
					String("old", "", Short("o"), Deprecated(true)).
					String("new", "", DeprecatedAliases("older")),
				Exec(func() {}),
				Args(test.args...),
				Pipe(nil, &stdout, &stderr),
			)
			switch {
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case stderr.String() != test.exp:
				t.Errorf("expected %q, got: %q", test.exp, stderr.String())
			}
		})
	}
}

// TestOnErrOption holds that an OnErr value passed to Run sets the error
// handling of the context, so that a run can panic in place of an exit (B50).
func TestOnErrOption(t *testing.T) {
	tests := []struct {
		name  string
		onErr OnErr
		panic bool
	}{
		{"panic", OnErrPanic, true},
		{"continue", OnErrContinue, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, err := NewContext(Usage("app", ""), test.onErr, Args("--nope"))
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if ctx.OnErr != test.onErr {
				t.Errorf("expected %v, got: %v", test.onErr, ctx.OnErr)
			}
			defer func() {
				if r := recover(); (r != nil) != test.panic {
					t.Errorf("expected a panic %t, got: %v", test.panic, r)
				}
			}()
			var buf bytes.Buffer
			RunContext(context.Background(), Usage("app", ""), test.onErr, Args("--nope"), Pipe(nil, &buf, &buf), Exec(func() {}))
		})
	}
}

// TestLocal holds that a sub command does not inherit a local flag, and that
// it inherits every other flag of its parents (D38).
func TestLocal(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"local flag of the root", ss(`--root//r`), "app [] [root=r]"},
		{"local short flag of the root", ss(`-r//r`), "app [] [root=r]"},
		{"global flag of the root", ss(`--global//g`), "app [] [global=g]"},
		{"local flag of the root after a sub command", ss(`sub//--root//r`), "--root: unknown flag"},
		{"local short flag of the root after a sub command", ss(`sub//-r//r`), "-r: unknown flag"},
		{"global flag of the root after a sub command", ss(`sub//--global//g`), "sub [] [global=g]"},
		{"local flag of the root before a sub command", ss(`--root//r//sub`), "sub [] [root=r]"},
		{"local flag of a sub command", ss(`sub//--only//o//a`), `sub ["a"] [only=o]`},
		{"local flag of a sub command after its sub command", ss(`sub//child//--only//o`), "--only: unknown flag"},
		{"global flag of the root after two sub commands", ss(`sub//child//--global//g`), "child [] [global=g]"},
		{"global flag of a sub command after its sub command", ss(`sub//child//--shared//s`), "child [] [shared=s]"},
		{"a flag of the root behind a local flag of the same name", ss(`sub//child//--name//n`), "child [] [name=n]"},
		{"a local flag of the same name", ss(`sub//--name//n`), "sub [] [name=n]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var rootName, subName string
			s := parseResult(
				t,
				test.args,
				Usage("app", ""),
				Flags().
					String("root", "", Short("r"), Local()).
					String("global", "").
					String("name", "", Bind(&rootName)),
				Sub(
					Usage("sub", ""),
					Flags().
						String("only", "", Local()).
						String("shared", "").
						String("name", "", Local(), Bind(&subName)),
					Sub(Usage("child", "")),
				),
			)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
			// the root has --name, and sub has a local --name
			switch {
			case slices.Equal(test.args, ss(`sub//child//--name//n`)) && (rootName != "n" || subName != ""):
				t.Errorf("expected the --name of the root, got: root %q, sub %q", rootName, subName)
			case slices.Equal(test.args, ss(`sub//--name//n`)) && (rootName != "" || subName != "n"):
				t.Errorf("expected the --name of sub, got: root %q, sub %q", rootName, subName)
			}
		})
	}
}

// TestLocalTag holds that the tag option local makes a local flag.
func TestLocalTag(t *testing.T) {
	flags, err := FlagsFrom(&struct {
		Dir  string `ox:"the dir,local"`
		Name string `ox:"the name"`
	}{})
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case !flags[0].Local || flags[1].Local:
		t.Errorf("expected only --dir to be local, got: %t %t", flags[0].Local, flags[1].Local)
	}
}

// TestStopAtArg holds that the first positional argument of a command with
// StopAtArg ends its flags (D38).
func TestStopAtArg(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"a flag after the first arg", ss(`exec//cmd//--flag//x`), `exec ["cmd" "--flag" "x"] []`},
		{"flags before the first arg", ss(`exec//-v//--dir//d//cmd//-v`), `exec ["cmd" "-v"] [dir=d verbose=true]`},
		{"a short flag after the first arg", ss(`exec//cmd//-v`), `exec ["cmd" "-v"] []`},
		{"end of the flags before the first arg", ss(`exec//--dir//d//--//cmd//--dir//e`), `exec ["cmd" "--dir" "e"] [dir=d]`},
		{"end of the flags after the first arg", ss(`exec//cmd//--//x`), `exec ["cmd" "--" "x"] []`},
		{"a dash", ss(`exec//-//-v`), `exec ["-" "-v"] []`},
		{"a sub command name after the first arg", ss(`exec//cmd//exec`), `exec ["cmd" "exec"] []`},
		{"an unknown flag before the first arg", ss(`exec//--nope//cmd`), "--nope: unknown flag"},
		{"no stop in a sub command without it", ss(`run//cmd//-v`), `run ["cmd"] [verbose=true]`},
		{"no stop in the root without it", ss(`cmd//-v`), `app ["cmd"] [verbose=true]`},
		{"no stop in a sub command of a command with it", ss(`exec//inner//cmd//-v`), `inner ["cmd"] [verbose=true]`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := parseResult(
				t,
				test.args,
				Usage("app", ""),
				Flags().Bool("verbose", "", Short("v")),
				Sub(
					Usage("exec", ""),
					StopAtArg(),
					Flags().String("dir", "", Short("d")),
					Sub(Usage("inner", "")),
				),
				Sub(Usage("run", "")),
			)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestStopAtArgRoot holds that StopAtArg on a root with sub commands still
// finds a sub command name before the first positional argument (D38).
func TestStopAtArgRoot(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"a sub command", ss(`sub//-v`), `sub [] [verbose=true]`},
		{"a flag before a sub command", ss(`-v//sub//a//-v`), `sub ["a"] [verbose=true]`},
		{"a sub command name after the first arg", ss(`a//sub//-v`), `app ["a" "sub" "-v"] []`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := parseResult(
				t,
				test.args,
				Usage("app", ""),
				StopAtArg(),
				Flags().Bool("verbose", "", Short("v")),
				Sub(Usage("sub", "")),
			)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestAbbrev holds that a unique prefix of a long flag gives the flag when
// the command has Abbrev (D38).
func TestAbbrev(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"a prefix", ss(`--verbo`), "app [] [verbose=true]"},
		{"another prefix", ss(`--verba`), "app [] [verbatim=true]"},
		{"a prefix with a value", ss(`--conf//c`), "app [] [config=c]"},
		{"a prefix with an equal", ss(`--conf=c`), "app [] [config=c]"},
		{"a one letter prefix", ss(`--n//n`), "app [] [name=n]"},
		{"an ambiguous prefix", ss(`--verb`), "--verb: ambiguous flag: --verbatim, --verbose"},
		{"an ambiguous prefix with an equal", ss(`--verb=true`), "--verb: ambiguous flag: --verbatim, --verbose"},
		{"a prefix of a name and an alias of one flag", ss(`--col//red`), "app [] [color=red]"},
		{"a prefix of an alias", ss(`--colou//red`), "app [] [color=red]"},
		{"a prefix of three names", ss(`--co//x`), "--co: ambiguous flag: --color, --colour, --config"},
		{"an exact name wins", ss(`--port//1`), "app [] [port=1]"},
		{"a prefix of the longer name", ss(`--porta//1`), "app [] [portal=1]"},
		{"a prefix of a hidden alias", ss(`--mon//m`), "app [] [name=m]"},
		{"no prefix of a hidden flag", ss(`--secr//s`), "--secr: unknown flag"},
		{"the exact name of a hidden flag", ss(`--secret//s`), "app [] [secret=s]"},
		{"no prefix of a deprecated flag", ss(`--ol//o`), "--ol: unknown flag"},
		{"no prefix of a deprecated alias", ss(`--lab//l`), "--lab: unknown flag"},
		{"no prefix that matches nothing", ss(`--zz`), "--zz: unknown flag"},
		{"no empty prefix", ss(`--=x`), "--: unknown flag"},
		{"no prefix of a short flag", ss(`-x`), "-x: unknown flag"},
		{"a prefix in a sub command", ss(`sub//--verbo`), "sub [] [verbose=true]"},
		{"a prefix of a flag of a sub command", ss(`sub//--cou//3`), "sub [] [count=3]"},
		{"an ambiguous prefix across commands", ss(`sub//--colo//x`), "--colo: ambiguous flag: --color, --colorize, --colour"},
		{"a prefix of a flag that hides a flag of the root", ss(`sub//--conf//c`), "sub [] [config=c]"},
		{"no prefix of a local flag of the root", ss(`sub//--roo//r`), "--roo: unknown flag"},
		{"a prefix of a local flag of the root", ss(`--roo//r`), "app [] [root-dir=r]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := parseResult(
				t,
				test.args,
				Usage("app", ""),
				Abbrev(),
				Flags().
					Bool("verbose", "").
					Bool("verbatim", "").
					String("color", "", Aliases("colour")).
					String("config", "").
					String("name", "", HiddenAliases("moniker"), DeprecatedAliases("label")).
					String("port", "").
					String("portal", "").
					String("secret", "", Hidden(true)).
					String("old", "", Deprecated(true)).
					String("root-dir", "", Local()).
					Bool("x", "", Short("y")),
				Sub(
					Usage("sub", ""),
					Flags().
						Int("count", "").
						Bool("colorize", "").
						String("config", ""),
				),
			)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestAbbrevScope holds that Abbrev is off by default, and that a command
// gets it from itself or from a parent (D38).
func TestAbbrevScope(t *testing.T) {
	tests := []struct {
		name string
		args []string
		opts []Option
		exp  string
	}{
		{"off", ss(`--verbo`), []Option{Sub(Usage("sub", ""))}, "--verbo: unknown flag"},
		{"off in a sub command", ss(`sub//--verbo`), []Option{Sub(Usage("sub", ""))}, "--verbo: unknown flag"},
		{"on in the root", ss(`sub//--verbo`), []Option{Abbrev(), Sub(Usage("sub", ""))}, "sub [] [verbose=true]"},
		{"on in a sub command", ss(`sub//--verbo`), []Option{Sub(Usage("sub", ""), Abbrev())}, "sub [] [verbose=true]"},
		{"on in a sub command and not in the root", ss(`--verbo`), []Option{Sub(Usage("sub", ""), Abbrev())}, "--verbo: unknown flag"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := append([]Option{Usage("app", ""), Flags().Bool("verbose", "")}, test.opts...)
			if s := parseResult(t, test.args, opts...); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestAbbrevError holds that a program finds an ambiguous prefix with
// errors.Is (D38).
func TestAbbrevError(t *testing.T) {
	err := RunError(
		context.Background(),
		Usage("app", ""),
		Abbrev(),
		Flags().Bool("verbose", "").Bool("verbatim", ""),
		Exec(func() {}),
		Args("--verb"),
	)
	if !errors.Is(err, ErrAmbiguousFlag) {
		t.Errorf("expected %v, got: %v", ErrAmbiguousFlag, err)
	}
}

// parseResult parses args with a command of opts, and returns the name of the
// command that runs, its args and the flags that the command line set, or
// the error.
func parseResult(t *testing.T, args []string, opts ...Option) string {
	t.Helper()
	ctx, err := NewContext(append(opts, Args(args...))...)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	ctx.Continue = func(*Command, error) bool {
		return false
	}
	if err := ctx.Parse(); err != nil {
		return err.Error()
	}
	var set []string
	for _, k := range slices.Sorted(maps.Keys(ctx.Vars)) {
		if v := ctx.Vars[k]; v.WasSet() {
			s, _ := v.Get()
			set = append(set, k+"="+s)
		}
	}
	return fmt.Sprintf("%s %q [%s]", ctx.Exec.Name, ctx.Args, strings.Join(set, " "))
}
