package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// example usaage: sledge redis://<user>:<pass>@localhost:6379/<db>

var rootCmd = &cobra.Command{
	Use:   "sledge",
	Short: "\n🛷 Sledge is a Redis TUI for data viewing.\n",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// debug
		if len(os.Getenv("DEBUG")) > 0 {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				fmt.Println("fatal", err)
				os.Exit(1)
			}
			defer f.Close()
		}

		// run
		url := args[0]
		p := tea.NewProgram(initialModel(url))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
