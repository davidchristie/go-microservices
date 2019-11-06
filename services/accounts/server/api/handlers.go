package api

import (
	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/davidchristie/go-microservices/services/accounts/server/core"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handlers interface {
	CreateAccount(ctx context.Context, request *accounts.CreateAccountRequest) (*accounts.CreateAccountResponse, error)
}

type handlers struct{}

var h Handlers = handlers{}

func NewHandlers() Handlers {
	return handlers{}
}

func (h handlers) CreateAccount(ctx context.Context, request *accounts.CreateAccountRequest) (*accounts.CreateAccountResponse, error) {
	account, err := core.CreateAccount(&core.CreateAccountInput{Context: ctx, Email: request.Email, Name: request.Name, Password: request.Password})
	if err != nil {
		switch err.(type) {
		case *core.WeakPasswordError:
			return nil, status.Error(codes.InvalidArgument, "Password was too weak")
		default:
			return nil, status.Error(codes.Unknown, "Unknown error occurred")
		}
	}

	response := &accounts.CreateAccountResponse{Email: account.Email, Id: account.ID.String(), Name: account.Name}
	return response, nil
}
