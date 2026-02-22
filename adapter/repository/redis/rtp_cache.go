package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
	"github.com/redis/go-redis/v9"
)

type RTPRepository struct {
	client *redis.Client
}

func NewRTPRepository(client *redis.Client) *RTPRepository {
	return &RTPRepository{
		client: client,
	}
}

func (r *RTPRepository) GetByRoomID(ctx context.Context, roomID string) (*entity.RTPState, error) {
	key := fmt.Sprintf("rtp:%s", roomID)
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	var state entity.RTPState
	if err := json.Unmarshal([]byte(val), &state); err != nil {
		return nil, err
	}
	return &state, nil
}

func (r *RTPRepository) Save(ctx context.Context, roomID string, state *entity.RTPState) error {
	key := fmt.Sprintf("rtp:%s", roomID)
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, string(data), 0).Err()
}
