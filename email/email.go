package email

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"os"

	"github.com/go-mail/mail"
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
		Name string
		ID   int32
	}{
		Name: name,
		ID:   id,
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

	user := "7baf9402d99a97"           //config.EnvVars.EMAIL_USER
	password := "2bf0a481a09bb7"       //config.EnvVars.EMAIL_PASSWORD
	host := "sandbox.smtp.mailtrap.io" //config.EnvVars.EMAIL_HOST
	port := 2525                       //config.EnvVars.EMAIL_PORT

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
