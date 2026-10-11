// Package server builds and starts the Fiber application.
package server

import (
	"fmt"

	"github.com/FriscoNP/go-fiber-boilerplate/internal/middleware"
	"github.com/FriscoNP/go-fiber-boilerplate/internal/routes"
	"github.com/gofiber/fiber/v3"
)

// New creates a Fiber app with middleware and routes registered.
func New() *fiber.App {
	app := fiber.New()

	// middleware
	middleware.Register(app)

	// routes
	routes.Register(app)

	return app
}

// Start listens on the given port.
func Start(app *fiber.App, port string) error {
	return app.Listen(fmt.Sprintf(":%s", port))
}
