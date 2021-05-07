package tracker

import (
	"errors"

	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

var ErrEventUnknown error = errors.New("Event Unknown")

type DomainEvent interface {
	GetType() string
}

type EventPublisher interface {
	issue.EventPublisher
	project.EventPublisher
}
