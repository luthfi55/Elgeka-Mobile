package models

import (
	"net/smtp"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type GmailSender struct {
	FromEmailAddress  string
	FromEmailPassword string
}

// NewGmailSender creates a new GmailSender instance
func NewGmailSender(emailAddress, password string) *GmailSender {
	return &GmailSender{
		FromEmailAddress:  emailAddress,
		FromEmailPassword: password,
	}
}

// SendEmail sends an email using the GmailSender's configuration
func (g *GmailSender) SendEmail(subject, content string, to, cc, bcc []string) error {
	auth := smtp.PlainAuth("", g.FromEmailAddress, g.FromEmailPassword, smtpAuthAddress)

	msg := "From: " + g.FromEmailAddress + "\n" +
		"Content-type: text/html; charset=UTF-8\r\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n\n" +
		content

	err := smtp.SendMail(smtpServerAddress, auth, g.FromEmailAddress, to, []byte(msg))
	return err
}
