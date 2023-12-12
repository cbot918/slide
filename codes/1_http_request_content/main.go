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
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read failed")
			continue
		}

		fmt.Println(string(buf[:n]))

	}
}
