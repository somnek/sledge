package cmd

import (
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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

		// update vp items, reflect update from external actions
		// if deleted item is under cursor, assign cursor to the last item
		if m.cursor >= len(m.records) {
			m.cursor = len(m.records) - 1
			m.vpCur = len(m.records) - 1
			vpRec := m.records
			currentBodyHeight = len(m.records)
			m.body, m.table, m.selected = m.updateVP(vpRec)
		}

		// handle item less than minHeight (new item etc)
		if len(m.records) <= fixedBodyHeight {
			vpRec := m.records
			currentBodyHeight = len(m.records)
			m.body, m.table, m.selected = m.updateVP(vpRec)
		}

		return m, doTick()

	// key events
	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "G":
			m.cursor, m.vpCur = len(m.records)-1, fixedBodyHeight-1

			// handle item less than minHeight
			if len(m.records) < fixedBodyHeight {
				currentBodyHeight = len(m.records)
			}

			vpRec := m.records[len(m.records)-currentBodyHeight:]
			m.body, m.table, m.selected = m.updateVP(vpRec)

		case "g":
			m.cursor, m.vpCur = 0, 0

			// handle item less than minHeight
			if len(m.records) < fixedBodyHeight {
				currentBodyHeight = len(m.records)
			}
			vpRec := m.records[:currentBodyHeight]
			m.body, m.table, m.selected = m.updateVP(vpRec)

		case "j", "down":
			if m.cursor < len(m.records)-1 {
				m.cursor++
			}
			if m.vpCur < currentBodyHeight-1 {
				m.vpCur++
			}

			var vpRec []Record
			// handle item less than minHeight
			if len(m.records) <= fixedBodyHeight {
				vpRec = m.records
			} else {
				offsetL, offsetR := calculateOffsets(m.cursor, m.vpCur)
				vpRec = m.records[offsetL:offsetR]
			}
			m.body, m.table, m.selected = m.updateVP(vpRec)

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
			if m.vpCur > 0 {
				m.vpCur--
			}

			var vpRec []Record
			// handle item less than minHeight
			if len(m.records) <= fixedBodyHeight {
				vpRec = m.records
			} else {
				offsetL, offsetR := calculateOffsets(m.cursor, m.vpCur)
				vpRec = m.records[offsetL:offsetR]
			}
			m.body, m.table, m.selected = m.updateVP(vpRec)
		}
	}

	return m, cmd
}

// grab the value of current selected key and set m.selected.val
func (m model) updateSelected() Record {

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

	return newS
}

func (m model) updateVP(recs []Record) (string, table.Model, Record) {
	m.body = BuildBody(recs, m.vpCur)
	m.selected = m.updateSelected()
	m.table = recordToTable(m.selected)
	return m.body, m.table, m.selected
}

func calculateOffsets(cur, vpCur int) (int, int) {
	var offsetL, offsetR int
	if vpCur == currentBodyHeight || vpCur == 0 {
		offsetL, offsetR = cur, cur+currentBodyHeight
	} else {
		offsetL, offsetR = cur-vpCur, cur+(currentBodyHeight-vpCur)
	}
	return offsetL, offsetR
}
