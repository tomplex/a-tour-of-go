package main

import "fmt"

type vertex struct {
	x, y int
}

var (
	v1 = vertex{1, 2}  // has type vertex
	v2 = vertex{x: 1}  // implies y = 0
	v3 = vertex{}      // x, y = 0
	pv = &vertex{3, 4} // has type *vertex
)

func main() {
	pv.x = 9384
	v4 := *pv
	fmt.Println(v1, v2, v3, v4)
}
