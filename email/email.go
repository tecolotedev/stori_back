package email

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"os"
	"strconv"

	"github.com/go-mail/mail"
	"github.com/tecolotedev/stori_back/config"
)

func SendSignupEmail(name string, id int32, to string) {
	f, err := os.ReadFile("email/signup.html") // just pass the file name
	if err != nil {
		fmt.Println("err loading template: ", err)
	}
	tmpl, err := template.New("template").Parse(string(f))
	if err != nil {
		fmt.Println(err)
	}

	var bodyContentBuffer bytes.Buffer

	err = tmpl.Execute(&bodyContentBuffer, struct {
		Name      string
		UrlSignup string
	}{
		Name:      name,
		UrlSignup: config.EnvVars.FRONT_URL + "/verifyAccount?id=" + strconv.Itoa(int(id)),
	})
	if err != nil {
		fmt.Println(err)
	}

	SendEmail(to, bodyContentBuffer.String())

}

type Record struct {
	Date        string
	Transaction float64
	Reason      string
}

func SendReportEmail(to string, balance float64, records []Record) {
	f, err := os.ReadFile("email/report.html") // just pass the file name
	if err != nil {
		fmt.Println("err loading template: ", err)
	}
	tmpl, err := template.New("template").Parse(string(f))
	if err != nil {
		fmt.Println(err)
	}

	var bodyContentBuffer bytes.Buffer

	err = tmpl.Execute(&bodyContentBuffer, struct {
		Records []Record
		Balance float64
	}{
		Records: records,
		Balance: math.Round(balance*100) / 100,
	})
	if err != nil {
		fmt.Println(err)
	}

	SendEmail(to, bodyContentBuffer.String())
}

func SendEmail(to, htmlContent string) {

	user := config.EnvVars.EMAIL_USER
	password := config.EnvVars.EMAIL_PASSWORD
	host := config.EnvVars.EMAIL_HOST
	port, err := strconv.Atoi(config.EnvVars.EMAIL_PORT)

	if err != nil {
		port = 587
	}

	from := "hello@tecolotedev.com"

	m := mail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", htmlContent)
	m.Attach("lolcat.jpg")

	d := mail.NewDialer(host, port, user, password)

	if err := d.DialAndSend(m); err != nil {

		fmt.Println("err seding email: ", err)

	}

}
