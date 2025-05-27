// _examples/embedded/main.go
package main

import (
	"fmt"

	"github.com/xo/ox"
)

func main() {
	var args struct {
		Verbose bool `ox:"enable verbose,short:v"`
		More    struct {
			Foo string `ox:"foo,short:f"`
		} `ox:""`
	}
	ox.Run(
		ox.Usage("embedded", "demonstrates using embebbed structs with ox's bind"),
		ox.Defaults(),
		ox.From(&args),
		ox.Exec(func() {
			fmt.Println("Verbose:", args.Verbose)
			fmt.Println("More.Foo:", args.More.Foo)
		}),
	)
}
