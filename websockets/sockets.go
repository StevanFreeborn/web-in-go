package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Println("Error upgrading connection")
		}

		for {
			msgType, msg, err := conn.ReadMessage()

			if err != nil {
				fmt.Println("Error reading message")
				return
			}

			fmt.Printf("%s sent %s\n", conn.RemoteAddr(), string(msg))

			writeErr := conn.WriteMessage(msgType, msg)

			if writeErr != nil {
				fmt.Println("Error writing message")
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	fmt.Println("Server listening on port 8080")

	http.ListenAndServe(":8080", nil)
}
