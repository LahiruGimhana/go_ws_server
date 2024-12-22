package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/websock", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
		}

		defer conn.Close()
		defer fmt.Println("message send stopped")

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			conn.WriteMessage(websocket.TextMessage, []byte("Hello World"))
		}

	})
	fmt.Println("server running on port 8050")
	log.Fatal(http.ListenAndServe(":8050", nil))
}