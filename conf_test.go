package ox

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestFill tests the order of the sources of a flag value: the command line,
// the environment, the config files, then the default.
func TestFill(t *testing.T) {
	dir := t.TempDir()
	first := writeFile(t, dir, "first.json", `{"port": 1, "name": "first", "verbose": 3, "tags": ["a,b", "c"],
		"labels": {"y": "2", "x": "1"}, "server": {"host": "db"}, "debug": true}`)
	second := writeFile(t, dir, "second.json", `{"name": "second"}`)
	tests := []struct {
		name string
		env  map[string]string
		args []string
		opts []Option
		flag string
		exp  string
	}{
		{"command line wins", map[string]string{"APP_PORT": "2"}, []string{"--port", "3"}, nil, "port", "3"},
		{"environment wins over file", map[string]string{"APP_PORT": "2"}, nil, nil, "port", "2"},
		{"file wins over default", nil, nil, nil, "port", "1"},
		{"default, and a missing file is skipped", nil, nil, []Option{ConfigFiles(filepath.Join(dir, "missing.json"))}, "fallback", "4"},
		{"later file wins", nil, nil, []Option{ConfigFiles(second)}, "name", "second"},
		{"no env", map[string]string{"APP_HOST": "env"}, nil, nil, "host", "db"},
		{"env var", map[string]string{"CUSTOM": "custom"}, nil, nil, "custom", "custom"},
		{"no config", nil, nil, nil, "local", "local default"},
		{"count from file assigns", nil, nil, nil, "verbose", "3"},
		{"count from env assigns", map[string]string{"APP_VERBOSE": "2"}, nil, nil, "verbose", "2"},
		{"bool from env", map[string]string{"APP_DEBUG": "false"}, nil, nil, "debug", "false"},
		{"bool from file", nil, nil, nil, "debug", "true"},
		{"list sets each element whole", nil, nil, nil, "tags", "[a,b c]"},
		{"env string splits", map[string]string{"APP_TAGS": "x,y"}, nil, nil, "tags", "[x y]"},
		{"command line replaces default", nil, []string{"--items", "c"}, nil, "items", "[c]"},
		{"default slice", nil, nil, nil, "items", "[a b]"},
		{"map in key order", nil, nil, nil, "labels", "[x:1 y:2]"},
		{"default reads file", nil, nil, nil, "from-file", "db:1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for k, v := range test.env {
				t.Setenv(k, v)
			}
			ctx := runFill(t, test.args, append([]Option{
				EnvPrefix("APP"),
				ConfigFiles(first),
				Flags().
					Int("port", "port", Default(4)).
					Int("fallback", "fallback", Default(4)).
					String("name", "name").
					String("host", "host", NoEnv(), Key("", "server.host")).
					String("custom", "custom", EnvVar("CUSTOM")).
					String("local", "local", NoConfig(), Default("local default")).
					Count("verbose", "verbosity", Short("v")).
					Bool("debug", "debug").
					Slice("tags", "tags").
					Slice("items", "items", Default("a,b")).
					Map("labels", "labels").
					String("from-file", "from file", Default("${json::server.host}:${json::port}")),
			}, test.opts...)...)
			v, ok := ctx.Vars[test.flag]
			if !ok {
				t.Fatalf("expected a value for %s", test.flag)
			}
			if s := v.String(); s != test.exp {
				t.Errorf("expected %s to be %q, got: %q", test.flag, test.exp, s)
			}
			if set := v.WasSet(); set != (len(test.args) != 0 && strings.TrimLeft(test.args[0], "-") == test.flag) {
				t.Errorf("expected WasSet to be true only for the command line, got: %t", set)
			}
		})
	}
}

