/*
https://tour.golang.org/methods/16
 */

package main

import "fmt"

func do(i interface{}) {
	switch  v := i.(type) {
	case int:
		fmt.Printf("twice %v is %v\n", v, 2*v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello tom!")
	do(false)
}
