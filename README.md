# xo/ox

`xo/ox` is a Go and TinyGo package for command-line argument and flag parsing.

[Using][] | [Example][] | [Applications][] | [About][] | [License][]

[Using]: #using "Using"
[Example]: #example "Example"
[Applications]: #applications "Applications"
[About]: #about "About"
[License]: #license "License"

[![Unit Tests][ox-ci-status]][ox-ci]
[![Go Reference][goref-ox-status]][goref-ox]
[![Discord Discussion][discord-status]][discord]

[ox-ci]: https://github.com/xo/ox/actions/workflows/test.yml
[ox-ci-status]: https://github.com/xo/ox/actions/workflows/test.yml/badge.svg
[goref-ox]: https://pkg.go.dev/github.com/xo/ox
[goref-ox-status]: https://pkg.go.dev/badge/github.com/xo/ox.svg
[discord]: https://discord.gg/yJKEzc7prt "Discord Discussion"
[discord-status]: https://img.shields.io/discord/829150509658013727.svg?label=Discord&logo=Discord&colorB=7289da&style=flat-square "Discord Discussion"

## Features

- Long (`--arg`, `--arg val`) and Short (`-a`, `-a val`) flag parsing
- POSIX-style/compatible flag parsing (`-vvv`, `-mfoo=bar` `-m foo=bar`, `--map=foo=bar`)
- Flags can have optional arguments (via `NoArg`) and type specific defaults
- Full command tree and sub command heirarchy
- Support for all scalars:
  - `[]byte`, `string`, `[]rune`, `byte`, `rune`
  - `int64`, `int32`, `int16`, `int8`, `int`
  - `uint64`, `uint32`, `uint16`, `uint8`, `uint`
  - `float64`, `float32`
  - `complex128`, `complex64`
- Support for standard library types:
  - `time.Time`, `time.Duration`
  - `*big.Int`, `*big.Float`, `*big.Rat`
  - `*url.URL`
  - `*netip.Addr`, `*netip.AddrPort`, `*netip.Prefix`
- Support for compound types of all above (slices/maps):
  - `[]int`, `[][]byte`, `[]string`, `[]float64`, `[]*big.Int`, etc.
  - `map[string]string`, `map[int]string`, `map[float64]*url.URL`, etc...
- Additional types:
  - `ox.FormattedTime` - formatted time values, set to a specific `time.Layout`
  - `ox.CountT` / `type:count` - incrementing counter, such as for verbosity `-vvvv`
  - `ox.PathT` / `type:path` - a file system path
  - `ox.HookT` - argument `func` hook
- Registerable user defined types, which work with all API styles
- Testable commands/sub-commands
- Simple/flexible APIs for Reflection, Bind, and Context style use cases
- Generics used where it makes sense
- Fast
- Environment, YAML, TOML, HCL config loading
- Deferred default value expansion
- Standard help, version and shell completion
- Suggestions for command names, aliases, and suggested names
- Argument validation and advanced shell completion support
- TinyGo compatible

## Using

Add to a Go project in the usual way:

```sh
$ go get -u github.com/xo/ox@latest
```

## Example

Examples are available in [the package overview examples][pkg-overview], as
well as [in the `_examples` directory](_examples).

[pkg-overview]: https://pkg.go.dev/github.com/xo/ox#pkg-overview

## Applications

The following applications make use of the `xo/ox` package for command-line
parsing:

- [`usql`][usql] - a universal command-line interface for SQL databases
- [`xo`][xo] - a templated code generator for databases
- [`iv`][iv] - a command-line terminal graphics image viewer
- [`fv`][fv] - a command-line terminal graphics font viewer
- [`wallgrab`][wallgrab] - a Apple Aerial wallpaper downloader

[usql]: https://github.com/xo/usql
[xo]: https://github.com/xo/xo
[iv]: https://github.com/xo/iv
[fv]: https://github.com/xo/fv
[wallgrab]: https://github.com/kenshaw/wallgrab

## About

`ox` aims to provide a robust and simple command-line package for the most
common command-line use-cases in [Go][golang].

`ox` was built to streamline/simplify complexity found in the powerful (and
popular!) [`cobra`][cobra]/ [`pflag`][pflag]/[`viper`][viper] combo. `ox` is
written in pure Go, with no non-standard package dependencies, and provides a
robust, extensible type system, as well as configuration loaders for
[YAML][yaml], [TOML][toml], [HCL][hcl] that can be optionally enabled/disabled
through imports.

`ox` avoids "magic", and has sane, sensible defaults. No interfaces, type
members or other internal logic is hidden or obscured. When using `ox`, the
user can manually build commands/flags however they see fit.

Wherever a non-standard package has been used, such as for the [YAML][yaml],
[TOML][toml], or [HCL][hcl] loaders, or for the built-in support for
[colors](color), and [UUIDs](uuid), the external dependencies are optional,
requiring a import of a `xo/ox` subpackage, for example:

```go
import (
  // base package
  "github.com/xo/ox"

  // import config loaders
  _ "github.com/xo/ox/hcl"
  _ "github.com/xo/ox/toml"
  _ "github.com/xo/ox/yaml"

  // well-known types
  _ "github.com/xo/ox/color"
  _ "github.com/xo/ox/uuid"
)
```

`ox` has been designed to use generics, and is built with Go 1.23+ applications
in mind and works with [TinyGo][tinygo].

Specific design considerations of the `ox` package:

- Constrained "good enough" feature set, no ambition to support every use
  case/scenario
- No magic, sane defaults, overrideable defaults
- Functional option and interface smuggling
- Use generics, iterators and other go1.23+ features where prudent
- Work with TinyGo out of the box
- Minimal use of reflection (unless TinyGo supports it)
- Case sensitive
- Enable registration for config file loaders, types with minimal hassle
- Man page generation
- **Optional** support for common use-cases, via package dependencies

Other command-line packages:

- [spf13/cobra][cobra] + [spf13/viper][viper] + [spf13/pflag][pflag]
- [urfave/cli][urfave]
- [alecthomas/kong][kong]
- [alecthomas/kingpin][kingpin]
- [jessevdk/go-flags][go-flags]
- [mow.cli][mowcli]
- [peterbourgon/ff][pbff]

[cobra]: https://github.com/spf13/cobra
[go-flags]: https://github.com/jessevdk/go-flags
[golang]: https://go.dev
[kingpin]: https://github.com/alecthomas/kingpin
[kong]: https://github.com/alecthomas/kong
[mowcli]: https://github.com/jawher/mow.cli
[pbff]: https://github.com/peterbourgon/ff
[pflag]: https://github.com/spf13/pflag
[tinygo]: https://tinygo.org
[urfave]: https://github.com/urfave/cli
[viper]: https://github.com/spf13/viper

Articles:

- [Matt Turner, Choosing a Go CLI Library][mtgocli]

[mtgocli]: https://mt165.co.uk/blog/golang-cli-library/

## License

`xo/ox` is licensed under the [MIT License](LICENSE). `ox`'s completion scripts
are originally cribbed from the [cobra][cobra] project, and are made available
under the [Apache License](LICENSE.completions.txt).

[hcl]: https://github.com/hashicorp/hcl
[toml]: https://toml.io
[yaml]: https://yaml.org
