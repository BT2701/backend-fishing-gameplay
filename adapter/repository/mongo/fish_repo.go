package mongo

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FishRepository struct {
	collection *mongo.Collection
}

func NewFishRepository(db *mongo.Database) *FishRepository {
	return &FishRepository{
		collection: db.Collection("fish_types"),
	}
}

func (f *FishRepository) GetTypeByID(ctx context.Context, fishID int) (*entity.FishType, error) {
	var fishType entity.FishType
	err := f.collection.FindOne(ctx, bson.M{"fish_id": fishID}).Decode(&fishType)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}
	return &fishType, nil
}
