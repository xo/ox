package ox

import (
	"bytes"
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestComps(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		exp   []string
		dir   CompDirective
		isErr bool
	}{
		// commands
		{"commands", ss(``), ss(`shell//serve//completion//help`), CompNoFileComp | CompKeepOrder, false},
		{"command prefix", ss(`sh`), ss(`shell`), CompNoFileComp | CompKeepOrder, false},
		{"command correction", ss(`shlel`), ss(`shell`), CompNoFileComp | CompKeepOrder, false},
		{"deprecated command", ss(`legacy`), nil, CompDefault, false},
		// the value after a flag
		{"value after a long flag", ss(`--config//`), nil, CompDefault, false},
		{"value after a short flag", ss(`-c//`), nil, CompDefault, false},
		{"value after a short cluster", ss(`-vc//`), nil, CompDefault, false},
		{"value in a short cluster", ss(`-cfile//`), ss(`shell//serve//completion//help`), CompNoFileComp | CompKeepOrder, false},
		{"value with a long equal", ss(`--config=`), nil, CompDefault, false},
		{"value with a long equal and a prefix", ss(`--config=fo`), nil, CompDefault, false},
		{"value of an unknown flag", ss(`--nope=fo`), nil, CompDefault, false},
		{"arg after a bool flag", ss(`--verbose//`), ss(`shell//serve//completion//help`), CompNoFileComp | CompKeepOrder, false},
		{"arg after a bool short flag", ss(`-v//sh`), ss(`shell`), CompNoFileComp | CompKeepOrder, false},
		// allowed values
		{"allowed values", ss(`--format//`), ss(`json//yaml//table`), CompNoFileComp | CompKeepOrder, false},
		{"allowed values after a short flag", ss(`-f//y`), ss(`yaml`), CompNoFileComp | CompKeepOrder, false},
		{"allowed values with an equal", ss(`--format=j`), ss(`json`), CompNoFileComp | CompKeepOrder, false},
		{"allowed values in a short cluster with an equal", ss(`-vf=t`), ss(`table`), CompNoFileComp | CompKeepOrder, false},
		{"allowed value correction", ss(`--format//jsno`), ss(`json`), CompNoFileComp | CompKeepOrder, false},
		{"no allowed value", ss(`--format//zzzzzz`), nil, CompNoFileComp, false},
		{"allowed values after a separator", ss(`--tag//a,`), ss(`a,a//a,b//a,c`), CompNoFileComp | CompKeepOrder, false},
		{"allowed values after a separator and a prefix", ss(`--tag//a,b,c`), ss(`a,b,c`), CompNoFileComp | CompKeepOrder, false},
		// valid args
		{"valid args", ss(`shell//`), ss(`bash//zsh//fish`), CompNoFileComp | CompKeepOrder, false},
		{"valid args prefix", ss(`shell//z`), ss(`zsh`), CompNoFileComp | CompKeepOrder, false},
		{"valid args correction", ss(`shell//basj`), ss(`bash`), CompNoFileComp | CompKeepOrder, false},
		{"valid args at the second position", ss(`shell//bash//f`), ss(`fish`), CompNoFileComp | CompKeepOrder, false},
		{"valid args after the maximum", ss(`shell//bash//zsh//`), nil, CompNoFileComp, false},
		{"valid args after a flag", ss(`shell//--verbose//`), ss(`bash//zsh//fish`), CompNoFileComp | CompKeepOrder, false},
		// completion funcs
		{"flag func", ss(`--color//`), ss(`red//green//blue`), CompNoFileComp, false},
		{"flag func prefix", ss(`--color//g`), ss(`green`), CompNoFileComp, false},
		{"flag func no correction", ss(`--color//rde`), nil, CompNoFileComp, false},
		{"flag func no filter", ss(`--any//q`), ss(`x//y`), CompDefault, false},
		{"flag func with an equal", ss(`--color=b`), ss(`blue`), CompNoFileComp, false},
		{"command func", ss(`--config//a.yaml//serve//one//`), ss(`one:a.yaml:x//one:a.yaml:y`), CompKeepOrder, false},
		{"command func prefix", ss(`serve//one:a`), nil, CompKeepOrder, false},
		{"command func error", ss(`serve//err`), nil, CompError, true},
		// path types
		{"path", ss(`--path//`), nil, CompDefault, false},
		{"file with extensions", ss(`--file//`), ss(`yaml//yml`), CompFilterFileExt, false},
		{"file with extensions and an equal", ss(`--file=a`), ss(`yaml//yml`), CompFilterFileExt, false},
		{"file", ss(`--any-file//`), nil, CompDefault, false},
		{"dir", ss(`--dir//`), nil, CompFilterDirs, false},
		// a bad earlier value
		{"bad value", ss(`--format//bad//sh`), ss(`shell`), CompNoFileComp | CompKeepOrder, false},
		{"bad value before a command", ss(`--int//bad//shell//z`), ss(`zsh`), CompNoFileComp | CompKeepOrder, false},
		// the end of the flags
		{"after the end of the flags", ss(`--//sh`), nil, CompDefault, false},
		{"flag after the end of the flags", ss(`shell//--//-`), nil, CompNoFileComp, false},
		{"valid args after the end of the flags", ss(`shell//--//`), ss(`bash//zsh//fish`), CompNoFileComp | CompKeepOrder, false},
		// flag names
		{"flag names", ss(`--no-`), ss(`--no-init//--no-plain//--no-psqlrc//--no-usqlrc`), CompNoFileComp | CompKeepOrder, false},
		{"hidden aliases", ss(`--va`), ss(`--var//--variable`), CompNoFileComp | CompKeepOrder, false},
		{"deprecated alias", ss(`--no-r`), nil, CompDefault, false},
		{"deprecated flag", ss(`--ol`), nil, CompDefault, false},
		{"hidden flag", ss(`--secre`), nil, CompDefault, false},
		{"flag correction", ss(`--verbsoe`), ss(`--verbose`), CompNoFileComp | CompKeepOrder, false},
		{"hidden alias correction", ss(`--varible`), ss(`--variable`), CompNoFileComp | CompKeepOrder, false},
		{"short flags", ss(`-X`), ss(`-X`), CompNoFileComp | CompKeepOrder, false},
		{"deprecated alias parses", ss(`--no-rc//sh`), ss(`shell`), CompNoFileComp | CompKeepOrder, false},
		{"hidden alias takes a value", ss(`--var//`), nil, CompDefault, false},
		// the help command completes a command path
		{"help", ss(`help//`), ss(`shell//serve//completion//help`), CompNoFileComp | CompKeepOrder, false},
		{"help prefix", ss(`help//se`), ss(`serve`), CompNoFileComp | CompKeepOrder, false},
		{"help correction", ss(`help//shlel`), ss(`shell`), CompNoFileComp | CompKeepOrder, false},
		{"help of a leaf", ss(`help//shell//`), nil, CompNoFileComp | CompKeepOrder, false},
		{"help of a path", ss(`help//completion//`), ss(`bash//fish//powershell//zsh`), CompNoFileComp | CompKeepOrder, false},
		// errors
		{"no word", nil, nil, CompError, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := &Context{
				Root: testCompCommand(t),
				Args: append([]string{DefaultCompName}, test.args...),
			}
			if test.args == nil {
				ctx.Args = []string{DefaultCompName}
			}
			comps, dir, err := ctx.Comps()
			switch {
			case test.isErr && err == nil:
				t.Fatalf("expected an error")
			case !test.isErr && err != nil:
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

func TestDefaultComp(t *testing.T) {
	tests := []struct {
		args []string
		env  string
		exp  string
		err  string
	}{
		{ss(`__complete//--format//y`), "", "yaml\n:36\n", ""},
		{ss(`__complete//--no-i`), "", "--no-init\tno init\n:36\n", ""},
		{ss(`__completeNoDesc//--no-i`), "", "--no-init\n:36\n", ""},
		{ss(`__complete//--any//`), "", "x\ny\n:0\n", ""},
		{ss(`__complete//serve//err`), "", ":1\n", "error: bad word\n"},
		{ss(`__complete//--size//s`), "", "_activeHelp_ pick a size\nsmall\n:4\n", ""},
		{ss(`__complete//--size//s`), "0", "small\n:4\n", ""},
		{ss(`__complete//--size//s`), "1", "_activeHelp_ pick a size\nsmall\n:4\n", ""},
	}
	for _, test := range tests {
		t.Run(strings.Join(test.args, " ")+" "+test.env, func(t *testing.T) {
			t.Setenv("MY_APP_ACTIVE_HELP", test.env)
			var stdout, stderr bytes.Buffer
			ctx := &Context{
				Root:   testCompCommand(t),
				Args:   test.args,
				Stdout: &stdout,
				Stderr: &stderr,
				Comp:   DefaultComp,
			}
			if err := ctx.Parse(); !errors.Is(err, ErrExit) {
				t.Fatalf("expected ErrExit, got: %v", err)
			}
			if s := stdout.String(); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
			if s := stderr.String(); s != test.err {
				t.Errorf("expected %q on standard error, got: %q", test.err, s)
			}
		})
	}
}

func TestHiddenAliases(t *testing.T) {
	tests := []struct {
		args []string
		exp  string
	}{
		{ss(`--no-init`), "true"},
		{ss(`--no-psqlrc`), "true"},
		{ss(`--no-rc`), "true"},
		{ss(`-X`), "true"},
		{ss(`--var//a=b`), "a=b"},
		{ss(`--variable=c=d`), "c=d"},
	}
	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			ctx := &Context{
				Root: testCompCommand(t),
				Args: test.args,
				Vars: make(Vars),
				Continue: func(*Command, error) bool {
					return false
				},
			}
			if err := ctx.Parse(); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			name := "no-init"
			if strings.Contains(test.args[0], "var") {
				name = "set"
			}
			v, ok := ctx.Vars[name]
			if !ok {
				t.Fatalf("expected %s to be set", name)
			}
			if s, _ := v.Get(); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

func TestHiddenAliasesTag(t *testing.T) {
	args := struct {
		Set string `ox:"set a variable,hiddenaliases:var|variable,deprecatedaliases:v1"`
	}{}
	flags, err := FlagsFrom(&args)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if exp := ss(`var//variable`); !slices.Equal(flags[0].HiddenAliases, exp) {
		t.Errorf("expected %q, got: %q", exp, flags[0].HiddenAliases)
	}
	if exp := ss(`v1`); !slices.Equal(flags[0].DeprecatedAliases, exp) {
		t.Errorf("expected %q, got: %q", exp, flags[0].DeprecatedAliases)
	}
}

// testCompCommand returns the command tree that the completion tests use.
func testCompCommand(t *testing.T) *Command {
	t.Helper()
	colors := func(*Context, string) ([]Completion, CompDirective, error) {
		return []Completion{
			NewCompletion("red", "", 0),
			NewCompletion("green", "", 0),
			NewCompletion("blue", "", 0),
		}, CompNoFileComp, nil
	}
	anything := func(*Context, string) ([]Completion, CompDirective, error) {
		return []Completion{
			NewCompletion("x", "", 0),
			NewCompletion("y", "", 0),
		}, CompNoFilter, nil
	}
	sizes := func(*Context, string) ([]Completion, CompDirective, error) {
		return []Completion{
			NewActiveHelp("pick a size"),
			NewCompletion("small", "", 0),
			NewCompletion("large", "", 0),
		}, CompNoFileComp, nil
	}
	serve := func(ctx *Context, toComplete string) ([]Completion, CompDirective, error) {
		if toComplete == "err" {
			return nil, CompDefault, errors.New("bad word")
		}
		var config string
		if v, ok := ctx.Vars["config"]; ok {
			config, _ = v.Get()
		}
		prefix := strings.Join(ctx.Args, "+") + ":" + config + ":"
		return []Completion{
			NewCompletion(prefix+"x", "", 0),
			NewCompletion(prefix+"y", "", 0),
		}, CompKeepOrder, nil
	}
	cmd, err := NewCommand(
		Usage("my-app", "an app"),
		Flags().
			String("config", "config file", Short("c")).
			Bool("verbose", "verbose", Short("v")).
			Int("int", "an int").
			String("format", "output format", Short("f"), Valid("json", "yaml", "table")).
			Slice("tag", "tags", Valid("a", "b", "c"), Split(",")).
			Bool("no-init", "no init", Short("X"), Aliases("no-plain"), HiddenAliases("no-psqlrc", "no-usqlrc"), DeprecatedAliases("no-rc")).
			String("set", "set a variable", Short("s"), HiddenAliases("var", "variable")).
			String("old", "old", Deprecated(true)).
			String("secret", "secret", Hidden(true)).
			String("color", "a color", CompFunc(colors)).
			String("any", "anything", CompFunc(anything)).
			String("size", "a size", CompFunc(sizes)).
			Path("path", "a path").
			File("file", "a file", Ext(".yaml", ".yml")).
			File("any-file", "any file").
			Dir("dir", "a directory"),
		Sub(Usage("shell", "run a shell"), ValidArgs(0, 2, "bash", "zsh", "fish")),
		Sub(Usage("serve", "serve"), CompFunc(serve)),
		Sub(Usage("legacy", "legacy"), Deprecated(true)),
		Comp(),
		Help(),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return cmd
}

// TestCompFoldAndDist holds that DefaultFoldCase, DefaultMaxDist and a
// negative MaxDist change how completion matches a name (B10).
func TestCompFoldAndDist(t *testing.T) {
	tests := []struct {
		name    string
		fold    bool
		dist    int
		cmdDist int
		args    []string
		exp     []string
	}{
		{"fold", true, 2, 0, ss(`SH`), ss(`shell`)},
		{"fold a flag", true, 2, 0, ss(`--VERB`), ss(`--verbose`)},
		{"no fold", false, 2, 0, ss(`SH`), nil},
		{"no fold for a flag", false, 2, 0, ss(`--VERB`), nil},
		{"no fold with a prefix", false, 2, 0, ss(`sh`), ss(`shell`)},
		{"correction", true, 2, 0, ss(`shlel`), ss(`shell`)},
		{"no distance", true, 0, 0, ss(`shlel`), nil},
		{"no distance for the command", true, 2, -1, ss(`shlel`), nil},
		{"no distance for a flag", true, 2, -1, ss(`--verbsoe`), nil},
		{"distance for the command", true, 0, 2, ss(`shlel`), ss(`shell`)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fold, dist := DefaultFoldCase, DefaultMaxDist
			DefaultFoldCase, DefaultMaxDist = test.fold, test.dist
			t.Cleanup(func() {
				DefaultFoldCase, DefaultMaxDist = fold, dist
			})
			root := testCompCommand(t)
			if test.cmdDist != 0 {
				root.Help.(*CommandHelp).MaxDist = test.cmdDist
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

// TestSuggestFold holds that DefaultFoldCase changes the suggestion of a
// command name.
func TestSuggestFold(t *testing.T) {
	tests := []struct {
		name string
		fold bool
		arg  string
		exp  bool
	}{
		{"fold", true, "SHELL", true},
		{"no fold", false, "SHELL", false},
		{"no fold near", false, "shel", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fold := DefaultFoldCase
			DefaultFoldCase = test.fold
			t.Cleanup(func() {
				DefaultFoldCase = fold
			})
			err := testCompCommand(t).Suggest(test.arg)
			var e *SuggestionError
			if ok := errors.As(err, &e); ok != test.exp {
				t.Errorf("expected a suggestion %t, got: %v", test.exp, err)
			}
		})
	}
}
