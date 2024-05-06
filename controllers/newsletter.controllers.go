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

type NewsletterCreateRequest struct {
	Name string `json:"name" form:"name"`
}

func CreateNewsletter(c *fiber.Ctx) error {
	nlr := new(NewsletterCreateRequest)

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

type NewsletterUpdateRequest struct {
	Name string `json:"name" form:"name"`
}

func UpdateNewsletter(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	nlr := new(NewsletterUpdateRequest)

	if err := c.BodyParser(nlr); err != nil {
		return err
	}

	models.DB.Model(&models.Newsletter{}).Where("id = ?", newsletterID).Update("Name", nlr.Name)

	return c.JSON(fiber.Map{"ok": true, "message": "updated"})
}

func DeleteNewsletter(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	models.DB.Where("newsletter_id = ?", newsletterID).Delete(&models.NewsletterVersion{})

	models.DB.Where("id = ?", newsletterID).Delete(&models.Newsletter{})

	return c.JSON(fiber.Map{"ok": true, "message": "deleted"})
}
