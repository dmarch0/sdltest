package game

import "github.com/veandco/go-sdl2/sdl"

func PollEvents(running *bool, context *Context) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			*running = false
			break
		case *sdl.KeyboardEvent:
			HandleKeyboardInput(event.(*sdl.KeyboardEvent), context)
			break
		}
	}
}

func HandleKeyboardInput(event *sdl.KeyboardEvent, context *Context) {
	switch event.Type {
	case sdl.KEYDOWN:
		HandleKeyDown(event, context)
		break
	case sdl.KEYUP:
		HandleKeyUp(event, context)
		break
	}
}

func HandleKeyDown(event *sdl.KeyboardEvent, context *Context) {
	if event.Repeat == 0 {
		switch event.Keysym.Scancode {
		case sdl.SCANCODE_UP:
			context.player.velocity = Vector{0, -1}
			break
		case sdl.SCANCODE_DOWN:
			context.player.velocity = Vector{0, 1}
			break
		case sdl.SCANCODE_LEFT:
			context.player.velocity = Vector{-1, 0}
			break
		case sdl.SCANCODE_RIGHT:
			context.player.velocity = Vector{1, 0}
			break
		}

	}
}

func HandleKeyUp(event *sdl.KeyboardEvent, context *Context) {
	if event.Repeat == 0 {
		switch event.Keysym.Scancode {
		case sdl.SCANCODE_UP, sdl.SCANCODE_DOWN, sdl.SCANCODE_LEFT, sdl.SCANCODE_RIGHT:
			context.player.velocity = Vector{0, 0}
		}
	}
}
