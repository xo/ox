# Backlog

The work in `ox` that is known and not done, ordered by importance. A
decision about any item goes in [PLAN.md](PLAN.md), and the design of the
planned work is in [DESIGN.md](DESIGN.md). When an item is done, remove it
here and name it in the commit message.

The number of an item never changes, so the order of the numbers is not the
order of importance. B1 to B21 were collected on 2026-09-27, B22 to B29 on
2026-11-09, and B30 to B45 on 2026-11-26. A line number is at `eae81b9` for
B1 to B21, and at `bb2641a` after that. Each item names where it came from:

1. Probe: a small program ran `ox` and showed the fault.
2. Read: the fault comes from reading the code, and no program ran it.
3. TODO: a `TODO` comment in the code.
4. Lint: a finding of `go vet` or `golangci-lint`.
5. README, or `notes.txt`: the text names something that the code does not
   do. `notes.txt` is a maintainer's file in the repository folder. Git
   ignores it, because `.gitignore` ignores `*.txt`.
6. A session: another session asked for it, such as the `usql` session.

Issues and pull requests are off on GitHub (D25), so no item comes from
there.

Eleven items closed on 2026-11-26. The split into modules of D33 and the
config loading of D34 closed B5, B6, B8, B23, B44 and B45. Every module
passes `go vet`, which closed B15. B29 closed, because a maintainer asked
that config loading work properly (D34). B31 replaces B9. B23 holds what B24
asked, because `usql` uses its only annotation to keep an option out of the
environment. B27 closed, because the `usql` session confirmed that a flag of
type `HookT` and `NoArg` cover its custom value types.

## Priority 1: needed before usql moves, or wrong today

The `usql` session said on 2026-11-26 that a maintainer decided that `usql`
moves from `cobra` to `ox` now (D22). The `usql` session built a prototype with every option of `usql` on
`bb2641a` and the staged changes, and sent its list in this order. It writes
its side against a local checkout. Tell the session named `usql` when each
item lands, with its API names.

The prototype found that these work today, with no change to `ox`:

1. A `HookT` flag with a `func(string)` default appends in the order of the
   command line. A probe on 2026-11-09 gave `[set=A=1 pset=x=2 set=B=3]`
   for `-v A=1 -P x=2 --set B=3`, and `-F` can append two entries.
2. Combined short options, such as `-Aq` and `-zX`, and `-vA=1` and `-F=,`.
3. `NoArg(true, "")` switches, and `NoArg(true, true)` for an optional
   bool. `--no-init=maybe` fails.
4. Aliases, `--`, and `ValidArgs(0, 1)`, which refuses a second argument.
5. A hidden flag with `Short("h")` and a hook that fails, so that psql's
   `-h` stops at once.
6. Help sections, `Spec`, `Banner` and `Footer`.
7. Help and version flags that the program defines with
   `Special("hook:help")` or `Special("hook:version")`, its own short name,
   usage and section, together with the `Help()` and `Version()` options. A
   probe found that `-V` fails with "missing argument" without `Version()`.
8. A root command with no sub commands gets the hidden flags
   `--completion-script-bash`, `--completion-script-zsh`,
   `--completion-script-fish`, `--completion-script-powershell` and
   `--no-descriptions` from `Comp()`, and no `completion` sub command. So a
   positional argument named `completion` stays a DSN. These are the names
   that `usql` uses with `cobra`.
9. The usage as text, and a run that returns its error. See B28.

### B33. A flag of type `HookT` loses its placeholder in help.

- Where: `Flag.SpecString` in `cmd.go:811-823`.
- Plain description: `SpecString` returns only the name when the type is
  `HookT`, so `Spec("NAME=VALUE")` is not shown, and `-v, --set NAME=VALUE`
  is written as `-v, --set`. `usql` builds 21 options as hooks. Before
  2026-11-26 this backlog said that `Spec` gives each hook its placeholder.
  That holds for parsing, and not for help.
