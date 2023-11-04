package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	frog string
}

func initialModel() model {
	return model{
		frog: "frog",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := m.frog
	s += "\nPress q to quit.\n"
	return s
}
