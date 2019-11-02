package gateway

import (
	"context"
	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/davidchristie/go-microservices/services/accounts/client"
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
	createdAccount, err := clientInstance.CreateAccount(context.Background(), &accounts.CreateAccountRequest{Email: "", Name: "", Password: ""})
	if err != nil {
		return nil, err
	}

	return &Account{ID: createdAccount.Id}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Account(ctx context.Context) (*Account, error) {
	return nil, nil
}