- Found: Read, on 2026-11-26.

### B32. Help cannot look like psql's help.

- Plain description: a diff of `_examples/gen/psql --help` against psql 18.7
  found these differences, which D29 asks `ox` to remove:
  1. The built-in `--help` and `--version` go in a "Flags:" section of their
     own. A program can move them today, and no document says how: define
     a flag with `Special("hook:help")` or `Special("hook:version")`, its
     own short name, usage and `Section`, and pass `Help()` and `Version()`.
     The `usql` prototype and a probe both did it on 2026-11-26. `ox` gives
     `--help` the short name `-h` unless the program defines `-h` itself,
     and psql's `-h` is the host.
  2. The descriptions start after the longest option, after 33 characters for
     psql's options. psql starts them after 27, and puts a longer option on
     its own line.
  3. A line break in a description is lost. psql keeps it.
  4. Aliases are not written. See B25.
  5. An optional value, such as `--help[=options]`, cannot be shown, and
     neither can the values of that option, such as `--help=commands`.
  6. The default text cannot be turned off for one flag. psql and `usql`
     write their defaults by hand.
  7. `ox` wraps a long description at `DefaultWrapWidth`, which is 95. psql
     never wraps a description. It keeps each one on one line, or breaks it
     where the usage text has a line break.
  [DESIGN.md](DESIGN.md) has the layout that the `usql` session asked for.
- Found: Probe, and the prototype of the `usql` session, on 2026-11-26.

### B25. Completion does not offer hidden options.

- Plain description: `Context.Comps` passes `hidden = false`, so a hidden
  flag is never completed. `usql` hides a few options, such as `--no-psqlrc`
  and `--variable`, from help and wants them completed. Cobra does not
  complete a hidden flag, so `usql` un-hides them while it writes a
  completion script. `usql` also needs the value completion of B31 and the
  end of the debug lines of B20.
- On 2026-11-26 the `usql` session made this exact. While it completes,
  `usql` shows the hidden aliases `--no-psqlrc`, `--no-usqlrc`, `--var` and
  `--variable`, and keeps `--no-rc` hidden. In help, `-v` shows `--variable`
  and hides `--var`. So `ox` needs an alias that help hides and completion
  offers, such as `HiddenAliases(...)`, and help that writes the aliases
  that are not hidden, as psql does. `ox` writes no aliases in help today.
- Found: the `usql` session, item 4. Read: `ox.go:441`.

### B20. Completion writes debug lines to standard error.

- Where: `ox.go:225`, `ox.go:420`, `ox.go:432`, `ox.go:434` and `ox.go:437`.
- Text: `COMP ENDED: %s`, `COMP CONTINUE ERR: %v`, `COMP ARGS: %v -- %s`,
  `COMP PARSE ERR: %v` and `COMP COMMAND: %s`.
- Plain description: every completion request writes these lines. The
  scripts in `comp/` discard standard error, so a shell does not show them.
  A person who runs `__complete` by hand sees them.
- Found: Read. It is not known whether they stay on purpose.

### B30. ox requires Go 1.27, and D24 makes the minimum Go 1.25.

- Plain description: `go.mod` declares `go 1.27`. A build with Go 1.25 and
  with Go 1.26 fails on one import, the standard `uuid` package, in
  `type.go`, `otx/otx.go` and `type_test.go`. Move each use to a file with
  `//go:build go1.27`, declare `go 1.25`, and add Go 1.25, 1.26 and 1.27 to
  the CI job beside TinyGo 0.42.0. Change rule 2 in `AGENTS.md` to allow a
  Go release constraint. [DESIGN.md](DESIGN.md) has the details.
- Found: a build with `GOTOOLCHAIN=go1.25.0` and `go1.26.0` in a copy of the
  repository, on 2026-11-26.

### B22. ox has no tagged release.

