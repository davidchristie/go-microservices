package api

import (
	"net"

	"github.com/davidchristie/go-microservices/services/accounts"
	"google.golang.org/grpc"
)

// Serve starts the server.
func Serve(listener net.Listener, handlers Handlers) error {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	accounts.RegisterAccountsServer(grpcServer, handlers)
	return grpcServer.Serve(listener)
}
