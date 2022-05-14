package game

import (
	"sdl-test/src/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func Render(renderer *sdl.Renderer, context *Context) {
	utils.Blit(renderer, context.player.texture, int(context.player.position.X), int(context.player.position.Y))
}
