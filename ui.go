package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // item on list
	cursor   int              // which item cursor is pointing at
	selected map[int]struct{} // which one is selected
}

func initialModel() model {
	rdb := connect()
	keys := rdb.getKeys()
	return model{
		choices:  keys,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

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
				m.cursor = len(m.choices) - 1
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else if m.cursor == len(m.choices)-1 {
				m.cursor = 0
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "r": // refresh redis
			return initialModel(), nil
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Just choose:\n\n"

	// s := fmt.Sprintf("Selected: %s\n\n", m.selected)
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	rdb := connect()
	currentValue := rdb.get(m.choices[m.cursor])

	// display text in pink color
	s += fmt.Sprintf("\nvalue: \033[1;35m%s\033[0m", currentValue)
	return s
}
