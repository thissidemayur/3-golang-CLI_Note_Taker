package notes

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/config"
)

// Create file
func CreateFile(fileName string) (*os.File,error) {
	filePath:= filepath.Join(config.VaultDir,fileName+".md")
	if _,err:=os.Stat(filePath); err == nil{
		return nil, fmt.Errorf("file already exists")
	}
	file,err:=os.Create(filePath)
	if err!= nil{
		log.Printf("Could not create file: %v ; @/internal/notes/manager.go \n",err)
		return nil,err
	}
	return file,nil
}

// Save file
func SaveFile(file *os.File, content string) error {
	   if file == nil {
        return fmt.Errorf("file is not open")
    }
	// delete existing content. size get truncated to 0
	if err:=file.Truncate(0); err != nil {
		log.Printf("Could not truncate file: %v ; @/internal/notes/manager.go \n",err)
		return err
	}
	// move the cursor to the beginning of the file
	if _,err := file.Seek(0,0); err != nil {
		log.Printf("Could not seel file: %v ; @/internal/notes/manager.go \n",err)
		return err
	}

	// write a new content
	if _,err := file.WriteString(content); err != nil {
		log.Printf("Could not Write file: %v ; @/internal/notes/manager.go \n",err)
		return err
	}
	if err:= file.Sync(); err != nil {
		log.Printf("Could not Sync file: %v ; @/internal/notes/manager.go \n",err)
		return err
	}
	return file.Close()
}

// Delete file
func DeleteFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		log.Printf("Could not delete file: %v ; @/internal/notes/manager.go \n",err)
		return err
	}
	return nil
}

// Read file
func ReadFile(filePath string) (string, error) {
	data,err:=os.ReadFile(filePath)
	if err != nil {
		log.Printf("Could not read file: %v ; @/internal/notes/manager.go \n",err)
		return "" ,err
	}
	return string(data),nil
}