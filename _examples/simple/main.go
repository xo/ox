// _examples/simple/main.go
package main

import (
	"context"
	"fmt"
	"net/url"
	"reflect"

	k "github.com/xo/kobra"
	_ "github.com/xo/kobra/toml"
	_ "github.com/xo/kobra/yaml"
)

func main() {
	k.Run(
		context.Background(),
		run,
		k.Usage("simple", "a simple demo of the kobra api"),
		k.Version("0.0.0-dev"),
		k.Help(),
		k.Comp(),
		k.UserConfigFile(),
		k.Flags().
			Var("param-a", "parameter a", k.Short("a"), k.Key("", "param-a"), k.Key("yaml", "my_param_a"), k.Key("toml", "paramA")).
			Int("param-b", "parameter b", k.Short("b"), k.Default(125)).
			Slice("floats", "a slice of float64", k.Float64T, k.Short("f")).
			URL("url", "a url", k.Alias("my-url", "other url flag alias"), k.Short("u")).
			Count("verbose", "verbose", k.Short("v")),
		k.Sub(
			sub,
			k.Usage("sub", "a sub command"),
			k.Alias("subCommand", "a sub command alias"),
			k.Flags().
				Var("sub", "sub param").
				Slice("strings", "a slice of strings", k.Short("s")).
				Slice("urls", "a slice of URLs", k.URLT, k.Short("u")).
				Map("ints", "a map of integers", k.IntT, k.Short("i")),
			k.Args(0, 10),
		),
	)
}

func run(ctx context.Context, args []string) error {
	fmt.Println("run args:", args)

	// get param-a
	paramA := k.Get[string](ctx, "param-a")
	fmt.Println("paramA:", paramA)

	// convert param-b (int) into a string
	paramB := k.String(ctx, "param-b")
	fmt.Println("paramB:", paramB)

	// a slice
	floats := k.Slice[float64](ctx, "floats")
	fmt.Println("floats:", floats)

	// convert a slice's values to strings
	floatStrings := k.Slice[string](ctx, "floats")
	fmt.Println("floatStrings:", floatStrings)

	// sub param is not available in this command, as it was defined on a sub
	// command and not on the root command
	sub := k.Get[string](ctx, "sub")
	fmt.Println("sub:", sub)

	// a url
	if u := k.URL(ctx, "url"); u != nil {
		// NOTE: this is wrapped in a if block, because when no flag has been
		// NOTE: passed, tinygo's fmt.Println will panic with a *url.URL(nil),
		// NOTE: however Go's fmt.Println does not
		fmt.Println("url:", u)
		// url alternate
		urlAlt := k.Get[*url.URL](ctx, "url")
		fmt.Println("urlAlt:", urlAlt)
	}

	// verbose as its own type
	type Verbosity int64
	v := k.Get[Verbosity](ctx, "verbose")
	fmt.Println("verbosity:", v, reflect.TypeOf(v))

	return nil
}

func sub(ctx context.Context, args []string) error {
	fmt.Println("sub args:", args)

	// get param-a, as any parent's
	paramA := k.Get[string](ctx, "param-a")
	fmt.Println("paramA:", paramA)

	// the floats param is available, as this is a sub command
	floats := k.Slice[float64](ctx, "floats")
	fmt.Println("floats:", floats)

	// sub is available here
	sub := k.Get[string](ctx, "sub")
	fmt.Println("subParam:", sub)

	// get strings
	slice := k.Slice[string](ctx, "strings")
	fmt.Println("slice:", slice)

	// slice of URLs
	urls := k.Slice[*url.URL](ctx, "urls")
	fmt.Println("urls:", urls)

	// map of ints, converted to int64
	ints := k.Map[int64](ctx, "ints")
	fmt.Println("ints:", ints)
	return nil
}
