package config

import (
	"fmt"
	"log"
	"os"
)

const (
	VaultDir string="note-taker";
)

func InitDirPath()  {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get Home directory: ",err)
	}
	noteTakerPath := fmt.Sprintf("%s/%s",&homeDir,VaultDir)

	os.MkdirAll(noteTakerPath,0644)

}