package client

import (
	"os"

	accounts "github.com/davidchristie/go-microservices/services/accounts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const defaultTarget = "accounts:5000"

func New() accounts.AccountsClient {
	target := os.Getenv("ACCOUNTS_URL")
	if target == "" {
		target = defaultTarget
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}

	return accounts.NewAccountsClient(conn)
}
