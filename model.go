package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	BOARD_WIDTH  = 70
	BOARD_HEIGHT = 30
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

const (
	BLANK_CELL = iota
	LOG_CELL
)

type TickMsg time.Time

type coord struct {
	x int
	y int
}

type wood struct {
	body      []coord
	direction int
}

// this is the model used by bubbletea
type frogGame struct {
	gameBoard   [][]int
	frog        coord
	score       int
	gameOver    bool
	startScreen bool
	testLog     wood

	width  int
	height int
}

func newFrogGame() frogGame {
	frog := coord{x: (BOARD_WIDTH / 2) + 1, y: BOARD_HEIGHT / 2}

	testLog := wood{
		body: []coord{
			{x: 1, y: 4},
			{x: 2, y: 4},
			{x: 3, y: 4},
		},

		direction: RIGHT,
	}

	game := frogGame{
		testLog:     testLog,
		frog:        frog,
		startScreen: true,
		gameOver:    false,
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

			curCell := coord{j, i}

			if f.testLog.coordInBody(curCell) {
				cellType = LOG_CELL
			}

			row = append(row, cellType)
		}

		gameBoard = append(gameBoard, row)
	}

	f.gameBoard = gameBoard
}

func (w wood) coordInBody(c coord) bool {
	for _, woodPart := range w.body {
		if woodPart == c {
			return true
		}
	}

	return false
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
	if f.width == 0 {
		return "loading"
	}
	scoreLabel := scoreStyle.Render("score")
	scoreText := fmt.Sprintf("\n%s: %d\n\n", scoreLabel, f.score)

	if f.startScreen {
		return lipgloss.Place(f.width, f.height, lipgloss.Center, lipgloss.Center,
			startBorder.Render(startScreenStyle.Render("> Frog Game")+
				"\n\nenter to play"))
	}

	if f.gameOver {
		return gameOverStyle.Render(gameOverText) + scoreText + "ctrl+c/q to quit\n"
	}

	screen := ""

	for i := 0; i < BOARD_HEIGHT; i++ {
		for j := 0; j < BOARD_WIDTH; j++ {
			if i == f.frog.y && j == f.frog.x {
				screen += "🐸"
			} else if f.gameBoard[i][j] == LOG_CELL {
				screen += logStyle.Render(" ")
			} else {
				screen += " "
			}
		}

		if i != BOARD_HEIGHT-1 {
			screen += "\n"
		}
	}

	helpMsg := "arrows to move | ctrl+c/q to quit\n"

	return lipgloss.Place(f.width, f.height, lipgloss.Center, lipgloss.Center, boardStyle.Render(screen)+scoreText+helpMsg)
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
		case "enter":
			f.startScreen = false
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

		prevWoodPartPos := f.testLog.body[0]

		switch f.testLog.direction {
		case RIGHT:
			f.testLog.body[0].x += 1
		case LEFT:
			f.testLog.body[0].x -= 1
		}

		for i := 1; i < len(f.testLog.body); i++ {
			prevPos := f.testLog.body[i]
			f.testLog.body[i] = prevWoodPartPos
			prevWoodPartPos = prevPos
		}

		f.updateBoard()

		return f, tickEvery()
	}
	return f, nil
}
