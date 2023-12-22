package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/utils"
)

func Auth(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	payload, err := utils.VerifyToken(accessToken)
	if err != nil {
		c.ClearCookie("access_token")
		return err
	}
	c.Locals("userId", payload.USERID)
	return c.Next()
}