- Plain description: `usql` must require a release, not a pseudo-version.
  `go vet ./...` passed in every module on 2026-11-26. D33 adds a tag for
  each module, in the order that [DESIGN.md](DESIGN.md) gives. Open
  question 8 in [PLAN.md](PLAN.md) asks when the first tag is.
- Found: the `usql` session, item 1.

### B31. Completion does not complete flag values or positional arguments.

- Where: `Context.Comps` in `ox.go:440-453`.
- Text: `// TODO: expose flags to allow hidden/deprecated`,
  `// TODO: logic incorrect; need to strip the -/-- and = from the flag`,
  and `// TODO: handle completion for flags with completion definitions`.
- Plain description: a probe program measured each case on 2026-10-15.
  These are wrong or missing today:
  1. The value after a flag, as in `--config `, `-c ` or `-vc `, offers
     command names with file completion off. The flag lookup keeps the
     dashes, so it never finds the flag.
  2. `--config=` offers `--config` again, and `--config=fo` offers nothing.
     The scripts add the `--config=` prefix back themselves, so `ox` must
     not add it.
  3. The allowed values of `Valid(...)` are not offered.
  4. The values of `ValidArgs` are not offered.
  5. A path type gives no file or directory completion. See B35.
  The API for a completion func is B37. D27 decides the parity with cobra,
  and [DESIGN.md](DESIGN.md) has the measurements for each shell.
- Found: TODO, and Probe. This replaces B9.

### B34. A parse error must be easy to tell apart.

- Plain description: `usql` prints a hint after a parse error, and finds one
  with `errors.Is`. Its prototype found on 2026-11-26 that the parse errors
  wrap `ErrUnknownFlag`, `ErrMissingArgument`, `ErrInvalidArgCount` and
  `ErrInvalidValue`. Keep these stable, add a test that holds each one, and
  list them in the documentation of `Parse`. D18 allows a rename, and a
  rename of one of these breaks `usql`.
- Found: the `usql` session, on 2026-11-26.

### B26. The `Hook` option does not work.

- Where: `Hook` at `opts.go:694`, and `newHook` at `value.go:579`.
- Plain description: `Hook` stores a `func(context.Context) error`, and
  `newHook` has no case for that type. A flag built with `Hook` fails with
  `invalid value: invalid hook func func(context.Context) error`. The flag
  also takes an argument, so a bare `--flag` fails with `missing argument`.
  A flag of type `HookT` with a `func(string)` default works.
- Found: Probe, on 2026-11-09.

## Priority 2: the planned direction

### B28. Run does not return its error, and there is no usage string.

- Plain description:
  1. `Run` and `RunContext` handle the error themselves and exit. `usql`
     needs a run func that returns the error, so that it chooses the exit
     status and the message. `NewContext`, `Context.Parse`,
     `Context.Validate` and `Context.Run` each return an error, so a caller
     can call them in turn today, but no one func does it.
  2. `usql` prints the usage inside `\? options` and needs a `Usage()
     string` helper. `CommandHelp.WriteTo` writes the help to an
     `io.Writer`, so a `strings.Builder` works today.
  3. Help has no section that describes the positional arguments, such as
     `[DSN]`. `Spec` changes only the spec line. On 2026-11-26 the `usql`
     session said that it no longer needs this, because its "Connection
     options" section can say what a DSN is. It stays here for other
     programs, at a lower priority.
- On 2026-11-26 the `usql` prototype did both without a change to `ox`: it
  calls `NewContext`, its own `Context.Handler`, `Parse`, `Validate` and
  `Run` in turn, and it writes `cmd.HelpContext(ctx)` to a
  `strings.Builder`. A run func and a usage string are easier to find, and
  they no longer block `usql`.
- Found: the `usql` session, item 6.

### B35. The path types of D26 do not exist.

- Plain description: add `FileT` and `DirT` beside `PathT`, and the options
  `Abs`, `Real`, `Home`, `Ext` and `NotExist`. [DESIGN.md](DESIGN.md) has the
  table, and questions 1 to 6 there are still open.
