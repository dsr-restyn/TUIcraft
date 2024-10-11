package pkg

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
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

// Constants
const (
	Name = iota
	Role
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
		Pbar bool
	}

	Choices struct {
		ChoicesSlice []Choice
	}

	Menu struct {
		Name    string
		Choices Choices
		Chosen  bool
	}

	MenuChoices struct {
		Choices Choices
	}

	GameChoices struct {
		Choices Choices
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

	CombatEntity struct {
		Name  string
		Stats stats
		Gold  int
		Items []Item
	}

	stats struct {
		Health int
		Mana   int
		Level  int
		Dmg    int
		Def    int
		Exp    float64
	}

	Player struct {
		Name      string
		Role      string
		Stats     stats
		Gold      int
		Inventory []Item
	}
)

func (p Player) IsLoaded() bool {
	return p.Name != "" && p.Role != ""
}

func (p *Player) LiquidateInventory() int {
	gold := p.Gold
	for _, item := range p.Inventory {
		gold = gold + item.SalePrice
	}
	p.Inventory = []Item{}
	log.Printf("Liquidating, gold earned: %d", gold)
	return gold
}

func (cs Choices) contains(choice Choice) bool {
	for _, c := range cs.ChoicesSlice {
		if c == choice {
			return true
		}
	}
	return false
}

func (cs Choices) GetChoiceByName(name string) Choice {
	for _, choice := range cs.ChoicesSlice {
		if choice.Name == name {
			return choice
		}
	}
	return Choice{}
}

func (cs Choices) GetChoiceById(id int) Choice {
	for _, choice := range cs.ChoicesSlice {
		if choice.Id == id {
			return choice
		}
	}
	return Choice{}
}

// Model
type Model struct {
	Choice           Choice
	Chosen           bool
	Ticks            int
	Frames           int
	Progress         float64
	Loaded           bool
	Quitting         bool
	MainMenu         Menu
	GameMenu         Menu
	ItemTable        ItemTable
	CombatEncounters []CombatEntity
	DroppedItems     []Item
	inputs           []textinput.Model
	focused          int
	Player           Player
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

func (m *Model) savePlayer() error {
	file, err := os.Create("player_save.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(m.Player)
}

func (m *Model) loadPlayer() error {
	file, err := os.Open("player_save.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&m.Player)
}

func (m *Model) saveModel() error {
	file, err := os.Create("model_save.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(m)
}

// nextInput focuses the next input field
func (m *Model) nextInput() {
	m.focused++
	// Wrap around
	if m.focused >= len(m.inputs) {
		m.focused = 0
	}
}

// prevInput focuses the previous input field
func (m *Model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func (m *Model) nextChoice(choices Choices) {
	c := m.Choice.Id
	newChoice := choices.GetChoiceById(c + 1)
	if newChoice.Id == 0 {
		newChoice = choices.ChoicesSlice[0]
	}
	m.Choice = newChoice
}

func (m *Model) previousChoice(choices Choices) {
	c := m.Choice.Id
	newChoice := choices.GetChoiceById(c - 1)
	if newChoice.Id == 0 {
		newChoice = choices.ChoicesSlice[len(choices.ChoicesSlice)-1]
	}
	m.Choice = newChoice
}

func (m *Model) InitPlayer() {
	m.Player = Player{
		Stats: stats{
			Health: 20,
			Mana:   5,
			Level:  1,
			Dmg:    5,
			Def:    2,
			Exp:    0,
		},
		Gold:      0,
		Inventory: []Item{},
	}
}

func InitEncounters() []CombatEntity {
	var encounters []CombatEntity
	for i := 0; i < 5; i++ {
		encounter := CombatEntity{
			Name: "Goblin",
			Stats: stats{
				Health: 10,
				Mana:   0,
				Level:  1,
				Dmg:    3,
				Def:    1,
				Exp:    10,
			},
			Gold: rand.New(rand.NewSource(time.Now().Unix())).Intn(50),
		}
		encounters = append(encounters, encounter)
	}
	return encounters
}

func InitalModel() Model {
	initItemTable := initItemTable()

	initMainMenu := initMainMenu()

	initGameMenu := initGameMenu()

	initEncounters := InitEncounters()

	var inputs []textinput.Model = make([]textinput.Model, 2)
	inputs[Name] = textinput.New()
	inputs[Name].Placeholder = ""
	inputs[Name].Focus()
	inputs[Name].Prompt = ""
	inputs[Name].Width = 11
	inputs[Name].CharLimit = 10
	inputs[Role] = textinput.New()
	inputs[Role].Placeholder = ""
	inputs[Role].Prompt = ""
	inputs[Role].Width = 11
	inputs[Role].CharLimit = 10

	return Model{
		Choice:           initMainMenu.Choices.ChoicesSlice[0],
		Chosen:           false,
		Ticks:            10,
		Frames:           0,
		Progress:         0,
		Loaded:           false,
		Quitting:         false,
		MainMenu:         initMainMenu,
		GameMenu:         initGameMenu,
		CombatEncounters: initEncounters,
		ItemTable:        initItemTable,
		Player:           Player{},
		inputs:           inputs,
		focused:          0,
	}
}

func initItemTable() ItemTable {
	initItemTable := ItemTable{
		Items: []Item{
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
	return initItemTable
}

func initMainMenu() Menu {
	initMainMenu := Menu{
		Choices: Choices{
			[]Choice{
				{Name: "New Game", Id: 1, Pbar: false},
				{Name: "Load Game", Id: 2, Pbar: false},
			},
		},
		Name: "Main Menu",
	}
	return initMainMenu
}

func initGameMenu() Menu {
	initGameMenu := Menu{
		Choices: Choices{
			[]Choice{
				{Name: "Wander Around", Id: 1, Pbar: true},
				{Name: "Fight Some Stuff", Id: 2, Pbar: true},
				{Name: "Talk to a Stranger", Id: 3, Pbar: true},
				{Name: "Take a Nap", Id: 4, Pbar: true},
				{Name: "Go to The Store", Id: 5, Pbar: false},
				{Name: "Craft", Id: 6, Pbar: false},
				{Name: "View Inventory", Id: 7, Pbar: false},
			},
		},
		Name: "Game Menu",
	}
	return initGameMenu
}
