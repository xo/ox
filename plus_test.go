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

// bashOpts returns the command line of plusOpts with bash's rules: each `+`
// word is flags, letters take their value from the next word, and -o and -O
// can be the last word without a value. A hook records `<none>` for no value
// (D53). letters turns LetterClusters on.
func bashOpts(events, args *[]string, letters bool) []Option {
	sign := func(ctx *Context) string {
		if ctx.IsPlus() {
			return "+"
		}
		return "-"
	}
	value := func(letter string) func(*Context, string) error {
		return func(ctx *Context, s string) error {
			if ctx.NoValue() {
				s = "<none>"
			}
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
	opts := []Option{
		Usage("sh", "a shell"),
		StopAtArg(),
		PlusWords(),
		Flags().
			Var("set-option", "set an option", HookT, Default(value("o")), Short("o"), AllowPlus(), OptionalAtEnd()).
			Var("shopt", "set a shell option", HookT, Default(value("O")), Short("O"), AllowPlus(), OptionalAtEnd()).
			Hook("errexit", "exit on an error", bare("e"), Short("e"), AllowPlus()).
			Hook("xtrace", "trace each command", bare("x"), Short("x"), AllowPlus()).
			String("name", "the name", Short("n")),
		Exec(func(_ context.Context, a []string) {
			*args = a
		}),
	}
	if letters {
		opts = append(opts, LetterClusters())
	}
	return opts
}

// TestBash tests PlusWords, LetterClusters and OptionalAtEnd (D53).
func TestBash(t *testing.T) {
	tests := []struct {
		name    string
		letters bool
		args    []string
		events  []string
		rest    []string
		err     error
		errText string
	}{
		// PlusWords
		{"a plus word without a plus flag", true, ss(`+z`), nil, nil, ErrUnknownFlag, "+z: unknown flag"},
		{"a plus word of the date kind", true, ss(`+%Y`), nil, nil, ErrUnknownFlag, ""},
		{"a lone plus is skipped", true, ss(`+//-e`), []string{"-e"}, nil, nil, ""},
		{"a lone plus before an arg", true, ss(`+//script`), nil, ss(`script`), nil, ""},
		{"a plus word after the first arg", true, ss(`script//+z`), nil, ss(`script//+z`), nil, ""},
		// LetterClusters
		{"a letter takes the next word", true, ss(`-oe//errexit`), []string{"-o errexit", "-e"}, nil, nil, ""},
		{"two letters take two words", true, ss(`-oO//errexit//extglob`), []string{"-o errexit", "-O extglob"}, nil, nil, ""},
		{"a plus cluster takes the next word", true, ss(`+oe//errexit`), []string{"+o errexit", "+e"}, nil, nil, ""},
		{"an equal is a letter", true, ss(`-e=`), []string{"-e"}, nil, ErrUnknownFlag, "-=: unknown flag"},
		{"getopt takes the rest of the word", false, ss(`-oe//errexit`), []string{"-o e"}, ss(`errexit`), nil, ""},
		// OptionalAtEnd
		{"a value flag at the end", true, ss(`-o`), []string{"-o <none>"}, nil, nil, ""},
		{"a plus value flag at the end", true, ss(`+O`), []string{"+O <none>"}, nil, nil, ""},
		{"a value flag at the end after a flag", true, ss(`-e//+o`), []string{"-e", "+o <none>"}, nil, nil, ""},
		{"a cluster at the end", true, ss(`-oe`), []string{"-o <none>", "-e"}, nil, nil, ""},
		{"a value flag at the end with getopt", false, ss(`-o`), []string{"-o <none>"}, nil, nil, ""},
		{"a long value flag at the end", false, ss(`--set-option`), []string{"-o <none>"}, nil, nil, ""},
		{"a value flag not at the end", true, ss(`-o//errexit`), []string{"-o errexit"}, nil, nil, ""},
		{"no OptionalAtEnd", true, ss(`-n`), nil, nil, ErrMissingArgument, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var events, rest []string
			ctx, err := NewContext(append(bashOpts(&events, &rest, test.letters), Args(test.args...))...)
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
			case test.errText != "" && err.Error() != test.errText:
				t.Errorf("expected %q, got: %q", test.errText, err.Error())
			}
			if !slices.Equal(events, test.events) {
				t.Errorf("expected events %q, got: %q", test.events, events)
			}
			if test.err == nil && !slices.Equal(rest, test.rest) {
				t.Errorf("expected args %q, got: %q", test.rest, rest)
			}
			if ctx.IsPlus() || ctx.NoValue() {
				t.Error("expected IsPlus and NoValue to be false after the parse")
			}
		})
	}
}

