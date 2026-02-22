package database

import (
	"context"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/contract"
	infmongo "github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/persistence/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDatabase wraps *mongo.Client to implement contract.Database interface
type MongoDatabase struct {
	client *mongo.Client
}

// NewMongoDatabase creates a new MongoDB adapter
func NewMongoDatabase(client *mongo.Client) contract.Database {
	return &MongoDatabase{
		client: client,
	}
}

func (m *MongoDatabase) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, nil)
}

func (m *MongoDatabase) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoDatabase) GetNative() interface{} {
	return m.client
}

// MongoDatabaseFactory implements contract.DatabaseFactory
type MongoDatabaseFactory struct{}

// NewMongoDatabaseFactory creates a new MongoDB factory
func NewMongoDatabaseFactory() contract.DatabaseFactory {
	return &MongoDatabaseFactory{}
}

func (f *MongoDatabaseFactory) CreateDatabase(cfg contract.Config, logger contract.Logger) (contract.Database, error) {
	client, err := infmongo.ConnectWithRetry(
		cfg.GetMongoURI(),
		cfg.GetMongoDatabase(),
		cfg.GetMongoTimeout(),
		cfg.GetMongoMaxRetries(),
		cfg.GetMongoRetryDelay(),
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewMongoDatabase(client), nil
}

func (f *MongoDatabaseFactory) CloseDatabase(db contract.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.Close(ctx)
}
