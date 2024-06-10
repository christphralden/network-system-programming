package main

import (
	"fmt"
	"io"
	"net"
)

func proxyFroward(from io.Reader, to io.Writer) error {

	fromWriter, fromIsWriter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	if fromIsWriter && toIsReader {
		go func(){
			io.Copy(fromWriter, toReader)
		}()
	}
	_, err := io.Copy(to, from)

	return err
}

func main(){
	// dia dengerin di port 5123
	ln, err := net.Listen("tcp", "localhost:5123")

	if err !=nil{
		fmt.Println("Error", err)
		return
	}

	for{
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		go func(from net.Conn){
			// dia mau nerusin ke servernya
			to, err := net.Dial("tcp", "localhost:1234")

			if err != nil {
				fmt.Println("Error", err)
				return
			}

			//proxyForward
			err = proxyFroward(from, to)

			if err != nil {
				fmt.Println("Error", err)
				return
			}

		}(conn)
	}

}