package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/models"
)

func GetRecipients(c *fiber.Ctx) error {
	return c.SendString("Hello, from recipient controller")
}

type RecipientRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RecipientCreateRequest struct {
	Recipients    []RecipientRequest `json:"recipients"`
	NewsletterIDS []int              `json:"newsletter_ids"`
}

func CreateRecipients(c *fiber.Ctx) error {
	recipientRequest := new(RecipientCreateRequest)

	if err := c.BodyParser(recipientRequest); err != nil {
		return err
	}

	var newsletters []models.Newsletter

	models.DB.Find(&newsletters, "id in ? ", recipientRequest.NewsletterIDS)

	for _, n := range newsletters {
		fmt.Println(n)
	}

	for _, recipient := range recipientRequest.Recipients {

		r := models.Recipient{
			Name:        recipient.Name,
			Email:       recipient.Email,
			Newsletters: newsletters,
		}

		models.DB.Create(&r)
	}

	return c.JSON(fiber.Map{"ok": true})
}
