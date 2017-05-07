/*
https://tour.golang.org/methods/10
 */

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello tom!"}
	i.M()
}