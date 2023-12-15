package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net"
)

var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

func Upgrade(conn net.Conn) (net.Conn, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read failed")
		return nil, err
	}

	lines := string(buf[:n])
	httpy := NewHTTPy().GetHTTPy(lines)
	key := httpy.SecWebKey[1:]
	fmt.Println(string(buf[:n]))

	responseHTTPString := GetUpgradeResponseString(key)
	_, err = conn.Write(responseHTTPString)
	if err != nil {
		log.Fatal(err)
	}

	return conn, nil
}

func GetUpgradeResponseString(webSecSocketkey string) []byte {
	h := sha1.New()
	h.Write([]byte(webSecSocketkey))
	h.Write(keyGUID)
	secWebSocketAccept := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return []byte(fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", secWebSocketAccept))
}
