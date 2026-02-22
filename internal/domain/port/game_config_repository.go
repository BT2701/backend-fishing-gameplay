package port

import (
	"context"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/games/game_base/models"
)

// GameConfigRepository defines methods to access game configuration data
type GameConfigRepository interface {
	// GetBulletConfig retrieves bullet configuration for a game
	GetBulletConfig(ctx context.Context, gameName string) (*gameBaseModels.BulletConfig, error)

	// GetGameConfig retrieves game betting and general configuration
	GetGameConfig(ctx context.Context, gameName string) (*gameBaseModels.GameConfig, error)

	// GetGameFeatures retrieves game-specific features (skills, rewards, multipliers)
	GetGameFeatures(ctx context.Context, gameName string) (*gameBaseModels.GameFeatures, error)

	// GetGamePaths retrieves fish paths for a game
	GetGamePaths(ctx context.Context, gameName string) (*gameBaseModels.GamePaths, error)

	// GetGameRTP retrieves RTP configuration for a game
	GetGameRTP(ctx context.Context, gameName string) (*gameBaseModels.GameRTP, error)

	// GetGameFishTypes retrieves all fish types for a game
	GetGameFishTypes(ctx context.Context, gameName string) (*gameBaseModels.GameFishTypes, error)
}
