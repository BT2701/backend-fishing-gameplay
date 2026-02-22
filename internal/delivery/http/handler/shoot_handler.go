package handler

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ShootHandler struct {
	shootUsecase *usecase.ShootUsecase
}

func NewShootHandler(shootUsecase *usecase.ShootUsecase) *ShootHandler {
	return &ShootHandler{
		shootUsecase: shootUsecase,
	}
}

func (h *ShootHandler) RegisterRoutes(app *fiber.App) {
	shootAPI := app.Group("/api/v1/shoot")
	shootAPI.Post("/fire", h.Fire)
}

func (h *ShootHandler) Fire(c *fiber.Ctx) error {
	var req struct {
		RoomID   string `json:"room_id"`
		PlayerID string `json:"player_id"`
		FishUID  string `json:"fish_uid"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if req.RoomID == "" || req.PlayerID == "" || req.FishUID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request: room_id, player_id, and fish_uid are required"})
	}

	shot, fish, player, err := h.shootUsecase.Fire(c.Context(), req.RoomID, req.PlayerID, req.FishUID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"shot":   shot,
		"fish":   fish,
		"player": player,
	})
}
