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

		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nDate: Fri, 15 Dec 2023 03:58:55 GMT\r\nContent-Length: 27\r\nContent-Type: text/plain; charset=utf-8\r\nHello, you've requested: /\r\n\r\n")))

		// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		// })

		// // Start the HTTP server on port 8080
		// fmt.Println(port)
		// http.ListenAndServe(port, nil)

	}
}
