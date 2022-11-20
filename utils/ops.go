package utils

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

var (
	opsFilePath string = path.Join(".", "config", "ops.json")
	ops         []string
)

func GetOpsAuto() []string {
	if len(ops) == 0 {
		return GetOps()
	}
	return ops
}

func Encode(ops []string) []byte {
	content, err := json.MarshalIndent(ops, "", "  ")
	if err != nil {
		panic(err)
	}
	return content
}

func CreateOpsFile(ops ...string) {
	f, err := os.Create(opsFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(Encode(ops))
}

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

func Decode() []string {
	err := json.Unmarshal(ReadOpsFile(), &ops)
	if err != nil {
		panic(err)
	}
	return ops
}

func GetOps() []string {
	return Decode()
}
