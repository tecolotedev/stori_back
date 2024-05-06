package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/models"
)

func GetNewsletter(c *fiber.Ctx) error {
	return c.SendString("Hello, from newsletter controller")
}

type CreateNewsLetter struct {
	Name string `json:"name" form:"name"`
}

func CreateNewsletter(c *fiber.Ctx) error {
	nl := new(CreateNewsLetter)

	if err := c.BodyParser(nl); err != nil {
		return err
	}

	newsletter := models.Newsletter{
		Name:               nl.Name,
		NewsletterVersions: []models.NewsletterVersion{},
	}

	models.DB.Create(&newsletter)

	return c.JSON(newsletter)
}
