package port

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
)

type PlayerRepository interface {
	GetByID(ctx context.Context, playerID string) (*entity.Player, error)
	Save(ctx context.Context, player *entity.Player) error
}
