package handler

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type RTPHandler struct {
	rtpUsecase *usecase.RTPUsecase
}

func NewRTPHandler(rtpUsecase *usecase.RTPUsecase) *RTPHandler {
	return &RTPHandler{
		rtpUsecase: rtpUsecase,
	}
}

func (h *RTPHandler) RegisterRoutes(app *fiber.App) {
	rtpAPI := app.Group("/api/v1/rtp")
	rtpAPI.Get("/:roomID", h.GetRTPState)
	rtpAPI.Post("/:roomID/update", h.UpdateRTPState)
}

func (h *RTPHandler) GetRTPState(c *fiber.Ctx) error {
	roomID := c.Params("roomID")
	if roomID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "room_id is required"})
	}

	state, err := h.rtpUsecase.GetState(c.Context(), roomID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(state)
}

func (h *RTPHandler) UpdateRTPState(c *fiber.Ctx) error {
	var req struct {
		TotalBetDelta int64 `json:"total_bet_delta"`
		TotalWinDelta int64 `json:"total_win_delta"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	roomID := c.Params("roomID")
	if roomID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "room_id is required"})
	}

	if req.TotalBetDelta < 0 || req.TotalWinDelta < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "total_bet_delta and total_win_delta must be non-negative"})
	}

	state, err := h.rtpUsecase.Add(c.Context(), roomID, req.TotalBetDelta, req.TotalWinDelta)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(state)
}
