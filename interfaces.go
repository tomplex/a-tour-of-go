package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
	resize(float64)
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type square struct {
	width float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r *rect) resize(value float64) {
	r.width += value
	r.height += value
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c *circle) resize(value float64) {
	c.radius += value
}

func (s square) area() float64 {
	return s.width * s.width
}

func (s square) perim() float64 {
	return 4 * s.width
}

func (s *square) resize(value float64) {
	s.width += value
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 10, height: 4}
	c := circle{radius: 9}
	s := square{width: 10}

	measure(&r)
	r.resize(4)
	measure(&r)

	measure(&c)
	c.resize(10)
	measure(&c)

	measure(&s)
	s.resize(-1)
	measure(&s)
}
