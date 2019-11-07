package scenario

import (
	"context"
	"testing"

	"github.com/davidchristie/go-microservices/libraries/testing/data"
	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/davidchristie/go-microservices/services/accounts/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestScenario(t *testing.T) {
	clientInstance := client.New()
	email := data.NewUniqueEmail()

	// Send first CreateAccount gRPC request with a unique email
	_, err := clientInstance.CreateAccount(context.Background(), &accounts.CreateAccountRequest{
		Email:    email,
		Name:     "First User",
		Password: "gUGqyEfWjh3wvK4!QFmr",
	})
	if err != nil {
		t.Fatalf("error creating first user: %v", err)
	}

	// Send second CreateAccount gRPC request with the same email
	response, err := clientInstance.CreateAccount(context.Background(), &accounts.CreateAccountRequest{
		Email:    email,
		Name:     "Second User",
		Password: "aYP!9im6nCTQuge-4aRb",
	})
	if err == nil {
		t.Logf("response: %v", response)
		t.Fatalf("no error creating user with duplicate email: %v", err)
	}

	// Verify the gRPC error
	t.Logf("error: %v", err)
	status := status.Convert(err)
	code := status.Code()
	message := status.Message()
	details := status.Details()
	t.Logf("status code: %v", status)
	t.Logf("status message: %v", message)
	t.Logf("status details: %v", details)
	const expectedCode = codes.AlreadyExists
	const expectedMessage = "An account with that email already exists"
	const expectedDetailsLength = 0
	if code != expectedCode {
		t.Fatalf("incorrect code: expected=%v, actual=%v", expectedCode, code)
	}
	if message != expectedMessage {
		t.Fatalf("incorrect message: expected=%v, actual=%v", expectedMessage, message)
	}
	if len(details) != expectedDetailsLength {
		t.Fatalf("incorrect details length: expected=%v, actual=%v", expectedDetailsLength, len(details))
	}
}
