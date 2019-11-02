module github.com/davidchristie/go-microservices/services/gateway

go 1.13

require (
	github.com/99designs/gqlgen v0.10.1
	github.com/davidchristie/go-microservices/services/accounts v0.0.0
	github.com/vektah/gqlparser v1.1.2
)

replace github.com/davidchristie/go-microservices/services/accounts => ../accounts
