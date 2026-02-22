package handler

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type GameConfigHandler struct {
	gameConfigUsecase *usecase.GameConfigUsecase
}

func NewGameConfigHandler(gameConfigUsecase *usecase.GameConfigUsecase) *GameConfigHandler {
	return &GameConfigHandler{
		gameConfigUsecase: gameConfigUsecase,
	}
}

func (h *GameConfigHandler) RegisterRoutes(app *fiber.App) {
	gameConfigAPI := app.Group("/api/v1/game-config")
	gameConfigAPI.Get("/:gameName/bullets", h.GetBulletConfig)
	gameConfigAPI.Get("/:gameName/config", h.GetGameConfig)
	gameConfigAPI.Get("/:gameName/features", h.GetGameFeatures)
	gameConfigAPI.Get("/:gameName/paths", h.GetGamePaths)
	gameConfigAPI.Get("/:gameName/rtp", h.GetGameRTP)
	gameConfigAPI.Get("/:gameName/fish-types", h.GetGameFishTypes)
}

func (h *GameConfigHandler) GetBulletConfig(c *fiber.Ctx) error {
	gameName := c.Params("gameName")
	if gameName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "game_name is required"})
	}

	config, err := h.gameConfigUsecase.GetBulletConfig(c.Context(), gameName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(config)
}

func (h *GameConfigHandler) GetGameConfig(c *fiber.Ctx) error {
	gameName := c.Params("gameName")
	if gameName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "game_name is required"})
	}

	config, err := h.gameConfigUsecase.GetGameConfig(c.Context(), gameName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(config)
}

func (h *GameConfigHandler) GetGameFeatures(c *fiber.Ctx) error {
	gameName := c.Params("gameName")
	if gameName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "game_name is required"})
	}

	features, err := h.gameConfigUsecase.GetGameFeatures(c.Context(), gameName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(features)
}

func (h *GameConfigHandler) GetGamePaths(c *fiber.Ctx) error {
	gameName := c.Params("gameName")
	if gameName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "game_name is required"})
	}

	paths, err := h.gameConfigUsecase.GetGamePaths(c.Context(), gameName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(paths)
}

func (h *GameConfigHandler) GetGameRTP(c *fiber.Ctx) error {
	gameName := c.Params("gameName")
	if gameName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "game_name is required"})
	}

	rtp, err := h.gameConfigUsecase.GetGameRTP(c.Context(), gameName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(rtp)
}

func (h *GameConfigHandler) GetGameFishTypes(c *fiber.Ctx) error {
	gameName := c.Params("gameName")
	if gameName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "game_name is required"})
	}

	fishTypes, err := h.gameConfigUsecase.GetGameFishTypes(c.Context(), gameName)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fishTypes)
}
