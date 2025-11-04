package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/config"
	"github.com/thissidemayur/3-golang-CLI_Note_Taker/internal/ui"
)

func main() {
	config.InitDirPath()

	// bubble tea 
	p:=tea.NewProgram(ui.InitializeModel())
	if _,err:=p.Run(); err != nil {
		fmt.Println("Error while Running Bubble Tea framework' programme: ",err)
		os.Exit(1)
	}

}
