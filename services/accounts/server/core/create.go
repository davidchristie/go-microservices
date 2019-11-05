package core

import (
	"context"
	"errors"

	db "github.com/davidchristie/go-microservices/services/accounts/server/data/repositories/accounts"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

// Account is the core representation of an account.
type Account struct {
	Email string
	ID    uuid.UUID
	Name  string
}

// CreateAccountInput is the input passed into the CreateAccount function.
type CreateAccountInput struct {
	Context  context.Context
	Email    string
	Name     string
	Password string
}

// CreateAccount attempts to create a new account.
func CreateAccount(input *CreateAccountInput) (*Account, error) {
	if input.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if input.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if input.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	id := uuid.New()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	account, err := db.CreateAccount(&db.CreateAccountInput{
		Context:      input.Context,
		Email:        input.Email,
		ID:           id,
		Name:         input.Name,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		return nil, err
	}

	return &Account{Email: account.Email, ID: account.ID, Name: account.Name}, nil
}
