package main

import "fmt"

type vertex struct {
	x, y int
}

func main() {
	v := vertex{1, 2}
	fmt.Println(v)
	v.x = 4
	fmt.Println(v)

	p := &v
	p.y = 1e9

	fmt.Println(v)
}
