package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	position Vector
	velocity Vector
	texture  *sdl.Texture
}
