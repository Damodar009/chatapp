package socket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Server *ChatServer
	Conn   *websocket.Conn
	Send   chan []byte
}

func (c *Client) Read() {
	defer func() {
		c.Server.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Panicln(err)
			return
		}
		c.Server.Broadcast <- message
	}
}
