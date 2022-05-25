package controllers

import (
	"WayhaSMS/api/presenter"
	"WayhaSMS/pkg/entities"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Access Tokenz get user and password
func AccessToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var auth entities.Auth
		var Failed bool = false
		err := c.BodyParser(&auth)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse("Unauthorization"))
		}
		identity := auth.Username
		newPass := auth.Password
		if auth.Username != os.Getenv("my_username") {
			return c.Status(fiber.StatusNotFound).JSON(presenter.AuthErrorResponse("Username Is Not Found"))
		}
		if matchPass := CheckPasswordHash(newPass, os.Getenv("my_password")); matchPass == Failed {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse("Password Incorrect"))
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims[os.Getenv("my_username")] = identity
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

		t, sigerr := token.SignedString([]byte(os.Getenv("signature")))
		if sigerr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrorResponse("Server Error"))
		}
		return c.Status(fiber.StatusOK).JSON(presenter.AuthSuccessResponse(&auth, t))
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}