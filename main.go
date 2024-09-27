package main

import (
	"fmt"
	"tuicraft/pkg"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initalItemTable := pkg.ItemTable{
		Items: []pkg.Item{
			{Name: "Crazy Orb", Desc: "Makes ya crazy!", SalePrice: 1000},
			{Name: "Magic Sword", Desc: "Slice and Dice", SalePrice: 500},
			{Name: "Golden Key", Desc: "Must open a Golden Door", SalePrice: 250},
			{Name: "Rusty Dagger", Desc: "Like a nice dagger, except rusty", SalePrice: 50},
			{Name: "Shiny Shield", Desc: "You can see your face in it", SalePrice: 100},
			{Name: "Old Book", Desc: "Dusty", SalePrice: 25},
			{Name: "Strange Potion", Desc: "What IS a normal potion?", SalePrice: 200},
			{Name: "Silver Coin", Desc: "Not a gold coin", SalePrice: 10},
		},
	}

	initalMenuChoices := pkg.MenuChoices{
		Choices: []pkg.Choice{
			{Name: "New Game", Id: 1},
			{Name: "Load Game", Id: 2},
		},
	}

	initialModel := pkg.Model{
		Choice:      initalMenuChoices.Choices[0],
		Chosen:      false,
		Ticks:       10,
		Frames:      0,
		Progress:    0,
		Loaded:      false,
		Quitting:    false,
		MenuChoices: initalMenuChoices,
		ItemTable:   initalItemTable,
		Player: pkg.Player{
			Name:       "Zelda",
			Role:       "Warrior",
			Health:     100,
			Mana:       50,
			Level:      1,
			Experience: 0,
			Inventory:  []pkg.Item{},
		},
	}

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
