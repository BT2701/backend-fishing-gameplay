package mongo

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomRepository struct {
	collection *mongo.Collection
}

func NewRoomRepository(db *mongo.Database) *RoomRepository {
	return &RoomRepository{
		collection: db.Collection("rooms"),
	}
}

func (r *RoomRepository) GetByID(ctx context.Context, roomID string) (*entity.Room, error) {
	var room entity.Room
	err := r.collection.FindOne(ctx, bson.M{"room_id": roomID}).Decode(&room)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) Save(ctx context.Context, room *entity.Room) error {
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"room_id": room.RoomID},
		bson.M{"$set": room},
		opts,
	)
	return err
}
