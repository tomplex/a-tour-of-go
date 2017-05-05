package main

/*
 *Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
 *
 *The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
 *
 *(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
 *
 *(Use uint8(intValue) to convert between types.)
 */

import "golang.org/x/tour/pic"
import "fmt"

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
	fmt.Println(Pic(10, 5))
	pic.Show(Pic)
}
