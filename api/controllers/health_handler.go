package controllers

import "github.com/gofiber/fiber/v2"

// HealthHandler is a controller that handles health checks
func HealthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	}
}
