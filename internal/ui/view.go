package ui

import (
	"fmt"
)

func (n NoteModel) View() string {
	theme:=n.Theme

	var output string
	// 1️⃣ Render message box (if any)
	// output += n.renderMessageBox()

	// show msg bo at top if any msg is active
	if n.MsgType != MessageNone  && n.MsgText != "" {
		output += theme.RenderMsgBox(n.MsgType,n.MsgText + "\n")
	}

	// if showing list of files
	if n.ShowingList  {
		output += theme.BorderStyle.Render(n.ListFiles.View())

		return output
	}

	// if showing title input
	if n.IsTitleInput {
		titleName := theme.TitleStyle.Render("\n📝 Write a file name : ")
		output +=  fmt.Sprintf("%s\n%s\n", titleName, n.TitleInput.View())
		return output
	}

	// editing mode
	if n.currentFile != nil {
		currentFileContent:= theme.ContentStyle.Render(n.ContentInput.View())
		modeOfFile:= theme.TitleStyle.Render(fmt.Sprintf("Editing File: %s", n.currentFile.Name()))
		output+= fmt.Sprintf("%s\n%s\n", modeOfFile,currentFileContent)
		return output
	}

	// default idle view
	welcomeMsg := theme.PromptStyle.Render(fmt.Sprintf("%s \n","Welcome to CLI NoteTaker "))
	commands := instructions(theme)
	output += fmt.Sprintf("\n%s \n%s\n",welcomeMsg,commands)
	return output
}

// instruction view
func instructions(th Theme) string {
	help := th.InactiveStyle.Render("Ctrl+N: new file | Ctrl+L: List | Esc: back | Ctrl+S: save | Ctrl+Q: quit")
	return help
}

