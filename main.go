package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// init redis data
	dummyIns()

	p := tea.NewProgram(initialModel(false, 0))
	if err := p.Start(); err != nil {
		fmt.Printf("skull :%v", err)
		os.Exit(1)
	}
}
