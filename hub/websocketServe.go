package hub

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]chatClient)

func WsMessager(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	clients[conn] = chatClient{Name: ""}

	greetingMessage := `{"sender":"SERVER","text":"Hello from the server"}`
	err = conn.WriteMessage(websocket.TextMessage, []byte(greetingMessage))
	if err != nil {
		log.Println("Failed to write message to client", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var receivedMessage receivedMessage
	for {
		err = conn.ReadJSON(&receivedMessage)
		if err != nil {
			log.Println("Failed to read from connection", err.Error())
			delete(clients, conn)

			return

		}

		clients[conn] = chatClient{Name: receivedMessage.Sender}
		// Broadcast the message to all connected clients

		for client := range clients {
			if clients[client].Name == receivedMessage.Sender {
				continue
			}

			err = client.WriteJSON(receivedMessage)
			if err != nil {
				log.Println(err.Error())
				client.Close()
				delete(clients, client)

				return
			}
		}
	}
}
