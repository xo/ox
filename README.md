<div align="center">
  <img src="docs/ox.svg" alt="ox" width="200"/>
  <br/>
  <a href="#using" title="Using">Using</a> |
  <a href="#about" title="About">About</a> |
  <a href="#features" title="Features">Features</a> |
  <a href="#comparison" title="Comparison">Comparison</a> |
  <a href="#applications" title="Applications">Applications</a> |
  <a href="#documents" title="Documents">Documents</a> |
  <a href="https://pkg.go.dev/github.com/xo/ox" title="Go Reference">Reference</a> |
  <a href="#license" title="License">License</a>
</div>

<br/>

[![Unit Tests][ox-ci-status]][ox-ci]
[![Go Reference][goref-ox-status]][goref-ox]
[![Discord Discussion][discord-status]][discord]

[ox-ci]: https://github.com/xo/ox/actions/workflows/test.yml "Test CI"
[ox-ci-status]: https://github.com/xo/ox/actions/workflows/test.yml/badge.svg "Test CI"
[goref-ox]: https://pkg.go.dev/github.com/xo/ox "Go Reference"
[goref-ox-status]: https://pkg.go.dev/badge/github.com/xo/ox.svg "Go Reference"
[discord]: https://discord.gg/yJKEzc7prt "Discord Discussion"
[discord-status]: https://img.shields.io/discord/829150509658013727.svg?label=Discord&logo=Discord&colorB=7289da&style=flat-square "Discord Discussion"

# ox

`ox` is a Go and TinyGo package for command-line argument and flag parsing.

## Using

Add `ox` to a Go project in the usual way:

```sh
$ go get -u github.com/xo/ox@latest
```

A simple program, from [`_examples/greet`](_examples/greet):

```sh
$ cat _examples/greet/main.go
// Command greet says hello. It is the example in the README of ox.
package main

import (
	"context"
	"fmt"

	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

func main() {
	ox.Run(
		ox.Usage("greet", "say hello"),
		ox.Defaults(),
		ox.Flags().
			String("name", "the name to greet", ox.Short("n"), ox.Default("world")).
			Int("times", "the number of greetings", ox.Short("t"), ox.Default(1)),
		ox.ValidArgs(0, 0),
		ox.Exec(run),
	)
}

func run(ctx context.Context) error {
	for range otx.Int(ctx, "times") {
		fmt.Println("hello", otx.String(ctx, "name"))
	}
	return nil
}
$ go run ./_examples/greet --help
greet say hello

Usage:
  greet [flags]

Flags:
  -n, --name string  the name to greet (default: world)
  -t, --times int    the number of greetings (default: 1)
  -v, --version      show version, then exit
  -h, --help         show help, then exit
$ go run ./_examples/greet -n ox -t 2
hello ox
hello ox
```

A program built from a struct, with `ox.AppFrom`, from
[`_examples/paint`](_examples/paint). Each field is a flag, and the `Exec`
method runs the program. A field of a type that reads its own text, such as
`slog.Level`, needs no registration, and help writes its type in lower case:

```sh
$ cat _examples/paint/main.go
// Command paint is the struct example in the README of ox. ox.AppFrom builds
// the program from the struct: each field is a flag, and the Exec method
// runs the program.
package main

import (
	"fmt"
	"log/slog"

	"github.com/xo/ox"
)

// Paint paints a thing.
type Paint struct {
	Thing string     `ox:"the thing to paint,short:t,default:fence"`
	Coats int        `ox:"the number of coats,short:c,default:2"`
	Log   slog.Level `ox:"the log level,short:l,default:info"`
}

// Exec runs the program.
func (p *Paint) Exec() {
	fmt.Printf("painting the %s with %d coats (log %v)\n", p.Thing, p.Coats, p.Log)
}

func main() {
	ox.AppFrom(
		new(Paint),
		ox.Usage("paint", "paint a thing"),
		ox.ValidArgs(0, 0),
	)
}
$ go run ./_examples/paint --help
paint paint a thing

Usage:
  paint [flags]

Flags:
  -t, --thing string  the thing to paint (default: fence)
  -c, --coats int     the number of coats (default: 2)
  -l, --log level     the log level (default: info)
  -v, --version       show version, then exit
  -h, --help          show help, then exit
$ go run ./_examples/paint -t door -c 3 --log debug
painting the door with 3 coats (log DEBUG)
```

### Example

More examples are in [the package overview examples][pkg-overview],
[`docs/EXAMPLES.md`](docs/EXAMPLES.md), and [the `_examples`
directory](_examples).

