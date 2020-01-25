package main

import (
	"github.com/eriknikulski/2048/game"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(game.Run)
}
