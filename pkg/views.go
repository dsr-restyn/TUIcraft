package pkg

import (
	"fmt"
	"math"
	"strings"
)

func (m Model) View() string {
	var s string
	if m.Quitting {
		s += "\nGoodbye!\n\n"
		return s
	}

	if !m.Chosen {
		s = menuView(m)
	} else if m.Choice.Name == "New Game" {
		s = newGameView(m)
	} else {
		s = "You chose... wisely."
	}

	return mainStyle.Render("\n" + s + "\n\n")
}

func menuView(m Model) string {
	tpl := "Welcome to the TuiCraft!\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render(" Use j/k to select") + dotStyle + subtleStyle.Render("Press enter to confirm") + dotStyle + subtleStyle.Render("Press q, esc, or ctrl+c to quit")
	c := m.Choice.Id
	var choices string
	for _, choice := range m.MenuChoices.Choices {
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
		inputStyle.Width(30).Render("Name: \n"),
		inputStyle.Width(30).Render("Role: \n"),
	}
	return fmt.Sprintf(tpl, strings.Join(textFields, "\n"))
}

// func choicesView(m Model) string {
// 	c := m.Choice

// 	tpl := "\n Let's Begin\n\n"
// 	tpl += "%s\n\n"
// 	tpl += "Task Completes in %s seconds\n\n"
// 	tpl += subtleStyle.Render(" Use j/k to select") + dotStyle + subtleStyle.Render("Press enter to confirm") + dotStyle + subtleStyle.Render("Press q, esc, or ctrl+c to quit")

// 	choices := fmt.Sprintf(
// 		"%s\n%s\n%s\n%s\n%s\n%s\n",
// 		checkbox("Wander Around", c == 0),
// 		checkbox("Fight Some Stuff", c == 1),
// 		checkbox("Talk to a Stranger", c == 2),
// 		checkbox("Take a Nap", c == 3),
// 		checkbox("Craft", c == 4),
// 		checkbox("View Inventory", c == 5),
// 	)

// 	return fmt.Sprintf(tpl, choices, ticksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
// }

// func chosenView(m Model) string {
// 	var msg string
// 	label := "Working on it..."
// 	done := "Done!"

// 	switch m.Choice {
// 	case 0:
// 		msg = "Wandering around...\n\n"
// 		label = "Wandering..."
// 		if len(m.DroppedItems) == 3 {
// 			done = fmt.Sprintf("You found a %s %s, a %s %s, and a %s %s!",
// 				renderRarity(m.DroppedItems[0].Rarity),
// 				keywordStyle.Render(m.DroppedItems[0].Name),
// 				renderRarity(m.DroppedItems[1].Rarity),
// 				keywordStyle.Render(m.DroppedItems[1].Name),
// 				renderRarity(m.DroppedItems[2].Rarity),
// 				keywordStyle.Render(m.DroppedItems[2].Name),
// 			)
// 		}
// 	case 1:
// 		msg = "Fighting some stuff...\n\n"
// 		label = "Fighting..."
// 		done = "You won the fight!"
// 	case 2:
// 		msg = "Talking to a stranger...\n\n"
// 		label = "Talking..."
// 		done = "You made a new friend!"
// 	case 3:
// 		msg = "Taking a nap...\n\n"
// 		label = "Napping..."
// 		done = "You feel refreshed!"
// 	case 4:
// 		msg = craftingView(m)
// 		label = "Crafting...intensly..."
// 		done = "You crafted a new item!"
// 	case 5:
// 		// TODO: Utilize bubbles list component for a cooler inventory view
// 		label = "Viewing Inventory...\n\n"
// 		var items []string
// 		for _, item := range m.Player.Inventory {
// 			items = append(items, fmt.Sprintf("%s %s\n%s", renderRarity(item.Rarity), keywordStyle.Render(item.Name), subtleStyle.Render(item.Desc)))
// 		}
// 		msg = "Inventory: \n\n" + strings.Join(items, "\n")
// 		m.Loaded = true
// 		m.Ticks = 500
// 	default:
// 		msg = "You chose... wisely."
// 	}

// 	if m.Choice != 5 && m.Loaded {
// 		label = fmt.Sprintf("%s, returning to menu in %s seconds...", done, ticksStyle.Render(fmt.Sprintf("%d", m.Ticks)))
// 		return msg + "\n\n" + label + "\n\n" + progressbar(m.Progress) + "%"

// 	} else if !m.Loaded {
// 		return msg + "\n\n" + label + "\n\n" + progressbar(m.Progress) + "%"
// 	} else {
// 		return msg + "\n\n" + label + "\n\n" + progressbar(m.Progress) + "%"
// 	}

// }

// func craftingView(m Model) string {
// 	c := m.CraftChoice
// 	// msg := "You crafted a new item!\n\n"

// 	tpl := "\n Let's Begin\n\n"

// 	choices := fmt.Sprintf(
// 		"%s\n%s\n%s\n",
// 		checkbox("Craft a Sword", c == 0),
// 		checkbox("Craft a Wand", c == 1),
// 		checkbox("Craft a Shield", c == 2),
// 	)

// 	return tpl + "\n\n" + choices + "\n\n" + progressbar(m.Progress) + "%"

// }

// func menuView(m Model) string {
// 	c := m.MainChoice
// 	tmp := "Welcome to the game!\n\n"

// 	choices := fmt.Sprintf(
// 		"%s\n%s\n",
// 		checkbox("New Game", c == 0),
// 		checkbox("Load Game", c == 1),
// 	)

// 	return tmp + "\n\n" + choices + "\n\n"
// }

// func renderRarity(rarity string) string {
// 	switch rarity {
// 	case "Common":
// 		return commonRarityStyle.Render(rarity)
// 	case "Uncommon":
// 		return uncommonRarityStyle.Render(rarity)
// 	case "Rare":
// 		return rareRarityStyle.Render(rarity)
// 	case "Epic":
// 		return epicRarityStyle.Render(rarity)
// 	case "Legendary":
// 		return legendaryRarityStyle.Render(rarity)
// 	default:
// 		return lipgloss.NewStyle().Render(rarity)
// 	}
// }

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
