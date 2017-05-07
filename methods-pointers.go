package main

import (
	"fmt"
	"math"
)

type point interface {
	abs() float64
	scale(float64)
}

type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func measure(p point) {
	fmt.Println(p.abs())
}

func main() {
	v := vertex{10, 14}
	measure(&v)

	v.scale(3)
	measure(&v)
}
