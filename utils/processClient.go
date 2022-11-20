package utils

import (
	"log"
	"net"
	"strings"
)

// Structure contenant le socket et le pseudo d'un utilisateur
type userSocket struct {
	socket net.Conn
	pseudo string
}

// Renvoie true si l'utilisateur est administrateur, false sinon
func (u userSocket) isOp() bool {
	return strings.Contains(","+strings.Join(GetOpsAuto(), ",")+",", u.pseudo)
}

// Liste des sockets connectés
var sockets []userSocket

// Liste des pseudos connectés
var onlinePseudos []string

// Fonction exécutée à chaque fois qu'un client se connecte
func ProcessClient(conn net.Conn) {
	log.Println("New connection from " + conn.RemoteAddr().String())
	pseudo, err := CheckPseudoAndPassword(conn)
	if err != nil {
		return
	}
	conn.Write([]byte("pseudook"))
	log.Printf("Pseudo for %s is now %s !\n", conn.RemoteAddr().String(), pseudo)
	tempSocket := userSocket{pseudo: pseudo, socket: conn}
	addElementToSockets(tempSocket)
	listenForMessages(conn, tempSocket)
	removeElementFromSockets(tempSocket)
	log.Printf("%s (with IP %s) has disconnected !\n", pseudo, conn.RemoteAddr().String())
}

// Envoyer un message à tous les sockets connectés et l'affiche dans la console
func broadcastToEveryone(sender userSocket, message string) {
	if sender.isOp() {
		sender.pseudo = "[Admin] " + sender.pseudo
	}
	for _, user := range sockets {
		user.socket.Write([]byte(sender.pseudo + "\n" + message))
	}
	log.Println(sender.pseudo + ": " + message)
}
