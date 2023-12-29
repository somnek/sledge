package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	kindStyleMap = map[string]lipgloss.Style{
		"string": styleString,
		"list":   styleList,
		"hash":   styleHash,
	}
)

func (m model) View() string {
	title := "Sledge 🛷 - Redis TUI"
	sb := strings.Builder{}

	var kindStyle, keyStyle lipgloss.Style
	var bodyContent string

	for i, record := range m.records {
		kindStyle = kindStyleMap[record.kind]
		if i == m.cursor {
			keyStyle = styleSelected
		} else {
			keyStyle = styleNormal
		}

		kindContent := kindStyle.Render(record.kind)
		keyContent := keyStyle.Render(record.key)
		rowContent := fmt.Sprintf("%s%s", kindContent, keyContent)

		bodyContent += rowContent + "\n"
	}

	titleContent := styleTitle.Render(title)
	sb.WriteString(styleBody.Render(titleContent + "\n" + bodyContent + "\n" + m.table.View()))
	return sb.String()
}
