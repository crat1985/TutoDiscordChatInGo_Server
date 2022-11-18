package main

import (
	"log"
	"net"

	"github.com/RIC217/TutoDiscordChatInGo_Server/src"
)

func main() {
	log.Println("Starting server on port 8080...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("Server listening on port 8080...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go src.ProcessClient(conn)
	}
}
