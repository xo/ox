package ox

import (
	"context"
	"errors"
	"slices"
	"testing"
)

// TestArg holds that each positional argument converts to its type, is
// checked, and is in the vars by its name.
func TestArg(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		args []string
		exp  string
		err  error
	}{
		{"string", opts(Arg("name", "a name")), ss(`x`), "[name:x]", nil},
		{"missing", opts(Arg("name", "a name")), nil, "", ErrMissingArgument},
		{"too many", opts(Arg("name", "a name")), ss(`x//y`), "", ErrInvalidArgCount},
		{"int", opts(Arg("n", "a number", IntT)), ss(`12`), "[n:12]", nil},
		{"bad int", opts(Arg("n", "a number", IntT)), ss(`x`), "", ErrInvalidValue},
		{"file", opts(Arg("file", "a file", FileT)), ss(`go.mod`), "[file:go.mod]", nil},
		{"missing file", opts(Arg("file", "a file", FileT)), ss(`missing.txt`), "", ErrNotExist},
		{"dir", opts(Arg("dir", "a dir", DirT)), ss(`go.mod`), "", ErrNotDir},
		{"valid", opts(Arg("mode", "a mode", Valid("fast", "slow"))), ss(`fast`), "[mode:fast]", nil},
		{"not valid", opts(Arg("mode", "a mode", Valid("fast", "slow"))), ss(`quick`), "", ErrInvalidValue},
		{"default", opts(Arg("a", "a"), Arg("n", "a number", IntT, Default(3))), ss(`x`), "[a:x n:3]", nil},
		{"default not used", opts(Arg("n", "a number", IntT, Default(3))), ss(`4`), "[n:4]", nil},
		{"default expands", opts(Arg("s", "a string", Default("${APPNAME}"))), nil, "[s:app]", nil},
		{"optional", opts(Arg("a", "a"), Arg("b", "b", Optional())), ss(`x`), "[a:x]", nil},
		{"variadic", opts(Arg("files", "files", SliceT, Elem(FileT))), ss(`go.mod//ox.go`), "[files:[go.mod ox.go]]", nil},
		{"variadic missing file", opts(Arg("files", "files", SliceT, Elem(FileT))), ss(`go.mod//missing.txt`), "", ErrNotExist},
		{"variadic required", opts(Arg("files", "files", SliceT)), nil, "", ErrMissingArgument},
		{"variadic optional", opts(Arg("files", "files", SliceT, Optional())), nil, "[]", nil},
		{"variadic after an arg", opts(Arg("a", "a"), Arg("rest", "rest", SliceT, Elem(IntT))), ss(`x//1//2//3`), "[a:x rest:[1 2 3]]", nil},
		{"valid args too", opts(ValidArgs(0, 1), Arg("a", "a", Optional())), ss(`x//y`), "", ErrInvalidArgCount},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var vars Vars
			err := RunError(
				context.Background(),
				append(test.opts,
					Usage("app", "an app"),
					Exec(func(ctx context.Context) {
						c, _ := Ctx(ctx)
						vars = c.Vars
					}),
					Args(test.args...),
				)...,
			)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.err == nil && vars.String() != test.exp:
				t.Errorf("expected %s, got: %s", test.exp, vars)
			}
		})
	}
}

// TestArgBind holds that a bound variable gets the value of an argument, and
// that a variadic argument takes each value whole.
func TestArgBind(t *testing.T) {
	var n int
	var rest []string
	err := RunError(
		context.Background(),
		Usage("app", "an app"),
		Arg("n", "a number", IntT, Bind(&n)),
		Arg("rest", "the rest", SliceT, Bind(&rest)),
		Exec(func() {}),
		Args("12", "a,b", "c"),
	)
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case n != 12:
		t.Errorf("expected 12, got: %d", n)
	case !slices.Equal(rest, []string{"a,b", "c"}):
		t.Errorf("expected [a,b c], got: %q", rest)
	}
}

