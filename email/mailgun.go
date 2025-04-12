package email

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v5"
)

type MailgunSender struct {
	domain string
	apiKey string
	from   string
}

func NewMailgunSender(domain, apiKey, from string) *MailgunSender {
	return &MailgunSender{
		domain: domain,
		apiKey: apiKey,
		from:   from,
	}
}

func (s *MailgunSender) Send(to string, subject string, body string, html string) error {
	mg := mailgun.NewMailgun(s.apiKey)

	var message *mailgun.PlainMessage

	if body != "" {
		message = mailgun.NewMessage(s.domain, s.from, subject, body, to)
	} else if html != "" {
		message = mailgun.NewMessage(s.domain, s.from, subject, "", to)
		message.SetHTML(html)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := mg.Send(ctx, message)
	return err
}
