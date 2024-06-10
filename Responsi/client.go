package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/christopher-alden/responsi/client"
)

func main() {

	client := client.NewClient("localhost:5123")

	client.Dial()
	// aksi kita mau ngapain

	fmt.Println("Input your messages:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		if message == "quit" {
			fmt.Println("Disconnecting from server")
			client.Stop()
			break
		}

		err := client.SendMessage(message)

		if err != nil {
			fmt.Println("Error in sending message")
			continue
		}
	}

	client.Stop()
}
