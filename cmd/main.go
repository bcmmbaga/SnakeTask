package main

import (
	"github.com/JoelOtter/termloop"
	"github.com/bcmmbaga/snaketask"
)

func main() {
	baseLevel := termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorDefault,
	})

	var (
		board = snaketask.NewBoard(80, 30)
		snake = snaketask.NewSnake()
		food  = snaketask.NewFood()
	)
	baseLevel.AddEntity(board)
	baseLevel.AddEntity(snake)
	baseLevel.AddEntity(food)

	g := termloop.NewGame()

	g.Screen().SetLevel(baseLevel)
	g.Screen().SetFps(10)
	g.Start()
}
