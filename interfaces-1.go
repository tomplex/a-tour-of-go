/*
https://tour.golang.org/methods/9
 */

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	abs() float64
}

type someFloat float64

func (s someFloat) abs() float64 {
	if s < 0 {
		return float64(-s)
	}
	return float64(s)
}

type vertex struct {
	x, y float64
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main() {
	var a Abser
	f := someFloat(-math.Sqrt2)
	v := vertex{8, 8}

	// the following two objects can be assigned to an Abser because they both implement abs.
	// note, specifically, that because of the pointer receiver for the vertex's abs method,
	// a pointer to the vertext object needs to be used.
	a = f
	a = &v

	// here, v is a vertex, not a *vertex, so this will cause an error.
	a = v

	fmt.Println(a.abs())
}
