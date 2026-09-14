package ox

import (
	"os"
	"path/filepath"
	"runtime"
)

// UserStateDir returns the application's state directory.
//
// Linux/BSD: $XDG_STATE_HOME/appName or ~/.local/state/appName
// macOS:     $XDG_STATE_HOME/appName, or ~/Library/Application Support/appName
// Windows:   %LOCALAPPDATA%\appName or %USERPROFILE%\AppData\Local\appName
func UserStateDir(appName string) (string, error) {
	if dir := os.Getenv("XDG_STATE_HOME"); dir != "" {
		return filepath.Join(dir, appName), nil
	}
	switch runtime.GOOS {
	case "windows":
		if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
			return filepath.Join(localAppData, appName), nil
		}
		dir, err := userHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "AppData", "Local", appName), nil
	case "darwin":
		dir, err := userHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "Library", "Application Support", appName), nil
	}
	// Linux, FreeBSD, OpenBSD, etc.
	dir, err := userHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, ".local", "state", appName), nil
}
