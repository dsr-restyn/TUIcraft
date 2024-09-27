package pkg

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
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

		}
		return m, tick()

	case frameMsg:
		return m, frame()
	case tickMsg:
		return m, tick()
	}

	return m, nil
}
