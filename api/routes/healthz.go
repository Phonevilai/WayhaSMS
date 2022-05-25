package routes

import (
	"WayhaSMS/api/controllers"

	"github.com/gofiber/fiber/v2"
)

// HealthRouter is a health check endpoint
func HealthRouter(app fiber.Router) {
	// GET /healthz
	app.Get("/healthz", controllers.HealthCheck())
}
