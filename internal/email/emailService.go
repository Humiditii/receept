package email

import (
	"fmt"
	"net/smtp"

	"github.com/Humiditii/receept/internal/config"
)

type EmailService struct {
	smtpHost string
	smtpPort string
	auth     smtp.Auth
	from     string
}

func NewEmailService(cfg *config.AppConfig) *EmailService {
	auth := smtp.PlainAuth("", (*cfg).EmailUsername, (*cfg).EmailPassword, (*cfg).EmailHost)
	return &EmailService{
		smtpHost: (*cfg).EmailHost,
		smtpPort: (*cfg).EmailPort,
		auth:     auth,
		from:     (*cfg).EmailUsername,
	}
}

func (e *EmailService) Send(to, subject, body string) error {
	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))
	addr := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)

	return smtp.SendMail(addr, e.auth, e.from, []string{to}, msg)
}
