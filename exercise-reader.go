/*
https://tour.golang.org/methods/22
 */

package main

import (
	"golang.org/x/tour/reader"
)

type AReader struct {}

func (r AReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(AReader{})
}