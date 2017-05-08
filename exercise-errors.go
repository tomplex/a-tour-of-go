/*
https://tour.golang.org/methods/20
 */

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt a negative number: %f", float64(e))
}

const Zi = 100.0

func Sqrt(x float64) (z float64, err error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z = Zi
	step := func() float64 {
		return z - (z*z - x) / (2 * z)
	}

	for i := 0; i <= 10; i++ {
		z = step()
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}