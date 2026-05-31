// Package retry wraps around execution with retryable options.
package retry

func DefaultConfig() Config {
	return Config{
		Attempts: 5,
	}
}

func Do(fn func() error, opts ...Option) error {
	// each option will mutate cfg
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	var lastError error
	for range cfg.Attempts {
		err := fn()

		if err == nil {
			return nil
		} else {
			lastError = err
			// TODO: backoff options here
		}

	}

	return lastError
}
