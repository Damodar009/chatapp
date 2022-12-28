package socket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Room *Room
	Conn *websocket.Conn
	Send chan *Message
}

func (c *Client) Read() {
	defer func() {
		print("closing this connection")
		c.Room.Unregister <- c
		c.Conn.Close()
	}()
	for {
		messageType, p, err := c.Conn.ReadMessage()
		print("new message has been read")
		if err != nil {
			log.Panicln(err)
			return
		}

		message := Message{
			Message: string(p),
			Type:    string(messageType),
		}
		println("new message hurray")

		c.Room.Broadcast <- &message

	}

}
