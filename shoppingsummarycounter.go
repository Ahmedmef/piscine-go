package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	b := make(map[string]int)
	c := 0
	d := 0
	for ; d < len(str); d++ {
		if str[d] == ' ' || str[d] == '\t' {
			if c < d {
				word := str[c:d]
				b[word]++
			}
			c = d + 1
		}
	}
	if c < d {
		word := str[c:d]
		b[word]++
	}
	return b
}
