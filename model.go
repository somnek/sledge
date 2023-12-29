package main

import (
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Record struct {
	key  string
	val  interface{}
	kind string
}
type model struct {
	table   table.Model
	records []Record
	cursor  int
}

func initialModel() model {
	var err error

	// redis
	rdb, err := NewClient("localhost:6379", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	t := makeTable()
	records, err := rdb.GetRecords(ctx, "*")
	if err != nil {
		log.Fatal(err)
	}

	return model{
		table:   t,
		records: records,
		cursor:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
