FROM golang

# Copy source code into the container
COPY ./services/accounts /go/src/github.com/davidchristie/go-microservices/services/accounts
COPY ./libraries/testing/data /go/src/github.com/davidchristie/go-microservices/libraries/testing/data
COPY ./system-tests /go/src/github.com/davidchristie/go-microservices/system-tests

# Set the working directory
WORKDIR /go/src/github.com/davidchristie/go-microservices/system-tests

# Install dependencies
RUN go get ./...

# Run test tests
CMD [ "go", "test", "./...", "-count", "1"]
