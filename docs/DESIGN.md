# Design of the planned work

This document holds the design of the work that [PLAN.md](PLAN.md) plans and
[BACKLOG.md](BACKLOG.md) lists. Each section names the decision that it
follows. A part that no maintainer has decided yet says "proposed", and it is
also an open question in [PLAN.md](PLAN.md). Code in this document marks each
identifier that does not exist yet with the comment `// planned`.

The measurements were made on 2026-11-26, at `bb2641a` with the staged
changes, with Go 1.27.2, psql 18.7 and the `usql` main branch at `26b2b73`.

## The review of the API

On 2026-11-26 a maintainer asked that the API changes be designed, and
reviewed by the flagship models, before any code is written. A draft of
every area in this document went to Gemini 3.1 Pro and to GPT-6 Astra. Their
answers changed the draft in the places below, and each section that follows
holds the result. Where the two models disagreed, this document names the
choice and the reason.

1. Config: one `Config` interface for a decoded file replaces the per-key
   loaders for files. Both models agreed.
2. Config: a value from a file assigns the flag, and does not act as one
   use of the flag. GPT-6 Astra found that a count adds one for each use, so
   `verbose: 2` in a file would give 1. A slice from a file replaces the
   value, and does not append.
3. Config: a default is expanded only when it wins. Today every default is
   expanded before parsing, and a probe found a fault from it: a slice with
   the default `a,b` and the argument `--tag c` gave `[a b c]`. That was B44,
   and D34 fixed it.
4. Config: the config flag is resolved first, from the command line, then
   from the environment, then from its default, and its default cannot read
   a config file.
5. Config: a hook flag is filled from the environment or a file only when
   the program binds it by name, with `Key` or `EnvVar`. A hook can have side
   effects, so a file must not call it by accident.
6. Config: an unknown key in a file is not an error. Both models agreed. A
   strict mode can come later.
7. Config: the names. Gemini proposed `ConfigPaths`, `IsConfig`,
   `IgnoreEnv` and `IgnoreConfig`. GPT-6 Astra kept `ConfigFlag`, `NoEnv`
   and `NoConfig`, and proposed `ConfigFiles`. This document uses GPT-6
   Astra's names, because they match the names that `ox` has, such as
   `NoArg`.
8. Modules: Gemini said never to commit a `go.work`, and GPT-6 Astra said
   that one can be committed if CI also tests each module with
   `GOWORK=off`. This document does not commit one, so that a developer's own
   workspace wins, and CI makes one.
9. AppFrom: both models recommended the `go generate` tool over embedding
   the source, because the Go parser would bloat every program and can fail
   with TinyGo. Gemini said that struct tags are more robust than doc
   comments. A maintainer asked for doc comments, so a tag still wins over a
   doc comment.
10. Completion: Gemini proposed `Completer` for `CompFunc`. This document
    keeps `CompFunc`, because every completion name in `ox` starts with
    `Comp`.
11. Help: Gemini proposed a formatter func in place of `HelpColumn(n)`,
    because a fixed column is rigid. psql's layout is a fixed column, so
    `HelpColumn` stays, and a formatter func can come later.
12. Completion: Gemini said that `ValidArgs(min, max, values...)` mixes the
    count of the arguments with their values. That is true, and it is the
    API that `usql` and `dbtpl` use today, so it stays until a maintainer
    decides otherwise.

## Modules

D33 decides this section.

| Module | Path | Requires | Tag |
| --- | --- | --- | --- |
| `github.com/xo/ox` | `/` | the standard library only | `v0.1.0` |
| `github.com/xo/ox/color` | `color/` | `ox`, `github.com/kenshaw/colors` | `color/v0.1.0` |
| `github.com/xo/ox/glob` | `glob/` | `ox`, `github.com/kenshaw/glob` | `glob/v0.1.0` |
| `github.com/xo/ox/yaml` | `yaml/` | `ox`, `github.com/goccy/go-yaml` | `yaml/v0.1.0` |
| `github.com/xo/ox/toml` | `toml/` | `ox`, `github.com/pelletier/go-toml/v2` | `toml/v0.1.0` |
| `github.com/xo/ox/test` | `test/` | every module above, with `replace` directives | never |

