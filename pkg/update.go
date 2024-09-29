package pkg

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))
	if msg, ok := msg.(tea.KeyMsg); ok {
		if m.Choice.Name != "New Game" {
			k := msg.String()
			if k == "q" || k == "esc" || k == "ctrl+c" {
				m.Quitting = true
				return m, tea.Quit
			}
		}
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Chosen && m.Choice.Name == "New Game" {
			switch msg.Type {
			case tea.KeyEnter:
				if m.focused == len(m.inputs)-1 {
					m.Player.Name = m.inputs[0].Value()
					m.Player.Role = m.inputs[1].Value()
					if err := m.savePlayer(); err != nil {
						fmt.Println("Error saving player:", err)
						return m, tick()
					}
					return m, tick()
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
			m.inputs[m.focused].Focus()
			for i := range m.inputs {
				m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
			}
			return m, tea.Batch(cmds...)
		} else if m.Chosen && m.Choice.Name == "Load Game" {
			m.loadPlayer()
			return m, tick()
		}
		k := msg.String()
		switch k {
		case "j", "down":
			if !m.Chosen {
				id := m.Choice.Id
				newChoice := m.MenuChoices.GetChoiceById(id + 1)
				if newChoice.Id != 0 {
					m.Choice = newChoice
					return m, tick()
				} else {
					return m, tick()
				}
			}
		case "k", "up":
			if !m.Chosen {
				id := m.Choice.Id
				newChoice := m.MenuChoices.GetChoiceById(id - 1)
				if newChoice.Id != 0 {
					m.Choice = newChoice
					return m, tick()
				} else {
					return m, tick()
				}
			}
		case "enter":
			if !m.Chosen {
				m.Chosen = true
				return m, tick()
			}

		}

		return m, tick()

	case frameMsg:
		return m, frame()
	case tickMsg:
		return m, tick()
	}

	return m, nil
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