- Found: D26.

### B37. There is no completion func for a flag or for arguments.

- Plain description: cobra has `ValidArgsFunction` and
  `RegisterFlagCompletionFunc`. D28 proposes one option, `CompFunc`, for
  both, and a directive, `CompNoFilter`. [DESIGN.md](DESIGN.md) has the
  signature and the questions.
- Found: D27.

### B38. Completion fails after a bad value, and after `--`.

- Plain description: three cases that a probe found on 2026-10-15:
  1. After a bad value, as in `--format bad sh`, completion prints nothing,
     because `Comps` returns the parse error.
  2. After `--`, completion offers command names.
  3. A deprecated flag is offered. Cobra does not offer one.
- Found: Probe.

### B36. `AppFrom` does not exist.

- Plain description: D30 plans `ox.AppFrom(*T)`, which reads the usage, the
  flags and the sub commands from a struct and its doc comments. It replaces
  the sketch of D20: `RunApp`, `RunAppContext`, `NewApp` and
  `ParseCommand`, with the faults of B3 and B4. How `AppFrom` gets the doc
  comments is an open question in [PLAN.md](PLAN.md).
- Found: D30.

### B39. The importer must replace `_examples/gen.go`.

- Plain description: D31 plans `_tools/oximport`. It reads a cobra program
  through `__complete` and its help, and every other program through its
  help, and it writes the hand-built, reflection and app styles. It keeps
  the help that it read in `testdata/import/`, so that its tests run without
  the programs. [DESIGN.md](DESIGN.md) has the design. B14 holds the
  faults of the old generator.
- Found: D31.

### B40. The old generator misreads psql's help.

- Plain description: the program that `_examples/gen.go` wrote for psql
  lost `-V`, `-1`, `-z`, `-0`, `--help=commands` and `--help=variables`. It
  read `-v, --set=, --variable=NAME=VALUE` as a flag `set` with the usage
  `--variable=NAME=VALUE`, and it changed "home page:" in the footer to
  "home page -". The importer of B39 must read each of these right. Each
  lost option is one whose description goes on the next line, or whose
  short name is followed by text, such as `-1 ("one")`.
- Found: Probe, a diff against psql 18.7 on 2026-11-26.

### B14. `_examples/gen.go` does not build, and CI builds no example.

- Where: `_examples/gen.go:392`.
- Text: `not enough arguments in call to c.Interpolate`
- Plain description: the signature of `Context.Interpolate` changed, and
  the generator did not follow. `go test ./...` skips a directory that starts
  with `_`, so CI does not build the examples. Every other example and every
  generated program in `_examples/gen/` passes `go vet`.
- Found: Lint: `go vet _examples/gen.go`. See D13, and open question 13 in
  [PLAN.md](PLAN.md).

### B41. There are few examples for a coding agent.

- Plain description: D32 plans `docs/EXAMPLES.md`, with an `Example` func
  in a test for each example, so that `go test` proves it. The first set,
  and the escaping rules of a struct tag, are in [DESIGN.md](DESIGN.md).
  Today the escaping of a comma in a tag is in one test only, and no
  document has it.
- Found: a maintainer, on 2026-11-26.

### B42. There is no sample of usql built with ox.

- Plain description: `_examples/usql` will hold every option of `usql`,
  built as `usql` will build it on `ox`, with a test that compares its help
  with a golden file. It needs B25, B28, B32 and B33 first. The
  `usql` session reviews it before `usql` moves. [DESIGN.md](DESIGN.md) has
  its structure and the list of options from `usql`.
- Found: a maintainer, on 2026-11-26.

### B16. `golangci-lint` reports 263 findings on main, and no CI job runs it.

