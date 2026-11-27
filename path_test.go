package ox

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// TestPathTypes holds the checks and the options of [PathT], [FileT] and
// [DirT] (D26).
func TestPathTypes(t *testing.T) {
	dir, err := Realpath(t.TempDir())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	file := filepath.Join(dir, "file.yaml")
	if err := os.WriteFile(file, nil, 0o600); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	sub := filepath.Join(dir, "sub")
	if err := os.Mkdir(sub, 0o700); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	link := filepath.Join(dir, "link")
	if err := os.Symlink(sub, link); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	missing := filepath.Join(dir, "missing")
	t.Setenv("HOME", dir)
	prev := DefaultGetwd
	DefaultGetwd = func() (string, error) {
		return dir, nil
	}
	t.Cleanup(func() {
		DefaultGetwd = prev
	})
	tests := []struct {
		name string
		typ  Type
		opts []Option
		args []string
		exp  string
		err  error
	}{
		{"path", PathT, nil, []string{"--v", file}, file, nil},
		{"path that does not exist", PathT, nil, []string{"--v", missing}, missing, nil},
		{"path is relative", PathT, nil, []string{"--v", "sub"}, "sub", nil},
		{"path expands a tilde", PathT, nil, []string{"--v", "~/sub"}, sub, nil},
		{"path expands a lone tilde", PathT, nil, []string{"--v", "~"}, dir, nil},
		{"path keeps a tilde with a name", PathT, nil, []string{"--v", "~user"}, "~user", nil},
		{"path abs", PathT, []Option{Abs()}, []string{"--v", "sub"}, sub, nil},
		{"path real", PathT, []Option{Real()}, []string{"--v", "link"}, sub, nil},
		{"path real that does not exist", PathT, []Option{Real()}, []string{"--v", "missing"}, missing, nil},
		{"path not exist", PathT, []Option{NotExist()}, []string{"--v", missing}, missing, nil},
		{"path not exist that exists", PathT, []Option{NotExist()}, []string{"--v", file}, "", ErrExist},
		{"file", FileT, nil, []string{"--v", file}, file, nil},
		{"file that does not exist", FileT, nil, []string{"--v", missing}, "", ErrNotExist},
		{"file is a directory", FileT, nil, []string{"--v", sub}, "", ErrIsDir},
		{"file with another extension", FileT, []Option{Ext(".toml")}, []string{"--v", file}, file, nil},
		{"file not exist", FileT, []Option{NotExist()}, []string{"--v", missing}, missing, nil},
		{"file default is not checked", FileT, []Option{Default(missing)}, nil, missing, nil},
		{"dir", DirT, nil, []string{"--v", sub}, sub, nil},
		{"dir through a link", DirT, nil, []string{"--v", link}, link, nil},
		{"dir real", DirT, []Option{Real()}, []string{"--v", link}, sub, nil},
		{"dir that does not exist", DirT, nil, []string{"--v", missing}, "", ErrNotExist},
		{"dir is a file", DirT, nil, []string{"--v", file}, "", ErrNotDir},
		{"dir from a tilde", DirT, []Option{Abs()}, []string{"--v", "~"}, dir, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var s string
			err := RunError(
				context.Background(),
				Usage("app", "an app"),
				Flags().Var("v", "a path", append([]Option{Option(test.typ), Bind(&s)}, test.opts...)...),
				Exec(func() {}),
				Args(test.args...),
			)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.err != nil && !errors.Is(err, ErrInvalidValue):
				t.Fatalf("expected error %v, got: %v", ErrInvalidValue, err)
			case s != test.exp:
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestPathHome holds that [Home] shows a default under the home directory
// as `~` in help, and keeps the value.
func TestPathHome(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("HOME", dir)
	tests := []struct {
		name string
		opts []Option
		exp  string
	}{
		{"home", []Option{Home()}, "~/x"},
		{"no home", nil, filepath.Join(dir, "x")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := runHelp(t,
				Usage("app", "an app"),
				Help(),
				Flags().Path("config", "the config", append([]Option{Default("$HOME/x")}, test.opts...)...),
			)
			if exp := "(default: " + test.exp + ")"; !strings.Contains(s, exp) {
				t.Errorf("expected %q in:\n%s", exp, s)
			}
		})
	}
}

// TestRealpath holds that [Realpath] resolves each symbolic link, `.` and
// `..`.
func TestRealpath(t *testing.T) {
	dir, err := Realpath(t.TempDir())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	sub := filepath.Join(dir, "sub")
	file := filepath.Join(sub, "file")
	for _, f := range []func() error{
		func() error { return os.Mkdir(sub, 0o700) },
		func() error { return os.WriteFile(file, nil, 0o600) },
		func() error { return os.Symlink(sub, filepath.Join(dir, "abs")) },
		func() error { return os.Symlink("sub", filepath.Join(dir, "rel")) },
		func() error { return os.Symlink("../rel/file", filepath.Join(sub, "up")) },
		func() error { return os.Symlink("loop", filepath.Join(dir, "loop")) },
	} {
		if err := f(); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	}
	prev := DefaultGetwd
	DefaultGetwd = func() (string, error) {
		return dir, nil
	}
	t.Cleanup(func() {
		DefaultGetwd = prev
	})
	tests := []struct {
		name string
		path string
		exp  string
		err  error
	}{
		{"dir", sub, sub, nil},
		{"file", file, file, nil},
		{"relative", "sub/file", file, nil},
		{"dot", filepath.Join(dir, ".", "sub"), sub, nil},
		{"dot dot", sub + "/../sub", sub, nil},
		{"trailing separator", sub + "/", sub, nil},
		{"absolute link", filepath.Join(dir, "abs"), sub, nil},
		{"relative link", filepath.Join(dir, "rel"), sub, nil},
		{"link inside the path", filepath.Join(dir, "rel", "file"), file, nil},
		{"link to a link", filepath.Join(sub, "up"), file, nil},
		{"missing", filepath.Join(dir, "missing"), "", os.ErrNotExist},
		{"loop", filepath.Join(dir, "loop"), "", os.ErrInvalid},
		{"empty", "", "", os.ErrInvalid},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, err := Realpath(test.path)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case s != test.exp:
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestPathTags holds the tag options of a path flag.
func TestPathTags(t *testing.T) {
	var v struct {
		Config string `ox:"the config,type:file,ext:.yaml|.yml,abs,home"`
		Out    string `ox:"the output,type:path,notexist,real"`
		Dir    string `ox:"the directory,type:dir"`
	}
	cmd, err := NewCommand(Usage("app", "an app"), From(&v))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		name string
		exp  Flag
	}{
		{"config", Flag{Type: FileT, Exts: []string{".yaml", ".yml"}, Abs: true, Home: true}},
		{"out", Flag{Type: PathT, NotExist: true, Real: true}},
		{"dir", Flag{Type: DirT}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := cmd.Flag(test.name, false, false)
			switch {
			case g == nil:
				t.Fatalf("expected a flag %s", test.name)
			case g.Type != test.exp.Type, g.Abs != test.exp.Abs, g.Real != test.exp.Real,
				g.Home != test.exp.Home, g.NotExist != test.exp.NotExist,
				!slices.Equal(g.Exts, test.exp.Exts):
				t.Errorf("expected %+v, got: %+v", test.exp, *g)
			}
		})
	}
}
