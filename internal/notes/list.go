package notes

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/config"
)

type NoteItem struct{
	title string
	desc string
}

func (i NoteItem) Title() string       { return i.title }
func (i NoteItem) Description() string  { return i.desc }
func (i NoteItem) FilterValue() string { return i.title }


func ListFiles() []list.Item {
	items := make([]list.Item, 0)
	entries, err := os.ReadDir(config.VaultDir)
	if err != nil {
		log.Fatalln("Error while reading notes: ",err)
		return nil
	}
	for _,entry := range entries {
		if !entry.IsDir() {
			fileInfo,err := entry.Info()
			if err != nil {
				// TODO: handle it later so that every file will be listed
				continue
			}

			modifiedTime := fileInfo.ModTime().Format("2006-01-02 15:04")
			items =append(items,NoteItem{
				title: entry.Name(),
				desc:  "Modified: " + modifiedTime,
			})
		}
	}
	log.Println("Listing all notes : ", items)
	return items
}
