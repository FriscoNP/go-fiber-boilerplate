package routes

import (
	"github.com/FriscoNP/go-fiber-boilerplate/internal/handler"
	"github.com/gofiber/fiber/v3"
)

func Register(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", handler.HealthCheck)
}
