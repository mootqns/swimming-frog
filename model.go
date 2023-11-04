package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

const (
	BoardWidth  = 40
	BoardHeight = 20
)

type TickMsg time.Time

type FrogGame struct {
	Frog     Coord
	GameOver bool
}

type Coord struct {
	X int
	Y int
}

func NewFrogGame() FrogGame {
	frog := Coord{
		X: 10, // Initial X position
		Y: 10, // Initial Y position
	}

	return FrogGame{
		Frog:     frog,
		GameOver: false,
	}
}

func tickEvery() tea.Cmd {
	return tea.Every(time.Millisecond*100, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (f FrogGame) Init() tea.Cmd {
	return tickEvery()
}

func (f FrogGame) View() string {
	if f.GameOver {
		return "Game Over!\nPress 'q' to quit."
	}

	screen := ""

	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
			if i == f.Frog.Y && j == f.Frog.X {
				screen += "🐸"
			} else {
				screen += " "
			}
		}

		if i != BoardHeight-1 {
			screen += "\n"
		}
	}

	helpMsg := "Use arrow keys to move. Press 'q' to quit."

	return screen + helpMsg
}

func (f FrogGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "q":
			return f, tea.Quit
		case "up":
			if f.Frog.Y > 0 {
				f.Frog.Y--
			}
		case "right":
			if f.Frog.X < BoardWidth-1 {
				f.Frog.X++
			}
		case "down":
			if f.Frog.Y < BoardHeight-1 {
				f.Frog.Y++
			}
		case "left":
			if f.Frog.X > 0 {
				f.Frog.X--
			}
		}
	case TickMsg:
		// Game logic can be added here if needed.
	}

	return f, nil
}