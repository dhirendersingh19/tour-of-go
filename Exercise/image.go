package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBA64Model
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 300, 300)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(255), uint8(0), uint8(0), 255}
}
func main() {
	m := Image{}
	pic.ShowImage(m)
}
