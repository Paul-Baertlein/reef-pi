package telemetry

import (
	"net/smtp"
	"strings"
	"testing"
)

func TestMailer(t *testing.T) {
	GMailMailer.Mailer()
	m := &mailer{
		config:   &GMailMailer,
		sendMail: func(_ string, _ smtp.Auth, _ string, _ []string, _ []byte) error { return nil },
	}
	msg := m.msg("Hi", "")
	if !strings.Contains(msg, "Subject: Hi") {
		t.Error("subject mismatch", msg)
	}
	if err := m.Email("", ""); err != nil {
		t.Error(err)
	}
}
