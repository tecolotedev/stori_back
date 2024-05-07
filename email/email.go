package email

import (
	"fmt"
	"strconv"

	"github.com/go-mail/mail"
	"github.com/tecolotedev/stori_back/config"
	"github.com/tecolotedev/stori_back/utils"
)

var Dialer *mail.Dialer

type Email struct {
	To      string
	Subject string
}

type NewsletterEmail struct {
	Email
	Name         string
	File         string
	Content      string
	UserID       int
	NewsletterID int
}

type EmailHandlerStruct struct {
	NewsletterEmailChan chan NewsletterEmail
	DoneChan            chan bool
}

func (e *EmailHandlerStruct) ListenEmails() {
	fmt.Println("listen emails")
	for {

		select {
		case email := <-e.NewsletterEmailChan:
			go sendNewsletterEmail(email)

		case <-e.DoneChan:
			return
		}

	}
}

func (e *EmailHandlerStruct) InitDialer() {
	user := config.EnvVars.EMAIL_USER
	password := config.EnvVars.EMAIL_PASSWORD
	host := config.EnvVars.EMAIL_HOST
	port, err := strconv.Atoi(config.EnvVars.EMAIL_PORT)

	if err != nil {
		port = 587
	}
	d := mail.NewDialer(host, port, user, password)
	Dialer = d
}

func sendNewsletterEmail(email NewsletterEmail) {
	htmlContent := GetNewsletterHTMLBody(email.Name, email.Content, email.UserID, email.NewsletterID)
	sendEmail(email.To, email.Subject, email.File, htmlContent)

}

func sendEmail(to, subject, file, htmlContent string) {

	from := "hello@tecolotedev.com"

	m := mail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlContent)

	if file != "" {
		m.Attach("files/" + file)
	}

	if err := Dialer.DialAndSend(m); err != nil {
		utils.ErrorLog(err)
	}
}

var EmailHandler EmailHandlerStruct
