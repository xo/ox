package main

import (
	"bytes"
	"flag"
	"testing"
	"time"
)

func TestGoFlag(t *testing.T) {
	root := importTestdata(t, "goimports")
	if root.Format != FormatGoFlag {
		t.Fatalf("expected format %q, got: %q", FormatGoFlag, root.Format)
	}
	if exp := "[flags] [path ...]"; root.Spec != exp {
		t.Errorf("expected spec %q, got: %q", exp, root.Spec)
	}
	if n := len(root.Flags); n != 12 {
		t.Errorf("expected 12 flags, got: %d", n)
	}
	checkFlags(t, root, []testFlag{
		{"cpuprofile", "", "string", "", "", -1, nil},
		{"d", "d", "bool", "", "", -1, nil},
		{"format-only", "", "bool", "", "", -1, nil},
		{"memrate", "", "int", "", "", -1, nil},
		{"srcdir", "", "string", "dir", "", -1, nil},
	})
}

// TestGoFlagDefaults reads the help that the flag package of this Go writes.
func TestGoFlagDefaults(t *testing.T) {
	fs := flag.NewFlagSet("prog", flag.ContinueOnError)
	var buf bytes.Buffer
	fs.SetOutput(&buf)
	_ = fs.String("name", "ken", "the `user` name")
	_ = fs.Int("n", 10, "number of items")
	_ = fs.Duration("wait", 5*time.Second, "time to wait,\nwith a second line")
	_ = fs.Bool("x", false, "turn on x")
	_ = fs.Float64("ratio", 0.5, "a ratio")
	_ = fs.String("dir", "/home/user/.cache", "cache directory")
	_, _ = buf.WriteString("usage: prog [flags]\n")
	fs.PrintDefaults()
	help := buf.String()
	if !isGoFlag(help) {
		t.Fatalf("expected the help of the flag package:\n%s", help)
	}
	cmd := &Command{Name: "prog", Section: -1}
	parseGoFlag(cmd, help, "/home/user")
	checkFlags(t, cmd, []testFlag{
		{"name", "", "string", "user", "ken", -1, nil},
		{"n", "n", "int", "", "10", -1, nil},
		{"wait", "", "duration", "", "5s", -1, nil},
		{"x", "x", "bool", "", "", -1, nil},
		{"ratio", "", "float64", "", "0.5", -1, nil},
		{"dir", "", "string", "", "$HOME/.cache", -1, nil},
	})
	if g := cmd.Flag("wait"); g != nil {
		if exp := "time to wait,\nwith a second line"; g.Usage != exp {
			t.Errorf("expected usage %q, got: %q", exp, g.Usage)
		}
	}
	if g := cmd.Flag("name"); g != nil {
		if exp := "the user name"; g.Usage != exp {
			t.Errorf("expected usage %q, got: %q", exp, g.Usage)
		}
	}
}
