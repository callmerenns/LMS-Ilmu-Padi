package service

import (
	"net/smtp"

	"github.com/kelompok-2/ilmu-padi/config"
)

type MailService struct {
	cfg config.SmtpConfig
}

func (m *MailService) SendMail(subject, html string, to []string) error {
	auth := smtp.PlainAuth(
		"",
		m.cfg.EmailName,
		m.cfg.EmailAppPswd,
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		m.cfg.EmailName,
		to,
		[]byte(msg),
	)
	if err != nil {
		return err
	}

	return nil
}

func NewMailService(config config.SmtpConfig) *MailService {
	return &MailService{
		cfg: config,
	}
}
