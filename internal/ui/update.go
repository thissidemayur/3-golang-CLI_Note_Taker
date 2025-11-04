package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/notes"
)

func (n NoteModel) Update(msg tea.Msg)(tea.Model, tea.Cmd) {
	var cmd tea.Cmd;

	switch msg := msg.(type) {
		case tea.KeyMsg :
			switch msg.String() {
				// quit terminal
				case "ctrl+c":
					return n  , tea.Quit

				// create a file
				case "ctrl+n":
					n.IsTitleInput = true
					return n, nil

				// save file
				case "ctrl+s":
					if n.currentFile != nil {
						if err:=notes.SaveFile(n.currentFile,n.ContentInput.Value()); err != nil{
							log.Fatalf("Could not save file: %v \n",err)
						}
						n.currentFile = nil
					}
					return n,nil
				
				// list files
				case "ctrl+l":
					n.ShowingList = true
					return n, nil

				// CREATE FILE	
				case "enter":
					if n.IsTitleInput {
						fileName := n.TitleInput.View()
						file,err:=notes.CreateFile(fileName)
						if err == nil {
							n.currentFile=file
						}
						n.IsTitleInput = false
						
					}
					return n,nil

				// back (without saving)
				case "esc":
					// Back from writing the file:
					if n.IsTitleInput {
						n.IsTitleInput = false
					}
					// file isnt created
					if n.currentFile != nil {
						n.currentFile=nil
					}
					// back from list view and filtering
					if n.ShowingList {
						n.ShowingList = false
						// back from filtering
						if n.ListFiles.FilterState() == list.Filtering {
							break
						}
					}
					return n,nil
			}

		}
		// now we can write inside titleInput using textinput
		if n.IsTitleInput {
			n.TitleInput , cmd= n.TitleInput.Update(msg)
		}

		// now we can write inside file  using textarea
		if n.currentFile!= nil {
			n.ContentInput, cmd= n.ContentInput.Update(msg)
		}

		// we can intereacct with list using list
		if n.ShowingList {
			n.ListFiles , cmd = n.ListFiles.Update(msg)
		}
		return n,cmd
}