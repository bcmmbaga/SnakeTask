package main

import (
	"flag"

	"github.com/bcmmbaga/snaketask"
)

var (
	width  = flag.Int("width", 80, "game board width")
	height = flag.Int("height", 30, "game board height")
)

func main() {
	flag.Parse()

	snaketask.Start(*width, *height)

}
