package ui

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/notes"
)

type MsgType int
const (
	MessageNone MsgType = iota
	MessageError
	MessageInfo
	MessageSuccess
	MessageWarning
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
	MsgText string
	MsgType  MsgType

	// Window dimensions	
	Width  int
Height int
}


func InitializeModel () NoteModel {
	n := NoteModel{}
	// textinput instance
	ti := textinput.New()
	ti.Placeholder = "Enter the file name"
	ti.Focus()

	// textarea instance
	ta:= textarea.New()
	ta.Placeholder = "Start typing your notes here..."
	ta.Focus()

	// list instance
	fileLists := notes.ListFiles()
	lists := list.New(fileLists, list.NewDefaultDelegate(),n.Width, n.Height)
	lists.Title = "🗒️ ALL NOTES LIST: "

	return NoteModel{
		TitleInput: ti,
		IsTitleInput: false,
		ContentInput: ta,
		ListFiles: lists,
		ShowingList: false,
		Theme: NewTheme(),
		MsgText: "",
		MsgType: MessageNone,
	}
}


func (n NoteModel) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
     return tea.Batch(
        tea.EnterAltScreen,
        tea.WindowSize(), // proactive resize event
    )

}