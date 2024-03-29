package api

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type mockHandlers struct {
	Handlers
}

func (m mockHandlers) CreateAccount(ctx context.Context, request *accounts.CreateAccountRequest) (*accounts.CreateAccountResponse, error) {
	response := &accounts.CreateAccountResponse{Email: "mock@email.com", Id: uuid.New().String(), Name: "Mock"}
	return response, nil
}

func startServer(listener net.Listener) {
	go func() {
		if err := Serve(listener, mockHandlers{}); err != nil {
			log.Printf("Server exited with error: %v", err)
		}
	}()
}

func TestCreateAccount(t *testing.T) {
	listener := bufconn.Listen(bufSize)
	defer listener.Close()

	startServer(listener)
	ctx := context.Background()

	bufDialer := func(string, time.Duration) (net.Conn, error) {
		return listener.Dial()
	}

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := accounts.NewAccountsClient(conn)
	resp, err := client.CreateAccount(ctx, &accounts.CreateAccountRequest{
		Email:    "test.user@email.com",
		Name:     "Test User",
		Password: "test_p@ssword_123",
	})
	if err != nil {
		t.Fatalf("CreateAccount failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	_, err = uuid.Parse(resp.Id)
	if err != nil {
		t.Fatalf("Invalid account ID: %v", err)
	}
}
