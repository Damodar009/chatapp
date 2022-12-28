package main

import (
	"net/http"
	"test/socket"
)

func serveWs(room *socket.Room, w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrade(w, r)
	if err != nil {
		println("the errror is", err)
	}

	client := &socket.Client{
		Room: room,
		Conn: conn,
	}
	println("room is registed serveWs0")
	room.Register <- client
	println("room is registed serveWs1")
	client.Read()
}
func setUpRoute() {
	room := socket.NewRoom()
	println(room)

	http.HandleFunc("/rotue/get_user", func(w http.ResponseWriter, r *http.Request) {
		serveWs(room, w, r)
	})

}

func main() {
	setUpRoute()
	//database.SetUpDBConnection()

	// query := `Insert into user_by_id(userid,username, photourl , status , email)
	// values (uuid() , 'damodar' , 'photourlchaina' , true , 'damodar@gmail.com');`

	// database.ExecuteQuery(query)

	http.ListenAndServe(":8000", nil)

}
