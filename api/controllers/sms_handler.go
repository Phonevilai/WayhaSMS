package controllers

import (
	"WayhaSMS/api/presenter"
	"WayhaSMS/pkg/entities"
	"WayhaSMS/pkg/ltc"
	"github.com/gofiber/fiber/v2"
	"os"
)

func SMS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqData entities.SMS
		if err := c.BodyParser(&reqData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.SendSMSErrorResponse("REQUEST BODY IS INVALID"))
		}
		q := entities.SMSReq{
			PrivateKey: os.Getenv("PRIVATE_KEY"),
			UserID:     os.Getenv("USER_ID"),
			Trans_ID:   reqData.Trans_ID,
			MsisDN:     reqData.MsisDN,
			HeaderSMS:  os.Getenv("HEADER_SMS"),
			Message:    reqData.Message,
		}
		resultSMS := ltc.SendSMS(&q)
		if resultSMS != "20" {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.SendSMSErrorResponse("Failed"))
		}
		return c.Status(fiber.StatusOK).JSON(presenter.SendSMSSuccessResponse())
	}
}
