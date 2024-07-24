package setupRouter

import (
	"fmt"
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("Pool Size: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				log.Println("New client registered:", client.ID)
				client.Conn.WriteJSON(Message{Type: 0, Body: "New User Joined..."})
			}
			break

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("Pool Size: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				log.Println("Client disconnected", client.ID)
				client.Conn.WriteJSON(Message{Type: 0, Body: fmt.Sprintf("User %s disconnected ", client.ID)})
			}
			break

		case message := <-pool.Broadcast:
			log.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Fatalln("Error: ", err)
					return
				}
			}
		}
	}
}
