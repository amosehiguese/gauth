package email

import (
	"github.com/amosehiguese/gauth/config"
)

func NewEmailSender(cfg config.Config) EmailProvider {
	switch cfg.EmailProvider {
	case "smtp":
		return NewSMTPSender(cfg)
	case "sendgrid":
		return NewSendGridSender(cfg.SendGridAPIKey, cfg.EmailFrom, cfg.EmailFrom)
	case "mailgun":
		return NewMailgunSender(cfg.MailGunDomain, cfg.MailGunAPIKey, cfg.EmailFrom)
	default:
		panic("unsupported email provider")
	}
}
