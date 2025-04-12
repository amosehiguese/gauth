package email

import (
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/amosehiguese/gauth/config"
)

type SMTPSender struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewSMTPSender(cfg config.Config) *SMTPSender {
	return &SMTPSender{
		host:     cfg.SMTPHost,
		port:     cfg.SMTPPort,
		username: cfg.SMTPUsername,
		password: cfg.SMTPPassword,
		from:     cfg.EmailFrom,
	}
}

func (s *SMTPSender) Send(to string, subject string, body string, html string) error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	var buf bytes.Buffer

	// Add headers
	buf.WriteString(fmt.Sprintf("From: %s\r\n", s.from))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", to))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("MIME-Version: 1.0\r\n")

	// If HTML is provided, create a multipart message
	if html != "" {
		boundary := "next-part"
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/alternative; boundary=%s\r\n", boundary))
		buf.WriteString("\r\n")

		// Text part
		buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		buf.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
		buf.WriteString("\r\n")
		buf.WriteString(body)
		buf.WriteString("\r\n\r\n")

		// HTML part
		buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		buf.WriteString("Content-Type: text/html; charset=utf-8\r\n")
		buf.WriteString("\r\n")
		buf.WriteString(html)
		buf.WriteString("\r\n\r\n")

		buf.WriteString(fmt.Sprintf("--%s--", boundary))
	} else {
		// Plain text message only
		buf.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
		buf.WriteString("\r\n")
		buf.WriteString(body)
	}

	// Send email
	return smtp.SendMail(addr, auth, s.from, []string{to}, buf.Bytes())
}
