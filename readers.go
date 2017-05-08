/*
https://tour.golang.org/methods/21
 */

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("The Pittsburgh Penguins are the best hockey team in the NHL")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

