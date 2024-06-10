package client

import (
	"fmt"
	"net"
	"time"

	"github.com/christopher-alden/responsi/types"
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

// dial the server
func (c *Client) Dial() error {
	// DIAL BIASA
	// conn, err := net.Dial("tcp", c.ConnectionAddress)

	// DIAL TIMEOUT : kalau dia tidak connect ke server dalam durasi timeout dia bakal close connection
	timeout := time.Second * 5
	conn, err := net.DialTimeout("tcp", c.ConnectionAddress, timeout)
	c.Conn = conn

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Stop() error {
	if c.Conn != nil {
		c.Conn.Close()
	}

	return nil
}

func (c *Client) SendMessage(msg string) error {
	// client.SendMessage("ini message kita")
	// convert ke payload
	payload := types.Binary(msg)

	// write ke server
	_, err := payload.WriteTo(c.Conn)

	if err != nil {
		fmt.Println("Error sending message", err)
		return err
	}

	// nerima reply dari server
	reply, err := types.Decode(c.Conn)

	if err != nil{
		fmt.Println("Error", err)
		return err
	}

	fmt.Println(string(reply.Bytes()))
	return err
}
