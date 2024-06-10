package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listenAddr := ":8080" // Proxy server listens on port 8080

	// Start listening for incoming connections
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Proxy server listening on %s\n", listenAddr)

	// Accept and handle incoming connections
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(clientConn)
	}
}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	// Establish connection to the destination server
	serverConn, err := net.Dial("tcp", "https://www.google.com")
	if err != nil {
		fmt.Printf("Error connecting to destination server: %v\n", err)
		return
	}
	defer serverConn.Close()

	// Proxy data between client and server
	go copyData(clientConn, serverConn)
	go copyData(serverConn, clientConn)
}

func copyData(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Printf("Error copying data: %v\n", err)
	}
}
