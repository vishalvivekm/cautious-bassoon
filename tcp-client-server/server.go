package main

import (
	"fmt"
	"log"
	"net"
)

const (
	protocol = "tcp"
	address  = "127.0.0.1"
	port     = ":8080"
)

func main() {
	listener, err := net.Listen("tcp", address+port) // initiate TCP server socket on adr 127.0.0.1:8080
	if err != nil {
		log.Fatal("cannot open server socket, ", err.Error())
	}
	defer listener.Close()
	log.Printf("Server socker is listening on %s%s\n", address, port)

	for {
		connection, err := listener.Accept() // server continuously listens for incoming connections
		if err != nil {
			log.Println("cannot accept client connection", err)
			continue
		}

		go func(connection net.Conn) {
			err = handleConnection(connection)
			if err != nil {
				log.Println(err)
			}
		}(connection)
	}
}

func handleConnection(connection net.Conn) error {
	buffer := make([]byte, 1024)

	clientMsgLen, err := connection.Read(buffer)
	if err != nil {
		return err
	}
	log.Printf("Receieved from client: %s\n", string(buffer[:clientMsgLen]))

	err = sendResponse(connection, "Server socket says hello!")
	if err != nil {
		return err
	}
	return connection.Close()
}

func sendResponse(connection net.Conn, message string) error {
	_, err := connection.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("cannot write server data to connection: %w", err)
	}
	return nil
}
