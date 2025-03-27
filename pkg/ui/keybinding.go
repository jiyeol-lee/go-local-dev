package ui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (u *Ui) Keybinding(message string) tea.Cmd {
	switch message {
	case "ctrl+c", "q", "esc":
		return tea.Quit

	case "tab":
		u.State.focused = (u.State.focused + 1) % len(u.State.viewports)

		break

	case "1", "2", "3", "4", "5", "6", "7", "8":
		num, err := strconv.Atoi(message)
		if err == nil {
			if num-1 >= 0 && num-1 < len(u.State.viewports) {
				u.State.focused = num - 1
			}
		}

		break
	}

	return tea.Batch()
}
