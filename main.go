package main

import (
	"gochat/chat"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)
	// websocket server
	server := chat.NewServer("/entry")
	go server.Listen()
	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
