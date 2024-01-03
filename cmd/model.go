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
	table   table.Model
	records []Record
	cursor  int
	url     string
}

func initialModel(url string) model {
	var err error

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

	// table
	t := recordToTable(records[0])

	return model{
		table:   t,
		records: records,
		cursor:  0,
		url:     url,
	}
}

func (m model) Init() tea.Cmd {
	return doTick()
}
