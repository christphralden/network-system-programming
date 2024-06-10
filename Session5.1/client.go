package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/christopher-alden/session3/client"
)

func main() {
	client := client.NewClient("localhost:9876")

	client.Dial()
	// client.SendMessage("hello nama saya budi")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		if message == "quit" {
			fmt.Println("Disconnecting...")
			client.Stop()
			break
		}

		err := client.SendMessage(message)
		if err != nil {
			fmt.Println("Error in sending message")
			continue
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Scanner error")
	}

	client.Stop()
}
