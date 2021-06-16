package main

import (
	"log"
	"os"

	"github.com/Peshowe/issue-tracker/tracker-service/api/grpc"
	"github.com/Peshowe/issue-tracker/tracker-service/event-publisher/kafka"
	"github.com/Peshowe/issue-tracker/tracker-service/repository/mongo"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

func main() {
	mongoAddress := "mongodb://localhost:27017"
	if os.Getenv("MONGO_ADDRESS") != "" {
		mongoAddress = os.Getenv("MONGO_ADDRESS")
	}
	mongoDB := "test-mongo"
	if os.Getenv("MONGO_DB") != "" {
		mongoDB = os.Getenv("MONGO_DB")
	}
	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASS")
	repo, err := mongo.NewMongoRepository(mongoAddress, mongoDB, mongoUser, mongoPass, 5)
	if err != nil {
		log.Fatalln(err)
	}

	kafkaAddress := "localhost:9092"
	if os.Getenv("KAFKA_ADDRESS") != "" {
		kafkaAddress = os.Getenv("KAFKA_ADDRESS")
	}
	eventPublisher, kafkaWritter, err := kafka.NewEventPublisher(kafkaAddress)
	if err != nil {
		log.Fatalln(err)
	}
	defer kafkaWritter.Close()

	projectService := project.NewProjectService(repo, eventPublisher)
	issueService := issue.NewIssueService(repo, eventPublisher, projectService)

	grpc.StartServer(projectService, issueService)

}
