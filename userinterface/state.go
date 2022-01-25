package userinterface

import (
	"autobubble/api"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	categories     []string
	selected       bool
	categoryCursor int
	entityCursor   int
	entities       []*api.Entity
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.categoryCursor > 0 {
				m.categoryCursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.categoryCursor < len(m.categories)-1 {
				m.categoryCursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.categories {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.categoryCursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
