package ox

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	// DefaultMaxPathLinks are the max symlinks to follow in [Path].
	DefaultMaxPathLinks = 16
	// DefaultGetwd is the func used to read the working directory by [Path].
	// Normally [os.Getwd].
	DefaultGetwd = os.Getwd
	// DefaultLstat is the func used to lstat a path by [Path]. Normally
	// [os.Lstat].
	DefaultLstat = os.Lstat
	// DefaultReadlink is the func used to read a link by [Path]. Normally
	// [os.Readlink].
	DefaultReadlink = os.Readlink
)

// Realpath resolves a path to a fully qualified ("real") path on disk.
func Realpath(path string) (string, error) {
	switch {
	case len(path) == 0:
		return "", os.ErrInvalid
	case !filepath.IsAbs(path):
		wd, err := DefaultGetwd()
		if err != nil {
			return "", err
		}
		path = filepath.Join(wd, path)
	}
	buf := []byte(path)
	for count, n, prev := 0, 1, 1; n < len(buf); {
		comp := buf
		if i := bytes.IndexByte(comp[n:], os.PathSeparator); i != -1 {
			comp = comp[:n+i]
		}
		b := comp[n:]
		switch {
		case len(b) == 0:
			copy(buf[n:], buf[n+1:])
			buf = buf[:len(buf)-1]
		case len(b) == 1 && b[0] == '.':
			if n+2 < len(buf) {
				copy(buf[n:], buf[n+2:])
			}
			buf = buf[:len(buf)-2]
		case len(b) == 2 && b[0] == '.' && b[1] == '.':
			copy(buf[prev:], buf[n+2:])
			buf, prev, n = buf[:len(buf)+prev-(n+2)], 1, 1
		default:
			switch fi, err := DefaultLstat(string(comp)); {
			case err != nil:
				return "", err
			case fi.Mode()&os.ModeSymlink == os.ModeSymlink:
				if count++; DefaultMaxPathLinks < count {
					return "", os.ErrInvalid
				}
				target, err := DefaultReadlink(string(comp))
				if err != nil {
					return "", err
				}
				// start again with the target in place of the link
				rest := string(buf[len(comp):])
				if filepath.IsAbs(target) {
					buf = []byte(filepath.Join(target, rest))
				} else {
					buf = []byte(filepath.Join(string(buf[:n]), target, rest))
				}
				prev, n = 1, 1
			default:
				prev, n = n, len(comp)+1
			}
		}
	}
	for len(buf) > 1 && buf[len(buf)-1] == os.PathSeparator {
		buf = buf[:len(buf)-1]
	}
	return string(buf), nil
}

// isPath returns true when typ is [PathT], [FileT] or [DirT].
func isPath(typ Type) bool {
	return typ == PathT || typ == FileT || typ == DirT
}

// pathValue returns the value of a path flag. It expands a leading `~` to the
// home directory, then makes the path absolute or real, as [Flag.Abs] and
// [Flag.Real] ask. When check is true, it checks that a [FileT] is a file
// that exists, that a [DirT] is a directory that exists, and that a path with
// [Flag.NotExist] does not exist.
func (g *Flag) pathValue(s string, check bool) (string, error) {
	if s == "" {
		return s, nil
	}
	if s == "~" || strings.HasPrefix(s, "~/") || strings.HasPrefix(s, "~"+string(os.PathSeparator)) {
		home, err := userHomeDir()
		if err != nil {
			return "", fmt.Errorf("%w: %s: %w", ErrInvalidValue, s, err)
		}
		s = filepath.Join(home, s[1:])
	}
	var err error
	switch {
	case g.Real:
		var r string
		switch r, err = Realpath(s); {
		case err == nil:
			s = r
		case errors.Is(err, fs.ErrNotExist):
			// a path that does not exist has no link to resolve
			s, err = absPath(s)
		}
	case g.Abs:
		s, err = absPath(s)
	}
	if err != nil {
		return "", fmt.Errorf("%w: %s: %w", ErrInvalidValue, s, err)
	}
	if !check {
		return s, nil
	}
	fi, err := os.Stat(s)
	switch {
	case err != nil && !errors.Is(err, fs.ErrNotExist):
		return "", fmt.Errorf("%w: %s: %w", ErrInvalidValue, s, err)
	case err == nil && g.NotExist:
		return "", fmt.Errorf("%w: %s: %w", ErrInvalidValue, s, ErrExist)
	case err != nil && !g.NotExist && g.Type != PathT:
		return "", fmt.Errorf("%w: %s: %w", ErrInvalidValue, s, ErrNotExist)
	case err == nil && g.Type == FileT && fi.IsDir():
		return "", fmt.Errorf("%w: %s: %w", ErrInvalidValue, s, ErrIsDir)
	case err == nil && g.Type == DirT && !fi.IsDir():
		return "", fmt.Errorf("%w: %s: %w", ErrInvalidValue, s, ErrNotDir)
	}
	return s, nil
}

// absPath joins a relative path to the working directory of [DefaultGetwd].
func absPath(s string) (string, error) {
	if filepath.IsAbs(s) {
		return filepath.Clean(s), nil
	}
	wd, err := DefaultGetwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, s), nil
}

// homePath returns s with the home directory written as `~`, when s is
// under the home directory.
func homePath(s string) string {
	home, err := userHomeDir()
	if err != nil || home == "" {
		return s
	}
	switch {
	case s == home:
		return "~"
	case strings.HasPrefix(s, home+string(os.PathSeparator)):
		return "~" + s[len(home):]
	}
	return s
}

// compPath returns the completions of a path flag: the extensions of
// [Flag.Exts] for a file, and only directories for a [DirT].
func (g *Flag) compPath() ([]Completion, CompDirective) {
	switch {
	case g.Type == DirT:
		return nil, CompFilterDirs
	case len(g.Exts) != 0:
		comps := make([]Completion, len(g.Exts))
		for i, ext := range g.Exts {
			comps[i] = Completion{Name: strings.TrimPrefix(ext, ".")}
		}
		return comps, CompFilterFileExt
	}
	return nil, CompDefault
}
