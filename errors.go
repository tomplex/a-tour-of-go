/*
https://tour.golang.org/methods/19
 */

package main

import (
	"fmt"
	"time"
)

type anError struct {
	when time.Time
	what string
}

func (e  *anError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &anError{
		time.Now(),
		"It didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

