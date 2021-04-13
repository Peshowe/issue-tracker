package main

import (
	"fmt"
	"os"

	"github.com/Peshowe/issue-tracker/tracker-service/api/grpc"
	"github.com/Peshowe/issue-tracker/tracker-service/repository/mongo"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

func main() {
	mongoAddress := "mongodb://mongo:27017/test-mongo"
	if os.Getenv("MONGO_ADDRESS") != "" {
		mongoAddress = os.Getenv("MONGO_ADDRESS")
	}
	repo, err := mongo.NewMongoRepository(mongoAddress, "test-mongo", 5)
	if err != nil {
		fmt.Println(err)
	}

	projectService := project.NewProjectService(repo)
	issueService := issue.NewIssueService(repo)

	grpc.StartServer(projectService, issueService)

}
