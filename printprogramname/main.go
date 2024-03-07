package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	for _, x := range args[2:] {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}
