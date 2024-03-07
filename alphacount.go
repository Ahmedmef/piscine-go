package piscine

func AlphaCount(s string) int {
	a := 0
	for _, i := range s {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
			a++
		}
	}
	return a
}
