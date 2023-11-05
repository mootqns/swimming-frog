package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	BOARD_WIDTH  = 44
	BOARD_HEIGHT = 20
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

const (
	BLANK_CELL = iota
)

type TickMsg time.Time

type coord struct {
	x int
	y int
}

// this is the model used by bubbletea
type frogGame struct {
	gameBoard [][]int
	frog      coord
	score     int
	gameOver  bool

	width  int
	height int
}

func newFrogGame() frogGame {
	frog := coord{x: (BOARD_WIDTH / 2) + 1, y: BOARD_HEIGHT / 2}

	game := frogGame{
		frog:     frog,
		gameOver: false,
	}

	game.updateBoard()

	return game
}

func (f *frogGame) updateBoard() {
	gameBoard := [][]int{}

	for i := 0; i < BOARD_HEIGHT; i++ {
		row := []int{}
		for j := 0; j < BOARD_WIDTH; j++ {
			cellType := BLANK_CELL

			row = append(row, cellType)
		}

		gameBoard = append(gameBoard, row)
	}

	f.gameBoard = gameBoard
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
	scoreLabel := scoreStyle.Render("score")
	scoreText := fmt.Sprintf("\n%s: %d\n\n", scoreLabel, f.score)

	if f.gameOver {
		return gameOverStyle.Render(gameOverText) + scoreText + "ctrl+c to quit\n"
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

	helpMsg := "arrow keys to move\ncontrol + c\n"

	return boardStyle.Render(screen) + scoreText + helpMsg
}

func (f frogGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		f.width = msg.Width
		f.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return f, tea.Quit
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
	}
	return f, nil
}
