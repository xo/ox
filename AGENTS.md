# ox

`ox` parses command-line flags and arguments for Go and TinyGo programs. It
builds a tree of commands and sub commands, converts each flag value to a Go
type, and writes the help, the version and the shell completion for the tree.
A flag default can hold variables, such as `$HOME/.${APPNAME|lower}rc`, and
`ox` expands them from the environment and from config loaders. `dbtpl`,
`iv` and `wallgrab` use `ox`, and `usql` moves to it. See D22.

## Standing rules

These hold in every `xo` repository, for every coding agent. dbmeta D110
records them.

1. Stage changes for review. Commit and push only when a maintainer says so.
2. Load the simple-english skill before you write any text that a person
   reads: project documentation, a code comment, an error message or a
   commit message.
3. In a Go project, load the go-pedantry skill before you write or review Go
   code. A rule in this file wins where the two conflict.

`CLAUDE.md` holds one line that imports this file, so that Claude Code and
every other agent read the same rules. Edit this file, not that one.

## Which document to read

| If you are | Read |
| --- | --- |
| asking why something is the way it is | the index at the top of [docs/PLAN.md](docs/PLAN.md) |
| looking for work that is known and not done | [docs/BACKLOG.md](docs/BACKLOG.md), in order of importance |
| designing or building planned work | [docs/DESIGN.md](docs/DESIGN.md), the section that the backlog item names |
| writing a program with ox, or an example of one | [docs/EXAMPLES.md](docs/EXAMPLES.md), by task. Each example there is a test |
| writing a struct tag for `From` | the tag options and escaping in [docs/EXAMPLES.md](docs/EXAMPLES.md) |
| changing config loading or the environment | D34, then "Configuration" in [docs/DESIGN.md](docs/DESIGN.md) and `conf.go` |
| adding a module or a dependency | D33, then "Modules" in [docs/DESIGN.md](docs/DESIGN.md) |
| deciding what to work on next | "Direction" and "Open questions for the maintainers" in [docs/PLAN.md](docs/PLAN.md) |
| adding an import outside the standard library | rule 1 below, and D1 |
| changing code that TinyGo builds | rule 2 below, D2 and D17 |
| adding a default or a string that `ox` writes | rule 3 below, and D8 |
| renaming or removing an exported identifier | rule 4 below, and D18 |
| adding an option | D5, then an option in `opts.go` |
| changing interpolation | D14, D15, D19, and items 31 to 35 of D36 |
| changing help, version or completion | D9 and D10, then `defs.go` and `comp/` |
| importing the completion scripts of cobra again | [docs/COBRA-IMPORT.md](docs/COBRA-IMPORT.md), every step in order |
| adding or changing a test | "Tests" below |
| answering a lint finding | "Linting" below |
| writing a document, a code comment, an error message or a commit message | the `simple-english` skill. Load it first. See "Writing documentation" below |
| writing or reviewing Go code | the `go-pedantry` skill. Load it first. Then "Go conventions" below |
| adding or updating an agent skill | "Agent skills" in [CONTRIBUTING.md](CONTRIBUTING.md) |
| looking for what the library does | [README.md](README.md) |

`CONTRIBUTING.md` holds the same rules for a person, and it is shorter.

A document that is not in that table does not exist. If you cannot find where
something is written down, it is not written down. Ask a maintainer. Do not decide it
yourself, and do not write it as though it were settled.

A number such as D3 or B5 names a decision in `docs/PLAN.md` or an item in
`docs/BACKLOG.md` of this repository. A decision of another repository names
that repository, such as dbmeta D110 or tblfmt D37.

## Rules from the decisions

1. The root module requires nothing, and the root package imports the
   standard library and its own `strcase` and `text` packages, and nothing
   else. A type or a loader that needs another module goes in a sub package
   that is a module of its own, such as `hcl`, `yaml` and `toml`, and
   registers itself from `init`. A loader that needs only the standard
   library, such as `ini`, is a sub package of the root module. A caller
   turns each on with a blank import. `TestTheRootModuleRequiresNothing`
   fails on a `require` in the root `go.mod`. See D1, D33 and D43.
