//go:build !tinygo

package ox_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestGenHelpIsTheSame builds the programs in _examples/gen and compares the
// help of each with a file in testdata/gen. The files hold the help from
// before HelpColumn existed, so the test holds that a program that does not
// ask for HelpColumn writes the same help as before. The build constraint
// keeps the test out of the TinyGo run, because it runs the go command.
func TestGenHelpIsTheSame(t *testing.T) {
	t.Parallel()
	if runtime.GOOS != "linux" {
		t.Skip("the defaults of the programs name the config and cache folders of linux")
	}
	if testing.Short() {
		t.Skip("the test builds 11 programs")
	}
	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("the go command is not in PATH")
	}
	dir := t.TempDir()
	if buf, err := exec.Command(gobin, "build", "-o", dir+string(filepath.Separator), "./_examples/gen/...").CombinedOutput(); err != nil {
		t.Fatalf("building _examples/gen: %v\n%s", err, buf)
	}
	tests := [][]string{
		{"docker"},
		{"docker", "run"},
		{"doctl"},
		{"gh"},
		{"helm"},
		{"helm", "install"},
		{"hugo"},
		{"kubectl"},
		{"kubectl", "create"},
		{"omnictl"},
		{"podman"},
		{"psql"},
		{"rclone"},
		{"talosctl"},
	}
	for _, test := range tests {
		name := strings.Join(test, "-")
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			exp, err := os.ReadFile(filepath.Join("testdata", "gen", name+".txt"))
			if err != nil {
				t.Fatal(err)
			}
			cmd := exec.Command(filepath.Join(dir, test[0]), append(test[1:], "--help")...)
			cmd.Env = []string{"HOME=/home/user", "USER=user", "XDG_CONFIG_HOME=", "XDG_CACHE_HOME="}
			buf, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("running %s --help: %v\n%s", strings.Join(test, " "), err, buf)
			}
			if s := string(buf); s != string(exp) {
				t.Errorf("the help changed:\nexpected:\n%s\ngot:\n%s", exp, s)
			}
		})
	}
}
