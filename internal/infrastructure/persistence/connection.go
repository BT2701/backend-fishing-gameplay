package persistence

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Logger interface for dependency injection
type Logger interface {
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

// LoggerAdapter wraps zap.Logger to implement Logger interface
type LoggerAdapter struct {
	logger *zap.Logger
}

// NewLoggerAdapter creates a new logger adapter
func NewLoggerAdapter(zapLogger *zap.Logger) Logger {
	return &LoggerAdapter{logger: zapLogger}
}

func (l *LoggerAdapter) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Warnw(msg, keysAndValues...)
}

func (l *LoggerAdapter) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Errorw(msg, keysAndValues...)
}

// RetryConfig holds retry parameters
type RetryConfig struct {
	MaxRetries int // number of retry attempts
	RetryDelay int // delay between retries in milliseconds
}

// RetryWithBackoff executes a function with exponential backoff retry logic
func RetryWithBackoff(maxRetries int, retryDelayMs int, logger Logger, operation string, fn func() error) error {
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
					"attempt", attempt,
					"max_attempts", maxRetries,
					"delay", delay.String(),
					"error", err.Error(),
				)
				time.Sleep(delay)
			}
		}
	}

	return fmt.Errorf("failed to %s after %d attempts: %w", operation, maxRetries, lastErr)
}
