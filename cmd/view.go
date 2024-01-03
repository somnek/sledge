package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

var (
	kindStyleMap = map[string]lipgloss.Style{
		"string": styleString,
		"list":   styleList,
		"hash":   styleHash,
		"set":    styleSet,
	}
)

func (m model) View() string {
	title := "Sledge 🛷 - Redis TUI"
	var kindStyle, keyStyle lipgloss.Style
	var bodyView, top, bottom string
	var selected Record

	for i, record := range m.records {
		kindStyle = kindStyleMap[record.kind]
		if i == m.cursor {
			selected = record
			keyStyle = styleSelected
		} else {
			keyStyle = styleNormal
		}

		kindView := kindStyle.Render(record.kind)
		keyView := keyStyle.Render(record.key)
		rowView := fmt.Sprintf("%s%s", kindView, keyView)

		bodyView += rowView + "\n"
	}

	titleView := styleTitle.Render(title)
	top = titleView + "\n" + bodyView

	// values
	switch selected.kind {
	case "string":
		val := selected.val.(string)
		bottom += runewidth.Truncate(val, maxWidth, "…")
	case "hash":
		bottom += m.table.View()
	case "list":
		bottom += m.table.View()
	case "set":
		bottom += m.table.View()
	}

	// fill bottom
	botHeight := countRune(bottom, '\n')
	bottom += strings.Repeat("\n", minBottomHeight-botHeight-1)

	finalView := styleApp.Render(lipgloss.JoinVertical(lipgloss.Left, top, bottom))
	return finalView + "\n"
}
