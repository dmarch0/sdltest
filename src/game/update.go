package game

func Update(context *Context) {
	context.player.position.Add(&context.player.velocity)
}
