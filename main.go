package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matg94/go-ha/api"
	"github.com/matg94/go-ha/userinterface"

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
