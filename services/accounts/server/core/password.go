package core

import (
	"errors"

	"github.com/nbutton23/zxcvbn-go"
	"github.com/nbutton23/zxcvbn-go/scoring"
)

// WeakPasswordError is produced when a weak password is validated.
type WeakPasswordError struct{}

func (e *WeakPasswordError) Error() string {
	return "weak password"
}

func validatePassword(input *CreateAccountInput) error {
	if input.Password == "" {
		return errors.New("empty password")
	}
	result := passwordStrength(input)
	if result.Score < 4 {
		return &WeakPasswordError{}
	}
	return nil
}

func passwordStrength(input *CreateAccountInput) scoring.MinEntropyMatch {
	userInput := []string{input.Email, input.Name}
	return zxcvbn.PasswordStrength(input.Password, userInput)
}
