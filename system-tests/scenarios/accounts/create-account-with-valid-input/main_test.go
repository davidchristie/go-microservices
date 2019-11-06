package scenario

import (
	"context"
	"database/sql"
	"testing"

	"github.com/davidchristie/go-microservices/libraries/testing/data"
	"github.com/davidchristie/go-microservices/services/accounts"
	"github.com/davidchristie/go-microservices/services/accounts/client"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
	"golang.org/x/crypto/bcrypt"
)

const accountsDatabaseURL = "postgres://accounts:acc0unts_secret123@accounts-database:5432/accounts?sslmode=disable"

func TestScenario(t *testing.T) {
	// Send the CreateAccount gRPC request
	const name = "Test User"
	const password = "2J3xtr_Z9xnobGQPJkDW"
	email := data.NewUniqueEmail()
	clientInstance := client.New()
	response, err := clientInstance.CreateAccount(context.Background(), &accounts.CreateAccountRequest{
		Email:    email,
		Name:     name,
		Password: password,
	})
	if err != nil {
		t.Fatalf("error response from accounts service: %s", err)
	}

	// Verify the gRPC response
	t.Logf("response id: %v", response.Id)
	t.Logf("response email: %v", response.Email)
	t.Logf("response name: %v", response.Name)
	if response.Email != email {
		t.Fatalf("response email does not match input: expected=%s, actual=%s", email, response.Email)
	}
	if response.Name != name {
		t.Fatalf("respone name does not match input: expected=%s, actual=%s", name, response.Name)
	}

	// Verify the account was inserted into the database
	db, err := sql.Open("postgres", accountsDatabaseURL)
	if err != nil {
		t.Fatalf("could not connect to database: %s", err)
	}
	defer db.Close()
	const query = `
		SELECT * FROM accounts
		WHERE account_id = $1
		AND email = $2
	`
	row := db.QueryRow(query, response.Id, email)
	var rowAccountID uuid.UUID
	var rowEmail string
	var rowName string
	var rowPasswordHash string
	switch err := row.Scan(&rowAccountID, &rowEmail, &rowName, &rowPasswordHash); err {
	case sql.ErrNoRows:
		t.Fatal("no rows were returned")
	case nil:
		t.Log("row was inserted into the database")
	default:
		panic(err)
	}

	// Verify the data saved in the database
	t.Logf("row account_id: %v", rowAccountID)
	t.Logf("row email: %v", rowEmail)
	t.Logf("row name: %v", rowName)
	t.Logf("row password_hash: %v", rowPasswordHash)
	if rowAccountID.String() != response.Id {
		t.Fatalf("row account_id does not match response: expected=%s, actual=%s", response.Id, rowName)
	}
	if rowEmail != email {
		t.Fatalf("row account_id does not match response: expected=%s, actual=%s", response.Id, rowName)
	}
	if rowName != name {
		t.Fatalf("row name does not match input: expected=%s, actual=%s", name, rowName)
	}
	err = bcrypt.CompareHashAndPassword([]byte(rowPasswordHash), []byte(password))
	if err != nil {
		t.Fatal("password hash does not match input")
	}
}
