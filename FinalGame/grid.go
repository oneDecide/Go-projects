package main

import rl "github.com/gen2brain/raylib-go/raylib"

// DrawGrid renders infinite grid around camera
func DrawGrid(cam *Camera2DWrapper) {
	bounds := float32(5000)
	ox := float32(int(cam.Camera.Target.X) % int(GridSpacing))
	oy := float32(int(cam.Camera.Target.Y) % int(GridSpacing))
	for x := -bounds; x <= bounds; x += GridSpacing {
		rl.DrawLine(int32(x-ox), -int32(bounds), int32(x-ox), int32(bounds), rl.LightGray)
	}
	for y := -bounds; y <= bounds; y += GridSpacing {
		rl.DrawLine(-int32(bounds), int32(y-oy), int32(bounds), int32(y-oy), rl.LightGray)
	}
}
