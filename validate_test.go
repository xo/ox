package ox

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// TestValidateFlags tests the required flags and the rules between flags. A
// value counts when it comes from the command line, the environment or a
// config file, and a default does not count.
func TestValidateFlags(t *testing.T) {
	dir := t.TempDir()
	config := writeFile(t, dir, "config.json", `{"name": "file", "user": "file"}`)
	required := func() []Option {
		return []Option{
			Flags().
				String("name", "the name", Required(), Short("n")).
				String("other", "other").
				String("token", "the token", Required(), Default("secret")),
		}
	}
	formats := func(opts ...Option) []Option {
		return append([]Option{
			Flags().
				Bool("json", "json").
				Bool("yaml", "yaml", Aliases("yml")).
				Bool("toml", "toml", Default(true)).
				String("user", "user").
				String("pass", "pass").
				String("file", "file").
				String("url", "url"),
		}, opts...)
	}
	tests := []struct {
		name string
		env  map[string]string
		args []string
		opts []Option
		err  error
		msg  string
	}{
		{"required from the command line", nil, []string{"--name", "a", "-n", "b", "--token", "t"}, required(), nil, ""},
		{"required from the environment", map[string]string{"APP_NAME": "a", "APP_TOKEN": "t"}, nil, required(), nil, ""},
		{"required from a config file", map[string]string{"APP_TOKEN": "t"}, nil, append(required(), ConfigFiles(config)), nil, ""},
		{"required without a value", nil, []string{"--other", "x"}, required(), ErrRequiredFlag, "command app: required flag: --name, --token"},
		{"a default is not a value", nil, []string{"-n", "a"}, required(), ErrRequiredFlag, "command app: required flag: --token"},
		{"the help flag works", nil, []string{"--help"}, append(required(), Defaults()), ErrExit, ""},
		{"the version flag works", nil, []string{"--version"}, append(required(), Defaults()), ErrExit, ""},
		{"only one of none", nil, nil, formats(OnlyOneOf("json", "yaml")), nil, ""},
		{"only one of one", nil, []string{"--json"}, formats(OnlyOneOf("json", "yaml")), nil, ""},
		{"only one of two", nil, []string{"--json", "--yml"}, formats(OnlyOneOf("json", "yml", "toml")), ErrOnlyOneOf, "command app: only one of the flags can be set: --json, --yaml"},
		{"only one of the command line and the environment", map[string]string{"APP_YAML": "true"}, []string{"--json"}, formats(OnlyOneOf("json", "yaml")), ErrOnlyOneOf, "command app: only one of the flags can be set: --json, --yaml"},
		{"only one of a default and a value", nil, []string{"--json"}, formats(OnlyOneOf("json", "toml")), nil, ""},
		{"all or none of none", nil, nil, formats(AllOrNoneOf("user", "pass")), nil, ""},
		{"all or none of all", nil, []string{"--user", "u", "--pass", "p"}, formats(AllOrNoneOf("user", "pass")), nil, ""},
		{"all or none of one", nil, []string{"--user", "u"}, formats(AllOrNoneOf("user", "pass")), ErrAllOrNoneOf, "command app: all or none of the flags must be set: --user, --pass (missing --pass)"},
		{"all or none of a file and the command line", nil, []string{"--pass", "p"}, formats(AllOrNoneOf("user", "pass"), ConfigFiles(config)), nil, ""},
		{"at least one of none", nil, nil, formats(AtLeastOneOf("file", "url")), ErrAtLeastOneOf, "command app: at least one of the flags must be set: --file, --url"},
		{"at least one of one", nil, []string{"--url", "u"}, formats(AtLeastOneOf("file", "url")), nil, ""},
		{"at least one of a default", nil, nil, formats(AtLeastOneOf("json", "toml")), ErrAtLeastOneOf, "command app: at least one of the flags must be set: --json, --toml"},
		{"the required flag comes first", nil, []string{"--json", "--yaml"}, []Option{
			Flags().
				String("name", "the name", Required()).
				Bool("json", "json").
				Bool("yaml", "yaml"),
			OnlyOneOf("json", "yaml"),
		}, ErrRequiredFlag, "command app: required flag: --name"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for k, v := range test.env {
				t.Setenv(k, v)
			}
			err := runValidate(t, test.args, append([]Option{
				Usage("app", "an app"),
				EnvPrefix("APP"),
				Exec(func(context.Context) error { return nil }),
			}, test.opts...)...)
			switch {
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.msg != "" && err.Error() != test.msg:
				t.Errorf("expected message %q, got: %q", test.msg, err.Error())
			}
		})
	}
}

