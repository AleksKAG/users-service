package grpc

import (
	"log"
	"net"

	userpb "github.com/AleksKAG/project-protos/proto/user"
	"github.com/AleksKAG/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.Service) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	grpcSrv := grpc.NewServer()
	handler := NewHandler(svc)
	userpb.RegisterUserServiceServer(grpcSrv, handler)
	log.Println("Server running at :50051")
	return grpcSrv.Serve(lis)
}
