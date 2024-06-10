package main

import (
	"fmt"
	"log"

	"github.com/christopher-alden/sesi3/pkg/server"
)

func main() {
	srv := server.NewServer(":0")

	go func(){
		for msg:= range srv.MsgChan{
			fmt.Printf("received message %s: %s\n", msg.From,string(msg.Payload))
		}
	}()

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}