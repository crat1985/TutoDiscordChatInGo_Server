package main

import (
	"log"
	"net"

	"github.com/RIC217/TutoDiscordChatInGo_Server/src"
	"github.com/RIC217/TutoDiscordChatInGo_Server/utils"
)

const port = "8888"

var ops []string

func main() {
	log.Printf("Starting server on port %s...\n", port)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	log.Printf("Server listening on port %s...\n", port)
	utils.GetAccounts()
	ops = utils.GetOpsAuto()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go src.ProcessClient(conn, &ops)
	}
}
