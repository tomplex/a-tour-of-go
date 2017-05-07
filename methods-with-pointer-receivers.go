/*
https://tour.golang.org/methods/8

!!!
In general, all methods on a type should have either value or pointer receivers, but not a mixture of both.
!!!
 */

package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &vertex{5, 16}
	fmt.Printf("Before scaling: %+v, abs: %v\n", v, v.abs())
	v.scale(6)
	fmt.Printf("After scaling: %+v, abs: %v\n", v, v.abs())
}
