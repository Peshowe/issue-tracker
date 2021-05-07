package project

import (
	"context"
)

var ProjectEventChannel string = "project-channel"

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
