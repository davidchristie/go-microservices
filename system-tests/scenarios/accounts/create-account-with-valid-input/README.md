# Create Account With Valid Input

When the `accounts` service receives a `CreateAccount` gRPC request with valid input:

1. It inserts a row in the `accounts` table of the accounts database.
2. It returns the newly created account in the gRPC response.
