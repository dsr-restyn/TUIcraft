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
		if !m.Chosen {
			s = gameChoicesView(m)
		} else {
			s = chosenView(m)
		}
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

func chosenView(m Model) string {
	var msg string
	label := m.Choice.Name
	done := "Done!"

	if m.GameChoices.Choices.contains(m.Choice) {
		switch m.Choice.Name {
		case "Wander Around":
			msg = "You wander around the forest, looking for adventure."
		case "Fight Some Stuff":
			msg = "You fight some stuff. It's a good time."
		case "Talk to a Stranger":
			msg = "You talk to a stranger. They're strange."
		case "Take a Nap":
			msg = "You take a nap. It's a good nap."
		case "Craft":
			msg = "You craft something. It's a good craft."
		case "View Inventory":
			msg = "You view your inventory. It's a good inventory."
		default:
			msg = "You do something. It's a good something."
		}
		if m.Loaded {
			msg += "\n\n" + label + "\n\n" + done
		} else {
			msg += "\n\n" + subtleStyle.Render("Loading...") + dotStyle
		}
	} else {
		return menuChoicesView(m)
	}

	return msg
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
