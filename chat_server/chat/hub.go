// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	// "encoding/json"
	"chat-server/db"
	"chat-server/utils"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan db.Message

	// Inbound votes for messages from the clients.
	vote chan db.Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan db.Message, 10),
		vote:       make(chan db.Message, 10),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:

			id, err := utils.Generate_id()
			if err != nil {
				panic(err)
			}
			message.Id = id
			message.Timestamp = utils.Generate_timestamp()

			if err := db.StoreMessage(message); err != nil {
				continue
			}
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.vote:
			new_message, _ := db.UpdateVotes(message)
			h.broadcast <- new_message
		}
	}
}
