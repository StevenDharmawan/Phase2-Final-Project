package helper

import (
	"fmt"
	"net/smtp"
	"phase2-final-project/config"
)

func SendEmail(email string, subject string, emailContent string) error {
	smtpHost := "smtp.gmail.com"
	smtpUsername := config.GetEnv("EMAIL")
	smtpPassword := config.GetEnv("PASSWORD")
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	msg := fmt.Sprintf("Subject: %s\n\n%s", subject, emailContent)
	err := smtp.SendMail("smtp.gmail.com:587", auth, smtpUsername, []string{email}, []byte(msg))
	return err
}
