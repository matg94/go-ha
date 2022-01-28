package main

import (
	"autobubble/api"
	"autobubble/userinterface"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	entities, err := api.GetEntityStates()
	if err != nil {
		fmt.Printf("API Call Error: %v", err)
	}
	p := tea.NewProgram(userinterface.InitialModel(entities))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
