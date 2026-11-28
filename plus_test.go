package ox

import (
	"bytes"
	"context"
	"errors"
	"slices"
	"strings"
	"testing"
)

// plusOpts returns a command line in the shape of bash's: -o and -O take a
// name, -e, -x and -v are hooks, -C is a plain bool, -n takes a string, and
// -c has no plus form. Each hook records its sign and its value in events, in
// the order of the command line (D52).
func plusOpts(events *[]string, args *[]string) []Option {
	sign := func(ctx *Context) string {
		if ctx.IsPlus() {
			return "+"
		}
		return "-"
	}
	value := func(letter string) func(*Context, string) error {
		return func(ctx *Context, s string) error {
			*events = append(*events, sign(ctx)+letter+" "+s)
			return nil
		}
	}
	bare := func(letter string) func(*Context) error {
		return func(ctx *Context) error {
			*events = append(*events, sign(ctx)+letter)
			return nil
		}
	}
	return []Option{
		Usage("sh", "a shell"),
		StopAtArg(),
		Flags().
			Var("set-option", "set an option", HookT, Default(value("o")), Short("o"), AllowPlus(), Spec("NAME")).
			Var("shopt", "set a shell option", HookT, Default(value("O")), Short("O"), AllowPlus(), Spec("NAME")).
			Hook("errexit", "exit on an error", bare("e"), Short("e"), AllowPlus()).
			Hook("xtrace", "trace each command", bare("x"), Short("x"), AllowPlus()).
			Hook("verbose", "write each line", bare("v"), Short("v"), AllowPlus()).
			Bool("noclobber", "do not overwrite a file", Short("C"), AllowPlus()).
			String("name", "the name", Short("n"), AllowPlus()).
			Bool("command", "read the commands from the arg", Short("c")),
		Exec(func(_ context.Context, a []string) {
			*args = a
		}),
	}
}

// TestPlus tests the plus form of short flags, as bash takes `+o name` and
// `+ex` (D52).
func TestPlus(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		events []string
		rest   []string
		vars   string
		err    error
	}{
		{"order of one option", ss(`-o//errexit//+o//errexit`), []string{"-o errexit", "+o errexit"}, nil, "", nil},
		{"order across options", ss(`-e//+o//errexit`), []string{"-e", "+o errexit"}, nil, "", nil},
		{"shopt", ss(`-O//extglob//+O//extglob`), []string{"-O extglob", "+O extglob"}, nil, "", nil},
		{"a value in the word", ss(`+oerrexit`), []string{"+o errexit"}, nil, "", nil},
		{"a value after an equal", ss(`+o=errexit`), []string{"+o errexit"}, nil, "", nil},
		{"one letter", ss(`+e`), []string{"+e"}, nil, "", nil},
		{"a cluster", ss(`+xv`), []string{"+x", "+v"}, nil, "", nil},
		{"a cluster that ends with a value", ss(`+eo//nounset`), []string{"+e", "+o nounset"}, nil, "", nil},
		{"a bool", ss(`-C//+C`), nil, nil, "noclobber:false", nil},
		{"a bool on", ss(`+C//-C`), nil, nil, "noclobber:true", nil},
		{"a string", ss(`+n//x`), nil, nil, "name:x", nil},
		{"a lone plus is an arg", ss(`+`), nil, ss(`+`), "", nil},
		{"a lone dash is an arg", ss(`-`), nil, ss(`-`), "", nil},
		{"a word without a plus flag is an arg", ss(`+%Y`), nil, ss(`+%Y`), "", nil},
		{"an unknown first letter is an arg", ss(`+z//-e`), nil, ss(`+z//-e`), "", nil},
		{"a dash dash ends the flags", ss(`--//+e`), nil, ss(`+e`), "", nil},
		{"the first arg ends the flags", ss(`script//+e`), nil, ss(`script//+e`), "", nil},
		{"a letter without a plus form", ss(`+ec`), []string{"+e"}, nil, "", ErrUnknownFlag},
		{"a missing value", ss(`+o`), nil, nil, "", ErrMissingArgument},
		{"a value for a bool", ss(`+e=1`), nil, nil, "", ErrInvalidValue},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var events, rest []string
			ctx, err := NewContext(append(plusOpts(&events, &rest), Args(test.args...))...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			err = ctx.Parse()
			if err == nil {
				err = ctx.Run(context.Background())
			}
			switch {
			case test.err != nil && !errors.Is(err, test.err):
				t.Fatalf("expected %v, got: %v", test.err, err)
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			}
			if !slices.Equal(events, test.events) {
				t.Errorf("expected events %q, got: %q", test.events, events)
			}
			if test.err == nil && !slices.Equal(rest, test.rest) {
				t.Errorf("expected args %q, got: %q", test.rest, rest)
			}
			if test.vars != "" && !strings.Contains(ctx.Vars.String(), test.vars) {
				t.Errorf("expected vars to hold %q, got: %s", test.vars, ctx.Vars)
			}
			if ctx.IsPlus() {
				t.Error("expected IsPlus to be false after the parse")
			}
		})
	}
}

