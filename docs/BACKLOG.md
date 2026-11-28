# Backlog

The work in `ox` that is known and not done, ordered by importance. A
decision about any item goes in [PLAN.md](PLAN.md), and the design of the
planned work is in [DESIGN.md](DESIGN.md). When an item is done, remove it
here and name it in the commit message.

The number of an item never changes, so the order of the numbers is not the
order of importance. B1 to B21 were collected on 2026-09-27, B22 to B29 on
2026-11-09, B30 to B51 on 2026-11-26, B52 on 2026-11-27, and B53 to B55 on
2026-11-28. A line number is at `eae81b9` for
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

On 2026-11-27 one agent took the open items in order, and closed B32, B28,
B35, B36, B14, B42, B16, B1, B2, B7, B3, B4, B10, B11, B12, B17, B18, B21,
B43, B46, B47, B48, B49, B50 and B51. D36 lists the choices that it made.
A maintainer tagged `v0.1.0` the same day, which closed B22. B13 waits
for the `go` line of D24. B55 closed on 2026-11-28 (D44), and B19, B52, B53
and B54 closed the same day (D45, D46, D47 and D48).

## Priority 1: needed before usql moves, or wrong today

The `usql` session said on 2026-11-26 that a maintainer decided that `usql`
moves from `cobra` to `ox` now (D22). Later that day it said that a
maintainer decided two more things: `usql` drops its `USQL_<OPTION>`
variables, so it uses none of `EnvPrefix`, `NoEnv` and `EnvVar`, and `usql`
commits the move only against a tag of `ox`. What `usql` needs before that
tag is, in order: B33, B32, B25, and the tag (B22). Agents finished B33,
B32 for psql's layout, B25, B20, B31, B37, B38, B30, B34, B26, B39, B40 and B41
the same day, and the work was merged in the main checkout. The backlog agent
finished the rest of B32 on 2026-11-27 (D36). A maintainer tagged
`v0.1.0` the same day, so no item is left for `usql`. The `usql` session built a prototype with every option of `usql` on
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
9. The usage as text, and a run that returns its error. Since 2026-11-27,
   `Command.UsageString` and `RunError` do each one in one call (D36).

## Priority 2: the planned direction

No item is open.

## Priority 3: config loading and smaller faults

### B13. Two stand-ins wait for TinyGo to support Go 1.26.

- Where: `fields` at `cmd.go:1166-1179`, and `asExitError` at
  `ox.go:728-736`.
- Text: `// TODO: remove this after tinygo supports go1.26+` and
  `// TODO: remove when tinygo supports go1.26.`
- Plain description: CI runs TinyGo 0.44.0. Use `reflect.Type.Fields` and
  `errors.AsType` when every compiler that `ox` supports has them.
- Found: TODO. Lint: `modernize` reports the `NumField` loop at
  `cmd.go:1173`.
- Still open on 2026-11-27: a probe with TinyGo 0.44.0 and Go 1.27 found
  that TinyGo has both `reflect.Type.Fields` and `errors.AsType`. But each
  `go.mod` declares `go 1.25` (D24), and Go 1.25 has neither. So the
  stand-ins stay until the `go` line moves to `go 1.26`, when Go 1.28 comes
  out. CI pins TinyGo 0.44.0 since 2026-11-27, so TinyGo does not block it.

