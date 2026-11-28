package ox

import (
	"bytes"
	"context"
	"errors"
	"slices"
	"testing"
)

// TestNegateBools holds that NegateBools gives each plain bool flag of the
// command tree a negated form, which sets the flag to false and takes no
// value (D47).
func TestNegateBools(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"the negated form", ss(`--no-debug`), "app [] [debug=false]"},
		{"the flag", ss(`--debug`), "app [] [debug=true]"},
		{"the negated form after the flag", ss(`--debug//--no-debug`), "app [] [debug=false]"},
		{"the flag after the negated form", ss(`--no-debug//--debug`), "app [] [debug=true]"},
		{"no value of true", ss(`--no-debug=true`), "--no-debug: invalid value: the negated form of --debug takes no value"},
		{"no value of false", ss(`--no-debug=false`), "--no-debug: invalid value: the negated form of --debug takes no value"},
		{"no empty value", ss(`--no-debug=`), "--no-debug: invalid value: the negated form of --debug takes no value"},
		{"not the next arg", ss(`--no-debug//false`), `app ["false"] [debug=false]`},
		{"no negated form of an alias", ss(`--no-dbg`), "--no-dbg: unknown flag"},
		{"no negated form of a short name", ss(`--no-d`), "--no-d: unknown flag"},
		{"no short negated form", ss(`-no-d`), "-n: unknown flag"},
		{"a flag that starts with the prefix", ss(`--no-psqlrc`), "app [] [no-psqlrc=true]"},
		{"no second negated form", ss(`--no-no-psqlrc`), "--no-no-psqlrc: unknown flag"},
		{"not a flag with Negatable false", ss(`--no-cache`), "--no-cache: unknown flag"},
		{"not a count flag", ss(`--no-verbose`), "--no-verbose: unknown flag"},
		{"not a string flag", ss(`--no-name`), "--no-name: unknown flag"},
		{"not a hook flag", ss(`--no-help`), "--no-help: unknown flag"},
		{"not a bool flag with an optional value", ss(`--no-mode`), "--no-mode: unknown flag"},
		{"a hidden flag", ss(`--no-secret`), "app [] [secret=false]"},
		{"a local flag", ss(`--no-root-only`), "app [] [root-only=false]"},
		{"an inherited flag", ss(`sub//--no-debug`), "sub [] [debug=false]"},
		{"a flag of a sub command", ss(`sub//--no-color`), "sub [] [color=false]"},
		{"no flag of a sub command in the root", ss(`--no-color`), "--no-color: unknown flag"},
		{"no local flag of the root in a sub command", ss(`sub//--no-root-only`), "--no-root-only: unknown flag"},
		{"a flag of a sub command that hides a flag of the root", ss(`sub//--no-quiet`), "--no-quiet: unknown flag"},
		{"the flag of the root that a flag of a sub command hides", ss(`--no-quiet`), "app [] [quiet=false]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := parseResult(
				t,
				test.args,
				Usage("app", ""),
				NegateBools(),
				Flags().
					Bool("debug", "", Short("d"), Aliases("dbg")).
					Bool("no-psqlrc", "").
					Bool("cache", "", Negatable(false)).
					Count("verbose", "", Short("v")).
					String("name", "").
					Hook("help", "", func() {}).
					Bool("mode", "", OptionalValue("MODE")).
					Bool("secret", "", Hidden(true)).
					Bool("root-only", "", Local()).
					Bool("quiet", ""),
				Sub(
					Usage("sub", ""),
					Flags().
						Bool("color", "").
						String("quiet", ""),
				),
			)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestNegateBoolsDeprecated holds that the negated form of a deprecated flag
// writes the warning of the flag, with the name that the command line gave
// (D47).
func TestNegateBoolsDeprecated(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := RunError(
		context.Background(),
		Usage("app", "an app"),
		NegateBools(),
		Flags().Bool("old", "", Deprecated(true)),
		Exec(func() {}),
		Args("--no-old"),
		Pipe(nil, &stdout, &stderr),
	)
	switch exp := "warning: --no-old is deprecated\n"; {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case stderr.String() != exp:
		t.Errorf("expected %q, got: %q", exp, stderr.String())
	}
}

