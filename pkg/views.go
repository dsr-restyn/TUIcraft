package pkg

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func (m Model) View() string {
	var s string
	if m.Quitting {
		s += "\nGoodbye!\n\n"
		return s
	}

	if m.Player.Name == "" {
		if !m.Chosen {
			s = menuChoicesView(m)
		} else {
			switch m.Choice.Name {
			case "New Game":
				s = newGameView(m)
			case "Load Game":
				log.Print("Loading game, going back to menu to check for player name...")
				s = menuChoicesView(m)
			}
		}
	} else {
		log.Print("Player name found, showing game choices...")
		s = gameChoicesView(m)
	}

	// if !m.Chosen {
	// 	s = menuChoicesView(m)
	// } else if m.Choice.Name == "New Game" {
	// 	s = newGameView(m)
	// } else if m.Choice.Name == "Load Game" {
	// 	s = gameChoicesView(m)
	// }

	return mainStyle.Render("\n" + s + "\n\n")
}

func menuChoicesView(m Model) string {
	tpl := "Welcome to the TuiCraft!\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render(" Use j/k to select") + dotStyle + subtleStyle.Render("Press enter to confirm") + dotStyle + subtleStyle.Render("Press q, esc, or ctrl+c to quit")
	c := m.Choice.Id
	var choices string
	for _, choice := range m.MenuChoices.Choices.ChoicesSlice {
		choices += "\n"
		choices += checkbox(choice.Name, c == choice.Id)
	}

	return fmt.Sprintf(tpl, choices)
}

func newGameView(m Model) string {
	tpl := "New Game!\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render(" Use j/k to select") + dotStyle + subtleStyle.Render("Press enter to confirm") + dotStyle + subtleStyle.Render("Press q, esc, or ctrl+c to quit")
	textFields := []string{
		inputStyle.Width(6).Render("Name ↴"),
		m.inputs[Name].View(),
		inputStyle.Width(6).Render("Role ↴"),
		m.inputs[Role].View(),
	}
	return fmt.Sprintf(tpl, strings.Join(textFields, "\n"))
}

func gameChoicesView(m Model) string {
	c := m.Choice

	tpl := "\nWelcome, %s: Let's Begin\n\n"
	tpl += "%s\n\n"
	tpl += "Task Completes in %s seconds\n\n"
	tpl += subtleStyle.Render(" Use j/k to select") + dotStyle + subtleStyle.Render("Press enter to confirm") + dotStyle + subtleStyle.Render("Press q, esc, or ctrl+c to quit")

	var choices string
	for _, choice := range m.GameChoices.Choices.ChoicesSlice {
		choices += "\n"
		choices += checkbox(choice.Name, c.Id == choice.Id)
	}

	return fmt.Sprintf(tpl, m.Player.Name, choices, ticksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[✔] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func progressbar(percent float64) string {
	w := float64(progressBarWidth)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += ramp[i].Render(progressFullChar)
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}
