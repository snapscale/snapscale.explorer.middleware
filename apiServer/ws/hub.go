package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"snapscale-api/dashBoard"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

type Hub struct {
	clients    map[*Client]bool
	Broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	dashboard  *dashBoard.DataCenterS
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
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

func (h *Hub) BroadcastDashboard() {
	data := &Msg{}
	data.Action = "dashboard"
	data.Data = h.dashboard
	p, _ := json.Marshal(data)
	h.Broadcast <- p
}

func (h *Hub) BroadcastBlock(in interface{}) {
	data := &Msg{}
	data.Action = "block"
	data.Data = in
	p, _ := json.Marshal(data)
	h.Broadcast <- p
}

func newHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}
