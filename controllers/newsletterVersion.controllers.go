package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/email"
	"github.com/tecolotedev/stori_back/models"
	"github.com/tecolotedev/stori_back/utils"
)

func GetAllNewslettersVersions(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")
	if err != nil {
		utils.ErrorLog(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	var newsletterVersions []models.NewsletterVersion

	result := models.DB.Where("newsletter_id = ? ", newsletterID).Find(&newsletterVersions)

	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
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
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
	}

	file, err := c.FormFile("file")
	var fileName string
	if err != nil {
		utils.ErrorLog(err)
	} else {
		fileName = utils.AddPrefixToFilename(file.Filename)
		c.SaveFile(file, fmt.Sprintf("./files/%s", fileName))
	}

	nlvr := new(NewsletterVersionRequest)

	if err := c.BodyParser(nlvr); err != nil {
		utils.ErrorLog(err)
		return c.Status(int(BodyError.code)).JSON(fiber.Map{"ok": false, "message": BodyError.message})
	}

	newsletterVersion := models.NewsletterVersion{
		Title:        nlvr.Title,
		Content:      nlvr.Content,
		File:         fileName,
		NewsletterID: uint(newsletterID),
	}

	result := models.DB.Create(&newsletterVersion)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(newsletterVersion)
}

type NewsletterVersionUpdateRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func UpdateNewsletterVersion(c *fiber.Ctx) error {

	newsletterVersionID, err := c.ParamsInt("newsletter_version_id")

	if err != nil {
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
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
		utils.ErrorLog(err)
		return c.Status(int(BodyError.code)).JSON(fiber.Map{"ok": false, "message": BodyError.message})
	}

	newsletterVersion := models.NewsletterVersion{
		Title:   nlvr.Title,
		Content: nlvr.Content,
		File:    fileName,
	}

	result := models.DB.Where("id = ?", newsletterVersionID).Updates(&newsletterVersion)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(newsletterVersion)
}

func DeleteNewsletterVersion(c *fiber.Ctx) error {
	newsletterVersionID, err := c.ParamsInt("newsletter_version_id")

	if err != nil {
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
	}

	result := models.DB.Where("id = ?", newsletterVersionID).Delete(&models.NewsletterVersion{})
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	return c.JSON(fiber.Map{"ok": true, "message": "deleted"})
}

func SendNewsletter(c *fiber.Ctx) error {
	newsletterVersionID, err := c.ParamsInt("newsletter_version_id")
	if err != nil {
		utils.ErrorLog(err)
		return c.Status(int(ParamsError.code)).JSON(fiber.Map{"ok": false, "message": ParamsError.message})
	}

	var newsletterVersion models.NewsletterVersion
	result := models.DB.First(&newsletterVersion, newsletterVersionID)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	var newsletter models.Newsletter
	result = models.DB.Model(&models.Newsletter{}).Preload("Recipients").Find(&newsletter, "id = ?", newsletterVersion.NewsletterID)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
		return c.Status(int(DatabaseError.code)).JSON(fiber.Map{"ok": false, "message": DatabaseError.message})
	}

	for _, recipient := range newsletter.Recipients {

		to := recipient.Email
		name := recipient.Name

		subject := newsletterVersion.Title
		content := newsletterVersion.Content
		file := newsletterVersion.File

		// sent to email to channel
		newEmail := email.NewsletterEmail{
			Email: email.Email{
				To:      to,
				Subject: subject,
			},
			File:    file,
			Name:    name,
			Content: content,
		}

		email.EmailHandler.NewsletterEmailChan <- newEmail

	}

	// update email sent
	result = models.DB.Model(&models.NewsletterVersion{}).Where("id = ?", newsletterVersionID).Update("sent", true)
	if result.Error != nil {
		utils.ErrorLog(result.Error)
	}

	return c.JSON(fiber.Map{"ok": true, "message": "email sent"})

}
