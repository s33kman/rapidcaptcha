// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Registered clients.
	captchas map[string]Captcha

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		captchas:   make(map[string]Captcha),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				fmt.Println("Captchas in Pipe: ", h.captchas[client.RequestorID])
				delete(h.clients, client)
				delete(h.captchas, client.RequestorID)
				fmt.Println("Client left: ", client.RequestorID)
				fmt.Println("Captchas in Pipe: ", h.captchas[client.RequestorID])
				close(client.send)
			}
		case message := <-h.broadcast:
			c := Captcha{}
			err := json.Unmarshal(message, &c)
			if err != nil {
				fmt.Println(err)
			}
			h.captchas[c.GroupID] = c
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
