package utils

import (
	"net"
	"strings"
)

// Ecoute les messages envoyés par le client
func listenForMessages(conn net.Conn, tempSocket userSocket) {
	slice := make([]byte, 1024)
	var message string
	for {
		n, err := conn.Read(slice)
		if err != nil {
			break
		}
		message = string(slice[:n])
		message = strings.ReplaceAll(message, "\n", "")
		message = strings.ReplaceAll(message, "\t", "")
		message = strings.Trim(message, " ")
		broadcastToEveryone(tempSocket, message)
	}
}
