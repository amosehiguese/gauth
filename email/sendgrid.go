package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridSender struct {
	apiKey   string
	from     string
	fromName string
}

func NewSendGridSender(apiKey, from, fromName string) *SendGridSender {
	return &SendGridSender{
		apiKey:   apiKey,
		from:     from,
		fromName: fromName,
	}
}

func (s *SendGridSender) Send(to string, subject string, body string, html string) error {
	message := mail.NewSingleEmail(
		mail.NewEmail(s.fromName, s.from),
		subject,
		mail.NewEmail("", to),
		body,
		html,
	)

	client := sendgrid.NewSendClient(s.apiKey)
	_, err := client.Send(message)
	return err
}