- Where: `.golangci.yml`.
- Plain description: the linter renamed `exhaustruct` to `exhaustruct_v5`
  and `wsl` to `wsl_v5`, so two disable lines no longer match. The two names
  give 98 of the findings. Most of the rest are style. These look like real
  faults or real questions:
  1. The findings of B1, B3 and B13.
  2. `recvcheck`: `Rate` in `size.go:193` and `FormattedTime` in
     `value.go:655` mix pointer and value receivers.
  3. `gosec` G115: `conv.go:560`, `587` and `614` convert a `uint64` to an
     `int64` with no bound.
  4. `errorlint`: `ox_test.go:186` uses a type assertion on an error.
  5. `wastedassign`: `size.go:80` and `strcase/initialisms.go:81`.
- Found: Lint, with golangci-lint 2.13.3. See D12, and open question 12 in
  [PLAN.md](PLAN.md).

### B1. The operators `##` and `%%` return the value unchanged.

- Where: `apply` in `interp.go:349-352`.
- Plain description: `DefaultOps` has no `##` and no `%%`. For a two
  character operator that it does not have, `apply` sets `f`, `op` and `str`
  from the first character, and then leaves the `switch` without a call to
  `f`. So the value passes through unchanged, and no error is returned.
- Found: Probe. With `X` set to `aXbXc`, `${X##*X}` and `${X%%X*}` both give
  `aXbXc`. Lint: `staticcheck` SA4006 reports that `f` and `op` are never
  used.
- The comment on `InterpolateVar` at `interp.go:229` lists `##` and `%%` as
  built in. `notes.txt` item 1 describes them as the bash operators, which
  remove the longest match of a pattern.
- Also: `#` and `%` remove literal text with `strings.TrimPrefix` and
  `strings.TrimSuffix`. In bash they take a pattern. `${X#aX}` gives `bXc`.

### B2. The operators `^`, `^^`, `,` and `,,` are not implemented.

- Where: `DefaultOps` in `interp.go`, and `isOp` at `interp.go:476`, which
  does not know `,`.
- Plain description: `${X^^}` gives the value
  `!(ERROR: ${X^^}: not implemented: invalid op "^^")`, and `${X,,}` gives the
  empty string with no error.
- Found: Probe. The comment on `InterpolateVar` lists `^` and `^^` as built
  in, and `notes.txt` item 1 lists all four as the bash case operators. The
  `|upper` and `|lower` filters already do the whole string.

## Priority 3: config loading and smaller faults

### B7. There is no `hcl` package.

- Where: the README, twice, and `conf.go:22`, which declares `HCLT`.
- Plain description: the README says that `ox` loads HCL, and its import
  example has `_ "github.com/xo/ox/hcl"`, which does not build. `notes.txt`
  item 3 names `github.com/hashicorp/hcl/v2`, and also JSON and Jsonnet.
- Found: README. Open question 11 in [PLAN.md](PLAN.md) asks whether it comes.

### B3. `RunApp` ignores its `context.Context`.

- Where: `ox.go:312-314`.
- Plain description: `RunApp(ctx, v, opts...)` calls
  `RunAppContext(DefaultContext, v, opts...)`, so the caller's `ctx` never
  reaches the command. Both funcs have an empty doc comment.
- Found: Read. Lint: `contextcheck` reports `ox.go:313`.
- See D20. `RunApp` is part of a sketch, so its signature is open too.

### B4. `RunAppContext` calls a method on a nil `*Context`.

- Where: `ox.go:304-306`, with `NewApp` at `ox.go:521-525`.
- Plain description: if `NewContext` fails, `NewApp` returns a nil
  `*Context`, and `RunAppContext` then calls `c.Handler(err)` on it. The
  program panics in place of the error.
- Found: Read.

### B10. Case sensitivity and the distance match cannot be turned off.

- Where: `cmd.go:288`.
- Text: `// TODO: settings for toggling case sensitivity / disabling ldist matching`
- Found: TODO. See D11.

### B11. The help output does not show argument validation or completion.

