package client

import (
	"fmt"
	"net"

	"github.com/christopher-alden/session3/types"
)

type Client struct {
	ConnectionAddress string
	Conn              net.Conn
}

func NewClient(connectionAddress string) *Client {
	return &Client{
		ConnectionAddress: connectionAddress,
	}
}

func (c *Client) Dial() error {
	// timeoutDuration := 5 * time.Second

	// // Use net.DialTimeout instead of net.Dial
	// dial, err := net.DialTimeout("tcp", c.ConnectionAddress, timeoutDuration)
	// c.Conn = dial
	dial, err := net.Dial("tcp", c.ConnectionAddress)
	c.Conn = dial

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendMessage(msg string) error {
	// _, err := c.Conn.Write([]byte(msg))

	payload := types.Binary(msg)

	_, err := payload.WriteTo(c.Conn)

	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
		return err
	}

	p, err := types.Decode(c.Conn)

	if err != nil {
		fmt.Printf("Error reading message: %v\n", err)
		return err
	}

	fmt.Println(string(p.Bytes()))
	return nil
}

func (s *Client) Stop() error {
	if s.Conn != nil {
		s.Conn.Close()
	}
	return nil
}
