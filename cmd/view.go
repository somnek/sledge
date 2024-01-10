package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
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

	// title
	titleView := styleTitle.Render(title)

	for i, record := range m.records {
		kindStyle = kindStyleMap[record.kind]
		if i == m.cursor {
			keyStyle = styleSelected
		} else {
			keyStyle = styleNormal
		}

		kindView := kindStyle.Render(record.kind)
		keyView := keyStyle.Render(record.key)
		rowView := fmt.Sprintf("%s%s", kindView, keyView)

		bodyView += rowView + "\n"
	}

	top = titleView + "\n" + bodyView

	// values
	switch m.selected.kind {
	case "string":
		val := m.selected.val.(string)
		wrap := wordwrap.String(val, maxWidth)
		splits := strings.Split(wrap, "\n")

		if len(splits) > fixedBottomHeight {
			// replace last line with ellipsis
			splits = splits[:fixedBottomHeight-1]
			bottom += strings.Join(splits, "\n")
			bottom += "\n" + strings.Repeat(" ", 20) + "..."
		} else {
			bottom += strings.Join(splits, "\n")
		}

	case "hash":
		bottom += m.table.View()

	case "list":
		bottom += m.table.View()

	case "set":
		bottom += m.table.View()
	}

	// fill bottom
	botHeight := countRune(bottom, '\n')
	bottom += strings.Repeat("\n", fixedBottomHeight-botHeight-1)

	finalView := styleApp.Render(lipgloss.JoinVertical(lipgloss.Left, top, bottom))
	return finalView + "\n"
}
