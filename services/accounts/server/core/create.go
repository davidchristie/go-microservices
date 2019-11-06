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
	if err := validate(input); err != nil {
		return nil, err
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

func validate(input *CreateAccountInput) error {
	validations := []func(*CreateAccountInput) error{validateEmail, validateName, validatePassword}
	for _, validation := range validations {
		if err := validation(input); err != nil {
			return err
		}
	}
	return nil
}

func validateEmail(input *CreateAccountInput) error {
	if input.Email == "" {
		return errors.New("empty email")
	}
	return nil
}

func validateName(input *CreateAccountInput) error {
	if input.Name == "" {
		return errors.New("empty name")
	}
	return nil
}
