package main

import (
	"log"
	"net"
	"os"

	"github.com/bettinson/chat_client/api"
)

func main() {
	dial()
}

func parseCommandLine() *api.Client {
	if len(os.Args) != 3 {
		log.Fatal("usage: client <name> <channel>")
	}

	return &api.Client{Name: os.Args[1], Channel: os.Args[2]}
}

func dial() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	client := parseCommandLine()
	bytes_to_send, err := client.Serialize()
	bytes_to_send

	if err != nil {
		log.Fatal(err)
	}
	conn.Write(bytes_to_send)
}
