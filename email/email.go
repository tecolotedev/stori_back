package email

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/go-mail/mail"
)

func SendEmail() {

	user := "7baf9402d99a97"           //config.EnvVars.EMAIL_USER
	password := "2bf0a481a09bb7"       //config.EnvVars.EMAIL_PASSWORD
	host := "sandbox.smtp.mailtrap.io" //config.EnvVars.EMAIL_HOST
	port := 2525                       //config.EnvVars.EMAIL_PORT

	from := "hello@tecolotedev.com"
	to := "wayaksdron@gmail.com"

	f, err := os.ReadFile("email/template.html") // just pass the file name

	if err != nil {
		fmt.Println("err: ", err)
	}

	tmpl, err := template.New("test").Parse(string(f))
	if err != nil {
		fmt.Println(err)
	}

	name := "asmdfkas alksmd"
	message := "This is a sample email sent from a Go program."
	var bodyContentBuffer bytes.Buffer

	err = tmpl.Execute(&bodyContentBuffer, struct {
		Name    string
		Message string
	}{
		Name:    name,
		Message: message,
	})
	if err != nil {
		fmt.Println(err)
	}

	m := mail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", bodyContentBuffer.String())
	m.Attach("lolcat.jpg")

	d := mail.NewDialer(host, port, user, password)

	if err := d.DialAndSend(m); err != nil {

		fmt.Println(err)

	}

	fmt.Println("Email sent successfully")
}