// TestNegateBoolsPrefixes holds that each prefix of NegateBools adds a
// negated form, and that Negatable gives one flag a negated form (D47).
func TestNegateBoolsPrefixes(t *testing.T) {
	tests := []struct {
		name string
		args []string
		opts []Option
		exp  string
	}{
		{"off", ss(`--no-debug`), nil, "--no-debug: unknown flag"},
		{"the default prefix", ss(`--no-debug`), []Option{NegateBools()}, "app [] [debug=false]"},
		{"a prefix", ss(`--without-debug`), []Option{NegateBools("without")}, "app [] [debug=false]"},
		{"not the default prefix with a prefix", ss(`--no-debug`), []Option{NegateBools("without")}, "--no-debug: unknown flag"},
		{"the first of two prefixes", ss(`--no-debug`), []Option{NegateBools("no", "without")}, "app [] [debug=false]"},
		{"the second of two prefixes", ss(`--without-debug`), []Option{NegateBools("no", "without")}, "app [] [debug=false]"},
		{"a prefix with a hyphen", ss(`--do-not-debug`), []Option{NegateBools("do-not")}, "app [] [debug=false]"},
		{"no negated form of a flag that starts with a second prefix", ss(`--no-without-tls`), []Option{NegateBools("no", "without")}, "--no-without-tls: unknown flag"},
		{"Negatable without NegateBools", ss(`--no-cache`), nil, "app [] [cache=false]"},
		{"no negated form of a flag without Negatable", ss(`--no-debug`), nil, "--no-debug: unknown flag"},
		{"Negatable takes the prefix of the root", ss(`--without-cache`), []Option{NegateBools("without")}, "app [] [cache=false]"},
		{"Negatable in a sub command", ss(`sub//--no-cache`), nil, "sub [] [cache=false]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := append([]Option{
				Usage("app", ""),
				Flags().
					Bool("debug", "").
					Bool("cache", "", Negatable(true)).
					Bool("without-tls", ""),
				Sub(Usage("sub", "")),
			}, test.opts...)
			if s := parseResult(t, test.args, opts...); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestNegateBoolsErrors holds that a bad prefix, and a negated form with the
// name of another flag, are errors when the command tree is built (D47).
func TestNegateBoolsErrors(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  error
	}{
		{"an empty prefix", []Option{NegateBools("")}, ErrInvalidValue},
		{"a hyphen at the start", []Option{NegateBools("-no")}, ErrInvalidValue},
		{"a hyphen at the end", []Option{NegateBools("no-")}, ErrInvalidValue},
		{"a space", []Option{NegateBools("n o")}, ErrInvalidValue},
		{"an equal", []Option{NegateBools("n=o")}, ErrInvalidValue},
		{"a prefix twice", []Option{NegateBools("no", "no")}, ErrInvalidValue},
		{"in a sub command", []Option{Sub(Usage("sub", ""), NegateBools())}, ErrCanOnlyBeUsedWithRootCommand},
		{"a flag", []Option{NegateBools(), Flags().Bool("cache", "").String("no-cache", "")}, ErrDuplicateFlag},
		{"a bool flag", []Option{NegateBools(), Flags().Bool("cache", "").Bool("no-cache", "")}, ErrDuplicateFlag},
		{"an alias", []Option{NegateBools(), Flags().Bool("cache", "").String("x", "", Aliases("no-cache"))}, ErrDuplicateFlag},
		{"an alias of the flag", []Option{NegateBools(), Flags().Bool("cache", "", Aliases("no-cache"))}, ErrDuplicateFlag},
		{"a hidden alias", []Option{NegateBools(), Flags().Bool("cache", "").String("x", "", HiddenAliases("no-cache"))}, ErrDuplicateFlag},
		{"a deprecated alias", []Option{NegateBools(), Flags().Bool("cache", "").String("x", "", DeprecatedAliases("no-cache"))}, ErrDuplicateFlag},
		{"a flag of the root", []Option{NegateBools(), Flags().String("no-cache", ""), Sub(Usage("sub", ""), Flags().Bool("cache", ""))}, ErrDuplicateFlag},
		{"a flag of a sub command", []Option{NegateBools(), Flags().Bool("cache", ""), Sub(Usage("sub", ""), Flags().String("no-cache", ""))}, ErrDuplicateFlag},
		{"a sub command before NegateBools", []Option{Sub(Usage("sub", ""), Flags().Bool("cache", "").String("no-cache", "")), NegateBools()}, ErrDuplicateFlag},
		{"a sub command of a sub command", []Option{NegateBools(), Sub(Usage("sub", ""), Sub(Usage("child", ""), Flags().Bool("cache", "").String("no-cache", "")))}, ErrDuplicateFlag},
		{"a second prefix", []Option{NegateBools("no", "skip"), Flags().Bool("cache", "").String("skip-cache", "")}, ErrDuplicateFlag},
		{"Negatable without NegateBools", []Option{Flags().Bool("cache", "", Negatable(true)).String("no-cache", "")}, ErrDuplicateFlag},
		{"no error with Negatable false", []Option{NegateBools(), Flags().Bool("cache", "", Negatable(false)).String("no-cache", "")}, nil},
		{"no error without NegateBools", []Option{Flags().Bool("cache", "").String("no-cache", "")}, nil},
		{"no error with a local flag of the root", []Option{NegateBools(), Flags().Bool("cache", "", Local()), Sub(Usage("sub", ""), Flags().String("no-cache", ""))}, nil},
		{"no error in two sub commands", []Option{NegateBools(), Sub(Usage("one", ""), Flags().Bool("cache", "")), Sub(Usage("two", ""), Flags().String("no-cache", ""))}, nil},
		{"no error with a prefix with a hyphen", []Option{NegateBools("do-not")}, nil},
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

// TestNegatableErrors holds that a flag that is not a plain bool flag
// refuses Negatable(true), and takes Negatable(false) (D47).
func TestNegatableErrors(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  error
	}{
		{"a string flag", []Option{Negatable(true)}, ErrAppliedToInvalidType},
		{"a count flag", []Option{CountT, Negatable(true)}, ErrAppliedToInvalidType},
		{"a hook flag", []Option{Hook(func(context.Context) error { return nil }), Negatable(true)}, ErrAppliedToInvalidType},
		{"a bool flag with an optional value", []Option{BoolT, OptionalValue("WHEN"), Negatable(true)}, ErrAppliedToInvalidType},
		{"a bool flag", []Option{BoolT, Negatable(true)}, nil},
		{"a string flag with Negatable false", []Option{Negatable(false)}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewFlag("cache", "", test.opts...)
			switch {
			case test.exp == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.exp):
				t.Errorf("expected error %v, got: %v", test.exp, err)
			}
		})
	}
}

