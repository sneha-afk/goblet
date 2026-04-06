// Package paths provides utilities for handling and retrieving filepaths.
package paths

import (
	"os"
	"path/filepath"
)

// XDGConfigHome returns the path of `$XDG_CONFIG_HOME` if available,
// else falling back to OS defaults (i.e `$HOME/.config` on Linux/BSD),
// which then falls back to `$HOME/.config`
//
// Does NOT guarantee the directory exists.
func XDGConfigHome() (string, error) {
	envvar, present := os.LookupEnv("XDG_CONFIG_HOME")
	if present {
		return envvar, nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, ".config"), nil
	}
	return dir, err
}

// XDGCacheHome returns the path of `$XDG_CACHE_HOME` if available,
// else falling back to OS defaults (i.e `$HOME/.cache` on Linux/BSD),
// which then falls back to `$HOME/.cache`
//
// Does NOT guarantee the directory exists.
func XDGCacheHome() (string, error) {
	envvar, present := os.LookupEnv("XDG_CACHE_HOME")
	if present {
		return envvar, nil
	}

	dir, err := os.UserCacheDir()
	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, ".cache"), nil
	}
	return dir, err

}

// XDGDataHome returns the path of `$XDG_DATA_HOME` if available,
// else falling back to `$HOME/.local/share`
//
// Does NOT guarantee the directory exists.
func XDGDataHome() (string, error) {
	envvar, present := os.LookupEnv("XDG_DATA_HOME")
	if present {
		return envvar, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".local", "share"), nil
}
