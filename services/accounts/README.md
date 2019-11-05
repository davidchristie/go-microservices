# Accounts

An internal service for managing user accounts.

## Scripts

Generate protocol buffer source code:

```console
protoc -I . accounts.proto --go_out=plugins=grpc:.
```
