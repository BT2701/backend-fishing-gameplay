package handler

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	fiber "github.com/gofiber/fiber/v2"
)

type RoomHandler struct {
	roomUsecase *usecase.RoomUsecase
}

func NewRoomHandler(roomUsecase *usecase.RoomUsecase) *RoomHandler {
	return &RoomHandler{
		roomUsecase: roomUsecase,
	}
}

func (h *RoomHandler) RegisterRoutes(app *fiber.App) {
	roomsAPI := app.Group("/api/v1/rooms")
	roomsAPI.Post("", h.CreateRoom)
	roomsAPI.Post("/:roomID/join", h.JoinRoom)
	roomsAPI.Post("/:roomID/leave", h.LeaveRoom)
	roomsAPI.Get("/:roomID", h.GetRoom)
}

func (h *RoomHandler) CreateRoom(c *fiber.Ctx) error {
	var req struct {
		RoomID     string `json:"room_id"`
		MaxPlayers int    `json:"max_players"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if req.RoomID == "" || req.MaxPlayers <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	room, err := h.roomUsecase.CreateRoom(c.Context(), req.RoomID, req.MaxPlayers)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(room)
}

func (h *RoomHandler) JoinRoom(c *fiber.Ctx) error {
	var req struct {
		PlayerID       string `json:"player_id"`
		SeatID         int    `json:"seat_id"`
		InitialBalance int64  `json:"initial_balance"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	roomID := c.Params("roomID")
	room, player, err := h.roomUsecase.JoinRoom(c.Context(), roomID, req.PlayerID, req.SeatID, req.InitialBalance)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"room": room, "player": player})
}

func (h *RoomHandler) LeaveRoom(c *fiber.Ctx) error {
	var req struct {
		PlayerID string `json:"player_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	roomID := c.Params("roomID")
	room, player, err := h.roomUsecase.LeaveRoom(c.Context(), roomID, req.PlayerID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"room": room, "player": player})
}

func (h *RoomHandler) GetRoom(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "not implemented yet"})
}
