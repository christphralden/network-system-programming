package server

import (
	"fmt"
	"io"
	"net"

	"github.com/christopher-alden/session3/types"
)

type Server struct {
	ListenAddress string
	Ln            net.Listener
	QuitChan      chan struct{}
}

func NewServer(listenAddress string) *Server {
	return &Server{
		ListenAddress: listenAddress,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddress)

	if err != nil {
		return err
	}
	defer ln.Close()
	s.Ln = ln
	_, port, _ := net.SplitHostPort(ln.Addr().String())
	fmt.Println("listening on port:", port)

	go s.AcceptLoop()

	<-s.QuitChan

	return nil
}

func (s *Server) Stop() error {
	if s.Ln != nil {
		s.Ln.Close()
	}
	return nil
}

func (s *Server) AcceptLoop() error {
	for {
		conn, err := s.Ln.Accept()

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client has disconnected", conn.RemoteAddr().String())
			} else {
				fmt.Printf("Error: %s\n", err)
			}
			// break ini penting buat keluarin dari loop
			continue
		} else {
			go s.readLoop(conn)
		}

	}
}

func (s *Server) readLoop(conn net.Conn) {
	// fmt.Println("Message Accepted")
	defer conn.Close()

	for {
		payload, err := types.Decode(conn)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client has disconnected", conn.RemoteAddr().String())
			} else {
				fmt.Printf("Error: %s\n", err)
			}
			// break ini penting buat keluarin dari loop
			break
		}

		// contohin tanpa bytes dulu
		fmt.Println(string(payload.Bytes()))

		p := types.Binary("Siap")

		// ini kalo dia mau write response balik
		_, err = p.WriteTo(conn)

		if err != nil {
			fmt.Println("error", err)
			break
		}

	}
	/*
		buf := make([]byte, 4)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Client has disconnected", conn.RemoteAddr().String())
				}else{
					fmt.Printf("Error: %s\n", err)
				}
				break
			}
			fmt.Printf("read %d bytes: %s\n", n, buf[:n])

		}

	*/
}
