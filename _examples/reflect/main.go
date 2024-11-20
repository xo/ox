// _examples/reflect/main.go
package main

import (
	"context"
	"fmt"
	"net/url"

	k "github.com/xo/kobra"
)

func main() {
	args := struct {
		Arg     string   `kobra:"an argument,short:a"`
		URL     *url.URL `kobra:"a url,short:u"`
		Strings []string `kobra:"a string,short:s"`
	}{}
	k.Run(context.Background(),
		func(ctx context.Context, v []string) error {
			fmt.Println("arg:", args.Arg)
			if args.URL != nil {
				fmt.Println("u:", args.URL)
			} else {
				fmt.Println("u: not set")
			}
			fmt.Println("strings:", args.Strings)
			return nil
		},
		k.Usage("reflect", "demonstrates using kobra's reflect"),
		k.FlagsFrom(&args),
	)
}