2. `ox` builds with TinyGo, and with the newest Go release and the two
   releases before it. Today these are Go 1.25, 1.26 and 1.27, so each
   `go.mod` declares `go 1.25`, or `go 1.25.0` when a dependency declares
   that, as `hcl/v2` does. When Go 1.28 comes out, each `go` line moves to
   `go 1.26`. CI runs `tinygo test -v ./...` with TinyGo 0.44.0. Code that
   differs between the compilers goes in `go.go`, behind
   `//go:build !tinygo`, and in `tinygo.go`, behind `//go:build tinygo`. Code
   that needs a feature that only a newer Go release has goes in a file with
   a constraint for that release. For example, `type_go127.go` is behind
   `//go:build go1.27`, because the standard `uuid` package is new in Go
   1.27. Such a file registers the feature from `init`. A program that builds
   with an older Go does not get the feature, and everything else works. Use
   no other build constraint in the package. If TinyGo lacks a function, copy
   it from the Go source tree. Say so in its comment, with a `TODO` that says
   when to remove the copy, as `fields` in `cmd.go` does. Do not raise the
   `go` line in `go.mod` beyond what the pinned TinyGo builds. See D2, D17
   and D24.
3. Each default is an exported package level variable with a `Default`
   prefix, in the `var` block at the top of `ox.go`, `interp.go` or
   `conf.go`. Each string that `ox` writes for a person is a variable in
   `text/text.go`. A caller changes the behavior when it assigns one. See D8.
4. `ox` makes no promise of backward compatibility. Rename or remove an
   exported identifier when that makes the code better. `dbtpl` requires
   `ox`, so name each use in `dbtpl` that a change breaks in the commit
   message. See D18.

## Layout

- `ox.go` holds `Run`, `RunContext`, `RunError`, `AppFrom`, `Context`, the
  `Default` variables, the `Error` type and its constants, and `Ldist`.
- `cmd.go` holds `Command`, `FlagSet` and `Flag`, and the reflect style
  (`FlagsFrom`).
- `opts.go` holds every option func and the `option` type.
- `parse.go` holds `Parse` and `Vars`, the parsed values. `conv.go` converts
  values between types, with `As[T]` and its relatives.
- `type.go` and `value.go` hold the flag types and their values. `size.go`
  holds `Size` and `Rate`. `path.go` holds `Realpath`.
- `defs.go` holds help, version and completion, with `CommandHelp`, and
  embeds `comp/`.
- `interp.go` holds `InterpolateVar`, its operators and its filters. `conf.go`
  holds the config loaders, `DefaultKeyLoader`, the config file formats,
  `MapConfig`, and `Context.Fill`, which gives each flag its value from the
  environment, the config files or its default. See D34.
- `app.go` holds `RegisterDoc` and the code of `App` and `AppFrom`, which
  build a program from a struct and its doc comments (D30).
  `cmd/oxdoc` is the tool that writes the calls to `RegisterDoc` from
  `go generate`.
- `misc.go` holds `UserStateDir`. `go.go` and `tinygo.go` hold the code that
  differs between the compilers. `onerr_string.go` is written by `stringer`,
  from the `go:generate` line in `ox.go`.
- `otx/` holds the typed accessors that read a value from a
  `context.Context`, such as `otx.Get[T]`. See D7.
- `strcase/` changes the case of identifiers. `text/` holds the strings that
  `ox` writes.
- `ini/` reads INI config files. It imports only the standard library and
  `ox`, so it is a package of the root module, and a program turns it on
  with a blank import (D43).
- `hcl/`, `yaml/` and `toml/` are modules of their own,
  each with its own `go.mod` and its own tags, such as `yaml/v0.1.0`.
  `hcl`, `yaml` and `toml` register their config file formats. See D1 and
  D33.
- `test/` is a module that nothing imports. It holds the tests that need
  more than one module, such as a config file of each format read end to
  end. Its `go.mod` replaces each module with its folder in the repository.
- No `go.work` is committed. Make one in the root, which `.gitignore`
  ignores: `go work init . ./hcl ./toml ./yaml ./test`. A sub
  module requires a root version from after the split of D33, so it builds
  without a workspace only after a maintainer pushes that commit.
- `comp/` holds the completion scripts, which come from the scripts of
  `cobra`. `LICENSE.completions.txt` holds their license. Do not edit a
  script by hand. Change the importer and import again. See D9, D23 and
  `docs/COBRA-IMPORT.md`.
- `_tools/cobraimport` imports the completion scripts from a cobra release.
  `_tools/oximport` reads the interface of a program and writes an `ox`
  program in each style (D31). `go test ./...` does not build `_tools/`, so
  test it with `go test ./_tools/oximport/...`.
