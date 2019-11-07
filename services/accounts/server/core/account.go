package core

import (
	"github.com/davidchristie/go-microservices/services/accounts/server/data"
	"github.com/google/uuid"
)

// Account is a core-layer representation of an account.
type Account interface {
	ID() uuid.UUID
	Name() string
	Email() string
}

type account struct {
	data *data.Account
}

func (a *account) ID() uuid.UUID {
	return (*a.data).ID()
}

func (a *account) Name() string {
	return (*a.data).Name()
}

func (a *account) Email() string {
	return (*a.data).Email()
}

func convertAccount(acc *data.Account) *Account {
	var a Account = &account{data: acc}
	return &a
}
