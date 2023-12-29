package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "j", "down":
			if m.cursor < len(m.records)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.records) - 1
			}
		}
	}

	return m, nil
}
