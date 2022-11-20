package utils

import "unicode"

// Renvoie true si la chaîne de caractère renseignée ne contient que des lettres, des chiffres ou des underscores, faux sinon
func IsAlphaNum(text string) bool {
	for _, value := range text {
		if !unicode.IsLetter(value) && !unicode.IsDigit(value) && string(value) != "_" {
			return false
		}
	}
	return true
}