// TestBashInherit tests that a sub command inherits PlusWords and
// LetterClusters, that OptionalAtEnd refuses a flag that takes no value, and
// that completion finds the letter that takes the next word.
func TestBashInherit(t *testing.T) {
	var got []string
	err := RunError(context.Background(),
		Usage("app", "an app"),
		PlusWords(),
		LetterClusters(),
		Sub(
			Usage("sh", "a shell"),
			Flags().
				Var("set-option", "set an option", HookT, Default(func(ctx *Context, s string) {
					got = append(got, s)
				}), Short("o"), AllowPlus()).
				Bool("errexit", "exit on an error", Short("e"), AllowPlus()),
			Exec(func() {}),
		),
		Args("sh", "-oe", "errexit", "+", "+e"),
	)
	if err != nil || !slices.Equal(got, []string{"errexit"}) {
		t.Errorf("expected [errexit] and no error, got: %q %v", got, err)
	}
	if _, err := NewFlag("errexit", "exit", BoolT, Short("e"), OptionalAtEnd()); !errors.Is(err, ErrAppliedToInvalidType) {
		t.Errorf("expected %v for a flag that takes no value, got: %v", ErrAppliedToInvalidType, err)
	}
	for _, letters := range []bool{false, true} {
		var events, rest []string
		cmd, err := NewCommand(bashOpts(&events, &rest, letters)...)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		switch g := compPrevFlag(cmd, "-oe"); {
		case letters && (g == nil || g.Name != "set-option"):
			t.Errorf("expected -oe to complete the value of -o, got: %v", g)
		case !letters && g != nil:
			t.Errorf("expected -oe to complete no value with getopt, got: %v", g.Name)
		}
	}
}

// TestBashLongFlags tests that a command with BashLongFlags refuses a long
// flag after a short flag, and a long flag with a value after an equal, as
// bash does (D54).
func TestBashLongFlags(t *testing.T) {
	tests := []struct {
		name    string
		bash    bool
		args    []string
		rest    []string
		err     error
		errText string
	}{
		{"a long switch", true, ss(`--norc`), nil, nil, ""},
		{"a long switch with a value", true, ss(`--norc=1`), nil, ErrInvalidValue, "--norc: invalid value: a long flag takes its value from the next word, and not after ="},
		{"a long flag and the next word", true, ss(`--rcfile//x`), nil, nil, ""},
		{"a long flag with a value", true, ss(`--rcfile=x`), nil, ErrInvalidValue, ""},
		{"a long flag after a short flag", true, ss(`-e//--norc`), nil, ErrUnknownFlag, "--norc: unknown flag: a long flag comes before the short flags"},
		{"a long flag after a plus form", true, ss(`+e//--norc`), nil, ErrUnknownFlag, ""},
		{"a long flag before a short flag", true, ss(`--norc//-e//+e`), nil, nil, ""},
		{"a dash dash after a short flag", true, ss(`-e//--//--norc`), ss(`--norc`), nil, ""},
		{"a sub command inherits it", true, ss(`sub//-e//--norc`), nil, ErrUnknownFlag, ""},
		{"no BashLongFlags", false, ss(`-e//--norc=1//--rcfile=x`), nil, nil, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var rest []string
			exec := Exec(func(_ context.Context, a []string) {
				rest = a
			})
			opts := []Option{
				Usage("sh", "a shell"),
				Flags().
					Bool("errexit", "exit on an error", Short("e"), AllowPlus()).
					Bool("norc", "read no startup file").
					String("rcfile", "the startup file"),
				Sub(Usage("sub", "a sub command"), exec),
				exec,
				Args(test.args...),
			}
			if test.bash {
				opts = append(opts, BashLongFlags())
			}
			err := RunError(context.Background(), opts...)
			switch {
			case test.err != nil && !errors.Is(err, test.err):
				t.Fatalf("expected %v, got: %v", test.err, err)
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case test.errText != "" && err.Error() != test.errText:
				t.Errorf("expected %q, got: %q", test.errText, err.Error())
			}
			if test.err == nil && !slices.Equal(rest, test.rest) {
				t.Errorf("expected args %q, got: %q", test.rest, rest)
			}
		})
	}
}
