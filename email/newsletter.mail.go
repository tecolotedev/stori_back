package email

import (
	"bytes"
	"os"
	"text/template"

	"github.com/tecolotedev/stori_back/utils"
)

func GetNewsletterHTMLBody(name, content string) string {

	// read template as file
	f, err := os.ReadFile("email/templates/newsletter.html")
	if err != nil {
		utils.ErrorLog(err)
	}

	// parse the file to template object
	tmpl, err := template.New("template").Parse(string(f))
	if err != nil {
		utils.ErrorLog(err)
	}

	// create html content and insert data into the template
	var bodyContentBuffer bytes.Buffer
	err = tmpl.Execute(&bodyContentBuffer, struct {
		Name    string
		Content string
	}{
		Name:    name,
		Content: content,
	})
	if err != nil {
		utils.ErrorLog(err)
	}

	return bodyContentBuffer.String()

}
