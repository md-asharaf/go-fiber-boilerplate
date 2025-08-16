package services

import (
	"github.com/go-mail/mail"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
)

type EmailService struct {
	mailer *mail.Dialer
	From   string
}

func NewEmailService(emailConfig config.SMTPConfig) *EmailService {
	mailer := mail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.Username, emailConfig.Password)
	return &EmailService{
		mailer: mailer,
		From:   emailConfig.From,
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	if err := e.mailer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
