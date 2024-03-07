package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
}
