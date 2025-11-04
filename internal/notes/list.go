package notes

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/config"
)

type NoteItem struct{
	Title string
	Desc string
}

func (i NoteItem) TitleString() string       { return i.Title }
func (i NoteItem) DescriptionString() string  { return i.Desc }
func (i NoteItem) FilterValue() string { return i.Title }


func ListFiles() []list.Item {
	items:= make([]list.Item,0)
    entries,err :=os.ReadDir(config.VaultDir)
	if err != nil {
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
				Title: entry.Name(),
				Desc:  "Modified: " + modifiedTime,
			})			
		}
	}
	return items
}
