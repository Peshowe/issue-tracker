package issue

import (
	"context"
)

var IssueEventChannel string = "issue-channel"

var (
	IssueCreatedEventType = "IssueCreated"
	IssueUpdatedEventType = "IssueUpdated"
	IssueDeletedEventType = "IssueDeleted"
)

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
