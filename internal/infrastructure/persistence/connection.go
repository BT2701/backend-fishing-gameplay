package persistence

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// RetryConfig holds retry parameters
type RetryConfig struct {
	MaxRetries int // number of retry attempts
	RetryDelay int // delay between retries in milliseconds
}

// RetryWithBackoff executes a function with exponential backoff retry logic
func RetryWithBackoff(maxRetries int, retryDelayMs int, logger *zap.Logger, operation string, fn func() error) error {
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		if err := fn(); err == nil {
			return nil
		} else {
			lastErr = err
			if attempt < maxRetries {
				delay := time.Duration(retryDelayMs*(attempt-1)) * time.Millisecond
				logger.Warn(
					fmt.Sprintf("Failed to %s, retrying...", operation),
					zap.Int("attempt", attempt),
					zap.Int("max_attempts", maxRetries),
					zap.Duration("delay", delay),
					zap.Error(err),
				)
				time.Sleep(delay)
			}
		}
	}

	return fmt.Errorf("failed to %s after %d attempts: %w", operation, maxRetries, lastErr)
}