// TestArgErrors holds the errors of a list of arguments that a command cannot
// parse.
func TestArgErrors(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		err  error
	}{
		{"after a variadic", opts(Arg("a", "a", SliceT), Arg("b", "b")), ErrInvalidArg},
		{"required after optional", opts(Arg("a", "a", Optional()), Arg("b", "b")), ErrInvalidArg},
		{"two names", opts(Arg("a", "a"), Arg("a", "a")), ErrInvalidArg},
		{"name of a flag", opts(Flags().String("a", "a"), Arg("a", "a")), ErrInvalidArg},
		{"option of a command", opts(Arg("a", "a", ConfigFiles("x"))), ErrAppliedToInvalidType},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewCommand(append(test.opts, Usage("app", "an app"))...)
			if !errors.Is(err, test.err) {
				t.Errorf("expected error %v, got: %v", test.err, err)
			}
		})
	}
	if _, err := NewFlag("a", "a", Optional()); !errors.Is(err, ErrAppliedToInvalidType) {
		t.Errorf("expected error %v, got: %v", ErrAppliedToInvalidType, err)
	}
}

// TestArgFrom holds that the tag option `arg` makes a field a positional
// argument for From and App, and that FlagsFrom leaves it out.
func TestArgFrom(t *testing.T) {
	var v struct {
		Src     string   `ox:"the source,arg,type:file"`
		Dst     string   `ox:"the target,arg,optional"`
		Rest    []string `ox:"the rest,arg,optional"`
		Verbose bool     `ox:"more output"`
	}
	flags, err := FlagsFrom(&v)
	switch {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case len(flags) != 1 || flags[0].Name != "verbose":
		t.Fatalf("expected only the flag verbose, got: %v", flags)
	}
	if err := RunError(
		context.Background(),
		Usage("app", "an app"),
		From(&v),
		Exec(func() {}),
		Args("--verbose", "go.mod", "out", "a,b", "c"),
	); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	switch {
	case v.Src != "go.mod", v.Dst != "out", !v.Verbose:
		t.Errorf("expected go.mod, out and true, got: %q, %q and %t", v.Src, v.Dst, v.Verbose)
	case !slices.Equal(v.Rest, []string{"a,b", "c"}):
		t.Errorf("expected [a,b c], got: %q", v.Rest)
	}
	err = RunError(context.Background(), Usage("app", "an app"), From(&v), Exec(func() {}), Args("missing.txt"))
	if !errors.Is(err, ErrNotExist) {
		t.Errorf("expected error %v, got: %v", ErrNotExist, err)
	}
	var bad struct {
		A string `ox:"a,optional"`
	}
	if _, err := FlagsFrom(&bad); !errors.Is(err, ErrInvalidValue) {
		t.Errorf("expected error %v, got: %v", ErrInvalidValue, err)
	}
}

// TestArgApp holds that App reads the tag option `arg`.
func TestArgApp(t *testing.T) {
	v := new(testArgApp)
	if err := RunError(context.Background(), App(v), Args("go.mod")); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if v.File != "go.mod" || v.ran != "go.mod" {
		t.Errorf("expected go.mod, got: %q and %q", v.File, v.ran)
	}
}

type testArgApp struct {
	File string `ox:"the file,arg,type:file"`
	ran  string
}

func (v *testArgApp) Exec(ctx context.Context) error {
	c, _ := Ctx(ctx)
	s, err := c.Vars["file"].Get()
	v.ran = s
	return err
}

