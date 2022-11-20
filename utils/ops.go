package utils

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

var (
	//Chemin du fichier contenant la liste des administrateurs
	opsFilePath string = path.Join(".", "config", "ops.json")
	//Liste des administrateurs
	ops []string
)

// Retourne les administrateurs, mais en ne lisant le fichier de configuration que si nécessaire
func GetOpsAuto() []string {
	if len(ops) == 0 {
		return GetOps()
	}
	return ops
}

// Encode la liste des administrateurs en JSON
func Encode(ops... string) []byte {
	content, err := json.MarshalIndent(ops, "", "  ")
	if err != nil {
		panic(err)
	}
	return content
}

// Créer le fichier contenant la liste des administrateurs
func CreateOpsFile(ops ...string) {
	f, err := os.Create(opsFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(Encode(ops...))
}

// Lire le fichier contenant la liste des administrateurs
func ReadOpsFile() []byte {
	_, err := os.Stat(opsFilePath)
	if err != nil {
		log.Println("Creating ops file...")
		CreateOpsFile("admin")
		log.Println("Created ops file !")
	}
	content, err := os.ReadFile(opsFilePath)
	if err != nil {
		panic(err)
	}
	return content
}

// Décoder le fichier de configuration JSON vers une liste de strings
func Decode() []string {
	err := json.Unmarshal(ReadOpsFile(), &ops)
	if err != nil {
		panic(err)
	}
	return ops
}

// Retourne les administrateurs en lisant le fichier de configuration
func GetOps() []string {
	return Decode()
}