[pkg-overview]: https://pkg.go.dev/github.com/xo/ox#pkg-overview

## About

`ox` aims to provide a robust and simple command-line package for the most
common command-line use-cases in [Go][golang].

`ox` was built to streamline/simplify complexity found in the powerful (and
popular!) [`cobra`][cobra]/ [`pflag`][pflag]/[`viper`][viper] combo. `ox` is
written in pure Go, with no non-standard package dependencies, and provides a
robust, extensible type system, as well as configuration loaders for
[YAML][yaml], [TOML][toml], [HCL][hcl] and INI that can be optionally
enabled/disabled through imports.

`ox` avoids "magic", and has sane, sensible defaults. No interfaces, type
members or other internal logic is hidden or obscured. When using `ox`, the
user can manually build commands and flags however they see fit.

Wherever a non-standard package has been used, such as for the [YAML][yaml],
[TOML][toml], or [HCL][hcl] loaders, the external dependencies are optional,
requiring an import of an `ox` subpackage, for example:

```go
import (
  // base package
  "github.com/xo/ox"

  // import config loaders
  _ "github.com/xo/ox/hcl"
  _ "github.com/xo/ox/ini"
  _ "github.com/xo/ox/toml"
  _ "github.com/xo/ox/yaml"
)
```

Each of these subpackages, other than `ini`, is a Go module of its own, with
its own tags, such as `yaml/v0.1.0`, so that a program that imports only
`github.com/xo/ox` gets no other dependency. `ini` uses only the standard
library, so it is in the module of `ox`, and a program reads INI files only
when it imports it. `github.com/xo/ox/metadoc` also uses only the standard
library. A program imports it and adds the option
`metadoc.Command(metadoc.ManPage(), metadoc.Markdown())`, and then
`app __docs man` writes a man page. See
[`metadoc/example_test.go`](metadoc/example_test.go). Upgrade `ox` and its
modules together:

```sh
$ go get github.com/xo/ox@latest github.com/xo/ox/yaml@latest
```

A program reads its config files, and the environment, with options on the
root command. Each flag takes its value from the command line, then the
environment, then the config files, then its default:

```go
ox.Run(
  ox.Usage("app", "an app"),
  ox.EnvPrefix("APP"),                                // APP_LOG_LEVEL fills --log-level
  ox.ConfigFiles("/etc/app.yaml", "$APPCONFIG/config.yaml"), // a later file wins
  ox.Flags().
    String("config", "config file", ox.ConfigFlag()). // --config replaces the files
    String("log-level", "log level", ox.Default("info")).
    Int("port", "port", ox.Key("", "server.port")),   // reads server.port
  ox.Exec(run),
)
```

A flag of a sub command reads its key under the path of the command, so
`--steps` of `app migrate` reads `migrate.steps`. A field of a nested struct
reads the path of the field, so the field `Source.Dir` of `app migrate`
reads `migrate.source.dir`. Each config loader has two examples, which read
the same settings: one with flags built by hand, and one with a struct:

| Format | Import | Examples |
| --- | --- | --- |
| JSON | `github.com/xo/ox` | [`examples_test.go`](examples_test.go), `ExampleConfigFiles_subCommand` |
| YAML | `github.com/xo/ox/yaml` | [`yaml/example_test.go`](yaml/example_test.go) |
| TOML | `github.com/xo/ox/toml` | [`toml/example_test.go`](toml/example_test.go) |
| HCL | `github.com/xo/ox/hcl` | [`hcl/example_test.go`](hcl/example_test.go) |
| INI | `github.com/xo/ox/ini` | [`ini/example_test.go`](ini/example_test.go) |

`ox` has been designed to use generics, and works with [TinyGo][tinygo]. It
supports the newest Go release and the two releases before it, so it builds
with Go 1.25 and later. A feature that needs a newer release, such as the
`uuid` type of Go 1.27, is there only when a program is built with that
release.

Specific design considerations of the `ox` package:

- Constrained "good enough" feature set, no ambition to support every use
  case/scenario
- No magic, sane defaults, overrideable defaults
- Functional option and interface smuggling
- Use generics, iterators and other features of the oldest supported Go
  release where prudent
- Work with TinyGo out of the box
- Minimal use of reflection (unless TinyGo supports it)
- Case sensitive
- Enable registration for config file loaders, types with minimal hassle
- Man page generation, which `github.com/xo/ox/metadoc` does
- **Optional** support for common use-cases, via package dependencies

