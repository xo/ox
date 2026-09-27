# ox

`ox` parses command-line flags and arguments for Go and TinyGo programs. It
builds a tree of commands and sub commands, converts each flag value to a Go
type, and writes the help, the version and the shell completion for the tree.
A flag default can hold variables, such as `$HOME/.${APPNAME|lower}rc`, and
`ox` expands them from the environment and from config loaders. `dbtpl` is a
consumer.

## Standing rules

These hold in every `xo` repository, for every coding agent. dbmeta D110
records them.

1. Stage changes for review. Commit and push only when Ken says so.
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
| looking for work that is known and not done | [docs/BACKLOG.md](docs/BACKLOG.md) |
| deciding what to work on next | "Direction" and "Open questions for Ken" in [docs/PLAN.md](docs/PLAN.md) |
| adding an import outside the standard library | rule 1 below, and D1 |
| changing code that TinyGo builds | rule 2 below, D2 and D17 |
| adding a default or a string that `ox` writes | rule 3 below, and D8 |
| renaming or removing an exported identifier | rule 4 below, and D18 |
| adding an option | D5, then an option in `opts.go` |
| changing interpolation | D14, D15 and D19, then B1 and B2 in [docs/BACKLOG.md](docs/BACKLOG.md) |
| changing help, version or completion | D9 and D10, then `defs.go` and `comp/` |
| adding or changing a test | "Tests" below |
| answering a lint finding | "Linting" below |
| writing a document, a code comment, an error message or a commit message | the `simple-english` skill. Load it first. See "Writing documentation" below |
| writing or reviewing Go code | the `go-pedantry` skill. Load it first. Then "Go conventions" below |
| adding or updating an agent skill | "Agent skills" in [CONTRIBUTING.md](CONTRIBUTING.md) |
| looking for what the library does | [README.md](README.md) |

`CONTRIBUTING.md` holds the same rules for a person, and it is shorter.

A document that is not in that table does not exist. If you cannot find where
something is written down, it is not written down. Ask Ken. Do not decide it
yourself, and do not write it as though it were settled.

A number such as D3 or B5 names a decision in `docs/PLAN.md` or an item in
`docs/BACKLOG.md` of this repository. A decision of another repository names
that repository, such as dbmeta D110 or tblfmt D37.

## Rules from the decisions

1. The root package imports the standard library and its own `strcase` and
   `text` packages, and nothing else. A type or a loader that needs another
   module goes in a subpackage that registers itself from `init`, such as
   `color`, `glob` and `yaml`. A caller turns it on with a blank import. See
   D1.
2. `ox` builds with TinyGo. CI runs `tinygo test -v ./...` with TinyGo 0.42.0.
   Code that differs between the compilers goes in `go.go`, behind
   `//go:build !tinygo`, and in `tinygo.go`, behind `//go:build tinygo`. Use
   no other build constraint in the package. If TinyGo lacks a function, copy
   it from the Go source tree. Say so in its comment, with a `TODO` that says
   when to remove the copy, as `fields` in `cmd.go` does. Do not raise the
   `go` line in `go.mod` beyond what the pinned TinyGo builds. See D2 and D17.
3. Each default is an exported package level variable with a `Default`
   prefix, in the `var` block at the top of `ox.go`, `interp.go` or
   `conf.go`. Each string that `ox` writes for a person is a variable in
   `text/text.go`. A caller changes the behavior when it assigns one. See D8.
4. `ox` makes no promise of backward compatibility. Rename or remove an
   exported identifier when that makes the code better. `dbtpl` requires
   `ox`, so name each use in `dbtpl` that a change breaks in the commit
   message. See D18.

## Layout

- `ox.go` holds `Run`, `RunContext`, `Context`, the `Default` variables, the
  `Error` type and its constants, and `Ldist`.
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
  holds the config loaders and `DefaultKeyLoader`.
- `misc.go` holds `UserStateDir`. `go.go` and `tinygo.go` hold the code that
  differs between the compilers. `onerr_string.go` is written by `stringer`,
  from the `go:generate` line in `ox.go`.
- `otx/` holds the typed accessors that read a value from a
  `context.Context`, such as `otx.Get[T]`. See D7.
- `strcase/` changes the case of identifiers. `text/` holds the strings that
  `ox` writes.
- `color/`, `glob/`, `yaml/` and `toml/` are the optional subpackages of D1.
  `yaml` and `toml` are not finished. See B6.
- `comp/` holds the completion scripts, which started as the scripts of
  `cobra`. `LICENSE.completions.txt` holds their license. See D9.
- `_examples/` holds example programs, one directory each.
  `_examples/gen.go` writes the programs in `_examples/gen/` from the help of
  real tools. `go test ./...` does not build `_examples/`. See D13 and B14.
- `repo_test.go` holds the tests that read the repository and not the
  package. See D21.

## Tests

Run these in the repository root:

```sh
go test -race -count=1 ./...
tinygo test -v ./...
```

CI runs `go test -v ./...` with the stable Go release, and
`tinygo test -v ./...` with TinyGo 0.42.0, both on `ubuntu-latest`. A local
TinyGo older than 0.42.0 refuses the module, because `go.mod` declares Go
1.27.

The tests of the package are in `package ox`, so they reach unexported
names. `example_test.go` is in `package ox_test`, and its examples are the
package overview on pkg.go.dev. A test is a table of cases that it runs with
`t.Run`.

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

Main has 263 findings today. `exhaustruct` and `wsl` are renamed
`exhaustruct_v5` and `wsl_v5`, so their disable lines no longer match. See
B16. To see only the findings in your change, run:

```sh
golangci-lint run --new-from-rev=HEAD ./...
```

## Before you stage

Standing rule 1 applies. Stage the change with `git add`, and give Ken a
proposed commit message. Do not tag. All of these must pass first:

```sh
gofmt -l . && go test -race -count=1 ./...
golangci-lint run --new-from-rev=HEAD ./...
```

`gofmt -l .` must print nothing. `go vet ./...` fails on main with two
self-assignments, so it is not in the list until B15 is done.

## Writing documentation

A new document goes in `docs/`. Only `README.md`, `AGENTS.md`, `CLAUDE.md`
and `CONTRIBUTING.md` belong in the repository root, and
`TestTheRootHoldsFourDocuments` fails on any other. Add a new document to
the table at the top of this file and to the list in `README.md`.

A decision goes in `docs/PLAN.md` and nowhere else. `ox` is a small library,
so its decisions stay in that one file (dbmeta D111). Give the decision the
next number and a heading such as `### D22. Title. Decided.` The status is
`Decided` only when Ken chose it. If you do not know whether Ken chose it,
write `Proposed` and ask him. If a decision changes an earlier one, write
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
