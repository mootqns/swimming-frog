package main

import (
	"fmt"
	"time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	BOARD_WIDTH  = 50
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
	logOne     wood
	logTwo	 wood
	logThree	wood
	logFour	 wood
	logFive	 wood
	logSix	 wood
	logSeven	 wood
	logEight	 wood
	logNine	 wood
	logTen	 wood

	width  int
	height int
}

func newFrogGame() frogGame {
	frog := coord{x: (BOARD_WIDTH / 2) + 1, y: BOARD_HEIGHT - 1}

	logOne := wood{
		body: []coord{
			{x: 1, y: 0},
			{x: 2, y: 0},
			{x: 3, y: 0},
		},
		direction: RIGHT,
	}
	
	logTwo := wood{
		body: []coord{
			{x: 67, y: BOARD_HEIGHT - 3},
			{x: 68, y: BOARD_HEIGHT - 3},
			{x: 69, y: BOARD_HEIGHT - 3},
		},
		direction: LEFT,
	}
	
	logThree := wood{
		body: []coord{
			{x: 9, y: 3},
			{x: 10, y: 3},
			{x: 11, y: 3},
		},
		direction: RIGHT,
	}
	
	logFour := wood{
		body: []coord{
			{x: 15, y: BOARD_HEIGHT - 6},
			{x: 16, y: BOARD_HEIGHT - 6},
			{x: 17, y: BOARD_HEIGHT - 6},
		},
		direction: RIGHT,
	}
	
	logFive := wood{
		body: []coord{
			{x: 54, y: 6},
			{x: 55, y: 6},
			{x: 56, y: 6},
		},
		direction: RIGHT,
	}
	
	logSix := wood{
		body: []coord{
			{x: 47, y: BOARD_HEIGHT - 9},
			{x: 48, y: BOARD_HEIGHT - 9},
			{x: 49, y: BOARD_HEIGHT - 9},
		},
		direction: LEFT,
	}
	
	logSeven := wood{
		body: []coord{
			{x: 23, y: 9},
			{x: 24, y: 9},
			{x: 25, y: 9},
		},
		direction: RIGHT,
	}
	
	logEight := wood{
		body: []coord{
			{x: 10, y: BOARD_HEIGHT - 12},
			{x: 11, y: BOARD_HEIGHT - 12},
			{x: 12, y: BOARD_HEIGHT - 12},
		},
		direction: RIGHT,
	}
	
	logNine := wood{
		body: []coord{
			{x: 62, y: 12},
			{x: 63, y: 12},
			{x: 64, y: 12},
		},
		direction: RIGHT,
	}
	
	logTen := wood{
		body: []coord{
			{x: 49, y: BOARD_HEIGHT - 15},
			{x: 50, y: BOARD_HEIGHT - 15},
			{x: 51, y: BOARD_HEIGHT - 15},
		},
		direction: LEFT,
	}

	game := frogGame{
		logOne:     logOne,
		logTwo:    logTwo,
		logThree:	logThree,
		logFour:	logFour,
		logFive:	logFive,
		logSix:	logSix,
		logSeven:	logSeven,
		logEight:	logEight,
		logNine:	logNine,
		logTen:	logTen,
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

			if f.logOne.coordInBody(curCell) || f.logTwo.coordInBody(curCell) || f.logThree.coordInBody(curCell) || f.logFour.coordInBody(curCell) ||
                f.logFive.coordInBody(curCell) || f.logSix.coordInBody(curCell) || f.logSeven.coordInBody(curCell) || f.logEight.coordInBody(curCell) ||
                f.logNine.coordInBody(curCell) || f.logTen.coordInBody(curCell) {
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
			startBorder.Render(menuScreenStyle.Render(menuTextStyle.Render("> frog game"))+
				altTextStyle.Render("\n\npress enter to play")))
	}

	if f.gameOver {
		return lipgloss.Place(f.width, f.height, lipgloss.Center, lipgloss.Center,
			startBorder.Render(menuScreenStyle.Render(menuTextStyle.Render("> game over"))+scoreText+italicsTextStyle.Render("q quit\n")))
	}

	screen := ""

	for i := 0; i < BOARD_HEIGHT; i++ {
		for j := 0; j < BOARD_WIDTH; j++ {
			if i == f.frog.y && j == f.frog.x {
				screen += "🐸"
			} else if f.gameBoard[i][j] == LOG_CELL {
				screen += logStyle.Render(" ")
			} else {
				screen += baseStyle.Render(" ")
			}
		}

		if i != BOARD_HEIGHT-1 {
			screen += "\n"
		}
	}

	if (f.logOne.body[0].x > 70 || f.logOne.body[0].x < -2) {
		f.logOne.body[0].x = 1
		f.logOne.body[1].x = 2
		f.logOne.body[2].x = 3
	}
	
	if (f.logTwo.body[0].x > 70 || f.logTwo.body[0].x < -2) {
		f.logTwo.body[0].x = 67
		f.logTwo.body[1].x = 68
		f.logTwo.body[2].x = 69
	}
	
	if (f.logThree.body[0].x > 70 || f.logThree.body[0].x < -2) {
		f.logThree.body[0].x = 1
		f.logThree.body[1].x = 2
		f.logThree.body[2].x = 3
	}
	
	if (f.logFour.body[0].x > 70 || f.logFour.body[0].x < -2) {
		f.logFour.body[0].x = 67
		f.logFour.body[1].x = 68
		f.logFour.body[2].x = 69
	}
	
	if (f.logFive.body[0].x > 70 || f.logFive.body[0].x < -2) {
		f.logFive.body[0].x = 1
		f.logFive.body[1].x = 2
		f.logFive.body[2].x = 3
	}
	
	if (f.logSix.body[0].x > 70 || f.logSix.body[0].x < -2) {
		f.logSix.body[0].x = 67
		f.logSix.body[1].x = 68
		f.logSix.body[2].x = 69
	}
	
	if (f.logSeven.body[0].x > 70 || f.logSeven.body[0].x < -2) {
		f.logSeven.body[0].x = 1
		f.logSeven.body[1].x = 2
		f.logSeven.body[2].x = 3
	}
	
	if (f.logEight.body[0].x > 70 || f.logEight.body[0].x < -2) {
		f.logEight.body[0].x = 67
		f.logEight.body[1].x = 68
		f.logEight.body[2].x = 69
	}
	
	if (f.logNine.body[0].x > 70 || f.logNine.body[0].x < -2) {
		f.logNine.body[0].x = 1
		f.logNine.body[1].x = 2
		f.logNine.body[2].x = 3
	}
	
	if (f.logTen.body[0].x > 70 || f.logTen.body[0].x < -2) {
		f.logTen.body[0].x = 67
		f.logTen.body[1].x = 68
		f.logTen.body[2].x = 69
	}

	helpMsg := "arrows move | q quit\n"

	return lipgloss.Place(f.width, f.height, lipgloss.Center, lipgloss.Center, boardStyle.Render(screen)+scoreText+italicsTextStyle.Render(helpMsg))
}

