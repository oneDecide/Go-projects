package main

import rl "github.com/gen2brain/raylib-go/raylib"

type ProgressBar struct {
	X          int32
	Y          int32
	Width      int32
	Height     int32
	progress   float32
	colorTheme *ColorTheme
}

func (pb *ProgressBar) SetProgress(newProgress float32) {
	pb.progress = rl.Clamp(newProgress, 0, 1)
}

func (pb ProgressBar) DrawBar() {
	rl.DrawRectangle(pb.X, pb.Y, pb.Width, pb.Height, pb.colorTheme.baseColor)
	rl.DrawRectangle(pb.X, pb.Y, int32(pb.progress*float32(pb.Width)), pb.Height, pb.colorTheme.accentColor)
}

func NewProgressBar(newX, newY, newWidth, newHeight int32, newTheme *ColorTheme) ProgressBar {
	pb := ProgressBar{X: newX, Y: newY, Width: newWidth, Height: newHeight}
	pb.colorTheme = newTheme
	pb.progress = 0
	return pb
}