[cobra]: https://github.com/spf13/cobra
[golang]: https://go.dev
[pflag]: https://github.com/spf13/pflag
[tinygo]: https://tinygo.org
[viper]: https://github.com/spf13/viper

## Features

- Long (`--arg`, `--arg val`) and Short (`-a`, `-a val`) flag parsing
- POSIX-style/compatible flag parsing (`-vvv`, `-mfoo=bar` `-m foo=bar`, `--map=foo=bar`)
- Flags can have optional arguments (via `NoArg`) and type specific defaults
- Full command tree and sub command heirarchy
- Sub commands inherit the flags of their parents, and `ox.Local` keeps a
  flag to its command. `ox.InheritedFlags` shows the inherited flags in a
  `Global Flags:` section of the help
- `ox.StopAtArg` ends the flags at the first positional argument, for a
  wrapper command such as `app exec cmd --its-own-flag`
- `ox.Abbrev` takes a unique prefix of a long flag, such as `--verb` for
  `--verbose`, as GNU `getopt_long` does
- `ox.NegateBools` gives each bool flag a negated form, such as `--no-debug`,
  and help writes the flag as `--[no-]debug`
- Support for `builtin` types:
  - `[]byte`, `string`, `[]rune`, `byte`, `rune`
  - `int64`, `int32`, `int16`, `int8`, `int`
  - `uint64`, `uint32`, `uint16`, `uint8`, `uint`
  - `float64`, `float32`
  - `complex128`, `complex64`
- Support for standard library types:
  - `time.Time`, `time.Duration`
  - `*big.Int`, `*big.Float`, `*big.Rat`
  - `*url.URL`, `*regexp.Regexp`
  - `*netip.Addr`, `*netip.AddrPort`, `*netip.Prefix`
  - `*uuid.UUID`, when built with Go 1.27 or later
- Non-standard types:
  - `ox.Size` - a byte size (`15 MiB`, `1 GB`, `1536Mi`, `6e9`, ...). The
    verbs `%q` and `%Q` write the exact form of a Kubernetes quantity, such
    as `1536Mi` or `6G`, and `ox.ParseSize` reads it back
  - `ox.Rate` - a byte rate (`15 MiB/s`, `1 GB/h`, `1536Mi/s`, ...), with
    the same verbs
  - `ox.PortRange` / `type:portrange` - a port or a range of ports (`80`,
    `100-200`, `100:200`), with `Contains` and `Ports`
- Support for aggregate types of above (slices/arrays/maps):
  - `[]int`, `[][]byte`, `[]string`, `[]float64`, `[]*big.Int`, etc.
  - `map[string]string`, `map[int]string`, `map[float64]*url.URL`, etc.
- Additional type support:
  - `ox.DateTimeT`, `ox.DateT`, `ox.TimeT` / `type:datetime`, `type:date`, `type:time` - standard dates and times
  - `ox.FormattedTime` - any `time.Time` value using any `time.Layout` format
  - `ox.Strftime` - a time in a strftime layout, such as `%d/%m/%Y`
  - `ox.CountT` / `type:count` - incrementing counter, such as for verbosity `-vvvv`
  - `ox.Base64T` - a base64 encoded string
  - `ox.HexT` - a hex encoded string
  - `ox.PathT` / `type:path` - a file system path
  - `ox.FileT` / `type:file` - a file that exists
  - `ox.DirT` / `type:dir` - a directory that exists
  - `ox.HookT` - argument `func` hook, for hooking flags
- Any type that reads its own text, with `UnmarshalText`, `UnmarshalBinary`
  or `Set`, with no registration, such as `slog.Level`,
  [`colors.Color`][colors] or [`glob.Glob`][glob]. `ox.TypeFor[T]()` gives a
  flag the type, and help writes its name in lower case, such as
  `--fg color` for `colors.Color`
- Registerable user defined types, which work with all API styles
- Testable commands/sub-commands
- Simple/flexible APIs for Reflection, Bind, and Context style use cases
- A program built from a struct and its doc comments, with `ox.AppFrom` and
  `go generate`
- Generics used where it makes sense
- Fast: about as fast as `cobra` when a program starts, with less memory, as
  [`BenchmarkCommandLine`](test/bench_test.go) measures
- Environment, JSON, YAML, TOML, HCL and INI config loading, with the keys
  of a sub command under the path of the command
