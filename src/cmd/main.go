package main

import (
	"sdl-test/src/game"
	"sdl-test/src/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, renderer := utils.InitSdl()

	game.GameLoop(window, renderer)

	sdl.Quit()
	window.Destroy()
}
