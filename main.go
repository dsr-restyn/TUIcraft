package main

import (
	"fmt"
	"tuicraft/pkg"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(pkg.InitalModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
