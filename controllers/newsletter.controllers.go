package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/models"
	"github.com/tecolotedev/stori_back/utils"
)

func GetAllNewsletters(c *fiber.Ctx) error {
	var newsletters []models.Newsletter

	result := models.DB.Find(&newsletters)

	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(newsletters)
}

type NewsletterCreateRequest struct {
	Name string `json:"name" form:"name"`
}

func CreateNewsletter(c *fiber.Ctx) error {
	nlr := new(NewsletterCreateRequest)

	if err := c.BodyParser(nlr); err != nil {
		utils.ErrorLog(err)
		return c.Status(int(BodyError.code)).JSON(fiber.Map{"ok": false, "message": BodyError.message})
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
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
	}

	nlr := new(NewsletterUpdateRequest)

	if err := c.BodyParser(nlr); err != nil {
		utils.ErrorLog(err)
		return c.Status(int(BodyError.code)).JSON(fiber.Map{"ok": false, "message": BodyError.message})
	}

	result := models.DB.Model(&models.Newsletter{}).Where("id = ?", newsletterID).Update("Name", nlr.Name)

	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(fiber.Map{"ok": true, "message": "updated"})
}

func DeleteNewsletter(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		utils.ErrorLog(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	result := models.DB.Where("newsletter_id = ?", newsletterID).Delete(&models.NewsletterVersion{})
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	result = models.DB.Where("id = ?", newsletterID).Delete(&models.Newsletter{})
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(fiber.Map{"ok": true, "message": "deleted"})
}
