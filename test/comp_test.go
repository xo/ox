package test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestShellCompletion runs a completion of my-app, in testdata/myapp, in
// bash, zsh, fish and PowerShell, with the script that ox writes for each.
// Each case checks the word that the shell puts on the command line. A shell
// that is not installed is skipped. See D23, B47 and B48.
func TestShellCompletion(t *testing.T) {
	if testing.Short() {
		t.Skip("the test builds a program and runs shells")
	}
	if runtime.GOOS == "windows" {
		t.Skip("the harnesses are for a Unix system")
	}
	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("the go command is not in PATH")
	}
	dir := t.TempDir()
	if buf, err := exec.Command(gobin, "build", "-o", filepath.Join(dir, "my-app"), "./testdata/myapp").CombinedOutput(); err != nil {
		t.Fatalf("building my-app: %v\n%s", err, buf)
	}
	env := append(os.Environ(), "PATH="+dir+string(os.PathListSeparator)+os.Getenv("PATH"), "HOME="+dir)
	script := func(t *testing.T, shell, name string) {
		t.Helper()
		cmd := exec.Command(filepath.Join(dir, "my-app"), "completion", shell)
		cmd.Env = env
		buf, err := cmd.Output()
		if err != nil {
			t.Fatalf("my-app completion %s: %v", shell, err)
		}
		if err := os.WriteFile(filepath.Join(dir, name), buf, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	shells := []struct {
		name   string
		bin    string
		file   string
		script string
		args   func(line string) []string
	}{
		{
			"bash", "bash", "my-app.bash", "bash",
			func(line string) []string {
				return []string{"-c", `source /usr/share/bash-completion/bash_completion 2>/dev/null; source my-app.bash
COMP_LINE="$1"; read -ra COMP_WORDS <<< "$COMP_LINE"; COMP_CWORD=$((${#COMP_WORDS[@]}-1))
COMP_POINT=${#COMP_LINE}; COMP_TYPE=9; __start_my-app my-app "${COMP_WORDS[COMP_CWORD]}" "${COMP_WORDS[COMP_CWORD-1]}" 2>/dev/null
printf "%s\n" "${COMPREPLY[@]}"`, "bash", line}
			},
		},
		{
			"fish", "fish", "my-app.fish", "fish",
			func(line string) []string {
				return []string{"-c", "source my-app.fish; complete -C '" + line + "'"}
			},
		},
		{
			"zsh", "zsh", "_my-app", "zsh",
			func(line string) []string {
				// zsh completes only in a terminal, so zpty drives one. The
				// script waits for zsh to print each marker, and does not
				// sleep for a fixed time, because compinit and the completion
				// take longer on a slow or busy machine. The marker READY is
				// written as RE''ADY, so that the echo of the typed line does
				// not match it. After the tab, the line moves into an echo,
				// and the script waits for the line that the echo writes
				return []string{"-f", "-c", `zmodload zsh/zpty
zpty z "zsh -f -i"
zpty -w z "fpath=($PWD \$fpath); autoload -U compinit; compinit -u; PS1='> '; print RE''ADY"
zpty -r -m z out '*READY*' || exit 1
zpty -w -n z "$1"$'\t'
zpty -w -n z $'\x01'"echo RESULT: "$'\n'
zpty -r -m z out '*'$'\n''RESULT: *'$'\n''*' || exit 1
zpty -d z
print -r -- "$out" | tr -d '\r' | grep -a '^RESULT: ' | tail -1`, "zsh", line}
			},
		},
		{
			"powershell", "pwsh", "my-app.ps1", "powershell",
			func(line string) []string {
				return []string{"-NoProfile", "-Command", ". ./my-app.ps1; (TabExpansion2 -inputScript '" + line + "' -cursorColumn " + strconv.Itoa(len(line)) + ").CompletionMatches | ForEach-Object { $_.CompletionText }"}
			},
		},
	}
	tests := []struct {
		name string
		line string
		exp  string
	}{
		{"a prefix", "my-app sh", "shell"},
		{"a correction", "my-app compleiton", "completion"},
		{"a value", "my-app shell z", "zsh"},
		{"a corrected value", "my-app completion zhs", "zsh"},
	}
	for _, shell := range shells {
		t.Run(shell.name, func(t *testing.T) {
			bin, err := exec.LookPath(shell.bin)
			if err != nil {
				t.Skipf("%s is not installed", shell.bin)
			}
			if shell.name == "bash" {
				if _, err := os.Stat("/usr/share/bash-completion/bash_completion"); err != nil {
					t.Skip("bash-completion is not installed")
				}
			}
			script(t, shell.script, shell.file)
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					// a shell that waits for output that never comes stops
					// at the deadline
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					cmd := exec.CommandContext(ctx, bin, shell.args(test.line)...)
					cmd.Dir, cmd.Env = dir, env
					buf, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatalf("expected no error, got: %v\n%s", err, buf)
					}
					if !completes(shell.name, string(buf), test.line, test.exp) {
						t.Errorf("expected %q for %q, got:\n%s", test.exp, test.line, buf)
					}
				})
			}
		})
	}
}

// completes returns true when the output of a shell completes line with exp.
func completes(shell, out, line, exp string) bool {
	if shell == "zsh" {
		fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(out), "RESULT: "))
		return len(fields) != 0 && fields[len(fields)-1] == exp
	}
	for s := range strings.Lines(out) {
		// fish writes a description after a tab, and PowerShell and bash can
		// add a space
		name, _, _ := strings.Cut(strings.TrimRight(s, "\r\n"), "\t")
		if strings.TrimSpace(name) == exp {
			return true
		}
	}
	return false
}
