package main

import "fmt"

type vertex struct {
	lat, lon float64
}

var m = map[string]vertex{
	"Bell Labs" : vertex{40.68433, -74.39967},
	"Google" : vertex{37.42202, -122.08408},
	// when the value is a type, you can omit the declaration:
	"White River Junction" : {43.649145, -72.321124},
}

func main() {
	fmt.Println(m)
}