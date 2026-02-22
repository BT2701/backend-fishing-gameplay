package mongo

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GunRepository struct {
	collection *mongo.Collection
}

func NewGunRepository(db *mongo.Database) *GunRepository {
	return &GunRepository{
		collection: db.Collection("guns"),
	}
}

func (g *GunRepository) GetByID(ctx context.Context, gunID int) (*entity.Gun, error) {
	var gun entity.Gun
	err := g.collection.FindOne(ctx, bson.M{"gun_id": gunID}).Decode(&gun)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}
	return &gun, nil
}
