# xo/ox

`xo/ox` is a Go and TinyGo package for command-line argument and flag parsing.

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
	"math/big"
	"net/url"
	"reflect"

	"github.com/xo/ox"
	_ "github.com/xo/ox/toml"
	_ "github.com/xo/ox/yaml"
)

func main() {
	ox.Run(
		ox.Exec(run),
		ox.Usage("simple", "a simple demo of the ox api"),
		ox.Version("0.0.0-dev"),
		ox.Help(),
		ox.Comp(),
		ox.UserConfigFile(),
		ox.Flags().
			Var("param-a", "parameter a", ox.Short("a"), ox.Key("", "param-a"), ox.Key("yaml", "my_param_a"), ox.Key("toml", "paramA")).
			Int("param-b", "parameter b", ox.Short("b"), ox.Default(125)).
			Slice("floats", "a slice of float64", ox.Float64T, ox.Short("f")).
			URL("url", "a url", ox.Alias("my-url", "other url flag alias"), ox.Short("u")).
			Count("verbose", "verbose", ox.Short("v")),
		ox.Sub(
			ox.Exec(sub),
			ox.Usage("sub", "a sub command"),
			ox.Alias("subCommand", "a sub command alias"),
			ox.Flags().
				Var("sub", "sub param").
				Slice("strings", "a slice of strings", ox.Short("s")).
				Slice("bigint", "a slice of big ints", ox.BigIntT, ox.Short("t")).
				Map("ints", "a map of integers", ox.IntT, ox.Short("i")),
			ox.ValidArgs(0, 10),
		),
	)
}

func run(ctx context.Context, args []string) error {
	fmt.Println("run args:", args)

	// get param-a
	paramA := ox.Get[string](ctx, "param-a")
	fmt.Println("paramA:", paramA)

	// convert param-b (int) into a string
	paramB := ox.String(ctx, "param-b")
	fmt.Println("paramB:", paramB)

	// a slice
	floats := ox.Slice[float64](ctx, "floats")
	fmt.Println("floats:", floats)

	// convert a slice's values to strings
	floatStrings := ox.Slice[string](ctx, "floats")
	fmt.Println("floatStrings:", floatStrings)

	// sub param is not available in this command, as it was defined on a sub
	// command and not on the root command
	sub := ox.Get[string](ctx, "sub")
	fmt.Println("sub:", sub)

	// a url
	if u := ox.URL(ctx, "url"); u != nil {
		// NOTE: this is wrapped in a if block, because when no flag has been
		// NOTE: passed, tinygo's fmt.Println will panic with a *url.URL(nil),
		// NOTE: however Go's fmt.Println does not
		fmt.Println("url:", u)
		// url alternate
		urlAlt := ox.Get[*url.URL](ctx, "url")
		fmt.Println("urlAlt:", urlAlt)
	}

	// verbose as its own type
	type Verbosity int64
	v := ox.Get[Verbosity](ctx, "verbose")
	fmt.Println("verbosity:", v, reflect.TypeOf(v))

	return nil
}

func sub(ctx context.Context, args []string) error {
	fmt.Println("sub args:", args)

	// get param-a, as any parent's
	paramA := ox.Get[string](ctx, "param-a")
	fmt.Println("paramA:", paramA)

	// convert param-b (int) into a uint32
	paramB := ox.Uint32(ctx, "param-b")
	fmt.Println("paramB:", paramB)

	// the floats param is available, as this is a sub command
	floats := ox.Slice[float64](ctx, "floats")
	fmt.Println("floats:", floats)

	// sub is available here
	sub := ox.Get[string](ctx, "sub")
	fmt.Println("subParam:", sub)

	// get strings
	slice := ox.Slice[string](ctx, "strings")
	fmt.Println("slice:", slice)

	// slice of *big.Int
	bigint := ox.Slice[*big.Int](ctx, "bigint")
	fmt.Println("bigint:", bigint)

	// map of ints, converted to int64
	ints := ox.Map[string, int64](ctx, "ints")
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

`ox` aims to provide a robust and simple command-line package for the most
common command-line use-cases in [Go][golang].

`ox` was built to streamline/simplify complexity found in the powerful (and
popular!) [`cobra`][cobra]/ [`pflag`][pflag]/[`viper`][viper] combo. `ox` is
written in pure Go, with no external package dependencies, and provides a
robust, extensible type system, as well as configuration loaders for
[YAML][yaml], [TOML][toml], [HCL][hcl] that can be optionally enabled/disabled
through imports.

`ox` avoids "magic", and has sane, sensible defaults. No interfaces, type
members or other internal logic is hidden or obscured. When using `ox`, the
user can manually build commands/flags however they see fit.

Wherever an external package is used, the logic is encapsulated via a simple,
anonymous import, thus allowing downstream users of the package to have a clean
dependency heirarchy and to pick and choose the functionality that `ox` has.

`ox` has been designed to use generics, and is built with Go 1.23+ applications
in mind and works with [TinyGo][tinygo].

Specific design considerations of the `ox` package:

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

Other command-line packages:

- [spf13/cobra][cobra] + [spf13/viper][viper] + [spf13/pflag][pflag]
- [urfave/cli][urfave]
- [alecthomas/kong][kong]
- [alecthomas/kingpin][kingpin]

[cobra]: https://github.com/spf13/cobra
[golang]: https://go.dev
[hcl]: https://github.com/hashicorp/hcl
[kingpin]: https://github.com/alecthomas/kingpin
[kong]: https://github.com/alecthomas/kong
[pflag]: https://github.com/spf13/pflag
[tinygo]: https://tinygo.org
[toml]: https://toml.io
[urfave]: https://github.com/urfave/cli
[viper]: https://github.com/spf13/viper
[yaml]: https://yaml.org
