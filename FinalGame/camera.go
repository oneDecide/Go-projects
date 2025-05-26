package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Camera2DWrapper follows the player
type Camera2DWrapper struct {
	Camera rl.Camera2D
}

func NewCamera2DWrapper() *Camera2DWrapper {
	return &Camera2DWrapper{Camera: rl.Camera2D{
		Target: rl.Vector2{X: 0, Y: 0},
		Offset: rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2},
		Zoom:   1,
	}}
}

func (c *Camera2DWrapper) Update(target rl.Vector2) {
	c.Camera.Target = target
}
