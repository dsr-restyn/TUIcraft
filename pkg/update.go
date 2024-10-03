package pkg

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		if m.Choice.Name != "New Game" {
			k := msg.String()
			if k == "q" || k == "esc" || k == "ctrl+c" {
				m.Quitting = true
				return m, tea.Quit
			}
		}
	}

	if m.Player.Name == "" {
		if !m.Chosen {
			return m.updateChoices(msg)
		} else {
			switch m.Choice.Name {
			case "New Game":
				return m.updateInputs(msg)
			case "Load Game":
				m.loadPlayer()
				m.Chosen = false
				m.Choice = m.GameChoices.Choices.ChoicesSlice[0]
				return m.updateChoices(msg)
			}
		}

	} else {
		if m.MenuChoices.Choices.contains(m.Choice) {
			m.Choice = m.GameChoices.Choices.ChoicesSlice[0]
			return m.updateChoices(msg)
		}
		return m.updateChoices(msg)
	}
	return m, nil
}

// if !m.Chosen {
// 	return m.updateChoices(msg)
// } else if m.Choice.Name == "New Game" {
// 	return m.updateInputs(msg)
// }

// 	return m, nil
// }

func (m Model) updateInputs(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))
	if m.Chosen && m.Choice.Name == "New Game" {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				if m.focused == len(m.inputs)-1 {
					m.Player.Name = m.inputs[0].Value()
					m.Player.Role = m.inputs[1].Value()
					if err := m.savePlayer(); err != nil {
						fmt.Println("Error saving player:", err)
					}
					m.loadPlayer()
					m.Chosen = false
					log.Printf("Player: %v", m.Player)
				}
			case tea.KeyShiftTab, tea.KeyCtrlP:
				m.inputs[m.focused].Blur()
				m.prevInput()
			case tea.KeyTab, tea.KeyCtrlN:
				m.inputs[m.focused].Blur()
				m.nextInput()
			case tea.KeyCtrlC:
				return m, tea.Quit
			}
		}
		m.inputs[m.focused].Focus()
		for i := range m.inputs {
			m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		}

	}
	return m, tick()
}

func (m Model) updateChoices(msg tea.Msg) (Model, tea.Cmd) {
	c := m.Choice
	inMenu := m.MenuChoices.Choices.contains(c)
	log.Printf("In menu: %v", inMenu)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if inMenu {
				m.nextChoice(m.MenuChoices.Choices)
			} else {
				m.nextChoice(m.GameChoices.Choices)
			}
		case "k", "up":
			if inMenu {
				m.previousChoice(m.MenuChoices.Choices)
			} else {
				m.previousChoice(m.GameChoices.Choices)
			}
		case "q", "esc", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		case "enter":
			m.Chosen = true
			if m.Choice.Name == "Load Game" {
				m.loadPlayer()
				m.Chosen = false
			}
		}
	}
	return m, tick()
}

// nextInput focuses the next input field
func (m *Model) nextInput() {
	m.focused++
	// Wrap around
	if m.focused >= len(m.inputs) {
		m.focused = 0
	}
}

// prevInput focuses the previous input field
func (m *Model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func (m *Model) nextChoice(choices Choices) {
	c := m.Choice.Id
	newChoice := choices.GetChoiceById(c + 1)
	if newChoice.Id == 0 {
		newChoice = choices.ChoicesSlice[0]
	}
	m.Choice = newChoice
}

func (m *Model) previousChoice(choices Choices) {
	c := m.Choice.Id
	newChoice := choices.GetChoiceById(c - 1)
	if newChoice.Id == 0 {
		newChoice = choices.ChoicesSlice[len(choices.ChoicesSlice)-1]
	}
	m.Choice = newChoice
}
