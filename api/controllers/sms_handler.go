package controllers

import (
	"WayhaSMS/api/presenter"
	"WayhaSMS/pkg/entities"
	"WayhaSMS/pkg/ltc"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SMS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqData entities.SMS
		if err := c.BodyParser(&reqData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.SendSMSErrorResponse("REQUEST BODY IS INVALID"))
		}
		q := entities.SMSReq{
			PrivateKey: os.Getenv("privateKey"),
			UserID:     os.Getenv("userid"),
			Trans_ID:   reqData.Trans_ID,
			MsisDN:     reqData.MsisDN,
			HeaderSMS:  os.Getenv("headerSMS"),
			Message:    reqData.Message,
		}
		resultSMS := ltc.SendSMS(&q)
		if resultSMS != "20" {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.SendSMSErrorResponse("Failed"))
		}
		return c.Status(fiber.StatusOK).JSON(presenter.SendSMSSuccessResponse())
	}
}
