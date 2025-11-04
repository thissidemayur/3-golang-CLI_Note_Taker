package notes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/config"
)

// Create file
func CreateFile(fileName string) (*os.File,error) {
	filePath:= filepath.Join(config.VaultDir,fileName+".md")

	if _,err:=os.Stat(filePath); err != nil{
		return nil, fmt.Errorf("file already exists")
	}
	file,err:=os.Create(filePath)
	if err!= nil{
		return nil,err
	}
	return file,nil
}

// Save file
func SaveFile(file *os.File, content string) error {
	if err:=file.Truncate(0); err != nil {
		return err
	}

	if _,err := file.Seek(0,0); err != nil {
		return err
	}

	return file.Close()
}

// Delete file
func DeleteFile(file *os.File) error {
	if err := os.Remove(file.Name()); err != nil {
		return err
	}
	return nil

}