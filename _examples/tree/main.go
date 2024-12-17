package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xo/ox"
)

var name = "tree"

func main() {
	f, err := os.OpenFile(name+".txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "---------------------------------------------")
	for _, s := range os.Environ() {
		fmt.Fprintln(f, s)
	}
	fmt.Fprintln(f)
	fmt.Fprintln(f, os.Args)
	fmt.Fprintln(f)
	ox.RunContext(
		context.Background(),
		ox.Defaults(),
		ox.Usage(name, "runs a command"),
		ox.Flags().
			String("config", "config file").
			String("my-string", "my string on "+name).
			Int("int", "an int", ox.Short("i")),
		ox.Sub(
			ox.Usage("sub1", "a sub command"),
			ox.Flags().
				String("my-string", "my string on sub1"),
		),
		ox.Sub(
			ox.Usage("sub2", "sub2 command"),
		),
	)
}
