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

		return m, doTick()

	// key events
	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "G":
			m.cursor, m.vpCur = len(m.records)-1, fixedBodyHeight-1
			vpRec := m.records[len(m.records)-fixedBodyHeight:]
			m.body, m.table = m.updateVP(vpRec)

		case "g":
			m.cursor, m.vpCur = 0, 0

			vpRec := m.records[:fixedBodyHeight]
			m.body, m.table = m.updateVP(vpRec)

		case "j", "down":
			if m.cursor < len(m.records)-1 {
				m.cursor++
			}
			if m.vpCur < fixedBodyHeight-1 {
				m.vpCur++
			}

			offsetL, offsetR := calculateOffsets(m.cursor, m.vpCur)
			vpRec := m.records[offsetL:offsetR]
			m.body, m.table = m.updateVP(vpRec)

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
			if m.vpCur > 0 {
				m.vpCur--
			}

			offsetL, offsetR := calculateOffsets(m.cursor, m.vpCur)
			vpRec := m.records[offsetL:offsetR]
			m.body, m.table = m.updateVP(vpRec)
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

func (m model) updateVP(recs []Record) (string, table.Model) {
	m.body = BuildBody(recs, m.vpCur)
	m.updateSelected(&m.selected)
	m.table = recordToTable(m.selected)
	return m.body, m.table
}

func calculateOffsets(cur, vpCur int) (int, int) {
	var offsetL, offsetR int
	if vpCur == fixedBodyHeight || vpCur == 0 {
		offsetL, offsetR = cur, cur+fixedBodyHeight
	} else {
		offsetL, offsetR = cur-vpCur, cur+(fixedBodyHeight-vpCur)
	}
	return offsetL, offsetR
}
