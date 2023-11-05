package main

import "github.com/charmbracelet/lipgloss"

var baseStyle = lipgloss.NewStyle().
	Width(2).
	Height(1).
	Background(lipgloss.Color("#458588"))

var logStyle = lipgloss.NewStyle().
	Width(2).
	Height(1).
	Background(lipgloss.Color("#996633"))

var boardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#d79921")).
	BorderBackground(lipgloss.Color("#d79921"))

var scoreStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#98971a"))

var menuScreenStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#79740e")).
	Background(lipgloss.Color("#98971a")).
	Margin(1).
	Padding(1).
	Width(50)

var startBorder = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#d79921")).
	Align(lipgloss.Center)

var menuTextStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#fbf1c7"))

var altTextStyle = lipgloss.NewStyle().
	Italic(true).
	Foreground(lipgloss.Color("#fbf1c7")).
	MarginBottom(2)

var italicsTextStyle = lipgloss.NewStyle().
	Italic(true).
	Foreground(lipgloss.Color("#fbf1c7"))
