package email

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/EMPAT94/pingmon/internal/config"
)

type Mailer struct {
	emailId    string
	password   string
	mailServer string
	auth       smtp.Auth
	recipients config.Recipients
}

func New(cfg *config.Config) *Mailer {
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
	return &Mailer{cfg.Mailer.Id, cfg.Mailer.Password, cfg.Mailer.Server, auth, cfg.Recipients}
}

func (m *Mailer) Alert(msg string) error {
	content := []byte("To: " + strings.Join(m.recipients.Id, ",") + "\r\n" +
		"Subject: Monitoring Alert\r\n" +
		"\r\n" + msg + "\r\n")

	err := smtp.SendMail(m.mailServer, m.auth, m.emailId, m.recipients.Id, content)
	if err != nil {
		return fmt.Errorf("Error while sending email %v:", err)
	}

	return nil
}
