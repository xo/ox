package main

import (
	"context"

	"github.com/xo/ox"
)

func main() {
	ox.RunContext(
		context.Background(),
		ox.Defaults(),
		ox.Usage("helm", "runs a command"),
		ox.Flags().
			String("config", "config file").
			Int("int", "an int", ox.Short("i")),
		ox.Sub(
			ox.Usage("sub1", "a sub command"),
			ox.Flags().
				String("my-string", "my string"),
		),
		ox.Sub(
			ox.Usage("sub2", "sub2 command"),
		),
	)
}
