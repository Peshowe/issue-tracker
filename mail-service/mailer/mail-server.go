package mailer

import "context"

type MailServer interface {
	SendMail(ctx context.Context, receiver string, message string) error
}
