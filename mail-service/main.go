package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Peshowe/issue-tracker/mail-service/event-decoder/bson"
	"github.com/Peshowe/issue-tracker/mail-service/event-subscriber/kafka"
	"github.com/Peshowe/issue-tracker/mail-service/mail-server/smtp"
	"github.com/Peshowe/issue-tracker/mail-service/mailer"
)

func main() {
	kafkaAddress := "localhost:9092"
	if os.Getenv("KAFKA_ADDRESS") != "" {
		kafkaAddress = os.Getenv("KAFKA_ADDRESS")
	}

	smtpAddress := "localhost"
	if os.Getenv("SMTP_ADDRESS") != "" {
		smtpAddress = os.Getenv("SMTP_ADDRESS")
	}
	smtpPort := 25
	if os.Getenv("SMTP_PORT") != "" {
		smtpPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	}

	mailServer := smtp.NewMailServer(smtpAddress, smtpPort)
	eventSubscriber := kafka.NewEventSubscriber(kafkaAddress)
	eventDecoder := bson.NewEventDecoder()

	mailService := mailer.NewMailService(mailServer, eventSubscriber, eventDecoder)

	// done channel to stop listening for events
	done := make(chan bool)

	// c is a channel where we'll be listening for OS interupt signals
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		//send a bool to the done channel to stop the subsriber listening
		done <- true
		os.Exit(1)
	}()

	mailService.ListenForEvents(context.Background(), done)

}
