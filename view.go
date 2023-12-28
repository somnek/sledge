package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	lava  = lipgloss.Color("#F56E0F")
	void  = lipgloss.Color("#151419")
	dust  = lipgloss.Color("#878787")
	slate = lipgloss.Color("#262626")
	snow  = lipgloss.Color("#FBFBFB")

	styleSelected = lipgloss.NewStyle().
			Foreground(lipgloss.Color(slate)).
			Background(lipgloss.Color(lava)).
			Bold(true).
			Width(24)
	styleNormal = lipgloss.NewStyle().
			Foreground(lipgloss.Color(snow)).
			Background(lipgloss.Color(slate)).
			Width(24)
	styleBody = lipgloss.NewStyle().
			Padding(0, 1, 0, 1)
)

func (m model) View() string {
	sb := strings.Builder{}

	for i, k := range m.records {
		if i == m.cursor {
			sb.WriteString(styleSelected.Render(k.key))
		} else {
			sb.WriteString(styleNormal.Render(k.key))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString(m.table.View())
	return sb.String()
}
