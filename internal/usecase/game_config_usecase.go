package usecase

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/games/game_base/models"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
)

type GameConfigUsecase struct {
	gameConfigRepo port.GameConfigRepository
}

func NewGameConfigUsecase(gameConfigRepo port.GameConfigRepository) *GameConfigUsecase {
	return &GameConfigUsecase{
		gameConfigRepo: gameConfigRepo,
	}
}

// GetBulletConfig retrieves bullet configuration for a game
func (uc *GameConfigUsecase) GetBulletConfig(ctx context.Context, gameName string) (*gameBaseModels.BulletConfig, error) {
	return uc.gameConfigRepo.GetBulletConfig(ctx, gameName)
}

// GetGameConfig retrieves game betting and general configuration
func (uc *GameConfigUsecase) GetGameConfig(ctx context.Context, gameName string) (*gameBaseModels.GameConfig, error) {
	return uc.gameConfigRepo.GetGameConfig(ctx, gameName)
}

// GetGameFeatures retrieves game-specific features
func (uc *GameConfigUsecase) GetGameFeatures(ctx context.Context, gameName string) (*gameBaseModels.GameFeatures, error) {
	return uc.gameConfigRepo.GetGameFeatures(ctx, gameName)
}

// GetGamePaths retrieves fish paths for a game
func (uc *GameConfigUsecase) GetGamePaths(ctx context.Context, gameName string) (*gameBaseModels.GamePaths, error) {
	return uc.gameConfigRepo.GetGamePaths(ctx, gameName)
}

// GetGameRTP retrieves RTP configuration for a game
func (uc *GameConfigUsecase) GetGameRTP(ctx context.Context, gameName string) (*gameBaseModels.GameRTP, error) {
	return uc.gameConfigRepo.GetGameRTP(ctx, gameName)
}

// GetGameFishTypes retrieves all fish types for a game
func (uc *GameConfigUsecase) GetGameFishTypes(ctx context.Context, gameName string) (*gameBaseModels.GameFishTypes, error) {
	return uc.gameConfigRepo.GetGameFishTypes(ctx, gameName)
}
