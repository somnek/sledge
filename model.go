package main

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items  []string
	cursor int
	marked []int
	value  string
	// selected map[int]struct{} // {2: 'ball'}
}

func initModel(ctx context.Context) model {
	items := keys(ctx)
	return model{items: items}
}

func (m model) Init() tea.Cmd {
	return nil
}
