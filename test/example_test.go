package test

// The examples in this file are the examples of docs/EXAMPLES.md that need
// the yaml, the toml or the hcl module. Change an example here and in
// docs/EXAMPLES.md at the same time.

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xo/ox"
	_ "github.com/xo/ox/hcl"
	"github.com/xo/ox/otx"
	_ "github.com/xo/ox/toml"
	_ "github.com/xo/ox/yaml"
)

// Example_yaml reads a YAML config file. A key that starts with $ is a YAML
// path. A default reads a key with ${yaml::key}, and it escapes the $ of a
// YAML path as \$, because a $ starts a variable.
func Example_yaml() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.yaml")
	data := "server:\n  port: 8080\nstore:\n  book:\n    - author: jon\n    - author: ken\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		// a program uses a path such as "$APPCONFIG/config.yaml"
		ox.ConfigFiles(path),
		ox.Flags().
			Int("port", "the port", ox.Key("yaml", "server.port")).
			String("first", "the first author", ox.Key("yaml", "$.store.book[0].author")).
			String("authors", "the authors", ox.Default(`${yaml::\$.store.book[*].author}`)).
			String("listen", "the address", ox.Default("localhost:${yaml::server.port}")),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Println("first:", otx.String(ctx, "first"))
			fmt.Println("authors:", otx.String(ctx, "authors"))
			fmt.Println("listen:", otx.String(ctx, "listen"))
		}),
		ox.Args(),
	)
	// Output:
	// port: 8080
	// first: jon
	// authors: jon,ken
	// listen: localhost:8080
}

// Example_yamlTag reads a YAML path in the default of a struct tag. The tag
// writes the \$ of the default as \\\\$, because Go and ox each remove one
// level of backslashes.
func Example_yamlTag() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.yaml")
	data := "server:\n  port: 8080\nstore:\n  book:\n    - author: jon\n    - author: ken\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	args := struct {
		Port    int    `ox:"the port,key:yaml|server.port"`
		Authors string `ox:"the authors,default:${yaml::\\\\$.store.book[*].author}"`
	}{}
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.From(&args),
		ox.Exec(func() {
			fmt.Println("port:", args.Port)
			fmt.Println("authors:", args.Authors)
		}),
		ox.Args(),
	)
	// Output:
	// port: 8080
	// authors: jon,ken
}

// Example_toml reads a TOML config file. A default reads a key of the TOML
// files with ${toml::key}, and a filter such as |upper changes the value.
func Example_toml() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.toml")
	data := "region = \"us\"\ntags = [\"a\", \"b\"]\n\n[server]\nport = 8080\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.Flags().
			Int("port", "the port", ox.Key("", "server.port")).
			Slice("tags", "the tags").
			String("region", "the region", ox.Default("${toml::region|upper}"), ox.NoConfig()),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Println("tags:", otx.Slice[string](ctx, "tags"))
			fmt.Println("region:", otx.String(ctx, "region"))
		}),
		ox.Args(),
	)
	// Output:
	// port: 8080
	// tags: [a b]
	// region: US
}

// Example_hcl reads an HCL config file. A block nests its body under its type
// and its labels, so the key of port is server.db.port.
func Example_hcl() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.hcl")
	data := "region = \"us\"\ntags = [\"a\", \"b\"]\n\nserver \"db\" {\n  port = 5432\n}\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.Flags().
			Int("port", "the port", ox.Key("", "server.db.port")).
			Slice("tags", "the tags").
			String("region", "the region", ox.Default("${hcl::region|upper}"), ox.NoConfig()),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Println("tags:", otx.Slice[string](ctx, "tags"))
			fmt.Println("region:", otx.String(ctx, "region"))
		}),
		ox.Args(),
	)
	// Output:
	// port: 5432
	// tags: [a b]
	// region: US
}
