package ui

import (
	"fmt"
)

func (n NoteModel) View() string {
	theme:=n.Theme

	// if showing list of files
	if n.ShowingList {
		return theme.BorderStyle.Render(n.ListFiles.View())
	}

	// if showing title input
	if n.IsTitleInput {
		titleName := theme.TitleStyle.Render("📝 Write a file name : \n")
		return fmt.Sprintf("%s\n%s", titleName, n.TitleInput.View())
	}

	// editing mode
	if n.currentFile != nil {
		return fmt.Sprintf("%s\n%s", theme.TitleStyle.Render(fmt.Sprintf("Editing File: %s", n.currentFile.Name())), theme.ContentStyle.Render(n.ContentInput.View()))
	}

	// default idle view
	welcomeMsg := "Welcome to the Note Taker CLI!"
	commands := instructions(theme)
	return theme.ContentStyle.Render(fmt.Sprintf("%s\n\n%s", welcomeMsg, commands))
}

// instruction view
func instructions(th Theme) string {
	help := th.InactiveStyle.Render("Ctrl+N: new file | Ctrl+L: List | Esc: back | Ctrl+S: save | Ctrl+Q: quit")
	return help
}