package issue

import (
	"context"
)

//IssueEventChannel is the name of the channel (or topic) to which we'll be publishing events for this subdomain
var IssueEventChannel string = "issue-channel"

//The different types of events we can have
var (
	IssueCreatedEventType = "IssueCreated"
	IssueUpdatedEventType = "IssueUpdated"
	IssueDeletedEventType = "IssueDeleted"
)

//IssueEvent is the event that gets published to the IssueEventChannel
type IssueEvent struct {
	Type  string `bson:"type,omitempty"`
	Issue *Issue `bson:"issue,omitempty"`
}

func (event IssueEvent) GetType() string {
	return event.Type
}

type EventPublisher interface {
	PublishEvent(context.Context, interface{}) error
}
