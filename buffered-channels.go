/*
https://tour.golang.org/concurrency/3
 */

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 30
	ch <- 7563

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
