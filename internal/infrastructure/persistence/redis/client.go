package redis

import (
	"context"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/persistence"
	redislib "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func ConnectWithRetry(addr string, password string, db int, maxRetries int, retryDelayMs int, logger persistence.Logger) (*redislib.Client, error) {
	var client *redislib.Client

	err := persistence.RetryWithBackoff(maxRetries, retryDelayMs, logger, "connect to Redis", func() error {
		c := redislib.NewClient(&redislib.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		// Verify connection with ping
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		status := c.Ping(ctx)
		if status.Err() != nil {
			c.Close()
			return status.Err()
		}

		client = c
		return nil
	})

	return client, err
}

// ConnectWithRetryZap is a helper function that wraps zap.Logger for backwards compatibility
func ConnectWithRetryZap(addr string, password string, db int, maxRetries int, retryDelayMs int, zapLogger *zap.Logger) (*redislib.Client, error) {
	return ConnectWithRetry(addr, password, db, maxRetries, retryDelayMs, persistence.NewLoggerAdapter(zapLogger))
}

func Connect(addr string, password string, db int) *redislib.Client {
	client := redislib.NewClient(&redislib.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}
