package mailer

import (
	"context"
	"fmt"
	"log"
)

type mailService struct {
	mailServer      MailServer
	eventSubscriber EventSubscriber
	eventDecoder    EventDecoder
}

func NewMailService(mailServer MailServer, eventSubscriber EventSubscriber, eventDecoder EventDecoder) MailService {
	return &mailService{
		mailServer,
		eventSubscriber,
		eventDecoder,
	}
}

func (m *mailService) ListenForEvents(ctx context.Context) error {
	done := make(chan bool)

	for {
		select {

		case <-done:
			return nil

		case event := <-m.eventSubscriber.GetIssueEventChannel(ctx):
			//Handle incoming issue events
			log.Println("Issue event received")

			issueEvent := &IssueEvent{}
			m.eventDecoder.DecodeEvent(event, issueEvent)

			message, err := m.BuildMessageFromIssueEvent(ctx, issueEvent)
			if err != nil {
				log.Println(err)
				break
			}

			receiver, err := m.GetReceiverFromIssueEvent(ctx, issueEvent)
			if err != nil {
				log.Println(err)
				break
			}

			m.SendMail(ctx, receiver, message)

		case event := <-m.eventSubscriber.GetProjectEventChannel(ctx):
			//Handle incoming project events
			log.Println("Project event received")

			projectEvent := &ProjectEvent{}
			m.eventDecoder.DecodeEvent(event, projectEvent)

			message, err := m.BuildMessageFromProjectEvent(ctx, projectEvent)
			if err != nil {
				log.Println(err)
				break
			}

			receiver, err := m.GetReceiverFromProjectEvent(ctx, projectEvent)
			if err != nil {
				log.Println(err)
				break
			}

			m.SendMail(ctx, receiver, message)
		}
	}
}

func (m *mailService) BuildMessageFromIssueEvent(ctx context.Context, event *IssueEvent) (string, error) {
	var message string
	switch event.Type {

	case IssueCreatedEventType:
		message = fmt.Sprintf("Issue created in project %v", event.Issue.Project)

	case IssueUpdatedEventType:
		message = fmt.Sprintf("Issue %v updated in project %v", event.Issue.Name, event.Issue.Project)

	case IssueDeletedEventType:
		message = fmt.Sprintf("Issue %v deleted in project %v", event.Issue.Name, event.Issue.Project)

	default:
		return "", ErrEventUnknown
	}

	return message, nil
}

func (m *mailService) GetReceiverFromIssueEvent(ctx context.Context, event *IssueEvent) (string, error) {
	return event.Issue.User, nil
}

func (m *mailService) BuildMessageFromProjectEvent(ctx context.Context, event *ProjectEvent) (string, error) {
	return "", nil
}

func (m *mailService) GetReceiverFromProjectEvent(ctx context.Context, event *ProjectEvent) (string, error) {
	return "", nil
}

func (m *mailService) SendMail(ctx context.Context, receiver string, message string) error {
	return m.mailServer.SendMail(ctx, receiver, message)
}
