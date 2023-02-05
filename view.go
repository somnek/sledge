package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	yellow    = lipgloss.Color("#f0ead2")
	melon     = lipgloss.Color("#E0F2E9")
	red       = lipgloss.Color("#EF798A")
	darkGreen = lipgloss.Color("#4B644A")
	brown     = lipgloss.Color("#49393B")
	darkGray  = lipgloss.Color("#595959")
)

var (
	styTitle     = lipgloss.NewStyle().Foreground(brown).MarginLeft(10).Bold(true)
	styUnderline = lipgloss.NewStyle().Foreground(red).MarginLeft(10).Bold(true)
	styItem      = lipgloss.NewStyle().Foreground(darkGreen)
	styStatus    = lipgloss.NewStyle().Foreground(darkGray)
	styValue     = lipgloss.NewStyle().Foreground(darkGray)
)

func (m model) View() string {
	var sBody string
	var body string

	// -----> header
	sTitle := styTitle.Render("Welcome to red list... 🥥")
	sUnder := styUnderline.Render("─────────────────────────")

	// -----> body
	for i, k := range m.items {
		var arrow string   // " " or >
		var checked string //" "  or x

		if m.cursor == i {
			arrow = ">"
			checked = "x"
		} else {
			arrow = " "
			checked = " "
		}

		body += fmt.Sprintf("%s[%s] %s\n", arrow, checked, k)
	}
	sBody += styItem.Render(body)

	// -----> footer
	ctx := context.Background()

	value := get(ctx, m.items[m.cursor])
	sValue := styValue.Render(value)
	sStatus := styStatus.Render(m.status)

	return fmt.Sprintf(
		"\n%s\n%s\n%s\n%s\n\n%s",
		sTitle,
		sUnder,
		sBody,
		sValue,
		sStatus,
	)
}
