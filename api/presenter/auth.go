package presenter

import (
	"WayhaSMS/pkg/entities"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Auth is the presenter object which will be passed in the response by Handler
type Auth struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// AuthSuccessResponse is the singular SuccessResponse that will be passed in the response by
////Handler
func AuthSuccessResponse(data *entities.Auth, token string) *fiber.Map {
	t := time.Now()
	d := Auth{
		Username: data.Username,
		Token:    token,
	}
	return &fiber.Map{
		"timestamp": fmt.Sprintf("%d%02d%02d%02d%02d%02d",
			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()),
		"status": "Success",
		"data":   d,
		"error":  nil,
	}
}

// AuthErrorResponse is the ErrorResponse that will be passed in the response by Handler
func AuthErrorResponse(errMsg string) *fiber.Map {
	t := time.Now()
	return &fiber.Map{
		"timestamp": fmt.Sprintf("%d%02d%02d%02d%02d%02d",
			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()),
		"status": "Failed",
		"error":  errMsg,
	}
}