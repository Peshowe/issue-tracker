package main

import (
	"fmt"
	"github.com/Peshowe/issue-tracker/tracker-service/api/grpc"
	"github.com/Peshowe/issue-tracker/tracker-service/repository/mongo"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

func main() {
	repo, err := mongo.NewMongoRepository("mongodb://mongo:27017/test-mongo", "test-mongo", 5)
	if err != nil {
		fmt.Println(err)
	}

	projectService := project.NewProjectService(repo)
	issueService := issue.NewIssueService(repo)

	grpc.StartServer(projectService, issueService)

}
