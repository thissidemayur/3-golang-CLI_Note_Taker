package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/config"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/ui"
)

func main() {
	config.InitDirPath()

	
	

	log.Println("🪶 Starting CLI Note Taker")

	p := tea.NewProgram(ui.InitializeModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error while Running Bubble Tea program: %v \n", err)
	}
}
