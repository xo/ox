package ox

import (
	"os"
	"path/filepath"
	"runtime"
)

// UserStateDir returns the user's state directory.
//
// Linux/BSD: $XDG_STATE_HOME or ~/.local/state
// macOS:     $XDG_STATE_HOME, or ~/Library/Application Support
// Windows:   %LOCALAPPDATA% or %USERPROFILE%\AppData\Local
func UserStateDir() (string, error) {
	if dir := os.Getenv("XDG_STATE_HOME"); dir != "" {
		return dir, nil
	}
	switch runtime.GOOS {
	case "windows":
		if dir := os.Getenv("LOCALAPPDATA"); dir != "" {
			return dir, nil
		}
		dir, err := userHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "AppData", "Local"), nil
	case "darwin":
		dir, err := userHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "Library", "Application Support"), nil
	}
	// Linux, FreeBSD, OpenBSD, etc.
	dir, err := userHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, ".local", "state"), nil
}
