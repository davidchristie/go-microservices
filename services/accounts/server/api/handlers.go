package api

import (
	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/davidchristie/go-microservices/services/accounts/server/core"
	"golang.org/x/net/context"
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
		return nil, err
	}

	response := &accounts.CreateAccountResponse{Email: account.Email, Id: account.ID.String(), Name: account.Name}
	return response, nil
}
