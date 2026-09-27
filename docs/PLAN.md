# Plan

This file holds the plan for `ox`, then every decision that shapes it, then
the open questions for Ken. The index of the decisions comes first, so that a
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
| [D13](#d13-the-examples-include-command-trees-generated-from-the-help-of-real-tools-decided) | The examples include command trees generated from the help of real tools | Decided |
| [D14](#d14-a-default-value-is-expanded-with-bash-like-interpolation-decided) | A default value is expanded with bash-like interpolation | Decided |
| [D15](#d15-an-unknown-key-expands-to-nothing-and-a-failed-lookup-writes-its-error-into-the-value-decided) | An unknown key expands to nothing, and a failed lookup writes its error into the value | Decided |
| [D16](#d16-a-config-loader-is-registered-for-a-configtype-decided) | A config loader is registered for a ConfigType | Decided |
| [D17](#d17-the-go-directive-stays-at-a-release-that-the-pinned-tinygo-can-build-decided) | The go directive stays at a release that the pinned TinyGo can build | Decided |
| [D18](#d18-ox-makes-no-promise-of-backward-compatibility-decided) | ox makes no promise of backward compatibility | Decided |
| [D19](#d19-the-substring-operator-takes-an-end-index-and-not-a-length-proposed) | The substring operator takes an end index and not a length | Proposed |
| [D20](#d20-an-app-builds-its-command-tree-from-the-doc-comment-of-a-struct-proposed) | An app builds its command tree from the doc comment of a struct | Proposed |
| [D21](#d21-ox-is-set-up-for-coding-agents-as-every-xo-repository-is-decided) | ox is set up for coding agents as every xo repository is | Decided |

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

This section is at `eae81b9` (2026-09-14).

1. The root package imports the standard library and two of its own
   subpackages, `strcase` and `text`. See D1.
2. Flag parsing, the type system, help, version, completion and suggestions
   work, and the tests cover them. `go test ./...` and `tinygo test ./...` pass
   in CI.
3. Config file loading is not finished. The `yaml` loader returns no value
   for any key, and the `toml` package registers no loader. There is no `hcl`
   package, and the README names one. See B6 and B7 in
   [BACKLOG.md](BACKLOG.md).
4. The app style of D20 is a sketch. `UserConfigFile` does nothing. See B5.
5. The repository has no tag. `dbtpl` requires the pseudo-version for
   `2d66424` (2026-06-22), which is older than the move to Go 1.27 in D17.
6. `golangci-lint` reports 263 findings on main, and no CI job runs it. See
   D12 and B16.

## Direction

The direction comes from the README, the code and `notes.txt`, a file of
Ken's that git ignores. Ken has not written it down as a plan, so each item
is a direction and not a decision.

1. Finish config loading: the `yaml`, `toml` and `hcl` loaders, config files
   that a command reads from the user's directories, and the app style.
2. Finish the bash-like operators that the `InterpolateVar` comment lists.
   See B1 and B2.
3. Finish completion for flags. See B9.
4. Add the types that `notes.txt` lists. See B19.

## Decisions

Each decision that shapes `ox`, in the order that it was made. The entries are
append only. A decision is never edited to change its conclusion. When a later
decision changes one, the older entry keeps its text, and its heading names
the decision that changed it.

Each heading ends with a status. `Decided` means that the history shows that
Ken chose it. `Proposed` means that the code does it, and no record shows that
Ken chose it. A proposed decision is also an open question at the end of this
file. `Amends D3` and `Amended by D7` name the other decision, and both
decisions carry the link.

D1 to D20 were written on 2026-09-27 from the history of the repository, from
`61d64ed` (2024-11-19) to `eae81b9` (2026-09-14). Ken wrote every commit. No
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
0.40.1 in `1f5bd91`, then 0.42.0 in `6c99d89`.

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

### D13. The examples include command trees generated from the help of real tools. Decided.

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

### D17. The go directive stays at a release that the pinned TinyGo can build. Decided.

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

Ken said that no `xo` project makes a promise about its API. tblfmt D37
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

The code does it, and no record shows that Ken chose the difference from
bash. This is open question 1.

### D20. An app builds its command tree from the doc comment of a struct. Proposed.

Commits: `49a01e6` (2026-03-27), `2e2b6f4` (2026-03-30), `23b1eda`
(2026-03-30), `739fd0c` (2026-03-30).

`RunApp`, `RunAppContext` and `NewApp` take a pointer to a struct, and
`ParseCommand` is to build the command tree from its doc comment.
`_examples/app` shows the idea. The title of `23b1eda` is "Rough pencil in of
how app should work", `RunApp` and `RunAppContext` have empty doc comments,
and `UserConfigFile` does nothing. See B3, B4 and B5.

This is a sketch and not a finished decision. This is open question 2.

### D21. ox is set up for coding agents as every xo repository is. Decided.

Ken decided on 2026-09-27 that every repository in the `xo` namespace is set
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

## Open questions for Ken

An open question is not a decision. When Ken answers one, the answer becomes
a decision above, and the question goes.

1. The substring operator takes an end index, and bash takes a length. Is the
   difference deliberate, or does `ox` follow bash? See D19.
2. Is the app style of D20 the direction, and what must it do before it is
   finished? `RunApp` ignores its `context.Context`. See B3.
3. The README says that `ox` loads HCL, and imports `github.com/xo/ox/hcl` in
   its example. There is no such package. Does an `hcl` loader come, or does
   the README change? See B7.
4. The README lists `usql` as a program that uses `ox`. `usql` requires
   `cobra` and not `ox`. `dbtpl` requires `ox`. Is `usql` going to move to
   `ox`, or does the line go? See B17.
5. `ox` has no tag, and `dbtpl` requires a pseudo-version from before D17.
   Does `ox` get tagged releases?
6. No CI job runs `golangci-lint`, and main has 263 findings. Must CI run it,
   and must the disabled list follow the renamed linters? See B16.
7. `_examples/gen.go` does not build, and no CI job builds the examples. Must
   CI build them? See B14.
8. The repository tests in `repo_test.go` do not run under TinyGo, because a
   build constraint keeps them out. TinyGo 0.42.0 was not available here to
   find out whether they run there. Is that the right choice? See D2 and D21.
