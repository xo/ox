// Command paint is the struct example in the README of ox. ox.AppFrom builds
// the program from the struct: each field is a flag, and the Exec method
// runs the program.
package main

import (
	"fmt"
	"log/slog"

	"github.com/xo/ox"
)

// Paint paints a thing.
type Paint struct {
	Thing string     `ox:"the thing to paint,short:t,default:fence"`
	Coats int        `ox:"the number of coats,short:c,default:2"`
	Log   slog.Level `ox:"the log level,short:l,default:info"`
}

// Exec runs the program.
func (p *Paint) Exec() {
	fmt.Printf("painting the %s with %d coats (log %v)\n", p.Thing, p.Coats, p.Log)
}

func main() {
	ox.AppFrom(
		new(Paint),
		ox.Usage("paint", "paint a thing"),
		ox.ValidArgs(0, 0),
	)
}
