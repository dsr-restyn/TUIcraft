package pkg

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/ease"
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

			if k == "m" {
				m.Chosen = false
				return m.updateChoices(msg)
			}
		}
	}

	if !m.Player.IsLoaded() {
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
		} else {
			if !m.Chosen {
				return m.updateChoices(msg)
			} else {
				return m.updateChosen(msg)
			}
		}
	}
	return m, nil
}

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
	// TODO: Figure out more elegant way to handle player load state
	c := m.Choice
	inMenu := m.MenuChoices.Choices.contains(c)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if inMenu {
				m.nextChoice(m.MenuChoices.Choices)
			} else {
				m.nextChoice(m.GameChoices.Choices)
				return m, frame()
			}
		case "k", "up":
			if inMenu {
				m.previousChoice(m.MenuChoices.Choices)
			} else {
				m.previousChoice(m.GameChoices.Choices)
				return m, frame()
			}
		case "q", "esc", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		case "enter":
			m.Chosen = true
			if m.Choice.Name == "Load Game" && !m.Player.IsLoaded() {
				m.loadPlayer()
				m.Chosen = false
			}
			return m, frame()
		}
	case tickMsg:
		m.Ticks--
		if m.Ticks <= 0 {
			m.Chosen = true
			m.Ticks = 100
		}
		return m, tick()
	}
	return m, nil
}

func (m Model) updateChosen(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {
	case frameMsg:
		log.Printf("frameMsg received: %d", m.Frames)
		if !m.Loaded {
			divisor := 100
			m.Frames++
			m.Progress = ease.OutBounce(float64(m.Frames) / float64(divisor))
			if m.Progress >= 1 {
				m.Progress = 1
				m.Loaded = true
				m.Ticks = 3
				if m.Choice.Name == "" {
					rand.New(rand.NewSource(time.Now().UnixNano()))
					item_1 := rand.Intn(len(m.ItemTable.Items))
					item_2 := rand.Intn(len(m.ItemTable.Items))
					item_3 := rand.Intn(len(m.ItemTable.Items))
					rarity_1 := rarities[rand.Intn(len(rarities))]
					rarity_2 := rarities[rand.Intn(len(rarities))]
					rarity_3 := rarities[rand.Intn(len(rarities))]
					m.ItemTable.Items[item_1].Rarity = rarity_1
					m.ItemTable.Items[item_2].Rarity = rarity_2
					m.ItemTable.Items[item_3].Rarity = rarity_3
					m.DroppedItems = []Item{m.ItemTable.Items[item_1], m.ItemTable.Items[item_2], m.ItemTable.Items[item_3]}
					m.Player.Inventory = append(m.Player.Inventory, m.DroppedItems...)
				}
				return m, tick()
			}
		}
	case tickMsg:
		if m.Loaded {
			if m.Ticks == 0 {
				m.Chosen = false
				m.Loaded = false
				m.Progress = 0
				m.Frames = 0
				m.Ticks = 100
				return m, tick()
			}
			m.Ticks--
			return m, tick()
		} else {
			m.Ticks--
		}
	}
	return m, tick()
}
