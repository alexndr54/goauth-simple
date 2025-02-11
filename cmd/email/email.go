package email

import (
	"bytes"
	"errors"
	"html/template"
	"net/smtp"
	"os"
)

func SendHTMLEmail(to, htmlBody, subject string, data map[string]interface{}) error {
	mailServer := os.Getenv("MAIL_SERVER")
	mailPort := os.Getenv("MAIL_PORT")
	mailUsername := os.Getenv("MAIL_USERNAME")
	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailFrom := os.Getenv("MAIL_FROM")

	if mailServer == "" || mailPort == "" || mailUsername == "" || mailPassword == "" || mailFrom == "" {
		return errors.New("Cek file .env pastikan anda sudah mengatur MAIL")
	}

	auth := smtp.PlainAuth("", mailUsername, mailPassword, mailServer)

	t := template.Must(template.New("email").Parse(htmlBody))
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return errors.New("Failed to execute template: " + err.Error())
	}

	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		buf.String())

	err := smtp.SendMail(mailServer+":"+mailPort, auth, mailFrom, []string{to}, message)
	if err != nil {
		return errors.New("Failed to send email: " + err.Error())
	}

	return nil
}