- `_examples/` holds example programs, one directory each.
  `_examples/greet` and `_examples/paint` are the programs in the Using
  section of `README.md`.
  `_examples/usql` is a sample of every option of `usql`, built as `usql`
  builds it (B42).
  `_examples/gen/` holds programs that the removed `_examples/gen.go` wrote
  from the help of real tools, and `TestGenHelpIsTheSame` builds them.
  `go test ./...` does not build `_examples/`, so CI runs `go vet` on each
  folder of it. See D13 and D31.
- `repo_test.go` holds the tests that read the repository and not the
  package. See D21.

## Tests

`go test ./...` in the root does not reach the other modules. Run each
module, with the workspace above:

```sh
for mod in . hcl toml yaml test; do (cd $mod && go test -race -count=1 ./...); done
tinygo test -v ./...
```

To test with an older Go release, set `GOTOOLCHAIN`. The `go` command
downloads that release from the Go module proxy. `go work init` writes the
local Go release in `go.work`, and an older release refuses that file. So
lower the `go` line in `go.work` first. It is `1.25.0` and not `1.25`,
because `hcl` and `test` declare `go 1.25.0`:

```sh
go work edit -go=1.25.0
for mod in . hcl toml yaml test; do (cd $mod && GOTOOLCHAIN=go1.25.0 go test -count=1 ./...); done
```

CI makes the same workspace and runs `go test -v ./...` in each module with
Go 1.25, Go 1.26 and the stable Go release. Then it tests `_tools/` and runs
`go vet` on each folder of `_examples/`, because `./...` skips a folder that
starts with `_`:

```sh
go test ./_tools/cobraimport ./_tools/oximport/...
go vet $(find _examples -name '*.go' -exec dirname {} \; | sort -u | sed 's|^|./|')
```

It also runs
`tinygo test -v ./...` in the root module with TinyGo 0.44.0. Each job runs
on `ubuntu-latest`. TinyGo refuses a Go release that is newer than the
releases that it knows. So if the local Go is 1.27, use TinyGo 0.44.0 or
later.

The tests of the package are in `package ox`, so they reach unexported
names. `example_test.go` is in `package ox_test`, and its examples are the
package overview on pkg.go.dev. `examples_test.go`, also in `package
ox_test`, holds the examples of `docs/EXAMPLES.md`, and
`test/example_test.go` holds the ones for YAML and TOML. A test is a table of
cases that it runs with `t.Run`.

Some tests compare output with a file in `testdata/`. `.gitignore` ignores
`*.txt`, and `!/testdata/**/*.txt` brings those files back:

1. `TestHelpPsql` in `help_test.go` compares a copy of psql's options, with
   `HelpColumn(27)`, with `testdata/psql-help.txt`, the help of psql 18.7.
   A comment names the one line that still differs, psql's `-1 ("one")`.
2. `TestGenHelpIsTheSame` in `gen_test.go` builds each program in
   `_examples/gen/` and compares its help with `testdata/gen/`, so that a
   change to help cannot change the help of a program that does not ask for
   it. It runs the `go` command, and it skips with `-short`.
3. `TestUsqlSample` in `usql_test.go` builds the sample of `usql` in
   `_examples/usql`, and compares its help with `testdata/usql/help.txt`
   and its parse of some command lines with the test. It runs the `go`
   command, and it skips with `-short`.
4. `comp_test.go` holds the cases of completion, and `conf_test.go` holds
   the sources of a flag value. `TestShellCompletion` in `test/comp_test.go`
   runs a completion in each shell that is installed, and skips with
   `-short`. `BenchmarkCommandLine` in `test/bench_test.go` runs the
   same command line with `ox` and with other Go flag packages, such as
   cobra and kong, and `TestCommandLine` holds that each one parses it. Run
   the benchmark with `go test -run '^$' -bench CommandLine -benchmem` in
   `test/`. No CI job runs it.

`repo_test.go` is in `package ox_test`, and a `//go:build !tinygo` line keeps
it out of the TinyGo run. It reads files and not the package:

1. `TestSkillsAreCopies` fails on a symbolic link in a skill, on a missing
   copy, on two copies that differ, and on a skill folder that
   `skills-lock.json` does not name.
2. `TestClaudeImportsAgents` fails unless `CLAUDE.md` is an ordinary file
   that holds only `@AGENTS.md`.
3. `TestTheRootHoldsFourDocuments` fails on any `.md` file in the root other
   than `README.md`, `AGENTS.md`, `CLAUDE.md` and `CONTRIBUTING.md`.
4. `TestTheDecisionIndexIsComplete` fails unless each decision heading in
   `docs/PLAN.md` has a row in the index, with the same title, status and
   link.
5. `TestAnAmendmentPointsBothWays` fails unless a decision that another
   decision amends names it back.
