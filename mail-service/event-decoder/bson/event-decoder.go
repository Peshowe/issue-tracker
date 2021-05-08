package bson

import (
	"github.com/Peshowe/issue-tracker/mail-service/mailer"
	"go.mongodb.org/mongo-driver/bson"
)

type eventDecoder struct{}

func (e *eventDecoder) DecodeEvent(data []byte, event interface{}) error {
	return bson.Unmarshal(data, event)
}

func NewEventDecoder() mailer.EventDecoder {
	return &eventDecoder{}
}
