package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/bettinson/chat_client/api"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	client := &api.Client{}
	// var data []byte
	// c.Read(data)

	message, _ := bufio.NewReader(c).ReadBytes('\n')

	err := client.Deserialize(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %v, Channel: %v\n", client.Name, client.Channel)
	c.Write([]byte("pong"))
}
