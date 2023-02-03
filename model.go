package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	keys     []string
	cursor   int
	selected map[int]struct{}
	status   string
}

func initModel() model {
	keys := []string{"foo", "bar"}
	return model{keys: keys}
}

func (m model) Init() tea.Cmd {
	return nil
}
