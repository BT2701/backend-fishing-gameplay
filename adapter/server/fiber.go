package server

import (
	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/contract"
	"github.com/gofiber/fiber/v2"
)

// FiberServer wraps *fiber.App to implement contract.Server interface
type FiberServer struct {
	app *fiber.App
}

// NewFiberServer creates a new Fiber server adapter
func NewFiberServer(app *fiber.App) contract.Server {
	return &FiberServer{
		app: app,
	}
}

func (f *FiberServer) Start() error {
	return f.app.Listen(":8080")
}

func (f *FiberServer) Stop() error {
	return f.app.Shutdown()
}

func (f *FiberServer) GetNative() interface{} {
	return f.app
}

// FiberServerFactory implements contract.ServerFactory
type FiberServerFactory struct{}

// NewFiberServerFactory creates a new Fiber server factory
func NewFiberServerFactory() contract.ServerFactory {
	return &FiberServerFactory{}
}

func (f *FiberServerFactory) CreateServer(host string, port int, logger contract.Logger) (contract.Server, error) {
	// Create a basic Fiber app
	// Note: This is a basic implementation. In real usage, you might want to configure it further
	app := fiber.New()
	return NewFiberServer(app), nil
}
