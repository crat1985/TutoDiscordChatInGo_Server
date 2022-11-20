package utils

import (
	"strings"
)

var commands map[string]func(u userSocket) = make(map[string]func(u userSocket))

func SetCommands() {
	commands["reload"] = func(u userSocket) {
		if u.isOp() {
			GetAccounts()
			GetOps()
			SetCommands()
			writeToClient("serv", "Configuration rechargée !", u.socket)
		} else {
			writeToClient("serv", "Vous n'êtes pas un administrateur !", u.socket)
		}
	}
}

func isCommand(msg string, usersocket userSocket) bool {
	msg = strings.ToLower(msg)
	if !strings.HasPrefix(msg, "/") {
		return false
	}
	if commands[msg[1:]] == nil {
		writeToClient("serv", "Commande introuvable !\n", usersocket.socket)
		return false
	}
	commands[msg[1:]](usersocket)
	return true
}
