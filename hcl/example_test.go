package hcl_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xo/ox"
	_ "github.com/xo/ox/hcl"
	"github.com/xo/ox/otx"
)

// config is the HCL config file of the examples.
const config = `verbose = true

server {
  host = "db.example.com"
  port = 5432
}

migrate {
  steps = 3
  source {
    dir = "./migrations"
  }
}
`

// Example_build reads a HCL config file into flags that a program builds by
// hand. A flag of the root reads the key with its name, such as verbose. A
// flag of a sub command reads the key under the path of the command, so the
// flag --steps of "db migrate" reads migrate.steps. ox.Key gives a flag a
// whole key of its own, such as server.host. A default reads a key with
// ${hcl::key}.
func Example_build() {
	dir, err := os.MkdirTemp("", "db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.hcl")
	if err := os.WriteFile(path, []byte(config), 0o600); err != nil {
		fmt.Println(err)
		return
	}
	ox.Run(
		ox.Usage("db", "a database tool"),
		// a program uses a path such as "$APPCONFIG/config.hcl"
		ox.ConfigFiles(path),
		ox.Flags().
			Bool("verbose", "write more").
			String("host", "the host", ox.Key("", "server.host")).
			Int("port", "the port", ox.Key("", "server.port")).
			String("url", "the url", ox.Default("postgres://${hcl::server.host}:${hcl::server.port}/app")),
		ox.Sub(
			ox.Usage("migrate", "migrate the database"),
			ox.Flags().
				Int("steps", "the number of steps").
				String("dir", "the folder of the migrations", ox.Key("", "migrate.source.dir")),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("verbose:", otx.Bool(ctx, "verbose"))
				fmt.Println("server:", otx.String(ctx, "host"), otx.Int(ctx, "port"))
				fmt.Println("url:", otx.String(ctx, "url"))
				fmt.Println("steps:", otx.Int(ctx, "steps"))
				fmt.Println("dir:", otx.String(ctx, "dir"))
			}),
		),
		ox.Args("migrate"),
	)
	// Output:
	// verbose: true
	// server: db.example.com 5432
	// url: postgres://db.example.com:5432/app
	// steps: 3
	// dir: ./migrations
}

// App is the root command of Example_reflect.
type App struct {
	Verbose bool `ox:"write more"`
	// the flags --server-host and --server-port read server.host and
	// server.port
	Server struct {
		Host string `ox:"the host"`
		Port int    `ox:"the port"`
	} `ox:""`
}

// Commands returns the sub commands of the app.
func (app *App) Commands() []any {
	return []any{&Migrate{app: app}}
}

// Migrate is the migrate command of Example_reflect.
type Migrate struct {
	app *App
	// the flag --steps reads migrate.steps
	Steps int `ox:"the number of steps"`
	// the flag --source-dir reads migrate.source.dir
	Source struct {
		Dir string `ox:"the folder of the migrations"`
	} `ox:""`
}

// Exec runs the command.
func (m *Migrate) Exec() {
	fmt.Println("verbose:", m.app.Verbose)
	fmt.Println("server:", m.app.Server.Host, m.app.Server.Port)
	fmt.Println("steps:", m.Steps)
	fmt.Println("dir:", m.Source.Dir)
}

// Example_reflect reads a HCL config file into a struct and the struct of
// its sub command, with ox.App. Each field reads the key of its path, under
// the path of its command: the field Server.Port of the root reads
// server.port, and the field Source.Dir of the command "db migrate" reads
// migrate.source.dir. No field needs a key of its own.
func Example_reflect() {
	dir, err := os.MkdirTemp("", "db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.hcl")
	if err := os.WriteFile(path, []byte(config), 0o600); err != nil {
		fmt.Println(err)
		return
	}
	ox.Run(
		ox.Usage("db", "a database tool"),
		ox.App(new(App)),
		ox.ConfigFiles(path),
		ox.Args("migrate"),
	)
	// Output:
	// verbose: true
	// server: db.example.com 5432
	// steps: 3
	// dir: ./migrations
}
