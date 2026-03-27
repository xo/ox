package ox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestSuggestions(t *testing.T) {
	for i, test := range []struct {
		args []string
		exp  string
	}{
		{
			ss(`subfoo`), `error: unknown command "subfoo" for "cmd"`,
		},
		{
			ss(`on`), `error: unknown command "on" for "cmd"

Did you mean this?
  one
`,
		},
		{
			ss(`remove`), `error: unknown command "remove" for "cmd"

Did you mean this?
  one
`,
		},
		{
			ss(`rmove`), `error: unknown command "rmove" for "cmd"

Did you mean this?
  one
`,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var code int
			c := testContext(t, &code, test.args...)
			err := c.Run(context.Background())
			if err == nil {
				t.Fatalf("expected non-nil error")
			}
			if !c.Handler(err) {
				t.Fatalf("expected Handler to return true")
			}
			if code == 0 {
				t.Fatalf("expected Handler to set non-zero code")
			}
			if s := strings.TrimSuffix(c.Stderr.(*bytes.Buffer).String(), "\n"); s != test.exp {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", test.exp, s)
			}
		})
	}
}

func TestLdist(t *testing.T) {
	tests := []struct {
		a   string
		b   string
		exp int
	}{
		{"", "", 0},
		{"", "a", 1},
		{"ab", "aa", 1},
		{"ab", "aaa", 2},
		{"ab", "ba", 2},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"bar", "br", 1},
		{"bbb", "a", 3},
		{"book", "back", 2},
		{"br", "bar", 1},
		{"distance", "difference", 5},
		{"fod", "Food", 1},
		{"foo", "", 3},
		{"foo", "bar", 3},
		{"foo", "Food", 1},
		{"Hafþór Júlíus Björnsson", "Hafþor Julius Bjornsson", 4},
		{"", "hello", 5},
		{"hello", "hello", 0},
		{"kitten", "sitting", 3},
		{"levenshtein", "frankenstein", 6},
		{"mississippi", "swiss miss", 8},
		{"resume and cafe", "resumé and café", 2},
		{"resume and cafe", "resumes and cafes", 2},
		{"resumé and café", "resumés and cafés", 2},
		{"rosettacode", "raisethysword", 8},
		{"saturday", "sunday", 3},
		{"sitten", "sittin", 1},
		{"sittin", "sitting", 1},
		{"sleep", "fleeting", 5},
		{"stop", "tops", 2},
		{"test", "t", 3},
		{"།་གམ་འས་པ་་མ།", "།་གམའས་པ་་མ", 2},
		{"こんにちは", "こんちは", 1},
		{"👀😀", "😀", 1},
	}
	for _, test := range tests {
		t.Run(test.a+"::"+test.b, func(t *testing.T) {
			a, b := []rune(strings.ToLower(test.a)), []rune(strings.ToLower(test.b))
			if i := Ldist(a, b); i != test.exp {
				t.Errorf("expected %d, got: %d", test.exp, i)
			}
			if i := Ldist(b, a); i != test.exp {
				t.Errorf("expected %d, got: %d", test.exp, i)
			}
		})
	}
}

func TestEmitExitError(t *testing.T) {
	if isTinyGo {
		t.Log("skipping: tinygo incompatible")
		return
	}
	code := new(int)
	c := testContext(t, code)
	c.Exec = testExitError(t)
	c.EmitExecExitError = true
	err := c.Run(context.Background())
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	switch ok := c.Handler(err); {
	case !ok:
		t.Errorf("expected to continue")
	case *code == 0, *code == 1:
		t.Errorf("expected error code != 0 and error code != 1, got: %d", *code)
	}
	t.Logf("err: %v, code: %d", err, *code)
	stdout, ok := c.Stdout.(*bytes.Buffer)
	if !ok {
		t.Fatal("expected *bytes.Buffer")
	}
	if s := stdout.String(); s != "" {
		t.Errorf("expected no output, got: %q", s)
	}
	stderr, ok := c.Stderr.(*bytes.Buffer)
	if !ok {
		t.Fatal("expected *bytes.Buffer")
	}
	if s, exp := stderr.String(), "error: (my error): exit status 2: expr: division by zero\n"; s != exp {
		t.Errorf("expected %q, got: %q", exp, s)
	}
}

func testContext(t *testing.T, code *int, args ...string) *Context {
	t.Helper()
	cmd := testCommand(t)
	cmd.Exec = nil
	c := &Context{
		Exit: func(c int) {
			*code = c
		},
		Panic: func(v any) {
			t.Fatalf("context panic: %v", v)
		},
		Root:   cmd,
		Stdout: new(bytes.Buffer),
		Stderr: new(bytes.Buffer),
		Args:   args,
		Vars:   make(Vars),
	}
	c.Handler = DefaultErrorHandler(c)
	if err := c.Parse(); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return c
}

// testExitError is a command that when executed always returns [exec.ExitError].
func testExitError(t *testing.T) *Command {
	t.Helper()
	cmd, err := NewCommand(Exec(func(ctx context.Context) error {
		c := exec.CommandContext(ctx, "expr", "1", "/", "0")
		buf, err := c.Output()
		t.Logf("command output %q", string(buf))
		if err != nil {
			t.Logf("error: %v (%T)", err, err)
			if e, ok := err.(*exec.ExitError); ok {
				t.Logf("  exec.ExitError.Stderr: %q", string(e.Stderr))
			}
		}
		return fmt.Errorf("(my error): %w", err)
	}))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return cmd
}
