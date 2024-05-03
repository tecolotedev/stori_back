package controllers

import "github.com/gofiber/fiber/v2"

func GetNewsletter(c *fiber.Ctx) error {
	return c.SendString("Hello, from newsletter controller")
}
