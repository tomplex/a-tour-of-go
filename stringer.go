/*
https://tour.golang.org/methods/17
 */

package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	s := person{"Sidney Crosby", 29}
	e := person{"Marc-Andre Fleury", 32}
	fmt.Println(s, e)
}