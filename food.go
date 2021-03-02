package snaketask

import (
	"math/rand"

	"github.com/JoelOtter/termloop"
)

type Food struct {
	*termloop.Entity
	coord coordinate
}

func NewFood() *Food {
	f := new(Food)
	f.Entity = termloop.NewEntity(1, 1, 1, 1)
	f.shiftToNewPosition()

	return f
}

func (f *Food) Draw(screen *termloop.Screen) {
	if f == nil {
		return
	}

	screen.RenderCell(f.coord.x, f.coord.y, &termloop.Cell{
		Fg: termloop.ColorGreen,
		Ch: '■',
	})
}

func (f *Food) Collide(collision termloop.Physical) {
	switch collision.(type) {
	case *Snake:
		f.shiftToNewPosition()
	}
}

// shiftToNewPosition moves food unit to next random position within board.
func (f *Food) shiftToNewPosition() {
	w, h := board.Rectangle.Size()

	x := rand.Intn((w-1)-1) + 1
	y := rand.Intn((h-1)-1) + 1

	f.coord = coordinate{x: x, y: y}

	f.SetPosition(x, y)
}
