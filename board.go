package snaketask

import (
	"github.com/JoelOtter/termloop"
)

var board *Board

type Board struct {
	*termloop.Rectangle
	coords map[coordinate]int
}

func newBoard(width, height int) *Board {
	b := new(Board)

	b.Rectangle = termloop.NewRectangle(1, 1, width, height, termloop.ColorWhite)

	// add all position of the boarders into map
	b.coords = make(map[coordinate]int)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			b.coords[coordinate{x: i, y: j}] = 1
		}
	}

	board = b

	return board
}

func (b *Board) Draw(screen *termloop.Screen) {
	if b == nil {
		return
	}

	b.Rectangle.Draw(screen)
}

// Start render base playground and start game.
func Start(width, height int) {
	baseLevel := termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlue,
	})

	var (
		board = newBoard(width, height)
		snake = newSnake()
		food  = newFood()
	)
	baseLevel.AddEntity(board)
	baseLevel.AddEntity(snake)
	baseLevel.AddEntity(food)

	g := termloop.NewGame()

	g.Screen().SetLevel(baseLevel)
	g.Screen().SetFps(10)
	g.Start()
}
