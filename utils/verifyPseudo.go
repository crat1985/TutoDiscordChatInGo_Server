package utils

import (
	"net"
	"unicode"
)

func CheckPseudo(conn net.Conn) (bool, string, error) {
	slice := make([]byte, 1024)
	n, err := conn.Read(slice)
	if err != nil {
		conn.Write([]byte("Impossible de lire le pseudo !"))
		return false, "Impossible de lire le pseudo !", err
	}
	pseudo := string(slice[:n])
	if len(pseudo) < 5 {
		conn.Write([]byte("Pseudo trop court (5 caractères minimum) !"))
		return false, "Pseudo trop court (5 caractères minimum) !", nil
	}
	if !unicode.IsLetter(rune([]byte(pseudo)[0])) {
		conn.Write([]byte("Le pseudo doit commencer par une lettre !"))
		return false, "Le pseudo doit commencer par une lettre !", nil
	}
	if !IsAlphaNum(pseudo) {
		return false, "Le pseudo ne peut contenir que des lettres, des underscores et des chiffres !", nil
	}
	return true, pseudo, nil
}
