package web

import (
	"fmt"
	//"golang.org/x/net/websocket"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func Echo(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		checkErr(err)
		fmt.Println(message, message[:])
		err = conn.WriteMessage(mt, append([]byte("hello , "), message[:]...))
		checkErr(err)
	}
}
