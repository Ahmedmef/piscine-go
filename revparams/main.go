package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	f := os.Args
	res := ""
	for _, a := range f[1:] {
		res = a + "\n" + res
	}
	for _, b := range res {
		z01.PrintRune(b)
	}
}
