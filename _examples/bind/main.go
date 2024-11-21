// _examples/bind/main.go
package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/xo/ox"
)

func main() {
	var (
		arg     string
		u       *url.URL
		urlSet  bool
		ints    []int
		strings []string
		m       map[int]int
	)
	run := func(ctx context.Context, args []string) error {
		fmt.Println("arg:", arg)
		if urlSet {
			fmt.Println("u:", u)
		} else {
			fmt.Println("u: not set")
		}
		fmt.Println("ints:", ints)
		fmt.Println("strings:", strings)
		fmt.Println("map:", m)
		return nil
	}
	ox.Run(
		context.Background(),
		run,
		ox.Usage("bind", "demonstrates using ox's binds"),
		ox.Flags().
			String("arg", "an arg", ox.Bind(&arg)).
			URL("url", "a url", ox.Short("u"), ox.BindSet(&u, &urlSet)).
			Slice("int", "a slice of ints", ox.Short("i"), ox.Uint64T, ox.Bind(&ints), ox.Bind(&strings)).
			Map("map", "a map", ox.Short("m"), ox.Bind(&m), ox.IntT, ox.MapKey(ox.IntT)),
	)
}