// TestNegateBoolsErrorMessage holds that the error of a negated form with the
// name of another flag names the command and both flags, and says what to do
// (D47).
func TestNegateBoolsErrorMessage(t *testing.T) {
	_, err := NewCommand(
		Usage("app", ""),
		NegateBools(),
		Sub(Usage("sub", ""), Flags().Bool("cache", "").String("no-cache", "")),
	)
	exp := "Sub (*ox.Command): command sub: duplicate flag: --no-cache is the negated form of --cache and the name of another flag (use Negatable(false))"
	if err == nil || err.Error() != exp {
		t.Errorf("expected %q, got: %v", exp, err)
	}
}

// TestNegatableTag holds that the tag options negatable and negatable:false
// set Negatable (D47).
func TestNegatableTag(t *testing.T) {
	type args struct {
		Debug bool `ox:"debug,negatable,default:true"`
		Cache bool `ox:"cache,negatable:false"`
		Color bool `ox:"color"`
	}
	tests := []struct {
		name string
		args []string
		opts []Option
		exp  string
	}{
		{"negatable", ss(`--no-debug`), nil, "app [] [debug=false]"},
		{"no negated form without the tag", ss(`--no-color`), nil, "--no-color: unknown flag"},
		{"negatable with NegateBools", ss(`--no-debug`), []Option{NegateBools()}, "app [] [debug=false]"},
		{"negatable false with NegateBools", ss(`--no-cache`), []Option{NegateBools()}, "--no-cache: unknown flag"},
		{"no tag with NegateBools", ss(`--no-color`), []Option{NegateBools()}, "app [] [color=false]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var v args
			opts := append([]Option{Usage("app", ""), From(&v)}, test.opts...)
			if s := parseResult(t, test.args, opts...); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
	t.Run("bind", func(t *testing.T) {
		var v args
		if err := RunError(context.Background(), Usage("app", ""), From(&v), Exec(func() {}), Args("--no-debug")); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if v.Debug {
			t.Errorf("expected Debug to be false")
		}
	})
	t.Run("a bad value", func(t *testing.T) {
		_, err := FlagsFrom(&struct {
			Debug bool `ox:"debug,negatable:maybe"`
		}{})
		if !errors.Is(err, ErrInvalidConversion) {
			t.Errorf("expected %v, got: %v", ErrInvalidConversion, err)
		}
	})
	t.Run("a string field", func(t *testing.T) {
		_, err := FlagsFrom(&struct {
			Name string `ox:"name,negatable"`
		}{})
		if !errors.Is(err, ErrAppliedToInvalidType) {
			t.Errorf("expected %v, got: %v", ErrAppliedToInvalidType, err)
		}
	})
}

// TestNegateBoolsAbbrev holds that with Abbrev, a unique prefix of a negated
// form gives the negated form, and that a prefix of a flag and of its
// negated form is ambiguous (D47).
func TestNegateBoolsAbbrev(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
	}{
		{"a prefix of a negated form", ss(`--no-deb`), "app [] [debug=false]"},
		{"a prefix of the flag", ss(`--deb`), "app [] [debug=true]"},
		{"a prefix of another negated form", ss(`--no-c`), "app [] [color=false]"},
		{"a prefix of a flag that starts with the prefix", ss(`--no-p`), "app [] [no-psqlrc=true]"},
		{"an ambiguous prefix", ss(`--no-`), "--no-: ambiguous flag: --no-color, --no-debug, --no-nox, --no-psqlrc"},
		{"a prefix of a flag and of its negated form", ss(`--no`), "--no: ambiguous flag: --no-color, --no-debug, --no-nox, --no-psqlrc, --nox"},
		{"no value after a prefix of a negated form", ss(`--no-deb=true`), "--no-deb: invalid value: the negated form of --debug takes no value"},
		{"no prefix of the negated form of a hidden flag", ss(`--no-sec`), "--no-sec: unknown flag"},
		{"the negated form of a hidden flag", ss(`--no-secret`), "app [] [secret=false]"},
		{"a prefix of an inherited negated form", ss(`sub//--no-deb`), "sub [] [debug=false]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := parseResult(
				t,
				test.args,
				Usage("app", ""),
				Abbrev(),
				NegateBools(),
				Flags().
					Bool("debug", "").
					String("deploy", "").
					Bool("color", "").
					Bool("no-psqlrc", "").
					Bool("nox", "").
					Bool("secret", "", Hidden(true)),
				Sub(Usage("sub", "")),
			)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestNegateBoolsEnv holds that the environment sets the flag itself, and
// that no environment variable has the name of a negated form (D47).
func TestNegateBoolsEnv(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		args []string
		exp  string
	}{
		{"the flag from the environment", map[string]string{"APP_DEBUG": "true"}, nil, "true"},
		{"no negated form from the environment", map[string]string{"APP_NO_DEBUG": "true"}, nil, ""},
		{"the negated form wins over the environment", map[string]string{"APP_DEBUG": "true"}, ss(`--no-debug`), "false"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for k, v := range test.env {
				t.Setenv(k, v)
			}
			ctx, err := NewContext(
				Usage("app", ""),
				EnvPrefix("APP"),
				NegateBools(),
				Flags().Bool("debug", ""),
				Args(test.args...),
			)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if err := ctx.Parse(); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			var s string
			if v, ok := ctx.Vars["debug"]; ok {
				s, _ = v.Get()
			}
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestNegateBoolsHelp holds that help writes a flag with a negated form as
// `--[no-]debug`, with the first prefix, in both layouts (D47).
func TestNegateBoolsHelp(t *testing.T) {
	flags := func() *FlagSet {
		return Flags().
			Bool("debug", "turn on debug output", Short("d")).
			Bool("color", "use color", Aliases("colour")).
			Bool("no-psqlrc", "do not read the startup file").
			Bool("cache", "use the cache", Negatable(false)).
			String("name", "the name").
			Bool("secret", "a secret", Hidden(true))
	}
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{
			"default layout",
			[]Option{NegateBools(), flags()},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command\n\nFlags:\n  -d, --[no-]debug   turn on debug output\n      --[no-]color   use color\n      --no-psqlrc    do not read the startup file\n      --cache        use the cache\n      --name string  the name\n  -h, --help         show help, then exit\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"column",
			[]Option{NegateBools(), flags(), HelpColumn(24)},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help                  show help for any command\n  sub                   a sub command\n\nFlags:\n  -d, --[no-]debug      turn on debug output\n      --[no-]color, --colour\n                        use color\n      --no-psqlrc       do not read the startup file\n      --cache           use the cache\n      --name string     the name\n  -h, --help            show help, then exit\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"the first prefix",
			[]Option{NegateBools("without", "no"), Flags().Bool("debug", "turn on debug output")},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command\n\nFlags:\n      --[without-]debug  turn on debug output\n  -h, --help             show help, then exit\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"Negatable without NegateBools",
			[]Option{Flags().Bool("debug", "turn on debug output", Negatable(true)).Bool("color", "use color")},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command\n\nFlags:\n      --[no-]debug  turn on debug output\n      --color       use color\n  -h, --help        show help, then exit\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
		{
			"off",
			[]Option{flags()},
			"app an app\n\nUsage:\n  app [flags] [command] [args]\n\nAvailable Commands:\n  help    show help for any command\n  sub     a sub command\n\nFlags:\n  -d, --debug        turn on debug output\n      --color        use color\n      --no-psqlrc    do not read the startup file\n      --cache        use the cache\n      --name string  the name\n  -h, --help         show help, then exit\n\nUse \"app [command] --help\" for more information about a command.\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := []Option{
				Usage("app", "an app"),
				Help(),
				Sub(Usage("sub", "a sub command")),
			}
			if s := runHelp(t, append(opts, test.opts...)...); s != test.exp {
				t.Errorf("expected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", test.exp, s, test.exp, s)
			}
		})
	}
	t.Run("inherited flags", func(t *testing.T) {
		var buf bytes.Buffer
		err := RunError(
			context.Background(),
			Usage("app", "an app"),
			NegateBools(),
			Flags().Bool("debug", "turn on debug output"),
			Sub(Usage("sub", "a sub command"), Flags().Bool("color", "use color")),
			Help(InheritedFlags(true)),
			Args("sub", "--help"),
			Pipe(nil, &buf, &buf),
		)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		s := buf.String()
		exp := "sub a sub command\n\nUsage:\n  app sub [flags] [args]\n\nFlags:\n      --[no-]color  use color\n  -h, --help        show help, then exit\n\nGlobal Flags:\n  --[no-]debug  turn on debug output\n"
		if s != exp {
			t.Errorf("expected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", exp, s, exp, s)
		}
	})
}

// TestNegateBoolsComps holds that completion offers the negated form of a
// flag after the flag, and never the negated form of a hidden flag, of an
// alias or of a flag with Negatable false (D47).
func TestNegateBoolsComps(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  []string
		dir  CompDirective
	}{
		{"each flag", ss(`--`), ss(`--debug//--no-debug//--color//--colour//--no-color//--no-psqlrc//--cache//--name`), CompNoFileComp | CompKeepOrder},
		{"a prefix of a negated form", ss(`--no-d`), ss(`--no-debug`), CompNoFileComp | CompKeepOrder},
		{"the prefix", ss(`--no`), ss(`--no-debug//--no-color//--no-psqlrc`), CompNoFileComp | CompKeepOrder},
		{"a correction of a negated form", ss(`--no-debgu`), ss(`--no-debug`), CompNoFileComp | CompKeepOrder},
		{"no value after a negated form", ss(`--no-debug//`), ss(`sub`), CompNoFileComp | CompKeepOrder},
		{"no value of a negated form", ss(`--no-debug=`), nil, CompDefault},
		{"an inherited negated form", ss(`sub//--no-d`), ss(`--no-debug`), CompNoFileComp | CompKeepOrder},
		{"no negated form of a hidden flag", ss(`--no-s`), nil, CompDefault},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root, err := NewCommand(
				Usage("app", "an app"),
				NegateBools(),
				Flags().
					Bool("debug", "debug", Short("d")).
					Bool("color", "color", Aliases("colour")).
					Bool("no-psqlrc", "no psqlrc").
					Bool("cache", "cache", Negatable(false)).
					String("name", "the name").
					Bool("secret", "secret", Hidden(true)),
				Sub(Usage("sub", "a sub command")),
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
