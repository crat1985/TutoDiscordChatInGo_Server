package utils

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

// Liste des comptes
var Accounts map[string]string = make(map[string]string)

// Rétablir la configuration d'usine des comptes
func ResetAccounts() {
	createConfigFile()
}

// Créer ou recréer le fichier de configuration avec les comptes par défaut (admin:password et example:example)
func createConfigFile() error {
	Accounts["admin"] = "password"
	Accounts["example"] = "example"
	bytes, err := json.MarshalIndent(Accounts, "", "  ")
	if err != nil {
		return err
	}
	return writeToConfigFile(bytes)
}

// Ecrire des données à l'intérieur du fichier de configuration
func writeToConfigFile(bytes []byte) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(bytes)
	return nil
}

// Chemin vers le fichier de configuration
var filePath = path.Join(".", "config", "config.json")

// Lire et retourner le contenu du fichier de configuration
func ReadConfig() ([]byte, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		log.Println("Creating config file...")
		err = createConfigFile()
		if err != nil {
			panic(err)
		}
		log.Println("Created config file !")
	}
	log.Println("Reading config file...")
	content, err := os.ReadFile(filePath)
	log.Println("Read config file !")
	return content, err
}

// Retourne les comptes en lisant le fichier de configuration
func GetAccounts() map[string]string {
	createConfigDir()
	content, err := ReadConfig()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, &Accounts)
	return Accounts
}

// Retourne les comptes, mais en ne lisant le fichier de configuration que si nécessaire
func GetAccountsAuto() map[string]string {
	if len(Accounts) == 0 {
		return GetAccounts()
	}
	return Accounts
}

// Rétablit la configuration d'usine puis retourne les comptes
func ResetAndGetAccounts() map[string]string {
	ResetAccounts()
	return GetAccounts()
}

// Créer le dossier où se trouvent les fichiers de configuration
func createConfigDir() {
	_, err := os.Stat("config")
	if err != nil {
		os.Mkdir("config", 0775)
	}
}
