package main

import (
	"fmt"
	"math"
)

type someFloat float64

func (s someFloat) abs() float64 {
	if s < 0 {
		return float64(-s)
	}
	return float64(s)
}

func main() {
	f := someFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}
