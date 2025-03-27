package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
)

type UiViewState struct {
	focused   int
	viewports []viewport.Model
	ready     bool
}

type Ui struct {
	State *UiViewState
}

func NewUi(viewportCounts int) Ui {
	if viewportCounts > 8 {
		fmt.Println("viewportCounts should be less than 8")
		os.Exit(1)
	}
	return Ui{
		State: &UiViewState{
			focused:   0,
			viewports: make([]viewport.Model, viewportCounts),
			ready:     false,
		},
	}
}
