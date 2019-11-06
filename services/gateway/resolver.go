package gateway

import (
	"context"
	"errors"

	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/davidchristie/go-microservices/services/accounts/client"
	"google.golang.org/grpc/status"
)

type Resolver struct{}

var clientInstance = client.New()

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAccount(ctx context.Context, input CreateAccount) (*Account, error) {
	newAccount, err := clientInstance.CreateAccount(context.Background(), &accounts.CreateAccountRequest{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, errors.New(st.Message())
		}
		return nil, errors.New("Unknown error occured")
	}

	return &Account{Email: newAccount.Email, ID: newAccount.Id, Name: newAccount.Name}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Account(ctx context.Context) (*Account, error) {
	return nil, nil
}
