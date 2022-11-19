package src

import (
	"log"
	"net"
	"strings"

	"github.com/RIC217/TutoDiscordChatInGo_Server/utils"
)

type userSocket struct {
	socket net.Conn
	pseudo string
}

var sockets []userSocket

func ProcessClient(conn net.Conn) {
	log.Println("New connection from " + conn.RemoteAddr().String())
	var valid bool
	var pseudo string
	var err error
	for {
		valid, pseudo, err = utils.CheckPseudo(conn)
		if err != nil {
			conn.Close()
			log.Printf("%s disconnected without logging in !\n", conn.RemoteAddr().String())
			return
		}
		if !valid {
			log.Println("Invalid pseudo : " + pseudo)
		} else {
			break
		}
	}
	conn.Write([]byte("pseudook"))
	log.Printf("Pseudo for %s is now %s !\n", conn.RemoteAddr().String(), pseudo)
	tempSocket := userSocket{pseudo: pseudo, socket: conn}
	addElementToSockets(tempSocket)
	slice := make([]byte, 1024)
	var message string
	for {
		n, err := conn.Read(slice)
		if err != nil {
			log.Println(err)
			break
		}
		message = string(slice[:n])
		message = strings.ReplaceAll(message, "\n", "")
		message = strings.ReplaceAll(message, "\t", "")
		message = strings.Trim(message, " ")
		broadcastToEveryone(tempSocket.pseudo, message)
	}
	removeElementFromSockets(tempSocket)
	log.Printf("%s (with IP %s) has disconnected !\n", pseudo, conn.RemoteAddr().String())
}

func addElementToSockets(e userSocket) {
	sockets = append(sockets, e)
}

func removeElementFromSockets(e userSocket) {
	var index int = -1
	for key, value := range sockets {
		if value == e {
			index = key
		}
	}
	if index == -1 {
		return
	}
	sockets = append(sockets[:index], sockets[index+1:]...)
}

func broadcastToEveryone(sender, message string) {
	for _, user := range sockets {
		user.socket.Write([]byte(sender + "\n" + message))
	}
	log.Println(sender + ": " + message)
}
