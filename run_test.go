package ox

import (
	"bytes"
	"context"
	"errors"
	"slices"
	"strings"
	"testing"
)

// TestBeforeAfter holds that the before funcs run from the root down, and the
// after funcs from the leaf up, that an error of a before func stops the run,
// that the after funcs run after an error of the exec func, and that help,
// the version and completion run none of them.
func TestBeforeAfter(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		fail  []string
		calls []string
		err   string
	}{
		{"leaf", ss(`sub//leaf//a`), nil, ss(`before app//before sub//before leaf//exec leaf//after leaf//after sub//after app`), ""},
		{"sub", ss(`sub//a`), nil, ss(`before app//before sub//exec sub//after sub//after app`), ""},
		{"before error", ss(`sub//leaf`), ss(`before sub`), ss(`before app//before sub`), "before sub"},
		{"exec error", ss(`sub//leaf`), ss(`exec leaf`), ss(`before app//before sub//before leaf//exec leaf//after leaf//after sub//after app`), "exec leaf"},
		{"after error", ss(`sub//leaf`), ss(`after sub`), ss(`before app//before sub//before leaf//exec leaf//after leaf//after sub//after app`), "after sub"},
		{"exec and after errors", ss(`sub//leaf`), ss(`exec leaf//after leaf//after app`), ss(`before app//before sub//before leaf//exec leaf//after leaf//after sub//after app`), "exec leaf\nafter leaf\nafter app"},
		{"no exec", nil, nil, nil, ""},
		{"help flag", ss(`sub//leaf//--help`), nil, nil, ""},
		{"help command", ss(`help//sub`), nil, nil, ""},
		{"version command", ss(`version`), nil, nil, ""},
		{"completion command", ss(`completion//bash`), nil, nil, ""},
		{"completion", ss(`__complete//sub//`), nil, nil, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var calls []string
			step := func(name string) func(context.Context, []string) error {
				return func(ctx context.Context, _ []string) error {
					if _, ok := Ctx(ctx); !ok {
						t.Errorf("%s: expected the context", name)
					}
					calls = append(calls, name)
					if slices.Contains(test.fail, name) {
						return errors.New(name)
					}
					return nil
				}
			}
			var stdout bytes.Buffer
			err := RunError(
				context.Background(),
				Usage("app", "an app"),
				VersionString("1.2.3"),
				Before(step("before app")),
				After(step("after app")),
				Sub(
					Usage("sub", "a sub command"),
					Before(step("before sub")),
					After(step("after sub")),
					Exec(step("exec sub")),
					Sub(
						Usage("leaf", "a leaf command"),
						Before(step("before leaf")),
						After(step("after leaf")),
						Exec(step("exec leaf")),
					),
				),
				Defaults(),
				Args(test.args...),
				Pipe(nil, &stdout, &stdout),
			)
			var s string
			if err != nil {
				s = err.Error()
			}
			if s != test.err {
				t.Errorf("expected error %q, got: %q", test.err, s)
			}
			for _, name := range test.fail {
				if slices.Contains(calls, name) && !strings.Contains(s, name) {
					t.Errorf("expected the error to hold %q", name)
				}
			}
			if !slices.Equal(calls, test.calls) {
				t.Errorf("expected %q, got: %q", test.calls, calls)
			}
		})
	}
}

// TestBeforeAfterError holds that the run returns the error of the exec func
// as it is when no after func fails, so that errors.As finds its type.
func TestBeforeAfterError(t *testing.T) {
	errExec := &SuggestionError{Parent: &Command{Name: "app"}, Arg: "x", Command: &Command{Name: "y"}}
	err := RunError(
		context.Background(),
		Usage("app", "an app"),
		Before(func() {}),
		After(func() {}),
		Exec(func() error {
			return errExec
		}),
		Args(),
	)
	if err != errExec { //nolint:errorlint
		t.Errorf("expected the error of the exec func, got: %v", err)
	}
}

// TestAppBeforeAfter holds that App takes the Before and After methods of a
// struct.
func TestAppBeforeAfter(t *testing.T) {
	v := &testHooks{calls: new([]string)}
	if err := RunError(context.Background(), App(v), Args("sub")); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if exp := ss(`before app//before sub//exec sub//after sub//after app`); !slices.Equal(*v.calls, exp) {
		t.Errorf("expected %q, got: %q", exp, *v.calls)
	}
}

type testHooks struct {
	calls *[]string
}

func (v *testHooks) Before(context.Context) error {
	*v.calls = append(*v.calls, "before app")
	return nil
}

func (v *testHooks) After() {
	*v.calls = append(*v.calls, "after app")
}

func (v *testHooks) Commands() []any {
	return []any{&testHooksSubCommand{calls: v.calls}}
}

type testHooksSubCommand struct {
	calls *[]string
}

func (v *testHooksSubCommand) Before([]string) error {
	*v.calls = append(*v.calls, "before sub")
	return nil
}

func (v *testHooksSubCommand) After(context.Context, []string) error {
	*v.calls = append(*v.calls, "after sub")
	return nil
}

func (v *testHooksSubCommand) Exec() {
	*v.calls = append(*v.calls, "exec sub")
}
