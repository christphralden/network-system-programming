package main

import (
	"context"
	"fmt"
	"net"
)

func streamingEchoServer(ctx context.Context, network string, addr string) (net.Addr, error) {
	ln, err := net.Listen(network, addr)

	if err != nil {
		return nil, err
	}

	go func() {

		go func() {
			<-ctx.Done()
			_ = ln.Close()
		}()
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go func() {
				defer conn.Close()

				for {
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						return
					}
					rtnMsg := []byte(fmt.Sprintf("Your ping: %v", string(buf[:n])))
					_, err = conn.Write(rtnMsg)

					if err != nil {
						return
					}
				}
			}()
		}
	}()

	return ln.Addr(), nil
}