6. `TestTheRootModuleRequiresNothing` fails on a `require` or a `replace` in
   the root `go.mod`, and on a sub module that is missing (D33).
7. `TestEachExampleIsInTheDocument` fails when a code block in
   `docs/EXAMPLES.md` differs from its `Example` func, or when an `Example`
   func is not in the document. Change both together.
8. `TestTheReadmeShowsTheExamples` fails when a program that `README.md`
   shows with `$ cat _examples/...` differs from its file, or when the
   output of a `$ go run ./_examples/...` differs from what the program
   writes. Change both together.

## Go conventions

Match the code around you. These are the patterns that the code follows.

A sentinel error is a constant of `type Error string`, in the `const` block in
`ox.go`, named with an `Err` prefix. Do not use `errors.New` for one. An error
with more to say is a struct type whose `Unwrap` returns the constant, such
as `SuggestionError`. Wrap an error with `%w`, and put the constant first
where it names the kind of fault:

```go
return "", fmt.Errorf("%w %q: missing /", ErrInvalidOp, v)
```

Compare errors with `errors.Is` and `errors.As`. An error message is lower
case.

Every option is a func that returns an `option` value with a `name` and the
funcs for the targets that it changes. Return the alias that names where it
goes, such as `CommandOption` or `FlagOption`. See D5.

A receiver name is a short word for its type, and it is the same on every
method of the type: `cmd` for `Command`, `fs` for `FlagSet`, `g` for `Flag`,
`ctx` for `Context`, `typ` for `Type`, `help` for `CommandHelp`, `val` for
the value types and `opt` for `option`. The `go-pedantry` skill asks for one
or two letters. Keep the names that the code uses.

Use generics where one func serves many types, such as `As[T]`, `otx.Get[T]`
and `Ldist[T]`. Return an `iter.Seq` to walk a tree, such as
`Command.WalkFlags`.

Discard the result of a write to an `io.Writer` with `_, _ =`. `errcheck` is
disabled.

Every exported identifier has a doc comment that starts with its name. Link
another identifier with brackets, such as `[Context.Stdout]`. A comment in a
func body is short and lower case.

## Linting

`.golangci.yml` enables every linter and disables a list, and no CI job runs
it. See D12. Read a finding as a question and not as a task. If a linter
makes idiomatic Go worse, disable it in `.golangci.yml`. Only a real defect
gets a code change.

`golangci-lint run ./...` in the root reports 156 findings today, with its
limit of 50 for each linter, and 270 without the limit. Most of them are
`paralleltest` and `goconst` in the tests. `.golangci.yml` disables both
`exhaustruct` and `exhaustruct_v5`, and both `wsl` and `wsl_v5`, because the
linter has both names. To see only the findings in your change, run:

```sh
golangci-lint run --new-from-rev=HEAD ./...
```

## Before you stage

Standing rule 1 applies. Stage the change with `git add`, and give a
maintainer a proposed commit message. Do not tag. All of these must pass first:

```sh
gofmt -l .
for mod in . hcl toml yaml test; do (cd $mod && go vet ./... && go test -race -count=1 ./...); done
golangci-lint run --new-from-rev=HEAD ./...
```

`gofmt -l .` must print nothing.

## Writing documentation

A new document goes in `docs/`. Only `README.md`, `AGENTS.md`, `CLAUDE.md`
and `CONTRIBUTING.md` belong in the repository root, and
`TestTheRootHoldsFourDocuments` fails on any other. Add a new document to
the table at the top of this file and to the list in `README.md`.

A decision goes in `docs/PLAN.md` and nowhere else. `ox` is a small library,
so its decisions stay in that one file (dbmeta D111). Give the decision the
next number and a heading such as `### D22. Title. Decided.` The status is
`Decided` only when a maintainer chose it. If you do not know whether a
maintainer chose it, write `Proposed` and ask a maintainer. If a decision changes an earlier one, write
`Amends D3` in its heading, and write `Amended by D22` in the heading of D3.
Add a row to the index at the top: the number with a link to the heading,
the title, and the status. `TestTheDecisionIndexIsComplete` prints the row
when it is missing.

Work that is known and not done goes in `docs/BACKLOG.md`, and each item
names where it came from. When an item is done, delete it, and record
anything decided in `docs/PLAN.md`.

Standing rule 2 applies to every such text. In brief: short sentences, the
active voice, no contractions, no semicolons and no em dashes. Use `can`,
`will` and `must`, never `should`, `may` or `might`. Put the condition
before the command.
