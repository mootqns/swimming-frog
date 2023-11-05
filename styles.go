package main

import "github.com/charmbracelet/lipgloss"

var baseStyle = lipgloss.NewStyle().
	Width(2).
	Height(1)

var frogStyle = lipgloss.NewStyle().
	Width(2).
	Height(1).
	Background(lipgloss.Color("#14db8f"))

var pelletStyle = lipgloss.NewStyle().
	Width(2).
	Height(1).
	Bold(true).
	Background(lipgloss.Color("#db1481")).
	AlignHorizontal(lipgloss.Center)

var boardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#bf00ff"))

var gameOverStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a32a2a"))

var scoreStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#14db8f"))

var startScreenStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Margin(3).
	Padding(2).
	Width(60)

var startBorder = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("83")).
	Align(lipgloss.Center)
					
var gameOverText = "game over"
