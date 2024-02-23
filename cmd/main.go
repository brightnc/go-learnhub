package main

import (
	"log"

	"github.com/brightnc/go-learnhub/protocol"
)

func main() {
	cleanup, err := protocol.Initialize()
	if err != nil {
		log.Fatal("Error initializing application:", err)
	}
	defer cleanup()
	protocol.ServeREST()
}
