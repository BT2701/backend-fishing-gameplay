package handler

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	fiber "github.com/gofiber/fiber/v2"
)

type FishHandler struct {
	fishUsecase *usecase.FishUsecase
}

func NewFishHandler(fishUsecase *usecase.FishUsecase) *FishHandler {
	return &FishHandler{
		fishUsecase: fishUsecase,
	}
}

func (h *FishHandler) RegisterRoutes(app *fiber.App) {
	fishAPI := app.Group("/api/v1/fish")
	fishAPI.Post("/:roomID/spawn", h.SpawnFish)
}

func (h *FishHandler) SpawnFish(c *fiber.Ctx) error {
	var req struct {
		FishID  int    `json:"fish_id"`
		FishUID string `json:"fish_uid"`
		PathID  int    `json:"path_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if req.FishID <= 0 || req.FishUID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	roomID := c.Params("roomID")
	fish, err := h.fishUsecase.SpawnFish(c.Context(), roomID, req.FishID, req.FishUID, req.PathID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fish)
}
