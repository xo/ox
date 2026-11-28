// Package test holds the tests of ox that need more than one module.
package test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/xo/ox"
	_ "github.com/xo/ox/hcl"
	_ "github.com/xo/ox/ini"
	_ "github.com/xo/ox/toml"
	_ "github.com/xo/ox/yaml"
)

// TestConfigFormats reads the same settings from a YAML, a TOML, a JSON, an
// HCL and an INI file, one at a time.
func TestConfigFormats(t *testing.T) {
	dir := t.TempDir()
	files := map[string]string{
		"config.yaml": "port: 8080\nserver:\n  host: db\ntags:\n  - a,b\n  - c\nverbose: 2\n",
		"config.yml":  "port: 8080\nserver:\n  host: db\ntags: [\"a,b\", c]\nverbose: 2\n",
		"config.toml": "port = 8080\ntags = [\"a,b\", \"c\"]\nverbose = 2\n[server]\nhost = \"db\"\n",
		"config.json": `{"port": 8080, "server": {"host": "db"}, "tags": ["a,b", "c"], "verbose": 2}`,
		"config.hcl":  "port = 8080\ntags = [\"a,b\", \"c\"]\nverbose = 2\nserver {\n  host = \"db\"\n}\n",
		"config.ini":  "port = 8080\ntags = \"a,b\"\ntags = c\nverbose = 2\n[server]\nhost = db\n",
	}
	for name, data := range files {
		t.Run(name, func(t *testing.T) {
			path := write(t, dir, name, data)
			ctx := parse(t, nil,
				ox.ConfigFiles(path),
				ox.Flags().
					Int("port", "port").
					String("host", "host", ox.Key("", "server.host")).
					Slice("tags", "tags").
					Count("verbose", "verbosity"),
			)
			for flag, exp := range map[string]string{
				"port":    "8080",
				"host":    "db",
				"tags":    "[a,b c]",
				"verbose": "2",
			} {
				if v, ok := ctx.Vars[flag]; !ok || v.String() != exp {
					t.Errorf("expected %s to be %q, got: %v", flag, exp, v)
				}
			}
		})
	}
}

// TestConfigCascade reads files of three formats in order, so that a later
// file wins, and reads a key of one format in a default.
func TestConfigCascade(t *testing.T) {
	dir := t.TempDir()
	system := write(t, dir, "system.toml", "name = \"system\"\nlevel = \"info\"\nregion = \"us\"\n")
	user := write(t, dir, "user.yaml", "name: user\nlevel: debug\nstore:\n  book:\n    - author: jon\n    - author: ken\n")
	local := write(t, dir, "local.json", `{"name": "local"}`)
	ctx := parse(t, nil,
		ox.ConfigFiles(system, user, filepath.Join(dir, "missing.yaml"), local),
		ox.Flags().
			String("name", "name").
			String("level", "level").
			String("region", "region").
			String("authors", "authors", ox.Default(`${yaml::\$.store.book[*].author}`)).
			String("from-toml", "a key of the TOML file", ox.Default("${toml::region|upper}")),
	)
	for flag, exp := range map[string]string{
		"name":      "local",
		"level":     "debug",
		"region":    "us",
		"authors":   "jon,ken",
		"from-toml": "US",
	} {
		if v, ok := ctx.Vars[flag]; !ok || v.String() != exp {
			t.Errorf("expected %s to be %q, got: %v", flag, exp, v)
		}
	}
}

// parse parses args with opts and fails the test on an error.
func parse(t *testing.T, args []string, opts ...ox.Option) *ox.Context {
	t.Helper()
	ctx, err := ox.NewContext(append([]ox.Option{
		ox.Usage("app", "an app"),
		ox.Args(args...),
		ox.Exec(func(context.Context) error { return nil }),
	}, opts...)...)
	if err != nil {
		t.Fatal(err)
	}
	if err := ctx.Parse(); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return ctx
}

// write writes a file in dir and returns its path.
func write(t *testing.T, dir, name, data string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