// TestFillConfigFlag tests a flag that names the config file.
func TestFillConfigFlag(t *testing.T) {
	dir := t.TempDir()
	listed := writeFile(t, dir, "listed.json", `{"name": "listed"}`)
	chosen := writeFile(t, dir, "chosen.json", `{"name": "chosen"}`)
	opts := func() []Option {
		return []Option{
			ConfigFiles(listed),
			Flags().
				String("config", "config file", ConfigFlag(), Default("${env::TEST_OX_CONFIG}")).
				String("name", "name"),
		}
	}
	t.Run("the list", func(t *testing.T) {
		if s := runFill(t, nil, opts()...).Vars["name"].String(); s != "listed" {
			t.Errorf("expected listed, got: %q", s)
		}
	})
	t.Run("the flag replaces the list", func(t *testing.T) {
		if s := runFill(t, []string{"--config", chosen}, opts()...).Vars["name"].String(); s != "chosen" {
			t.Errorf("expected chosen, got: %q", s)
		}
	})
	t.Run("the default of the flag reads the environment", func(t *testing.T) {
		t.Setenv("TEST_OX_CONFIG", chosen)
		if s := runFill(t, nil, opts()...).Vars["name"].String(); s != "chosen" {
			t.Errorf("expected chosen, got: %q", s)
		}
	})
	t.Run("a file that the flag names must exist", func(t *testing.T) {
		err := fillError(t, []string{"--config", filepath.Join(dir, "missing.json")}, opts()...)
		if !errors.Is(err, ErrConfigNotFound) {
			t.Errorf("expected ErrConfigNotFound, got: %v", err)
		}
	})
}

// TestFillErrors tests the errors that filling a flag returns.
func TestFillErrors(t *testing.T) {
	dir := t.TempDir()
	list := writeFile(t, dir, "list.json", `{"name": ["a", "b"]}`)
	bad := writeFile(t, dir, "bad.json", `{"name": `)
	unknown := writeFile(t, dir, "config.conf", `name = x`)
	tests := []struct {
		name string
		opts []Option
		exp  error
	}{
		{"a list for a string flag", []Option{ConfigFiles(list)}, ErrInvalidValue},
		{"no config file exists", []Option{ConfigFiles(filepath.Join(dir, "missing.json")), ConfigRequired()}, ErrConfigNotFound},
		{"an extension with no format", []Option{ConfigFiles(unknown)}, ErrUnknownLoader},
		{"a file that cannot be decoded", []Option{ConfigFiles(bad)}, ErrInvalidConfig},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := fillError(t, nil, append(test.opts, Flags().String("name", "name"))...)
			if !errors.Is(err, test.exp) {
				t.Errorf("expected %v, got: %v", test.exp, err)
			}
			if !strings.Contains(err.Error(), "--name") && !errors.Is(err, ErrConfigNotFound) && !errors.Is(err, ErrUnknownLoader) && !errors.Is(err, ErrInvalidConfig) {
				t.Errorf("expected the error to name the flag, got: %v", err)
			}
		})
	}
	t.Run("a file that cannot be decoded", func(t *testing.T) {
		err := fillError(t, nil, ConfigFiles(bad), Flags().String("name", "name"))
		if err == nil || !strings.Contains(err.Error(), bad) {
			t.Errorf("expected an error that names %s, got: %v", bad, err)
		}
	})
}

// TestFillHook tests that a file calls a hook only when the program binds the
// hook by name.
func TestFillHook(t *testing.T) {
	dir := t.TempDir()
	file := writeFile(t, dir, "config.json", `{"hook": "from file", "bound": "from file"}`)
	var calls []string
	f := func(name string) func(string) {
		return func(s string) {
			calls = append(calls, name+"="+s)
		}
	}
	runFill(t, nil,
		ConfigFiles(file),
		Flags().
			Var("hook", "hook", HookT, Default(f("hook"))).
			Var("bound", "bound hook", HookT, Default(f("bound")), Key("json", "bound")),
	)
	if len(calls) != 1 || calls[0] != "bound=from file" {
		t.Errorf("expected only the bound hook to be called, got: %q", calls)
	}
}

