package main

import (
	"log"

	"github.com/christopher-alden/session3/server"
)

func main() {
	server := server.NewServer("localhost:1234")

	if err := server.Start(); err != nil {
		log.Fatal("server failed to start")
	}
}
