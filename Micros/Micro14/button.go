package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ClickBound interface {
	OnClick()
	//I can put return types here
}

type Button struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
	ColorTheme
	text         string
	textSize     int32
	mouseOver    bool
	mouseClicked bool
	onClickFuncs []func()

	pressed bool
}

func NewButton(x, y, width, height int32, newTheme ColorTheme) Button {
	nb := Button{X: x, Y: y, Width: width, Height: height, ColorTheme: newTheme}
	nb.onClickFuncs = make([]func(), 0)
	nb.pressed = false
	nb.mouseClicked = false
	return nb
}

func (b *Button) SetText(newText string, newTextSize int32) {
	b.text = newText
	b.textSize = newTextSize
}

func (b *Button) DrawButton() {
	currentColor := b.baseColor

	if b.mouseOver {
		currentColor = rl.NewColor(b.baseColor.R/2, b.baseColor.G/2, b.baseColor.B/2, 255)
	}
	if b.pressed {
		currentColor = b.accentColor
	}

	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, currentColor)

	textWidth := rl.MeasureText(b.text, int32(b.textSize))
	rl.DrawText(
		b.text,
		b.X+(b.Width/2)-(textWidth/2),
		b.Y+(b.Height/2)-(b.textSize/2),
		int32(b.textSize),
		b.textColor,
	)
}

func (b *Button) CenterButtonX() {
	b.X = int32(rl.GetScreenWidth()/2) - b.Width/2
}

func (b *Button) CenterButtonY() {
	b.Y = int32(rl.GetScreenHeight()/2) - b.Height/2
}

func (b *Button) CenterButton() {
	b.CenterButtonX()
	b.CenterButtonY()
}

func (b *Button) CheckMouseOver() {
	mousePos := rl.GetMousePosition()
	b.mouseOver = false
	if int32(mousePos.X) < b.X || int32(mousePos.X) > b.X+b.Width {
		return
	}
	if int32(mousePos.Y) < b.Y || int32(mousePos.Y) > b.Y+b.Height {
		return
	}
	b.mouseOver = true
}

func (b *Button) CheckMouseClick() {
	if b.mouseOver && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		b.pressed = !b.pressed
		fmt.Println("Button toggled:", b.pressed)
		for _, onClickFunc := range b.onClickFuncs {
			onClickFunc()
		}
	}
}

func (b *Button) AddOnClickFunc(newFunc func()) {
	b.onClickFuncs = append(b.onClickFuncs, newFunc)
}

func (b *Button) UpdateButton() {

	b.CheckMouseOver()

	b.DrawButton()

	b.CheckMouseClick()

}
