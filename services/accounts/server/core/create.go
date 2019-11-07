package core

import (
	"context"
	"errors"

	"github.com/davidchristie/go-microservices/services/accounts/server/data"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

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

	acc, err := data.CreateAccount(&data.CreateAccountInput{
		Context:      input.Context,
		Email:        input.Email,
		ID:           id,
		Name:         input.Name,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		return nil, err
	}

	return convertAccount(acc), nil
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

func validateName(input *CreateAccountInput) error {
	if input.Name == "" {
		return errors.New("empty name")
	}
	return nil
}
