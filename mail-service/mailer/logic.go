package mailer

import (
	"context"
	"fmt"
	"log"

	"github.com/Peshowe/issue-tracker/mail-service/utils"
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
)

//ErrNoRecepients should be returned if there is no one to send an email to
var ErrNoRecepients error = errors.New("No recepients to notify")

type mailService struct {
	mailServer      MailServer
	eventSubscriber EventSubscriber
	eventDecoder    EventDecoder
	mailRepo        MailRepo
}

func NewMailService(mailServer MailServer, eventSubscriber EventSubscriber, eventDecoder EventDecoder, mailRepo MailRepo) MailService {
	return &mailService{
		mailServer,
		eventSubscriber,
		eventDecoder,
		mailRepo,
	}
}

//ListenForEvents listens on all the channels for incoming events and handles them appropriately
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

//BuildMessageFromIssueEvent builds the entire gomail.Message from the Issue event, including body, subject, receivers, etc
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
		return nil, errors.Wrap(ErrEventUnknown, "mailer.BuildMessageFromIssueEvent")
	}

	message := gomail.NewMessage()

	// receivers = []string{event.Issue.User}
	receivers = []string{"nedelevbg@gmail.com"}

	//check user notification preferences
	i := 0 // output index
	for _, receiver := range receivers {
		preference, err := m.mailRepo.GetUserPreference(ctx, receiver)
		if err != nil {
			return nil, errors.Wrap(err, "mailer.BuildMessageFromIssueEvent")
		}
		if preference != nil && preference.IsMailNotificationOn {
			//keep a user in the receivers list if they've turned on mail notifications (if the user is not in the repo, their notification are off by default)
			//	(here we're rewriting the slice in-place, keeping only the valid receivers)
			receivers[i] = receiver
			i++
		}
	}
	receivers = receivers[:i]

	if len(receivers) == 0 {
		//in case everyone is filtered out
		return nil, errors.Wrap(ErrNoRecepients, "mailer.BuildMessageFromIssueEvent")
	}

	message.SetBody("text/plain", messageBody)

	message.SetHeaders(map[string][]string{
		"From":    {message.FormatAddress("noreply@parvusjira.com", "Parvus JIRA")},
		"To":      receivers,
		"Subject": {"Parvus JIRA Notification"},
	})

	return message, nil
}

//BuildMessageFromProjectEvent builds the entire gomail.Message from the Project event, including body, subject, receivers, etc
func (m *mailService) BuildMessageFromProjectEvent(ctx context.Context, event *ProjectEvent) (*gomail.Message, error) {
	return nil, nil
}

//SendMail sends the mail using the injected mailServer
func (m *mailService) SendMail(ctx context.Context, message *gomail.Message) error {
	return m.mailServer.SendMail(ctx, message)
}

//userAllowed checks if the currently authenticated user can access/change the targeted user's settings
func (m *mailService) userAllowed(ctx context.Context, targetUser string) bool {
	return utils.GetUserFromContext(ctx) == targetUser
}

//SetUserPreference sets the notification preferences for the given user
func (m *mailService) SetUserPreference(ctx context.Context, userPreference *UserPreference) error {
	if m.userAllowed(ctx, userPreference.User) {
		return m.mailRepo.SetUserPreference(ctx, userPreference)
	} else {
		return errors.Wrap(utils.ErrNotAllowed, "service.Mailer.SetUserPreference")
	}
}

//GetUserPreference returns the notification preferences of the given user
func (m *mailService) GetUserPreference(ctx context.Context, user string) (*UserPreference, error) {
	if m.userAllowed(ctx, user) {
		return m.mailRepo.GetUserPreference(ctx, user)
	} else {
		return nil, errors.Wrap(utils.ErrNotAllowed, "service.Mailer.GetUserPreference")
	}
}
