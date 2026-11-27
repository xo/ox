//go:build !tinygo

package ox_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestUsqlSample builds the sample of usql in _examples/usql, and compares
// what it writes for each command line. The help is in
// testdata/usql/help.txt. The build constraint keeps the test out of the
// TinyGo run, because it runs the go command. See B42.
func TestUsqlSample(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("the test builds a program")
	}
	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("the go command is not in PATH")
	}
	bin := filepath.Join(t.TempDir(), "usql")
	if buf, err := exec.Command(gobin, "build", "-o", bin, "./_examples/usql").CombinedOutput(); err != nil {
		t.Fatalf("building _examples/usql: %v\n%s", err, buf)
	}
	help, err := os.ReadFile(filepath.Join("testdata", "usql", "help.txt"))
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		args []string
		exp  string
		fail bool
	}{
		{"help", []string{"--help"}, string(help), false},
		{"help with -?", []string{"-?"}, string(help), false},
		{"version", []string{"-V"}, "usql 0.0.0-dev\n", false},
		{
			"options in order", []string{"-v", "A=1", "-P", "x=2", "--set", "B=3", "-F", ",", "-c", "select 1", "-f", "a.sql", "-Aq", "--var=C=4", "pg://h"},
			"dsn: pg://h\ncommands and files: [command:select 1 file:a.sql]\noptions: [vars.A=1 print.x=2 vars.B=3 print.fieldsep=\",\" print.csv_fieldsep=\",\" print.format=unaligned usql.quiet=on vars.C=4]\nout: \nconfig: \nno init: false single transaction: false\nno password: false force password: false\ndump: false []\n",
			false,
		},
		{
			"switches and dump options", []string{"-X", "-1", "-w", "-o", "out.txt", "--config", "u.yaml", "-D", "--target", "duckdb", "--schema-only", "-d", "pg://h"},
			"dsn: pg://h\ncommands and files: []\noptions: []\nout: out.txt\nconfig: u.yaml\nno init: true single transaction: true\nno password: true force password: false\ndump: true [target=duckdb schema-only=true]\n",
			false,
		},
		{"hidden and deprecated aliases", []string{"--no-psqlrc", "--no-rc"}, "dsn: \ncommands and files: []\noptions: []\nout: \nconfig: \nno init: true single transaction: false\nno password: false force password: false\ndump: false []\n", false},
		{"psql's -h fails", []string{"-h", "host"}, "usql: -h: unknown flag\nTry \"usql --help\" for more information.\n", true},
		{"two DSNs", []string{"a", "b"}, "usql: command usql: wrong number of arguments\nTry \"usql --help\" for more information.\n", true},
		{"an optional value is refused", []string{"--no-init=maybe"}, "usql: --no-init: invalid value: strconv.ParseBool: parsing \"maybe\": invalid syntax\nTry \"usql --help\" for more information.\n", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cmd := exec.Command(bin, test.args...)
			cmd.Env = []string{"HOME=/home/user", "USER=user"}
			buf, err := cmd.CombinedOutput()
			switch {
			case test.fail && err == nil:
				t.Fatalf("expected an error")
			case !test.fail && err != nil:
				t.Fatalf("expected no error, got: %v\n%s", err, buf)
			}
			if s := string(buf); s != test.exp {
				t.Errorf("expected:\n%s\ngot:\n%s\nexpected: %q\ngot:      %q", test.exp, s, test.exp, s)
			}
		})
	}
}
