/*
https://tour.golang.org/methods/13
 */

package main

import "fmt"

type I interface {
	M()
}

// this file panics!
func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}