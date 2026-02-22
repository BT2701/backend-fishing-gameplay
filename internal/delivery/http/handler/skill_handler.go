package handler

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type SkillHandler struct {
	skillUsecase *usecase.SkillUsecase
}

func NewSkillHandler(skillUsecase *usecase.SkillUsecase) *SkillHandler {
	return &SkillHandler{
		skillUsecase: skillUsecase,
	}
}

func (h *SkillHandler) RegisterRoutes(app *fiber.App) {
	skillAPI := app.Group("/api/v1/skill")
	skillAPI.Post("/use", h.UseSkill)
}

func (h *SkillHandler) UseSkill(c *fiber.Ctx) error {
	var req struct {
		PlayerID   string `json:"player_id"`
		SkillType  string `json:"skill_type"`
		Cost       int    `json:"cost"`
		CooldownMs int    `json:"cooldown_ms"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if req.PlayerID == "" || req.SkillType == "" || req.Cost <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request: player_id, skill_type, and cost are required and cost must be positive"})
	}

	skill := &entity.Skill{
		SkillType:  req.SkillType,
		Cost:       req.Cost,
		CooldownMs: req.CooldownMs,
	}

	err := h.skillUsecase.UseSkill(c.Context(), req.PlayerID, skill)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":    "skill used successfully",
		"skill_type": req.SkillType,
	})
}
