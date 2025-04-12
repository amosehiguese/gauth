package email

type EmailProvider interface {
	Send(to, subject, body, html string) error
}
