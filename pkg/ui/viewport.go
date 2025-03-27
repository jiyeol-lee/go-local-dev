package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

func (u *Ui) GetViewportsLayout() string {
	if len(u.State.viewports) == 2 {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(lipgloss.Top, u.getViewPort(0)),
			lipgloss.JoinHorizontal(lipgloss.Top, u.getViewPort(1)),
		)
	}

	if len(u.State.viewports) == 4 {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(lipgloss.Top, u.getViewPort(0), u.getViewPort(2)),
			lipgloss.JoinHorizontal(lipgloss.Top, u.getViewPort(1), u.getViewPort(3)),
		)
	}

	if len(u.State.viewports) == 6 {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				u.getViewPort(0),
				u.getViewPort(2),
				u.getViewPort(4),
			),
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				u.getViewPort(1),
				u.getViewPort(3),
				u.getViewPort(5),
			),
		)
	}

	if len(u.State.viewports) == 8 {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				u.getViewPort(0),
				u.getViewPort(2),
				u.getViewPort(4),
				u.getViewPort(6),
			),
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				u.getViewPort(1),
				u.getViewPort(3),
				u.getViewPort(5),
				u.getViewPort(7),
			),
		)
	}

	return ""
}

// UpdateViewportsLayout method is used to update the viewports layout when the window size changes
func (u *Ui) UpdateViewportsLayout(windowWidth, windowHeight int) {
	divNum := len(u.State.viewports) / 2

	if !u.State.ready {
		for i := range u.State.viewports {
			vp := viewport.New(windowWidth/divNum-2, windowHeight/2-2)
			if i%2 == 0 {
				vp.YPosition = windowHeight
			}
			u.State.viewports[i] = vp
		}
		u.State.ready = true
	} else {
		for i := range u.State.viewports {
			u.State.viewports[i].Width = windowWidth/divNum - 2
			u.State.viewports[i].Height = windowHeight/2 - 2
		}
	}
}

func (u *Ui) getViewPort(index int) string {
	vp := u.State.viewports[index]

	v := vp.View()
	s := lipgloss.NewStyle()
	if u.State.focused == index {
		s = s.Foreground(lipgloss.Color("69"))
	}
	lt := s.Render(fmt.Sprintf("[%d]%s", index+1, strings.Repeat("─", vp.Width-3)))
	lb := s.Render((strings.Repeat("─", vp.Width)))
	ll := s.Render(("┌\n" + strings.Repeat("│\n", vp.Height) + "└"))
	lr := s.Render(("┐\n" + strings.Repeat("│\n", vp.Height) + "┘"))

	return lipgloss.JoinHorizontal(lipgloss.Top,
		ll,
		lipgloss.JoinVertical(lipgloss.Top, lt, v, lb),
		lr,
	)
}
