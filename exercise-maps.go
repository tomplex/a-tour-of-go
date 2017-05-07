package main

/*
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
 */

import (
	"golang.org/x/tour/wc"
	"strings"
)

func wordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, w := range strings.Fields(s) {
		count[w] += 1
	}
	return count
}

func main () {
	wc.Test(wordCount)
}