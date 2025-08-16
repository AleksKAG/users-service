package grpc

import (
	"net"

	userpb "github.com/AleksKAG/project-protos/proto/user"
	"github.com/AleksKAG/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.Service) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))
	return grpcServer.Serve(listener)
}