- Deferred default value expansion
- Standard help, version and shell completion
- Command, argument, and flag completion
- Suggestions for command names, aliases, and suggested names
- Argument validation and advanced shell completion support
- Plus forms of short flags, as bash takes `+o name` and `+ex`, with
  `ox.AllowPlus()`, and a hook that reads the sign with `ctx.IsPlus()`.
  `ox.PlusWords()`, `ox.LetterClusters()` and `ox.OptionalAtEnd()` read the
  rest of bash's command line
- Required flags, and rules between flags: only one of, all or none of, and
  at least one of a group
- Positional arguments with a name and a type, with `ox.Arg`, which are
  converted, checked, completed and written in help as flags are
- Funcs that run before and after a command and each of its sub commands,
  with `ox.Before` and `ox.After`
- A man page and Markdown docs of the command tree, with
  `github.com/xo/ox/metadoc`, as one page or as a file for each command
- TinyGo compatible

## Comparison

Other command-line packages:

- [spf13/cobra][cobra] + [spf13/viper][viper] + [spf13/pflag][pflag]
- [urfave/cli][urfave]
- [alecthomas/kong][kong]
- [alecthomas/kingpin][kingpin]
- [jessevdk/go-flags][go-flags]
- [mow.cli][mowcli]
- [peterbourgon/ff][pbff]

[go-flags]: https://github.com/jessevdk/go-flags
[kingpin]: https://github.com/alecthomas/kingpin
[kong]: https://github.com/alecthomas/kong
[mowcli]: https://github.com/jawher/mow.cli
[pbff]: https://github.com/peterbourgon/ff
[urfave]: https://github.com/urfave/cli

[`test/bench_test.go`](test/bench_test.go) runs the same command line with
`ox` and with each of these packages, and checks that each one parses it.
Run it in the `test` folder:

```sh
$ go test -run '^$' -bench CommandLine -benchmem
```

Articles:

- [Matt Turner, Choosing a Go CLI Library][mtgocli]

[mtgocli]: https://mt165.co.uk/blog/golang-cli-library/

## Applications

The following applications make use of the `ox` package for command-line
parsing:

- [`usql`][usql] - a universal command-line interface for SQL databases
- [`dbtpl`][dbtpl] - a templated code generator for databases
- [`iv`][iv] - a command-line terminal graphics image viewer
- [`wallgrab`][wallgrab] - a Apple Aerial wallpaper downloader

`usql` and `iv` show two different ways to use `ox`:

- [`usql`][usql-cmdline] - builds each flag by hand with `ox.Flags()`, with
  hooks, sections and help values, so that its help matches the help of
  `psql`
- [`iv`][iv-ivcmd] - reads its flags from the `ox` tags of a struct with
  `ox.From`, and [registers a type of its own][iv-ivctx] for a flag value

[usql]: https://github.com/xo/usql
[dbtpl]: https://github.com/xo/dbtpl
[iv]: https://github.com/kenshaw/iv
[wallgrab]: https://github.com/kenshaw/wallgrab
[usql-cmdline]: https://github.com/xo/usql/blob/main/cmdline.go
[iv-ivcmd]: https://github.com/kenshaw/iv/blob/main/ivcmd/ivcmd.go
[iv-ivctx]: https://github.com/kenshaw/iv/blob/main/ivctx/ivctx.go

## Documents

| Document | Holds |
| --- | --- |
| [CONTRIBUTING.md](CONTRIBUTING.md) | How to ask for a change, what a change must do, and the commands to run |
| [docs/PLAN.md](docs/PLAN.md) | The plan, each decision that shapes `ox` and the reason for it, and the open questions |
| [docs/BACKLOG.md](docs/BACKLOG.md) | Known faults and work that is not done, in order of importance |
| [docs/EXAMPLES.md](docs/EXAMPLES.md) | Short examples by task, each of which is a test |
| [docs/DESIGN.md](docs/DESIGN.md) | The design of the planned work |
| [docs/COBRA-IMPORT.md](docs/COBRA-IMPORT.md) | How the completion scripts are imported from cobra |
| [AGENTS.md](AGENTS.md) | The rules for a coding agent. `CLAUDE.md` imports it |

## License

`ox` is licensed under the [MIT License](LICENSE). `ox`'s completion scripts
are originally cribbed from the [cobra][cobra] project, and are made available
under the [Apache License](LICENSE.completions.txt).

[colors]: https://github.com/kenshaw/colors
[glob]: https://github.com/kenshaw/glob
[hcl]: https://github.com/hashicorp/hcl
[toml]: https://toml.io
[yaml]: https://yaml.org
