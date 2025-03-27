package main

import (
	"fmt"
	"jiyeol-lee/go-local-dev/pkg/ui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	u := ui.NewUi(4)

	p := tea.NewProgram(
		&u,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