// TestPlusErrors tests the flags that AllowPlus refuses.
func TestPlusErrors(t *testing.T) {
	if _, err := NewFlag("errexit", "exit on an error", AllowPlus()); !errors.Is(err, ErrInvalidValue) {
		t.Errorf("expected %v without a short name, got: %v", ErrInvalidValue, err)
	}
	if _, err := NewFlag("verbose", "verbosity", CountT, Short("v"), AllowPlus()); !errors.Is(err, ErrAppliedToInvalidType) {
		t.Errorf("expected %v for a count, got: %v", ErrAppliedToInvalidType, err)
	}
}

// TestPlusTag tests the struct tag option plus.
func TestPlusTag(t *testing.T) {
	var args struct {
		Noclobber bool `ox:"do not overwrite a file,short:C,plus"`
		Command   bool `ox:"read the commands from the arg,short:c"`
	}
	err := RunError(context.Background(), Usage("sh", "a shell"), From(&args), Exec(func() {}), Args("-C", "+C", "-c"))
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case args.Noclobber || !args.Command:
		t.Errorf("expected noclobber false and command true, got: %+v", args)
	}
	err = RunError(context.Background(), Usage("sh", "a shell"), From(&args), Exec(func() {}), Args("+c"))
	if err != nil {
		t.Errorf("expected +c to be an arg without plus, got: %v", err)
	}
}

// TestPlusHelp tests that help writes the plus form in the short column, in
// both layouts, and lines up each flag.
func TestPlusHelp(t *testing.T) {
	var events, rest []string
	for _, test := range []struct {
		name string
		opts []Option
		exp  []string
	}{
		{"default", nil, []string{
			"  -o, +o, --set-option NAME  set an option\n",
			"  -C, +C, --noclobber        do not overwrite a file\n",
			"  -c,     --command          read the commands from the arg\n",
		}},
		{"column", []Option{HelpColumn(30)}, []string{
			"  -o, +o, --set-option NAME   set an option\n",
			"  -c, --command               read the commands from the arg\n",
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			cmd, err := NewCommand(plusOpts(&events, &rest)...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			help, err := NewCommandHelp(cmd, test.opts...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			var buf bytes.Buffer
			help.AddFlags(&buf)
			buf.WriteString("\n")
			for _, exp := range test.exp {
				if !strings.Contains(buf.String(), exp) {
					t.Errorf("expected help to hold %q, got:\n%s", exp, buf.String())
				}
			}
		})
	}
}

// TestPlusComp tests the completion of the plus forms and of their values.
func TestPlusComp(t *testing.T) {
	var events, rest []string
	tests := []struct {
		args []string
		exp  []string
	}{
		{ss(`+`), ss(`+o//+O//+e//+x//+v//+C//+n`)},
		{ss(`+x`), ss(`+x`)},
		{ss(`+%`), nil},
		{ss(`-e//+`), ss(`+o//+O//+e//+x//+v//+C//+n`)},
	}
	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			root, err := NewCommand(plusOpts(&events, &rest)...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			ctx := &Context{
				Root: root,
				Args: append([]string{DefaultCompName}, test.args...),
			}
			comps, _, err := ctx.Comps()
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			var names []string
			for _, c := range comps {
				names = append(names, c.Name)
			}
			if !slices.Equal(names, test.exp) {
				t.Errorf("expected %q, got: %q", test.exp, names)
			}
		})
	}
}
