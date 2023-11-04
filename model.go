package main

import (
	"fmt"
	"math/rand"
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
	SNAKE_CELL
	PELLET_CELL
)

type TickMsg time.Time

type coord struct {
	x int
	y int
}

// this is the model used by bubbletea
type snakeGame struct {
	gameBoard [][]int
	frog      coord
	pellet    coord
	score     int
	gameOver  bool

	width  int
	height int
}

func newSnakeGame() snakeGame {
	frog := coord{x: (BOARD_WIDTH / 2) + 1, y: BOARD_HEIGHT / 2}

	game := snakeGame{
		frog:     frog,
		gameOver: false,
	}

	game.spawnPellet()
	game.updateBoard()

	return game
}

func (s *snakeGame) updateBoard() {
	gameBoard := [][]int{}

	for i := 0; i < BOARD_HEIGHT; i++ {
		row := []int{}
		for j := 0; j < BOARD_WIDTH; j++ {
			cellType := BLANK_CELL
			curCell := coord{j, i}

			if curCell == s.pellet {
				cellType = PELLET_CELL
			}

			row = append(row, cellType)
		}

		gameBoard = append(gameBoard, row)
	}

	s.gameBoard = gameBoard
}

// place pellet at random position that doesn't overlap with snake
func (s *snakeGame) spawnPellet() {
	pellet := coord{
		x: rand.Intn(BOARD_WIDTH),
		y: rand.Intn(BOARD_HEIGHT),
	}

	s.pellet = pellet
}

func tickEvery() tea.Cmd {
	return tea.Every(time.Millisecond*100, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (s snakeGame) Init() tea.Cmd {
	return tickEvery()
}

func (s snakeGame) View() string {
	scoreLabel := scoreStyle.Render("score")
	scoreText := fmt.Sprintf("\n%s: %d\n\n", scoreLabel, s.score)

	if s.gameOver {
		return gameOverStyle.Render(gameOverText) + scoreText + "ctrl+c to quit\n"
	}

	screen := ""

	for i := 0; i < BOARD_HEIGHT; i++ {
		for j := 0; j < BOARD_WIDTH; j++ {
			if i == s.frog.y && j == s.frog.x {
				screen += "🐸"
			} else {
				screen += " "
			}
		}

		if i != BOARD_HEIGHT-1 {
			screen += "\n"
		}
	}

	helpMsg := "arrow keys to move\nesc to quit\n"

	return boardStyle.Render(screen) + scoreText + helpMsg
}

func (s snakeGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit

		// when moving snake, don't allow movement in opposite direction
		// this is done by checking if snake body part directly behind head
		// is in the direction player is trying to move
		case "q":
			return s, tea.Quit
		case "up":
			if s.frog.y > 0 {
				s.frog.y--
			}
		case "right":
			if s.frog.x < BOARD_WIDTH-1 {
				s.frog.x++
			}
		case "down":
			if s.frog.y < BOARD_HEIGHT-1 {
				s.frog.y++
			}
		case "left":
			if s.frog.x > 0 {
				s.frog.x--
			}
		}
		// case TickMsg:
		// 	// move snake head in direction
		// 	prevSnakePartPos := s.snake.body[0]

		// 	if s.snake.body[0].x < 0 || s.snake.body[0].x >= BOARD_WIDTH ||
		// 		s.snake.body[0].y < 0 || s.snake.body[0].y >= BOARD_HEIGHT ||
		// 		s.snake.isHeadCollidingWithBody() {

		// 		s.gameOver = true
		// 		break
		// 	}

		// 	atePellet := s.snake.body[0] == s.pellet

		// 	/*
		// 		move the rest of the snake
		// 		temporarily save position of current part as prevPos
		// 		move part to prevSnakePartPos
		// 		set prevSnakePartPos to prevPos for the next iteration
		// 	*/
		// 	for i := 1; i < len(s.snake.body); i++ {
		// 		prevPos := s.snake.body[i]
		// 		s.snake.body[i] = prevSnakePartPos
		// 		prevSnakePartPos = prevPos
		// 	}

		// 	if atePellet {
		// 		s.snake.body = append(s.snake.body, prevSnakePartPos)
		// 		s.spawnPellet()
		// 		s.score++
		// 	}

		//s.updateBoard()

		// return s, tickEvery()
	}
	return s, nil
}