`otx`, `strcase` and `text` stay in the root module, because they import
only the standard library and `ox`.

The rules:

1. The root module has no `require` block. A test checks it, as dbmeta and
   dburl do.
2. A sub module requires the release of the root that it was tested
   against. Minimal version selection can choose a newer root, so CI tests a
   sub module against the root in the same commit too.
3. `test/` holds the tests that need more than one module or a real
   program: config files of each format read end to end, the completion
   scripts run in bash, zsh and fish, the importer against captured help,
   and a build of each example. Its `replace` directives point at the
   folders in the repository, and nothing imports it.
4. No `go.work` is committed, and `.gitignore` ignores it. A developer makes
   one with `go work init . ./color ./glob ./yaml ./toml ./test`. CI makes
   the same one and runs the tests of every module in it. A second CI job
   runs each module with `GOWORK=off` against the versions in its `go.mod`.
5. To release, tag the root first. Then run
   `go get github.com/xo/ox@v0.1.0` in each sub module, stage that change,
   and tag each sub module after a maintainer commits it.
6. A program that upgrades the root does not upgrade a sub module, and the
   reverse. The README tells a program to upgrade them together:
   `go get github.com/xo/ox@latest github.com/xo/ox/yaml@latest`.

## Configuration

D34 decides this section.

Each flag gets its value from the first source that has one, in this order:

1. The command line.
2. The environment, when the root has `EnvPrefix` or the flag has
   `EnvVar`.
3. The config files, from the last file that has the key.
4. The default of the flag, expanded with interpolation. The default can
   read a key of a config file, such as `${yaml::server.port}`.

### The API

```go
// Config is one decoded config file.
type Config interface {
	// Lookup returns the value of a key. A key is a dotted path, such as
	// "server.port". A format can accept a syntax of its own as well, such as
	// a YAML path that starts with "$.".
	Lookup(key string) (any, bool, error)
}

// MapConfig is a Config over a decoded tree of maps. Lookup walks the tree
// one dotted part at a time.
type MapConfig map[string]any

// RegisterConfigFormat registers the decoder of a config file format, for
// the file extensions that follow. The root package registers JSONT with
// encoding/json. The yaml and toml modules register YAMLT and TOMLT.
func RegisterConfigFormat(typ ConfigType, decode func([]byte) (Config, error), exts ...string)

// ConfigFiles is a root Command option. It adds config files, in order. Each
// path is interpolated, such as "$APPCONFIG/config.yaml". A later file wins
// over an earlier one. A file that does not exist is skipped.
func ConfigFiles(paths ...string) CommandOption

// ConfigRequired is a root Command option. It makes it an error when no
// config file exists.
func ConfigRequired() CommandOption

// ConfigFlag is a Flag option. The value of the flag names the config file,
// and it replaces the files of ConfigFiles. Struct tag: `config`.
func ConfigFlag() FlagOption

// EnvPrefix is a root Command option. It fills each flag from the
// environment variable PREFIX_NAME, with the name in upper case and each "-"
// as "_".
func EnvPrefix(prefix string) CommandOption

// EnvVar is a Flag option. It fills the flag from the environment variable
// name, whatever the prefix. Struct tag: `env:NAME`.
func EnvVar(name string) FlagOption

// NoEnv and NoConfig are Flag options. They keep a flag out of the
// environment, or out of the config files. Struct tags: `noenv`, `noconfig`.
func NoEnv() FlagOption
func NoConfig() FlagOption

// Key, which exists, gives the config key of a flag for one format, or for
// every format when typ is empty. Without it, the key is the name of the flag.
func Key(typ, key string) FlagOption
```

`RegisterConfigLoader` stays for the keys that are not in a file, such as
`$HOME` and `${env::NAME}`. A registered format also answers its own keys in
interpolation: `${yaml::server.port}` reads only the YAML files, and
`${server.port}` reads every source in order. `UserConfigFile` goes, and
`ConfigFiles("$APPCONFIG/config.yaml")` does its job. JSON is in the root,
because `encoding/json` is in the standard library (D1).

