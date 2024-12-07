package ox

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestSuggestions(t *testing.T) {
	stderr := new(bytes.Buffer)
	var code int
	cmd := testCommand(t)
	cmd.Exec = nil
	c := &Context{
		Exit: func(c int) {
			code = c
		},
		Panic: func(v any) {
			panic(v)
		},
		Root:   cmd,
		Stderr: stderr,
		Args:   []string{"on"},
		Vars:   make(Vars),
	}
	c.Handler = DefaultErrorHandler(c)
	if err := c.Parse(); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	err := c.Run(context.Background())
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
	if s, exp := err.Error(), `unknown command "on" for "cmd"`; s != exp {
		t.Fatalf("expected %q, got: %q", exp, s)
	}
	if !c.Handler(err) {
		t.Fatalf("expected Handler to return true")
	}
	if code == 0 {
		t.Fatalf("expected Handler to set non-zero code")
	}
	exp := `error: unknown command "on" for "cmd"

Did you mean this?
  one

`
	if s := stderr.String(); s != exp {
		t.Errorf("\nexpected:\n%s\ngot:\n%s", exp, s)
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
		{"a", "", 1},
		{"foo", "", 3},
		{"foo", "bar", 3},
		{"foo", "food", 1},
		{"Food", "foo", 1},
		{"fod", "Food", 1},
		{"br", "bar", 1},
		{"bar", "br", 1},
		{"test", "t", 3},
		{"book", "back", 2},
		{"kitten", "sitting", 3},
		{"Kitten", "Sitting", 3},
	}
	for _, test := range tests {
		t.Run(test.a+"::"+test.b, func(t *testing.T) {
			a, b := []rune(strings.ToLower(test.a)), []rune(strings.ToLower(test.b))
			if i := Ldist(a, b); i != test.exp {
				t.Errorf("expected %d, got: %d", test.exp, i)
			}
		})
	}
}
