package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	r := gin.Default()

	r.GET("/ws", wsHandler)

	r.Run(":8887")
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	conn.WriteMessage(1, []byte("hihi"))
	conn.WriteJSON(map[string]string{"a": "b"})
}
