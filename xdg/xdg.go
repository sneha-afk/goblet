// Package xdg provides utilities for working with the XDG specification
// See https://specifications.freedesktop.org/basedir/latest/
package xdg

import (
	"os"
	"path/filepath"
)

// ConfigHome returns the path of `$XDG_CONFIG_HOME` if available,
// else falling back to OS defaults (i.e `$HOME/.config` on Linux/BSD),
// which then falls back to `$HOME/.config`
//
// Does NOT guarantee the directory exists.
func ConfigHome() (string, error) {
	envvar, present := os.LookupEnv("XDG_CONFIG_HOME")
	if present && envvar != "" {
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

// CacheHome returns the path of `$XDG_CACHE_HOME` if available,
// else falling back to OS defaults (i.e `$HOME/.cache` on Linux/BSD),
// which then falls back to `$HOME/.cache`
//
// Does NOT guarantee the directory exists.
func CacheHome() (string, error) {
	envvar, present := os.LookupEnv("XDG_CACHE_HOME")
	if present && envvar != "" {
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

// DataHome returns the path of `$XDG_DATA_HOME` if available,
// else falling back to `$HOME/.local/share`
//
// Does NOT guarantee the directory exists.
func DataHome() (string, error) {
	envvar, present := os.LookupEnv("XDG_DATA_HOME")
	if present && envvar != "" {
		return envvar, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".local", "share"), nil
}
