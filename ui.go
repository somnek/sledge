package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/muesli/termenv"
	"github.com/tidwall/pretty"
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
	db         int
	logs       []string
}

func initialModel(db int, isRefresh bool, curPos int) model {
	rdb := connect(db)
	keys := rdb.getKeys()

	statusBar := "\n"
	if isRefresh {
		// current time
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		statusBar = fmt.Sprintf("\nlast refresh: \033[1;32m%v\033[0m", currentTime)

		return model{
			db:        db,
			choices:   keys,
			cursor:    curPos,
			selected:  make(map[int]struct{}),
			statusBar: statusBar,
		}
	}
	// init (not refresh)
	return model{
		db:        db,
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
		case "ctrl+d":
			// check if 'foo' key exists
			rdb := connect(m.db)
			if rdb.exists("foo") {
				// * add to log here, replace in place
			}
			rdb.add("foo", "bar") // insert no matter what
			// refresh
			return initialModel(m.db, true, 0), nil

		case "x", "d":
			rdb := connect(m.db)

			// deletion methods
			// - delete those that marked wih 'X'
			// - when no item is marked, delete the current one
			// - do not allow delettion if len(key) == 0

			if len(rdb.getKeys()) > 0 {
				if len(m.selected) > 0 {
					for k := range m.selected {
						rdb.del(m.choices[k])
					}
					return initialModel(0, false, 0), nil
				} else { // none selected
					rdb.del(m.choices[m.cursor])
					return initialModel(0, false, 0), nil
				}
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
			return initialModel(m.db, isRefresh, curPos), nil
		case "l", "right", "h", "left":
			// switch db
			if m.db == 0 {
				if Ping(1) == "PONG" {
					return initialModel(1, true, 0), nil
				}
			} else {
				// refresh anyeway
				return initialModel(0, true, 0), nil
			}
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
	s += colorFg("Welcome to red list... 🥥\n", "#f0ead2")
	s += colorFg("─────────────────────────\n", "#E0F2E9")

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = colorFg(">", "#E0F2E9")
			// choice = colorFg(choice, "#DC965A")
			choice = colorFg(choice, "#a98467")
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = colorFg("x", "#ED7B84")
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	rdb := connect(m.db)
	currentValue := "Empty"
	if len(m.choices) > 0 {
		currentValue = rdb.get(m.choices[m.cursor])
	}
	// currentValue = colorFg(currentValue, "#F991CC")
	instruction := colorFg("\nj:down, k:up, d:del, space:mark, d:dummy, r:refresh\n", "#8D8D8D")

	var footer string
	if valType(currentValue) == "map" {
		footer += style.Render("value: " + colorFg(string(pretty.Pretty([]byte(currentValue))), "#adc178"))
	} else {
		currentValue = colorFg(currentValue, "#adc178")
		footer = style.Render(fmt.Sprintf("value: %v", currentValue))
	}

	m.logs = append(m.logs, fmt.Sprintf("\nlogs:\ndb: %d", m.db))
	s += m.statusBar
	s += instruction
	s += m.actionMenu
	s += m.testText
	s += footer
	s += glueLogs(&m.logs)
	return s
}
