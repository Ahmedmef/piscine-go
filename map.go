package piscine

func Map(f func(int) bool, a []int) []bool {
	len := 0
	for l := range a {
		len = l + 1
	}
	bol := make([]bool, len)
	for i := range a {
		bol[i] = f(a[i])
	}
	return bol
}
