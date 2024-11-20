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
		strings []string
		ints    []int
		m       map[string]string
	)
	k.Run(context.Background(),
		func(ctx context.Context, args []string) error {
			fmt.Println("arg:", arg)
			if urlSet {
				fmt.Println("u:", u)
			} else {
				fmt.Println("u: not set")
			}
			fmt.Println("strings:", strings)
			fmt.Println("ints:", ints)
			fmt.Println("map:", m)
			return nil
		},
		k.Usage("bind", "demonstrates using kobra's binds"),
		k.Flags().
			String("arg", "an arg", k.Bind(&arg)).
			URL("url", "a url", k.Short("u"), k.BindSet(&u, &urlSet)).
			Slice("str", "a string", k.Short("s"), k.Bind(&strings), k.Bind(&ints)).
			Map("map", "a map", k.Short("m"), k.Bind(&m)),
	)
}
