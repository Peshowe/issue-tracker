package mailer

import (
	"context"
	"github.com/go-gomail/gomail"
)

type MailServer interface {
	SendMail(ctx context.Context, message *gomail.Message) error
}
