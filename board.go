package snaketask

import (
	"github.com/JoelOtter/termloop"
)

type Board struct {
	*termloop.Rectangle
}

func NewBoard(width, height int) *Board {
	board := new(Board)

	board.Rectangle = termloop.NewRectangle(0, 0, width, height, termloop.ColorWhite)

	return board
}

func (b *Board) Draw(screen *termloop.Screen) {
	if b == nil {
		return
	}

	w, h := screen.Size()
	b.Rectangle.SetSize(w/2, h/2)
	b.Rectangle.SetPosition(w/4, h/4)

	b.Rectangle.Draw(screen)
}
