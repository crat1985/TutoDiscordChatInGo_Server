package utils

import (
	"log"
	"strings"
)

type Command struct {
	Action      func(u userSocket)
	Description string
	NeedOp      bool
}

// Liste des commandes
var commands map[string]Command = make(map[string]Command)

// Met à jour la liste des commandes
func SetCommands() {
	commands["reload"] = Command{Action: func(u userSocket) {
		if u.isOp() {
			GetAccounts()
			GetOps()
			SetCommands()
			writeToClientAsServer("Configuration rechargée !", u.socket)
		} else {
			writeToClientAsServer("Vous n'êtes pas un administrateur !", u.socket)
		}
	}, Description: "Recharger la configuration.", NeedOp: true}
	commands["help"] = Command{Action: func(u userSocket) {
		if u.isOp() {
			writeToClientAsServer(GetAllCommands(), u.socket)
		} else {
			writeToClientAsServer(GetCommandsNotOp(), u.socket)
		}
	}, Description: "Afficher cet aide.", NeedOp: false}
}

func GetCommandsNotOp() string {
	var commandDescList string
	var x uint8
	for k, command := range commands {
		if command.NeedOp {
			continue
		}
		if x == 0 {
			commandDescList = k + " - " + command.Description
			x++
			continue
		}
		commandDescList += "\n" + k + " - " + command.Description
	}
	log.Println(commandDescList)
	return commandDescList
}

func GetAllCommands() string {
	var commandDescList string
	var x uint8
	for k, command := range commands {
		if x == 0 {
			commandDescList = k + " - " + command.Description
			x++
			continue
		}
		commandDescList += "\n"+k + " - " + command.Description
	}
	log.Println(commandDescList)
	return commandDescList
}

// Vérifie si un utilisateur a entré une commande et si c'est le cas, exécute la commande correspondante
func isCommand(msg string, usersocket userSocket) bool {
	msg = strings.ToLower(msg)
	if !strings.HasPrefix(msg, "/") {
		return false
	}
	if commands[msg[1:]].Action == nil {
		writeToClientAsServer("Commande introuvable !\n", usersocket.socket)
		return false
	}
	commands[msg[1:]].Action(usersocket)
	return true
}
