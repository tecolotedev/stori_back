package controllers

import "github.com/gofiber/fiber/v2"

func GetRecipients(c *fiber.Ctx) error {
	return c.SendString("Hello, from recipient controller")
}
