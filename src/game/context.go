package game

import (
	"sdl-test/src/utils"

	"github.com/veandco/go-sdl2/sdl"
)

type Context struct {
	player Player
}

func SetupContext(renderer *sdl.Renderer) Context {
	position := Vector{0, 0}
	velocity := Vector{0, 0}
	texture := utils.LoadTexture(renderer, "resources/player.png")
	player := Player{position, velocity, texture}
	context := Context{player}
	return context
}
