package main

import (
	"fmt"
	"log"
	"net"
)

const (
	network = "tcp"
	port    = ":8887"
)

func main() {

	fmt.Println(port)
	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var err error
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}

		c, err := NewClient(conn)
		if err != nil {
			log.Fatal(err)
		}
		msg, err := c.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg))

	}
}
