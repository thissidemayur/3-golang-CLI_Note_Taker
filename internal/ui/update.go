package ui

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/config"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/notes"
)

const messageTimeout = 3 * time.Second
func (n NoteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

		// message expiry
	case clearMessageMsg:
		n.MsgText = ""
		n.MsgType = MessageNone
		return n, nil

	// MAKE TERMINAL RESIZEABLE
	case tea.WindowSizeMsg:
		h, v := n.Theme.BorderStyle.GetFrameSize()
		   // Save window dimensions
    n.Width = msg.Width
    n.Height = msg.Height

    // Dynamically set list size
    n.ListFiles.SetSize(msg.Width-h-4, msg.Height-v-10)
    return n, nil
	
		// KEYBOARD EVENT TRIGGERED
	case tea.KeyMsg:
		switch msg.String() {
		// quit terminal
		case "ctrl+q":
			return n, tea.Quit

		// create a file
		case "ctrl+n":
			log.Println("Enter pressed: ", n.IsTitleInput)
			n.IsTitleInput = true
			

				// save filecl
				case "ctrl+s":
					if n.currentFile != nil {

						if err:=notes.SaveFile(n.currentFile,n.ContentInput.Value()); err != nil{
							n.MsgText = fmt.Sprintf("Save failed: %v \n",err)
							n.MsgType = MessageError
							return n, clearMsgAfter()
						}
						n.MsgText = "💾 File saved successfully!"
						n.MsgType = MessageSuccess
						n.currentFile = nil
						return n, clearMsgAfter()
					}
					return n, nil

				// list files
				case "ctrl+l":
					n.ShowingList = true
					// list instance
					fileLists := notes.ListFiles()
					lists := list.New(fileLists, list.NewDefaultDelegate(), n.Width, n.Height)
					lists.Title =
					 "🗒️ ALL NOTES LIST: "
					n.ListFiles = lists
					// n.MsgText = "📂 Listing all notes..."
					// n.MsgType = MessageInfo
					return n, clearMsgAfter()

				// CREATE FILE, OPEN FILE from list,
				case "enter":
					// writing a file (teaxtarea mode)
					if n.currentFile != nil {
						break;
					}
					// create a new file
					fileName := n.TitleInput.Value()
					if fileName !="" {
						file,err:=notes.CreateFile(fileName)
						if err != nil {
							n.MsgText = fmt.Sprintf(" %v: \n",err)
							n.MsgType = MessageError
							return n , clearMsgAfter()
						}
						n.IsTitleInput = false
						n.currentFile = file

						n.TitleInput.Blur()
						n.ContentInput.Focus()
						n.TitleInput.SetValue("")
						n.MsgText = fmt.Sprintf("✅ Created file: %s.md", fileName)
						n.MsgType = MessageSuccess
						return n, clearMsgAfter()
					}

					// TODO: open file from list by pressing enter
					if n.ShowingList {
						fileName,ok:=n.ListFiles.SelectedItem().(notes.NoteItem)
						if ok{
								filePath := filepath.Join(config.VaultDir, fileName.Title())
							fileData,err:=notes.ReadFile(filePath)
							if err != nil {
								log.Fatalln("Could not read file:", err)
								n.MsgText = fmt.Sprintf(" Could not read file: %v \n",err)
								n.MsgType = MessageError
								return n, clearMsgAfter()
							}

							// set the content to textarea
							n.ContentInput.SetValue(fileData)

							// open the file
						selectedFFilePath := filepath.Join(config.VaultDir, fileName.Title())
						

							file,err:=os.OpenFile(selectedFFilePath, os.O_RDWR, 0644)
							if err != nil {
								log.Fatalln("Could not open file:", err)
								n.MsgText = fmt.Sprintf(" Could not open file: %v \n",err)
								n.MsgType = MessageError
								return n, clearMsgAfter()
							}
							n.currentFile = file
							n.ShowingList = false

						}
					}
					return n,nil


				// delete file from list
				case "ctrl+d":
					if n.ShowingList {
						selectedItem,ok :=n.ListFiles.SelectedItem().(notes.NoteItem)

						if ok {
							selectedFFilePath := filepath.Join(config.VaultDir, selectedItem.Title())
							if err := notes.DeleteFile(selectedFFilePath); err != nil {
								n.MsgText = fmt.Sprintf(" Could not delete file: %v \n",err)
								n.MsgType = MessageError
								return n, clearMsgAfter()

							}
							n.MsgText = fmt.Sprintf("🗑️ Deleted file: %s \n", selectedItem.Title())
							n.MsgType = MessageSuccess
							// refresh list
							fileLists := notes.ListFiles()
							lists := list.New(fileLists, list.NewDefaultDelegate(), n.Width, n.Height)
							lists.Title = "🗒️ ALL NOTES LIST: "
							n.ListFiles = lists

						}
					}

				// back (without saving)
				case "esc":
					// file isnt created yet
					if n.IsTitleInput {
						n.IsTitleInput = false
					}
					//  Back from writing the file:
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
					// back to delete mode

					return n,nil
			}

		}
		// now we can write inside titleInput using textinput
		if n.IsTitleInput {
			n.TitleInput , cmd= n.TitleInput.Update(msg)
			return n,cmd
		}

		// now we can write inside file  using textarea
		if n.currentFile!= nil {
			n.ContentInput, cmd= n.ContentInput.Update(msg)
			return n,cmd
		}

		// we can intereacct with list using list
		if n.ShowingList {
			n.ListFiles , cmd = n.ListFiles.Update(msg)
			return n,cmd
		}
		return n,cmd
}

// used to signal message expiry
type clearMessageMsg struct{}

// clear message after a timeout
func clearMsgAfter() tea.Cmd {
	return tea.Tick(messageTimeout, func(time.Time) tea.Msg {
		return clearMessageMsg{}
	})
}


