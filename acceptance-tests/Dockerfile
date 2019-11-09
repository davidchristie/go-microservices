FROM golang

# Copy source code into the container
COPY ./acceptance-tests /go/src/github.com/davidchristie/go-microservices/acceptance-tests

# Set the working directory
WORKDIR /go/src/github.com/davidchristie/go-microservices/acceptance-tests

# Install dependencies
RUN go get ./...

# Run test tests
CMD [ "go", "test", "./...", "-count", "1"]
