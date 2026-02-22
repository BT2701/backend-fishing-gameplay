package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
)

type SkillUsecase struct {
	playerRepo port.PlayerRepository
	now        func() time.Time
}

func NewSkillUsecase(playerRepo port.PlayerRepository) *SkillUsecase {
	return &SkillUsecase{
		playerRepo: playerRepo,
		now:        time.Now,
	}
}

func (uc *SkillUsecase) UseSkill(ctx context.Context, playerID string, skill *entity.Skill) error {
	if playerID == "" {
		return apperr.ErrInvalidPlayerID
	}

	ok, err := skill.IsValid()
	if !ok || err != nil {
		return err
	}

	player, err := uc.playerRepo.GetByID(ctx, playerID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			return apperr.ErrPlayerNotFound
		}
		return err
	}

	if !player.CanSpend(int64(skill.Cost)) {
		return apperr.ErrInsufficientBalance
	}

	if err := player.Spend(int64(skill.Cost)); err != nil {
		return err
	}

	player.LastActionAt = uc.now().Unix()

	if err := uc.playerRepo.Save(ctx, player); err != nil {
		return err
	}

	return nil
}
