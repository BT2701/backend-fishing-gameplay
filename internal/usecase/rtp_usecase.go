package usecase

import (
	"context"
	"errors"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
)

type RTPUsecase struct {
	rtpRepo port.RTPRepository
}

func NewRTPUsecase(rtpRepo port.RTPRepository) *RTPUsecase {
	return &RTPUsecase{rtpRepo: rtpRepo}
}

func (uc *RTPUsecase) GetState(ctx context.Context, roomID string) (*entity.RTPState, error) {
	state, err := uc.rtpRepo.GetByRoomID(ctx, roomID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return &entity.RTPState{}, nil
		}
		return nil, err
	}
	return state, nil
}

func (uc *RTPUsecase) Add(ctx context.Context, roomID string, totalBetDelta, totalWinDelta int64) (*entity.RTPState, error) {
	if totalBetDelta < 0 || totalWinDelta < 0 {
		return nil, apperr.ErrInvalidRTPDelta
	}

	state, err := uc.rtpRepo.GetByRoomID(ctx, roomID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			state = &entity.RTPState{}
		} else {
			return nil, err
		}
	}

	state.TotalBet += totalBetDelta
	state.TotalWin += totalWinDelta

	if err := uc.rtpRepo.Save(ctx, roomID, state); err != nil {
		return nil, err
	}

	return state, nil
}
