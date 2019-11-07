package core

import (
	"errors"

	"github.com/davidchristie/go-microservices/services/accounts/server/data"
)

// ErrDuplicateEmail is returned by CreateAccount when an account with that email already exists.
var ErrDuplicateEmail = errors.New("duplicate email")

func validateEmail(input *CreateAccountInput) error {
	if input.Email == "" {
		return errors.New("empty email")
	}
	_, err := data.QueryAccountByEmail(input.Context, input.Email)
	if err == nil {
		return ErrDuplicateEmail
	}
	if err != data.ErrNotFound {
		return errors.New("unable to verify account email is unique")
	}
	return nil
}
