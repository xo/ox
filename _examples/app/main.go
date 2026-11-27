// Command app is an example of a program that ox builds from a struct and
// its doc comments, with ox.AppFrom.
//
// Run `go generate` after a change to a doc comment. It writes
// oxdoc_gen.go.
package main

//go:generate go run github.com/xo/ox/cmd/oxdoc

import (
	"context"
	"fmt"

	"github.com/xo/ox"
)

var (
	name    = "app"
	version = "0.0.0-dev"
)

func main() {
	ox.AppFrom(new(App), ox.AppNameVersion(name, version))
}

// App is the ox example app.
//
// Usage:
//
//	app [flags] [command]
//
// Banner:
//
//	An example of an app built with ox.AppFrom.
//
// Examples:
//
//	app -vv sub1 a b
//	app sub2 --all
type App struct {
	// verbosity, such as -vvv
	Verbose int `ox:",type:count,short:v"`
}

// Commands returns the sub commands of the app.
func (app *App) Commands() []any {
	return []any{
		&AppSub1Command{app: app},
		new(AppSub2Command),
	}
}

// AppSub1Command is the app sub1 command.
//
// Description:
//
//	print the args, with the verbosity of the app
//
// Usage:
//
//	app sub1 [ARG]...
type AppSub1Command struct {
	app *App
}

// Exec runs the command.
func (cmd *AppSub1Command) Exec(_ context.Context, args []string) error {
	fmt.Println("verbose:", cmd.app.Verbose, "args:", args)
	return nil
}

// AppSub2Command is the app sub2 command. It has no usage line, so its name
// comes from its type: sub2.
//
// Description:
//
//	print a flag
type AppSub2Command struct {
	// enable all, or none
	All bool `ox:",short:a"`
}

// Exec runs the command.
func (cmd *AppSub2Command) Exec() {
	fmt.Println("all:", cmd.All)
}
