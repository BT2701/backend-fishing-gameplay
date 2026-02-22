package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
)

type ShootUsecase struct {
	roomRepo   port.RoomRepository
	playerRepo port.PlayerRepository
	fishRepo   port.FishRepository
	gunRepo    port.GunRepository
	rtpRepo    port.RTPRepository
	now        func() time.Time
}

func NewShootUsecase(roomRepo port.RoomRepository, playerRepo port.PlayerRepository, fishRepo port.FishRepository, gunRepo port.GunRepository, rtpRepo port.RTPRepository) *ShootUsecase {
	return &ShootUsecase{
		roomRepo:   roomRepo,
		playerRepo: playerRepo,
		fishRepo:   fishRepo,
		gunRepo:    gunRepo,
		rtpRepo:    rtpRepo,
		now:        time.Now,
	}
}

func (uc *ShootUsecase) Fire(ctx context.Context, roomID, playerID, fishUID string) (*entity.Shot, *entity.FishInstance, *entity.Player, error) {
	if roomID == "" {
		return nil, nil, nil, apperr.ErrInvalidRoomID
	}
	if playerID == "" {
		return nil, nil, nil, apperr.ErrInvalidPlayerID
	}
	if fishUID == "" {
		return nil, nil, nil, apperr.ErrInvalidFishUID
	}

	room, err := uc.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, nil, nil, apperr.ErrRoomNotFound
		}
		return nil, nil, nil, err
	}

	if room.FishMap == nil {
		return nil, nil, nil, apperr.ErrFishNotFound
	}

	fish, ok := room.FishMap[fishUID]
	if !ok {
		return nil, nil, nil, apperr.ErrFishNotFound
	}
	if !fish.Alive || fish.HP <= 0 {
		return nil, nil, nil, apperr.ErrFishAlreadyDead
	}

	player, ok := room.Players[playerID]
	if !ok {
		p, err := uc.playerRepo.GetByID(ctx, playerID)
		if err != nil {
			if errors.Is(err, apperr.ErrNotFound) {
				return nil, nil, nil, apperr.ErrPlayerNotFound
			}
			return nil, nil, nil, err
		}
		player = p
	}

	if player.RoomID != roomID {
		return nil, nil, nil, apperr.ErrPlayerNotInRoom
	}
	if player.Balance < 0 {
		return nil, nil, nil, apperr.ErrInvalidBalance
	}

	gun, err := uc.gunRepo.GetByID(ctx, player.GunID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return nil, nil, nil, apperr.ErrGunNotFound
		}
		return nil, nil, nil, err
	}

	if player.Balance < int64(gun.BulletCost) {
		return nil, nil, nil, apperr.ErrInsufficientBalance
	}

	player.Balance -= int64(gun.BulletCost)

	reward := int64(0)
	fish.HP -= gun.Damage
	if fish.HP <= 0 {
		fish.Alive = false
		fishType, err := uc.fishRepo.GetTypeByID(ctx, fish.FishID)
		if err != nil {
			if errors.Is(err, apperr.ErrNotFound) {
				return nil, nil, nil, apperr.ErrFishTypeNotFound
			}
			return nil, nil, nil, err
		}
		reward = int64(fishType.Reward)
		player.Balance += reward
	}

	if err := uc.roomRepo.Save(ctx, room); err != nil {
		return nil, nil, nil, err
	}
	if err := uc.playerRepo.Save(ctx, player); err != nil {
		return nil, nil, nil, err
	}

	if uc.rtpRepo != nil {
		state, err := uc.rtpRepo.GetByRoomID(ctx, roomID)
		if err != nil && !errors.Is(err, apperr.ErrNotFound) {
			return nil, nil, nil, err
		}
		if state == nil {
			state = &entity.RTPState{}
		}
		state.TotalBet += int64(gun.BulletCost)
		state.TotalWin += reward
		if err := uc.rtpRepo.Save(ctx, roomID, state); err != nil {
			return nil, nil, nil, err
		}
	}

	shot := &entity.Shot{
		BulletID: fmt.Sprintf("%s-%d", playerID, uc.now().UnixNano()),
		PlayerID: playerID,
		FishUID:  fishUID,
		GunID:    gun.GunID,
		FireTime: uc.now().Unix(),
	}

	return shot, fish, player, nil
}
