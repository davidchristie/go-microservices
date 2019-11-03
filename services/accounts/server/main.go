package main

import (
	"google.golang.org/grpc/grpclog"
)

var grpcFatalf = grpclog.Fatalf
var newListener = NewListener
var serve = Serve

func main() {
	listener, err := newListener()
	if err != nil {
		grpcFatalf("failed to listen: %v", err)
	}
	serve(listener)
}
