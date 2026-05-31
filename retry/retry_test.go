package retry_test

import (
	"fmt"
	"testing"

	"github.com/sneha-afk/goblet/retry"
)

func TestIntAttempts(t *testing.T) {
	tests := []struct {
		name        string
		numAttempts int
		expectErr   bool
	}{
		{
			name:        "retries a non-zero attempt of times",
			numAttempts: 3,
			expectErr:   false,
		},
		{
			name:        "zero attempts is okay",
			numAttempts: 0,
			expectErr:   false,
		},
		{
			name:        "negative attempts doesn't make sense",
			numAttempts: -1,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attempts := 0

			err := retry.Do(func() error {
				// this func will succeed on the last attempt
				if attempts+1 >= tt.numAttempts {
					return nil
				}

				attempts++
				return fmt.Errorf(":(")
			}, retry.Attempts(tt.numAttempts))

			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, but got none: test '%v'", tt.name)
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect error, but got one: %v", err)
				}

				if attempts > tt.numAttempts {
					t.Fatalf("took too many attempts: %v", attempts)
				}
			}
		})
	}
}