// TestFillSubCommand tests that a sub command fills its own flags and the
// flags of its parents, and that the key of a flag of a sub command starts
// with the path of the command.
func TestFillSubCommand(t *testing.T) {
	tests := []struct {
		name string
		data string
		exp  string
	}{
		{"nested", `{"root": "r", "sub": {"deep": {"child": "c"}}}`, "c"},
		{"dotted", `{"root": "r", "sub.deep.child": "c"}`, "c"},
		{"not top level", `{"root": "r", "child": "c"}`, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file := writeFile(t, t.TempDir(), "config.json", test.data)
			ctx := runFill(t, []string{"sub", "deep"},
				ConfigFiles(file),
				Flags().String("root", "root"),
				Sub(
					Usage("sub", "a sub command"),
					Sub(
						Usage("deep", "a sub command of sub"),
						Flags().String("child", "child"),
						Exec(func() {}),
					),
				),
			)
			if v, ok := ctx.Vars["root"]; !ok || v.String() != "r" {
				t.Errorf("expected root to be %q, got: %v", "r", v)
			}
			var s string
			if v, ok := ctx.Vars["child"]; ok {
				s = v.String()
			}
			if s != test.exp {
				t.Errorf("expected child to be %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestFillKeyPath tests the key of a field of a nested struct of a sub
// command, and that a key of Key is the whole key.
func TestFillKeyPath(t *testing.T) {
	file := writeFile(t, t.TempDir(), "config.json", `{
		"server": {"host": "db"},
		"migrate": {"steps": 3, "source": {"dir": "up", "url": "file:///x"}},
		"absolute": "a"
	}`)
	var root struct {
		Server struct {
			Host string `ox:"the host"`
		} `ox:""`
	}
	var sub struct {
		Steps  int `ox:"the steps"`
		Source struct {
			Dir string `ox:"the folder"`
			URL string `ox:"the url,name:url"`
		} `ox:""`
		Other string `ox:"a key of its own,key:absolute"`
	}
	ctx := runFill(t, []string{"migrate"},
		ConfigFiles(file),
		From(&root),
		Sub(
			Usage("migrate", "migrate"),
			From(&sub),
			Exec(func() {}),
		),
	)
	for name, exp := range map[string]string{
		"server-host": "db",
		"steps":       "3",
		"source-dir":  "up",
		"url":         "file:///x",
		"other":       "a",
	} {
		if v, ok := ctx.Vars[name]; !ok || v.String() != exp {
			t.Errorf("expected %s to be %q, got: %v", name, exp, v)
		}
	}
}

// TestMapConfig tests a lookup in a tree of maps and slices.
func TestMapConfig(t *testing.T) {
	m := MapConfig{
		"a.b": "literal",
		"a":   map[string]any{"b": "nested", "list": []any{"x", map[string]any{"c": "y"}}},
	}
	tests := []struct {
		key string
		exp any
		ok  bool
	}{
		{"a.b", "literal", true},
		{"a.list.0", "x", true},
		{"a.list.1.c", "y", true},
		{"a.list.2", nil, false},
		{"a.missing", nil, false},
		{"a.b.c", nil, false},
	}
	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			v, ok, err := m.Lookup(test.key)
			if err != nil || ok != test.ok || v != test.exp {
				t.Errorf("expected %v %t, got: %v %t %v", test.exp, test.ok, v, ok, err)
			}
		})
	}
}

// runFill parses args with opts and fails the test on an error.
func runFill(t *testing.T, args []string, opts ...Option) *Context {
	t.Helper()
	ctx, err := newFillContext(args, opts...)
	if err == nil {
		err = ctx.Parse()
	}
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return ctx
}

// fillError parses args with opts and returns the error.
func fillError(t *testing.T, args []string, opts ...Option) error {
	t.Helper()
	ctx, err := newFillContext(args, opts...)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if err := ctx.Parse(); err != nil {
		return err
	}
	t.Fatal("expected an error")
	return nil
}

// newFillContext creates a context for args with opts.
func newFillContext(args []string, opts ...Option) (*Context, error) {
	return NewContext(append([]Option{
		Usage("app", "an app"),
		Args(args...),
		Exec(func(context.Context) error { return nil }),
	}, opts...)...)
}

// writeFile writes a file in dir and returns its path.
func writeFile(t *testing.T, dir, name, data string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