// TestArgHelp holds the usage line and the Arguments: section of the
// positional arguments.
func TestArgHelp(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{
			"required",
			opts(Arg("src", "the file to copy", FileT)),
			"app an app\n\nUsage:\n  app src\n\nArguments:\n  src     the file to copy\n",
		},
		{
			"optional and variadic",
			opts(
				Arg("src", "the file to copy"),
				Arg("mode", "the mode", Default("fast")),
				Arg("files", "more files", SliceT, Optional(), Spec("FILE")),
			),
			"app an app\n\nUsage:\n  app src [mode] [FILE...]\n\nArguments:\n  src     the file to copy\n  mode    the mode (default: fast)\n  FILE    more files\n",
		},
		{
			"required variadic",
			opts(Arg("files", "the files", SliceT)),
			"app an app\n\nUsage:\n  app files...\n\nArguments:\n  files   the files\n",
		},
		{
			"with arg usage",
			opts(Arg("src", "the file to copy"), Help(ArgUsage("EXTRA", "an extra line"))),
			"app an app\n\nUsage:\n  app [flags] src\n\nArguments:\n  src     the file to copy\n  EXTRA   an extra line\n\nFlags:\n  -h, --help      show help, then exit\n",
		},
		{
			"spec wins",
			opts(Arg("src", "the file to copy"), Help(Spec("SOURCE"))),
			"app an app\n\nUsage:\n  app SOURCE\n\nArguments:\n  src     the file to copy\n\nFlags:\n  -h, --help      show help, then exit\n",
		},
		{
			"column",
			opts(Arg("src", "the file to copy"), HelpColumn(12)),
			"app an app\n\nUsage:\n  app src\n\nArguments:\n  src       the file to copy\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd, err := NewCommand(append(test.opts, Usage("app", "an app"))...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s := cmd.UsageString(); s != test.exp {
				t.Errorf("expected:\n%q\ngot:\n%q", test.exp, s)
			}
		})
	}
}

// TestArgComps holds that completion completes each position by the type of
// its argument, and falls back to the CompFunc of the command.
func TestArgComps(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  []string
		dir  CompDirective
	}{
		{"file", ss(`cp//`), ss(`go//mod`), CompFilterFileExt},
		{"allowed values", ss(`cp//a.go//`), ss(`fast//slow`), CompNoFileComp | CompKeepOrder},
		{"allowed values prefix", ss(`cp//a.go//s`), ss(`slow`), CompNoFileComp | CompKeepOrder},
		{"dir", ss(`cp//a.go//fast//`), nil, CompFilterDirs},
		{"variadic dir", ss(`cp//a.go//fast//a//b//`), nil, CompFilterDirs},
		{"command func", ss(`two//`), ss(`one//two`), CompNoFileComp},
		{"arg func", ss(`two//x//`), ss(`red//green`), CompNoFileComp},
		{"arg func prefix", ss(`two//x//g`), ss(`green`), CompNoFileComp},
		{"past the last", ss(`two//x//y//`), nil, CompNoFileComp},
		{"string", ss(`str//`), nil, CompDefault},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root, err := NewCommand(
				Usage("app", "an app"),
				Sub(
					Usage("cp", "copy"),
					Arg("src", "the file", FileT, Ext(".go", "mod")),
					Arg("mode", "the mode", Valid("fast", "slow")),
					Arg("dirs", "the dirs", SliceT, Elem(DirT)),
				),
				Sub(
					Usage("two", "two args"),
					Arg("a", "a"),
					Arg("b", "b", CompFunc(func(*Context, string) ([]Completion, CompDirective, error) {
						return []Completion{{Name: "red"}, {Name: "green"}}, CompNoFileComp, nil
					})),
					CompFunc(func(*Context, string) ([]Completion, CompDirective, error) {
						return []Completion{{Name: "one"}, {Name: "two"}}, CompNoFileComp, nil
					}),
				),
				Sub(
					Usage("str", "a string arg"),
					Arg("s", "s"),
				),
			)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			ctx := &Context{
				Root: root,
				Args: append([]string{DefaultCompName}, test.args...),
			}
			comps, dir, err := ctx.Comps()
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
			if dir != test.dir {
				t.Errorf("expected %s, got: %s", test.dir, dir)
			}
		})
	}
}

// opts returns the options.
func opts(v ...Option) []Option {
	return v
}
