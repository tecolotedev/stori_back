package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func SendEmail() {

	user := "7baf9402d99a97"           //config.EnvVars.EMAIL_USER
	password := "2bf0a481a09bb7"       //config.EnvVars.EMAIL_PASSWORD
	host := "sandbox.smtp.mailtrap.io" //config.EnvVars.EMAIL_HOST
	port := "2525"                     //config.EnvVars.EMAIL_PORT

	from := "hello@tecolotedev.com"
	to := []string{
		"wayaksdron@gmail.com",
	}

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

	msg := []byte("From: hello@tecolotedev.com\r\n" +
		"To: wayaksdron@gmail.com\r\n" +
		"Subject: Test mail\r\n\r\n" + bodyContentBuffer.String())

	auth := smtp.PlainAuth("", user, password, host)

	err = smtp.SendMail(host+":"+port, auth, from, to, msg)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Email sent successfully")
}
