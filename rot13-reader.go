/*
https://tour.golang.org/methods/23
 */

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func rot13(b byte) (byte) {
	// take a byte and apply the rot13 cipher
	switch {
		case b >= 110 || (97 > b && b >= 78):
			b -= 13
		case (b < 110 && b >= 97) || (b < 78 && b >= 65):
			b += 13
	}
	return b
}

func (r rot13Reader) Read(b []byte) (int, error) {
	for {
		_, err := r.reader.Read(b)
		if err == io.EOF {
			break
		}
	}

	for i := range b {
		b[i] = rot13(b[i])
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
