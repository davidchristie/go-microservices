package main

import (
	"github.com/davidchristie/go-microservices/services/accounts/server/api"
	"google.golang.org/grpc/grpclog"
)

var grpcFatalf = grpclog.Fatalf
var newHandlers = api.NewHandlers
var newListener = api.NewListener
var serve = api.Serve

func main() {
	listener, err := newListener()
	if err != nil {
		grpcFatalf("failed to listen: %v", err)
	}
	handlers := newHandlers()
	serve(listener, handlers)
}
