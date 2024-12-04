// _examples/simple/main.go
package main

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"reflect"

	"github.com/xo/ox"
	"github.com/xo/ox/otx"
	_ "github.com/xo/ox/toml"
	_ "github.com/xo/ox/yaml"
)

func main() {
	ox.Run(
		ox.Exec(run),
		ox.Usage("simple", "a simple demo of the ox api"),
		ox.Defaults(),
		ox.UserConfigFile(),
		ox.Flags().
			Var("param-a", "parameter a", ox.Short("a"), ox.Key("", "param-a"), ox.Key("yaml", "my_param_a"), ox.Key("toml", "paramA")).
			Int("param-b", "parameter b", ox.Short("b"), ox.Default(125)).
			Slice("floats", "a slice of float64", ox.Float64T, ox.Short("f")).
			URL("url", "a url", ox.Aliases("my-url"), ox.Short("u")).
			Count("verbose", "verbose", ox.Short("v")),
		ox.Sub(
			ox.Exec(sub),
			ox.Usage("sub", "a sub command"),
			ox.Aliases("subCommand"),
			ox.Flags().
				Var("sub", "sub param").
				Slice("strings", "a slice of strings", ox.Short("s")).
				Slice("bigint", "a slice of big ints", ox.BigIntT, ox.Short("t")).
				Map("ints", "a map of integers", ox.IntT, ox.Short("i")),
			ox.ValidArgs(0, 10),
			ox.Sub(
				ox.Usage("dub", "a dub command"),
			),
		),
	)
}

func run(ctx context.Context, args []string) error {
	fmt.Println("run args:", args)

	// get param-a
	paramA := otx.Get[string](ctx, "param-a")
	fmt.Println("paramA:", paramA)

	// convert param-b (int) into a string
	paramB := otx.String(ctx, "param-b")
	fmt.Println("paramB:", paramB)

	// a slice
	floats := otx.Slice[float64](ctx, "floats")
	fmt.Println("floats:", floats)

	// convert a slice's values to strings
	floatStrings := otx.Slice[string](ctx, "floats")
	fmt.Println("floatStrings:", floatStrings)

	// sub param is not available in this command, as it was defined on a sub
	// command and not on the root command
	sub := otx.Get[string](ctx, "sub")
	fmt.Println("sub:", sub)

	// a url
	if u := otx.URL(ctx, "url"); u != nil {
		// NOTE: this is wrapped in a if block, because when no flag has been
		// NOTE: passed, tinygo's fmt.Println will panic with a *url.URL(nil),
		// NOTE: however Go's fmt.Println does not
		fmt.Println("url:", u)
		// url alternate
		urlAlt := otx.Get[*url.URL](ctx, "url")
		fmt.Println("urlAlt:", urlAlt)
	}

	// verbose as its own type
	type Verbosity int64
	v := otx.Get[Verbosity](ctx, "verbose")
	fmt.Println("verbosity:", v, reflect.TypeOf(v))

	return nil
}

func sub(ctx context.Context, args []string) error {
	fmt.Println("sub args:", args)

	// get param-a, as any parent's
	paramA := otx.Get[string](ctx, "param-a")
	fmt.Println("paramA:", paramA)

	// convert param-b (int) into a uint32
	paramB := otx.Uint32(ctx, "param-b")
	fmt.Println("paramB:", paramB)

	// the floats param is available, as this is a sub command
	floats := otx.Slice[float64](ctx, "floats")
	fmt.Println("floats:", floats)

	// sub is available here
	sub := otx.Get[string](ctx, "sub")
	fmt.Println("subParam:", sub)

	// get strings
	slice := otx.Slice[string](ctx, "strings")
	fmt.Println("slice:", slice)

	// slice of *big.Int
	bigint := otx.Slice[*big.Int](ctx, "bigint")
	fmt.Println("bigint:", bigint)

	// map of ints, converted to int64
	ints := otx.Map[string, int64](ctx, "ints")
	fmt.Println("ints:", ints)
	return nil
}
