package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/models"
)

func GetAllNewsletters(c *fiber.Ctx) error {
	var newsletters []models.Newsletter

	result := models.DB.Find(&newsletters)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(500).JSON(fiber.Map{"ok": false, "message": "something went wrong, please try it later"})
	}

	return c.JSON(newsletters)
}

type NewsletterRequest struct {
	Name string `json:"name" form:"name"`
}

func CreateNewsletter(c *fiber.Ctx) error {
	nlr := new(NewsletterRequest)

	if err := c.BodyParser(nlr); err != nil {
		return err
	}

	newsletter := models.Newsletter{
		Name:               nlr.Name,
		NewsletterVersions: []models.NewsletterVersion{},
	}

	models.DB.Create(&newsletter)

	return c.JSON(newsletter)
}
