package middleware

import (
	"log"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(fromTitle string, 
		toTitle string, 
		toEmail string,
		subjectLine string, 
		contentPlain string, 
		contentHtml string) bool {
	from := mail.NewEmail(fromTitle, "wasev74013@procowork.com")
	subject := subjectLine
	to := mail.NewEmail(toTitle, toEmail)
	plainTextContent := contentPlain
	htmlContent := contentHtml
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	
	if err != nil {
		return true
	} else {
		log.Println(response.StatusCode, response.Body, response.Headers)
		log.Println(err)
	}
	return false
}