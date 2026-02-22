package server

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	host   string
	port   int
	logger *zap.Logger
}

func New(host string, port int, logger *zap.Logger) *Server {
	app := fiber.New()
	return &Server{
		app:    app,
		host:   host,
		port:   port,
		logger: logger,
	}
}

func (s *Server) GetApp() *fiber.App {
	return s.app
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.logger.Info("Starting server", zap.String("addr", addr))
	return s.app.Listen(addr)
}

func (s *Server) Stop() error {
	return s.app.Shutdown()
}
