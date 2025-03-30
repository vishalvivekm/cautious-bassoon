package models

import (
	"github.com/gorilla/websocket"
)
type Location struct {
	Latitutde float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DevideId string `json:"device_id"`
}

type Hub interface {
	// push will trigger the broadcast
	// and send the passed location to
	// all subscribed clients
	Push(Location) error

	// Register will pass the websocket
	// connection to an internally managed
	// list
	Register(string, *websocket.Conn)

	// De-register will be called when the connection
	// is closed. It will remove the client from the
	// broadcast list.
	DeRegister(string)
}