### The rules

1. The key of a flag is the name of the flag, such as `log-level`, and it
   is the same for a flag of a sub command. `Key` gives another key, such as
   `server.port`. Flat keys survive the rename of a command.
2. A scalar in a file is formatted as a string and set. A list sets a slice
   or an array, and replaces its value. A map sets a map flag, one
   `key=value` at a time, in key order. A list for a flag that is not a
   slice, or a map for a flag that is not a map, is an error.
3. A number in a file assigns a count. A bool in a file assigns a switch.
   `NoArg` is the syntax of the command line, and it does not apply to a
   file.
4. An error from a file names the file and the key. A file that is missing
   is skipped. A file that cannot be read or decoded is an error, and so is
   a file with an extension that no format has registered.
5. An unknown key in a file is not an error.
6. `Value.WasSet` stays true only for a value from the command line.

### The lifecycle

`Context.Parse` becomes four steps:

1. Find the command, and set each flag that the command line gives. A hook
   runs at once, as today, so that `--help` and `--version` still exit at
   once.
2. Resolve the config flag: the command line, then the environment, then its
   default. Its default can read the environment, and not a config file.
3. Read the config files once.
4. Fill each flag that the command line did not set, for the executed
   command and each of its parents: from the environment, then the files,
   then the default. A default is expanded only when it wins.

`Populate` no longer runs before parsing, which fixed B44. Completion runs
steps 2 to 4 too, so that a completion func can read a value from a file,
and it ignores their errors.

## Go versions

D24 decides this section.

`ox` supports the newest Go release and the two releases before it. On
2026-11-26 that is Go 1.25, 1.26 and 1.27, so `go.mod` declares `go 1.25`.
When Go 1.28 comes out, `go.mod` moves to `go 1.26`.

The code uses the newest features that the minimum release has. A feature
that only a newer release has goes in a file with a build constraint for
that release, such as `//go:build go1.27`. The file registers the feature from
`init`, the same way a subpackage of D1 does. A program that builds with an
older Go does not get the feature, and everything else works.

A build with Go 1.25 and with Go 1.26 measured one dependency on Go 1.27: the
standard `uuid` package. It is in three places:

1. `type.go` registers `*uuid.UUID` as a text type in `init`, and names it
   in `typeRef`, in `reflectType` and in the variable `uuidType`.
2. `otx/otx.go` has `otx.UUID`.
3. `type_test.go` tests the type.

Each moves to a file with `//go:build go1.27`: `type_go127.go`,
`otx/otx_go127.go` and `type_go127_test.go`. CI then builds and tests with Go
1.25, 1.26 and 1.27, and with TinyGo 0.42.0. Rule 2 in `AGENTS.md` allows only
`tinygo` and `!tinygo` as build constraints today, so D24 amends it to allow a
Go release constraint too.

## Help that looks like psql

D29 decides this section.

`usql` moves to `ox` (D22), and `usql --help` on `ox` must look as close to
`psql --help` as `usql` allows. `ox` already writes psql's sections when a
program names them with `Help(Sections(...))` and gives each flag a
`Section(n)`. `_examples/gen/psql` is a copy of psql's options that the
importer made. A diff of its `--help` against psql 18.7 found these
differences. Each is an item in [BACKLOG.md](BACKLOG.md), under B32:

1. `ox` writes its own `--help` and `--version` in a "Flags:" section above
   the others. psql writes `-?, --help` and `-V, --version` inside "General
   options". A program can do the same today, and no document says how: it
   defines a flag with `Special("hook:help")` or `Special("hook:version")`,
   with its own short name, usage and section, and passes `Help()` and
   `Version()`. `ox` gives `--help` the short name `-h` unless the program
   defines `-h`, and psql's `-h` is the host.
2. `ox` starts the descriptions after the longest option, which is after 33
   characters for psql's options. psql starts them after 27 characters. An
   option that is too long for that, such as `-F, --field-separator=STRING`,
   goes on its own line, and its description goes on the next line, after 27
   characters.
