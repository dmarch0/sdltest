package game

import "github.com/veandco/go-sdl2/sdl"

func PrepearScene(renderer *sdl.Renderer) {
	renderer.SetDrawColor(96, 128, 255, 255)
	renderer.Clear()
}
