package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	serverAddr := "localhost:" + strconv.Itoa(getServerPort())

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the server. Type your messages below:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		if message == "quit" {
			fmt.Println("Quitting the client.")
			break
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			continue
		}
	}

	if scanner.Err() != nil {
		fmt.Printf("Error reading from stdin: %v\n", scanner.Err())
	}
}


func getServerPort() int {
	var port int
	fmt.Print("Enter the server port: ")
	_, err := fmt.Scan(&port)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return port
}
