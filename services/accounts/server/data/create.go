package data

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
)

// CreateAccountInput is the input passed into the CreateAccount function.
type CreateAccountInput struct {
	Context      context.Context
	Email        string
	ID           uuid.UUID
	Name         string
	PasswordHash string
}

// CreateAccount attempts to create a new account in the repository.
func CreateAccount(input *CreateAccountInput) (*Account, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tx, err := db.BeginTx(input.Context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	const query = `
		INSERT INTO accounts (account_id, email, name, password_hash)
		VALUES ($1, $2, $3, $4);
	`

	_, err = tx.Exec(query, input.ID, input.Email, input.Name, input.PasswordHash)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return newAccount(input.ID, input.Name, input.Email), nil
}
