package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Peshowe/issue-tracker/mail-service/api/grpc"
	"github.com/Peshowe/issue-tracker/mail-service/event-decoder/bson"
	"github.com/Peshowe/issue-tracker/mail-service/event-subscriber/kafka"
	"github.com/Peshowe/issue-tracker/mail-service/mail-repo/mongo"
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

	mongoAddress := "mongodb://localhost:27017"
	if os.Getenv("MONGO_ADDRESS") != "" {
		mongoAddress = os.Getenv("MONGO_ADDRESS")
	}
	mongoDB := "mail-service"
	if os.Getenv("MONGO_DB") != "" {
		mongoDB = os.Getenv("MONGO_DB")
	}
	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASS")

	mailServer := smtp.NewMailServer(smtpAddress, smtpPort)
	eventSubscriber := kafka.NewEventSubscriber(kafkaAddress)
	eventDecoder := bson.NewEventDecoder()
	mailRepo, err := mongo.NewMongoRepository(mongoAddress, mongoDB, mongoUser, mongoPass, 5)
	if err != nil {
		log.Fatalln(err)
	}

	mailService := mailer.NewMailService(mailServer, eventSubscriber, eventDecoder, mailRepo)

	// start the GRPC server in a goroutine
	go grpc.StartServer(mailService)

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
