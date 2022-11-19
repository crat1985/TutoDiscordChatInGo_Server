package utils

import (
	"net"
	"strings"
)

func CheckPseudoAndPassword(conn net.Conn) (valid bool, infos string, err error) {
	slice := make([]byte, 1024)
	n, err := conn.Read(slice)
	if err != nil {
		conn.Write([]byte("Erreur lors de la lecture des informations !"))
		return false, "Erreur lors de la lecture des informations !", err
	}
	pseudoAndPassword := string(slice[:n])
	if !strings.Contains(pseudoAndPassword, "\n") {
		return false, "Pseudo invalide", nil
	}
	pseudo := strings.Split(pseudoAndPassword, "\n")[0]
	password := strings.Split(pseudoAndPassword, "\n")[1]

	accounts := GetAccounts()
	if accounts[pseudo] == "" {
		conn.Write([]byte("Pseudo invalide !"))
		return false, "Pseudo invalide !", nil
	}
	if accounts[pseudo] == password {
		return true, pseudo, nil
	}
	conn.Write([]byte("Mot de passe incorrect !"))
	return false, "Mot de passe incorrect !", nil
}
