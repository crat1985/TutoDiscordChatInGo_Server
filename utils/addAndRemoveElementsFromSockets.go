package utils

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
