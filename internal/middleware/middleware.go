// Package middleware registers the global Fiber middleware.
package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

// Register attaches recover and logger middleware to app.
func Register(app *fiber.App) {
	app.Use(recoverer.New())

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))
}
