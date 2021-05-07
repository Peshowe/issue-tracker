//This package contains logic for interfacing with the Kafka broker and publishing domain events
package kafka

import (
	"context"
	"log"
	"net"
	"strconv"

	"github.com/Peshowe/issue-tracker/tracker-service/tracker"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
	"github.com/Peshowe/issue-tracker/tracker-service/utils"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type eventPublisher struct {
	writer *kafka.Writer
}

//NewEventPublisher creates a new eventPublisher with a refrence to the Kafka broker
func NewEventPublisher(kafkaAddress string) (tracker.EventPublisher, *kafka.Writer, error) {

	// create the topics
	if err := createTopics(kafkaAddress, issue.IssueEventChannel, project.ProjectEventChannel); err != nil {
		return nil, nil, err
	}

	// make a writer that produces messages, using the least-bytes distribution (without specifying a topic here)
	addr := kafka.TCP(kafkaAddress)
	w := &kafka.Writer{
		Addr:     addr,
		Balancer: &kafka.LeastBytes{},
		// Topic:    issue.IssueEventChannel,
	}

	publisher := &eventPublisher{
		writer: w,
	}
	return publisher, w, nil
}

//PublishEvent publishes a given event to its respective channel (topic)
func (p *eventPublisher) PublishEvent(ctx context.Context, event interface{}) error {

	//message is the message we'll be creating and publishing
	var message kafka.Message

	//type switch to determine what event we're publishing and to build the correct message
	switch event := event.(type) {

	case issue.IssueEvent:
		message = kafka.Message{
			Topic: issue.IssueEventChannel,
			Key:   []byte(event.Issue.Id),
			Value: utils.EncodeToBytes(event),
		}
	case project.ProjectEvent:
		message = kafka.Message{
			Topic: project.ProjectEventChannel,
			Key:   []byte(event.Project.Id),
			Value: utils.EncodeToBytes(event),
		}
	default:
		log.Printf("Type %T is not a supported Event type!\n", event)
		return errors.Wrap(tracker.ErrEventUnknown, "event-publisher.PublishEvent")
	}

	// Publish the message
	err := p.writer.WriteMessages(ctx, message)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "event-publisher.PublishEvent")
	}

	log.Println("Published event: ", event.(tracker.DomainEvent).GetType(), string(message.Key))

	return nil
}

//createTopics creates the topics we'll be publishing to in the Kafka broker
func createTopics(kafkaAddress string, topics ...string) error {

	conn, err := kafka.Dial("tcp", kafkaAddress)
	if err != nil {
		return errors.Wrap(err, "event-publisher.createTopics")
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return errors.Wrap(err, "event-publisher.createTopics")
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return errors.Wrap(err, "event-publisher.createTopics")
	}
	defer controllerConn.Close()
	for _, topic := range topics {

		topicConfigs := []kafka.TopicConfig{
			kafka.TopicConfig{
				Topic:             topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err = controllerConn.CreateTopics(topicConfigs...)
		if err != nil {
			return errors.Wrap(err, "event-publisher.createTopics")
		}
	}

	return nil
}
