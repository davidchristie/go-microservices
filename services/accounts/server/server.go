package main

import (
	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	accounts.RegisterAccountsServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) CreateAccount(ctx context.Context, request *accounts.CreateAccountRequest) (*accounts.CreateAccountResponse, error) {
	id := uuid.New()
	response := &accounts.CreateAccountResponse{Id: id.String()}
	return response, nil
}
