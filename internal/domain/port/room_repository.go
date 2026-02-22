package port

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
)

type RoomRepository interface {
	GetByID(ctx context.Context, roomID string) (*entity.Room, error)
	Save(ctx context.Context, room *entity.Room) error
}
