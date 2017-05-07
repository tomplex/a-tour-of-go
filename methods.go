package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main() {
	v := vertex{4, 10}
	fmt.Println(v.abs())
}

