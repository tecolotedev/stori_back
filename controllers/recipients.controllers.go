package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/models"
	"github.com/tecolotedev/stori_back/utils"
)

func GetRecipients(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
	}

	var newsletter models.Newsletter

	result := models.DB.Model(&models.Newsletter{}).Preload("Recipients").Find(&newsletter, "id = ?", newsletterID)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(newsletter)
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
		utils.ErrorLog(err)
		return c.Status(int(BodyError.code)).JSON(fiber.Map{"ok": false, "message": BodyError.message})
	}

	var newsletters []models.Newsletter

	result := models.DB.Find(&newsletters, "id in ? ", recipientRequest.NewsletterIDS)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

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

	return c.JSON(fiber.Map{"ok": true, "message": "created"})
}

func DeleteRecipient(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
	}

	recipientID, err := c.ParamsInt("recipient_id")

	if err != nil {
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
	}

	var newsletter models.Newsletter
	result := models.DB.First(&newsletter, newsletterID)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	var recipient models.Recipient
	result = models.DB.First(&recipient, recipientID)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	err = models.DB.Model(&recipient).Association("Newsletters").Delete(newsletter)
	if err != nil {
		utils.ErrorLog(err)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(fiber.Map{"ok": true, "message": "deleted"})
}