- Where: `defs.go:374`.
- Text: `// TODO: better output when arg validation/completion has been set`
- Found: TODO.

### B12. A conversion to `[]byte` or `[]rune` is not implemented.

- Where: `conv.go:687`.
- Text: `// TODO: implement []byte/[]rune`
- Found: TODO.

### B13. Two stand-ins wait for TinyGo to support Go 1.26.

- Where: `fields` at `cmd.go:1166-1179`, and `asExitError` at
  `ox.go:728-736`.
- Text: `// TODO: remove this after tinygo supports go1.26+` and
  `// TODO: remove when tinygo supports go1.26.`
- Plain description: CI now runs TinyGo 0.42.0 and the module declares Go
  1.27. Find out whether TinyGo 0.42.0 has `reflect.Type.Fields` and
  `errors.AsType`, and if it has, use them. The local TinyGo is 0.40.1, which
  refuses Go 1.27, so this was not measured.
- Found: TODO. Lint: `modernize` reports the `NumField` loop at
  `cmd.go:1173`.

### B17. The README is out of date.

- Where: `README.md`.
- Plain description:
  1. It says that `ox` "is built with Go 1.23+ applications in mind". The
     module declares Go 1.27. See D17.
  2. It lists "Man page generation" as a design choice. No code writes a
     man page.
  3. It lists HCL loading. See B7.
- Found: README, compared with the code.

### B18. Some doc comments name the wrong thing.

- Plain description:
  1. `conf.go:63` says that `$STATE` comes from `os.UserStateDir`. It comes
     from `UserStateDir` in `misc.go:14`, which is part of `ox`.
  2. `ox.go:118` and `ox.go:333` link `[os.ExitError]`. The type is
     `exec.ExitError`, from `os/exec`.
  3. `ox.go:708` links `[DefaultVersionString]`, which does not exist.
  4. `ox.go:57` names `DefaultFlagNameMaper`, and the variable is
     `DefaultFlagNameMapper`.
  5. `RunApp` and `RunAppContext` have a doc comment that holds only the
     name. See B3.
- Found: Read.

### B19. The types and features in `notes.txt`.

- Plain description: `notes.txt` lists these, and the code does not have
  them:
  1. Cascading config levels for interpolation.
  2. A `$ANY{}` key that returns the first value from any loader.
  3. Interpolation of `NoArgDef` in `Parse`.
  4. Reflection faults with `time.Duration`.
  5. Global flags that turn propagation and display on or off.
  6. An error when a flag is defined twice.
  7. Any type whose pointer has `MarshalText` and `UnmarshalText`, or the
     binary pair.
  8. `--help` and `help <command path>` that work with completion.
  9. `Dir`, `File`, `Realdir` and `Realfile` types, with `Homepath` and
     `Abspath` forms.
  10. A strftime type, a semver type, and a port range type such as
      `100:200`.
- Found: `notes.txt`. Some items can be partly done. Measure each one
  before you start it.

### B21. `.gitignore` ignores every `.txt`, `.log` and shell script file.

- Where: `.gitignore:40-41` and `.gitignore:43-46`.
- Plain description: the patterns `*.txt`, `*.log`, `*.bash`, `*.fish`,
  `*.ps1` and `*.zsh` keep `notes.txt`, `_examples/unknown-flag-types.log`
  and the scripts that `_examples/tree/gen.sh` writes out of git. They also
  match `LICENSE.completions.txt` and the eight files in `comp/`, which are
  committed. Git goes on tracking a file that it already has. A new file in
  `comp/` needs `git add -f`.
- Found: Read.

### B43. ox writes no ActiveHelp.

- Plain description: the scripts of cobra show ActiveHelp, which is a line
  of help that a program prints during completion. `ox` gives the scripts
  the name of the ActiveHelp variable and never writes a line of it. No
  consumer has asked for it.
- Found: Read, on 2026-11-26.
