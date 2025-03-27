package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (u *Ui) Init() tea.Cmd {
	return nil
}

func (u *Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return u, u.Keybinding(msg.String())

	case tea.WindowSizeMsg:
		u.UpdateViewportsLayout(msg.Width, msg.Height)

		break
	}

	// Only acvite the focused viewport
	u.State.viewports[u.State.focused], cmd = u.State.viewports[u.State.focused].Update(msg)
	cmds = append(cmds, cmd)

	return u, tea.Batch(cmds...)
}

func (u *Ui) View() string {
	if !u.State.ready {
		return ""
	}

	return u.GetViewportsLayout()
}
