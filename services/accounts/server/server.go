package main

import (
	"net"

	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

// Serve starts the server.
func Serve(listener net.Listener) error {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	server := &server{}
	accounts.RegisterAccountsServer(grpcServer, server)
	return grpcServer.Serve(listener)
}

func (s *server) CreateAccount(ctx context.Context, request *accounts.CreateAccountRequest) (*accounts.CreateAccountResponse, error) {
	id := uuid.New()
	response := &accounts.CreateAccountResponse{Id: id.String()}
	return response, nil
}
