// _examples/bind/main.go
package main

import (
	"context"
	"fmt"
	"net/url"

	k "github.com/xo/kobra"
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
	k.Run(
		context.Background(),
		run,
		k.Usage("bind", "demonstrates using kobra's binds"),
		k.Flags().
			String("arg", "an arg", k.Bind(&arg)).
			URL("url", "a url", k.Short("u"), k.BindSet(&u, &urlSet)).
			Slice("int", "a slice of ints", k.Short("i"), k.Uint64T, k.Bind(&ints), k.Bind(&strings)).
			Map("map", "a map", k.Short("m"), k.Bind(&m), k.IntT, k.MapKey(k.IntT)),
	)
}
