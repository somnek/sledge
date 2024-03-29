package cmd

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// recordToTable converts slices of Records to a table.Model.
func recordToTable(record Record) table.Model {
	var rows []table.Row
	var columns []table.Column

	switch record.kind {
	case "hash":
		pairs := record.val.([]FVPair)

		for _, p := range pairs {
			rows = append(rows, table.Row{p.field, p.value})
		}

		columns = []table.Column{
			{Title: "Field", Width: 13},
			{Title: "Value", Width: 33},
		}

	case "list":
		vals := record.val.([]string)

		for _, v := range vals {
			rows = append(rows, table.Row{v})
		}
		columns = []table.Column{
			{Title: "Values", Width: maxWidth},
		}

	case "set":
		vals := record.val.([]string)

		for _, v := range vals {
			rows = append(rows, table.Row{v})
		}
		columns = []table.Column{
			{Title: "Members", Width: maxWidth},
		}
	}

	t := makeTable(columns, rows)
	return t
}

func makeTable(columns []table.Column, rows []table.Row) table.Model {

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(6),
	)

	// style
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color("81"))

	s.Selected = s.Selected.UnsetForeground().
		UnsetBold()

	t.SetStyles(s)
	return t
}
