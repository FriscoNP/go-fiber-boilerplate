package server

import (
	"fmt"

	"github.com/FriscoNP/go-fiber-boilerplate/internal/middleware"
	"github.com/FriscoNP/go-fiber-boilerplate/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func New() *fiber.App {
	app := fiber.New()

	// middleware
	middleware.Register(app)

	// routes
	routes.Register(app)

	return app
}

func Start(app *fiber.App, port string) error {
	return app.Listen(fmt.Sprintf(":%s", port))
}
