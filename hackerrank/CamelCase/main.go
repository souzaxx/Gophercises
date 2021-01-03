package main

import (
	"fmt"
	"unicode"
)

func main() {
	var words int
	s := "oneTwoThree"

	for _, l := range s {
		if unicode.IsUpper(l) {
			words++
		}
	}
	fmt.Print(words)
}
