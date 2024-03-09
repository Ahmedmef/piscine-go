package piscine

func BasicAtoi(s string) int {
	res := 0
	for _, x := range s {
		res = res*10 + int(x-'0')
	}
	return res
}
