package routes

import (
	"WayhaSMS/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SMSRouter(app fiber.Router) {
	app.Post("/send-sms", controllers.SMS())
}
