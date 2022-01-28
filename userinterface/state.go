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
	shownEntities  []*api.Entity
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
			if !m.selected {
				if m.categoryCursor > 0 {
					m.categoryCursor--
				}
			} else {
				if m.entityCursor > 0 {
					m.entityCursor--
				}
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if !m.selected {
				if m.categoryCursor < len(m.categories)-1 {
					m.categoryCursor++
				}
			} else {
				if m.entityCursor < len(m.shownEntities)-1 {
					m.entityCursor++
				}
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			if !m.selected {
				m.entityCursor = 0
				m.selected = true
				m.shownEntities = []*api.Entity{}
				for _, entity := range m.entities {
					if entity.Category == m.categories[m.categoryCursor] {
						m.shownEntities = append(m.shownEntities, entity)
					}
				}
			} else {
				for i, entity := range m.shownEntities {
					if m.entityCursor == i {
						api.ActivateService(entity.Category, "toggle", entity.ID)
						break
					}
				}
				newState, _ := api.GetEntityStates()
				m.UpdateEntities(newState)
			}
		case "b":
			if m.selected {
				m.shownEntities = []*api.Entity{}
				m.entityCursor = 0
				m.selected = false
			}
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

		if m.selected {

			for i, entity := range m.shownEntities {
				curs := " "
				if entity.Category == choice {
					if m.entityCursor == i {
						curs = ">"
					}

					s += fmt.Sprintf("%s - %s : %s\n", curs, entity.ID, entity.State)
				}
			}
		}
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
