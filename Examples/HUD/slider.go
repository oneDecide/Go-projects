package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Slider struct {
	X            int32
	Y            int32
	width        int32
	height       int32
	handleWidth  int32
	handleHeight int32
	progress     float32
	colorTheme   *ColorTheme
}

func (s *Slider) SetProgress(newProgress float32) {
	s.progress = newProgress
	if s.progress < 0 {
		s.progress = 0
	}
	if s.progress > 1 {
		s.progress = 1
	}
}

func (pb ProgressBar) DrawSlider() {

	rl.DrawRectangle(pb.X, pb.Y, int32(pb.progress*float32(pb.Width)), pb.Height, pb.colorTheme.accentColor)
}

func NewSlider(newX, newY, newWidth, newHeight, newHWidth, newHHeight int32, newTheme *ColorTheme) Slider {
	s := Slider{X: newX, Y: newY, width: newWidth, height: newHeight, handleWidth: newHWidth, handleHeight: newHHeight}
	s.colorTheme = newTheme
	s.progress = 0
	return s
}

func (s Slider) DrawSlider() {
	leftX := s.X
	rightX := s.X + s.width - s.handleWidth

	handleX := int32(rl.Lerp(float32(leftX), float32(rightX), s.progress))

	rl.DrawRectangle(s.X, s.Y, s.width, s.height, s.colorTheme.baseColor)
	rl.DrawRectangle(handleX, s.Y+(s.height/2)-(s.handleHeight/2), s.handleWidth, s.handleHeight, s.colorTheme.accentColor)
}
