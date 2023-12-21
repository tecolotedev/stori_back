package controllers

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"ok": true})
}
