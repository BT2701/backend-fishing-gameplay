package mongo

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerRepository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(db *mongo.Database) *PlayerRepository {
	return &PlayerRepository{
		collection: db.Collection("players"),
	}
}

func (p *PlayerRepository) GetByID(ctx context.Context, playerID string) (*entity.Player, error) {
	var player entity.Player
	err := p.collection.FindOne(ctx, bson.M{"player_id": playerID}).Decode(&player)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}
	return &player, nil
}

func (p *PlayerRepository) Save(ctx context.Context, player *entity.Player) error {
	opts := options.Update().SetUpsert(true)
	_, err := p.collection.UpdateOne(
		ctx,
		bson.M{"player_id": player.PlayerID},
		bson.M{"$set": player},
		opts,
	)
	return err
}
