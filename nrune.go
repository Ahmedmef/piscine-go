package piscine

func NRune(s string, n int) rune {
	if n < 0 || n > len(s) {
		return 0
	} else if n > 0 && n <= len(s) {
		return []rune(s)[n-1]
	}
	return 0
}
