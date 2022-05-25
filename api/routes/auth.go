package routes

import (
	"WayhaSMS/api/controllers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router){
	app.Post(os.Getenv("OAUTH_PATH"), controllers.AccessToken())
}