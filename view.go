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
	styTitle         = lipgloss.NewStyle().Foreground(melon).MarginLeft(10).Bold(true)
	styUnderline     = lipgloss.NewStyle().Foreground(red).MarginLeft(10).Bold(true)
	styBodyContainer = lipgloss.NewStyle().Foreground(yellow).MarginLeft(10)
	styStatus        = lipgloss.NewStyle().Foreground(darkGray)
	styBox           = lipgloss.NewStyle().Foreground(melon)
	styKey           = lipgloss.NewStyle().Foreground(yellow)
	styValue         = lipgloss.NewStyle().Foreground(darkGray)
	styArrow         = lipgloss.NewStyle().Foreground(red).Render
)

func (m model) View() string {
	var sBody string
	var body string

	// -----> header
	sTitle := styTitle.Render("Welcome to red list... 🥥")
	sUnder := styUnderline.Render("─────────────────────────")

	// -----> body
	for i, k := range m.items {
		var arrow string // " " or >
		var box string   //[ ] or [x]

		if m.cursor == i {
			arrow = styArrow(">")
			box = "[x]"
			styBox.Bold(true)
			styKey.Bold(true)
		} else {
			arrow = " "
			box = "[ ]"
			styBox.Bold(false)
			styKey.Bold(false)
		}

		body += fmt.Sprintf("%s%s %s\n", arrow, styBox.Render(box), styKey.Render(k))
	}
	// sBody = lipgloss.JoinVertical(lipgloss.Top, body)
	sBody = styBodyContainer.Render(body)

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
