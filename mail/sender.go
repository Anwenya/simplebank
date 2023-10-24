package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpGmailAuthAddress     = "smtp.gmail.com"
	smtpGmailServerAddress   = "smtp.gmail.com:587"
	smtpTencentAuthAddress   = "smtp.qq.com"
	smtpTencentServerAddress = "smtp.qq.com:25"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGamilSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s : %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpGmailAuthAddress)

	return e.Send(smtpGmailServerAddress, smtpAuth)
}

type TencentSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewTencentSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &TencentSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *TencentSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s : %w", f, err)
		}
	}
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpTencentAuthAddress)
	return e.Send(smtpTencentServerAddress, smtpAuth)
}
