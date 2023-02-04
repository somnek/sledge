package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	yellow    = lipgloss.Color("#f0ead2")
	melon     = lipgloss.Color("#E0F2E9")
	red       = lipgloss.Color("#EF798A")
	darkGreen = lipgloss.Color("#4B644A")
	brown     = lipgloss.Color("#49393B")
)

var (
	styTitle       = lipgloss.NewStyle().Foreground(melon).MarginLeft(10).Bold(true).Faint(true)
	styUnderline   = lipgloss.NewStyle().Foreground(yellow).MarginLeft(10).Faint(true)
	styItemDefault = lipgloss.NewStyle().Foreground(yellow)
	styItemSelect  = lipgloss.NewStyle().Foreground(yellow)
)

func (m model) View() string {
	var sBody string
	var arrow string
	var checked string
	var row string

	sTitle := styTitle.Render("Welcome to red list... 🥥")
	sUnder := styUnderline.Render("─────────────────────────")

	for i, k := range m.keys {
		// TODO:  do selectd later
		if m.cursor == i {
			arrow = ">"
			row = fmt.Sprintf("%s[%s] %s\n", arrow, checked, k)
		} else {
			row = fmt.Sprintf("%s[%s] %s\n", arrow, checked, k)
		}

		sBody += styItemDefault.Render(row)
	}

	return fmt.Sprintf(
		"\n%s\n%s\n%s\n",
		sTitle,
		sUnder,
		sBody,
	)
}
