package main

import (
	"io"

	"github.com/rubencaro/edward/lib/cnf"
	gomail "gopkg.in/gomail.v2"
)

// SendAlertEmail sends given image attached to an alert email using details in given Config,
// It receives an optional DialerSender to make it testable.
func SendAlertEmail(c *cnf.Config, img []byte, d DialerSender) error {
	m := gomail.NewMessage()
	m.SetHeader("From", c.From)
	m.SetHeader("To", c.To)
	m.SetHeader("Subject", "Movement detected")
	m.SetBody("text/html", "I'm Edward, I have detected some movement. Interesting image attached.")
	m.Attach("image.jpg", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(img)
		return err
	}))

	if d == nil {
		d = newDefaultDialerSender("smtp.gmail.com", 587, c.From, c.Pass)
	}

	return d.DialAndSend(m)
}

// DialerSender is the interface for the email senders
type DialerSender interface {
	// DialAndSend auths with the SMTP server and sends the email
	// Mimics gomail.Dialer's DialAndSend method
	DialAndSend(m ...*gomail.Message) error
}

func newDefaultDialerSender(host string, port int, username string, password string) DialerSender {
	return gomail.NewDialer(host, port, username, password)
}
