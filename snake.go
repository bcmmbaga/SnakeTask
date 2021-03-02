package snaketask

import (
	"fmt"
	"os"

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
	score     int
	scoreText *termloop.Text

	// length of the snake
	len int

	direction direction
}

// NewSnake return new Snake with default coordinate and length.
func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = termloop.NewEntity(1, 1, 1, 1)
	snake.coords = []coordinate{
		{x: 1, y: 4},
		{x: 2, y: 4},
		{x: 3, y: 4},
	}

	snake.len = len(snake.coords)
	snake.direction = right
	snake.scoreText = termloop.NewText(1, 1, "Score: 0", termloop.ColorBlack, termloop.ColorWhite)

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

	s.scoreText.Draw(screen)

	if s.selfCollissionCheck() || s.BorderCollisionCheck() {
		fmt.Printf("Your score: %d", s.score)
		os.Exit(1)
	}

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

func (s *Snake) Collide(collision termloop.Physical) {
	switch collision.(type) {
	case *Food:
		s.len += 3
		s.score++
		s.scoreText.SetText(fmt.Sprintf("Score: %d", s.score))
	}
}

func (s *Snake) selfCollissionCheck() bool {
	for i := 0; i < len(s.coords)-1; i++ {
		if s.coords[len(s.coords)-1] == s.coords[i] {
			return true
		}
	}
	return false
}

func (s *Snake) BorderCollisionCheck() bool {
	_, ok := board.coords[s.coords[len(s.coords)-1]]
	return !ok
}
