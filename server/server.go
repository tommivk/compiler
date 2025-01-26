package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type jsonMessage struct {
	Command string
	Code    string
}

func handleConnection(c net.Conn) {
	d := json.NewDecoder(c)
	var msg jsonMessage
	err := d.Decode(&msg)
	if err != nil {
		log.Fatal("Failed to decode message: ", err)
	}
	fmt.Println("COMMAND: ", msg.Command)
	fmt.Println("CODE:", msg.Code)
	c.Close()
}

func main() {
	PORT := ":3000"
	ln, err := net.Listen("tcp4", PORT)

	if err != nil {
		log.Fatal("Failed to listen to port: " + PORT)

	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error while accepting connection: ", err)
		}
		go handleConnection(conn)
	}
}
