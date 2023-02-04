package main

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ctx := context.Background()
	dummyIns(ctx)
	Ping(ctx)

	p := tea.NewProgram(initModel(ctx))
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

}
