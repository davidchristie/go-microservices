# Create Account With Unique Email

When the `accounts` service receives a `CreateAccount` gRPC request with a unique email:

1. It inserts a row in the `accounts` table of the accounts database.
2. It returns a copy of the newly created account in the gRPC response.
