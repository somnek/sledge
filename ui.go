package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/muesli/termenv"
)

var (
	term  = termenv.EnvColorProfile()
	style = lipgloss.NewStyle().
		Width(50).
		Border(lipgloss.RoundedBorder())
)

type model struct {
	choices    []string         // item on list
	cursor     int              // which item cursor is pointing at
	selected   map[int]struct{} // which one is selected
	statusBar  string
	actionMenu string
	testText   string
}

func initialModel(isRefresh bool, curPos int) model {
	rdb := connect()
	keys := rdb.getKeys()

	statusBar := "\n"
	if isRefresh {
		// current time
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		statusBar = fmt.Sprintf("\nlast refresh: \033[1;32m%v\033[0m", currentTime)

		return model{
			choices:   keys,
			cursor:    curPos,
			selected:  make(map[int]struct{}),
			statusBar: statusBar,
		}
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
		case "ctrl+c", "q", "esc":
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
		case "d", "x":
			rdb := connect()

			// there are 2 deletion method
			// 1. delete those that marked wih 'X'
			// 2. when no item is marked, delete the current one

			if len(m.selected) > 0 {
				for k := range m.selected {
					rdb.del(m.choices[k])
				}
				return initialModel(false, 0), nil
			} else { // none selected
				rdb.del(m.choices[m.cursor])
				return initialModel(false, 0), nil
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "r":
			curPos := m.cursor
			isRefresh := true
			return initialModel(isRefresh, curPos), nil
		case "l", "right":
			// switch db
		}

	}
	return m, nil
}

func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

func (m model) View() string {
	// colorFg will return string
	s := "\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = colorFg(">", "#E0F2E9")
			choice = colorFg(choice, "#DC965A")
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = colorFg("x", "#ED7B84")
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	rdb := connect()
	currentValue := "Empty"
	if len(m.choices) > 0 {
		currentValue = rdb.get(m.choices[m.cursor])
	}
	currentValue = colorFg(currentValue, "#F991CC")
	instruction := colorFg("\nj:down, k:up, d:del, r:refresh\n", "#8D8D8D")

	footer := style.Render(fmt.Sprintf("value: %s", currentValue))
	s += footer
	s += m.statusBar
	s += instruction
	s += m.actionMenu
	s += m.testText
	return s
}
