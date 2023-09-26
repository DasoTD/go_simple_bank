package main

import (
	"log"
	"net/http"
	// "strings"

	socketio "github.com/googollee/go-socket.io"
)

func mainx() {
	server := socketio.NewServer(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	server.OnConnect("/", func(so socketio.Conn) error {
		log.Printf("Client connected: %s\n", so.ID())
		so.SetContext("")
		return nil
	})

	server.OnEvent("/", "join", func(so socketio.Conn, room string) {
		log.Printf("Client %s joined room %s\n", so.ID(), room)
		so.Join(room)
		so.SetContext(room)
	})

	server.OnEvent("/", "chat", func(so socketio.Conn, msg string) {
		room := so.Context().(string)
		log.Printf("Received chat message '%s' from client %s in room %s\n", msg, so.ID(), room)
		server.BroadcastToRoom(room, "chat", so.ID()+": "+msg)
	})

	server.OnDisconnect("/", func(so socketio.Conn, reason string) {
		room := so.Context().(string)
		log.Printf("Client %s disconnected from room %s: %s\n", so.ID(), room, reason)
		so.Leave(room)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
