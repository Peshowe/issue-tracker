package smtp

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/Peshowe/issue-tracker/mail-service/mailer"
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
)

type mailServer struct {
	dialer *gomail.Dialer
}

func NewMailServer(smptpAddress string, smtpPort int) mailer.MailServer {
	// Settings for SMTP server
	d := gomail.NewDialer(smptpAddress, smtpPort, "", "")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &mailServer{d}
}

func (m *mailServer) SendMail(ctx context.Context, message *gomail.Message) error {
	// Send E-Mail
	if err := m.dialer.DialAndSend(message); err != nil {
		err = errors.Wrap(err, "mail-server.SendMail")
		log.Println(err)
		return err
	}
	return nil
}
