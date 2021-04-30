package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/Peshowe/issue-tracker/tracker-service/grpc-contract/tracker-service/v1/issue"
	"github.com/Peshowe/issue-tracker/tracker-service/grpc-contract/tracker-service/v1/project"
	"github.com/Peshowe/issue-tracker/tracker-service/utils"

	is "github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	pr "github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

// authInterceptor reads the user metadata from the incoming context, used for authorisation
func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(meta["token"]) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	if meta["token"][0] == "" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	ctx = context.WithValue(ctx, utils.TokenContexKey("token"), meta["token"][0])

	return handler(ctx, req)
}

func StartServer(projectService pr.ProjectService, issueService is.IssueService) {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))

	projectSrv := &projectServer{projectService: projectService}
	project.RegisterProjectServiceServer(srv, projectSrv)

	issueSrv := &issueServer{issueService: issueService}
	issue.RegisterIssueServiceServer(srv, issueSrv)

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
