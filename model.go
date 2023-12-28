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
	types, err := rdb.GetRecords(ctx, "*")
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range types {
		if t.kind == "hash" {
			for k, v := range t.val.(map[string]string) {
				log.Printf("key: %s, val: %s\n", k, v)
			}
		} else {
			log.Printf("key: %s, val: %s\n", t.key, t.val)
		}
	}

	return model{
		table:   t,
		records: []Record{},
		cursor:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
