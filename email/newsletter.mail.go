package email

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/tecolotedev/stori_back/utils"
)

func GetNewsletterHTMLBody(name, content string, userID, newsletterID int) string {

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

	url := fmt.Sprintf("http://localhost:3000/unsubscribe?user_id=%d&newsletter_id=%d", userID, newsletterID)

	// create html content and insert data into the template
	var bodyContentBuffer bytes.Buffer
	err = tmpl.Execute(&bodyContentBuffer, struct {
		Name    string
		Content string
		URL     string
	}{
		Name:    name,
		Content: content,
		URL:     url,
	})
	if err != nil {
		utils.ErrorLog(err)
	}

	return bodyContentBuffer.String()

}
