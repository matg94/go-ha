package main

import (
	"autobubble/api"
	"autobubble/userinterface"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	entities, _ := api.GetEntityStates()
	p := tea.NewProgram(userinterface.InitialModel(entities))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
