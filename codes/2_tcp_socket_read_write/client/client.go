package main

import (
	"fmt"
	"log"
	"net"
)

const (
	network = "tcp"
	domain  = "localhost:8887"
	message = "echo message"
)

func main() {

	conn, err := net.Dial(network, domain)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf[:n]))
}
