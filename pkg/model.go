package pkg

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Vars
var (
	// Fixed-length slice for rarities
	rarities = [...]string{
		"Common",
		"Uncommon",
		"Rare",
		"Epic",
		"Legendary",
	}
)

// Message types
type (
	tickMsg struct {
		time time.Time
	}
	frameMsg struct{}
)

// Game types

type (
	Choice struct {
		Name string
		Id   int
	}

	MenuChoices struct {
		Choices []Choice
	}

	Item struct {
		Name      string
		Desc      string
		Rarity    string
		SalePrice int
	}

	ItemTable struct {
		Items []Item
	}

	Player struct {
		Name       string
		Role       string
		Health     int
		Mana       int
		Level      int
		Experience float64
		Inventory  []Item
	}
)

func (mc MenuChoices) GetChoiceById(id int) Choice {
	for _, choice := range mc.Choices {
		if choice.Id == id {
			return choice
		}
	}
	return Choice{}
}

// Model
type Model struct {
	Choice       Choice
	Chosen       bool
	Ticks        int
	Frames       int
	Progress     float64
	Loaded       bool
	Quitting     bool
	MenuChoices  MenuChoices
	ItemTable    ItemTable
	DroppedItems []Item
	Player       Player
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{time: time.Now()}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m Model) Init() tea.Cmd {
	return tick()
}

// func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	if msg, ok := msg.(tea.KeyMsg); ok {
// 		k := msg.String()
// 		if k == "q" || k == "esc" || k == "ctrl+c" {
// 			m.Quitting = true
// 			return m, tea.Quit
// 		}

// 		if k == "m" {
// 			m.Chosen = false
// 			m.Loaded = false
// 			m.Progress = 0
// 			m.Frames = 0
// 			m.Ticks = 10
// 			return m, tick()
// 		}
// 	}

// 	if !m.Chosen {
// 		return updateChoices(msg, m)
// 	} else if m.Chosen && m.Choice == 4 {
// 		return updateChoices(msg, m)
// 	}

// 	return updateChosen(msg, m)

// }

// func updateChoices(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "j", "down":
// 			if !m.MainChosen {
// 				m.MainChoice++
// 				if m.MainChoice > 1 {
// 					m.MainChoice = 1
// 				}
// 				return m, frame()
// 			}
// 			if m.Chosen && m.Choice == 4 {
// 				m.CraftChoice++
// 				if m.CraftChoice > 2 {
// 					m.CraftChoice = 2
// 				}

// 			} else {
// 				m.Choice++
// 				if m.Choice > 5 {
// 					m.Choice = 5
// 				}
// 			}

// 		case "k", "up":
// 			if !m.MainChosen {
// 				m.MainChoice--
// 				if m.MainChoice < 0 {
// 					m.MainChoice = 0
// 				}
// 				return m, frame()
// 			}
// 			if m.Chosen && m.Choice == 4 {
// 				m.CraftChoice--
// 				if m.CraftChoice < 0 {
// 					m.CraftChoice = 0
// 				}
// 			} else {
// 				m.Choice--
// 				if m.Choice < 0 {
// 					m.Choice = 0
// 				}
// 			}

// 		case "enter":
// 			if !m.MainChosen {
// 				m.MainChosen = true
// 				return m, frame()
// 			}
// 			if m.Choice == 4 && m.Chosen {
// 				m.CraftChosen = true
// 				return m, frame()
// 			} else {
// 				m.Chosen = true
// 				return m, frame()
// 			}
// 		}

// 	case tickMsg:
// 		if m.Chosen && m.Choice == 4 {
// 			if m.Ticks == 0 {
// 				m.CraftChosen = true
// 				return m, tick()
// 			}
// 		}
// 		if m.Ticks == 0 {
// 			m.Chosen = true
// 			m.Ticks = 10
// 			return m, tick()
// 		}
// 		m.Ticks--
// 		return m, tick()
// 	}

// 	return m, nil
// }

// func updateChosen(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
// 	switch msg.(type) {
// 	case frameMsg:
// 		if !m.Loaded {
// 			divisor := 100
// 			m.Frames++
// 			if m.Choice == 0 {
// 				divisor = 1000
// 				m.Progress = ease.Linear(float64(m.Frames) / float64(divisor))
// 			} else if m.Choice == 4 || m.Choice == 5 {
// 				m.Progress = 0
// 			} else {
// 				m.Progress = ease.OutBounce(float64(m.Frames) / float64(divisor))
// 			}
// 			if m.Progress >= 1 {
// 				m.Progress = 1
// 				m.Loaded = true
// 				m.Ticks = 3
// 				if m.Choice == 0 {
// 					rand.New(rand.NewSource(time.Now().UnixNano()))
// 					item_1 := rand.Intn(len(m.ItemTable.Items))
// 					item_2 := rand.Intn(len(m.ItemTable.Items))
// 					item_3 := rand.Intn(len(m.ItemTable.Items))
// 					rarity_1 := rarities[rand.Intn(len(rarities))]
// 					rarity_2 := rarities[rand.Intn(len(rarities))]
// 					rarity_3 := rarities[rand.Intn(len(rarities))]
// 					m.ItemTable.Items[item_1].Rarity = rarity_1
// 					m.ItemTable.Items[item_2].Rarity = rarity_2
// 					m.ItemTable.Items[item_3].Rarity = rarity_3
// 					m.DroppedItems = []Item{m.ItemTable.Items[item_1], m.ItemTable.Items[item_2], m.ItemTable.Items[item_3]}
// 					m.Player.Inventory = append(m.Player.Inventory, m.DroppedItems...)
// 				}
// 				return m, tick()
// 			}
// 		}
// 	case tickMsg:
// 		if m.Loaded {
// 			if m.Ticks == 0 {
// 				m.Chosen = false
// 				m.Loaded = false
// 				m.Progress = 0
// 				m.Frames = 0
// 				m.Ticks = 10
// 				return m, tick()
// 			}
// 			m.Ticks--
// 			return m, tick()
// 		} else {
// 			m.Ticks--
// 		}
// 	}
// 	return m, frame()
// }
