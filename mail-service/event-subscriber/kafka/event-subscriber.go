package kafka

import (
	"context"
	"log"

	"github.com/Peshowe/issue-tracker/mail-service/mailer"
	"github.com/segmentio/kafka-go"
)

type eventSubscriber struct {
	issueReader   *kafka.Reader
	projectReader *kafka.Reader
}

func NewEventSubscriber(kafkaAddress string) mailer.EventSubscriber {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	issueReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaAddress},
		Topic:     mailer.IssueEventChannel,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	projectReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaAddress},
		Topic:     mailer.ProjectEventChannel,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	return &eventSubscriber{
		issueReader,
		projectReader,
	}

}

func startListening(ctx context.Context, r *kafka.Reader) <-chan []byte {

	//channel to which we'll be writting the values of the received messages
	c := make(chan []byte)

	go func() {
		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				break
			}
			log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			c <- m.Value
		}
	}()

	return c
}

func (e *eventSubscriber) GetIssueEventChannel(ctx context.Context) <-chan []byte {
	return startListening(ctx, e.issueReader)
}

func (e *eventSubscriber) GetProjectEventChannel(ctx context.Context) <-chan []byte {
	return startListening(ctx, e.projectReader)
}
