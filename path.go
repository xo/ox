package ox

import (
	"bytes"
	"os"
	"path/filepath"
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
				comp, err := DefaultReadlink(string(comp))
				if err != nil {
					return "", err
				}
				if l := string(buf[len(comp):]); l[0] == os.PathSeparator {
					buf = []byte(filepath.Join(l, string(buf[len(comp):])))
				} else {
					buf = []byte(filepath.Join(string(buf[:n]), l, string(buf[len(comp):])))
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
