package src

import (
	"log"
	"net"
	"strings"

	"github.com/RIC217/TutoDiscordChatInGo_Server/utils"
)

// Structure contenant le socket et le pseudo d'un utilisateur
type userSocket struct {
	socket net.Conn
	pseudo string
}

// Renvoie true si l'utilisateur est administrateur, false sinon
func (u userSocket) isOp() bool {
	return strings.Contains(","+strings.Join(utils.GetOpsAuto(), ",")+",", u.pseudo)
}

// Liste des sockets connectés
var sockets []userSocket

// Liste des pseudos connectés
var onlinePseudos []string

// Fonction exécutée à chaque fois qu'un client se connecte
func ProcessClient(conn net.Conn) {
	log.Println("New connection from " + conn.RemoteAddr().String())
	var valid bool
	var pseudo string
	var err error
	for {
		valid, pseudo, err = utils.CheckPseudoAndPassword(conn, &onlinePseudos)
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
		broadcastToEveryone(tempSocket, message)
	}
	removeElementFromSockets(tempSocket)
	log.Printf("%s (with IP %s) has disconnected !\n", pseudo, conn.RemoteAddr().String())
}

// Ajoute un élément à la liste des sockets et à la liste des pseudos connectés
func addElementToSockets(e userSocket) {
	sockets = append(sockets, e)
	onlinePseudos = append(onlinePseudos, e.pseudo)
}

// Supprime un élément à la liste des sockets et à la liste des pseudos connectés
func removeElementFromSockets(e userSocket) {
	var index int = -1
	for key, value := range sockets {
		if value == e {
			index = key
		}
	}
	if index != -1 {
		sockets = append(sockets[:index], sockets[index+1:]...)
	}
	index = -1
	for key, value := range onlinePseudos {
		if value == e.pseudo {
			index = key
		}
	}
	if index != -1 {
		onlinePseudos = append(onlinePseudos[:index], onlinePseudos[index+1:]...)
	}
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
