package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
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
			Bold(true)
	styleNormal = lipgloss.NewStyle().
			Foreground(lipgloss.Color(snow)).
			Background(lipgloss.Color(slate))
	styleBody = lipgloss.NewStyle().
			Padding(0, 1, 0, 1)
)

type model struct {
	table  table.Model
	keys   []string
	cursor int
}

func initialModel() model {
	keys := []string{
		"🥕 : carrot",
		"🍎 : apple",
		"🍐 : pear",
	}
	t := makeTable()
	return model{
		table:  t,
		keys:   keys,
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.keys)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	sb := strings.Builder{}

	for i, k := range m.keys {
		if i == m.cursor {
			sb.WriteString(styleSelected.Render(k))
		} else {
			sb.WriteString(styleNormal.Render(k))
		}
		sb.WriteString("\n")
	}
	sb.WriteString(m.table.View())
	return sb.String()
}
