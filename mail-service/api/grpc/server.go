package grpc

import (
	"context"
	"net"

	"github.com/Peshowe/issue-tracker/mail-service/grpc-contract/mail-service/v1/mailer"
	mlr "github.com/Peshowe/issue-tracker/mail-service/mailer"
	"github.com/Peshowe/issue-tracker/mail-service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

func StartServer(mailService mlr.MailService) {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))

	mailSrv := &mailServer{mailService: mailService}

	mailer.RegisterMailServiceServer(srv, mailSrv)

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
