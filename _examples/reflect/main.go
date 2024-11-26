// _examples/reflect/main.go
package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/xo/ox"
)

type args struct {
	Arg     string      `ox:"an argument,short:a"`
	URL     []*url.URL  `ox:"a url,short:u,set:URLSet"`
	URLSet  bool        ``
	Ints    []int       `ox:"a slice of ints,short:i"`
	Strings []string    `ox:"a slice of strings,short:s"`
	Map     map[int]int `ox:"a map of ints,short:m"`
	Other   []*url.URL  `ox:"a slice of urls,short:z"`
}

func main() {
	args := new(args)
	ox.Run(
		ox.Exec(run(args)),
		ox.Usage("reflect", "demonstrates using ox's From with struct tags"),
		ox.From(args),
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
		fmt.Println("ints:", args.Ints)
		fmt.Println("strings:", args.Strings)
		fmt.Println("map:", args.Map)
		fmt.Println("other:", args.Other)
		return nil
	}
}