3. `ox` joins a description that has a line break into one line. psql keeps
   the line break, and indents the next line by 27 characters.
4. `ox` writes no aliases in help. psql writes `-v, --set=, --variable=NAME=VALUE`.
5. `ox` has no way to show an optional value, as psql's `--help[=options]`,
   and no way to list the values of that option on lines of their own, such
   as `--help=commands`.
6. `ox` writes `(default: |)`. psql writes `(default: "|")` inside the
   description, by hand. `usql` writes its defaults by hand too, so `ox` needs
   a way to turn off the default text for one flag.
7. `ox` wraps a long description at `DefaultWrapWidth`, which is 95. psql
   never wraps one.
8. A flag of type `HookT` loses its placeholder in help. `Flag.SpecString`
   returns only the name for `HookT`, so `-v, --set NAME=VALUE` becomes
   `-v, --set`. `usql` builds 21 options this way.
9. psql writes `-1 ("one"), --single-transaction`. That text is an exception
   that only psql has, and `usql` does not copy it.

The importer has faults of its own in this diff, which B40 lists. It dropped
`-V`, `-1`, `-z`, `-0`, `--help=commands` and `--help=variables`, it read
`-v, --set=, --variable=` as a flag named `set` with the usage
`--variable=NAME=VALUE`, and it changed "home page:" to "home page -".

### What usql asked for

The `usql` session gave this layout on 2026-11-26. It is a proposal that a
maintainer of `usql` must confirm, and `ox` must make it possible:

1. "General options": `-c`, `-f`, `-v` (with `--set` and `--variable` shown
   as psql shows them), `-O`, `-U`, `-N`, `-D`, `--config`, `-V`, `-X`, `-1`
   and `-?`.
2. "Input and output options": `-o` and `-q`.
3. "Output format options": `-A`, `-C` and `--csv`, `-F`, `-G`, `-H`, `-J`,
   `-M`, `-P`, `-R`, `-t`, `-T`, `-x`, `-z` and `-0`.
4. "Connection options": `-w`, `-W`, and a line that says what a DSN is: a
   URL, or the name of a connection.
5. The usage line `usql [OPTION]... [DSN]`.
6. Text at the end, like psql's: type `\?` or `\help` inside `usql`, and a
   link to `https://github.com/xo/usql`.
7. The dump options stay hidden.

`usql` differs from psql on purpose: it has no `-d`, `-h`, `-p` or `-l`, `-U`
and `-N` mean other options, and `-h` must fail at once, so that a psql habit
such as `-h host` stops.

## A sample usql app

D29 and D22 decide this section.

`_examples/usql` will hold a program that has every option of `usql`, built
the way that `usql` will build it on `ox`. It does not connect to a database.
Its `Exec` prints the parsed options in order. A test compares its `--help`
with a golden file. The `usql` session said that the hand-built style fits,
because 21 options append to one ordered list, `-F` appends two entries, and
`-c` and `-f` append to a second list.

The structure, with the planned identifiers marked:

