package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"snapscale-api/config"
	"snapscale-api/dashBoard"
	log2 "snapscale-api/libs/log"
)

var Xhub *Hub

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		log2.I.Println(string(message))
		c.hub.Broadcast <- message
	}
}

func (c *Client) writePump() {
	defer func() {
		_ = c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_ = c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{hub: Xhub, conn: c, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func Start() {
	Xhub = newHub()
	Xhub.dashboard = dashBoard.DataCenter
	dashBoard.BroadcastDashboard = Xhub.BroadcastDashboard
	dashBoard.BroadcastBlock = Xhub.BroadcastBlock
	go Xhub.run()
	http.HandleFunc("/ws", ws)
	log.Fatal(http.ListenAndServe(config.WsPort, nil))
}
