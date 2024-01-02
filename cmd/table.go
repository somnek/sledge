package cmd

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// recordToTable converts slices of Records to a table.Model.
func recordToTable(record Record) table.Model {

	if records
	rows := []table.Row{
		{"Japan", "37,274,000"},
		{"India", "32,065,760"},
		{"China", "28,516,904"},
		{"Bangladesh", "22,478,116"},
		{"Brazil", "22,429,800"},
		{"Mexico", "22,085,140"},
		{"Egypt", "21,750,020"},
		{"China", "21,333,332"},
		{"India", "20,961,472"},
		{"Japan", "19,059,856"},
		{"China", "16,874,740"},
	}

	columns := []table.Column{
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	t := makeTable(columns, rows)
	return t
}

func makeTable(columns []table.Column, rows []table.Row) table.Model {

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

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("81")).
		Bold(false)

	t.SetStyles(s)
	return t
}
