package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketHub struct {
	clients map[string]*websocket.Conn
	mu sync.Mutex
}

func NewWebSocketHub() *WebsocketHub {
	return &WebsocketHub{
		clients: make(map[string]*websocket.Conn),
	}
}

func (hub *WebsocketHub) Push(payload Location) error {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	for clientId, conn := range hub.clients{
		// WriteJSON will encode the location and send it to the client
		if err := conn.WriteJSON(payload); err != nil {
			// Handle errors gracefully, such as removing disconnected clients
			conn.Close()
			delete(hub.clients, clientId)
			return err
		}
	}
	return nil
}

// Register adds a new client connection to the hub
func(hub *WebsocketHub) Register(clientID string, conn *websocket.Conn){
	hub.mu.Lock()
	defer hub.mu.Unlock()

	hub.clients[clientID] = conn
}

// DeRegister removes a client connection from the hub
func(hub *WebsocketHub) DeRegister(clientID string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	if conn, exists := hub.clients[clientID]; exists{
		// gracefully close the connection
		conn.Close()
		delete(hub.clients, clientID)
	}
}