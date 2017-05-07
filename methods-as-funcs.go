package main

import (
	"math"
	"fmt"
)

type vertex struct {
	x, y float64
}

func abs(v vertex) float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main() {
	// this is functionally the same as adding the method receiver to the abs function.
	// but, using the function on the vertex object (instead of as a method) will feel
	// less familiar to people like me who are coming from mostly object-oriented languages.
	v := vertex{4, 10}
	fmt.Println(abs(v))
}
