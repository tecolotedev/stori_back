package email

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func GetNewsletterHTMLBody(name, content string) string {

	// read template as file
	f, err := os.ReadFile("email/templates/newsletter.html")
	if err != nil {
		fmt.Println(err)
	}

	// parse the file to template object
	tmpl, err := template.New("template").Parse(string(f))
	if err != nil {
		fmt.Println(err)
	}

	var bodyContentBuffer bytes.Buffer
	err = tmpl.Execute(&bodyContentBuffer, struct {
		Name    string
		Content string
	}{
		Name:    name,
		Content: content,
	})
	if err != nil {
		fmt.Println(err)
	}

	return bodyContentBuffer.String()

}
