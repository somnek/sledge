package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

var (
	term = termenv.EnvColorProfile()
)

type model struct {
	choices   []string         // item on list
	cursor    int              // which item cursor is pointing at
	selected  map[int]struct{} // which one is selected
	statusBar string
}

func initialModel(firstRun bool) model {
	rdb := connect()
	keys := rdb.getKeys()

	statusBar := "\n"
	if firstRun == false {
		// current time
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		statusBar = fmt.Sprintf("\nlast refresh: \033[1;32m%v\033[0m", currentTime)
	}

	return model{
		choices:   keys,
		selected:  make(map[int]struct{}),
		statusBar: statusBar,
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
		case "r":
			firstRun := false
			return initialModel(firstRun), nil
		}
	}
	return m, nil
}

func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

func (m model) View() string {
	s := "Just choose:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			choice = colorFg(choice, "66")
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = colorFg("x", "79")
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	rdb := connect()
	currentValue := rdb.get(m.choices[m.cursor])

	footer := fmt.Sprintf("\nvalue: \033[1;35m%s\033[0m", currentValue)
	s += footer
	s += m.statusBar
	return s
}
