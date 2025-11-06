package config

import (
	"log"
	"os"
	"path/filepath"
)

var (
	VaultDir string
	dirName  = ".note-taker"
)


func InitDirPath()  {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("❌ Could not get Home directory: ",err)
	}
	VaultDir = filepath.Join(homeDir, dirName)
	if err:=os.MkdirAll(VaultDir,0755); err != nil {
		log.Fatal("❌ Could not create Vault directory: ",err)
	}

	log.Println("✅ Vault directory:", VaultDir)

}