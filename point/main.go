package main

import (
	"github.com/01-edu/z01"
)

var z = "x = 42, y = 21"

func main() {
	for _, e := range z {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
