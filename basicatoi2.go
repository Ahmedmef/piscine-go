package piscine

func BasicAtoi2(s string) int {
	res := 0
	for _, x := range s {
		if x < '0' || x > '9' {
			return 0
		}
		res = res*10 + int(x-'0')
	}
	return res
}