```go
// Command usql is a sample of the options of usql, built with ox.
package main

func main() {
	var args Args
	err := ox.RunError( // planned: a run that returns its error
		context.Background(),
		ox.Usage("usql", "the universal command-line interface for SQL databases"),
		ox.Banner("usql, the universal command-line interface for SQL databases"),
		ox.Spec("[OPTION]... [DSN]"),
		ox.Help(ox.Sections(
			"General options",
			"Input and output options",
			"Output format options",
			"Connection options",
		)),
		ox.HelpColumn(27), // planned: psql's column, B32
		ox.EnvPrefix("USQL"),
		ox.Footer(footer),
		ox.ValidArgs(0, 1),
		ox.VersionString(version),
		ox.Version(),
		ox.Flags().
			// help and version in section 0, as the usql prototype found
			Var("help", "show this help, then exit",
				ox.Special("hook:help"), ox.Short("?"), ox.Section(0)).
			Var("version", "output version information, then exit",
				ox.Special("hook:version"), ox.Short("V"), ox.Section(0)).
			// general options: -c and -f append to one list, in order
			Var("command", "run only single command (SQL or internal) and exit",
				ox.HookT, ox.Default(args.addCommand), ox.Short("c"), ox.Spec("COMMAND"), ox.Section(0)).
			Var("file", "execute commands from file, then exit",
				ox.HookT, ox.Default(args.addFile), ox.Short("f"), ox.Spec("FILE"), ox.Section(0)).
			// general options: -v, -O, -U and -N append to the options list
			Var("set", "set script variable NAME to VALUE (see \\set command)",
				ox.HookT, ox.Default(args.option("vars.")), ox.Short("v"), ox.Spec("NAME=VALUE"),
				ox.Aliases("variable"), ox.HiddenAliases("var"), ox.Section(0)). // planned: HiddenAliases, B25
			Bool("no-init", "do not execute initialization scripts",
				ox.Short("X"), ox.HiddenAliases("no-rc", "no-psqlrc", "no-usqlrc"), ox.Section(0)).
			// ...
			// output format options: each switch appends a fixed entry
			Var("no-align", "unaligned table output mode",
				ox.HookT, ox.NoArg(true, ""), ox.Default(args.fixed("print.format=unaligned")),
				ox.Short("A"), ox.Section(2)).
			// ...
			// the dump options: hidden, long only, and never from USQL_ variables
			String("target", "", ox.Hidden(true), ox.NoEnv(), ox.Spec("DRIVER")).
			// -h fails on purpose
			Var("bad-help", "", ox.HookT, ox.NoArg(true, ""), ox.Short("h"), ox.Hidden(true),
				ox.Default(func(string) error { return ox.ErrUnknownFlag })),
		ox.Exec(args.run),
	)
	if err != nil {
		// usql prints its own message and its own hint
	}
}
```

Each planned identifier is a proposal. B42 builds the sample when B25, B28
and B32 are done, and the `usql` session reviews it before `usql` moves.

## Path types

D26 decides this section.

| Type | Checks when the user gives a value | Completes |
| --- | --- | --- |
| `PathT` | nothing | files and directories |
| `FileT` | the file exists and is not a directory | files, and directories to go into |
| `DirT` | the directory exists | directories |

The value of each is a `string`, so `Bind(&s)` works with a `string` for all
three. Options change a path flag:

1. `Abs()` stores the absolute path.
2. `Real()` stores the path with each symbolic link resolved, with
   `Realpath`.
3. `Home()` shows a path under the home directory as `~/...` in help and in
   defaults. It does not change the stored value. Proposed: a maintainer must
   confirm that this is what the notes mean by "Homepath".
4. `Ext(".yaml", ".yml")` narrows file completion to those extensions, where
   the shell can do it.
5. `NotExist()` makes a path that must not exist yet, such as an output file.

Proposed: each path type expands a leading `~` to the home directory, and
checks only a value that the user gave, never a default. Questions 1 to 6 in
"Value completion" below list the rest.

## Value completion

D27 decides that `ox` completes flag values and positional arguments at
parity with cobra. D28 holds the API, which is proposed.

### What a user sees today

A probe program measured each case on 2026-10-15:

| Case | What the user sees | Priority |
| --- | --- | --- |
| Value after a flag (`--config `, `-c `, `-vc `) | Command names, with file completion off. The flag lookup keeps the dashes, so it never finds the flag | High |
| `--flag=` and `-f=` | `--config=` offers `--config` again. `--config=fo` offers nothing | High |
| Allowed values from `Valid(...)` | Nothing | High |
| Positional values from `ValidArgs` | Nothing | High |
| A completion func for a flag or for arguments | No API | High |
| Path types with file, directory and extension completion | `PathT` is a plain string | High |
| A bad earlier value, such as `--format bad sh` | No output, so completion stops | Medium |
| `--`, the end of the flags | Command names | Medium |
| Deprecated flags | Offered. Cobra does not offer them | Medium |
| Values after a separator, such as `--tag a,` | Nothing | Low |

### Parity with cobra

