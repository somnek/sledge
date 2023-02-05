package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else if m.cursor == 0 {
				m.cursor = len(m.items) - 1
			}
			return m, nil
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			} else if m.cursor == len(m.items)-1 {
				m.cursor = 0
			}
			return m, nil
		case "enter", " ":
			// check if cursor already in marked
			if contains(m.marked, m.cursor) {
				m.marked = remove(m.marked, m.cursor)
				return m, nil
			}
			m.marked = append(m.marked, m.cursor)
			return m, nil
		}
	}
	return m, nil
}
