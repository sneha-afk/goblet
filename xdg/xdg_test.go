package xdg

import (
	"os"
	"testing"
)

func TestConfigHomeFallback(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "")

	got, err := ConfigHome()
	if err != nil {
		t.Fatal(err)
	}

	want, err := os.UserConfigDir()
	if err != nil {
		t.Skip("UserConfigDir unavailable")
	}

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
