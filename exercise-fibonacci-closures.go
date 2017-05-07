package main

/*
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
 */

import "fmt"

// I struggled for a bit to figure out how to implement a fibonacci generator.
// I ended up using this *beautiful* answer: http://stackoverflow.com/a/25491573/4453925
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		defer func() {first, second = second, first+second}()
		return first
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}