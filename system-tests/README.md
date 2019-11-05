# System Tests

## Scripts

Run the tests:

```console
docker run \
    --network production_default \
    --rm \
    --volume "$(pwd)/..:/go/src/github.com/davidchristie/go-microservices" \
    --workdir /go/src/github.com/davidchristie/go-microservices/system-tests \
    golang \
    go test ./...
```
