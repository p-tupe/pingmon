package email

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/p-tupe/pingmon/internal/app/config"
)

func SendMail(config *config.Config) {
	host := config.Mailer.Host
	port := config.Mailer.Port
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	username := config.Mailer.Username
	password := config.Mailer.Password
	from := config.Mailer.From
	to := config.EmailTo
	subject := "Pingmon Alert"
	body := "Could not reach site"

	message := constructEmail(from, to, subject, body)
	auth := smtp.PlainAuth("", username, password, host)
	err := smtp.SendMail(serverAddr, auth, from, to, []byte(message))
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}
}

func constructEmail(from string, to []string, subject, body string) string {
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = strings.Join(to, "")
	headers["Subject"] = subject
	headers["Content-Type"] = "text/plain; charset=UTF-8"

	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + body

	return message
}
