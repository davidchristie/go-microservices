package scenario

import (
	"context"
	"testing"

	"github.com/davidchristie/go-microservices/libraries/testing/data"
	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/davidchristie/go-microservices/services/accounts/client"
	_ "github.com/lib/pq" // postgres driver
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestScenario(t *testing.T) {
	// Send the CreateAccount gRPC request
	const name = "Test User"
	const password = "test_u$er123"
	email := data.NewUniqueEmail()
	clientInstance := client.New()
	response, err := clientInstance.CreateAccount(context.Background(), &accounts.CreateAccountRequest{
		Email:    email,
		Name:     name,
		Password: password,
	})
	if err == nil {
		t.Logf("response: %v", response)
		t.Fatalf("no error from accounts service")
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
	const expectedCode = codes.InvalidArgument
	const expectedMessage = "Password was too weak"
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
