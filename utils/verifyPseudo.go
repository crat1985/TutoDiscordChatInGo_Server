package utils

import (
	"log"
	"net"
	"strings"
)

// Vérifie si le pseudo et le mot de passe sont valides et recommence tant qu'ils ne le sont pas.
func CheckPseudoAndPassword(conn net.Conn) (string, error) {
	slice := make([]byte, 1024)
	for {
		n, err := conn.Read(slice)
		if err != nil {
			logErrorToConsoleAndConn(conn, "Erreur lors de la lecture des informations !")
			return "", err
		}
		pseudoAndPassword := string(slice[:n])
		if !strings.Contains(pseudoAndPassword, "\n") {
			logErrorToConsoleAndConn(conn, "Pseudo invalide")
			continue
		}
		pseudo := strings.Split(pseudoAndPassword, "\n")[0]
		password := strings.Split(pseudoAndPassword, "\n")[1]

		if strings.Contains(","+strings.Join(onlinePseudos, ",")+",", ","+pseudo+",") {
			logErrorToConsoleAndConn(conn, "Déjà connecté !")
			continue
		}

		accounts := GetAccountsAuto()
		if accounts[pseudo] == "" {
			logErrorToConsoleAndConn(conn, "Pseudo non existant !")
			continue
		}
		if accounts[pseudo] == password {
			return pseudo, nil
		}
		logErrorToConsoleAndConn(conn, "Mot de passe incorrect !")
	}
}

func logErrorToConsoleAndConn(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
	log.Println(conn.RemoteAddr().String() + " : " + msg)
}
