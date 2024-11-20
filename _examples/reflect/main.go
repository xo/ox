// _examples/reflect/main.go
package main

import (
	"context"
	"fmt"
	"net/url"

	k "github.com/xo/kobra"
)

type args struct {
	Arg     string      `kobra:"an argument,short:a"`
	URL     *url.URL    `kobra:"a url,short:u,set:URLSet"`
	Ints    []int       `kobra:"a slice of ints,short:i"`
	Strings []string    `kobra:"a slice of strings,short:s"`
	Map     map[int]int `kobra:"a map of ints,short:m"`
	URLSet  bool
}

func main() {
	args := new(args)
	k.Run(
		context.Background(),
		run(args),
		k.Usage("reflect", "demonstrates using kobra's flags from"),
		k.FlagsFrom(&args),
	)
}

func run(args *args) func(context.Context, []string) error {
	return func(ctx context.Context, v []string) error {
		fmt.Println("arg:", args.Arg)
		if args.URL != nil {
			fmt.Println("u:", args.URL)
		} else {
			fmt.Println("u: not set")
		}
		fmt.Println("strings:", args.Strings)
		return nil
	}
}
