package port

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
)

type RTPRepository interface {
	GetByRoomID(ctx context.Context, roomID string) (*entity.RTPState, error)
	Save(ctx context.Context, roomID string, state *entity.RTPState) error
}
