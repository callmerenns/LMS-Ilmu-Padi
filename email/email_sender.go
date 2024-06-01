package email

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridEmailSender struct {
	apiKey string
	email  string
}

func NewSendGridEmailSender(apiKey string, email string) *SendGridEmailSender {
	return &SendGridEmailSender{
		apiKey: apiKey,
		email:  email,
	}
}

func (s *SendGridEmailSender) Send(to string, subject string, body string) error {
	from := mail.NewEmail("Example User", s.email)
	toEmail := mail.NewEmail("Recipient", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, body)
	client := sendgrid.NewSendClient(s.apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Printf("Email sent successfully with status code: %d", response.StatusCode)
	return nil
}
