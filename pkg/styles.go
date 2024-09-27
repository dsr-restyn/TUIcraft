package pkg

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	hotPink           = lipgloss.Color("#FF06B7")
	darkGray          = lipgloss.Color("#767676")
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
	dotChar           = " • "
)

var (
	keywordStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	commonRarityStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#918a90"))
	uncommonRarityStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#0588ed"))
	rareRarityStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#fcd112"))
	epicRarityStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#fc12cd"))
	legendaryRarityStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fc1212"))
	subtleStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ticksStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	checkboxStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	progressEmpty        = subtleStyle.Render(progressEmptyChar)
	dotStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle            = lipgloss.NewStyle().MarginLeft(2)
	spinnerStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	// animatedStyle        = lipgloss.NewStyle().MarginLeft(2).Width(1)

	// Gradient colors we'll use for the progress bar
	ramp = MakeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

// Utils

// Generate a blend of colors.
func MakeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(ColorToHex(c))))
	}
	return
}

// Convert a colorful.Color to a hexadecimal format.
func ColorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", ColorFloatToHex(c.R), ColorFloatToHex(c.G), ColorFloatToHex(c.B))
}

// Helper function for converting colors to hex. Assumes a value between 0 and
// 1.
func ColorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}
