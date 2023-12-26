package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func SendError(c *fiber.Ctx, message string, status int) error {

	c.Status(status)

	res := ErrorResponse{Ok: false, Message: message}

	return c.JSON(res)
}
