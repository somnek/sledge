package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	typeStyleMap = map[string]lipgloss.Style{
		"string": styleString,
		"list":   styleList,
		"hash":   styleHash,
	}
)

func (m model) View() string {
	sb := strings.Builder{}

	var typeStyle, keyStyle lipgloss.Style

	for i, record := range m.records {
		typeStyle = typeStyleMap[record.kind]
		if i == m.cursor {
			keyStyle = styleSelected
		} else {
			keyStyle = styleNormal
		}

		rowContent := fmt.Sprintf("%s%s", typeStyle.Render(record.kind), keyStyle.Render(record.key))
		sb.WriteString(rowContent)
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString(m.table.View())
	return sb.String()
}
