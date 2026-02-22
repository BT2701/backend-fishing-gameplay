package database

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/contract"
	infredis "github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/persistence/redis"
	"github.com/redis/go-redis/v9"
)

// RedisCache wraps *redis.Client to implement contract.Cache interface
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new Redis cache adapter
func NewRedisCache(client *redis.Client) contract.Cache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}

func (r *RedisCache) GetNative() interface{} {
	return r.client
}

// RedisCacheFactory implements contract.CacheFactory
type RedisCacheFactory struct{}

// NewRedisCacheFactory creates a new Redis cache factory
func NewRedisCacheFactory() contract.CacheFactory {
	return &RedisCacheFactory{}
}

func (f *RedisCacheFactory) CreateCache(cfg contract.Config, logger contract.Logger) (contract.Cache, error) {
	client, err := infredis.ConnectWithRetry(
		cfg.GetRedisAddr(),
		cfg.GetRedisPassword(),
		cfg.GetRedisDB(),
		cfg.GetRedisMaxRetries(),
		cfg.GetRedisRetryDelay(),
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewRedisCache(client), nil
}

func (f *RedisCacheFactory) CloseCache(cache contract.Cache) error {
	return cache.Close()
}
