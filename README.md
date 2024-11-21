# xo/ox

`xo/ox` is a Go and TinyGo package for command-line argument and flag parsing,
designed for [context based][context] applications.

[Using][] | [Example][] | [About][]

[Using]: #using "Using"
[Example]: #example "Example"
[About]: #about "About"

[![Unit Tests][ox-ci-status]][ox-ci]
[![Go Reference][goref-ox-status]][goref-ox]
[![Discord Discussion][discord-status]][discord]

[ox-ci]: https://github.com/xo/ox/actions/workflows/test.yml
[ox-ci-status]: https://github.com/xo/ox/actions/workflows/test.yml/badge.svg
[goref-ox]: https://pkg.go.dev/github.com/xo/ox
[goref-ox-status]: https://pkg.go.dev/badge/github.com/xo/ox.svg
[discord]: https://discord.gg/yJKEzc7prt "Discord Discussion"
[discord-status]: https://img.shields.io/discord/829150509658013727.svg?label=Discord&logo=Discord&colorB=7289da&style=flat-square "Discord Discussion"

## Using

Add to a Go project in the usual way:

```sh
$ go get -u github.com/xo/ox@latest
```

## Example

```go
// _examples/simple/main.go
package main

import (
	"context"
	"fmt"
	"net/url"
	"reflect"

	k "github.com/xo/ox"
	_ "github.com/xo/ox/toml"
	_ "github.com/xo/ox/yaml"
)

func main() {
	k.Run(
		context.Background(),
		run,
		k.Usage("simple", "a simple demo of the ox api"),
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
```

### Additional Examples

Please see the [`usql`][usql], [`iv`][iv], [`fv`][fv], and [`xo`][xo]
application's source trees for additional, real-world examples.

[usql]: https://github.com/xo/usql
[iv]: https://github.com/xo/iv
[fv]: https://github.com/xo/fv
[xo]: https://github.com/xo/xo

## About

`ox` was built to address limitations with the popular [`cobra`][cobra]/
[`pflag`][pflag]/[`viper`][viper] package, and similar limitations of the
[`kingpin`][kingpin] and [`urfave/cli`][urfave] packages.

Specific design goals of the `ox` package:

- Context based
- Work with TinyGo out of the box
- No reflection (unless TinyGo supports it)
- No magic, sane defaults, easy overrideable defaults
- Evolving, simple API and maintained
- Constrained "good enough" feature set, no ambition to support every use case/scenario
- Functional option and interface smuggling
- Use generics, iterators and other go1.23+ features where prudent
- Out-of-the-box command completion for bash, fish, zsh, powershell
- **Optional** YAML/TOML support, but easy to enable/disable
- Case sensitive key names from config files
- Allow configuring different lookup keys for different config file types
- Allow registering additional type handlers (marhsalers/unmarshalers) with
  minimal hassle, compatible with standard Go interfaces/types
- Man page generation

[context]: https://pkg.go.dev/context
[cobra]: https://github.com/spf13/cobra
[pflag]: https://github.com/spf13/pflag
[viper]: https://github.com/spf13/viper
[kingpin]: https://github.com/alecthomas/kingpin
[urfave]: https://github.com/urfave/cli
