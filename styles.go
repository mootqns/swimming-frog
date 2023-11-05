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
	BorderForeground(lipgloss.Color("#d79921"))

var scoreStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#98971a"))

var menuScreenStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#79740e")).
	Background(lipgloss.Color("#98971a")).
	Margin(3).
	Padding(2).
	Width(60)

var startBorder = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#d79921")).
	Align(lipgloss.Center)

var menuTextStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#fbf1c7"))
