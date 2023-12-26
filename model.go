package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	keys   []string
	cursor int
}

func initialModel() model {
	return model{
		keys:   []string{},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "test render"
}
