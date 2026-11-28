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
| [D12](#d12-golangci-lint-enables-every-linter-and-disables-a-list-decided-amended-by-d49) | golangci-lint enables every linter and disables a list | Decided. Amended by D49 |
| [D13](#d13-the-examples-include-command-trees-generated-from-the-help-of-real-tools-decided-amended-by-d31) | The examples include command trees generated from the help of real tools | Decided. Amended by D31 |
| [D14](#d14-a-default-value-is-expanded-with-bash-like-interpolation-decided) | A default value is expanded with bash-like interpolation | Decided |
| [D15](#d15-an-unknown-key-expands-to-nothing-and-a-failed-lookup-writes-its-error-into-the-value-decided) | An unknown key expands to nothing, and a failed lookup writes its error into the value | Decided |
| [D16](#d16-a-config-loader-is-registered-for-a-configtype-decided) | A config loader is registered for a ConfigType | Decided |
| [D17](#d17-the-go-directive-stays-at-a-release-that-the-pinned-tinygo-can-build-decided-amended-by-d24) | The go directive stays at a release that the pinned TinyGo can build | Decided. Amended by D24 |
| [D18](#d18-ox-makes-no-promise-of-backward-compatibility-decided) | ox makes no promise of backward compatibility | Decided |
| [D19](#d19-the-substring-operator-takes-an-end-index-and-not-a-length-superseded-by-d50) | The substring operator takes an end index and not a length | Superseded by D50 |
| [D20](#d20-an-app-builds-its-command-tree-from-the-doc-comment-of-a-struct-superseded-by-d30) | An app builds its command tree from the doc comment of a struct | Superseded by D30 |
| [D21](#d21-ox-is-set-up-for-coding-agents-as-every-xo-repository-is-decided) | ox is set up for coding agents as every xo repository is | Decided |
| [D22](#d22-ox-is-the-command-line-package-for-the-maintainers-programs-and-usql-moves-to-it-decided) | ox is the command-line package for the maintainers' programs, and usql moves to it | Decided |
| [D23](#d23-a-tool-imports-the-completion-scripts-from-each-cobra-release-decided) | A tool imports the completion scripts from each cobra release | Decided |
| [D24](#d24-ox-supports-the-newest-go-release-and-the-two-before-it-decided-amends-d17) | ox supports the newest Go release and the two before it | Decided. Amends D17 |
| [D25](#d25-the-github-project-takes-discussions-and-not-issues-or-pull-requests-decided) | The GitHub project takes discussions, and not issues or pull requests | Decided |
| [D26](#d26-a-path-flag-is-a-patht-a-filet-or-a-dirt-decided) | A path flag is a PathT, a FileT or a DirT | Decided |
| [D27](#d27-ox-completes-flag-values-and-positional-arguments-as-cobra-does-decided) | ox completes flag values and positional arguments as cobra does | Decided |
| [D28](#d28-one-compfunc-completes-a-flag-or-an-argument-decided) | One CompFunc completes a flag or an argument | Decided |
| [D29](#d29-a-program-can-make-its-help-look-like-psqls-help-decided) | A program can make its help look like psql's help | Decided |
| [D30](#d30-appfrom-builds-a-program-from-a-struct-and-its-doc-comments-decided-supersedes-d20) | AppFrom builds a program from a struct and its doc comments | Decided. Supersedes D20 |
| [D31](#d31-an-importer-writes-ox-programs-from-the-interface-of-well-known-programs-decided-amends-d13) | An importer writes ox programs from the interface of well-known programs | Decided. Amends D13 |
| [D32](#d32-ox-has-many-examples-for-coding-agents-decided) | ox has many examples for coding agents | Decided |
| [D33](#d33-each-sub-package-with-a-dependency-is-a-module-of-its-own-and-test-is-a-module-decided-amended-by-d42) | Each sub package with a dependency is a module of its own, and test is a module | Decided. Amended by D42 |
| [D34](#d34-a-flag-takes-its-value-from-the-command-line-the-environment-the-config-files-or-its-default-decided-amended-by-d43) | A flag takes its value from the command line, the environment, the config files or its default | Decided. Amended by D43 |
| [D35](#d35-the-choices-that-the-agents-made-on-2026-11-26-decided) | The choices that the agents made on 2026-11-26 | Decided |
| [D36](#d36-choices-that-the-backlog-agent-made-decided-amended-by-d41-amended-by-d49) | Choices that the backlog agent made | Decided. Amended by D41. Amended by D49 |
| [D37](#d37-required-flags-and-rules-between-flags-decided) | Required flags and rules between flags | Decided |
| [D38](#d38-inherited-flags-stopping-at-an-argument-and-prefixes-decided) | Inherited flags, stopping at an argument, and prefixes | Decided |
| [D39](#d39-before-and-after-funcs-typed-arguments-and-the-wrap-width-decided) | Before and after funcs, typed arguments, and the wrap width | Decided |
| [D40](#d40-the-kubernetes-form-of-a-size-decided) | The Kubernetes form of a size | Decided |
| [D41](#d41-a-type-that-reads-its-own-text-needs-no-registration-decided-amends-d36) | A type that reads its own text needs no registration | Decided. Amends D36 |
| [D42](#d42-the-color-and-glob-modules-are-removed-decided-amends-d33) | The color and glob modules are removed | Decided. Amends D33 |
| [D43](#d43-a-config-key-starts-with-the-path-of-its-command-and-ini-is-a-package-decided-amends-d34) | A config key starts with the path of its command, and INI is a package | Decided. Amends D34 |
| [D44](#d44-ox-builds-the-shells-of-completion-on-first-use-decided) | ox builds the shells of completion on first use | Decided |
| [D45](#d45-a-port-range-type-and-strftime-layouts-decided) | A port range type, and strftime layouts | Decided |
| [D46](#d46-a-value-of-helpvalue-can-run-funcs-decided-amended-by-d49-amended-by-d51) | A value of HelpValue can run funcs | Decided. Amended by D49. Amended by D51 |
| [D47](#d47-a-bool-flag-can-have-a-negated-form-such-as---no-debug-decided) | A bool flag can have a negated form, such as --no-debug | Decided |
| [D48](#d48-metadoc-writes-a-man-page-and-markdown-docs-decided) | metadoc writes a man page and Markdown docs | Decided |
| [D49](#d49-answers-of-2026-11-28-lint-in-ci-completion-of-help-values-and-fuzz-tests-decided-amends-d12-amends-d36-amends-d46) | Answers of 2026-11-28, lint in CI, completion of help values, and fuzz tests | Decided. Amends D12. Amends D36. Amends D46 |
| [D50](#d50-the-substring-operator-follows-bash-and-base64-and-hex-write-their-text-decided-supersedes-d19) | The substring operator follows bash, and base64 and hex write their text | Decided. Supersedes D19 |
| [D51](#d51-a-hook-flag-checks-its-value-against-valid-decided-amends-d46) | A hook flag checks its value against Valid | Decided. Amends D46 |
| [D52](#d52-a-short-flag-can-take-a-plus-form-as-bash-takes-o-name-decided-amended-by-d53) | A short flag can take a plus form, as bash takes +o name | Decided. Amended by D53 |
| [D53](#d53-a-command-can-read-plus-words-letters-and-a-value-at-the-end-as-bash-does-decided-amends-d52) | A command can read plus words, letters and a value at the end as bash does | Decided. Amends D52 |
| [D54](#d54-a-command-can-read-long-flags-as-bash-does-decided) | A command can read long flags as bash does | Decided |

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
8. `dbtpl`, `iv` and `wallgrab` use `ox`. `usql` requires `v0.1.0` since
   2026-11-27, in its commit `2aff65d`. See D22.

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
package does not link them. D33 made each subpackage a module of its own,
and D42 removed `color` and `glob`.

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

### D12. golangci-lint enables every linter and disables a list. Decided. Amended by D49.

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

### D19. The substring operator takes an end index and not a length. Superseded by D50.

Commit: `89032b1` (2026-03-22).

In bash, `${X:1:2}` of `aXbXc` is `Xb`: two characters from offset 1. In
`ox` it is `X`, because the second number is an end index, as in the Go
expression `s[1:2]`. A negative number counts from the end in both, so
`TestInterpolate` gives the same result in bash and in `ox`. `notes.txt`
describes the operator as `${VAR:offset:length}`.

The code does it, and no record shows that a maintainer chose the difference from
bash. This is open question 10.

### D20. An app builds its command tree from the doc comment of a struct. Superseded by D30.

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

### D28. One CompFunc completes a flag or an argument. Decided.

The API that D27 needs. One option, `CompFunc`, works on a command and on a
flag, and a directive, `CompNoFilter`, turns off the prefix match that `ox`
does now that the scripts do not. [DESIGN.md](DESIGN.md) has the signature
and 12 questions, each with a recommendation. The work followed the
recommendations (D35, D36), and a maintainer confirmed it as built on
2026-11-28 (D50).

### D29. A program can make its help look like psql's help. Decided.

A maintainer asked on 2026-11-26 that `usql --help` on `ox` look as close to
`psql --help` as it can. `ox` already writes psql's sections. A diff of a
copy of psql's options against psql 18.7 found six differences that `ox`
must remove, and one fault where a hook loses its placeholder. They are B32
and B33. The `usql` session proposed its own layout, which a maintainer of
`usql` must confirm. [DESIGN.md](DESIGN.md) has both.

### D30. AppFrom builds a program from a struct and its doc comments. Decided. Supersedes D20.

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

### D33. Each sub package with a dependency is a module of its own, and test is a module. Decided. Amended by D42.

A maintainer decided on 2026-11-26 that `ox` adds a `test` package with its
own `go.mod`, and that each sub package that needs another dependency gets
its own `go.mod` and its own tags. So `color`, `glob`, `yaml` and `toml`
become modules, tagged such as `yaml/v0.1.0`, and the root module requires
nothing. This changes how D1 is held, and not what it says: a program that
imports only `ox` still gets no other module. [DESIGN.md](DESIGN.md) has the
layout, the CI and the order of tagging, which two flagship models reviewed.

### D34. A flag takes its value from the command line, the environment, the config files or its default. Decided. Amended by D43.

A maintainer asked on 2026-11-26 that config loading work properly. This
answers what was open question 11. The design in [DESIGN.md](DESIGN.md) went
to Gemini 3.1 Pro and GPT-6 Astra before any code was written, and it
follows their review: one `Config` interface for a file, `ConfigFiles`,
`ConfigFlag`, `EnvPrefix`, `EnvVar`, `NoEnv` and `NoConfig`, flat keys by flag
name, a value from a file that assigns and does not append, and a default
that is expanded only when it wins. A maintainer confirmed the names and
the rules on 2026-11-28, and that `Value.WasSet` stays true only for a value
from the command line (D50). It replaces the stub `UserConfigFile`, and amends nothing that
was decided, because D16 decides only that a loader is registered for a
`ConfigType`.

### D35. The choices that the agents made on 2026-11-26. Decided.

A maintainer asked on 2026-11-26 that agents implement the plan. Five agents
did B33, B32, B25, B20, B31, B37, B38, B30, B34, B26, B39, B40 and B41, each in a
worktree of its own, and the work was merged in the main checkout. They
followed the recommendations of the open questions, and each made choices
that no maintainer had made. This entry lists them. A maintainer confirmed
each one on 2026-11-28 (D50). It holds D28 as built, and the answer of open
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

### D36. Choices that the backlog agent made. Decided. Amended by D41. Amended by D49.

A maintainer asked on 2026-11-27 that one agent close each item of
[BACKLOG.md](BACKLOG.md) in order. The agent followed the recommendation of
each open question, and made the choices below. A maintainer confirmed each
one on 2026-11-28 (D50).

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
    the full name when another type has that name. D41 widens this to every
    type that reads its own text.

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

### D37. Required flags and rules between flags. Decided.

A maintainer agreed on 2026-11-28 to two features, after a survey of other
Go flag packages, such as cobra, kong, go-flags, go-arg, kingpin and
urfave/cli. The first is a required flag. The second is a rule between the
flags of a command: only one of a group, all or none of a group, and at
least one of a group. An agent built both, and made the choices below. A
maintainer confirmed each choice on 2026-11-28.
The section "Required flags and rules between flags" in
[DESIGN.md](DESIGN.md) holds the API.

1. The names are `Required()`, a flag option, and `OnlyOneOf(names...)`,
   `AllOrNoneOf(names...)` and `AtLeastOneOf(names...)`, command options.
   The command options use the words of the request, because they say what
   the rule does at the call site. cobra calls them
   `MarkFlagsMutuallyExclusive`, `MarkFlagsRequiredTogether` and
   `MarkFlagsOneRequired`.
2. A flag counts as set when it gets a value from the command line, the
   environment or a config file. A default does not count. `Value.WasSet`
   stays true only for the command line, as rule 6 of the configuration in
   [DESIGN.md](DESIGN.md) says. So `Context.Fill` records the names of the
   flags that the environment or a config file filled, in a field of
   `Context` that is not exported. No exported func reads it.
3. `Context.Validate` runs the checks, after `Command.Validate` checks the
   args. So `ErrInvalidArgCount` comes first, then the required flags, then
   the groups in the order that the program adds them. cobra checks in the
   same order. `RunContext` and `RunError` call `Validate` after the parse,
   and before the exec func runs. Completion does not call it.
4. The checks cover the flags and the groups of the command and of each of
   its parents, as `Context.Fill` does. So a required flag of the root is
   required for each sub command, as a required persistent flag of cobra is.
5. A command without an exec func has no check, because `Context.Run` then
   writes its help. A command with a special value has no check either. Those
   are the help, version and completion commands that `Defaults` adds, so a
   required flag of the root does not break `app help`. `--help` and
   `--version` are hooks that stop the parse with `ErrExit`, so `Validate`
   does not run after them.
6. Four sentinel errors name the faults: `ErrRequiredFlag`, `ErrOnlyOneOf`,
   `ErrAllOrNoneOf` and `ErrAtLeastOneOf`. `ErrRequiredFlag` names each
   required flag that has no value, in one error, such as
   `command app: required flag: --name, --token`. `ErrOnlyOneOf` names the
   flags of the group that are set. `ErrAllOrNoneOf` names the group, then
   the flags that are missing. `ErrAtLeastOneOf` names the group.
7. Help writes `(required)`, the text of `text.FlagRequired`, after the
   usage of a required flag, in place of its default. A default never wins
   for a required flag, so help does not write it. Completion is not
   changed.
8. A name in a group is the long name or a long alias of a flag of the
   command or of a parent. The command resolves the names after it has its
   flags, and an alias becomes the name of its flag. A name that matches no
   flag fails `NewCommand` with `ErrUnknownFlag`.
9. In the reflection style, a tag option names the rule and a group:
   `onlyoneof:GROUP`, `allornoneof:GROUP` and `atleastoneof:GROUP`. The
   flags of a struct that have the same rule and the same group name are one
   group, nested structs too. A flag can be in more than one group. The
   struct then holds the whole rule, as the `xor` and `and` tags of kong do.
   The other way was a tag `group:GROUP` and a command option that names the
   group, which splits the rule between two places. A tag option without a
   group name fails with `ErrInvalidValue`.
10. `FlagsFrom` keeps the group of a tag on a field of `Flag` that is not
    exported. `From` and `App` add the groups to `Command.FlagGroups`. A
    program that calls `FlagsFrom` itself adds its groups with the command
    options.
11. A group can have one flag, and a required flag can have a default.
    Neither is an error. The default of a required flag is never used.
12. `_tools/oximport` still writes a code comment for a flag that cobra
    lists as required (D35, item 24). It can write `ox.Required()` now. This
    change does not touch the importer.

### D38. Inherited flags, stopping at an argument, and prefixes. Decided.

A maintainer agreed to three features on 2026-11-28, after a survey of other
Go flag packages: control over which flags a sub command inherits (item 5 of
B19), a stop at the first positional argument, and a unique prefix of a long
flag. An agent built them, and made the choices below. A maintainer
confirmed each choice on 2026-11-28, and items 11 and 13 in particular.
[DESIGN.md](DESIGN.md) holds the API.

The flags that a sub command inherits:

1. A sub command inherits each flag of its parents by default, as before, so
   that no program changes. `Local()`, or the tag option `local`, keeps a
   flag to its command, as a local flag of cobra does. The command line of a
   sub command refuses a local flag of a parent with `ErrUnknownFlag`. No
   option makes every flag of a command local.
2. A local flag of a parent does not hide a flag with the same name of a
   grandparent. The sub command finds the flag of the grandparent.
3. `Local` changes only the command line and completion. When a sub command
   runs, `Context.Fill` still gives a local flag of a parent its value from
   the environment, a config file or its default. So a variable that the
   flag binds keeps its default.
4. `InheritedFlags(true)` is a help option. It writes the flags that a sub
   command inherits in a section after its own flags, with the name of
   `text.GlobalFlags`, `Global Flags`, as cobra does. It is off by default,
   so that the help of each program stays the same. When the root has it,
   `Help` gives it to each sub command. `CommandHelp.InheritedFlags` holds
   it, and `CommandHelp.AddInheritedFlags` writes the section.
5. The section leaves out a local flag, a hidden flag, a deprecated flag,
   and a flag of a parent with the name of a nearer flag. It holds the flags
   of the nearest parent first, in one section, whatever their `Section`.
   It aligns its descriptions on its own, as cobra does.
6. Completion offers the flags that the parse accepts. `Command.Flag` does
   not find a local flag of a parent, and `Command.WalkFlags` skips it.

The stop at the first positional argument:

7. `StopAtArg()` is a command option, and `Command.StopAtArg` holds it. A
   sub command does not inherit it, because each command has its own
   arguments.
8. A sub command name comes before the first positional argument. So
   `StopAtArg` on a root with sub commands still finds `app sub`, and
   `app x sub` gives the root the args `x sub`.
9. A `--` before the first positional argument ends the flags, and the
   parse drops it, as before. A `--` after it is a positional argument, as
   pflag does, so that a wrapper command passes it on. A `-` is a
   positional argument, as before.
10. After the first positional argument, completion completes only
    positional arguments, as it does after `--`.

The prefix of a long flag:

11. `Abbrev()` is a command option, off by default, and `Command.Abbrev`
    holds it. A command takes a prefix when it or a parent has `Abbrev`, so
    a sub command inherits it and cannot turn it off.
12. An exact name always wins over a prefix. This includes the exact name of
    a hidden flag, a deprecated flag, an alias and a deprecated alias.
13. A prefix matches the long names that completion offers: the name, the
    aliases and the hidden aliases of each flag that is not hidden and not
    deprecated. It does not match a hidden flag, a deprecated flag or a
    deprecated alias. So a hidden flag never makes a prefix ambiguous, and
    an error never names it. GNU `getopt_long` has no hidden option, so this
    is a difference from it.
14. A prefix matches the flags of the command and the flags that it
    inherits as one set, because the command line takes them in the same
    place. A local flag of a parent is not in the set, and neither is a flag
    of a parent with the name of a nearer flag.
15. Two names of one flag, such as `--color` and its alias `--colour`, are
    one match. A prefix that matches more than one flag fails with the new
    constant `ErrAmbiguousFlag`. The error names each long name that
    matches, in sorted order: `--ver: ambiguous flag: --verbose, --version`.
16. A prefix compares with case, as the parse of a name does (D36). A short
    flag needs its exact name. The empty name of `--=x` matches nothing.
17. Completion of the value of a flag finds the flag by its prefix too, so
    `--form <tab>` completes the values of `--format`.

### D39. Before and after funcs, typed arguments, and the wrap width. Decided.

A maintainer agreed on 2026-11-28 to three features, after a survey of other
Go flag packages: funcs that run before and after a command, as cobra's
`PersistentPreRun` and urfave/cli's `Before` do, positional arguments with a
name and a type, as kong and go-arg have, and help that wraps at the width of
the terminal. An agent built them, and made the choices below. A maintainer
confirmed each choice on 2026-11-28, and removed the width of the terminal
(item 17).
[DESIGN.md](DESIGN.md) has the design.

The before and after funcs:

1. `Before(f)` and `After(f)` are command options, and set
   `Command.Before` and `Command.After`. Each takes any signature that
   `Exec` takes. A command has one of each, so a later option replaces an
   earlier one.
2. The before func of a command runs for the command and for each sub
   command below it. The before funcs run from the root down, then the exec
   func runs, then the after funcs run from the exec command up. Each func
   gets the `context.Context` and the args of the exec func.
3. When a before func returns an error, the run stops and returns it. No
   later before func, no exec func and no after func runs. This includes the
   after func of a command whose before func ran.
4. The after funcs run also when the exec func returns an error, as the
   `After` of urfave/cli does. Each after func runs, whatever an earlier one
   returns. The run joins the errors of the exec func and of the after funcs
   with `errors.Join`. When only one func returns an error, the run returns
   that error as it is, so that `DefaultErrorHandler` still finds its
   `ErrorCode` method.
5. The funcs run only for a command with an exec func. A command without
   one writes its help and runs no func. A panic in the exec func runs no
   after func.
6. Help, the version and completion run no before func and no after func.
   The `--help` and `--version` flags and completion stop the run in the
   parse. `Context.Run` runs only the exec func of a command with the
   special value `help`, `version`, `comp` or `comp:` and a shell, and of
   each command below one.
7. With `App`, a method `Before` or `After` of the struct sets the func.
   It can have any signature that `Exec` takes, as item 20 of D36 allows for
   `Exec`.

The typed positional arguments:

8. `Arg(name, usage, opts...)` is a command option that adds an `Argument`
   to `Command.Positional`. An `Argument` holds a `*Flag` and `Optional`.
   Each option of a flag applies to the flag of the argument, so a type,
   `Elem`, `Default`, `Valid`, `Bind`, `Spec`, `CompFunc` and the path
   options work as they do for a flag. An `ArgOption` applies to the
   argument itself. `Optional()` is the only one. The `option` struct has a
   new func `arg` for it (D5). `NewArgument` builds an argument.
9. An argument is required, unless it has a default or `Optional()`. A
   required argument cannot follow an optional one, and no argument can
   follow a variadic one. `NewCommand` returns `ErrInvalidArg` for these,
   for two arguments with the same name, and for an argument with the name
   of a flag of the same command, because both keep their value in `Vars`.
10. An argument of type `SliceT` or `ArrayT` is variadic. It takes each arg
    that is left, and each arg is one value, without a split. A variadic
    argument with `Elem(FileT)` or `Elem(DirT)` checks each path. A slice
    flag of paths checks none (item 11 of D36). A list of files is the
    usual case of a variadic argument, so it gets the check.
11. `Context.Validate` converts and checks the args, after
    `Command.Validate` counts them. A missing required argument gives
    `ErrMissingArgument`, and too many args give `ErrInvalidArgCount`. A
    value that does not convert gives `ErrInvalidValue`, as for a flag,
    such as `command cp: arg src: invalid value: x: does not exist`.
12. `Context.Vars` holds each value by the name of the argument. So
    `otx.Get`, `otx.Slice`, `Bind` and `Vars` read an argument as they read
    a flag. A missing optional argument without a default has no value. A
    default is expanded as the default of a flag is, and a default path is
    not checked (item 6 of D36). The environment and the config files never
    fill an argument.
13. The struct tag option `arg` makes a field an argument, in the order of
    the fields, and `optional` makes it optional. `From` and `App` add the
    arguments. `FlagsFrom` leaves them out, because it returns only flags.
    `optional` without `arg` is an error.
14. With no `Spec`, the usage line of help writes the arguments in place of
    `[args]` or the values of `ValidArgs`, such as `src`, `[mode]`,
    `files...` and `[files...]`. The texts are `text.CommandArgSpec`,
    `text.CommandArgOptionalSpec` and `text.CommandArgsMore`. The
    `Arguments:` section writes each argument with its default, then each
    line of `ArgUsage`. `Spec` on an argument sets the name that help
    writes, such as `FILE`.
15. Completion completes the argument at the position of the word: with its
    own `CompFunc`, then with the values of `Valid`, then by a path type, as
    for the value of a flag. A variadic argument completes each position
    after it. An argument of another type falls back to the `CompFunc` of
    the command, then to `ValidArgs`. Past the last argument, when it is not
    variadic, completion offers nothing, and no file names.
16. `ValidArgs`, `ArgsFunc`, `ArgUsage` and `CompFunc` work as before. With
    `Arg`, the funcs of `ValidArgs` and `ArgsFunc` also check the args,
    before the count of the arguments. So a program uses one or the other.

The wrap width:

17. Help does not read the width of the terminal. The agent built a width
    from `$COLUMNS`, then from `stty size`, because the root package has no
    dependency (D1). No other Go flag package starts a process for this.
    kong and kingpin call the `TIOCGWINSZ` ioctl behind a build constraint
    for each system, and cobra, pflag and urfave/cli take the width from the
    caller. A maintainer chose on 2026-11-28 that the caller sets the width,
    with `DefaultWrapWidth`. So a program such as `usql` reads the width of
    the terminal as it wants to, and `ox` starts no process and adds no
    build constraint.
18. `Wrap` let each line after the first pass the width by the indent of
    the description. A maintainer chose on 2026-11-28 to fix `Wrap`. Each
    line of a description now ends inside the width, unless a word is
    wider. This changed the help of each program with a long description,
    such as 6 of the files in `testdata/gen`.
19. `HelpColumn` does not wrap, as before.

### D40. The Kubernetes form of a size. Decided.

A maintainer asked on 2026-11-28 that `ox.Size` and `ox.Rate` write the
form of a Kubernetes quantity, such as `6Gi` or `6 Gi`. A Kubernetes
quantity is a number with a binary suffix (`Ki`, `Mi`, `Gi`, `Ti`, `Pi`,
`Ei`) or a decimal suffix (`k`, `M`, `G`, `T`, `P`, `E`). `%e` and `%E`
already wrote this form, but they round, so they can lose bytes. An agent
made the choices below. A maintainer confirmed each choice on 2026-11-28.

1. The new verb `q` writes the exact size with a binary suffix, and `Q`
   writes it with a decimal suffix. The name comes from "quantity". The
   lower case verb is binary, as with `e` and `z`. `AppendSize` did not use
   `q` or `Q`, and `fmt` gives both to `Size.Format`. `fmt` keeps `p` and
   `T` for itself.
2. The exact form is an integer with the largest suffix that divides the
   size, as in the canonical form of Kubernetes. So 1.5 GiB is `1536Mi`,
   1500 MiB is `1500Mi`, and 6,000,000,000 bytes is `6G` with `Q`. A size
   that no suffix divides has no suffix, such as `1536` for 1536 bytes. Zero
   is `0`. The precision does not apply.
3. Each verb keeps its base. Kubernetes writes a binary quantity between
   -1024 and 1024 in the decimal form, so 1000 bytes is `1k`. `ox` writes
   `1000` with `q`, because the verb names the base.
4. Without the space flag there is no space, as Kubernetes writes. With the
   space flag, as in `% q`, one space comes before the suffix. A value
   without a suffix gets no space.
5. `% e` and `% E` wrote `100 ` with a trailing space for a size below
   1 KiB or 1 kB, because the value has no suffix. They now write `100`. No
   other verb had this fault, because the others write `B`.
6. `ParseSize` reads the exponent form of Kubernetes, such as `6e9`, `1E3`
   and `1e+3`. The exponent follows the number with no space, and a suffix
   cannot follow it. This form does not conflict with the suffix `e` or `E`
   alone, which is still EB. `6e` is 6 EB, and `6e9` was an error before. A
   negative exponent can give a fraction of a byte, and `ParseSize` drops
   the fraction, as it does for `0.5`.
7. `m` and `M` stay MB. Kubernetes reads `500m` as 0.5, because `m` is
   milli there. A byte count has no milli, so this is a difference. A
   memory value of Kubernetes with `m` gives a different size in `ox`.
   `ox` refuses the suffixes `u` and `n` (micro and nano) of Kubernetes with
   `ErrUnknownSize`. `q` and `Q` never write these suffixes, so their output
   always parses back to the same size.
8. `ParseSize` multiplies an integer with a suffix, or without one, as a
   `uint64`, so it keeps each byte. Before, the value went through a
   `float64`, which lost bytes above 2^53, such as `9007199254740993`. A
   value with a fraction or an exponent still goes through a `float64`.
   A value that does not fit in an `int64` gives `ErrInvalidSize` and
   `strconv.ErrRange`. Before, such a value gave an undefined size.
9. A rate takes the same verbs, such as `1536Mi/s`, and `ParseRate` reads
   it back. Kubernetes has no rate.
10. `MarshalText` does not change. It writes `z` with a precision of -2 and
    a space, because help and config output depend on it. No `Default`
    variable chooses its verb. Rule 3 of `AGENTS.md` puts such a variable
    in `ox.go`, and this change stays in `size.go`. A caller that needs the
    exact form calls `FormatSize` or `FormatRate` with `q`. `MarshalText`
    rounds, so 12345678 bytes gives `11.77 MiB`, and a config file that it
    writes can lose bytes. If a maintainer wants the exact form there, a
    `DefaultSizeVerb` in `ox.go` can choose it.

### D41. A type that reads its own text needs no registration. Decided. Amends D36.

A maintainer asked on 2026-11-28 that `ox` take any type without
registration, so that a program does not need the `color` module to read a
`colors.Color`. The maintainer also asked that help write the name of such
a type in lower case, such as `--fg color`. Before, only a pointer field with
the text or the binary marshal pair worked (D36, item 51). A value field
such as `colors.Color` or `netip.Addr` was refused, a named int with
`UnmarshalText`, such as `slog.Level`, was read as an int, and a hand-built
flag could not name a Go type. The maintainer chose the plan below, and an
agent built it.

1. A type reads its own text when it, or a pointer to it, has
   `UnmarshalText`, `UnmarshalBinary` or `Set(string) error`, in that order.
   `Set` is the method of a `flag.Value`. A type needs no `MarshalText`.
2. The method comes before the kind of the type. So `slog.Level` reads
   `debug`, as `flag.TextVar` and `encoding/json` read it. A number that the
   method refuses fails, even when the kind of the type is int.
   `time.Time`, `Size` and `Rate` keep their own parse, which reads more
   forms than their `UnmarshalText`.
3. `From`, `App`, `Bind` and the new option `TypeFor[T]()` take such a type,
   as a value or as a pointer, and as the element of a slice or a map.
   `TypeFor` sets the type of a hand-built flag from a Go type, as `From`
   does for a field. It also takes the types that `From` takes, such as
   `[]string`.
4. The first use registers the type, with the name of the type in lower
   case, such as `color` for `colors.Color` and for `*colors.Color`. When
   another type has that name, the name is the package and the name, such
   as `ox.url`, then the package path and the name. A registered type keeps
   its name, so `*url.URL` stays `url`.
5. A value of such a type is always a pointer. A bind or a conversion to the
   type that is not a pointer copies the value that it points to.
6. A type with `Set` keeps one value, and calls `Set` for each value on the
   command line, as the `flag` package does. So a `flag.Value` that appends
   keeps each value. A type with `UnmarshalText` or `UnmarshalBinary` gets a
   new value each time, so the last value wins.
7. Help writes a value with `MarshalText`, `MarshalBinary` or `String`, and
   else with `fmt`.
8. `As[T]`, and so `otx.Get[T]`, converts a string to such a type without a
   registration, so that two goroutines can read values at the same time.
9. A type assertion, and not `reflect.Type.Implements`, finds the methods,
   so that the code works with TinyGo.
10. The first use of a type registers it, so two goroutines that build
    commands at the same time write the maps of registered types. A
    `sync.RWMutex` guards the maps. The race detector found this in the
    tests of the `test` module, and the old `registerReflect` had the same
    fault.
11. A program does not need the `color` and `glob` modules to read
    `colors.Color` and `glob.Glob`. D42 removed both modules.

### D42. The color and glob modules are removed. Decided. Amends D33.

A maintainer chose on 2026-11-28 to remove the `color` and `glob` modules,
because D41 reads `colors.Color` and `glob.Glob` without them. Before, each
module registered its type from `init`, and gave a name, a method and an
accessor. A program that did not import the module could not use the name or
the method.

1. The root loses `ColorT`, `GlobT`, `FlagSet.Color` and `FlagSet.Glob`. A
   program uses `ox.TypeFor[*colors.Color]()` and
   `ox.TypeFor[*glob.Glob]()`, or a field of the type with `From` or `App`.
   Help writes `color` and `glob`, as before.
2. The accessors `color.Color` and `glob.Glob` go. A program calls
   `otx.Get[*colors.Color]` and `otx.Get[*glob.Glob]`.
3. `color.Default` goes. The zero `colors.Color` is transparent, as
   `color.Default` was. A program sets another default with `ox.Default`.
4. The tags `color/v0.1.0` and `glob/v0.1.0` stay, so that a program that
   pins them still builds against `ox` `v0.1.0`. They do not build against a
   later root, because they name `ox.ColorT` and `ox.GlobT`.
5. `ox` makes no promise of backward compatibility (D18). `dbtpl` uses
   `ox.GlobT` and imports `github.com/xo/ox/glob`. `iv` imports
   `github.com/xo/ox/color`. A maintainer changes both when they move to the
   next `ox`.

### D43. A config key starts with the path of its command, and INI is a package. Decided. Amends D34.

A maintainer asked on 2026-11-28 that the key of a flag of a sub command
start with the path of the command, so that a flag of `a b c` reads
`a.b.c.flag`. Before, each flag read the key of its own name, and a flag of
a sub command read a key at the top of the file. The maintainer also asked
for an INI format, in a package that a program must import, as for the other
formats, and for tests and examples of each format.

1. Without `Key`, the key of a flag is the path of its command below the
   root, then the name of the flag, joined with dots. So `--steps` of
   `db migrate` reads `migrate.steps`, and a flag of the root reads its name,
   as before.
2. A field of a nested struct of `From` or `App` reads the path of the
   field, with `DefaultFlagNameMapper` on each part. So the field
   `Source.Dir` of `db migrate` reads `migrate.source.dir`, and its flag is
   `--source-dir`. A tag option `name:` names the last part. A flag that a
   program builds by hand has no such path, so it uses `Key`.
3. The command of a flag is the command that owns it. So a flag of the root
   that a sub command inherits keeps the key of the root.
4. `Key` gives the whole key. The path of the command does not apply to it.
5. `MapConfig` tries the whole key first, so a flat key with dots, such as
   `"migrate.steps"` in JSON, also matches.
6. The environment does not change. `EnvPrefix` gives `APP_STEPS` for
   `--steps` of a sub command, as before.
7. `github.com/xo/ox/ini` reads INI files, with the extension `.ini`, and
   `INIT` is its config type. It uses only the standard library, so it is a
   package of the root module and not a module of its own (D33). A program
   imports it for its side effect, as for `yaml`, `toml` and `hcl`, so a
   program that does not import it does not link it.
8. INI has sections, and a dot in the name of a section or of a key nests
   it. A key that is written more than once is a list. `;` and `#` start a
   comment on a line of its own. A value in double quotes is a Go string,
   and a value in single quotes is the text between them. Each value is a
   string, and the flag converts it. An error names its line.
9. Each of `yaml`, `toml`, `hcl` and `ini` has the examples `Example_build`
   and `Example_reflect`, which read the same settings, and the README links
   them. `yaml`, `toml` and `ini` have tests of their own.

### D44. ox builds the shells of completion on first use. Decided.

`BenchmarkCommandLine` in `test/bench_test.go` runs the same command line
with `ox` and other Go flag packages (B55). It found that a run of `ox`
with `Defaults` took 40 µs and 195 KB, and a run of cobra took 9 µs and
32 KB, on an AMD Ryzen 9 9950X. A maintainer asked on 2026-11-28 to fix
it.

1. `compTemplates` reads the embedded completion scripts once, and reuses
   them. Before, each command with completion read every script again,
   which was 86% of the memory of a run. A program that sets
   `DefaultCompTemplates` to its own files gets them read each time.
2. The `completion` command builds its sub commands, one for each shell,
   the first time that a method of `Command` reads them, such as `Command`,
   `Lookup`, `CompCommands` or the help of the command. A `sync.Once`
   guards the build. Each shell then gets a help flag, as `AddHelpFlag`
   gives the other commands, when the completion command has one. A
   program that reads the field `Command.Commands` of the completion command
   before such a method finds no shell. The doc comment of the field says
   so. `Context.Validate` returns an error of the build.
3. The help flag of each command, and its `CommandHelp`, are still built
   when the program starts. Each command writes its own `-h, --help` in its
   help, and the parse and the completion need the flag. Building it later
   would change the help of each program.
4. A `Type` used as an option carries the type, and not two funcs, so a
   flag of each type allocates no func. `SplitBy` uses `strings.Split` when
   the text has no backslash. `asBool` and `asDuration` try the word or the
   duration with a unit first, so a value such as `true` or `30s` does not
   try each number parse, which allocates an error. Each gives the same
   result for each value as before.
5. A run of `ox` now takes about 10 µs, 22 KB and 313 allocations, and a run
   of `ox` with a struct about 18 µs, 29 KB and 586 allocations. cobra takes
   about 9.5 µs, 32 KB and 156 allocations.

### D45. A port range type, and strftime layouts. Decided.

A maintainer asked on 2026-11-28 for item 10 of B19: a strftime type, a
semver type, and a port range type such as `100:200`. The maintainer asked
an agent to consult flagship models and to build the best design. Two
flagship models, GPT-6 Astra and Gemini 3.1 Pro, advised the design below,
and an agent built it. A maintainer accepted it on 2026-11-28.

1. `ox.PortRange` is a struct with the fields `Lo` and `Hi`, both of type
   `uint16`. It is in the root package, because it needs only the standard
   library (D1). Its type is `PortRangeT`, which help writes as
   `portrange`. `FlagSet.PortRange` adds a flag of the type. It is
   registered as `*url.URL` is, so a value of the type is a pointer (D41).
2. `ParsePortRange` reads one port, such as `80`, or a range, such as
   `100-200` or `100:200`. Each port is 1 to 65535, and the first port is not
   greater than the last. It refuses port 0, a sign, a space, an empty port,
   such as `100-`, and a port above 65535. Its error wraps the new
   `ErrInvalidPortRange`.
3. `-` is the canonical separator. `String` and `MarshalText` write
   `100-200`, or `80` when `Lo` and `Hi` are the same. The zero `PortRange`
   is not valid. It writes no text, and empty text reads back to it, as with
   `netip.Addr`.
4. `Contains` tells whether a valid range holds a port. `Ports` returns an
   `iter.Seq[uint16]` that walks each port of a valid range. The loop stops
   at `Hi` before it adds 1, so a range that ends at 65535 does not wrap to
   0.
5. A list of ranges is not a part of the type. A slice flag of port ranges
   splits its value at each comma, as each slice flag does, such as
   `22,80-90`.
6. `ox` has no semver type. Since D41, a type that reads its own text needs
   no registration. So a program uses the version type of its own semver
   package. A type in `ox` would need a dependency or a copy of a parser,
   and a program would then have two version types. "A type that nothing
   registers" in `docs/EXAMPLES.md` says this.
7. `ox` has no strftime type. The new flag option `Strftime(layout)`
   converts a strftime layout to a Go layout, and sets the new field
   `Flag.Layout`. A time flag reads and writes its value in that layout. A
   flag of type `TimestampT`, `DateTimeT`, `DateT` or `TimeT` keeps its
   type, and a flag of another type gets `TimestampT`. A slice, an array, a
   map or a hook flag refuses the option with `ErrAppliedToInvalidType`.
   Help writes the type, such as `timestamp`, and not the layout. A program
   can show the layout with `Spec`.
8. B19 said that a strftime type needs a module of its own (D33). That is
   true only for a strftime package from outside the standard library. The
   conversion uses only the standard library, so it is in the root.
9. `Strftime` takes only the directives that Go writes and reads in the
   same way: `%Y`, `%y`, `%m`, `%b`, `%B`, `%d`, `%e`, `%j`, `%a`, `%A`,
   `%H`, `%I`, `%p`, `%M`, `%S`, `%f`, `%L`, `%z`, `%Z`, `%T`, `%R`, `%D`,
   `%F` and `%%`. `%f` is 6 digits and `%L` is 3 digits of the second. Each
   must follow a `.` or a `,`, because Go reads a fraction only there.
10. `Strftime` refuses any other directive, such as `%c`, `%x`, `%l` or
    `%-d`. Its error names the directive and wraps the new
    `ErrInvalidLayout`. It never uses a Go layout that is only close to the
    directive.
11. A Go layout has no escape for text. So text such as `1`, `Jan` or `PM`,
    or text that joins a directive, such as `day` after `%a`, reads as a
    part of the time. `Strftime` refuses such a layout. It writes two times
    that differ in each field with the whole layout, and with each part on
    its own. The text must be the same for each time.
12. The struct tags have no option for a time layout today, so `Strftime`
    has no tag option.
13. `As[time.Time]`, and so `otx.Get[time.Time]`, returned the zero time
    for each time flag, because the conversion gave an `ox.FormattedTime`.
    It now gives the `time.Time`. `ExampleStrftime` found the fault.
14. The new field `Flag.Layout` made `Flag` 520 bytes, which moved each
    flag into the next size class of the Go allocator and cost about 64
    bytes for each flag of a program (D44). The fields `NoArg` and
    `Required` now follow `Optional`, and `NotExist` follows `Home`, so
    that the bool fields share their words. `Flag` is 496 bytes, and a run
    of `BenchmarkCommandLine` uses the same memory as before.
### D46. A value of HelpValue can run funcs. Decided. Amended by D49. Amended by D51.

Before, `HelpValue` wrote a line such as `--help=commands` in the help, but
the help flag wrote the same help for every value. So `usql` replaced the
default of its `hook:help` flag with a func that picked the topic (B52). A
maintainer asked on 2026-11-28 to build B52. An agent made these choices,
and a maintainer accepted them on 2026-11-28.

1. `HelpValue(value, usage string, funcs ...func(*Context) error)` takes the
   funcs as its last argument. A call with no funcs works as before. A
   separate option would add a second name that writes the same line of
   help.
2. When the command line gives the flag the value, the funcs run in order,
   in place of the flag. The flag does not take the value, and its hook
   does not run. Then the run stops with `ErrExit`, as help does, so
   `RunError` returns nil. The first error of a func stops the run with
   that error, and the funcs after it do not run.
3. The funcs run for any flag, and not only for the help flag. So
   `--format=list` can list the formats of a string flag. A value from the
   environment, a config file or a default does not run the funcs, because
   only the command line asks for help.
4. A value of the help flag without funcs, and a value that no `HelpValue`
   names, write the help of the command, as before. `--help` without a
   value writes the help of the command.
5. `HelpLine` has a new field `Funcs`, which holds the funcs. `ArgUsage`
   sets no funcs. The values stay in `Flag.HelpValues`, in one list with
   their usage.
6. Completion does not change. It does not offer the values of
   `HelpValue`, as item 2 of D36 says. A completion parses the words before
   the word that it completes, so a value with funcs on the command line
   runs them, as a hook runs. The completion discards what they write to
   `Context.Stdout`.
### D47. A bool flag can have a negated form, such as --no-debug. Decided.

B53 found that `--no-debug` fails with "unknown flag" for a bool flag named
`debug`. `kong` and `kingpin` add the `--no-` form to each bool flag. A
maintainer asked on 2026-11-28 for an option that a program adds when it
wants the form, and that `Defaults` does not add. The maintainer asked two
flagship models, GPT-6 Astra and Gemini 3.1 Pro, for the design. They
advised the names and the rules below, and an agent built them. A
maintainer accepted the choices on 2026-11-28.

The API:

1. `NegateBools(prefixes ...string)` is a root command option.
   `Command.NegateBools` holds the prefixes. A sub command refuses it with
   `ErrCanOnlyBeUsedWithRootCommand`. Without a prefix, the prefix is
   `DefaultNegatePrefix`, which is `no`. Each prefix adds a form, so
   `NegateBools("no", "without")` gives `--no-debug` and `--without-debug`.
2. A prefix is a word of letters, digits and hyphens, with no hyphen at
   either end, such as `no` or `do-not`. An empty prefix, a bad prefix and
   a prefix that is given twice are errors that wrap `ErrInvalidValue`.
3. `Negatable(false)` is a flag option that keeps one flag out of
   `NegateBools`. `Negatable(true)` gives one flag the negated form without
   `NegateBools`, with the prefix `DefaultNegatePrefix`. With `NegateBools`,
   the flag takes the prefixes of the root. `Flag.Negatable` and
   `Flag.NoNegate` hold the option. The tag options are `negatable` and
   `negatable:false`.
4. The name `NegateBools` was chosen over `NoBoolDefaults`. A name that
   starts with `Defaults` reads as a part of `Defaults`, which adds the
   standard flags and commands. `NegateBools` changes how the command line
   parses, so a program adds it on purpose.

The flags that get a negated form:

5. Only a plain bool flag gets the form: a `BoolT` flag that the command
   line sets to true without a value. A count flag, a hook flag such as
   `--help`, and a bool flag with `OptionalValue` do not. A flag that is not
   a plain bool flag refuses `Negatable(true)` with
   `ErrAppliedToInvalidType`.
6. The form comes from the long name of the flag. An alias and a short name
   get none, and no short flag has a negated form.
7. A flag whose name starts with a prefix and a hyphen gets no negated form,
   so psql's `--no-psqlrc` never reads `--no-no-psqlrc`. This holds for each
   prefix of the root. With `NegateBools("no", "without")`, `--no-psqlrc`
   gets no `--without-no-psqlrc` either.

Where the form is applied:

8. The form is applied when `ox` reads a long name, and not when the
   command tree is built. No flag is added to a `FlagSet`. The parse, the
   completion and the help read the prefixes of the root and the options of
   the flag when they need them. So a sub command or a flag that a program
   adds after `NegateBools` gets the form too. The parse looks for a negated
   form only after the exact name of a flag fails, so a program that does
   not use `NegateBools` does no more work in the parse.
9. A negated form with the long name of another flag, or of a long alias,
   or of another negated form, is an error that wraps `ErrDuplicateFlag`.
   `NewCommand` checks it when it builds each command, with the flags that
   the command accepts, and the error says to use `Negatable(false)`. A
   negated form never hides a flag. A root with `NegateBools` checks each
   sub command again, because a sub command can be built before the root
   gets the option. The check builds no shell of the completion command
   (D44). A flag that a program adds to a `FlagSet` after `NewCommand` is
   not checked. Then the exact name wins in the parse.

The parse:

10. The negated form sets the flag to false, and counts as a value from the
    command line, as the flag does, so `Value.WasSet` is true. A later
    `--debug` or `--no-debug` wins over an earlier one.
11. The negated form takes no value. `--no-debug=true`, `--no-debug=false`
    and `--no-debug=` are errors that wrap `ErrInvalidValue`. The negated
    form does not take the next arg, so `--no-debug false` gives the
    positional argument `false`.
12. The negated form follows the flag. The form of a hidden flag works and
    is hidden. The form of a local flag works only in its command. A sub
    command takes the form of each flag that it inherits. A flag of a parent
    that a nearer flag hides has no negated form in the sub command (D38).
    The form of a deprecated flag writes the warning of the flag.
13. With `Abbrev`, a prefix matches the negated forms as it matches the
    other long names, so `--no-deb` gives `--no-debug` when it is unique. A
    flag and its negated form are two matches, because they set different
    values. So with a flag `nox`, the prefix `--no` is ambiguous between
    `--nox` and `--no-nox`. An exact negated form wins over a prefix. A
    prefix does not match the negated form of a hidden or a deprecated flag.
14. The environment and the config files set the flag itself. No
    environment variable or config key has the name of a negated form.

Help and completion:

15. Help writes a flag with a negated form on one line, as `--[no-]debug`,
    with the first prefix. `text.FlagNegatedSpec` holds the text `[%s-]`.
    In the layout of `HelpColumn`, the text goes before the name of the
    flag, and the long aliases follow without it. The help of a program
    without `NegateBools` and without `Negatable(true)` does not change, so
    the files in `testdata/` do not change.
16. Completion offers each negated form after the names of its flag, and
    corrects a negated form as it corrects a name. It completes no value
    after a negated form, and nothing after `--no-debug=`.
### D48. metadoc writes a man page and Markdown docs. Decided.

B54 found that `ox` writes only help, and that `cobra/doc`, `urfave/cli`,
`kong` and `kingpin` write a man page or Markdown docs. A maintainer chose
this design on 2026-11-28. Two flagship models, GPT-6 Astra and Gemini 3.1
Pro, advised on the design, and then critiqued it. The maintainer chose the
output after a survey of these packages. The hidden `--help-man` flag of
`kingpin`, `mango-kong` for `kong`, and `ToMan` and `ToMarkdown` of
`urfave/cli` write the whole tree as one page. `cobra/doc` writes one file
for each command with `GenManTree`, and one command with `GenMan`.

The maintainer chose these:

1. The package is `github.com/xo/ox/metadoc`. It uses only the standard
   library, so it is a package of the root module, as `ini` is (D43). It
   builds with TinyGo. It has no `init` func, so an import of it changes
   nothing.
2. A program adds it with a root option, such as
   `ox.Run(..., metadoc.Command(metadoc.ManPage(), metadoc.Markdown()))`.
   `Command` adds one hidden command, `__docs`, with a sub command for each
   format: `__docs man` and `__docs markdown`. `ox.Defaults` does not add
   it. When the program has a command named `__docs`, `Command` returns an
   error that wraps `ErrExist`.
3. Help and completion do not show `__docs` or its sub commands, as for
   `__complete`.
4. `app __docs man` writes the whole tree to standard out as one man page:
   the root, then a COMMANDS section that holds each sub command with its
   usage, flags and arguments. `app __docs markdown` writes the whole tree
   as one Markdown document.
5. With `--dir DIR`, each writes one file for each command, such as `app.1`,
   `app-sub.1` and `app-sub-deep.1`, or `app.md` and `app-sub.md`. A man
   page has SEE ALSO references, and a Markdown file has relative links, to
   the parent and to the sub commands. The folder is created when it does
   not exist. Each file has the mode 0o644. A file name that is not in the
   folder is an error.
6. A library API writes the whole tree to an `io.Writer`, one command to an
   `io.Writer`, and the tree to a folder. A program calls it from a test or
   from `go generate`, and then the program does not link the writer.
7. The sections are NAME, SYNOPSIS, DESCRIPTION, OPTIONS, ARGUMENTS,
   COMMANDS, EXAMPLES, the footer and SEE ALSO. The data comes from
   `CommandHelp` and the command, and not from the text of help.
8. The docs leave out hidden commands and flags, deprecated commands and
   flags, `__complete` and `__docs`. They keep the `help`, `version` and
   `completion` commands.
9. The `.TH` line holds the name, the section, the date, the source and the
   manual. The section is 1, and the source is the name and the version of
   the root, by default. The date is the `Date` option, else
   `$SOURCE_DATE_EPOCH` read as Unix seconds in UTC, else today in UTC. A
   value of `$SOURCE_DATE_EPOCH` that is not a number is an error. The date
   is written as `2006-01-02`.
10. A man page escapes backslashes, a `.` or a `'` at the start of a line,
    and the hyphens in flag names. Markdown escapes `<`, `>`, `*`, `_`,
    backticks and `|`, and a flag spec is a code span.
11. The name of the program is the name of the root command, and never
    `os.Args[0]`.

The agent that built it chose these. A maintainer confirmed them on
2026-11-28 (D50):

12. `Command` with no format adds both formats, so that `metadoc.Command()`
    is enough for most programs.
13. `Format` is one type. `ManPage(opts ...ManOption)` and
    `Markdown(opts ...MarkdownOption)` return it. Each format has its own
    option func type, so an option of one format cannot go to the other. The
    man options are `Section`, `Date`, `Source` and `Manual`. The Markdown
    options are `FrontMatter` and `Link`. Each func gets the `*ox.Command`
    of the page.
14. The library API is three funcs that take a `Format`: `Write`,
    `WriteCommand` and `WriteDir`. Six funcs, one set for each format, would
    say the same thing twice.
15. The type of an `ox` option is not exported, so a package outside `ox`
    cannot write an option that changes the tree after the other options.
    So `ox` has a new root option, `ox.Extend(name, f)`, which runs `f`
    after each other option. `metadoc.Command` uses it. So `__docs` comes
    after the help, version and completion of `Defaults`, in any order of
    the options, and the help, version and completion do not change.
16. `ox` knows the special value `docs` of the hidden command. Help does not
    write `[command]` or the default footer for it, and a suggestion for an
    unknown command does not name it. The funcs of `Before` and `After` do
    not run for it, as for help. The flags are not checked for it, so a
    required flag does not stop `app __docs man`.
17. `CommandHelp` has two new methods, `SpecString` and `FlagUsage`, which
    help also uses. So the usage spec and the text of a default are the same
    in help and in the docs.
18. The docs show each default as the program wrote it, such as
    `$HOME/.apprc`. They do not expand it, so the docs are the same on each
    computer, and they show the home folder of no user.
19. The hidden command and its sub commands get a help flag when the root
    has one.
20. The source is the name of the root and its `Version`, from
    `ox.VersionString`. It does not use the version of the build info,
    which changes with each build, such as `v0.2.1+dirty`. The manual is empty by default, so `man` writes the title of
    the section, such as "General Commands Manual". The section is an int.
21. An empty `$SOURCE_DATE_EPOCH` counts as not set. An explicit date is
    written as it is, without a change to UTC.
22. A man page escapes each hyphen as `\-`, in all text and not only in flag
    names, so that a reader can copy a flag or an example. A backslash
    becomes `\e`, and a line that starts with `.` or `'` starts with `\&`.
    An empty line in a description starts a new paragraph. An example is
    written without fill, between `.nf` and `.fi`. The argument of a
    request, such as `.SS`, is in double quotes, and a `"` in it becomes
    `\(dq`.
23. Markdown puts a backslash before each `\`, backtick, `*`, `_`, `<`, `>`,
    `[`, `]`, `|` and `&` in text. It also puts one before a `#`, `-`, `+`
    or `=` at the start of a line, and before the `.` or `)` of a number at
    the start of a line, so that the line does not start a heading or a
    list. A flag spec and the name of an argument are code spans. The
    synopsis and the examples are fenced code blocks. The fence of a code
    span or a block is longer than each run of backticks in its text.
24. The whole Markdown document links each command to its heading, with the
    anchor that GitHub makes, such as `#app-sub`. `Link` applies only to
    the files of `WriteDir` and to the page of `WriteCommand`. `FrontMatter`
    applies to each document.
25. In the whole tree, a sub command lists only its own flags, because the
    flags that it inherits are in the entry of its parent. The page of one
    command has a GLOBAL OPTIONS section with the flags that it inherits.
26. A man page always has a DESCRIPTION, with the banner, or with the usage
    when the command has no banner. Markdown leaves out the description when
    it is the same as the usage.
27. The footer goes in a NOTES section, and only when the program sets it
    with `ox.Footer`. The default footer of help names `--help`, which does
    not apply to the docs.
28. A flag with allowed values, such as from `ox.Valid`, gets
    `(allowed: json, yaml)` after its usage. The text that `metadoc` writes
    is in `text/text.go`, with the prefix `Doc` (D8). `metadoc.DefaultName`
    holds the name `__docs`.
29. The order of the flags and of the commands follows `Sort` and
    `CommandSort` of each command, as help does. The docs do not keep the
    sections of the flags and of the commands.
30. A bool flag with a negated form (D47) shows it as help does, such as
    `--[no-]debug`. `Command.NegatedNames(g)` gives the negated forms of a
    flag in a command, so that `metadoc` reads them as help does. It was
    added when D47 and D48 were merged.

### D49. Answers of 2026-11-28, lint in CI, completion of help values, and fuzz tests. Decided. Amends D12. Amends D36. Amends D46.

A maintainer answered these on 2026-11-28, and asked for the lint fixes,
the completion of help values and the fuzz tests below.

The answers:

1. Open question 2: D24 means what was built. Code that needs a newer Go
   release, such as the `uuid` type, is in a file with a build constraint
   for that release, such as `//go:build go1.27`, and each `go.mod` declares
   `go 1.25`.
2. Open question 7: psql's layout of help, `HelpColumn(27)`, stays an
   option. It is not the default, because it would change the help of each
   program that does not ask for it.
3. Open question 9: a change from a person outside the project arrives as a
   GitHub discussion, and only so. A maintainer makes the change (D25).
4. Open question 14: the tests in `repo_test.go` do not run under TinyGo.
   They read files, and not the package.
5. The environment variable of a flag of a sub command does not take the
   path of the command. `--steps` of `app migrate` reads `APP_STEPS`, as
   item 6 of D43 says.

Lint, which answers open question 12:

6. `.golangci.yml` keeps `default: all` (D12), and gives the reason beside
   each linter that it disables, as in `github.com/xo/dbmeta`. It disables
   the linters that make idiomatic Go worse here, such as `goconst`,
   `paralleltest` (many tests change a `Default` variable), `dupword`, the
   complexity metrics, `dupl`, `godox` (a copy from the Go source tree has a
   TODO, by rule 2 of `AGENTS.md`) and `gomoddirectives` (the `replace` of
   the test module, D33).
7. `gosec` keeps its checks, except G304 (a program reads the files that it
   names), G301 and G306 (the docs of `metadoc` and the code of `oxdoc` are
   read by people and by the go command, so they are 0o644 in folders of
   0o755, as a man page or a source file is), and G401, G501 and G505 (the
   hash filters of interpolation are not for security). G115 is off only
   for `size.go`, where each conversion checks the range first. A test, a
   tool in `_tools/` and an example in `_examples/` are not the surface
   that gosec looks for. The six findings of `metadoc` were G301, G306 and
   G304, so they do not matter.
8. The real findings got a code change: the results of the writes in `has`
   and of a `Close` are discarded with `_ =`, as `AGENTS.md` asks; `Wrap`
   builds its text with a `strings.Builder`; a condition of `opIndex` in
   `interp.go` reads as what it means; two examples got a package comment;
   and a few unused names and conversions went.
9. A CI job, Lint, installs `golangci-lint` v2.15.0, the first release that
   reads the export data of Go 1.27, and lints the root, the tools and each
   example, then each module from its own folder. It also checks `gofmt`.
   Each module has no finding.

Completion:

10. Completion of the value of a flag offers the values of `HelpValue`,
    after the values of `Valid`, each with its usage. So `--help=<tab>`
    offers the topics. Before, item 2 of D36 and item 6 of D46 kept them out.
    B52 gave each value a meaning, so a maintainer asked for them.

The README:

11. The line "Fast" of the README says that `ox` is about as fast as cobra
    when a program starts, with less memory, and links
    `BenchmarkCommandLine` (D44).

Fuzz tests, in `fuzz_test.go`:

12. `FuzzCommandLine` parses, checks and runs any command line with a tree
    that uses most features of the parse: short and long flags, a count, a
    slice, a map, an optional value with a func, allowed values, a local
    flag, `Abbrev`, `NegateBools`, sub commands with aliases, typed and
    variadic arguments, `StopAtArg`, and help, version and completion. It
    must never panic. A first run of 13.6 million command lines found no
    fault.
13. `FuzzTypeValue` sets any text on a value of each built-in type. When
    `Set` takes the text, `Get` must not fail, and a new value must read the
    text of `Get` back to the same text. It found that `1e20` gave a size of
    the lowest `int64`, which help wrote as `--9223372036854775808 B`. The
    float paths of `asSize` and `asDuration`, and the seconds of
    `asDuration`, now refuse a value out of the range of an `int64`, and NaN.
    `bestSize` takes the size as a `uint64`, so the lowest size has one sign.
    `TestSizeOverflow` holds each case.
14. `FuzzTypeValue` does not wait for `Get` of a `big.Float` with a huge
    exponent, such as `8p97070970`. Its parse is fast, and `MarshalText` of
    `math/big` takes about a minute to write it. ox has no fault there.
15. `Base64T` and `HexT` decode their text, and `Get` gives the bytes, as
    `TestTypeNew` holds. So their text does not read back, and the fuzz test
    checks them only for a panic. Open question 15 asks if `Get` must give
    the encoded text.
16. `FuzzParseSize` holds that a size reads back from its exact form, with
    the verbs `q` and `Q` (D40), and `FuzzParsePortRange` holds that a port
    range reads back from its text. The inputs that failed are kept in
    `testdata/fuzz`, so `go test` runs them each time. Run a fuzz test with
    `go test -run '^$' -fuzz '^FuzzCommandLine$' -fuzztime 60s`. No CI job
    fuzzes.

### D50. The substring operator follows bash, and base64 and hex write their text. Decided. Supersedes D19.

A maintainer answered these on 2026-11-28.

The substring operator, which answers open question 10:

1. `${VAR:offset}` and `${VAR:offset:length}` follow bash, and the second
   number is a length, not an end index. So `${X:1:2}` of `aXbXc` is `Xb`,
   as in bash. Before, it was `X` (D19).
2. A negative offset counts from the end. An offset past either end gives an
   empty value, as in bash, and not an error. A length past the end stops at
   the end.
3. A negative length counts from the end of the value, so `${X:1:-1}` of
   `aXbXc` is `XbX`. An end before the offset is an error that wraps
   `ErrInvalidOp`, as bash fails with "substring expression < 0".
4. An empty offset or length is 0, as in bash, so `${X:1:}` is empty and
   `${X::2}` is `aX`.
5. The operator counts characters and not bytes, so `${X:1:3}` of `héllo`
   is `éll`, as bash does in a UTF-8 locale.
6. `${X:-1}` is still the default operator `:-`, as in bash. A negative
   offset needs a space, as in `${X: -1}`, as bash needs it.
   `TestSubstring` holds each case, and each was checked with bash.

Base64 and hex, which answers open question 15:

7. `Get` of a `Base64T` or a `HexT` value writes the encoded text that
   `Set` reads, so help and a config file write what the command line
   takes. Before, `Get` gave the decoded bytes, so help wrote a default of
   `aGVsbG8=` as `hello`. `otx.Bytes`, `otx.String` and `As[[]byte]` read
   the value, and still give the bytes. `FuzzTypeValue` now holds the round
   trip of both types (D49).

The confirmations:

8. D48 is decided as built, with each choice that the agent made, such as
   `ox.Extend`.
9. D34 is decided. `Value.WasSet` stays true only for a value from the
   command line, which answers open question 1. `Context.isSet` counts the
   environment and the config files for a required flag and a rule.
10. D28, D35 and D36 are decided as built, which answers open question 3.
11. D30 supersedes D20, because D30 holds the design as built. D20 is the
    sketch that came first.
12. Open questions 4, 5, 6, 11 and 13 were answered on 2026-11-27, when the
    work was built, and are marked answered.

### D51. A hook flag checks its value against Valid. Decided. Amends D46.

The `usql` session tried `v0.3.0` on 2026-11-28, and stayed on `v0.1.0` for
one reason. With the funcs of `HelpValue`, `--help=commands` works, but
`--help=bogus` writes the help and exits with 0. psql fails on a topic that
it does not know. `ox.Valid` on the help flag did nothing, because
`NewHelpFlag` makes it a hook flag, and the value of a hook was never
checked. A maintainer chose that a hook flag honors `Valid`, over an error
for each value that no `HelpValue` names.

1. The value of a hook flag checks its text against the validator of the
   flag, as each other value does, before the hook runs. So a program gives
   its help flag `ox.Valid("", "options", "commands", "variables")`, and
   `--help=bogus` fails with `ErrInvalidValue`, such as
   `--help: invalid value: "bogus": allowed values: "", "options",
   "commands", "variables"`. The empty value is in the list, so that a bare
   `--help` works.
2. A value of `HelpValue` with funcs runs its funcs before the check, so it
   needs no place in `Valid`. A program lists it anyway, so that the error
   names each topic.
3. A hook flag without `Valid` takes each value, as before, so no program
   changes unless it gives a hook flag `Valid`. Item 4 of D46 holds for such
   a flag.
4. The error quotes the value of a hook flag, as for a string, because the
   value is text from the command line. `Type.Formatter` of `HookT` is
   `%q`.
5. Completion leaves out an empty allowed value, such as the `""` of a bare
   `--help`, because it has nothing to complete.
   `TestHookValid` holds each case.

### D52. A short flag can take a plus form, as bash takes +o name. Decided. Amended by D53.

The `hail` session, a shell that copies bash's command line, asked on
2026-11-28 for bash's `+` options, with a maintainer's go-ahead for its
item W27. `bash +o name` and `+O name` turn a `set` or a `shopt` option
off, and `+e` and `+xv` turn single letters off. The order matters:
`-o errexit +o errexit` ends with errexit off. `ox` read each word that
starts with `+` as an argument. Two flagship models, GPT-6 Astra and
Gemini 3.1 Pro, critiqued two shapes: one flag for both forms, or a pair of
flags. Both chose one flag.

The maintainer chose these:

1. One flag takes both forms. The flag option `AllowPlus()` lets the short
   name of the flag take a `+`, and the struct tag option is `plus`.
2. A hook reads the sign with the method `Context.IsPlus()`, which is true
   while a hook runs for a plus form. No new hook signature is added.
   Hooks run in the order of the command line, so a program that needs the
   order of each `-` and `+` gives each option a hook.
3. A word that starts with `+` is read as flags only when its first letter
   names a flag with `AllowPlus`. Then each letter must name such a flag.
   Any other word, such as the `+%Y` of `date`, stays an argument, so a
   program without `AllowPlus` sees no change.
4. Help writes both forms in the short column, such as `-o, +o NAME` and
   `-e, +e`.
5. `OnSet`, a func that would run after each set of a flag that is not a
   hook, is left out until a program needs it.

An agent made these choices, and a maintainer accepted them on
2026-11-28:

6. A flag with `AllowPlus` needs a short name, and a count refuses it,
   because a count has no off. Each is an error from `NewFlag`.
7. The plus form of a plain bool flag sets it to false, as `--no-x` does
   (D47). It takes no value, so `+e=1` is an error that wraps
   `ErrInvalidValue`.
8. The plus form of any other flag takes its value as the `-` form does:
   `+o name`, `+oname` and `+o=name`. A cluster such as `+eo name` ends
   with the letter that takes the value. `+o` with no value is an error
   that wraps `ErrMissingArgument`. A `String` or a `Slice` flag keeps its
   value and loses the sign, so a program that needs the sign uses a hook.
9. A letter without `AllowPlus` in a cluster, such as the `c` of `+ec`, is
   an error that wraps `ErrUnknownFlag`, as bash fails on `bash +z`.
10. `ParseFlagPlus` parses a plus word, as `ParseFlagShort` parses a short
    word, and `IsPlus` is false again after it returns.
11. In the default layout, the short column grows from `-x, ` to
    `-x, +x, ` only when a flag of the list has a plus form, so the help of
    each other program does not change. A flag without a plus form then
    writes `-x,     ` so that each name lines up. The layout of
    `HelpColumn` writes the plus form after each short name.
12. Completion offers the plus forms when the word starts with `+` and the
    command has a flag with `AllowPlus`, with `Command.CompPlus`. It
    completes the value after `+o` and in `+o=`, as after `-o`.
13. `metadoc` writes the plus form beside the short name (D48).
14. The plus form is for the command line only. The environment and the
    config files set the flag as before, and there is no long form such as
    `++name`. `NegateBools`, `Abbrev` and `StopAtArg` work as before. A lone
    `+` is an argument, as a lone `-` is.
15. `Flag.AllowPlus` fits in the padding of `Flag`, which stays 496 bytes,
    and `BenchmarkCommandLine` shows no change (D44). `TestPlus`,
    `TestPlusHelp`, `TestPlusComp`, `TestPlusTag` and `TestPlusErrors` hold
    the cases, and `FuzzCommandLine` has flags with plus forms (D49).

### D53. A command can read plus words, letters and a value at the end as bash does. Decided. Amends D52.

The `hail` session moved to `v0.5.0` for its item W27 (D52), and found three
differences from bash on 2026-11-28 (B56, B57 and B58). A maintainer
accepted B57 and B58 as `hail` asked, and asked to revise D52 for B56. Each
of the three is opt in, so a program that does not ask reads its command
line as before. The names are the ones that `hail` proposed. A maintainer
accepted the design and each choice below on 2026-11-28.

B56, plus words:

1. The command option `PlusWords()` reads each word that starts with `+`
   as the plus forms of short flags, and not only a word whose first letter
   names a flag with `AllowPlus` (item 3 of D52). So `+z` and `+%Y` are
   errors that wrap `ErrUnknownFlag`, as `bash +z` fails with "invalid
   option", and a lone `+` is skipped, as bash skips it.
2. Without `PlusWords`, item 3 of D52 holds, so `date +%Y` keeps working.
3. A command reads plus words when it or a parent has `PlusWords`, so a
   sub command inherits it, as for `Abbrev` (D38). The first argument with
   `StopAtArg`, and `--`, still end the flags.

B57, letter clusters:

4. The command option `LetterClusters()` reads short letters as bash
   does. A letter that takes a value takes it only from the next word, so
   `-oe errexit` sets `e`, and `o` takes `errexit`. Without it, `ox` reads
   letters as `getopt` does, and `o` takes the value `e`.
5. Each letter that takes a value takes the next word in turn, so
   `-oO a b` gives `a` to `o` and `b` to `O`. The letters run in their order
   in the word.
6. An `=` is a letter, so `-e=` is an error that wraps `ErrUnknownFlag` and
   names `-=`, as bash refuses it. `-o=x` is refused too. A hook so never
   sees `-e=` as `-e`.
7. It applies to the `+` forms too, so `+oe errexit` turns `o` and `e` off.
   A long flag does not change. A sub command inherits it. Completion
   completes the value of the first letter that takes the next word.
8. `ParseFlagShort` and `ParseFlagPlus` share one parse, `parseLetters`,
   which reads either way.

B58, a value at the end:

9. The flag option `OptionalAtEnd()` lets a flag that takes a value be
   the last word of the command line without one, so a lone `-o` or `+o`
   reaches its hook, which lists the options as bash does. The flag is set
   from an empty value, and the hook reads the new method
   `Context.NoValue()` as true. Anywhere else, the flag takes its value as
   before.
10. It works for a short flag, its plus form, a letter at the end of a
    cluster with or without `LetterClusters`, and a long flag at the end,
    such as `--set-option`. A flag that takes no value refuses it with
    `ErrAppliedToInvalidType`.
11. No struct tag option is added for these three, because `hail` builds
    its flags by hand. One can be added when a program asks.
12. `TestBash` and `TestBashInherit` hold each case, `ExampleLetterClusters`
    shows them, and `FuzzCommandLine` has a sub command with the three
    (D49). The two fields of `Command` fit in its padding.
13. `warnDeprecated` takes the sign and the name of a flag apart, and joins
    them only when it writes a warning. Before, the parse of each flag built
    the name, deprecated or not. A run of `BenchmarkCommandLine` now takes
    307 allocations, where it took 313 (D44).

### D54. A command can read long flags as bash does. Decided.

The `hail` session moved to `v0.6.0` (D53), and found two more differences
from bash on 2026-11-28 (B59 and B60). bash reads a long option only before
the first short option, so `bash -e --noediting` fails with
`--: invalid option`. bash also refuses a value after `=` on a long option,
so `--norc=1` and `--rcfile=x` fail. `ox` took each. A maintainer chose one
command option for both, `BashLongFlags`, on 2026-11-28.

1. With `BashLongFlags()`, a long flag after the first short flag, or after
   the first plus form (D52), is an error that wraps `ErrUnknownFlag`, such
   as `--noediting: unknown flag: a long flag comes before the short flags`.
2. A long flag with a value after `=` is an error that wraps
   `ErrInvalidValue`. So a long flag takes its value only from the next
   word, as `--rcfile x`, and a long switch takes none.
3. `--` still ends the flags, after a short flag too.
4. A command reads its long flags so when it or a parent has
   `BashLongFlags`, so a sub command inherits it, as for `PlusWords` and
   `LetterClusters` (D53). A program without it reads long flags as before.
5. The check is in the parse of the command line, so completion and the
   environment do not change. `TestBashLongFlags` holds each case, and
   `FuzzCommandLine` has it on its sub command with bash's command line
   (D49).

## Open questions for the maintainers

An open question is not a decision. When a maintainer answers one, the
answer becomes a decision above, and the question goes. Each question has a
recommendation.

1. Answered on 2026-11-28: D34 is decided, and `Value.WasSet` stays true
   only for the command line (D50).
2. Answered on 2026-11-28: yes, D24 means what was built (D49).
3. Answered on 2026-11-28: D28, D35 and D36 are decided as built (D50).
4. Answered on 2026-11-27: `cmd/oxdoc` writes the doc comments into the
   program from `go generate` (D36).
5. Answered on 2026-11-27: the method is `Commands() []any`, and `Exec` can
   have any signature that `ox.Exec` takes (D36).
6. Answered on 2026-11-27: `ox.RunError` and `Command.UsageString` (D36).
7. Answered on 2026-11-28: `HelpColumn(27)` stays an option, and is not the
   default (D49).
8. Answered on 2026-11-27: a maintainer tagged `v0.1.0`, then `color`,
   `glob`, `hcl`, `toml` and `yaml`, in the order in [DESIGN.md](DESIGN.md).
9. Answered on 2026-11-28: a change arrives only as a GitHub discussion, and
   a maintainer makes it (D49).
10. Answered on 2026-11-28: the substring operator follows bash (D50).
11. Answered on 2026-11-27: the `hcl` module reads HCL files (D36).
12. Answered on 2026-11-28: CI runs `golangci-lint` on each module, and each
    module has no finding (D49).
13. Answered on 2026-11-27: CI runs `go vet` on each example, and
    `TestBuild` builds each style that the importer writes (D36).
14. Answered on 2026-11-28: the tests in `repo_test.go` do not run under
    TinyGo (D49).
15. Answered on 2026-11-28: `Get` of `Base64T` and `HexT` writes the encoded
    text (D50).
