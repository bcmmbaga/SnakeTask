package snaketask

import (
	"github.com/JoelOtter/termloop"
)

type direction int

const (
	right direction = iota
	left
	up
	down
)

// Snake ...
type Snake struct {
	*termloop.Entity

	// snake coordinate
	coords []coordinate

	// number of food units eaten
	score int

	// length of the snake
	len int

	direction direction
}

// NewSnake return new Snake with default coordinate and length.
func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = termloop.NewEntity(5, 5, 1, 1)
	x, y := 282/2, 69/2
	snake.coords = []coordinate{
		{x: x + 1, y: y},
		{x: x + 2, y: y},
		{x: x + 3, y: y},
	}

	snake.len = len(snake.coords)
	snake.direction = right

	return snake
}

func (s *Snake) Draw(screen *termloop.Screen) {
	if s == nil {
		return
	}
	headCoord := s.coords[len(s.coords)-1]

	switch s.direction {
	case right:
		headCoord.x++
	case left:
		headCoord.x--
	case up:
		headCoord.y--
	case down:
		headCoord.y++
	}

	if s.len > len(s.coords) {
		s.coords = append(s.coords, headCoord)
	} else {
		s.coords = append(s.coords[1:], headCoord)
	}

	s.SetPosition(headCoord.x, headCoord.y)

	for _, coord := range s.coords {
		screen.RenderCell(
			coord.x, coord.y, &termloop.Cell{
				Fg: termloop.ColorBlue,
				Ch: '■',
			},
		)
	}
}
func (s *Snake) Tick(event termloop.Event) {
	// snake can move to different direction but not from direction it came (no reverse).
	if event.Type == termloop.EventKey {
		switch event.Key {
		case termloop.KeyArrowRight:
			if s.direction != left {
				s.direction = right
			}
		case termloop.KeyArrowLeft:
			if s.direction != right {
				s.direction = left
			}
		case termloop.KeyArrowUp:
			if s.direction != down {
				s.direction = up
			}
		case termloop.KeyArrowDown:
			if s.direction != up {
				s.direction = down
			}
		}
	}
}
