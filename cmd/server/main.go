package main

import (
	"log"

	"github.com/AleksKAG/users-service/internal/database"
	transportgrpc "github.com/AleksKAG/users-service/internal/transport/grpc"
	"github.com/AleksKAG/users-service/internal/user"
)

func main() {
	log.Println("Starting server initialization...")
	database.InitDB()
	log.Println("Database initialized successfully")
	repo := user.NewRepository(database.DB)
	log.Println("Repository created")
	svc := user.NewService(repo)
	log.Println("Service created")
	log.Println("Starting gRPC server on :50051...")
	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("Users gRPC server error: %v", err)
	}
}
