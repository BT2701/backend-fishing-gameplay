package port

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
)

type GunRepository interface {
	GetByID(ctx context.Context, gunID int) (*entity.Gun, error)
}
