package main

import (
	"fmt"

	"github.com/JoelOtter/termloop"
	"github.com/bcmmbaga/snaketask"
)

func main() {
	fmt.Println("start from here")

	g := termloop.NewGame()

	g.Screen().AddEntity(snaketask.NewBoard(100, 100))
	g.Screen().AddEntity(snaketask.NewSnake())
	g.Screen().SetFps(10)
	g.Start()
}
