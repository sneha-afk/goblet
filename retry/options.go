package retry

import "fmt"

type Config struct {
	Attempts int
}

// Option is defined as such to mutate an running config struct
type Option func(*Config)

// Validate a Config struct, that may or may not be complete.
func (cfg Config) Validate() error {
	if cfg.Attempts < 0 {
		return fmt.Errorf("precondition: Attempts should be non-zero")
	}

	return nil
}

// Attempts defines a simple constant number of retries.
func Attempts(n int) Option {

	return func(c *Config) {
		c.Attempts = n
	}
}
