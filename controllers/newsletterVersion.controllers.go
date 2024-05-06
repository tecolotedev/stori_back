package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/models"
)

// func GetAllNewsletterVersions(c *fiber.Ctx) error {
// 	// var newsletters []models.Newsletter

// 	// result := models.DB.Find(&newsletters)

// 	// if result.Error != nil {
// 	// 	fmt.Println(result.Error)
// 	// 	c.Status(500).JSON(fiber.Map{"ok": false, "message": "something went wrong, please try it later"})
// 	// }

// 	// return c.JSON(newsletters)
// }

type NewsletterVersionRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	// File
}

func CreateNewsletterVersion(c *fiber.Ctx) error {
	newsletterID, err := c.ParamsInt("newsletter_id")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	nlvr := new(NewsletterVersionRequest)

	if err := c.BodyParser(nlvr); err != nil {
		return err
	}

	splitFile := strings.Split(file.Filename, ".")
	unix := time.Now().Unix()

	fileName := splitFile[0] + "_" + strconv.Itoa(int(unix)) + "." + splitFile[1]

	c.SaveFile(file, fmt.Sprintf("./files/%s", fileName))

	newsletterVersion := models.NewsletterVersion{
		Title:        nlvr.Title,
		Content:      nlvr.Content,
		File:         fileName,
		NewsletterID: uint(newsletterID),
	}

	models.DB.Create(&newsletterVersion)

	return c.JSON(newsletterVersion)
}
