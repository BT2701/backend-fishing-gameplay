package mongo

import (
	"context"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/persistence"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func ConnectWithRetry(uri string, database string, timeout int, maxRetries int, retryDelayMs int, logger persistence.Logger) (*mongo.Client, error) {
	var client *mongo.Client

	err := persistence.RetryWithBackoff(maxRetries, retryDelayMs, logger, "connect to MongoDB", func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()

		c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}

		// Verify connection with ping
		if err := c.Ping(ctx, nil); err != nil {
			c.Disconnect(context.Background())
			return err
		}

		client = c
		return nil
	})

	return client, err
}

// ConnectWithRetryZap is a helper function that wraps zap.Logger for backwards compatibility
func ConnectWithRetryZap(uri string, database string, timeout int, maxRetries int, retryDelayMs int, zapLogger *zap.Logger) (*mongo.Client, error) {
	return ConnectWithRetry(uri, database, timeout, maxRetries, retryDelayMs, persistence.NewLoggerAdapter(zapLogger))
}

func Connect(uri string, database string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Close(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}
