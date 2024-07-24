package setupRouter

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Fprintln(w, "Unable to setup WebSocket connection")
		return ws, err
	}
	return ws, nil
}

// // define our WebSocket endpoint
func serverWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	conn, err := Upgrade(w, r)
	if err != nil {
		log.Println(w, "Unable to setup WebSocket connection")
		return
	}

	client := &Client{
		ID:   fmt.Sprintf("%v", time.Now().UnixMilli()),
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

}

func SetupRouter() {
	pool := NewPool()
	go pool.Start()

	// handle our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})
}