func (f frogGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if f.frog.y == 1 {
		f.score += 10
		f.frog.y += BOARD_HEIGHT - 1
	}

	for i := 0; i < BOARD_HEIGHT; i++ {
		for j := 0; j < BOARD_WIDTH; j++ {
			if f.gameBoard[i][j] == LOG_CELL && i == f.frog.y && j == f.frog.x {
				f.gameOver = true
			}
		}
	}

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
		case "c":
			f.gameOver = true
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

		prevWoodOnePartPos := f.logOne.body[0]

		switch f.logOne.direction {
		case RIGHT:
			f.logOne.body[0].x++
		case LEFT:
			f.logOne.body[0].x--
		}

		for i := 1; i < len(f.logOne.body); i++ {
			prevPosOne := f.logOne.body[i]
			f.logOne.body[i] = prevWoodOnePartPos
			prevWoodOnePartPos = prevPosOne
		}

		prevWoodTwoPartPos := f.logTwo.body[0]

		switch f.logTwo.direction {
		case RIGHT:
			f.logTwo.body[0].x++
		case LEFT:
			f.logTwo.body[0].x--
		}

		for i := 1; i < len(f.logTwo.body); i++ {
			prevPosTwo := f.logTwo.body[i]
			f.logTwo.body[i] = prevWoodTwoPartPos
			prevWoodTwoPartPos = prevPosTwo
		}

		prevWoodThreePartPos := f.logThree.body[0]

		switch f.logThree.direction {
		case RIGHT:
		f.logThree.body[0].x++
		case LEFT:
		f.logThree.body[0].x--
		}

		for i := 1; i < len(f.logThree.body); i++ {
		prevPosThree := f.logThree.body[i]
		f.logThree.body[i] = prevWoodThreePartPos
		prevWoodThreePartPos = prevPosThree
		}

		prevWoodFourPartPos := f.logFour.body[0]

		switch f.logFour.direction {
		case RIGHT:
		f.logFour.body[0].x++
		case LEFT:
		f.logFour.body[0].x--
		}

		for i := 1; i < len(f.logFour.body); i++ {
		prevPosFour := f.logFour.body[i]
		f.logFour.body[i] = prevWoodFourPartPos
		prevWoodFourPartPos = prevPosFour
		}

		prevWoodFivePartPos := f.logFive.body[0]

		switch f.logFive.direction {
		case RIGHT:
		f.logFive.body[0].x++
		case LEFT:
		f.logFive.body[0].x--
		}

		for i := 1; i < len(f.logFive.body); i++ {
		prevPosFive := f.logFive.body[i]
		f.logFive.body[i] = prevWoodFivePartPos
		prevWoodFivePartPos = prevPosFive
		}

		prevWoodSixPartPos := f.logSix.body[0]

		switch f.logSix.direction {
		case RIGHT:
		f.logSix.body[0].x++
		case LEFT:
		f.logSix.body[0].x--
		}

		for i := 1; i < len(f.logSix.body); i++ {
		prevPosSix := f.logSix.body[i]
		f.logSix.body[i] = prevWoodSixPartPos
		prevWoodSixPartPos = prevPosSix
		}

		prevWoodSevenPartPos := f.logSeven.body[0]

		switch f.logSeven.direction {
		case RIGHT:
		f.logSeven.body[0].x++
		case LEFT:
		f.logSeven.body[0].x--
		}

		for i := 1; i < len(f.logSeven.body); i++ {
		prevPosSeven := f.logSeven.body[i]
		f.logSeven.body[i] = prevWoodSevenPartPos
		prevWoodSevenPartPos = prevPosSeven
		}

		prevWoodEightPartPos := f.logEight.body[0]

		switch f.logEight.direction {
		case RIGHT:
		f.logEight.body[0].x++
		case LEFT:
		f.logEight.body[0].x--
		}

		for i := 1; i < len(f.logEight.body); i++ {
		prevPosEight := f.logEight.body[i]
		f.logEight.body[i] = prevWoodEightPartPos
		prevWoodEightPartPos = prevPosEight
		}

		prevWoodNinePartPos := f.logNine.body[0]

		switch f.logNine.direction {
		case RIGHT:
		f.logNine.body[0].x++
		case LEFT:
		f.logNine.body[0].x--
		}

		for i := 1; i < len(f.logNine.body); i++ {
		prevPosNine := f.logNine.body[i]
		f.logNine.body[i] = prevWoodNinePartPos
		prevWoodNinePartPos = prevPosNine
		}

		prevWoodTenPartPos := f.logTen.body[0]

		switch f.logTen.direction {
		case RIGHT:
		f.logTen.body[0].x++
		case LEFT:
		f.logTen.body[0].x--
		}

		for i := 1; i < len(f.logTen.body); i++ {
		prevPosTen := f.logTen.body[i]
		f.logTen.body[i] = prevWoodTenPartPos
		prevWoodTenPartPos = prevPosTen
		}

		f.updateBoard()

		return f, tickEvery()
	}
	return f, nil
}
