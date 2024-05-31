package email

import (
	"gopkg.in/gomail.v2"
)

type SMTPEmailSender struct {
    host     string
    port     int
    username string
    password string
}

func NewSMTPEmailSender(host string, port int, username, password string) *SMTPEmailSender {
    return &SMTPEmailSender{
        host:     host,
        port:     port,
        username: username,
        password: password,
    }
}

func (s *SMTPEmailSender) Send(to, subject, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", s.username)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer(s.host, s.port, s.username, s.password)
    return d.DialAndSend(m)
}