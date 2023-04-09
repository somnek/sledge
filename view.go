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
	styBox           = lipgloss.NewStyle().Foreground(melon)
	styKey           = lipgloss.NewStyle().Foreground(yellow)
	styValue         = lipgloss.NewStyle().Foreground(darkGray)
	styArrow         = lipgloss.NewStyle().Foreground(red).Render
)

func (m model) View() string {
	var sBody string
	var body string

	// -----> header
	sTitle := styTitle.Render("Welcome to redlist... 🥥")
	sUnder := styUnderline.Render("─────────────────────────")

	// -----> body
	for i, k := range m.items {
		var arrow string // " " or >
		var box = "[ ]"  //[ ] or [x]

		// on cursor
		if m.cursor == i {
			arrow = styArrow(">")
			styBox.Bold(true)
			styKey.Bold(true)
		} else {
			arrow = " "
			styBox.Bold(false)
			styKey.Bold(false)
		}

		// on marked
		for _, n := range m.marked {
			if n == i {
				box = "[x]"
			}
		}

		body += fmt.Sprintf("%s%s %s\n", arrow, styBox.Render(box), styKey.Render(k))
	}
	sBody = styBodyContainer.Render(body)

	// -----> footer
	ctx := context.Background()

	value := get(ctx, m.items[m.cursor])
	sValue := styValue.Render(value)

	return fmt.Sprintf(
		"\n%s\n%s\n%s\n%s\n",
		sTitle,
		sUnder,
		sBody,
		sValue,
	)
}
