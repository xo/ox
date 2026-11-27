# Plan

This file holds the plan for `ox`, then every decision that shapes it, then
the open questions for the maintainers. The index of the decisions comes first, so that a
reader can find a decision by its number.

| Decision | Title | Status |
| --- | --- | --- |
| [D1](#d1-a-dependency-outside-the-standard-library-lives-in-a-subpackage-decided) | A dependency outside the standard library lives in a subpackage | Decided |
| [D2](#d2-ox-builds-with-tinygo-and-ci-tests-both-compilers-decided) | ox builds with TinyGo, and CI tests both compilers | Decided |
| [D3](#d3-the-package-is-named-ox-decided) | The package is named ox | Decided |
| [D4](#d4-a-sentinel-error-is-a-constant-of-type-error-decided) | A sentinel error is a constant of type Error | Decided |
| [D5](#d5-every-option-is-an-option-value-and-one-option-interface-carries-it-decided) | Every option is an option value, and one Option interface carries it | Decided |
| [D6](#d6-the-context-bind-and-reflect-styles-build-one-command-tree-decided) | The context, bind and reflect styles build one command tree | Decided |
| [D7](#d7-a-command-gets-the-run-context-from-its-contextcontext-and-otx-reads-it-decided) | A command gets the run context from its context.Context, and otx reads it | Decided |
| [D8](#d8-each-default-is-a-package-level-variable-that-a-caller-can-replace-decided) | Each default is a package level variable that a caller can replace | Decided |
| [D9](#d9-help-version-and-completion-are-built-in-and-the-completion-scripts-come-from-cobra-decided) | Help, version and completion are built in, and the completion scripts come from cobra | Decided |
| [D10](#d10-only-a-root-that-has-sub-commands-gets-a-help-command-and-a-version-command-decided) | Only a root that has sub commands gets a help command and a version command | Decided |
| [D11](#d11-an-unknown-command-gets-a-suggestion-within-a-levenshtein-distance-of-2-decided) | An unknown command gets a suggestion within a Levenshtein distance of 2 | Decided |
| [D12](#d12-golangci-lint-enables-every-linter-and-disables-a-list-decided) | golangci-lint enables every linter and disables a list | Decided |
| [D13](#d13-the-examples-include-command-trees-generated-from-the-help-of-real-tools-decided-amended-by-d31) | The examples include command trees generated from the help of real tools | Decided. Amended by D31 |
| [D14](#d14-a-default-value-is-expanded-with-bash-like-interpolation-decided) | A default value is expanded with bash-like interpolation | Decided |
| [D15](#d15-an-unknown-key-expands-to-nothing-and-a-failed-lookup-writes-its-error-into-the-value-decided) | An unknown key expands to nothing, and a failed lookup writes its error into the value | Decided |
| [D16](#d16-a-config-loader-is-registered-for-a-configtype-decided) | A config loader is registered for a ConfigType | Decided |
| [D17](#d17-the-go-directive-stays-at-a-release-that-the-pinned-tinygo-can-build-decided-amended-by-d24) | The go directive stays at a release that the pinned TinyGo can build | Decided. Amended by D24 |
| [D18](#d18-ox-makes-no-promise-of-backward-compatibility-decided) | ox makes no promise of backward compatibility | Decided |
| [D19](#d19-the-substring-operator-takes-an-end-index-and-not-a-length-proposed) | The substring operator takes an end index and not a length | Proposed |
| [D20](#d20-an-app-builds-its-command-tree-from-the-doc-comment-of-a-struct-proposed-amended-by-d30) | An app builds its command tree from the doc comment of a struct | Proposed. Amended by D30 |
| [D21](#d21-ox-is-set-up-for-coding-agents-as-every-xo-repository-is-decided) | ox is set up for coding agents as every xo repository is | Decided |
| [D22](#d22-ox-is-the-command-line-package-for-the-maintainers-programs-and-usql-moves-to-it-decided) | ox is the command-line package for the maintainers' programs, and usql moves to it | Decided |
| [D23](#d23-a-tool-imports-the-completion-scripts-from-each-cobra-release-decided) | A tool imports the completion scripts from each cobra release | Decided |
| [D24](#d24-ox-supports-the-newest-go-release-and-the-two-before-it-decided-amends-d17) | ox supports the newest Go release and the two before it | Decided. Amends D17 |
| [D25](#d25-the-github-project-takes-discussions-and-not-issues-or-pull-requests-decided) | The GitHub project takes discussions, and not issues or pull requests | Decided |
| [D26](#d26-a-path-flag-is-a-patht-a-filet-or-a-dirt-decided) | A path flag is a PathT, a FileT or a DirT | Decided |
| [D27](#d27-ox-completes-flag-values-and-positional-arguments-as-cobra-does-decided) | ox completes flag values and positional arguments as cobra does | Decided |
| [D28](#d28-one-compfunc-completes-a-flag-or-an-argument-proposed) | One CompFunc completes a flag or an argument | Proposed |
| [D29](#d29-a-program-can-make-its-help-look-like-psqls-help-decided) | A program can make its help look like psql's help | Decided |
| [D30](#d30-appfrom-builds-a-program-from-a-struct-and-its-doc-comments-decided-amends-d20) | AppFrom builds a program from a struct and its doc comments | Decided. Amends D20 |
| [D31](#d31-an-importer-writes-ox-programs-from-the-interface-of-well-known-programs-decided-amends-d13) | An importer writes ox programs from the interface of well-known programs | Decided. Amends D13 |
| [D32](#d32-ox-has-many-examples-for-coding-agents-decided) | ox has many examples for coding agents | Decided |
| [D33](#d33-each-sub-package-with-a-dependency-is-a-module-of-its-own-and-test-is-a-module-decided) | Each sub package with a dependency is a module of its own, and test is a module | Decided |
| [D34](#d34-a-flag-takes-its-value-from-the-command-line-the-environment-the-config-files-or-its-default-proposed) | A flag takes its value from the command line, the environment, the config files or its default | Proposed |
| [D35](#d35-the-choices-that-the-agents-made-on-2026-11-26-proposed) | The choices that the agents made on 2026-11-26 | Proposed |
| [D36](#d36-choices-that-the-backlog-agent-made-proposed) | Choices that the backlog agent made | Proposed |

## Purpose

`ox` parses command-line flags and arguments for Go and TinyGo programs. It
builds a tree of commands and sub commands, and it writes the help, the
version and the shell completion for that tree. It reads a flag default from
the environment and from other sources through interpolation, such as
`$HOME/.${APPNAME|lower}rc`.

The README says why it exists. `ox` makes the `cobra`, `pflag` and `viper`
combination simpler, with a limited set of features, no hidden behavior, and
defaults that a caller can change. It is not meant to support every use case.

## Current state

This section is at `bb2641a` (2026-09-27), with the changes staged on
2026-11-26, which split the modules and add config loading.

1. The root package imports the standard library and two of its own
   subpackages, `strcase` and `text`. See D1.
2. Flag parsing, the type system, help, version, the completion of command
   and flag names, and suggestions work, and the tests cover them.
   `go test ./...` and `tinygo test ./...` pass in CI.
3. Completion completes command and flag names, flag values, positional
   values, hidden aliases and the values of a `CompFunc`, and help can look
   like psql's help with `HelpColumn` (D35). zsh shows no corrections (B47).
4. Config files of JSON, YAML, TOML and HCL, and the environment, fill the
   flags (D34). The tests in `conf_test.go` and in the `test` module cover
   each format and the order of the sources.
5. `AppFrom` builds a program from a struct and its doc comments (D30), and
   `cmd/oxdoc` writes the doc comments from `go generate`. It replaces the
   sketch of D20.
6. The repository is seven modules (D33): the root, which requires nothing,
   `color`, `glob`, `hcl`, `toml`, `yaml` and `test`. Each module has the tag `v0.1.0`, such as `yaml/v0.1.0`, since 2026-11-27. Each `go.mod`
   declares `go 1.25`, and the `uuid` type needs Go 1.27 (D24). Every module
   passes `go vet` and its tests with Go 1.25, 1.26 and 1.27.
7. `golangci-lint` reports 156 findings in the root, and no CI job runs it.
   See D12, and D36 for the count.
8. `dbtpl`, `iv` and `wallgrab` use `ox`. `usql` has its move to `ox` staged,
   and it can commit the move against `v0.1.0`. See D22.

## Direction

A maintainer set this direction on 2026-11-26. [BACKLOG.md](BACKLOG.md) holds
each item in order of importance, and [DESIGN.md](DESIGN.md) holds the design
of each part.

1. Move `usql` to `ox`. That is priority 1 of the backlog: a release that
   builds with Go 1.25 to 1.27 (D24), value completion at parity with cobra
   (D27), help that looks like psql's help (D29), environment variables,
   hidden aliases, and a run that returns its error.
2. Show the structure of `usql` on `ox` in a sample, `_examples/usql`, which
   the `usql` session reviews before `usql` moves. See B42.
3. Add the path types of D26 and the completion API of D28.
4. Add `AppFrom`, which builds a program from a struct and its doc comments
   (D30).
5. Replace `_examples/gen.go` with an importer that reads well-known
   programs and writes each style of `ox` program (D31).
6. Write many more examples for coding agents (D32).
7. Then finish config loading, and the bash-like operators of D14.

## Decisions

Each decision that shapes `ox`, in the order that it was made. The entries are
append only. A decision is never edited to change its conclusion. When a later
decision changes one, the older entry keeps its text, and its heading names
the decision that changed it.

Each heading ends with a status. `Decided` means that the history shows that
a maintainer chose it. `Proposed` means that the code does it, and no record shows that
a maintainer chose it. A proposed decision is also an open question at the end of this
file. `Amends D3` and `Amended by D7` name the other decision, and both
decisions carry the link.

D1 to D20 were written on 2026-09-27 from the history of the repository, from
`61d64ed` (2024-11-19) to `eae81b9` (2026-09-14). One maintainer wrote every commit. No
commit message has a body, so a reason comes from the README, a code comment
or the code. If none of them gives a reason, the entry says "Reason not
recorded". A line number is at `eae81b9`.

A number such as D3 means a decision in this file. A decision of another
repository names that repository, such as dbmeta D110.

### D1. A dependency outside the standard library lives in a subpackage. Decided.

Commits: `61d64ed` (2024-11-19) added `color`, `yaml` and `toml`. `470f6c2`
(2024-12-27) added `glob`.

The root package imports only the standard library and the `strcase` and
`text` subpackages. Each subpackage that needs another module registers
itself from `init`, and a caller turns it on with a blank import:

1. `color` registers `*colors.Color` from `github.com/kenshaw/colors`.
2. `glob` registers `*glob.Glob` from `github.com/kenshaw/glob`.
3. `yaml` registers a loader that uses `github.com/goccy/go-yaml`.
4. `toml` holds a decoder for `github.com/pelletier/go-toml/v2`.

`go.mod` requires the four modules, and a program that imports only the root
package does not link them.

Reason, from the README: `ox` "is written in pure Go, with no non-standard
package dependencies", and the external dependencies "are optional, requiring
a import of a `xo/ox` subpackage".

### D2. ox builds with TinyGo, and CI tests both compilers. Decided.

Commits: `61d64ed` (2024-11-19), `74b466a` (2024-11-19), `8024ca4`
(2024-11-26), `9c1faab` (2026-03-20), `c0d24d5` (2026-03-21), `1f5bd91`
(2026-03-22), `6c99d89` (2026-09-14).

The first commit has a CI job that runs `tinygo test -v ./...` beside the job
that runs `go test -v ./...`. The TinyGo release is pinned: 0.34.0, then
0.40.1 in `1f5bd91`, then 0.42.0 in `6c99d89`, then 0.44.0 on
2026-11-27.

`go.go` and `tinygo.go` hold the code that differs between the two compilers,
behind `//go:build !tinygo` and `//go:build tinygo`. These are the only build
constraints in the module. Where TinyGo lacks a function, `ox` holds a copy of
it from the Go source tree, and a comment says so. `tinygo.go` copies
`os.UserCacheDir` and `os.UserConfigDir`, `fields` in `cmd.go` copies
`reflect.Type.Fields`, and `asExitError` in `ox.go` stands in for
`errors.AsType`.

Reason, from the README: "Work with TinyGo out of the box", and "Minimal use
of reflection (unless TinyGo supports it)".

### D3. The package is named ox. Decided.

Commit: `9152503` (2024-11-21).

The package was `kobra` for its first two days. `9152503` renamed it to `ox`
in every file, the module path and the README.

Reason not recorded.

### D4. A sentinel error is a constant of type Error. Decided.

Commits: `61d64ed` (2024-11-19), `d9844bc` (2024-11-21), `e641803`
(2024-11-21).

`ox.go` declares `type Error string` with an `Error` method, and every
sentinel error is a constant of that type with an `Err` prefix, such as
`ErrUnknownFlag`. There are 26 of them in one `const` block in `ox.go`. An
error with more to say is a struct type, such as `SuggestionError`, whose
`Unwrap` returns the constant. A caller compares with `errors.Is`.

Reason not recorded. A constant cannot be reassigned by an importer, and a
variable can.

### D5. Every option is an option value, and one Option interface carries it. Decided.

Commits: `61d64ed` (2024-11-19), `d571846` (2024-11-30).

Every option func, such as `Usage` or `Default`, returns the unexported
struct `option`. The struct holds a name and one func for each thing that an
option can change: a `*Context`, a `*Command`, a `*Flag`, a `*CommandHelp`
and a type. A `post` func runs after the command tree is built. `Option` is an
interface with one method, `Option() option`. `ContextOption`,
`CommandOption`, `FlagOption` and the others are aliases of `Option`, so
their names document where an option goes, and the compiler does not enforce
it. So one list of options can go to `Run`, and each target takes the funcs
that apply to it. `option.apply` returns `ErrAppliedToInvalidType`, with the
name of the option, when an option has no func for its target.

Reason, from the README: "Functional option and interface smuggling".

### D6. The context, bind and reflect styles build one command tree. Decided.

Commits: `6f3ff60` (2024-11-20), `8a7a9f9` (2024-11-21), `dc26437`
(2024-11-21), `8f9651f` (2025-01-12), `1f506d1` (2025-01-16).

A caller builds a command in one of three styles, and each builds the same
`Command` and `FlagSet`:

1. The context style reads each value from the `context.Context` in `Exec`.
   See D7 and `_examples/context`.
2. The bind style gives each flag a pointer to a variable. See `Bind`,
   `BindSet` and `_examples/bind`.
3. The reflect style reads the flags from the `ox` tags of a struct. See
   `From`, `FlagsFrom` and `_examples/reflect`.

Reason, from the README: "Simple/flexible APIs for Reflection, Bind, and
Context style use cases".

### D7. A command gets the run context from its context.Context, and otx reads it. Decided.

Commits: `54855dc` (2024-11-25), `e487856` (2024-11-26).

`Context.Run` passes a `context.Context` to `Command.Exec` that carries the
`*ox.Context`. `WithContext` adds it and `Ctx` returns it. `e487856` moved the
typed accessors, such as `otx.Get[T]` and `otx.Slice[E]`, to the `otx`
package. The message of `54855dc` calls it a "try at using separate context
object".

Reason not recorded.

### D8. Each default is a package level variable that a caller can replace. Decided.

Commits: every commit that added a `Default` variable, from `61d64ed`
(2024-11-19).

The `var` blocks at the top of `ox.go`, `interp.go` and `conf.go` hold every
default: widths, precisions, the version, error and completion funcs, the
interpolation ops and filters, and the config loaders. Each is exported and
named with a `Default` prefix. A program changes the behavior of `ox` when it
assigns one before it calls `Run`. `text` holds each string that `ox` writes,
as a variable, for the same reason.

Reason, from the README: "No magic, sane defaults, overrideable defaults".

### D9. Help, version and completion are built in, and the completion scripts come from cobra. Decided.

Commits: `2b20141` (2024-12-02), `b36e2ba` (2024-12-04), `2fccfad`
(2024-12-05), `a45c515` (2024-12-06), `c1991ca` (2024-12-06), `7ec4492`
(2024-12-12), `b01d503` (2025-01-23).

`ox` writes the help and the version of each command, and it writes a
completion script for bash, zsh, fish and PowerShell. The script calls the
program with `__complete`, and `Context.Comps` answers. `b01d503` moved the
scripts to `comp/`, and `defs.go` embeds them.

The scripts started as the scripts of `cobra`. `c1991ca` added
`LICENSE.completions.txt`, which holds the Apache License of `cobra`, and the
README says that the scripts are available under that license.

Reason, from the README: "Standard help, version and shell completion".

### D10. Only a root that has sub commands gets a help command and a version command. Decided.

Commits: `dd0fd4b` (2024-12-20), `97980da` (2025-05-09), `9732515`
(2025-05-09).

`dd0fd4b` makes the `Version` option add a `version` command when the command
has sub commands, and a `--version` flag when it has none. `9732515` stops
the `Help` option from adding a `help` command to a sub command. A sub command
still gets the `--help` flag.

Reason, from the title of `9732515`: "Stop adding help command to sub
commands by default". `97980da` fixed a "double help command issue" on the
same day.

### D11. An unknown command gets a suggestion within a Levenshtein distance of 2. Decided.

Commits: `833e93a` (2024-12-07), `e81aab3` (2024-12-17), `063e988`
(2024-12-18), `2171252` (2025-01-02).

If a command has sub commands and the first argument names none of them,
`Command.Suggest` looks for the nearest name, alias or suggested name with
`Ldist`. `DefaultMaxDist` is 2, and `CommandHelp.MaxDist` changes it for one
command. `DefaultSuggestionsEnabled` turns the search off. A command name
matches only in the same case, and the README lists "Case sensitive" as a
design choice. The search for a suggestion changes both names to lower case
first.

A `TODO` at `cmd.go:288` asks for settings that turn off case sensitivity
and the distance match. See B10.

Reason, from the README: "Suggestions for command names, aliases, and
suggested names".

### D12. golangci-lint enables every linter and disables a list. Decided.

Commits: `615436e` (2024-12-04), `cd71c50` (2024-12-07), `5ea04fb`
(2025-05-08).

`.golangci.yml` sets `default: all` and disables 24 linters by name, such as
`errcheck`, `gochecknoglobals` and `varnamelen`. No workflow runs
`golangci-lint`.

Reason not recorded. The disabled list matches the code: D8 needs package
level variables, and the code discards write errors with `_, _ =`.

### D13. The examples include command trees generated from the help of real tools. Decided. Amended by D31.

Commits: `dc5247a` (2024-12-21), `4b2a76f` (2024-12-22), `e83cd7d`
(2024-12-24), `a7bcbd9` (2025-05-16), and the other commits that change the
generated files, 15 in all.

`_examples/gen.go` runs a tool with `help`, parses the output, and writes an
`ox` program for it in `_examples/gen/<tool>`. Its default list is `docker`,
`doctl`, `gh`, `helm`, `hugo`, `kubectl`, `podman`, `psql`, `rclone` and
`talosctl`, and `_examples/gen/omnictl` is also committed. The generated
programs are committed.

`go test ./...` does not build a directory that starts with `_`, and no CI
job builds the examples. Today every generated program builds, and
`_examples/gen.go` does not. See B14.

Reason, from the comment in `_examples/gen.go`: it generates "xo/ox style run
entries for well-known, commmon commands". No text says why.

### D14. A default value is expanded with bash-like interpolation. Decided.

Commits: `b6239e0` (2024-12-04), `45e4d73` (2024-12-04), `4f9efa0`
(2024-12-05), `1dc9b16` (2024-12-07), `89032b1` (2026-03-22), `04b9df2`
(2026-03-24), `93e74e5` (2026-03-24), `0306c0c` (2026-03-25).

`Context.Populate` expands the default of each flag with
`Context.Interpolate`, which is `InterpolateVar` unless the caller changes
it. A variable is `$KEY`, `${KEY}`, `$TYPE{KEY}` or `${type::KEY}`. After the
key come operators, such as `#`, `/`, `||default` and `@%f`, and filters, such
as `|upper`. `DefaultKeyLoader` answers keys such as `HOME`,
`APPNAME`, `APPCONFIG`, `STATE` and `NUMCPU`, and `DefaultEnvLoader` answers
from the environment. `Context.Override` answers first. `0306c0c` refuses a
`${` that crosses a line.

The message of `89032b1` says "adding bash style string operations". The
comment on `InterpolateVar` lists more operators than the code has. See B1
and B2.

Reason, from the README: "Deferred default value expansion".

### D15. An unknown key expands to nothing, and a failed lookup writes its error into the value. Decided.

Commits: `1428040` (2026-03-27), `133ef03` (2026-03-27), `cf6e242`
(2026-03-27), `0b557cf` (2026-03-27).

`InterpolateVar` returns no error for a variable. If no loader knows the key,
the variable expands to the empty string, as in bash. If a loader or an
operator fails, the value holds the error in place of the variable, such as
`!(ERROR: ${X^^}: not implemented: invalid op "^^")`. Four commits in one day
made that form, and `cf6e242` added the `!`.

Reason not recorded.

### D16. A config loader is registered for a ConfigType. Decided.

Commits: `f32c6d5` (2024-12-24), `2f57aab` (2026-03-24), `ec9141c`
(2026-03-30), `caf9b5e` (2026-03-30).

`RegisterConfigLoader` takes a `ConfigType`, such as `YAMLT`, and a func that
answers one key. `DefaultLoader` asks the loaders in the order that they were
registered, and `RegisterConfigLoaderOrder` changes the order. `ox` registers
`AnyT` for `DefaultKeyLoader` and `EnvT` for `DefaultEnvLoader`. The `yaml`
subpackage registers `YAMLT`. `ConfigType` also names `JSONT`, `HCLT` and
`TOMLT`, and nothing registers them.

Reason, from the README: "Enable registration for config file loaders, types
with minimal hassle".

### D17. The go directive stays at a release that the pinned TinyGo can build. Decided. Amended by D24.

Commits: `55bbf56` (2025-05-08), `0a943b5` (2026-03-05), `9c1faab`
(2026-03-20), `c0d24d5` (2026-03-21), `25b315d` (2026-08-26).

The `go` line was 1.23, then 1.24.2, 1.24.3 and 1.26. `9c1faab`, "Fix
compatibility issue with tinygo", lowered it from 1.26 to 1.24.2. `c0d24d5`,
"Update tinygo support to go1.25", raised it to 1.25. `25b315d` raised it to
1.27, to use the standard `uuid` package, and `6c99d89` moved CI to TinyGo
0.42.0 after it.

Reason, from the two titles: the directive follows what TinyGo can build.
See D2.

### D18. ox makes no promise of backward compatibility. Decided.

A maintainer said that no `xo` project makes a promise about its API. tblfmt D37
records it on 2026-09-27. `ox` has no tag, and `30e013b` (2026-03-30), "Rename
Path -> Realpath", is one of the renames that the history holds.

A change can rename or remove an exported identifier when that makes the code
better. `dbtpl` requires `ox`, so a change that breaks `dbtpl` names the use
in its commit message.

### D19. The substring operator takes an end index and not a length. Proposed.

Commit: `89032b1` (2026-03-22).

In bash, `${X:1:2}` of `aXbXc` is `Xb`: two characters from offset 1. In
`ox` it is `X`, because the second number is an end index, as in the Go
expression `s[1:2]`. A negative number counts from the end in both, so
`TestInterpolate` gives the same result in bash and in `ox`. `notes.txt`
describes the operator as `${VAR:offset:length}`.

The code does it, and no record shows that a maintainer chose the difference from
bash. This is open question 10.

### D20. An app builds its command tree from the doc comment of a struct. Proposed. Amended by D30.

Commits: `49a01e6` (2026-03-27), `2e2b6f4` (2026-03-30), `23b1eda`
(2026-03-30), `739fd0c` (2026-03-30).

`RunApp`, `RunAppContext` and `NewApp` take a pointer to a struct, and
`ParseCommand` is to build the command tree from its doc comment.
`_examples/app` shows the idea. The title of `23b1eda` is "Rough pencil in of
how app should work", `RunApp` and `RunAppContext` have empty doc comments,
and `UserConfigFile` does nothing. See B3, B4 and B5.

This is a sketch and not a finished decision. This is open question 2.

### D21. ox is set up for coding agents as every xo repository is. Decided.

A maintainer decided on 2026-09-27 that every repository in the `xo` namespace is set
up for coding agents the same way. dbmeta D110 is the standard. This entry
records what `ox` adopted.

1. `AGENTS.md` holds the rules for a coding agent, because Codex reads that
   file. It opens with the three standing rules of dbmeta D110. `ox` had no
   `CLAUDE.md` and no `AGENTS.md` before.
2. `CLAUDE.md` holds one line, `@AGENTS.md`, which Claude Code reads as an
   import. It is an ordinary file and not a symbolic link.
3. `CONTRIBUTING.md` is new. Its Agent skills section holds the command that
   installs each skill.
4. `simple-english` and `go-pedantry` are ordinary folders in
   `.agents/skills` and in `.claude/skills`, from `npx skills@1.7.0 add` with
   `--copy`. `skills-lock.json` names the source of each.
5. `.gitignore` ignores `.claude/settings.local.json`, the Claude Code
   permissions of one person. `.gitattributes` holds `* text=auto eol=lf`.
6. `docs/PLAN.md` and `docs/BACKLOG.md` are new.

`ox` is a small library, so its decisions stay in this file, with the index at
the top. dbmeta D111 draws that line.

Five tests in `repo_test.go` hold the setup. `TestSkillsAreCopies` fails on a
symbolic link, on a missing copy, on two copies that differ, and on a skill
folder that `skills-lock.json` does not name. `TestClaudeImportsAgents` fails
unless `CLAUDE.md` is an ordinary file that holds only `@AGENTS.md`.
`TestTheRootHoldsFourDocuments` fails on any other `.md` file in the root.
`TestTheDecisionIndexIsComplete` fails unless each decision heading has a
row in the index above, with the same title and status.
`TestAnAmendmentPointsBothWays` fails unless an amended decision names the
decision that amends it. The tests use only the standard library.

A Windows checkout writes a symbolic link as a small text file. Claude Code
then finds a file where it expects a folder, and it loads no skill and says
nothing. That is why the skills and `CLAUDE.md` are ordinary files.

### D22. ox is the command-line package for the maintainers' programs, and usql moves to it. Decided.

A maintainer said on 2026-10-15 that `usql`, `iv` and `wallgrab` are all projects that
use `ox`, among many others. `usql` has not moved to `ox` yet.

On that date:

1. `dbtpl` requires `ox` at `2d66424` (2026-06-22).
2. `iv`, at `github.com/kenshaw/iv`, requires `ox` at `bb2641a` (2026-09-27).
3. `wallgrab`, at `github.com/kenshaw/wallgrab`, requires `ox` at `6fcd48b`
   (2026-08-26).
4. `usql` requires `cobra` v1.10.2 and not `ox`.

The README keeps `usql` in its list of applications. The README linked `iv`
to `github.com/xo/iv`, which does not exist, and now links to
`github.com/kenshaw/iv`.

This answers what was open question 4.

### D23. A tool imports the completion scripts from each cobra release. Decided.

A maintainer said on 2026-10-15 that the import of the cobra completion
scripts must be done regularly, and that it was done by hand before. The
maintainer asked for the import to be done again, and for the steps to be
written down.

`_tools/cobraimport` reads the script templates from a cobra release and
converts them for `ox`. Then it makes the ox edits. It downloads cobra with
`go mod download`, so `ox` does not require cobra (D1). It is in `_tools/`
and not in `comp/`, because `defs.go` embeds every file in `comp/`.
[COBRA-IMPORT.md](COBRA-IMPORT.md) holds the steps, the edits and the tests.

No commit message named the cobra version of the first import. The scripts
that `d1583b0` (2024-12-06) added match cobra v1.8.1 apart from the ox edits.
They match each cobra commit from v1.8.0 to `3a5efae` (2024-10-16) equally,
and `5a138f1` (2024-11-04), which changed the PowerShell script, matches
less well. The local cobra checkout is at `0745e55` (2024-12-28), which is
later than the import.

The first run imported cobra v1.10.2. Bash and PowerShell changed, and zsh
did not. It also corrected the fish template, which had `%[1]s` and `%[2]s`
swapped since the first import. A program with a `-` or a `:` in its name
got fish completion under the wrong name.

### D24. ox supports the newest Go release and the two before it. Decided. Amends D17.

A maintainer decided on 2026-11-26 that `ox` targets Go 1.25 and later, uses
the features of Go 1.27, and trails the newest Go release by two releases.
So `go.mod` declares `go 1.25` while Go 1.27 is the newest release, and it
moves to `go 1.26` when Go 1.28 comes out.

A feature that only a newer release has goes in a file with a build
constraint for that release, such as `//go:build go1.27`. The only one
today is the standard `uuid` package. This changes rule 2 of `AGENTS.md`,
which allowed no build constraint other than `tinygo`. D17 still holds: the
pinned TinyGo must build the minimum release. [DESIGN.md](DESIGN.md) has
the measurements, and B30 has the work. How to read "uses the features of
Go 1.27" is an open question below.

### D25. The GitHub project takes discussions, and not issues or pull requests. Decided.

A maintainer asked on 2026-11-26 that issues and pull requests be turned off
and discussions turned on. That was done the same day. The project also got
19 topics, so that it is easy to find. Two flagship models, GPT-6 Astra and
Gemini 3.1 Pro, each ranked 20 topics, and the topics are the ones that both
named. `yaml` and `toml` were left out, because config loading is not
finished (B6). The next ones that one model named took their place:

`go`, `golang`, `cli`, `command-line`, `command-line-interface`,
`command-line-parser`, `cli-library`, `flags`, `flag-parser`,
`argument-parser`, `tinygo`, `shell-completion`, `autocomplete`,
`subcommands`, `configuration`, `generics`, `posix`, `struct-tags` and
`args`.

Later on 2026-11-26, config loading worked (D34), so `yaml` and `toml` were
added, as both models had ranked them. GitHub allows 20 topics, so `args`
went, which only one model named, last. A fresh ranking was asked for, and
the model service failed to answer three times, so this follows the first
ranking. The topics are now the 17 that both models named, and
`subcommands`, `autocomplete` and `struct-tags`.

`CONTRIBUTING.md` now sends a person to the discussions.

### D26. A path flag is a PathT, a FileT or a DirT. Decided.

A maintainer agreed on 2026-10-15 to three path types: `PathT` for any path,
`FileT` for a file that exists, and `DirT` for a directory that exists. How
a path is resolved or shown is an option on each, and not a type of its
own: `Abs`, `Real` and `Home`. `Ext` narrows file completion, and
`NotExist` makes a path that must not exist yet. A maintainer's notes named
six types, `Dir`, `File`, `Path`, `Realdir`, `Realfile` and `Realpath`. With
the resolution as an option, three types give the same choices.
[DESIGN.md](DESIGN.md) has the table. Questions 1 to 6 of D28 are about
these types.

### D27. ox completes flag values and positional arguments as cobra does. Decided.

A maintainer asked on 2026-10-15 that `ox` complete flag values, with file
and directory completion for a path, and asked on 2026-11-26 that `ox`
support everything that the scripts of cobra support for positional
arguments. The scripts in `comp/` already do this. `ox` must give them the
completions and the directives. [DESIGN.md](DESIGN.md) maps each feature of
cobra to `ox`, and B31 has the faults that a probe found.

### D28. One CompFunc completes a flag or an argument. Proposed.

The API that D27 needs. One option, `CompFunc`, works on a command and on a
flag, and a directive, `CompNoFilter`, turns off the prefix match that `ox`
does now that the scripts do not. [DESIGN.md](DESIGN.md) has the signature
and 12 questions, each with a recommendation. A maintainer must answer them
before the work starts. This is open question 3 below.

### D29. A program can make its help look like psql's help. Decided.

A maintainer asked on 2026-11-26 that `usql --help` on `ox` look as close to
`psql --help` as it can. `ox` already writes psql's sections. A diff of a
copy of psql's options against psql 18.7 found six differences that `ox`
must remove, and one fault where a hook loses its placeholder. They are B32
and B33. The `usql` session proposed its own layout, which a maintainer of
`usql` must confirm. [DESIGN.md](DESIGN.md) has both.

### D30. AppFrom builds a program from a struct and its doc comments. Decided. Amends D20.

A maintainer decided on 2026-11-26 that a program can pass a pointer to a
struct to `ox.AppFrom`. The doc comments of the struct and of its fields
give the usage and the help, the fields give the flags, an interface that the
struct has gives the sub commands, and an `Exec` method gives the exec
func. This is the direction of D20, which was a sketch. Go keeps no doc
comment in a program, so how `AppFrom` gets them, and what the interface is
called, are open questions below. [DESIGN.md](DESIGN.md) has the design.

### D31. An importer writes ox programs from the interface of well-known programs. Decided. Amends D13.

A maintainer asked on 2026-11-26 that `ox` read the command-line interface of
well-known programs, written in Go or not, and write `ox` programs from it in
each style: built by hand, reflection, and app. The importer replaces
`_examples/gen.go` (D13) and lives in `_tools/oximport`. A program built with
cobra answers `__complete`, which gives its commands, flags and allowed
values more exactly than its help, so the importer reads both.
[DESIGN.md](DESIGN.md) has the design, and B39 and B40 have the work.

### D32. ox has many examples for coding agents. Decided.

A maintainer asked on 2026-11-26 for many more examples for coding agents,
in particular for the escaping of a comma in the description of a struct
tag. `docs/EXAMPLES.md` will hold them by task, and each will also be an
`Example` func, so that `go test` proves it. [DESIGN.md](DESIGN.md) has the
first set and the escaping rules, which a probe measured. B41 has the work.

### D33. Each sub package with a dependency is a module of its own, and test is a module. Decided.

A maintainer decided on 2026-11-26 that `ox` adds a `test` package with its
own `go.mod`, and that each sub package that needs another dependency gets
its own `go.mod` and its own tags. So `color`, `glob`, `yaml` and `toml`
become modules, tagged such as `yaml/v0.1.0`, and the root module requires
nothing. This changes how D1 is held, and not what it says: a program that
imports only `ox` still gets no other module. [DESIGN.md](DESIGN.md) has the
layout, the CI and the order of tagging, which two flagship models reviewed.

### D34. A flag takes its value from the command line, the environment, the config files or its default. Proposed.

A maintainer asked on 2026-11-26 that config loading work properly. This
answers what was open question 11. The design in [DESIGN.md](DESIGN.md) went
to Gemini 3.1 Pro and GPT-6 Astra before any code was written, and it
follows their review: one `Config` interface for a file, `ConfigFiles`,
`ConfigFlag`, `EnvPrefix`, `EnvVar`, `NoEnv` and `NoConfig`, flat keys by flag
name, a value from a file that assigns and does not append, and a default
that is expanded only when it wins. No maintainer has chosen the names or
the rules, so this is proposed. Open question 1 asks a maintainer to
confirm it. It replaces the stub `UserConfigFile`, and amends nothing that
was decided, because D16 decides only that a loader is registered for a
`ConfigType`.

### D35. The choices that the agents made on 2026-11-26. Proposed.

A maintainer asked on 2026-11-26 that agents implement the plan. Five agents
did B33, B32, B25, B20, B31, B37, B38, B30, B34, B26, B39, B40 and B41, each in a
worktree of its own, and the work was merged in the main checkout. They
followed the recommendations of the open questions, and each made choices
that no maintainer has made. This entry lists them, so that a maintainer can
confirm or change each one. It holds D28 as built, and the answer of open
question 7 as built.

Help, with `HelpColumn(n)` and `NoDefault()`:

1. `HelpColumn` is opt in, so the help of every other program is the same,
   byte for byte. `TestGenHelpIsTheSame` holds the help of 14 programs.
2. `HelpColumn` does not change `text.FlagSpecSpacer`. A program that wants
   psql's `--file=FILENAME` sets the spacer to `=` itself.
3. Under `HelpColumn`, the short names come first, then the name and the long
   aliases, and only the last name has the spec: `-v, --set=, --variable=NAME=VALUE`.
   A hidden or a deprecated alias is not written.
4. An option goes on its own line when it leaves fewer than 2 spaces before
   the column.
5. `HelpColumn` also lays out the section of sub commands.
6. A hook flag shows its spec only when it takes a value. The help agent
   first showed it for every hook, and the merge found that the spec of a
   command then reached `--help`, so it was changed.
7. An empty line inside a usage gets no indent.

Completion, which holds D28:

8. `HiddenAliases(...)` parses, and completion offers it, and help does not
   write it. `DeprecatedAliases(...)` parses, and nothing shows it. `usql`
   needed both.
9. The debug lines of B20 are removed, and not written when a variable is
   set.
10. When a command or a flag has a `CompFunc` and allowed values, the
    `CompFunc` wins. Cobra uses `ValidArgs` first.
11. When allowed values exist and none matches, or the position is past the
    maximum of `ValidArgs`, `ox` returns `CompNoFileComp`, as cobra does. File
    completion comes only when a value has no source of completions.
12. A value matches by prefix with case, as in cobra. A correction compares
    in lower case. Command and flag names keep their match without case.
13. Completions from commands, values and a `CompFunc` are merged, and only
    those with the smallest distance stay, so a correction shows only when
    nothing matches by prefix.
14. `CompNoFilter` is 64, the next bit, and `Comps` removes it before the
    shell sees it.
15. When `Comps` fails, `DefaultComp` writes the error to standard error and
    prints the directive `:1`.
16. The bash script stops its prefix match in the path without descriptions
    too (question 11 of D28).
17. A flag that the command line already set is still offered, and a hook
    still runs during completion.

Platform:

18. `FlagSet.UUID` and `UUIDT` have no build constraint, so with Go 1.25 or
    1.26 a `UUID` flag fails when the program runs, as `Color` and `Glob`
    do without their modules.
19. A hook built with `Hook` gets a context with `DefaultContext` as its
    parent, and not the context of `RunContext`, because `Context` keeps no
    parent.
20. `ErrInvalidArgCount` comes from `Context.Validate`, and not from `Parse`.
    The other parse errors wrap `ErrUnknownFlag`, `ErrMissingArgument` or
    `ErrInvalidValue`. `TestParseErrors` holds 24 of them, because `usql`
    matches them with `errors.Is`.

The importer, `_tools/oximport` (D31):

21. It asks `__complete` only when the help is not GNU, psql or Go flag
    help, because psql would try to connect to a database.
22. It runs each program with `LC_ALL=C`, and keeps the rest of the
    environment.
23. It asks `__complete` for sub commands only on the root, and on a
    command whose help lists sub commands, because the argument completion of
    a leaf command can use the network.
24. It reads short names from `__complete <path> -`. A flag that cobra lists
    before `--help` is required, and `ox` has no required option, so that is
    a code comment.
25. It writes `ox.Valid` only when the description names every value that
    completes. Otherwise it writes a `// completes:` comment.
26. A flag of unknown type that `__complete` follows with sub commands is a
    bool.
27. It takes defaults only from pflag, Go flag and kubectl help, which write
    them the same way each time.
28. It keeps the line breaks of a description.
29. An optional value, `--x[=ARG]`, becomes a string with `NoArg(true, "")`.
30. It writes a help or a version flag only on the root, and only when its
    short name or section differs from what `ox` adds.
31. It stops at a command whose help is the same as its parent's, and at 12
    levels, because helm plugins loop.
32. The reflection style uses a `string` field with `type:X` for a type that
    has no Go type that `ox` knows.
33. The app style takes the doc comment of a field as its usage, word for
    word, adds the doc sections Description, Aliases, Section and Sections,
    and takes sub commands from `Commands() []any`.
34. It writes `<prog>/{hand,reflect,app}/main.go`, with a "Code generated"
    line, and keeps the help that it read in
    `_tools/oximport/testdata/<prog>.txtar`. The recorder writes the home
    directory as `/home/user`, and `TestTestdata` fails if a real home
    directory appears.

### D36. Choices that the backlog agent made. Proposed.

A maintainer asked on 2026-11-27 that one agent close each item of
[BACKLOG.md](BACKLOG.md) in order. The agent followed the recommendation of
each open question, and made the choices below that no maintainer has made.
A maintainer can confirm or change each one.

B32, the optional value in help:

1. `OptionalValue(spec)` sets the spec, sets `NoArg`, and gives the flag the
   empty string as its value without an `=` when no `NoArg` default exists.
   Help writes `--name[=SPEC]` with an `=`, whatever
   `text.FlagSpecSpacer` is, because an optional value must follow an `=`.
2. `HelpValue(value, usage)` writes `--name=value` on a line of its own
   after the flag, in both layouts of help. Completion does not offer these
   values. `Valid` gives a flag the values that completion offers.

B28, the run that returns its error, and the usage string, with the names
of open question 6:

3. `RunError(ctx, opts...)` returns the first error of `NewContext`,
   `Parse`, `Validate` or `Run`, whatever `Context.OnErr` is, and never
   writes it or exits. It returns nil for `ErrExit`, so `--help`,
   `--version` and a completion return nil.
4. `Command.UsageString()` takes no `Context`, so it expands no variable in
   a default. `Context.Expand` on a nil `Context` returns the value as it
   is, where it panicked before.
5. Item 3 of B28 is `ArgUsage(name, usage)`, a help option. Help writes an
   `Arguments:` section after the usage, with the name of
   `text.Arguments`. `HelpLine` holds a name and its usage, for
   `ArgUsage` and for `HelpValue`.

B35, the path types of D26, built as questions 1 to 6 of D28 recommend:

6. A path type checks only a value from the command line. A value from the
   environment or a config file is not checked, as a default is not,
   because `Context.Fill` does not tell them apart.
7. A path type expands only `~` and `~/...`, and not `~user`.
8. `Real()` makes a path that does not exist absolute, as `Abs()` does,
   because a path that does not exist has no link to resolve.
9. A fault gives `ErrInvalidValue` and one of the new constants
   `ErrNotExist`, `ErrExist`, `ErrIsDir` and `ErrNotDir`.
10. `Ext` works on each path type. A `DirT` completes with
    `CompFilterDirs`, and a path with `Ext` completes its extensions with
    `CompFilterFileExt`, as cobra does.
11. A slice of paths gets no expansion and no check.
12. The tag options are `abs`, `real`, `home`, `notexist` and
    `ext:.yaml|.yml`.
13. `Realpath` lost the rest of the path after a symbolic link, and failed.
    The build fixed it, and `TestRealpath` holds it.

B36, `AppFrom` of D30, built as open questions 4 and 5 recommend:

14. `ox.App(v)` is a command option, so that `RunError` and the other
    options work with it. `ox.AppFrom(v, opts...)` is `RunContext` with
    `DefaultContext`, `App(v)` and `Defaults()`, then `opts`.
15. The tool is `github.com/xo/ox/cmd/oxdoc`, in the root module, because it
    needs only the standard library. It writes `oxdoc_gen.go`, with a call
    to `ox.RegisterDoc[T]` for each struct type that has a doc comment or a
    field with one. It skips a generic type.
16. The usage of a field is the first paragraph of its doc comment, so the
    `oximport:` note that the importer writes after a blank line is not
    usage. A tag with a description wins. `From` reads the doc comments too.
17. A field whose tag has no description, and that has no doc comment, is
    not a flag, as before.
18. A section of the doc comment starts at a line that holds only a known
    name and a colon, and ends at the next one. One tab is removed from each
    line. The help indents `Examples:` by two spaces, because gofmt changes
    the indent of a code block in a doc comment.
19. A sub command without a `Usage:` line takes its name from its type,
    without the suffix `Command` and without the name of the type of its
    parent.
20. `Exec` can have any signature that `ox.Exec` takes.
21. `RunApp`, `RunAppContext`, `NewApp` and `ParseCommand` are removed
    (D18). That closes B3 and B4. No program in the `xo` folder used them.
22. `_examples/app` uses `AppFrom`, with its `oxdoc_gen.go` committed. The
    importer writes the app style without the build constraint `ignore`, and
    `TestBuild` runs `go generate`, then builds and runs every style.

B14, the examples:

23. `_examples/gen.go` is removed, and the programs in `_examples/gen/` stay
    as it wrote them, because `TestGenHelpIsTheSame` builds them. Their
    comment names `_tools/oximport` as the replacement.
24. CI runs `go vet` on each folder of `_examples/`, as the importer agent
    recommended for open question 13. `./_examples/...` matches nothing,
    because `...` skips a folder that starts with `_`, so CI lists them with
    `find`.

B42, the sample of `usql`:

25. `_examples/usql` copies the options of the prototype of `usql` of
    2026-11-27, which has `-d`, `-l`, `-O`, `-U`, `-N`, `-D`, `--config` and
    the hidden options of `--dump`. The texts are in the sample, and not in
    a package `text`. It calls `RunError`, and its `Exec` prints the options
    in order.
26. `TestUsqlSample` in the root builds the sample and runs it, as
    `TestGenHelpIsTheSame` does.

B16, the findings of `golangci-lint` 2.13.3:

27. `.golangci.yml` disables `exhaustruct_v5` and `wsl_v5` beside
    `exhaustruct` and `wsl`. With only the new names, the old ones ran and
    gave 1,182 findings.
28. `.golangci.yml` disables `noinlineerr`, because `if err := f(); err !=
    nil` is idiomatic Go. `paralleltest` stays, because many tests change a
    `Default` variable, and that is a question for a maintainer.
29. Real faults that got a code change: `asDuration`, `asSize` and `asRate`
    refuse a `uint64` above `math.MaxInt64` with `ErrInvalidConversion`
    (gosec G115). `Rate.MarshalText` has a value receiver, so a `Rate`
    value is a `TextMarshaler` (recvcheck). A test in `ox_test.go` uses
    `errors.As` (errorlint). The doc comments of `Lookup` and `Override`
    name the func.
30. No change for the rest: `FormattedTime` needs a pointer for `Set`
    (recvcheck), `wastedassign` in `size.go` and `strcase` is style, and
    `badCond` in `opIndex` is a false positive. The finding of B1 closes
    with B1.

B1, the operators `#`, `##`, `%` and `%%`:

31. Each of the four takes a glob pattern, as bash does: `*`, `?` and
    `[...]`, with `[!...]` or `[^...]` for a set that is negated. Any other
    character matches itself, so a value without these characters works as
    before. A pattern cannot escape `*`, because the reader of a variable
    removes each backslash first.
32. As before, the pattern takes the rest of the variable, filters too:
    `${X##*X|upper}` has the pattern `*X|upper`. Write `${X|upper##*X}`.
33. `apply` now calls the op when the second character of a two character
    op starts its value, such as `/|`. It skipped the call before.

B2, the operators `^`, `^^`, `,` and `,,`:

34. They take a glob pattern, as `#` does, and an empty pattern matches each
    character. `^` and `,` change only the first character, and only when it
    matches.
35. `,` is now an op character, so a key cannot hold a `,`, such as the
    union `[0,1]` of a YAML path. `:`, `/` and `|` were op characters
    already.

B7, the `hcl` module, as open question 11 asks:

36. `github.com/xo/ox/hcl` is a module of its own (D33), and it requires
    `github.com/hashicorp/hcl/v2` v2.25.0. It registers `HCLT` for the
    extension `.hcl`.
37. An attribute is a key. A block nests its body under its type, then
    under each label, so `server "db" { port = 5432 }` gives
    `server.db.port`. A second block with the same type and labels adds its
    keys to the first. An expression is evaluated with no variable and no
    function.
38. A value goes through the JSON of `go-cty`, so a number is a
    `json.Number`, as the JSON format of `ox` gives.
39. `hcl/go.mod` and `test/go.mod` declare `go 1.25.0`, and not `go 1.25`,
    because `hcl/v2` v2.25.0 declares `go 1.25.0`, and `go mod tidy`
    raises a lower line. It is the same release.
40. `notes.txt` names JSON and Jsonnet beside HCL. JSON is in the root
    package. Jsonnet is not built.

B3 and B4 closed with B36, because `RunApp`, `RunAppContext` and `NewApp`
are removed.

B10, case and distance:

41. `DefaultFoldCase`, true by default, makes suggestions and the completion
    of command and flag names compare without case, as before. False
    compares with case. The parse of a command name always compares with
    case (D11), and a value of `Valid` matches its prefix with case (D35).
42. `MaxDist(-1)` turns off the distance match for one command, and
    `DefaultMaxDist = 0` turns it off for every command. `MaxDist(0)` keeps
    `DefaultMaxDist`, as before. A suggestion still matches a name that
    differs only in case.

B11, the usage line and `ValidArgs`:

43. With no `Spec`, the usage line follows `ValidArgs`: no `[args]` for
    `ValidArgs(0, 0)`, and the allowed values, such as `[bash|zsh]...`, in
    place of `[args]`. `...` follows when more than one arg is allowed. The
    texts are `text.CommandArgValuesSpec` and `text.CommandArgsMore`. A
    `CompFunc` does not change the help, because its values are not known
    until it runs. `TestGenHelpIsTheSame` holds that no program in
    `_examples/gen` changed.

B12, `[]byte` and `[]rune`:

44. A value converts to a `[]byte` and to a `[]rune`, and to a named type of
    either. reflect cannot tell `[]rune` from `[]int32`, so a `[]int32`
    element takes runes too. `Bind` on a `[]byte` sets its bytes, where it
    tried to append one byte for each value before. `Bind` on a `[]int32`
    still appends ints.

B17, the README:

45. Items 1 and 3 were already true: the README says Go 1.25 and later, and
    HCL loads since B7. The line "Man page generation" stays as a design
    goal, and it says that it is not built yet, because a maintainer wrote
    it.

B18 fixed the four doc comments that named the wrong thing. Its item 5
closed with B36.

B19, the items of `notes.txt`. Items 5, 7 and 10 stay in the backlog:

46. Item 1, the cascade of config files, is D34. Item 9, the path types, is
    B35. Item 4, `time.Duration` in reflection, is B49.
47. `$ANY{KEY}` and `${any::KEY}` ask each loader in order, as `$KEY` does.
48. The value that a flag takes without a value, `NoArgDef`, expands its
    variables, as a default does.
49. `NewCommand` refuses two flags of one command with the same long or
    short name, counting each alias, with `ErrDuplicateFlag`. The name is
    long, as `--H`, the short name is short, and an alias of one letter is
    short, as `Command.Flag` finds them. A sub command
    can have a flag with the name of a flag of its parent. The check found
    two `-L` flags in `Example_psql`, and the second is now psql's `-n`.
50. The `help` command completes the path of a command, such as
    `app help sub <tab>`, with the corrections of `CompCommands`.
51. `From` takes a field whose pointer type has the text or the binary
    marshal pair, and that no program registered. `registerReflect`
    registers the type the first time, with a `Value` that is not generic.
    Its name is the name of the type in lower case, such as `testtext`, or
    the full name when another type has that name.

B21, `.gitignore`:

52. `.gitignore` keeps its patterns for `*.txt`, `*.log` and the shell
    scripts, because they keep `notes.txt` and the files that
    `_examples/tree/gen.sh` writes out of git. Two lines bring back
    `/comp/*` and `/LICENSE.completions.txt`, so a new script in `comp/`
    needs no `git add -f`.

B43, ActiveHelp:

53. `NewActiveHelp(s)` returns a `Completion` that starts with
    `DefaultCompActiveHelpMarker`, `_activeHelp_ `, as cobra's
    `AppendActiveHelp` does. A `CompFunc` returns it beside its completions,
    and the prefix match keeps it. `DefaultComp` drops it when the variable
    of `DefaultCompActiveHelpSuffix`, such as `MY_APP_ACTIVE_HELP`, is `0`.
    `ox` writes no ActiveHelp of its own. cobra's global
    `COBRA_ACTIVE_HELP` has no `ox` match.

B46: help converts a default that is not a string with the conversion of
`ox`, and with `fmt.Sprint` when that fails, such as for a slice.

B47 and B48, completion in zsh and PowerShell:

54. The importer adds `-U` to the `_describe` call of the zsh script, and
    the scripts were imported again from cobra v1.10.2. Only `zsh.zsh`
    changed.
55. `TestShellCompletion` in the `test` module completes in bash, zsh, fish
    and PowerShell, and skips a shell that is not installed. The PowerShell
    harness of 2026-11-26 failed because `my-app` was not in `PATH`.

B49, the types of `Bind` and `From`:

56. A bind of a flag that is not a slice, an array, a map or a hook gets the
    value that the flag parsed, with `BindVal`, and not its text. So a
    count, a duration and a date with its layout reach the variable. A
    slice and a map still get the text, one element at a time.
57. The conversion to a field knows `time.Duration`, `ox.Size`, `ox.Rate`
    and `time.Time` before it looks at the kind. A `time.Time` from text
    tries the layouts of `TimestampT`, `DateTimeT`, `DateT` and `TimeT`.
58. `From` gives a `time.Duration` field `DurationT`, an `ox.Size` field
    `SizeT`, an `ox.Rate` field `RateT`, and a `complex64` field
    `Complex64T`.

B50, the smaller faults of the examples:

59. `Help(Footer(""))` sets `NoFooter`, as `Footer("")` on a command does.
60. On a root with sub commands, `Version()` adds the `version` command, and
    also hooks a flag that the program defines with `Special("hook:version")`.
    It adds no `--version` flag of its own there (D10).
61. An `OnErr` value, such as `ox.OnErrPanic`, sets `Context.OnErr` as well
    as `Command.OnErr`, so `Run` can panic or go on in place of an exit.
62. A flag with `Deprecated(true)` writes `warning: --old is deprecated` to
    `Context.Stderr` when the command line uses it, with
    `text.FlagDeprecated`. A deprecated alias writes nothing, so that an old
    name of `usql` keeps working as before (D35).
63. A flag error names the flag as the command line gave it, so a
    one-letter long flag is `--d`.

B51, the errors that wrap no constant:

64. An error that decodes a config file wraps `ErrInvalidConfig`.
65. An error that a hook or a validator of the program returns gets no
    constant. It passes through with `%w`, so the program finds its own
    error with `errors.Is`. A constant in front changes the message,
    such as the `-h: unknown flag` of `usql`. A validator of `Valid`
    already wraps `ErrInvalidValue`.

## Open questions for the maintainers

An open question is not a decision. When a maintainer answers one, the
answer becomes a decision above, and the question goes. Each question has a
recommendation.

1. D34 is proposed. Recommended: confirm the names and rules in
   [DESIGN.md](DESIGN.md), as two flagship models reviewed them. In
   particular, must `Value.WasSet` stay true only for the command line, and
   is a key the flat name of the flag?
2. D24 says that `ox` uses the features of Go 1.27 and declares `go 1.25`.
   This was built: code that needs a newer release, such as the `uuid` type,
   is in a file with the build constraint `//go:build go1.27`. Is that what
   "uses the features of Go 1.27" means?
3. D35 lists 20 choices that the agents made, which hold D28 as built.
   Recommended: confirm them. Questions 1 to 6 of D28, about the path types,
   are still open. The path types were built on 2026-11-27 as they
   recommend (D36).
4. How does `AppFrom` get the doc comments (D30)? Recommended: a
   `go generate` tool from `ox` writes them into the program, because it
   works with TinyGo and links no parser. The other way embeds the source and
   parses it when the program starts. This was built on 2026-11-27 as
   `cmd/oxdoc` (D36).
5. What does the interface for the sub commands of `AppFrom` call its
   method? Recommended: `Commands() []any`, because `Sub` is already an
   option. The sketch of D20 uses `Sub() []any`. Is the exec method
   `Exec(ctx context.Context, args []string) error`? Both were built
   on 2026-11-27 (D36), and `Exec` can have any signature that `ox.Exec`
   takes.
6. What are the run func that returns its error and the usage string called
   (B28)? Recommended: `ox.RunError(ctx, opts...) error`, and
   `Command.UsageString() string`. These names were built on 2026-11-27
   (D36).
7. psql's layout of help was built as an option, `HelpColumn(27)`, so that
   `dbtpl`, `iv` and `wallgrab` do not change. Recommended: keep it an option.
   Must it become the default for every program?
8. Answered on 2026-11-27: a maintainer tagged `v0.1.0`, then `color`,
   `glob`, `hcl`, `toml` and `yaml`, in the order in [DESIGN.md](DESIGN.md).
9. Issues and pull requests are off (D25). How does a change from a person
   outside the project arrive? Recommended: the person opens a discussion,
   and a maintainer makes the change. `CONTRIBUTING.md` says so now.
10. The substring operator takes an end index, and bash takes a length. Is the
   difference deliberate, or does `ox` follow bash? See D19.
11. The README says that `ox` loads HCL, and imports `github.com/xo/ox/hcl`
    in its example. There is no such package. Does an `hcl` loader come, or
    does the README change? See B7. The `hcl` module was built on 2026-11-27
    (D36).
12. No CI job runs `golangci-lint`, and main had 263 findings. Must CI run
    it, and must the disabled list follow the renamed linters? See B16.
    Since 2026-11-27, the disabled list names both the old and the new
    names, and the root has 156 findings (D36). CI does not run it yet.
13. No CI job builds the examples. Must CI build them, and the programs that
    the importer writes? See B14 and B39. Since 2026-11-27, CI runs `go vet`
    on each example, and `TestBuild` builds each style that the importer
    writes (D36).
14. The repository tests in `repo_test.go` do not run under TinyGo, because a
    build constraint keeps them out. Is that the right choice? See D2 and
    D21.