| cobra | ox |
| --- | --- |
| `ValidArgs` | `ValidArgs(min, max, values...)`, which also completes the values |
| `ValidArgsFunction` | `CompFunc(f)` on a command. Proposed |
| `RegisterFlagCompletionFunc` | `CompFunc(f)` on a flag. Proposed |
| `MarkFlagFilename(name, exts...)` | `FileT` with `Ext(exts...)` |
| `MarkFlagDirname` | `DirT` |
| `ShellCompDirective` values | `CompDirective`, with the same values |
| `CompletionWithDesc` | `Completion`, with `Name` and `Usage` |
| ActiveHelp lines | Not planned. B43 |

The scripts in `comp/` already handle every directive, so this work changes
only the Go side. Each shell measured on 2026-10-15:

1. bash and zsh filter by extension and list only directories.
2. fish and PowerShell ignore both filters and complete every file.
3. In every shell, file completion starts only when `ox` returns no
   completions and does not set `CompNoFileComp`. Any completion, even a
   correction, turns file completion off.

### The proposed API

```go
// CompFunc is a Command and Flag option. f returns the completions for the
// word under the cursor. ctx.Vars holds the flags parsed so far, and
// ctx.Args holds the positional arguments.
func CompFunc(f func(ctx *Context, toComplete string) ([]Completion, CompDirective, error)) CommandFlagOption // planned

// CompNoFilter is a directive that tells ox not to keep only the
// completions that start with toComplete.
const CompNoFilter CompDirective // planned
```

The questions, with the recommendation for each:

1. A path type checks only a value that the user gave, never a default.
2. The value of a path type is a `string`.
3. `Home()` changes only the display.
4. Each path type expands a leading `~`.
5. The option names are `Abs()`, `Real()`, `Home()`, `Ext(...)` and
   `NotExist()`.
6. `Ext` narrows completion and does not reject a value.
7. One `CompFunc` works on a command and on a flag.
8. `ox` keeps only the results that start with the typed text, for allowed
   values, `ValidArgs` values and `CompFunc` results, unless a func returns
   `CompNoFilter`. The scripts no longer filter by prefix.
9. A Levenshtein correction applies to allowed values and `ValidArgs`
   values, never to `CompFunc` results and never where files complete.
10. `ValidArgs` values are offered at each position up to the maximum, and
    stored on a new field, `Command.ValidArgs`.
11. The bash script stops its prefix filter for `--no-descriptions` too.
12. `usql` asks that completion offers hidden aliases. See B25.

## Apps from a struct

D30 decides this section.

`ox.AppFrom(v)` takes a pointer to a struct and builds the whole program
from it:

1. The doc comment of the struct type gives the usage, the banner, the spec,
   the examples and the footer, in the sections that `_examples/app` already
   sketches: `Usage:`, `Banner:`, `Examples:` and `Footer:`.
2. Each exported field with an `ox` tag is a flag, as `From` makes today. The
   doc comment of the field gives the usage, so that a description no longer
   needs a comma escaped in the tag. A tag that has a description still wins.
3. A struct that implements an interface for sub commands gives the sub
   commands. Each is a pointer to a struct, built the same way.
4. A struct that has an `Exec` method gives the exec func of its command.

```go
// App is the usql sample.
//
// Usage:
//
//	usql [OPTION]... [DSN]
type App struct {
	// Command runs only a single command (SQL or internal) and exits.
	Command []string `ox:",short:c,spec:COMMAND,section:0"`
	// Verbose turns on verbose output, such as -vvv.
	Verbose int `ox:",type:count,short:v"`
}

// Commands returns the sub commands. // proposed name
func (app *App) Commands() []any { return []any{new(DumpCommand)} }

// Exec runs the command.
func (app *App) Exec(ctx context.Context, args []string) error { return nil }

func main() {
	ox.AppFrom(new(App)) // planned
}
```

Go does not keep doc comments in a program, so reflection cannot read them.
Two designs can give them to `AppFrom`:

1. Proposed: `go generate` runs a tool from `ox`, such as
   `go run github.com/xo/ox/cmd/oxdoc`, which reads the source and writes a
   file that registers each doc comment from `init`. The program carries only
   the text. This works with TinyGo.
