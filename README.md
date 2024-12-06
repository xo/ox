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

- Long (`--arg`) and Short (`-a`) flag parsing
- POSIX-style/compatible flag parsing (`-vvv`, `-mfoo=bar` `-m foo=bar`, `--map=foo=bar`)
- Commands and sub commands heirarchy
- Support for standard library types, including slices and maps
- Register different types
- Testable commands/sub-commands
- Simple/flexible API [see example][Example].
- Generics enabled
- Extremely fast
- TinyGo compatible
- Standard help, version and shell completion
- Environment, YAML, TOML, HCL config loading
- Argument validation (count and values)
- Reflection, Bind, and Context style APIs

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

Wherever an a non-standard package has been used, such as for the [YAML][yaml],
[TOML][toml], or [HCL][hcl] loaders, or for the built-in support for
[colors](color), and [UUIDs](uuid), the external dependencies are fully
optional, requiring a import of a `xo/ox` subpackage, for example:

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
