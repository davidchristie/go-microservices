package data

import "github.com/google/uuid"

// Account is the data-layer representation of an account.
type Account interface {
	ID() uuid.UUID
	Name() string
	Email() string
}

type account struct {
	id    uuid.UUID
	name  string
	email string
}

func (s *account) ID() uuid.UUID {
	return s.id
}

func (s *account) Name() string {
	return s.name
}

func (s *account) Email() string {
	return s.email
}

// New creates a new Account instance.
func newAccount(id uuid.UUID, name string, email string) *Account {
	var a Account = &account{
		id:    id,
		name:  name,
		email: email,
	}
	return &a
}
