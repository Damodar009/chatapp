package socket

import (
	"math/rand"
	"time"
)

type ChatServer struct {
	Client     map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewChatServer() *ChatServer {
	rand.Seed(time.Now().UnixNano())
	chatServer := &ChatServer{
		Client:     make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
	go chatServer.run()
	return chatServer
}

func (r *ChatServer) run() {
	for {
		select {
		case client := <-r.Register:
			r.Client[client] = true
			for client := range r.Client {
				println(client)
				println("client has been registered")
			}
		case client := <-r.Unregister:
			if _, ok := r.Client[client]; ok {
				delete(r.Client, client)	
				close(client.Send)
			}
		case message := <-r.Broadcast:
			for client := range r.Client {
				println("this is okay ")
				if err := client.Conn.WriteJSON(string(message)); err != nil {
				}
			}
		}
	}
}
