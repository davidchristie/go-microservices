package data

import (
	"github.com/google/uuid"
)

// NewUniqueEmail creates a new unique email or panics.
func NewUniqueEmail() string {
	return "test.user+" + uuid.New().String() + "@email.com"
}
