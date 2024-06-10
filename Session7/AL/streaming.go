package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

func main() {

	dir, err := ioutil.TempDir("", "echo_unix")
	// fmt.Println(dir)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if rErr := os.RemoveAll(dir); rErr != nil {
			fmt.Println(rErr)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))

	rAddr, err := streamingEchoServer(ctx, "unix", socket)

	if err != nil {
		fmt.Println(err)
	}

	defer cancel()

	err = os.Chmod(socket, os.ModeSocket|0666)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.Dial("unix", rAddr.String())

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	msg := []byte("ping")

	for i := 0; i < 3; i++ {
		_, err := conn.Write(msg)

		if err != nil {
			fmt.Println(err)
		}
	}

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
	}

	expectedServer := bytes.Repeat(msg, 3)

	if n != 0 {
		fmt.Println(string(buf[:n]))
	}

	if !bytes.Equal(expectedServer, buf[:n]) {
		fmt.Printf("recieved reply: %q, expected reply:%q\n", buf[:n], expectedServer)
	}
}
