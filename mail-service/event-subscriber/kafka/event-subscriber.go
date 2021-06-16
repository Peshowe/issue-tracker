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
	// make new readers for each topic
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
				log.Println(err)
				break
			}
			log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			c <- m.Value
		}
		close(c)
	}()

	return c
}

func (e *eventSubscriber) GetIssueEventChannel(ctx context.Context) <-chan []byte {
	return startListening(ctx, e.issueReader)
}

func (e *eventSubscriber) GetProjectEventChannel(ctx context.Context) <-chan []byte {
	return startListening(ctx, e.projectReader)
}

func (e *eventSubscriber) CloseChannels() {
	if err := e.issueReader.Close(); err != nil {
		log.Println("failed to close issue reader:", err)
	}
	if err := e.projectReader.Close(); err != nil {
		log.Println("failed to close project reader:", err)
	}
}
