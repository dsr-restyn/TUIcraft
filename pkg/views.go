package pkg

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
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
				s = menuChoicesView(m)
			}
		}
	} else {
		if !m.Chosen {
			s = gameChoicesView(m)
		} else {
			s = chosenView(m)
		}
	}
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

	tpl := "\nWelcome, %s The %s: Let's Begin\n\n"
	tpl += "%s\n\n"
	tpl += "Task Completes in %s seconds\n\n"
	tpl += subtleStyle.Render(" Use j/k to select") + dotStyle + subtleStyle.Render("Press enter to confirm") + dotStyle + subtleStyle.Render("Press q, esc, or ctrl+c to quit")

	var choices string
	for _, choice := range m.GameChoices.Choices.ChoicesSlice {
		choices += "\n"
		choices += checkbox(choice.Name, c.Id == choice.Id)
	}

	return fmt.Sprintf(tpl, m.Player.Name, m.Player.Role, choices, ticksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
}

func chosenView(m Model) string {
	var msg string
	label := m.Choice.Name
	done := "Done!"

	if m.GameChoices.Choices.contains(m.Choice) {
		switch m.Choice.Name {
		case "Wander Around":
			msg = "Wandering around...\n\n"
			label = "Wandering..."
			if len(m.DroppedItems) == 3 {
				done = fmt.Sprintf("You found a %s %s, a %s %s, and a %s %s!",
					renderRarity(m.DroppedItems[0].Rarity),
					keywordStyle.Render(m.DroppedItems[0].Name),
					renderRarity(m.DroppedItems[1].Rarity),
					keywordStyle.Render(m.DroppedItems[1].Name),
					renderRarity(m.DroppedItems[2].Rarity),
					keywordStyle.Render(m.DroppedItems[2].Name),
				)
			}
		case "Fight Some Stuff":
			msg = "You fight some stuff. It's a good time."
		case "Talk to a Stranger":
			msg = "You talk to a stranger. They're strange."
		case "Take a Nap":
			msg = "You take a nap. It's a good nap."
		case "Craft":
			msg = "You craft something. It's a good craft."
		case "Go to The Store":
			msg = "You visit the store, what would you like to do?"
			label = subtleStyle.Render("Press F to sell all")
		case "View Inventory":
			label = "Viewing Inventory...\n\n"
			var items []string
			for _, item := range m.Player.Inventory {
				items = append(items, fmt.Sprintf("%s %s\n%s", renderRarity(item.Rarity), keywordStyle.Render(item.Name), subtleStyle.Render(item.Desc)))
			}
			msg = "Inventory: \n\n" + strings.Join(items, "\n")
			msg += fmt.Sprintf("\n\n Gold: %d", m.Player.Gold)
			m.Ticks = 500
		default:
			msg = "You do something. It's a good something."
		}
		if m.Loaded {
			label = fmt.Sprintf("%s\n\nreturning to menu in %s seconds...", done, ticksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
			return msg + "\n\n" + label + "\n\n" + progressbar(m.Progress) + "%"
		} else {
			msg += "\n\n" + subtleStyle.Render("Loading...") + dotStyle
			return msg + "\n\n" + label + "\n\n" + progressbar(m.Progress) + "%"

		}
	} else {
		return menuChoicesView(m)
	}

}

func renderRarity(rarity string) string {
	switch rarity {
	case "Common":
		return commonRarityStyle.Render(rarity)
	case "Uncommon":
		return uncommonRarityStyle.Render(rarity)
	case "Rare":
		return rareRarityStyle.Render(rarity)
	case "Epic":
		return epicRarityStyle.Render(rarity)
	case "Legendary":
		return legendaryRarityStyle.Render(rarity)
	default:
		return lipgloss.NewStyle().Render(rarity)
	}
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
