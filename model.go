package main

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items  []string
	cursor int
	// selected map[int]struct{} // {2: 'ball'}
	status string
}

func initModel(ctx context.Context) model {
	items := keys(ctx)
	return model{items: items}
}

func (m model) Init() tea.Cmd {
	return nil
}
