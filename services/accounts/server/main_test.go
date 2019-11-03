package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rc := m.Run()

	// rc 0 means the tests have passed
	// CoverMode will be non-empty if run with the -cover flag
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		// Require 100% coverage
		if c < 1 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}

	os.Exit(rc)
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
	original := serve
	called := false
	serve = func(listener net.Listener) error {
		called = true
		return nil
	}
	defer func() {
		serve = original
	}()
	main()
	if !called {
		t.Fatal("Serve function was not called")
	}
}
