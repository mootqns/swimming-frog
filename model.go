package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type TickMsg time.Time

type coord struct {
	x int
	y int
}

type frogGame struct {
	frog     coord
	gameOver bool
}

func newFrogGame() frogGame {
	frog := coord{
		x: 10, // Initial X position
		y: 10, // Initial Y position
	}

	return frogGame{
		frog:     frog,
		gameOver: false,
	}
}

func tickEvery() tea.Cmd {
	return tea.Every(time.Millisecond*100, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (f frogGame) Init() tea.Cmd {
	return tickEvery()
}

func (f frogGame) View() string {
	if f.gameOver {
		return "Game Over!\nPress 'q' to quit."
	}

	screen := ""

	for i := 0; i < BOARD_HEIGHT; i++ {
		for j := 0; j < BOARD_WIDTH; j++ {
			if i == f.frog.y && j == f.frog.x {
				screen += "🐸"
			} else {
				screen += " "
			}
		}

		if i != BOARD_HEIGHT-1 {
			screen += "\n"
		}
	}

	helpMsg := "Use arrow keys to move. Press 'q' to quit."

	return boardStyle.Render(screen) + screen + helpMsg
}

func (f frogGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "q":
			return f, tea.Quit
		case "up":
			if f.frog.y > 0 {
				f.frog.y--
			}
		case "right":
			if f.frog.x < BOARD_WIDTH-1 {
				f.frog.x++
			}
		case "down":
			if f.frog.y < BOARD_HEIGHT-1 {
				f.frog.y++
			}
		case "left":
			if f.frog.x > 0 {
				f.frog.x--
			}
		}
	case TickMsg:
		// Game logic can be added here if needed.
	}

	return f, nil
}

const (
	BOARD_WIDTH  = 44
	BOARD_HEIGHT = 20
)
