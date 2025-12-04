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
	pb.progress = newProgress
	if pb.progress < 0 {
		pb.progress = 0
	}
	if pb.progress > 0.97 {
		pb.progress = 0.97
	}
}

func (pb ProgressBar) DrawBar() {
	rl.DrawRectangle(pb.X, pb.Y, pb.Width, pb.Height, pb.colorTheme.baseColor)
	rl.DrawRectangle(pb.X+5, pb.Y+10, int32(pb.progress*float32(pb.Width)), pb.Height-15, pb.colorTheme.accentColor)
}

func NewProgressBar(newX, newY, newWidth, newHeight int32, newTheme *ColorTheme) ProgressBar {
	pb := ProgressBar{X: newX, Y: newY, Width: newWidth, Height: newHeight}
	pb.colorTheme = newTheme
	pb.progress = 0
	return pb
}
