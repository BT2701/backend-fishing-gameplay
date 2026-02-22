package mongo

import (
	"context"
	"errors"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/games/game_base/models"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameConfigRepository struct {
	db *mongo.Database
}

func NewGameConfigRepository(db *mongo.Database) port.GameConfigRepository {
	return &GameConfigRepository{
		db: db,
	}
}

func (r *GameConfigRepository) GetBulletConfig(ctx context.Context, gameName string) (*gameBaseModels.BulletConfig, error) {
	collection := r.db.Collection("bullets")
	var config gameBaseModels.BulletConfig

	err := collection.FindOne(ctx, bson.M{"game_name": gameName}).Decode(&config)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return &config, nil
}

func (r *GameConfigRepository) GetGameConfig(ctx context.Context, gameName string) (*gameBaseModels.GameConfig, error) {
	collection := r.db.Collection("configs")
	var config gameBaseModels.GameConfig

	err := collection.FindOne(ctx, bson.M{"game_name": gameName}).Decode(&config)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return &config, nil
}

func (r *GameConfigRepository) GetGameFeatures(ctx context.Context, gameName string) (*gameBaseModels.GameFeatures, error) {
	collection := r.db.Collection("features")
	var features gameBaseModels.GameFeatures

	err := collection.FindOne(ctx, bson.M{"game_name": gameName}).Decode(&features)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return &features, nil
}

func (r *GameConfigRepository) GetGamePaths(ctx context.Context, gameName string) (*gameBaseModels.GamePaths, error) {
	collection := r.db.Collection("paths")
	var paths gameBaseModels.GamePaths

	err := collection.FindOne(ctx, bson.M{"game_name": gameName}).Decode(&paths)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return &paths, nil
}

func (r *GameConfigRepository) GetGameRTP(ctx context.Context, gameName string) (*gameBaseModels.GameRTP, error) {
	collection := r.db.Collection("rtps")
	var rtp gameBaseModels.GameRTP

	err := collection.FindOne(ctx, bson.M{"game_name": gameName}).Decode(&rtp)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return &rtp, nil
}

func (r *GameConfigRepository) GetGameFishTypes(ctx context.Context, gameName string) (*gameBaseModels.GameFishTypes, error) {
	collection := r.db.Collection("types")
	var fishTypes gameBaseModels.GameFishTypes

	err := collection.FindOne(ctx, bson.M{"game_name": gameName}).Decode(&fishTypes)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return &fishTypes, nil
}
