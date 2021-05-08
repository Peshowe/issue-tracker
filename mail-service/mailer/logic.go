package mailer

import (
	"context"
	"fmt"
	"log"

	"github.com/go-gomail/gomail"
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

func (m *mailService) ListenForEvents(ctx context.Context, done chan bool) error {

	for {
		select {

		case <-done:
			m.eventSubscriber.CloseChannels()
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

			m.SendMail(ctx, message)

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

			m.SendMail(ctx, message)
		}
	}
}

func (m *mailService) BuildMessageFromIssueEvent(ctx context.Context, event *IssueEvent) (*gomail.Message, error) {
	var messageBody string
	var receivers []string
	switch event.Type {

	case IssueCreatedEventType:
		messageBody = fmt.Sprintf("Issue created in project %v", event.Issue.Project)

	case IssueUpdatedEventType:
		messageBody = fmt.Sprintf("Issue %v updated in project %v", event.Issue.Name, event.Issue.Project)

	case IssueDeletedEventType:
		messageBody = fmt.Sprintf("Issue %v deleted in project %v", event.Issue.Name, event.Issue.Project)

	default:
		return nil, ErrEventUnknown
	}

	message := gomail.NewMessage()

	// receivers = []string{event.Issue.User}
	receivers = []string{"nedelevbg@gmail.com"}

	message.SetBody("text/plain", messageBody)

	message.SetHeaders(map[string][]string{
		"From":    {message.FormatAddress("noreply@parvusjira.com", "Parvus JIRA")},
		"To":      receivers,
		"Subject": {"Parvus JIRA Notification"},
	})

	return message, nil
}

func (m *mailService) BuildMessageFromProjectEvent(ctx context.Context, event *ProjectEvent) (*gomail.Message, error) {
	return nil, nil
}

func (m *mailService) SendMail(ctx context.Context, message *gomail.Message) error {
	return m.mailServer.SendMail(ctx, message)
}
