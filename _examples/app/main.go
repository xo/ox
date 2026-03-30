package main

import (
	"context"

	"github.com/xo/ox"
)

var (
	name    = "app"
	version = "0.0.0-dev"
)

func main() {
	ox.RunAppContext(
		context.Background(),
		NewAppCommand(),
		ox.AppNameVersion(name, version),
	)
}

// AppCommand is the ox example app command.
//
// Usage:
//
//	app [options] [command]
//
// Banner:
//
// Examples:
//
// Footer:
type AppCommand struct {
	Verbose int  `ox:"verbosity,type:count,short:v"`
	All     bool `ox:"enable all,short:a"`
}

// NewAppCommand creates a new app command.
func NewAppCommand() *AppCommand {
	return new(AppCommand)
}

// Sub returns the sub commands for the app.
func (a *AppCommand) Sub() []any {
	return []any{
		NewSub1Command(),
		NewSub2Command(),
	}
}

// Sub1Command is the app sub1 command.
type Sub1Command struct{}

// NewSub1Command creates a app sub1 command.
func NewSub1Command() *Sub1Command {
	return new(Sub1Command)
}

// Sub2Command is the app sub2 command.
type Sub2Command struct{}

// NewSub2Command creates a app sub2 command.
func NewSub2Command() *Sub2Command {
	return new(Sub2Command)
}
