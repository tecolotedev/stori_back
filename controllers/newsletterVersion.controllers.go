package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/models"
	"github.com/tecolotedev/stori_back/utils"
)

func GetAllNewslettersVersions(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}
	var newsletterVersions []models.NewsletterVersion

	result := models.DB.Where("newsletter_id = ? ", newsletterID).Find(&newsletterVersions)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(500).JSON(fiber.Map{"ok": false, "message": "something went wrong, please try it later"})
	}

	return c.JSON(newsletterVersions)
}

type NewsletterVersionRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func CreateNewsletterVersion(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	file, err := c.FormFile("file")
	var fileName string
	if err != nil {
		fmt.Println(err)
	} else {
		fileName = utils.AddPrefixToFilename(file.Filename)
		c.SaveFile(file, fmt.Sprintf("./files/%s", fileName))
	}

	nlvr := new(NewsletterVersionRequest)

	if err := c.BodyParser(nlvr); err != nil {
		return err
	}

	newsletterVersion := models.NewsletterVersion{
		Title:        nlvr.Title,
		Content:      nlvr.Content,
		File:         fileName,
		NewsletterID: uint(newsletterID),
	}

	models.DB.Create(&newsletterVersion)

	return c.JSON(newsletterVersion)
}

type NewsletterVersionUpdateRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func UpdateNewsletterVersion(c *fiber.Ctx) error {

	newsletterVersionID, err := c.ParamsInt("newsletter_version_id")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	file, err := c.FormFile("file")
	var fileName string
	if err != nil {
		fmt.Println(err)
	} else {
		fileName = utils.AddPrefixToFilename(file.Filename)
		c.SaveFile(file, fmt.Sprintf("./files/%s", fileName))
	}

	nlvr := new(NewsletterVersionUpdateRequest)

	if err := c.BodyParser(nlvr); err != nil {
		return err
	}

	newsletterVersion := models.NewsletterVersion{
		Title:   nlvr.Title,
		Content: nlvr.Content,
		File:    fileName,
	}

	models.DB.Where("id = ?", newsletterVersionID).Updates(&newsletterVersion)

	return c.JSON(newsletterVersion)
}

func DeleteNewsletterVersion(c *fiber.Ctx) error {
	newsletterVersionID, err := c.ParamsInt("newsletter_version_id")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	models.DB.Where("id = ?", newsletterVersionID).Delete(&models.NewsletterVersion{})

	return c.JSON(fiber.Map{"ok": true, "message": "deleted"})
}

func SendNewsletter(c *fiber.Ctx) error {
	newsletterVersionID, err := c.ParamsInt("newsletter_version_id")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	var newsletterVersion models.NewsletterVersion
	models.DB.First(&newsletterVersion, newsletterVersionID)

	fmt.Println("newsletterVersion: ", newsletterVersion)

	var newsletter models.Newsletter
	models.DB.Model(&models.Newsletter{}).Preload("Recipients").Find(&newsletter, "id = ?", newsletterVersion.NewsletterID)

	fmt.Println("newsletter: ", newsletter)

	// sent to email to channel

	// update email sent
	models.DB.Model(&models.NewsletterVersion{}).Where("id = ?", newsletterVersionID).Update("sent", true)

	return c.JSON(fiber.Map{"ok": true, "message": "email sent"})

}
