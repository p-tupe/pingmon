package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func Alert(msg string) {
	if cfg.PostRequest != nil {
		_, err := http.Post(cfg.PostRequest.URL, cfg.PostRequest.ContentType, strings.NewReader(msg))
		if err != nil {
			log.Println("Error while alerting on ntfy:", err.Error())
		} else {
			log.Println("Alert sent on post request")
		}
	}

	if cfg.Mailer != nil {
		auth := smtp.PlainAuth("", cfg.Mailer.Username, cfg.Mailer.Password, cfg.Mailer.Host)
		msg := fmt.Appendf(make([]byte, 0, 50), "From: %s \r\nTo: %s\r\nSubject: Pingmon Alert\r\n\r\n %s \r\n", cfg.Mailer.From, strings.Join(cfg.MailTo, ","), msg)
		err := smtp.SendMail(cfg.Mailer.Host+":"+cfg.Mailer.Port, auth, cfg.Mailer.From, cfg.MailTo, msg)
		if err != nil {
			log.Println("Unable to send email:", err.Error())
		} else {
			log.Println("Alert sent on email")
		}
	}
}
