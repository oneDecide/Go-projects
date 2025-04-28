package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(720)
	rl.InitWindow(screenWidth, screenHeight, "Grass Shader Example")
	defer rl.CloseWindow()

	rl.DisableCursor()

	model := rl.LoadModel("spikeygrass.obj")
	defer rl.UnloadModel(model)

	grassShader := rl.LoadShader("grass.vs", "grass.fs")
	defer rl.UnloadShader(grassShader)

	baseColorLoc := rl.GetShaderLocation(grassShader, "baseColor")
	topColorLoc := rl.GetShaderLocation(grassShader, "topColor")
	timeLoc := rl.GetShaderLocation(grassShader, "time")

	baseColor := []float32{0.05, 0.25, 0.25, 1.0}
	topColor := []float32{0.1, 0.5, 0.5, 1.0}
	time := []float32{0.0}

	rl.SetShaderValue(grassShader, baseColorLoc, baseColor, rl.ShaderUniformVec4)
	rl.SetShaderValue(grassShader, topColorLoc, topColor, rl.ShaderUniformVec4)

	model.GetMaterials()[0].Shader = grassShader

	camera := rl.Camera{
		Position:   rl.NewVector3(0.0, 5, 10.0),
		Target:     rl.NewVector3(0.0, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		time[0] += rl.GetFrameTime()
		rl.SetShaderValue(grassShader, timeLoc, time, rl.ShaderUniformFloat)

		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)
		//rl.DrawGrid(10, 10.0)

		//rl.DrawModel(model, rl.Vector3Zero(), 2.0, rl.White)
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				rl.DrawModel(model, rl.NewVector3(2*float32(i), 0, 3*float32(j)), 1.0, rl.White)
			}
		}

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
