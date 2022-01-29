package main

import (
	"autobubble/api"
	"autobubble/userinterface"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	api := api.HaApi{
		HAUrl:   os.Getenv("HA_URL"),
		HAToken: os.Getenv("HA_TOKEN"),
	}
	status := api.GetAPIState()
	if status != `{"message": "API running."}` {
		log.Fatalf("API was not reachable: %v", status)
	}
	p := tea.NewProgram(userinterface.InitialModel(api))
	if err := p.Start(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
