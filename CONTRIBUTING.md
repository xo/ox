# Contributing to ox

`ox` parses command-line flags and arguments for Go and TinyGo programs.
`AGENTS.md` holds the full rules for a coding agent. This file is the short
form for a person.

## What a change must do

The root package imports only the standard library. A type or a config
loader that needs another module goes in a subpackage, such as `color` or
`yaml`, which a program turns on with a blank import. See D1 in
[docs/PLAN.md](docs/PLAN.md).

`ox` builds with TinyGo. Code that differs between Go and TinyGo goes in
`go.go` and `tinygo.go`. See D2.

`ox` makes no promise of backward compatibility. A change can rename or
remove an exported identifier when that makes the code better. `dbtpl`
requires `ox`, so name each such change in the pull request. See D18.

[docs/PLAN.md](docs/PLAN.md) holds each decision and the reason for it, and
the questions that are still open. [docs/BACKLOG.md](docs/BACKLOG.md) holds
the known faults and the work that is not done.

## Before you open a pull request

Run these in the repository root. Each must pass:

```sh
gofmt -l . && go test -race -count=1 ./...
golangci-lint run --new-from-rev=HEAD ./...
```

`gofmt -l .` must print nothing. CI also runs `tinygo test -v ./...` with
TinyGo 0.42.0.

## Writing

Write documents, code comments, error messages and commit messages in plain
English. Use short sentences and the active voice. Do not use contractions,
semicolons or em dashes. Use "must" and "can", not "should" and "may". The
`simple-english` skill in `.agents/skills` holds the full rules.

## Agent skills

The repository carries two agent skills. A skill is a set of instructions
that a coding agent loads for a task. `simple-english` sets how prose is
written, and `go-pedantry` sets how Go is written.

`skills-lock.json` names the source of each skill. The `skills` command from
npm writes that file. It writes each skill into two folders. Codex and other
agents read `.agents/skills/<name>`, and Claude Code reads
`.claude/skills/<name>`.

To add a skill or to update one, run its command in the repository root:

```sh
npx skills@1.7.0 add AminBlg/SimpleEnglish --skill simple-english --agent codex claude-code --copy -y
npx skills@1.7.0 add oborchers/fractional-cto --skill go-pedantry --agent codex claude-code --copy -y
```

Keep `--copy`. Without it, the command writes `.claude/skills/<name>` as a
symbolic link. A Windows checkout writes a symbolic link as a small text
file, and Claude Code then loads no skill and says nothing.
`TestSkillsAreCopies` fails on a link, and it fails when the two folders
differ.

`.claude/settings.local.json` holds the Claude Code permissions of one
person. The root `.gitignore` ignores it.
