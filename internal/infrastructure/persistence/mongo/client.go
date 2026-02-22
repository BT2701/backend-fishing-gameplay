package mongo

import (
	"context"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/persistence"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func ConnectWithRetry(uri string, database string, timeout int, maxRetries int, retryDelayMs int, logger *zap.Logger) (*mongo.Client, error) {
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
