package utils

import "unicode"

func IsAlphaNum(text string) bool {
	for _, value := range text {
		if !unicode.IsLetter(value) && !unicode.IsDigit(value) && string(value) != "_" {
			return false
		}
	}
	return true
}
