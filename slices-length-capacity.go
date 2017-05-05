package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// now extend it
	s = s[:6]
	printSlice(s)

	// drop the first two values
	s = s[2:]
	printSlice(s)

	// drop the last two values
	s = s[:len(s)-2]
	printSlice(s)

	// now extend it again
	s = s[:4]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
