package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	frog string
}

func initialModel() model {
	return model{
		frog: "🐸",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	s := m.frog
	return s
}
