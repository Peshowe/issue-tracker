package mailer

import (
	"context"
	"errors"
)

//IssueEventChannel is the name of the channel (or topic) from which we'll be listening for events for this subdomain
var IssueEventChannel string = "issue-channel"

//The different types of events we can have
var (
	IssueCreatedEventType = "IssueCreated"
	IssueUpdatedEventType = "IssueUpdated"
	IssueDeletedEventType = "IssueDeleted"
)

var ErrEventUnknown error = errors.New("Event Unknown")

//IssueEvent is the type of event received when something in the Issue subdomain has happened
type IssueEvent struct {
	Type  string `bson:"type,omitempty"`
	Issue struct {
		Id      string `bson:"_id,omitempty"`
		Name    string `bson:"name,omitempty"`
		User    string `bson:"user,omitempty"`
		Project string `bson:"project,omitempty"`
	} `bson:"issue,omitempty"`
}

//ProjectEventChannel is the name of the channel (or topic) from which we'll be listening for events for this subdomain
var ProjectEventChannel string = "project-channel"

//ProjectEvent is the type of event received when something in the Project subdomain has happened
type ProjectEvent struct {
	Type    string `bson:"type,omitempty"`
	Project struct {
		Id    string   `bson:"_id,omitempty"`
		Name  string   `bson:"name,omitempty"`
		Users []string `bson:"users,omitempty"`
	} `bson:"project,omitempty"`
}

type EventSubscriber interface {
	GetIssueEventChannel(context.Context) <-chan []byte
	GetProjectEventChannel(context.Context) <-chan []byte
	CloseChannels()
}
