# Importing the cobra completion scripts

`comp/` holds the four shell scripts that `ox` writes for
`<program> completion <shell>`. They come from
[cobra](https://github.com/spf13/cobra), and `ox` changes a few lines in
each. `_tools/cobraimport` imports them from a cobra release. Import them
again after each cobra release. See D9 and D23 in [PLAN.md](PLAN.md).

The current import is cobra v1.10.2, from 2026-10-15. The one before it was
cobra v1.8.1, from `d1583b0` (2024-12-06).

## The files

| File in `comp/` | Comes from |
| --- | --- |
| `bash.bash` | `bash_completionsV2.go` in cobra |
| `fish.fish` | `fish_completions.go` in cobra |
| `powershell.ps1` | `powershell_completions.go` in cobra |
| `zsh.zsh` | `zsh_completions.go` in cobra |
| `bash.txt`, `fish.txt`, `powershell.txt`, `zsh.txt` | `ox` itself. The importer does not change them |

Each `.txt` file is the help text of one `completion <shell>` command. Each
other file is the template of one script. `DefaultCompWrite` in `ox.go`
formats the template with four values:

| Verb | Value |
| --- | --- |
| `%[1]s` | The name of the root command, such as `my-app` |
| `%[2]s` | The name with each character that is not a letter, a digit or `_` changed to `_`, such as `my_app` |
| `%[3]s` | `__complete`, or `__completeNoDesc` for `--no-descriptions` |
| `%[4]s` | The environment variable for active help, such as `MY_APP_ACTIVE_HELP` |

A template is a Go format string, so a `%` in a script is written `%%`.

## How to import

Do these steps in the repository root.

1. Run the importer. Without a version, it imports the latest cobra release:

   ```sh
   go run ./_tools/cobraimport
   ```

   To import one release, give its version, such as `v1.10.2`. The importer
   downloads cobra with `go mod download`, so `ox` does not require cobra in
   `go.mod`. It prints the version that it imported.
2. Run `git diff comp/`. If it shows nothing, cobra did not change its
   scripts, and the import is done.
3. If the importer fails, read the message, and then do the step for it in
   "When the import fails" below. Then run the importer again.
4. Read the diff. Make sure that each `# ox:` line is still in place and
   still does what "The ox edits" below says.
5. Run the tests, and render and run each script. See "How to test" below.
6. Change "The current import" at the top of this file to the new version
   and the date.
7. Make sure that the license of cobra did not change. `LICENSE.completions.txt`
   holds it.
8. Stage the change. The commit message names the cobra version, such as
   "Import the completion scripts of cobra v1.10.2".

## How the importer works

The importer reads each cobra source file with `go/parser`. It finds the one
format string that is longer than 2000 bytes, which is the script, and the
arguments that cobra formats it with. Then it changes each verb in the
string:

1. An argument that `ox` also passes becomes the `ox` verb of the table
   above: `name` is `%[1]s`, `nameForVar` is `%[2]s`, `compCmd` is `%[3]s`
   and `activeHelpEnvVar(name)` is `%[4]s`.
2. A cobra directive, such as `ShellCompDirectiveNoSpace`, becomes the value
   of the `ox` directive with the same meaning, such as `CompNoSpace`. A
   script and the `ox` code that answers it then agree on each value.
3. A string constant of cobra, such as `activeHelpMarker`, becomes its value.
4. A `%%` stays `%%`.

The importer maps each argument by its name, and not by its position. Each
cobra script orders its arguments in its own way. The fish template from
cobra v1.8.1 was converted by hand and had `%[1]s` and `%[2]s` swapped. A
program named `my-app` then got fish completion for `my_app`, and completion
did not work. The v1.10.2 import corrected this.

Then the importer makes the ox edits. Each edit replaces a text that must
occur exactly once in the converted script.

## The ox edits

`edits` in `_tools/cobraimport/main.go` holds them. Each one that changes
behavior adds a line that starts with `# ox:` to the script.

1. The first comment line of the bash, PowerShell and zsh scripts is
   `# <shell> completion for <name>`. Cobra writes
   `# bash completion V2 for <name> -*- shell-script -*-`, with the name
   padded to 36 characters. The fish script of cobra has no such line, so the
   importer adds one.
2. The bash, fish and PowerShell scripts offer every completion that the
   program returns. Cobra keeps only the completions that start with the
   text that the user typed. `ox` suggests a command name when the user types
   it wrong, with a Levenshtein distance of 2 (D11). `8eb319f` (2024-12-20),
   "Allow correction in completion", turned off the prefix match so that the
   shell shows the corrected name. In bash, the importer changes the match in
   two places: where the shell inserts the completions, and where it shows
   the descriptions. Zsh did not need a change.

The bash script still keeps only the completions that start with the typed
text when the program returns no descriptions, as with `--no-descriptions`.
The v1.8.1 import kept that match too. Question 11 of D28, in
[DESIGN.md](DESIGN.md), asks whether it goes.

## When the import fails

1. "found 0 copies of ..." means that cobra changed the text that an edit
   replaces. Find the new text in the converted script, and change the edit
   in `edits` so that it does the same thing.
2. "found 2 copies of ..." means that the old text of an edit is no longer
   unique. Add a line before or after it to the edit.
3. "has no mapping for it" means that cobra formats a script with a new
   argument. If it is a directive, add it to `CompDirective` in `defs.go` and
   to `directives` in the importer. If it is a value that changes for each
   program, add it to `DefaultCompWrite` in `ox.go`, to `args` in the
   importer, and to the table of verbs above.
4. "no script template found" means that cobra moved or split a script. Read
   the cobra source file, and change `scripts` in the importer.

## How to test

`go test ./...` does not run a script. Build the tree example with a name
that has a hyphen, because a hyphen is where `%[1]s` and `%[2]s` differ:

```sh
mkdir -p /tmp/cobraimport/bin
(cd _examples/tree && go build -ldflags "-X main.name=my-app" -o /tmp/cobraimport/bin/my-app .)
export PATH=/tmp/cobraimport/bin:$PATH
cd /tmp/cobraimport
for sh in bash zsh fish powershell; do my-app completion $sh > my-app.$sh; done
grep -c '%!' my-app.bash my-app.zsh my-app.fish my-app.powershell
bash -n my-app.bash && zsh -n my-app.zsh && fish -n my-app.fish
```

`grep -c '%!'` must print 0 for each file. A `%!` is a verb that the
template has and `ox` cannot fill. Then run a completion in fish and in bash.
Each must offer `completion` for the typo `compleiton`:

```sh
fish -c "source my-app.fish; complete -C 'my-app compleiton'"
bash -c 'source /usr/share/bash-completion/bash_completion; source my-app.bash
COMP_WORDS=(my-app compleiton); COMP_CWORD=1; COMP_LINE="my-app compleiton"
COMP_POINT=${#COMP_LINE}; COMP_TYPE=9; __start_my-app my-app compleiton my-app
printf "[%s]\n" "${COMPREPLY[@]}"'
```

Bash prints two `compopt` warnings, because it does not run in a terminal.
If `pwsh` is installed, make sure that the PowerShell script parses:

```sh
pwsh -NoProfile -Command '$e=$null; [System.Management.Automation.Language.Parser]::ParseFile("$PWD/my-app.powershell",[ref]$null,[ref]$e) | Out-Null; $e.Count'
```

It must print 0.
