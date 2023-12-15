package main

import (
	"fmt"
	"net"
)

type Client struct {
	Conn  net.Conn
	frame *Frame
}

func NewClient(conn net.Conn) (*Client, error) {

	conn, err := Upgrade(conn)
	if err != nil {
		return nil, err
	}
	fmt.Printf("websocket upgrade success!\n\n")

	return &Client{
		Conn:  conn,
		frame: NewFrame(),
	}, nil
}

func (c *Client) Read() ([]byte, error) {
	var err error
	buf2 := make([]byte, 4096)
	_, err = c.Conn.Read(buf2)
	if err != nil {
		fmt.Println("read buffer failed")
		return nil, err
	}

	return c.frame.DecodeFrame(buf2), nil
}
