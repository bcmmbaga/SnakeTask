package snaketask

import (
	"github.com/JoelOtter/termloop"
)

var board *Board

type Board struct {
	*termloop.Rectangle
	coords map[coordinate]int
}

func NewBoard(width, height int) *Board {
	b := new(Board)

	b.Rectangle = termloop.NewRectangle(1, 1, width, height, termloop.ColorWhite)

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
