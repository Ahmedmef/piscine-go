package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0:]
	ox := len(arg)
	for i := 1; i < ox; i++ {
		for _, v := range arg[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
