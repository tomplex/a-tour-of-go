/*
https://tour.golang.org/flowcontrol/8
 */

package main

import (
	"fmt"
)

const Zi = 100.0

func Sqrt(x float64) (z float64) {

	z = Zi

	step := func() float64 {
		return z - (z*z - x) / (2 * z)
	}

	for i := 0; i <= 10; i++ {
		z = step()
	}
	return z
}

func main() {
	fmt.Println(Sqrt())
}
