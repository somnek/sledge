package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	kindStyleMapNormal = map[string]lipgloss.Style{
		"string": styleString,
		"hash":   styleHash,
		"set":    styleSet,
		"list":   styleList,
	}
	kindStyleMapDarker = map[string]lipgloss.Style{
		"string": styleString.Copy().Background(lipgloss.Color(palette8Darker)),
		"hash":   styleHash.Copy().Background(lipgloss.Color(palette7Darker)),
		"set":    styleSet.Copy().Background(lipgloss.Color(palette4Darker)),
		"list":   styleList.Copy().Background(lipgloss.Color(palette6Darker)),
	}
)

func BuildBody(records []Record, cursor int) string {
	var kindStyle, keyStyle lipgloss.Style
	var bodyView string

	for i, record := range records {
		if i == cursor {
			keyStyle = styleSelected
			kindStyle = kindStyleMapDarker[record.kind]
		} else {
			keyStyle = styleNormal
			kindStyle = kindStyleMapNormal[record.kind]
		}

		kindView := kindStyle.Render(record.kind)
		keyView := keyStyle.Render(record.key)
		rowView := fmt.Sprintf("%s%s", kindView, keyView)

		bodyView += rowView + "\n"
	}

	return bodyView
}

func (m model) View() string {
	title := "Sledge 🛷 - Redis TUI"
	var top, bottom string

	// title + body
	titleView := styleTitle.Render(title)
	top = titleView + "\n" + m.body

	// values
	switch m.selected.kind {
	case "string":
		var sText string
		header := styleStringHeader.Render("Value") + "\n"
		underline := styleUnderline.Render(strings.Repeat("─", 50)) + "\n"

		val := m.selected.val.(string)
		wrap := wordwrap.String(val, maxWidth)
		splits := strings.Split(wrap, "\n")

		if len(splits) > fixedBottomHeight {
			// replace last line with ellipsis
			splits = splits[:fixedBottomHeight-3]
			sText += strings.Join(splits, "\n")
			sText += "\n" + strings.Repeat(" ", 20) + "..."
		} else {
			sText += strings.Join(splits, "\n")
		}
		bottom += header
		bottom += underline
		bottom += styleStringVal.Render(sText)

	case "hash", "list", "set":
		bottom += m.table.View()
	}

	// fill bottom
	botHeight := countRune(bottom, '\n')
	bottom += strings.Repeat("\n", fixedBottomHeight-botHeight-1)

	finalView := styleApp.Render(lipgloss.JoinVertical(lipgloss.Left, top, bottom))
	return finalView + "\n"
}
