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
			conn.Write([]byte("Erreur lors de la lecture des informations !"))
			log.Println(conn.RemoteAddr().String() + " : Erreur lors de la lecture des informations !")
			return "", err
		}
		pseudoAndPassword := string(slice[:n])
		if !strings.Contains(pseudoAndPassword, "\n") {
			log.Println(conn.RemoteAddr().String() + ": Pseudo invalide")
			continue
		}
		pseudo := strings.Split(pseudoAndPassword, "\n")[0]
		password := strings.Split(pseudoAndPassword, "\n")[1]

		if strings.Contains(","+strings.Join(onlinePseudos, ",")+",", ","+pseudo+",") {
			conn.Write([]byte("Déjà connecté !"))
			log.Println(conn.RemoteAddr().String() + ": Déjà connecté !")
			continue
		}

		accounts := GetAccountsAuto()
		if accounts[pseudo] == "" {
			conn.Write([]byte("Pseudo non existant !"))
			log.Println(conn.RemoteAddr().String() + ": Pseudo non existant !")
			continue
		}
		if accounts[pseudo] == password {
			return pseudo, nil
		}
		conn.Write([]byte("Mot de passe incorrect !"))
	}
}
