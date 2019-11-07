package core

import (
	"errors"

	"github.com/nbutton23/zxcvbn-go"
	"github.com/nbutton23/zxcvbn-go/scoring"
)

// ErrWeakPassword is returned by CreateAccount when the password is too weak.
var ErrWeakPassword = errors.New("weak password")

func validatePassword(input *CreateAccountInput) error {
	if input.Password == "" {
		return errors.New("empty password")
	}
	result := passwordStrength(input)
	if result.Score < 4 {
		return ErrWeakPassword
	}
	return nil
}

func passwordStrength(input *CreateAccountInput) scoring.MinEntropyMatch {
	userInput := []string{input.Email, input.Name}
	return zxcvbn.PasswordStrength(input.Password, userInput)
}
