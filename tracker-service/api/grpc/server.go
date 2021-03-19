package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/issue"
	"github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/project"

	is "github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	pr "github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

func StartServer(projectService pr.ProjectService, issueService is.IssueService) {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	projectSrv := &projectServer{projectService: projectService}
	project.RegisterProjectServiceServer(srv, projectSrv)

	issueSrv := &issueServer{issueService: issueService}
	issue.RegisterIssueServiceServer(srv, issueSrv)

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
