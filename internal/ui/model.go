package ui

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)


type NoteModel struct {
	// notes []NoteItem
	TitleInput  textinput.Model
	IsTitleInput bool
	ContentInput textarea.Model
	currentFile *os.File
	ListFiles      list.Model
	ShowingList    bool

	Theme Theme
	// AppError error

}


func InitializeModel () NoteModel {
	// textinput instance
	ti := textinput.New()
	ti.Placeholder = "Enter the file name"
	ti.Focus()

	// textarea instance
	ta:= textarea.New()
	ta.Placeholder = "Start typing your notes here..."
	ta.Focus()

	// list instance
	noteList:= list.New([]list.Item{}, list.NewDefaultDelegate(), 0,0)
	noteList.Title="🗒️ ALL NOTES LIST: "
	
	return NoteModel{
		TitleInput: ti,
		IsTitleInput: false,
		ContentInput: ta,
		ListFiles: noteList,
		ShowingList: false,
		Theme: NewTheme(),
	}
}


func (n NoteModel) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}