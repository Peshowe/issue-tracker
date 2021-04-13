// package main

// import (
// 	"context"
// 	"fmt"
// 	// "github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/issue"
// 	"github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/project"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}

// 	client := project.NewProjectServiceClient(conn)

// 	req := &project.CreateRequest{Name: "Test project"}
// 	if _, err := client.CreateProject(context.Background(), req); err != nil {
// 		fmt.Println(err)
// 	}

// 	resp, err := client.GetProjectsAll(context.Background(), &project.ProjectsAllRequest{})
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(resp)

// }
