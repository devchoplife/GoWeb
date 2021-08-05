package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

//We will build a simple server that echoes back everything we send to it
//we will use the gorilla/websocket package

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) //error ignored for the sake of simplicity

		for {
			//Read messages from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			//Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			//Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
