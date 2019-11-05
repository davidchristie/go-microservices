package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}
