package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	size := len(m.records)

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

		case "G":
			m.cursor = len(m.records) - 1
			m.vpCur = fixedBodyHeight - 1
			vpRec := m.records[size-fixedBodyHeight:]
			m.body = BuildBody(vpRec, m.vpCur)

			m.updateSelected(&m.selected)
			m.table = recordToTable(m.selected)

		case "g":
			m.cursor = 0
			m.vpCur = 0
			vpRec := m.records[:fixedBodyHeight]
			m.body = BuildBody(vpRec, m.vpCur)

			m.updateSelected(&m.selected)
			m.table = recordToTable(m.selected)

		case "j", "down":
			if m.cursor < len(m.records)-1 {
				m.cursor++
			}
			if m.vpCur < fixedBodyHeight-1 {
				m.vpCur++
			}

			vpRec := make([]Record, fixedBodyHeight)
			var offsetL, offsetR int
			if m.vpCur == fixedBodyHeight {
				offsetL, offsetR = m.cursor, m.cursor+fixedBodyHeight
			} else {
				offsetL, offsetR = m.cursor-m.vpCur, m.cursor+(fixedBodyHeight-m.vpCur)
			}
			vpRec = m.records[offsetL:offsetR]

			m.body = BuildBody(vpRec, m.vpCur)
			m.updateSelected(&m.selected)
			m.table = recordToTable(m.selected)

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
			if m.vpCur > 0 {
				m.vpCur--
			}

			vpRec := make([]Record, fixedBodyHeight)
			var offsetL, offsetR int
			if m.vpCur == 0 {
				offsetL, offsetR = m.cursor, m.cursor+fixedBodyHeight
			} else {
				offsetL, offsetR = m.cursor-m.vpCur, m.cursor+(fixedBodyHeight-m.vpCur)
			}
			vpRec = m.records[offsetL:offsetR]

			m.body = BuildBody(vpRec, m.vpCur)
			m.updateSelected(&m.selected)
			m.table = recordToTable(m.selected)
		}
	}

	return m, cmd
}

// grab the value of current selected key and set m.selected.val
func (m model) updateSelected(selected *Record) {

	// connect
	rdb, err := NewClient(m.url)
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	newS := m.records[m.cursor]
	newS.val, err = rdb.ExtractVal(ctx, newS.key, newS.kind)
	if err != nil {
		log.Fatal(err)
	}

	*selected = newS
}