2. The program embeds its own source with `//go:embed *.go`, and `AppFrom`
   parses it with `go/parser` and `go/doc` when it starts. This needs no
   generate step, but every program then links the Go parser, and it is not
   known whether TinyGo builds `go/doc`.

`ParseCommand` and `RunApp` today look for `Sub() []any`. Proposed: rename
it to `Commands() []any`, because `Sub` is already the option that adds a
sub command. Both names are open questions.

## The importer

D31 decides this section.

The importer reads the command-line interface of a well-known program and
writes an `ox` program that has the same commands and flags. It replaces
`_examples/gen.go`, which does not build today (B14), and it moves to
`_tools/oximport`.

It reads a program in two ways:

1. A program built with cobra answers `<prog> __complete <args>`, as `gh` does.
   `gh __complete ''` lists each sub command with its description,
   `gh __complete pr list --` lists each flag with its description, and
   `gh __complete pr list --state ''` lists the allowed values `open`,
   `closed`, `merged` and `all`. This is more exact than the help text,
   because cobra writes it for a program.
2. Every program writes help. A parser for each format reads it: GNU
   `getopt_long` help (`du`, `cp`, `tar`), psql's help, the help of cobra
   (`--flag type`), and the help of Go's `flag` package (`-flag value`).

`__complete` gives no short names, types, defaults or sections, so the
importer reads both for a cobra program and joins them by the flag name.

It writes three styles of the same program, which D31 decides:

1. Built by hand: `ox.Run` with `ox.Flags()...`, as `_examples/gen` has.
2. Reflection: a struct with `ox` tags and `ox.From`.
3. App: a struct for `ox.AppFrom`, with doc comments.

Each run writes `testdata/import/<prog>/help.txt` with the text that it read,
so that a test can run the parsers without the program installed. A test
compares the `--help` of each program that it writes with the help that it
read, the same way that the psql diff above was made. That diff is the
measure of how well `ox` copies a program.

## Examples for coding agents

D32 decides this section.

An agent that writes an `ox` program needs short, correct examples that it
can copy. `docs/EXAMPLES.md` will hold them, by task, and each one is also an
`Example` func in an `example_test.go`, so that `go test` proves it. The
first set:

1. Each style: built by hand, bind, reflection, context and app.
2. Every tag option of `From`, with the escaping rules below.
3. Sub commands, aliases, hidden and deprecated flags, sections, and a
   psql-style help.
4. Each value type, including slices, maps, counts, hooks and path types.
5. Completion: allowed values, a `CompFunc`, and path types.
6. Testing a command with `Args(...)` and a buffer for `Stdout`.

### Escaping in a struct tag

These rules hold today, and a probe measured each one on 2026-11-26. A tag
is a Go string inside backquotes, so Go removes one level of backslashes
first, and then `ox` removes one more:

| To get this in the description | Write this in the tag | Why |
| --- | --- | --- |
| `read a file, then exit` | ``ox:"read a file\\, then exit"`` | A comma ends the description. `\\,` becomes `\,` for `ox`, which keeps the comma |
| `host:port to listen on` | ``ox:"host:port to listen on"`` | A colon in the description needs no escape. Only the options after it use `key:value` |
| `say "hi", loudly` | ``ox:"say \"hi\"\\, loudly"`` | A double quote ends the tag value, so Go needs `\"` |
| `a back\slash` | ``ox:"a back\\\\slash"`` | Each level removes one backslash |
| the default `http://x:80/a,b` | ``ox:"a url,default:http://x:80/a\\,b"`` | The first colon of an option splits the key, and the rest is the value |

A default that reads a YAML path escapes the `$` of the path, because a `$`
starts a variable: ``ox.Default(`${yaml::\$.store.book[*].author}`)``. A
dotted key needs no escape: `${yaml::server.port}`. A test in the `test`
module measured both on 2026-11-26.

Two aliases are `aliases:one|two`, and one is `alias:one`. A comma that is
not escaped is an error that names the field: `field X: unknown tag option:
" forgot to escape"`.
