# Backlog

The work in `ox` that is known and not done: the faults, the `TODO`
comments, and the features that the README or Ken's notes name and the code
does not have. A decision about any item goes in [PLAN.md](PLAN.md). When an
item is done, remove it here and name it in the commit message.

B1 to B21 were collected on 2026-09-27. Their line numbers are at `eae81b9`,
and they can move. Each item names where it came from:

1. Probe: a small program ran `ox` and showed the fault.
2. Read: the fault comes from reading the code, and no program ran it.
3. TODO: a `TODO` comment in the code.
4. Lint: a finding of `go vet` or `golangci-lint`.
5. README, or `notes.txt`: the text names something that the code does not
   do. `notes.txt` is a file of Ken's in the repository folder. Git ignores
   it, because `.gitignore` ignores `*.txt`.

The repository has no open issue and no pull request, open or closed, so no
item came from GitHub.

## Faults in interpolation

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

## Faults in the app style and the config loaders

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

### B5. `UserConfigFile` does nothing.

- Where: `opts.go:328-343`.
- Text: `dir = dir` and then `return nil`.
- Plain description: the option finds the user's config directory and then
  discards it. No file is loaded.
- Found: Lint: `go vet` reports "self-assignment of dir" at `opts.go:339`.

### B6. The `yaml` loader answers no key, and the `toml` package registers no loader.

- Where: `loadKey` in `yaml/yaml.go:20-38`, and `init` in `toml/toml.go:11-23`.
- Plain description: `loadKey` parses the key as a YAML path, then returns
  the empty string and false. The code that loads the file is in a comment.
  The body of `init` in `toml` is all in a comment, so the `decoder` type and
  its `Decode` method have no caller.
- Found: Lint: `go vet` reports "self-assignment of path" at
  `yaml/yaml.go:36`, and `unused` reports `decoder` in `toml/toml.go` and the
  `opts` and `once` fields in `yaml/yaml.go`.
- The README lists "Environment, YAML, TOML, HCL config loading" as a
  feature. `notes.txt` item 3 marks env as done and names the others.
  `TestInterpolate` reads `${yaml::...}` keys from a loader that the test
  registers itself.

### B7. There is no `hcl` package.

- Where: the README, twice, and `conf.go:22`, which declares `HCLT`.
- Plain description: the README says that `ox` loads HCL, and its import
  example has `_ "github.com/xo/ox/hcl"`, which does not build. `notes.txt`
  item 3 names `github.com/hashicorp/hcl/v2`, and also JSON and Jsonnet.
- Found: README. Open question 3 in [PLAN.md](PLAN.md) asks whether it comes.

### B8. `RegisterConfigLoader` ignores its `extensions` parameter.

- Where: `conf.go:162`.
- Plain description: the `yaml` package passes `"yaml", "yml"`, and nothing
  keeps them. They are the file extensions that a config file loader needs.
- Found: Read.

## Unfinished work named in the code

### B9. Completion does not complete the value of a flag.

- Where: `Context.Comps` in `ox.go:440-453`.
- Text: `// TODO: expose flags to allow hidden/deprecated`,
  `// TODO: logic incorrect; need to strip the -/-- and = from the flag`,
  and `// TODO: handle completion for flags with completion definitions`.
- Plain description: after a flag that takes a value, completion offers
  nothing. The lookup of the flag uses the text with its dashes. Hidden
  flags are never offered, and deprecated flags always are.
- Found: TODO. `notes.txt` item 7 names flag completion.

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

## Tooling

### B14. `_examples/gen.go` does not build, and CI builds no example.

- Where: `_examples/gen.go:392`.
- Text: `not enough arguments in call to c.Interpolate`
- Plain description: the signature of `Context.Interpolate` changed, and
  the generator did not follow. `go test ./...` skips a directory that starts
  with `_`, so CI does not build the examples. Every other example and every
  generated program in `_examples/gen/` passes `go vet`.
- Found: Lint: `go vet _examples/gen.go`. See D13, and open question 7 in
  [PLAN.md](PLAN.md).

### B15. `go vet ./...` fails on main.

- Plain description: it reports the two self-assignments of B5 and B6. When
  both are fixed, add `go vet ./...` to the commands that must pass in
  `AGENTS.md`.
- Found: Lint.

### B16. `golangci-lint` reports 263 findings on main, and no CI job runs it.

- Where: `.golangci.yml`.
- Plain description: the linter renamed `exhaustruct` to `exhaustruct_v5`
  and `wsl` to `wsl_v5`, so two disable lines no longer match. The two names
  give 98 of the findings. Most of the rest are style. These look like real
  faults or real questions:
  1. The findings of B1, B3, B5, B6 and B13.
  2. `recvcheck`: `Rate` in `size.go:193` and `FormattedTime` in
     `value.go:655` mix pointer and value receivers.
  3. `gosec` G115: `conv.go:560`, `587` and `614` convert a `uint64` to an
     `int64` with no bound.
  4. `errorlint`: `ox_test.go:186` uses a type assertion on an error.
  5. `wastedassign`: `size.go:80` and `strcase/initialisms.go:81`.
- Found: Lint, with golangci-lint 2.13.3. See D12, and open question 6 in
  [PLAN.md](PLAN.md).

## Documentation

### B17. The README is out of date.

- Where: `README.md`.
- Plain description:
  1. It says that `ox` "is built with Go 1.23+ applications in mind". The
     module declares Go 1.27. See D17.
  2. It lists `usql` as a program that uses `ox`. `usql` requires `cobra`
     and not `ox`. Open question 4 in [PLAN.md](PLAN.md) asks about it.
  3. It lists "Man page generation" as a design choice. No code writes a
     man page.
  4. It lists HCL loading. See B7.
- Found: README, compared with the code, the `usql` `go.mod` and the `dbtpl`
  `go.mod`.

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

## Planned features

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

## Other

### B20. Completion writes debug lines to standard error.

- Where: `ox.go:225`, `ox.go:420`, `ox.go:432`, `ox.go:434` and `ox.go:437`.
- Text: `COMP ENDED: %s`, `COMP CONTINUE ERR: %v`, `COMP ARGS: %v -- %s`,
  `COMP PARSE ERR: %v` and `COMP COMMAND: %s`.
- Plain description: every completion request writes these lines. The
  scripts in `comp/` discard standard error, so a shell does not show them.
  A person who runs `__complete` by hand sees them.
- Found: Read. It is not known whether they stay on purpose.

### B21. `.gitignore` ignores every `.txt`, `.log` and shell script file.

- Where: `.gitignore:40-41` and `.gitignore:43-46`.
- Plain description: the patterns `*.txt`, `*.log`, `*.bash`, `*.fish`,
  `*.ps1` and `*.zsh` keep `notes.txt`, `_examples/unknown-flag-types.log`
  and the scripts that `_examples/tree/gen.sh` writes out of git. They also
  match `LICENSE.completions.txt` and the eight files in `comp/`, which are
  committed. Git goes on tracking a file that it already has. A new file in
  `comp/` needs `git add -f`.
- Found: Read.
