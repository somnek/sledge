package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func makeTable() table.Model {

	rows := []table.Row{
		{"🌊 ", "water"},
		{"🔥 ", "fire"},
		{"🌱 ", "grass"},
	}

	columns := []table.Column{
		{Title: "⭐️", Width: 10},
		{Title: "name", Width: 10},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	// style
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	// s.Selected = s.Cell // treat selected row as normal

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("81")).
		Bold(false)

	t.SetStyles(s)
	return t
}
