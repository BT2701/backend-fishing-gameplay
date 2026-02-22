package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/games/game_base/models"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
	"github.com/redis/go-redis/v9"
)

type GameConfigCacheRepository struct {
	redisClient *redis.Client
	mongoRepo   port.GameConfigRepository
}

func NewGameConfigCacheRepository(redisClient *redis.Client, mongoRepo port.GameConfigRepository) *GameConfigCacheRepository {
	return &GameConfigCacheRepository{
		redisClient: redisClient,
		mongoRepo:   mongoRepo,
	}
}

// cacheKey generates Redis cache key for a config type
func (r *GameConfigCacheRepository) cacheKey(configType, gameName string) string {
	return fmt.Sprintf("game_config:%s:%s", configType, gameName)
}

// cacheTTL returns the cache TTL in seconds (24 hours)
const cacheTTL = 24 * 60 * 60

func (r *GameConfigCacheRepository) GetBulletConfig(ctx context.Context, gameName string) (*gameBaseModels.BulletConfig, error) {
	cacheKey := r.cacheKey("bullets", gameName)

	// Try to get from Redis
	cachedData, err := r.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Found in Redis, deserialize and return
		var config gameBaseModels.BulletConfig
		if err := json.Unmarshal([]byte(cachedData), &config); err != nil {
			return nil, err
		}
		return &config, nil
	}

	if !errors.Is(err, redis.Nil) {
		// Redis error occurred
		return nil, err
	}

	// Redis cache miss, try MongoDB
	config, err := r.mongoRepo.GetBulletConfig(ctx, gameName)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.New(apperr.Code("BULLET_CONFIG_NOT_FOUND"), fmt.Sprintf("bullet config not found for game: %s", gameName))
		}
		return nil, err
	}

	// Cache the data in Redis
	if data, err := json.Marshal(config); err == nil {
		_ = r.redisClient.Set(ctx, cacheKey, data, cacheTTL).Err()
	}

	return config, nil
}

func (r *GameConfigCacheRepository) GetGameConfig(ctx context.Context, gameName string) (*gameBaseModels.GameConfig, error) {
	cacheKey := r.cacheKey("configs", gameName)

	// Try to get from Redis
	cachedData, err := r.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var config gameBaseModels.GameConfig
		if err := json.Unmarshal([]byte(cachedData), &config); err != nil {
			return nil, err
		}
		return &config, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// Redis cache miss, try MongoDB
	config, err := r.mongoRepo.GetGameConfig(ctx, gameName)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.New(apperr.Code("GAME_CONFIG_NOT_FOUND"), fmt.Sprintf("game config not found for game: %s", gameName))
		}
		return nil, err
	}

	// Cache the data in Redis
	if data, err := json.Marshal(config); err == nil {
		_ = r.redisClient.Set(ctx, cacheKey, data, cacheTTL).Err()
	}

	return config, nil
}

func (r *GameConfigCacheRepository) GetGameFeatures(ctx context.Context, gameName string) (*gameBaseModels.GameFeatures, error) {
	cacheKey := r.cacheKey("features", gameName)

	// Try to get from Redis
	cachedData, err := r.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var features gameBaseModels.GameFeatures
		if err := json.Unmarshal([]byte(cachedData), &features); err != nil {
			return nil, err
		}
		return &features, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// Redis cache miss, try MongoDB
	features, err := r.mongoRepo.GetGameFeatures(ctx, gameName)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.New(apperr.Code("GAME_FEATURES_NOT_FOUND"), fmt.Sprintf("game features not found for game: %s", gameName))
		}
		return nil, err
	}

	// Cache the data in Redis
	if data, err := json.Marshal(features); err == nil {
		_ = r.redisClient.Set(ctx, cacheKey, data, cacheTTL).Err()
	}

	return features, nil
}

func (r *GameConfigCacheRepository) GetGamePaths(ctx context.Context, gameName string) (*gameBaseModels.GamePaths, error) {
	cacheKey := r.cacheKey("paths", gameName)

	// Try to get from Redis
	cachedData, err := r.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var paths gameBaseModels.GamePaths
		if err := json.Unmarshal([]byte(cachedData), &paths); err != nil {
			return nil, err
		}
		return &paths, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// Redis cache miss, try MongoDB
	paths, err := r.mongoRepo.GetGamePaths(ctx, gameName)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.New(apperr.Code("GAME_PATHS_NOT_FOUND"), fmt.Sprintf("game paths not found for game: %s", gameName))
		}
		return nil, err
	}

	// Cache the data in Redis
	if data, err := json.Marshal(paths); err == nil {
		_ = r.redisClient.Set(ctx, cacheKey, data, cacheTTL).Err()
	}

	return paths, nil
}

func (r *GameConfigCacheRepository) GetGameRTP(ctx context.Context, gameName string) (*gameBaseModels.GameRTP, error) {
	cacheKey := r.cacheKey("rtps", gameName)

	// Try to get from Redis
	cachedData, err := r.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var rtp gameBaseModels.GameRTP
		if err := json.Unmarshal([]byte(cachedData), &rtp); err != nil {
			return nil, err
		}
		return &rtp, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// Redis cache miss, try MongoDB
	rtp, err := r.mongoRepo.GetGameRTP(ctx, gameName)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.New(apperr.Code("GAME_RTP_NOT_FOUND"), fmt.Sprintf("game rtp not found for game: %s", gameName))
		}
		return nil, err
	}

	// Cache the data in Redis
	if data, err := json.Marshal(rtp); err == nil {
		_ = r.redisClient.Set(ctx, cacheKey, data, cacheTTL).Err()
	}

	return rtp, nil
}

func (r *GameConfigCacheRepository) GetGameFishTypes(ctx context.Context, gameName string) (*gameBaseModels.GameFishTypes, error) {
	cacheKey := r.cacheKey("types", gameName)

	// Try to get from Redis
	cachedData, err := r.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var fishTypes gameBaseModels.GameFishTypes
		if err := json.Unmarshal([]byte(cachedData), &fishTypes); err != nil {
			return nil, err
		}
		return &fishTypes, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// Redis cache miss, try MongoDB
	fishTypes, err := r.mongoRepo.GetGameFishTypes(ctx, gameName)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.New(apperr.Code("GAME_FISH_TYPES_NOT_FOUND"), fmt.Sprintf("game fish types not found for game: %s", gameName))
		}
		return nil, err
	}

	// Cache the data in Redis
	if data, err := json.Marshal(fishTypes); err == nil {
		_ = r.redisClient.Set(ctx, cacheKey, data, cacheTTL).Err()
	}

	return fishTypes, nil
}
