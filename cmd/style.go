package cmd

import "github.com/charmbracelet/lipgloss"

const (
	minBottomHeight = 9
	maxWidth        = 48
)

var (
	lava      = lipgloss.Color("#F56E0F")
	void      = lipgloss.Color("#151419")
	dust      = lipgloss.Color("#878787")
	slate     = lipgloss.Color("#262626")
	snow      = lipgloss.Color("#FBFBFB")
	palette1  = lipgloss.Color("#E2E2DF")
	palette2  = lipgloss.Color("#D2D2CF")
	palette3  = lipgloss.Color("#E2CFC4")
	palette4  = lipgloss.Color("#F7D9C4")
	palette5  = lipgloss.Color("#FAEDCB")
	palette6  = lipgloss.Color("#C9E4DE")
	palette7  = lipgloss.Color("#C6DEF1")
	palette8  = lipgloss.Color("#DBCDF0")
	palette9  = lipgloss.Color("#F2C6DE")
	palette10 = lipgloss.Color("#F9C6C9")

	styleTitle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(palette10)).
			Bold(true).
			Align(lipgloss.Center).
			Width(maxWidth + 2)
	styleSelected = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(palette2)).
			Bold(true).
			Width(42).PaddingLeft(2)
	styleNormal = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(palette1)).
			Width(42).
			PaddingLeft(2)
	styleApp = lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(palette2))
	styleHash = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(palette7)).
			Bold(true).
			Width(8).
			Align(lipgloss.Center)
	styleString = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(palette8)).
			Bold(true).
			Width(8).
			Align(lipgloss.Center)
	styleSet = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(palette4)).
			Bold(true).
			Width(8).
			Align(lipgloss.Center)
	styleList = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(palette6)).
			Bold(true).
			Width(8).
			Align(lipgloss.Center)
)
