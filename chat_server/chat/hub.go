// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	// "encoding/json"
	"chat-server/db"
	"context"
	"fmt"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan db.Message

	// Publish Inbound messages to other chat servers
	publish chan db.Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan db.Message, 10),
		publish:    make(chan db.Message, 10),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) GetSubChannel() chan db.Message {

	return h.broadcast
}

func (h *Hub) GetClientCount() int {

	return len(h.clients)

}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			cur, err := db.GetMessageHistory()
			if err != nil {
				continue
			}
			for cur.Next(context.Background()) {
				// To decode into a struct, use cursor.Decode()
				var message db.Message
				err := cur.Decode(&message)
				fmt.Println("history data")
				fmt.Println(message)
				if err != nil {
					log.Fatal(err)
				}
				// do something with result...
				client.send <- message
			}
			if err := cur.Err(); err != nil {
				panic(err)
			}
			cur.Close(context.Background())

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.publish:
			db.PublishMessage(message)
		}
	}
}
