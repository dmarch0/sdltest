package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

func GameLoop(window *sdl.Window, renderer *sdl.Renderer) {
	context := SetupContext(renderer)
	running := true

	for running {
		PrepearScene(renderer)
		PollEvents(&running, &context)
		Update(&context)
		Render(renderer, &context)
		PresentScene(renderer)
		sdl.Delay(16)
	}
}
