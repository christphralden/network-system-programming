package server

// net, io, bufio, fmt, time

import (
	"fmt"
	"io"
	"net"

	"github.com/christopher-alden/responsi/types"
)

type Server struct {
	ListenAddress string        // buat simpen address kita
	Ln            net.Listener  // buat listener kita
	QuitChan      chan struct{} // buat penanda server close
}

// buat ngebikin server
func NewServer(listenAddress string) *Server {
	return &Server{
		ListenAddress: listenAddress,
	}
}

// start server
func (s *Server) Start() error {
	ln, err := net.Listen("tcp",s.ListenAddress)

	if err != nil {
		return err
	}

	// close connection gracefully
	defer ln.Close()
	s.Ln = ln

	_, port, _ := net.SplitHostPort(ln.Addr().String())

	fmt.Println("Listening on port:", port)

	//loop pertama untuk denger connection yang masuk
		// loop kedua itu untuk melaksanakan action yang di inginkan oleh connection

	//go routine -> concurrency
	go s.AcceptLoop()

	// nandain bahwa server kita mau dimatiin
	<- s.QuitChan


	return nil
}


func (s *Server) AcceptLoop() error {
	fmt.Println("Accept Connection")
	for{
		conn, err := s.Ln.Accept()

		if err != nil{
			if err == io.EOF{
				fmt.Println("Client has disconnected", conn.RemoteAddr().String())
			}else{
				fmt.Println("Error", err)
			}
			//continue penting
			continue
		}

		// do something
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	// asdjajsdajsdnajsdnjasdn
	// buf := make([]byte, 64)

	for{
		// static allocation : mempunyai buffer yang fixed size
		// dynamic allocation : allocate buffer size based on payload nya
		payload, err := types.Decode(conn)

		if err != nil{
			if err == io.EOF{
				fmt.Println("Client has disconnected", conn.RemoteAddr().String())
			}else{
				fmt.Println("Error", err)
			}
			//break peting
			break
		}
		fmt.Println(string(payload.Bytes()))

		reply := types.Binary("Read")

		_, err = reply.WriteTo(conn)

		if err !=nil{
			fmt.Println("Error", err)
			break
		}
	}

}