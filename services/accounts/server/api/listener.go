package api

import (
	"net"
	"os"
)

const defaultPort = "5000"

// NewListener creates a new listener
func NewListener() (net.Listener, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return net.Listen("tcp", ":"+port)
}