// TestValidateFlagsCommands tests the required flags and the rules between
// flags of a command tree. The flags of a parent count for its sub commands,
// and the commands that [Defaults] adds do not check them.
func TestValidateFlagsCommands(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  error
	}{
		{"the root without an exec func writes help", nil, nil},
		{"a sub command checks a required flag of its parent", []string{"sub"}, ErrRequiredFlag},
		{"a sub command takes the flag of its parent", []string{"sub", "--name", "a"}, nil},
		{"a sub command checks the rules of its parent", []string{"sub", "--name", "a", "--json", "--yaml"}, ErrOnlyOneOf},
		{"a sub command checks its own rules", []string{"sub", "--name", "a"}, nil},
		{"the help command", []string{"help", "sub"}, nil},
		{"the version command", []string{"version"}, nil},
		{"the completion command", []string{"completion", "bash"}, nil},
		{"the help flag of a sub command", []string{"sub", "--help"}, ErrExit},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := runValidate(t, test.args,
				Usage("app", "an app"),
				Defaults(),
				Flags().
					String("name", "the name", Required()).
					Bool("json", "json").
					Bool("yaml", "yaml"),
				OnlyOneOf("json", "yaml"),
				Sub(
					Usage("sub", "a sub command"),
					Flags().
						String("file", "file").
						String("url", "url"),
					AllOrNoneOf("file", "url"),
					Exec(func(context.Context) error { return nil }),
				),
			)
			if !errors.Is(err, test.err) {
				t.Errorf("expected error %v, got: %v", test.err, err)
			}
		})
	}
}

// TestValidateFlagsFrom tests the struct tag options `required`,
// `onlyoneof`, `allornoneof` and `atleastoneof`.
func TestValidateFlagsFrom(t *testing.T) {
	type args struct {
		Name string `ox:"the name,required"`
		JSON bool   `ox:"json,onlyoneof:format"`
		YAML bool   `ox:"yaml,onlyoneof:format,atleastoneof:output"`
		Out  string `ox:"out,atleastoneof:output"`
		User string `ox:"user,allornoneof:auth"`
		Pass string `ox:"pass,allornoneof:auth"`
	}
	tests := []struct {
		name string
		args []string
		err  error
	}{
		{"valid", []string{"--name", "a", "--yaml", "--user", "u", "--pass", "p"}, nil},
		{"required", []string{"--yaml"}, ErrRequiredFlag},
		{"only one of", []string{"--name", "a", "--json", "--yaml"}, ErrOnlyOneOf},
		{"at least one of", []string{"--name", "a", "--json"}, ErrAtLeastOneOf},
		{"all or none of", []string{"--name", "a", "--out", "o", "--pass", "p"}, ErrAllOrNoneOf},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var v args
			err := runValidate(t, test.args,
				Usage("app", "an app"),
				From(&v),
				Exec(func(context.Context) error { return nil }),
			)
			if !errors.Is(err, test.err) {
				t.Errorf("expected error %v, got: %v", test.err, err)
			}
		})
	}
	t.Run("groups", func(t *testing.T) {
		cmd, err := NewCommand(Usage("app", "an app"), From(&args{}))
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		exp := []FlagGroup{
			{RuleOnlyOneOf, []string{"json", "yaml"}},
			{RuleAtLeastOneOf, []string{"yaml", "out"}},
			{RuleAllOrNoneOf, []string{"user", "pass"}},
		}
		if s, e := fmt.Sprint(cmd.FlagGroups), fmt.Sprint(exp); s != e {
			t.Errorf("expected %s, got: %s", e, s)
		}
	})
}

// TestFlagGroupErrors tests the errors of a group that names a flag that
// does not exist, and of a struct tag without a group name.
func TestFlagGroupErrors(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		err  error
		msg  string
	}{
		{"unknown flag", []Option{Flags().Bool("json", "json"), OnlyOneOf("json", "xml")}, ErrUnknownFlag, "OnlyOneOf: --xml: unknown flag"},
		{"no group name", []Option{From(&struct {
			JSON bool `ox:"json,onlyoneof"`
		}{})}, ErrInvalidValue, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewCommand(append([]Option{Usage("app", "an app")}, test.opts...)...)
			switch {
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.msg != "" && err.Error() != test.msg:
				t.Errorf("expected message %q, got: %q", test.msg, err.Error())
			}
		})
	}
}

// TestHelpRequired tests that help marks a required flag, in place of its
// default.
func TestHelpRequired(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{"default layout", nil, "  --name NAME   the name (required)\n"},
		{"column layout", []Option{HelpColumn(20)}, "  --name NAME       the name (required)\n"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd, err := NewCommand(append([]Option{
				Usage("app", "an app"),
				Flags().
					String("name", "the name", Required(), Default("x"), Spec("NAME")).
					String("other", "other", Default("y"), Spec("NAME")),
			}, test.opts...)...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			s := cmd.UsageString()
			if !strings.Contains(s, test.exp) {
				t.Errorf("expected %q in:\n%s", test.exp, s)
			}
			if strings.Contains(s, "(default: x)") {
				t.Errorf("expected no default for a required flag in:\n%s", s)
			}
		})
	}
}

// runValidate creates a context with opts, then parses args and validates
// them. It returns the first error.
func runValidate(t *testing.T, args []string, opts ...Option) error {
	t.Helper()
	var buf strings.Builder
	ctx, err := NewContext(append(opts, Args(args...), Pipe(nil, &buf, &buf))...)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if err := ctx.Parse(); err != nil {
		return err
	}
	return ctx.Validate()
}
