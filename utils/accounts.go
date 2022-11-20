package utils

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

var (
	Accounts map[string]string = make(map[string]string)
)

func ResetAccounts() {
	for k := range Accounts {
		delete(Accounts, k)
	}
}

func createConfigFile() error {
	//TODO You can delete this and use the JSON file instead
	Accounts["admin"] = "password"
	bytes, err := json.MarshalIndent(Accounts, "", "  ")
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(bytes)
	return nil
}

var filePath = path.Join(".", "config", "config.json")

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

func GetAccounts() map[string]string {
	createConfigDir()
	content, err := ReadConfig()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, &Accounts)
	return Accounts
}

func GetAccountsAuto() map[string]string {
	if len(Accounts) == 0 {
		return GetAccounts()
	}
	return Accounts
}

func ResetAndGetAccounts() map[string]string {
	ResetAccounts()
	return GetAccounts()
}

func createConfigDir() {
	_, err := os.Stat("config")
	if err != nil {
		os.Mkdir("config", 0775)
	}
}
