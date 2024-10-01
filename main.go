package main

import (
	"fmt"
	"os"
	"tuicraft/pkg"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	os.Setenv("DEBUG", "1")
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
	p := tea.NewProgram(pkg.InitalModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
