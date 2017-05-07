package main

import "fmt"

type vertex struct {
	x, y float64
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func scaleFunc(v *vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := vertex{3, 4}
	v.scale(2)
	scaleFunc(&v, 10)

	p := &vertex{2, 6}
	p.scale(8)
	scaleFunc(p, 9)

	fmt.Println(v, p)
}
