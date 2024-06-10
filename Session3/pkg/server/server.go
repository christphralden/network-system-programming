package server

import (
	"fmt"
	"io"
	"net"
)

type Message struct {
	From    string
	Payload []byte
}
type Server struct {
	ListenAddress string
	Ln            net.Listener
	QuitChan      chan struct{}
	MsgChan       chan Message
}

func NewServer(listenAddress string) *Server {
	return &Server{
		ListenAddress: listenAddress,
		QuitChan:      make(chan struct{}),
		MsgChan:       make(chan Message, 10),
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

	fmt.Println("Server listening on port", port)

	go s.AcceptLoop()

	<-s.QuitChan
	// if close server close the msg chan
	//cleanup, Closing the channel, notify all the people reading the channel that its closed gracefully
	close(s.MsgChan)

	return nil
}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.Ln.Accept()

		if err != nil {
			fmt.Print("Error", err)
			continue
		}

		fmt.Println("new connection to the server", conn.RemoteAddr())
		//each time we acc, then each time we gonna make a new goroutine to handle the new connection
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected", conn.RemoteAddr().String())
			} else {
				fmt.Println("Error reading from connection:", err)
			}
			break // Exit the loop on EOF or any other error
		}

		// read as much as the amoun that is read
		// msg := buf[:n]
		// fmt.Println(string(msg))
		s.MsgChan <- Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf[:n],
		}
	}
}
