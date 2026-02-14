package handler

import "github.com/gofiber/fiber/v3"

func HealthCheck(ctx fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"message": "Fiber API is running smoothly!",
	})
}
