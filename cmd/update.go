package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// tick events
	case TickMsg:
		// connect
		rdb, err := NewClient(m.url)
		if err != nil {
			log.Fatal(err)
		}
		defer rdb.Close()

		// records
		m.records, err = rdb.GetRecords(ctx, "*")
		if err != nil {
			log.Fatal(err)
		}

		return m, doTick()

	// key events
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
			m.table = recordToTable(m.records[m.cursor])

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.records) - 1
			}
			m.table = recordToTable(m.records[m.cursor])
		}
	}

	return m, nil
}
