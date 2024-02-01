package cmd

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg struct {
	Time time.Time
}

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg{Time: t}
	})
}

type Record struct {
	key  string
	val  interface{}
	kind string
}

type model struct {
	table    table.Model
	records  []Record
	body     string
	vpCur    int
	selected Record
	keys     string
	cursor   int
	url      string
}

func initialModel(url string) model {
	var (
		err    error
		cursor int
		body   string
	)

	// connect
	rdb, err := NewClient(url)
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	// records
	records, err := rdb.GetRecords(ctx, "*")
	if err != nil {
		log.Fatal(err)
	}

	// body
	if len(records) < fixedBodyHeight {
		currentBodyHeight = len(records)
	}
	vpRec := records[cursor:currentBodyHeight]
	body = BuildBody(vpRec, cursor)

	// get val for selected
	selected := records[cursor]
	selected.val, err = rdb.ExtractVal(ctx, selected.key, selected.kind)
	if err != nil {
		log.Fatal(err)
	}

	// table
	t := recordToTable(selected)

	return model{
		table:    t,
		records:  records,
		body:     body,
		selected: selected,
		cursor:   cursor,
		url:      url,
	}
}

func (m model) Init() tea.Cmd {
	return doTick()
}
