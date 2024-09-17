package services

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

var MAIL string
var MAIL_PASSWORD string
var SMTP string
var PORT int

// Function to send an EMAIL
func SendMail(to string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", MAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	fmt.Println(m, SMTP, MAIL_PASSWORD)

	d := gomail.NewDialer(SMTP, PORT, MAIL, MAIL_PASSWORD)

	d.TLSConfig = nil

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
