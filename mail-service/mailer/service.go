package mailer

import "context"

type MailService interface {
	ListenForEvents(ctx context.Context) error

	BuildMessageFromIssueEvent(ctx context.Context, event *IssueEvent) (string, error)
	GetReceiverFromIssueEvent(ctx context.Context, event *IssueEvent) (string, error)

	BuildMessageFromProjectEvent(ctx context.Context, event *ProjectEvent) (string, error)
	GetReceiverFromProjectEvent(ctx context.Context, event *ProjectEvent) (string, error)

	SendMail(ctx context.Context, receiver string, message string) error
}
