package main

import (
	"testing"

	"github.com/rubencaro/edward/lib/tst"

	"github.com/rubencaro/edward/lib/cnf"
	gomail "gopkg.in/gomail.v2"
)

func TestSendAlertMail(t *testing.T) {
	c := &cnf.Config{
		From: "fromme",
		To:   "tome",
	}
	d := &MockedDialerSender{T: t}
	img := []byte("")
	err := SendAlertEmail(c, img, d)
	tst.Ok(t, err)
}

type MockedDialerSender struct {
	T *testing.T
}

func (mds MockedDialerSender) DialAndSend(m ...*gomail.Message) error {
	t := mds.T
	tst.Eq(t, []string{"fromme"}, m[0].GetHeader("From"))
	tst.Eq(t, []string{"tome"}, m[0].GetHeader("To"))
	tst.Eq(t, []string{"Movement detected"}, m[0].GetHeader("Subject"))
	tst.Ok(t, nil)
	return nil
}
