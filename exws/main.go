package exws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSuck(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
func WebSeeHanlder(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi\n")
	http.ServeFile(w, r, "./exws/websockets.html")
}
