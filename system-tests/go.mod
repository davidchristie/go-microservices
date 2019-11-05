module system-tests

go 1.13

require (
	github.com/davidchristie/go-microservices/libraries/testing/data v0.0.0-00010101000000-000000000000
	github.com/davidchristie/go-microservices/services/accounts v0.0.0-20191103083057-24a9009a4ca6
	github.com/google/uuid v1.1.1
	github.com/lib/pq v1.2.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)

replace github.com/davidchristie/go-microservices/libraries/testing/data => ../libraries/testing/data

replace github.com/davidchristie/go-microservices/services/accounts => ../services/accounts
