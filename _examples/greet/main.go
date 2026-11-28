// Command greet says hello. It is the example in the README of ox.
package main

import (
	"context"
	"fmt"

	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

func main() {
	ox.Run(
		ox.Usage("greet", "say hello"),
		ox.Defaults(),
		ox.Flags().
			String("name", "the name to greet", ox.Short("n"), ox.Default("world")).
			Int("times", "the number of greetings", ox.Short("t"), ox.Default(1)),
		ox.ValidArgs(0, 0),
		ox.Exec(run),
	)
}

func run(ctx context.Context) error {
	for range otx.Int(ctx, "times") {
		fmt.Println("hello", otx.String(ctx, "name"))
	}
	return nil
}
