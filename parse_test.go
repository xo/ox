package ox

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
			switch err := c.Parse(); {
			case err != nil:
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
