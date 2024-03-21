package main

import (
	"log"
	"net/http"

	"github.com/vcscsvcscs/real-time-chat-application/hub"
)

func main() {
	// test is our app running
	log.Println("Hello from Go Chat App Backend")

	// Handles static files from public directory
	http.Handle("/ws/messages", http.HandlerFunc(hub.WsMessager))
	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
