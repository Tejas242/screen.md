package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"screen.md/internal/tui"
)

func main() {
    p := tea.NewProgram(
        tui.NewModel(),
        tea.WithAltScreen(),
        tea.WithMouseCellMotion(),
    )

    if err := p.Start(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
