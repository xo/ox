package ox

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
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

func TestContextExpand(t *testing.T) {
	root := &Command{
		Name: "myApp",
	}
	configDir, err := userConfigDir()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	cacheDir, err := userCacheDir()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		v   string
		exp string
		err error
	}{
		{``, ``, nil},
		{`a`, `a`, nil},
		{`$ABC`, `$ABC`, nil},
		{`$HOME`, os.Getenv(`HOME`), nil},
		{`$USER`, os.Getenv(`USER`), nil},
		{`$APPNAME`, root.Name, nil},
		{`$CONFIG`, configDir, nil},
		{`$APPCONFIG`, filepath.Join(configDir, root.Name), nil},
		{`$CACHE`, cacheDir, nil},
		{`$APPCACHE`, filepath.Join(cacheDir, root.Name), nil},
		{`$NUMCPU`, strconv.Itoa(runtime.NumCPU()), nil},
		{`$NUMCPU2`, strconv.Itoa(runtime.NumCPU() + 2), nil},
		{`$NUMCPU2X`, strconv.Itoa(runtime.NumCPU() * 2), nil},
		{`$ARCH`, runtime.GOARCH, nil},
		{`$OS`, runtime.GOOS, nil},
		{`$ENV{HOME}`, os.Getenv(`HOME`), nil},
		{`$ENV{NOT_DEFINED}`, ``, nil},
		{`$MY_OVERRIDE{foo}`, `bar`, nil},
		{`$HOME/$USER/$APPNAME`, os.Getenv("HOME") + "/" + os.Getenv("USER") + "/" + root.Name, nil},
		{`$HOME$USER$APPNAME`, os.Getenv("HOME") + os.Getenv("USER") + root.Name, nil},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Logf("test: %q", test.v)
			ctx := &Context{
				Root: root,
				Override: func(typ, key string) (string, bool) {
					if typ == "MY_OVERRIDE" {
						switch strings.ToLower(key) {
						case "foo":
							return "bar", true
						}
					}
					return "", false
				},
			}
			v, err := ctx.Expand(test.v)
			switch {
			case err != nil && !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case err == nil && test.err != nil:
				t.Fatalf("expected error %v, got nil", test.err)
			}
			t.Logf("v: %v (%T)", v, v)
			s, ok := v.(string)
			if !ok {
				t.Fatalf("expected string, got: %T", v)
			}
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
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
