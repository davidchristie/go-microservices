package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
)

// ErrNotFound is returned when a queried entity does not exist.
var ErrNotFound = errors.New("not found")

// QueryAccountByEmail attempts to find an account with the specified email.
// If the account does not exist a ErrNoRows error is returned.
func QueryAccountByEmail(ctx context.Context, email string) (*Account, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	const query = `
		SELECT account_id, name, email FROM accounts
		WHERE email = $1
	`

	row := db.QueryRow(query, email)
	var rowID uuid.UUID
	var rowEmail string
	var rowName string
	if err := row.Scan(&rowID, &rowName, &rowEmail); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return newAccount(rowID, rowEmail, rowName), nil
}
