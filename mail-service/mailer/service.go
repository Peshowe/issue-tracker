package mailer

import (
	"context"

	"github.com/go-gomail/gomail"
)

type MailService interface {
	ListenForEvents(ctx context.Context, done chan bool) error

	BuildMessageFromIssueEvent(ctx context.Context, event *IssueEvent) (*gomail.Message, error)
	BuildMessageFromProjectEvent(ctx context.Context, event *ProjectEvent) (*gomail.Message, error)

	SendMail(ctx context.Context, message *gomail.Message) error
}
