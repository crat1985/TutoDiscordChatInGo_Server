package utils

import (
	"strings"
)

var Commands []Command

type Command struct {
	Name        string
	Description string
	Action      func(u userSocket)
	NeedOp      bool
}

func GetOpsCommands() string {
	var textToSend string
	var x uint8
	for _, command := range Commands {
		if x == 0 {
			textToSend += command.Name + " - " + command.Description
			x++
			continue
		}
		textToSend += "\n" + command.Name + " - " + command.Description
	}
	return textToSend
}

func GetNonOpsCommands() string {
	var textToSend string
	var x uint8
	for _, command := range Commands {
		if command.NeedOp {
			continue
		}
		if x == 0 {
			textToSend += command.Name + " - " + command.Description
			x++
			continue
		}
		textToSend += "\n" + command.Name + " - " + command.Description
	}
	return textToSend
}

func SetCommands() {
	Commands = nil
	Commands = append(Commands, Command{Name: "reload", Description: "Cette commande permet de recharger la configuration afin de mettre à jour d'éventuels changements.", Action: func(u userSocket) {
		if !u.isOp() {
			writeToClientAsServer("Vous n'avez pas la permission de faire ça !", u.socket)
			return
		}
		GetOps()
		GetAccounts()
		writeToClientAsServer("Configuration rechargée !", u.socket)
	}, NeedOp: true})
	Commands = append(Commands, Command{Name: "help", Description: "Affiche cet aide.", Action: func(u userSocket) {
		if u.isOp() {
			writeToClientAsServer(GetOpsCommands(), u.socket)
			return
		}
		writeToClientAsServer(GetNonOpsCommands(), u.socket)
	}, NeedOp: false})
}

// Verifie si le texte envoyé est une commande et l'exécute si c'est le cas.
func executeCommand(cmd string, u userSocket) (isCommand bool) {
	for _, command := range Commands {
		if command.Name == cmd {
			command.Action(u)
			return true
		}
	}
	return false
}

// Vérifie si un utilisateur a entré une commande et si c'est le cas, exécute la commande correspondante.
func isCommand(msg string, usersocket userSocket) bool {
	msg = strings.ToLower(msg)
	if !strings.HasPrefix(msg, "/") {
		return false
	}
	return executeCommand(msg[1:], usersocket)
}
