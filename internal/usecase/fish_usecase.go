package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
)

type FishUsecase struct {
	roomRepo port.RoomRepository
	fishRepo port.FishRepository
	now      func() time.Time
}

func NewFishUsecase(roomRepo port.RoomRepository, fishRepo port.FishRepository) *FishUsecase {
	return &FishUsecase{
		roomRepo: roomRepo,
		fishRepo: fishRepo,
		now:      time.Now,
	}
}

func (uc *FishUsecase) SpawnFish(ctx context.Context, roomID string, fishID int, fishUID string, pathID int) (*entity.FishInstance, error) {
	if fishID <= 0 {
		return nil, apperr.ErrInvalidFishID
	}
	if fishUID == "" {
		return nil, apperr.ErrInvalidFishUID
	}

	room, err := uc.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.ErrRoomNotFound
		}
		return nil, err
	}

	if room.FishMap == nil {
		room.FishMap = map[string]*entity.FishInstance{}
	}
	if _, exists := room.FishMap[fishUID]; exists {
		return nil, apperr.ErrFishUIDExists
	}

	fishType, err := uc.fishRepo.GetTypeByID(ctx, fishID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, apperr.ErrFishTypeNotFound
		}
		return nil, err
	}

	instance := &entity.FishInstance{
		FishUID:   fishUID,
		FishID:    fishID,
		HP:        fishType.BaseHP,
		SpawnTime: uc.now().Unix(),
		PathID:    pathID,
		Alive:     true,
	}

	room.FishMap[fishUID] = instance

	if err := uc.roomRepo.Save(ctx, room); err != nil {
		return nil, err
	}

	return instance, nil
}
