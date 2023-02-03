package main

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ctx := context.Background()
	dummyIns()
	Ping(ctx)

	p := tea.NewProgram(initModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

}
