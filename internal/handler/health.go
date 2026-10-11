// Package handler contains the HTTP handlers.
package handler

import "github.com/gofiber/fiber/v3"

// HealthCheck reports that the service is up.
func HealthCheck(ctx fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"message": "Fiber API is running smoothly!",
	})
}
