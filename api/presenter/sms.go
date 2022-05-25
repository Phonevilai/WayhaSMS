package presenter

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
)

// Send SMS SuccessResponse
func SendSMSSuccessResponse() *fiber.Map {
	t := time.Now()
	return &fiber.Map{
		"timestamp": fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()),
		"status": "Success",
	}
}
// Send SMS Failed
func SendSMSErrorResponse(errMsg string) *fiber.Map {
	t := time.Now()
	return &fiber.Map{
		"timestamp": fmt.Sprintf("%d%02d%02d%02d%02d%02d",
			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()),
		"status": "Failed",
		"error":  errMsg,
	}
}
