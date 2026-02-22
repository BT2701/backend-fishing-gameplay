package http

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/delivery/http/handler"
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	roomUsecase *usecase.RoomUsecase,
	fishUsecase *usecase.FishUsecase,
	shootUsecase *usecase.ShootUsecase,
	rtpUsecase *usecase.RTPUsecase,
	skillUsecase *usecase.SkillUsecase,
	gameConfigUsecase *usecase.GameConfigUsecase,
) {
	roomHandler := handler.NewRoomHandler(roomUsecase)
	fishHandler := handler.NewFishHandler(fishUsecase)
	shootHandler := handler.NewShootHandler(shootUsecase)
	rtpHandler := handler.NewRTPHandler(rtpUsecase)
	skillHandler := handler.NewSkillHandler(skillUsecase)
	gameConfigHandler := handler.NewGameConfigHandler(gameConfigUsecase)

	roomHandler.RegisterRoutes(app)
	fishHandler.RegisterRoutes(app)
	shootHandler.RegisterRoutes(app)
	rtpHandler.RegisterRoutes(app)
	skillHandler.RegisterRoutes(app)
	gameConfigHandler.RegisterRoutes(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok"})
	})
}
