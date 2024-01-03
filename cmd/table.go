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
		m := record.val.(map[string]string)

		for k, v := range m {
			rows = append(rows, table.Row{k, v})
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
