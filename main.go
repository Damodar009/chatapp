package main

import (
	"log"
	"net/http"
	"test/socket"
)

func serveWs(chatServer *socket.ChatServer, w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrade(w, r)
	if err != nil {
		log.Fatalln(err)
	}

	client := &socket.Client{
		Server: chatServer,
		Conn:   conn,
	}		

	chatServer.Register <- client
	go client.Read()
}

func main() {
	chatServer := socket.NewChatServer()
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(chatServer, w, r)
	})
	http.ListenAndServe(":8000", nil)
}
	