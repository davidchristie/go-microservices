package main

import (
	"errors"
	"net"
	"testing"

	"github.com/davidchristie/go-microservices/services/accounts/server/api"
)

type mockListener struct {
	net.Listener
}

func TestNetworkError(t *testing.T) {
	originalGrpcFatalf := grpcFatalf
	originalNewListener := newListener

	grpcFatalf = func(format string, args ...interface{}) {
		t.Skip()
	}
	newListener = func() (net.Listener, error) {
		return nil, errors.New("<TEST_NETWORK_ERROR>")
	}

	defer func() {
		grpcFatalf = originalGrpcFatalf
		newListener = originalNewListener
	}()

	main()
	t.Fatal("grpc.Fatalf was not called")
}

func TestServe(t *testing.T) {
	originalNewListener := newListener

	newListener = func() (net.Listener, error) {
		return mockListener{}, nil
	}

	originalServe := serve
	called := false
	serve = func(listener net.Listener, handlers api.Handlers) error {
		called = true
		return nil
	}

	defer func() {
		serve = originalServe
		newListener = originalNewListener

	}()
	main()
	if !called {
		t.Fatal("serve function was not called")
	}
}
