package test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/kenshaw/colors"
	"github.com/kenshaw/glob"
	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

// TestColor holds that a flag of type *colors.Color reads each form of a
// color with no registration (D41, D42).
func TestColor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		s   string
		exp string
	}{
		{"red", "red"},
		{"BLACK", "black"},
		{"rgb(0,255,0)", "lime"},
		{"rgba( 0, 255, 0, 255)", "lime"},
		{"hex(0,ff,0)", "lime"},
		{"hex(0,ff,0,ff)", "lime"},
		{"#00ff00", "lime"},
		{"#00ff00ff", "lime"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			t.Parallel()
			var s string
			err := ox.RunError(
				context.Background(),
				ox.Usage("app", "an app"),
				ox.Flags().
					Var("fg", "the foreground color", ox.TypeFor[*colors.Color]()),
				ox.Exec(func(ctx context.Context) {
					s = otx.Get[*colors.Color](ctx, "fg").AsText()
				}),
				ox.Args("--fg", test.s),
			)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestGlob holds that a field of type []*glob.Glob reads each pattern with no
// registration, and that help writes the type as glob (D41, D42).
func TestGlob(t *testing.T) {
	t.Parallel()
	// From splits a slice value at each comma, so no pattern has a comma
	patterns := []string{"config.*", "file.txt", "[A-Z]*_test.go"}
	var args struct {
		Include []*glob.Glob `ox:"the files to include,short:i"`
		Color   colors.Color `ox:"the color"`
	}
	var out bytes.Buffer
	opts := []ox.Option{
		ox.Usage("app", "an app"),
		ox.From(&args),
		ox.Exec(func() {}),
	}
	cmdline := []string{"--color", "white"}
	for _, p := range patterns {
		cmdline = append(cmdline, "-i", p)
	}
	if err := ox.RunError(context.Background(), append(opts, ox.Args(cmdline...))...); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(args.Include) != len(patterns) {
		t.Fatalf("expected %d globs, got: %d", len(patterns), len(args.Include))
	}
	for i, p := range patterns {
		if s := args.Include[i].String(); s != p {
			t.Errorf("expected %q, got: %q", p, s)
		}
	}
	if s := args.Color.AsText(); s != "white" {
		t.Errorf("expected white, got: %q", s)
	}
	var help struct {
		Include []*glob.Glob `ox:"the files to include,short:i"`
		Color   colors.Color `ox:"the color"`
	}
	cmd, err := ox.NewCommand(ox.Usage("app", "an app"), ox.From(&help))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if _, err := cmd.WriteTo(&out); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	for _, exp := range []string{"--include glob", "--color color"} {
		if !strings.Contains(out.String(), exp) {
			t.Errorf("expected help to hold %q, got:\n%s", exp, out.String())
		}
	}
}
