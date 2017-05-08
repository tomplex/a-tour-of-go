/*
https://tour.golang.org/methods/25
 */

package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type picture struct{
	dx, dy int
	data [][]uint8
}

func (p picture) ColorModel() color.Model {
	return color.RGBAModel
}

func (p picture) Bounds() image.Rectangle {
	return image.Rect(0,0, p.dx, p.dy)
}

func (p picture) At(x, y int) color.Color {
	return color.RGBA{p.data[x][y],p.data[x][y],255,255}
}


func f(x, y int) uint8 {
	return uint8(x ^ y)
}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, 0, dy)

	for y := 1; y <= dy; y++ {
		sl := make([]uint8, 0, dx)
		for x := 1; x <= dx; x++ {
			sl = append(sl, f(x, y))
		}
		s = append(s, sl)
	}
	return s
}

func main() {
	m := picture{100, 100, Pic(100, 100)}
	pic.ShowImage(m)
}