package mail

import (
	"gohub/pkg/config"
	"sync"
)

type From struct {
	Name    string
	Address string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Plaintext Message
	HTML    []byte // Html Massage
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})
	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.stmp"))
}
