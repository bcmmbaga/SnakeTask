package snaketask

import (
	"github.com/JoelOtter/termloop"
)

var board *Board

type Board struct {
	*termloop.Rectangle
}

func NewBoard(width, height int) *Board {
	b := new(Board)

	b.Rectangle = termloop.NewRectangle(1, 1, width, height, termloop.ColorWhite)

	board = b

	return board
}

func (b *Board) Draw(screen *termloop.Screen) {
	if b == nil {
		return
	}

	b.Rectangle.Draw(screen)
}
