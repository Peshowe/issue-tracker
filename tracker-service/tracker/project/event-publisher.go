package project

import (
	"context"
)

//ProjectEventChannel is the name of the channel (or topic) to which we'll be publishing events for this subdomain
var ProjectEventChannel string = "project-channel"

//ProjectEvent is the event that gets published to the ProjectEventChannel
type ProjectEvent struct {
	Type    string   `bson:"type,omitempty"`
	Project *Project `bson:"project,omitempty"`
}

func (event ProjectEvent) GetType() string {
	return event.Type
}

type EventPublisher interface {
	PublishEvent(context.Context, interface{}) error
}
