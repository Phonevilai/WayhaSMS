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
		err := c.BodyParser(&auth)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse("Unauthorization"))
		}
		identity := auth.Username
		newPass := auth.Password
		if auth.Username != os.Getenv("USERNAME") {
			return c.Status(fiber.StatusNotFound).JSON(presenter.AuthErrorResponse("Username Is Not Found"))
		}
		//fmt.Println(newPass)
		match := CheckPasswordHash(os.Getenv("PASSWORD"), newPass)
		if match == false {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrorResponse("Password Incorrect"))
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims[os.Getenv("USERNAME")] = identity
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

		t, sigerr := token.SignedString([]byte(os.Getenv("SIGNATURE")))
		if sigerr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrorResponse("Server Error"))
		}
		return c.Status(fiber.StatusOK).JSON(presenter.AuthSuccessResponse(&auth, t))
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
