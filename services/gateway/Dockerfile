# This stage compiles the server into a static binary
FROM golang AS build

# Copy source code into the container
COPY ./services /go/src/github.com/davidchristie/go-microservices/services

# Set the working directory
WORKDIR /go/src/github.com/davidchristie/go-microservices/services/gateway/server

# Install dependencies
RUN go get -v

# Compile the server into a static binary
RUN CGO_ENABLED=0 go build -o /server

# This stage runs the server
FROM scratch

# Copy the static binary that was created during the build stage
COPY --from=build /server /

# Set the static binary as the container entrypoint
ENTRYPOINT ["/server"]
