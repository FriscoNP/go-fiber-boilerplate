// Package routes registers the HTTP routes.
package routes

import (
	"github.com/FriscoNP/go-fiber-boilerplate/internal/handler"
	"github.com/gofiber/fiber/v3"
)

// Register mounts the API routes on app.
func Register(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", handler.HealthCheck)
}
